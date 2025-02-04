# Step 4: Deploy Loki

Grafana Loki will be used to store our collected logs. In this tutorial we will deploy Loki with a minimal footprint and use the default storage backend provided by the Loki Helm chart, MinIO.

> **Note**: Due to the resource constraints of the Kubernetes cluster running in the playground, we are deploying Loki using a custom values file. This values file reduces the resource requirements of Loki. This turns off features such as cache and Loki Canary, and runs Loki with limited resources. This can take up to **1 minute** to complete.

To deploy Loki run the following command:

```bash
helm install --values killercoda/loki-values.yml loki grafana/loki -n meta
```{{exec}}

This command will deploy Loki in the `meta`{{copy}} namespace. The command also includes a `values`{{copy}} file that specifies the configuration for Loki. For more details on how to configure the Loki Helm chart refer to the Loki Helm [documentation](https://grafana.com/docs/loki/latest/setup/install/helm).

# Step 5: Deploy Grafana

Next we will deploy Grafana to the `meta`{{copy}} namespace. You will use Grafana to visualize the logs stored in Loki. To deploy Grafana run the following command:

```bash
helm install --values grafana-values.yml grafana grafana/grafana --namespace meta
```{{exec}}

As before the command also includes a `values`{{copy}} file that specifies the configuration for Grafana. There are two important configuration attributes to take note of:

1. `adminUser`{{copy}} & `adminPassword`{{copy}}: These are the credentials you will use to log in to Grafana. The values are `admin`{{copy}} and `adminadminadmin`{{copy}} respectively. The recommended practice is to either use a Kubernetes secret or allow Grafana to generate a password for you. For more details on how to configure the Grafana Helm chart, refer to the Grafana Helm [documentation](https://grafana.com/docs/grafana/latest/installation/helm/).

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

This configuration defines a data source named `Loki`{{copy}} that Grafana will use to query logs stored in Loki. The `url`{{copy}} attribute specifies the URL of the Loki gateway. The Loki gateway is a service that sits in front of the Loki API and provides a single endpoint for ingesting and querying logs. The URL is in the format `http://loki-gateway.meta.svc.cluster.local:80`{{copy}}. The `loki-gateway`{{copy}} service is created by the Loki Helm chart and is used to query logs stored in Loki. **If you choose to deploy Loki in a different namespace or with a different name, you will need to update the `url`{{copy}} attribute accordingly.**
