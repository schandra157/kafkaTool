# kafkaTool

- Tool to create/Modify Topics in kafka Broker


# Setting up configuration

- Refer to config/kafka.yaml file

```yaml
bootstrapServers: "localhost:9092"
# below config is appliciable for newly created topics
replication: 1
partitions: 2
topics:
  - name: Topic4
  # below config is Topic specific which will overrided if topic already exist
    configs:
      retention.ms: 12000   
      compression.type: gzip  
  - name: Topic2
    configs:
      retention.ms: 48000
      compression.type: gzip  
```


# How to Run
  ``go run main.go``
