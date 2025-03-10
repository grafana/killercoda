# Deploy the Kubernetes Monitoring Helm chart

The Kubernetes Monitoring Helm chart is used for gathering, scraping, and forwarding Kubernetes telemetry data to a Grafana stack. This includes the ability to collect metrics, logs, traces, and continuous profiling data. The scope of this tutorial is to deploy the Kubernetes Monitoring Helm chart to collect pod logs and Kubernetes events.

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
  labelsToKeep: ["app_kubernetes_io_name","container","instance","job","level","namespace","service_name","service_namespace","deployment_environment","deployment_environment_name"]
  structuredMetadata:
    pod: pod  # Set structured metadata "pod" from label "pod"
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
  # Required when using the Kubernetes API to pod logs
  alloy:
    mounts:
      varlog: false
    clustering:
      enabled: true

alloy-profiles:
  enabled: false

alloy-receiver:
  enabled: false
```{{copy}}

To break down the configuration file:

- Define the cluster name as `meta-monitoring-tutorial`{{copy}}. This a static label that will be attached to all logs collected by the Kubernetes Monitoring Helm chart.

- Define a destination named `loki`{{copy}} that will be used to forward logs to Loki. The `url`{{copy}} attribute specifies the URL of the Loki gateway. **If you choose to deploy Loki in a different namespace or in a different location entirely, you will need to update the `url`{{copy}} attribute accordingly.**

- Enable the collection of cluster events and pod logs:
  - `collector`{{copy}}: specifies which collector to use to collect logs. In this case, we are using the `alloy-logs`{{copy}} collector.

  - `labelsToKeep`{{copy}}: specifies the labels to keep when collecting logs. Note this does not drop logs. This is useful when you do not want to apply a high cardanility label. In this case we have removed `pod`{{copy}} from the labels to keep.

  - `structuredMetadata`{{copy}}: specifies the structured metadata to collect. In this case, we are setting the structured metadata `pod`{{copy}} so we can retain the pod name for querying. Though it does not need to be indexed as a label.zw

  - `namespaces`{{copy}}: specifies the namespaces to collect logs from. In this case, we are collecting logs from the `meta`{{copy}} and `prod`{{copy}} namespaces.

- Disable the collection of node logs for the purpose of this tutorial as it requires the mounting of `/var/log/journal`{{copy}}. This is out of scope for this tutorial.

- Lastly, define the role of the collector. The Kubernetes Monitoring Helm chart will deploy only what you need and nothing more. In this case, we are telling the Helm chart to only deploy Alloy with the capability to collect logs. If you need to collect K8s metrics, traces, or continuous profiling data, you can enable the respective collectors.
