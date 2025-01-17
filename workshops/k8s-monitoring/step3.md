# Step 3: Deploy Loki

Grafana Loki will be used to store our collected logs. In this tutorial we will deploy Loki with a minimal footprint and use the default storage backend provided by the Loki Helm (MinIO). As mentioned earlier, it is recommended to use a more production-ready storage backend like S3, GCS, or Azure Blob Storage for production use cases.

To deploy Loki run the following command:

```bash
helm install --values loki-values.yml loki grafana/loki -n meta
```{{exec}}

This command will deploy Loki in the `meta`{{copy}} namespace. The command also includes a `values`{{copy}} file that specifies the configuration for Loki. For more details on how to configure the Loki Helm refer to the Loki Helm [documentation](https://grafana.com/docs/loki/latest/setup/install/helm).

# Step 4: Deploy Grafana

Next, we will deploy Grafana to visualize the logs stored in Loki. The deployment of the Grafana Helm chart is similar to the Loki Helm chart. To deploy Grafana run the following command:

```bash
helm install --values grafana-values.yml grafana grafana/grafana --namespace meta
```{{exec}}

This command will deploy Grafana in the `meta`{{copy}} namespace. As before the command also includes a `values`{{copy}} file that specifies the configuration for Grafana. There are two important configurations attributes to take note of:

1. `adminUser`{{copy}} & `adminPassword`{{copy}}: These are the credentials you will use to log in to Grafana. The values are `admin`{{copy}} and `adminadminadmin`{{copy}} respectively. The recommended practice is to either use a Kubernetes secret or allow Grafana to generate a password for you. For more details on how to configure the Grafana Helm refer to the Grafana Helm [documentation](https://grafana.com/docs/grafana/latest/installation/helm/).

1. `datasources`{{copy}}: This section of the configuration allows for the definition of datasources that Grafana will use. In this tutorial, we will define a datasource for Loki. The datasource is defined as follows:

```yaml
datasources:
  datasources.yaml:
        apiVersion: 1
        datasources:
        - name: Loki
          type: loki
          access: proxy
          orgId: 1
          url: http://loki-gateway.meta.svc.cluster.local:80
          basicAuth: false
          isDefault: false
          version: 1
          editable: false
```{{copy}}

This configuration defines a data source named `Loki`{{copy}} that Grafana will use to query logs stored in Loki. The `url`{{copy}} attribute specifies the URL of the Loki gateway. The Loki gateway is a service that sits in front of the Loki API and provides a single endpoint for ingesting and querying logs. The URL is in the format `http://loki-gateway.meta.svc.cluster.local:80`{{copy}}. The `loki-gateway`{{copy}} service is created by the Loki Helm chart and is used to query logs stored in Loki. **If you choose to deploy Loki in a different namespace or with a different name, you will need to update the `url`{{copy}} attribute accordingly.**
