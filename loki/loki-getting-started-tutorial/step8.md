# Grafana Loki Data source

The final piece of the puzzle is the Grafana Loki datasource. This is used by Grafana to connect to Loki and query the logs. Grafana has multiple ways to define a datasource;

- **Direct**: This is where you define the datasource in the Grafana UI.

- **Provisioning**: This is where you define the datasource in a configuration file and have Grafana automatically create the datasource.

- **API**: This is where you use the Grafana API to create the datasource.

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

Within the entrypoint section of the `docker-compose.yml`{{copy}} file, we have defined a file called `run.sh`{{copy}} this runs on startup and creates the datasource configuration file `ds.yaml`{{copy}} in the Grafana provisioning directory.
This file defines the Loki datasource and tells Grafana to use it. Since Loki is running in the same Docker network as Grafana, we can use the service name `loki`{{copy}} as the URL.
