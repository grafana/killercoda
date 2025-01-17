# Step 6: Deploy the Kubernetes Monitoring Helm

The Kubernetes Monitoring Helm chart is used for gathering, scraping, and forwarding Kubernetes telemetry data to a Grafana Stack. This includes the ability to collect; metrics, logs, traces & continuous profiling data. The scope of this tutorial is to deploy the Kubernetes Monitoring Helm chart to collect pod logs and Kubernetes events.

To deploy the Kubernetes Monitoring Helm chart run the following command:

```bash
helm install --values ./k8s-monitoring-values.yml k8s grafana/k8s-monitoring -n meta 
```{{exec}}

Within the configuration file `k8s-monitoring-values.yml`{{copy}} we have defined the following:

```yaml
---
cluster:
  name: meta-monitoring-tutorial

destinations:
  - name: loki
    type: loki
    url: http://loki-gateway.meta.svc.cluster.local/loki/api/v1/push


clusterEvents:
  enabled: true
  collector: alloy-logs
  namespaces:
    - meta
    - prod

nodeLogs:
  enabled: false

podLogs:
  enabled: true
  gatherMethod: kubernetesApi
  collector: alloy-logs
  namespaces:
    - meta
    - prod

# Collectors
alloy-singleton:
  enabled: false

alloy-metrics:
  enabled: false

alloy-logs:
  enabled: true

alloy-profiles:
  enabled: false

alloy-receiver:
  enabled: false
```{{copy}}

To break down the configuration file:

- We define the cluster name as `meta-monitoring-tutorial`{{copy}}. This a static label that will be attached to all logs collected by the Kubernetes Monitoring Helm chart.

- We define a destination named `loki`{{copy}} that will be used to forward logs to Loki. The `url`{{copy}} attribute specifies the URL of the Loki gateway. **If you choose to deploy Loki in a different namespace or with a different name, you will need to update the `url`{{copy}} attribute accordingly.**

- We enable the collection of cluster events and pod logs:
  - `collector`{{copy}}: specifies which collector to use to collect logs. In this case, we are using the `alloy-logs`{{copy}} collector.

  - `namespaces`{{copy}}: specifies the namespaces to collect logs from. In this case, we are collecting logs from the `meta`{{copy}} and `prod`{{copy}} namespaces.

- We disable the collection of node logs for the purpose of this tutorial as it requires the mounting of `/var/log/journal`{{copy}}. This is out of scope for this tutorial.

- Lastly, we define the role of the collector. The Kubernetes Monitoring Helm chart is designed to deploy only what you need and nothing more. So in this case, we are telling the Helm chart to only deploy Alloy with the capability to collect logs. If you need to collect K8s metrics, traces, or continuous profiling data, you can enable the respective collectors.
