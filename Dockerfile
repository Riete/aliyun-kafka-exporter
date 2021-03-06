FROM riet/golang:1.13.10 as backend
COPY . .
RUN unset GOPATH && go build -mod=vendor

FROM riet/centos:7.4.1708-cnzone
COPY --from=backend /go/aliyun-kafka-exporter /opt/aliyun-kafka-exporter
EXPOSE 10004
CMD /opt/aliyun-kafka-exporter
