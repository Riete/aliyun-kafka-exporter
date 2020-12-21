package exporter

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	accessKeyId     = os.Getenv("ACCESS_KEY_ID")
	accessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
	regionId        = os.Getenv("REGION_ID")
)

const (
	PROJECT string = "acs_kafka"
)

type KafkaExporter struct {
	client     *cms.Client
	DataPoints []struct {
		InstanceId    string  `json:"instanceId"`
		Maximum       float64 `json:"Maximum,omitempty"`
		Value         float64 `json:"Value,omitempty"`
		ConsumerGroup string  `json:"consumerGroup,omitempty"`
		Topic         string  `json:"topic,omitempty"`
	}
	metrics        map[string]*prometheus.GaugeVec
	instances      map[string]string
	maxConnections map[string]float64
	metricMeta     []string
}
