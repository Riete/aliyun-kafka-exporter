### docker build
``` docker build . -t <image>:<tag> ```

### or pull 
``` docker pull riet/aliyun-kafka-exporter ```

### run
```
docker run \ 
  -d \ 
  --name aliyun-kafka-exporter \
  -e ACCESS_KEY_ID=<aliyun ak> \
  -e ACCESS_KEY_SECRET=<aliyun ak sk> \
  -e REGION_ID=<region id> \
  -p 10004:10004 \
  riet/aliyun-kafka-exporter 
```

visit http://localhost:10004/metrics