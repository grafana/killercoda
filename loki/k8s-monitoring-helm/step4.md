# Deploy Grafana

Next we will deploy Grafana to the `meta`{{copy}} namespace. You will use Grafana to visualize the logs stored in Loki. To deploy Grafana run the following command:

```bash
helm install --values grafana-values.yml grafana grafana/grafana --namespace meta
```{{exec}}

As before, the command also includes a `values`{{copy}} file that specifies the configuration for Grafana. There are two important configuration attributes to take note of:

1. `adminUser`{{copy}} and `adminPassword`{{copy}}: These are the credentials you will use to log in to Grafana. The values are `admin`{{copy}} and `adminadminadmin`{{copy}} respectively. The recommended practice is to either use a Kubernetes secret or allow Grafana to generate a password for you. For more details on how to configure the Grafana Helm chart, refer to the Grafana Helm [documentation](https://grafana.com/docs/grafana/latest/installation/helm/).
1. `datasources`{{copy}}: This section of the configuration lets you define the data sources that Grafana should use. In this tutorial, you will define a Loki data source. The data source is defined as follows:

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

This configuration defines a data source named `Loki`{{copy}} that Grafana will use to query logs stored in Loki. The `url`{{copy}} attribute specifies the URL of the Loki gateway. The Loki gateway is a service that sits in front of the Loki API and provides a single endpoint for ingesting and querying logs. The URL is in the format `http://loki-gateway.<NAMESPACE>.svc.cluster.local:80`{{copy}}. The `loki-gateway`{{copy}} service is created by the Loki Helm chart and is used to query logs stored in Loki. **If you choose to deploy Loki in a different namespace or with a different name, you will need to update the `url`{{copy}} attribute accordingly.**
