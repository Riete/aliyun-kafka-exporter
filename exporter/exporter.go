package exporter

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alikafka"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

var sleep = false

func wakeup() {
	lock := sync.RWMutex{}
	lock.Lock()
	time.Sleep(time.Minute * time.Duration(1))
	sleep = false
	lock.Unlock()
}

func (k *KafkaExporter) NewClient() {
	client, err := cms.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	k.client = client
}

func (k *KafkaExporter) GetInstance() {
	client, err := alikafka.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	request := alikafka.CreateGetInstanceListRequest()
	response, err := client.GetInstanceList(request)
	if err != nil {
		panic(err)
	}
	instances := make(map[string]string)
	for _, v := range response.InstanceList.InstanceVO {
		instances[v.InstanceId] = v.Name

	}
	k.instances = instances
}

func (k *KafkaExporter) GetMetricMeta() {
	request := cms.CreateDescribeMetricMetaListRequest()
	request.Namespace = PROJECT
	request.PageSize = requests.NewInteger(100)
	response, err := k.client.DescribeMetricMetaList(request)
	if err != nil {
		panic(err)
	}
	for _, v := range response.Resources.Resource {
		k.metricMeta = append(k.metricMeta, v.MetricName)
	}
}

func (k *KafkaExporter) GetMetric(metricName string) {
	var dimensions []map[string]string
	for k := range k.instances {
		d := map[string]string{"instanceId": k}
		dimensions = append(dimensions, d)
	}
	dimension, err := json.Marshal(dimensions)
	if err != nil {
		log.Println(err)
	}
	request := cms.CreateDescribeMetricLastRequest()
	request.Namespace = PROJECT
	request.MetricName = metricName
	request.Dimensions = string(dimension)
	request.Period = "120"
	response, err := k.client.DescribeMetricLast(request)
	if err != nil {
		log.Println(err)
	}
	k.DataPoints = nil
	err = json.Unmarshal([]byte(response.Datapoints), &k.DataPoints)
	if err != nil {
		log.Println(err)
	}
}

func (k *KafkaExporter) InitGauge() {
	k.NewClient()
	k.GetInstance()
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			k.GetInstance()
		}
	}()
	k.GetMetricMeta()
	k.metrics = map[string]*prometheus.GaugeVec{}
	for _, v := range k.metricMeta {
		k.metrics[v] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "aliyun_kafka",
			Name:      strings.ToLower(v),
		}, []string{"instance_id", "instance_name", "consumer_group", "topic"})
	}
}

func (k *KafkaExporter) Describe(ch chan<- *prometheus.Desc) {
	for _, v := range k.metrics {
		v.Describe(ch)
	}
}

func (k *KafkaExporter) Collect(ch chan<- prometheus.Metric) {
	if !sleep {
		sleep = true
		go wakeup()
		for _, v := range k.metricMeta {
			k.GetMetric(v)
			for _, d := range k.DataPoints {
				value := d.Value
				if v == "instance_disk_capacity" {
					value = d.Maximum
				}
				k.metrics[v].With(prometheus.Labels{
					"instance_id":    d.InstanceId,
					"instance_name":  k.instances[d.InstanceId],
					"consumer_group": d.ConsumerGroup,
					"topic":          d.Topic,
				}).Set(value)
			}
			time.Sleep(34 * time.Millisecond)
		}
	}
	for _, m := range k.metrics {
		m.Collect(ch)
	}
}
