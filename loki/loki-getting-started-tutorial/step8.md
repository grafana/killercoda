# Grafana Loki Data source

The final piece of the puzzle is the Grafana Loki data source. This is used by Grafana to connect to Loki and query the logs. Grafana has multiple ways to define a data source:

- **Direct**: This is where you define the data source in the Grafana UI.
- **Provisioning**: This is where you define the data source in a configuration file and have Grafana automatically create the data source.
- **API**: This is where you use the Grafana API to create the data source.

In this case we are using the provisioning method. Instead of mounting the Grafana configuration directory, we have defined the data source in this portion of the `docker-compose.yml`{{copy}} file:

```yaml
  grafana:
    image: grafana/grafana:latest
    environment:
      - GF_FEATURE_TOGGLES_ENABLE=grafanaManagedRecordingRules
      - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
      - GF_AUTH_ANONYMOUS_ENABLED=true
      - GF_AUTH_BASIC_ENABLED=false
    ports:
      - 3000:3000/tcp
    entrypoint:
      - sh
      - -euc
      - |
         mkdir -p /etc/grafana/provisioning/datasources
         cat <<EOF > /etc/grafana/provisioning/datasources/ds.yaml
         apiVersion: 1
         datasources:
         - name: Loki
           type: loki
           access: proxy
           orgId: 1
           url: 'http://loki:3100'
           basicAuth: false
           isDefault: true
           version: 1
           editable: true 
         EOF
         /run.sh
    networks:
      - loki
```{{copy}}

The entrypoint overrides the default startup command to first create the datasource provisioning file `ds.yaml`{{copy}}, which configures Loki as the default datasource. It then calls `/run.sh`{{copy}}, the default Grafana startup script included in the `grafana/grafana`{{copy}} Docker image, which starts the Grafana server. Since Loki is running in the same Docker network as Grafana, the datasource URL uses the service name `loki`{{copy}}.
