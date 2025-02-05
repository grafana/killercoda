# Deploy Loki

Grafana Loki will be used to store our collected logs. In this tutorial we will deploy Loki with a minimal footprint and use the default storage backend provided by the Loki Helm chart, MinIO.

> **Note**: Due to the resource constraints of the Kubernetes cluster running in the playground, we are deploying Loki using a custom values file. This values file reduces the resource requirements of Loki. This turns off features such as cache and Loki Canary, and runs Loki with limited resources. This can take up to **1 minute** to complete.

To deploy Loki run the following command:

```bash
helm install --values killercoda/loki-values.yml loki grafana/loki -n meta
```{{exec}}

This command will deploy Loki in the `meta`{{copy}} namespace. The command also includes a `values`{{copy}} file that specifies the configuration for Loki. For more details on how to configure the Loki Helm chart refer to the Loki Helm [documentation](https://grafana.com/docs/loki/latest/setup/install/helm).
