# Loki Configuration

Grafana Loki requires a configuration file to define how it should run. Within the `loki-fundamentals`{{copy}} directory, you will find a file called `loki-config.yaml`{{copy}}:

```yaml
auth_enabled: false

server:
  http_listen_port: 3100
  grpc_listen_port: 9096
  log_level: info
  grpc_server_max_concurrent_streams: 1000

common:
  instance_addr: 127.0.0.1
  path_prefix: /tmp/loki
  storage:
    filesystem:
      chunks_directory: /tmp/loki/chunks
      rules_directory: /tmp/loki/rules
  replication_factor: 1
  ring:
    kvstore:
      store: inmemory

query_range:
  results_cache:
    cache:
      embedded_cache:
        enabled: true
        max_size_mb: 100

limits_config:
  metric_aggregation_enabled: true
  allow_structured_metadata: true
  volume_enabled: true
  retention_period: 24h   # 24h

schema_config:
  configs:
    - from: 2020-10-24
      store: tsdb
      object_store: filesystem
      schema: v13
      index:
        prefix: index_
        period: 24h

pattern_ingester:
  enabled: true
  metric_aggregation:
    loki_address: localhost:3100

ruler:
  enable_alertmanager_discovery: true
  enable_api: true
  
frontend:
  encoding: protobuf

compactor:
  working_directory: /tmp/loki/retention
  delete_request_store: filesystem
  retention_enabled: true
```{{copy}}

To summarize the configuration file:

- **auth_enabled**: This is set to false, meaning Loki does not need a [tenant ID](https://grafana.com/docs/loki/latest/operations/multi-tenancy/) for ingest or query.

- **server**: Defines the ports Loki listens on, the log level, and the maximum number of concurrent gRPC streams.

- **common**:  Defines the common configuration for Loki. This includes the instance address, storage configuration, replication factor, and ring configuration.

- **query_range**: This is defined to tell Loki to use inbuilt caching for query results. In production environments of Loki this is handled by a seperate cache service such as memcached.

- **limits_config**: Defines the global limits for all Loki tenants. This includes enabling specific features such as metric aggregation and structured metadata. Limits can be defined on a per tenant basis, however this is considered an advanced configuration and for most use cases the global limits are sufficient.

- **schema_config**: Defines the schema configuration for Loki. This includes the schema version, the object store, and the index configuration.

- **pattern_ingester**: Enables pattern ingesters which are used to discover log patterns. Mostly used by Grafana Logs Drilldown.

- **ruler**: Enables the ruler component of Loki. This is used to create alerts based on log queries.

- **frontend**: Defines the encoding format for the frontend. In this case it is set to `protobuf`{{copy}}.

- **compactor**: Defines the compactor configuration. Used to compact the index and mange chunk retention.

The above configuration file is a basic configuration file for Loki. For more advanced configuration options, refer to the [Loki Configuration](https://grafana.com/docs/loki/latest/configuration/) documentation.
