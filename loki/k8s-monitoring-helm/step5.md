# Deploy the Kubernetes Monitoring Helm chart

The Kubernetes Monitoring Helm chart is used for gathering, scraping, and forwarding Kubernetes telemetry data to a Grafana stack. This includes the ability to collect metrics, logs, traces, and continuous profiling data. The scope of this tutorial is to deploy the Kubernetes Monitoring Helm chart to collect Pod logs and Kubernetes events.

> **Note:**
> This tutorial uses chart version 4.x, the current major version of the Kubernetes Monitoring Helm chart. If you have an existing `values.yml`{{copy}} file written for chart version 3.x or earlier, its configuration keys do not carry over directly. Refer to the [migration guide](https://grafana.com/docs/grafana-cloud/observe-and-act/monitor-infrastructure/kubernetes-monitoring/configuration/helm-chart-config/helm-chart/migrate-helm-chart/) before reusing an older values file with this tutorial.

To deploy the Kubernetes Monitoring Helm chart run the following command:

```bash
helm install --values ./k8s-monitoring-values.yml k8s grafana/k8s-monitoring -n meta --version "^4"
```{{exec}}

Pinning the chart to the 4.x line (`--version "^4"`{{copy}}) keeps this tutorial working even after a future major release changes the configuration format again.

Within the configuration file `k8s-monitoring-values.yml`{{copy}} we have defined the following:

```yaml
---
cluster:
  name: meta-monitoring-tutorial

destinations:
  loki:
    type: loki
    url: http://loki-gateway.meta.svc.cluster.local/loki/api/v1/push

clusterEvents:
  enabled: true
  collector: alloy-singleton
  namespaces:
    - meta
    - prod

podLogsViaKubernetesApi:
  enabled: true
  collector: alloy-logs
  namespaces:
    - meta
    - prod
  structuredMetadata:
    pod: pod  # Set structured metadata "pod" from label "pod"

# Collectors
collectors:
  alloy-singleton:
    presets: [singleton]
  alloy-logs:
    presets: [clustered]
```{{copy}}

To break down the configuration file:

- Define the cluster name as `meta-monitoring-tutorial`{{copy}}. This is a static label that will be attached to all logs collected by the Kubernetes Monitoring Helm chart.
- Define a destination named `loki`{{copy}} that will be used to forward logs to Loki. The `url`{{copy}} attribute specifies the URL of the Loki gateway. **If you choose to deploy Loki in a different namespace or in a different location entirely, you will need to update the `url`{{copy}} attribute accordingly.**
- Enable the collection of cluster events and Pod logs. Pod logs are gathered via the Kubernetes API, using the `podLogsViaKubernetesApi`{{copy}} feature:

  - `collector`{{copy}}: specifies which collector to use to collect logs. The value must match a name defined under `collectors:`{{copy}} below. In this case, `clusterEvents`{{copy}} uses the `alloy-singleton`{{copy}} collector, and `podLogsViaKubernetesApi`{{copy}} uses the `alloy-logs`{{copy}} collector.
  - `structuredMetadata`{{copy}}: specifies the structured metadata to collect. In this case, we are setting the structured metadata `pod`{{copy}} so we can retain the pod name for querying, without promoting it to an indexed label (which would add high [cardinality](https://grafana.com/docs/loki/latest/get-started/labels/cardinality/)).
  - `namespaces`{{copy}}: specifies the namespaces to collect logs from. In this case, we are collecting logs from the `meta`{{copy}} and `prod`{{copy}} namespaces.
- Node log collection is not enabled in this tutorial. The `nodeLogs`{{copy}} feature defaults to disabled, and requires mounting `/var/log/journal`{{copy}}, which is out of scope here.
- Define the `collectors:`{{copy}} map. Each entry is a name you choose for an Alloy instance, configured with one or more `presets:`{{copy}} that determine how that instance is deployed:

  - `alloy-singleton`{{copy}} uses the `singleton`{{copy}} preset, which deploys a single Alloy replica. This is sufficient for gathering cluster-wide data such as Kubernetes events.
  - `alloy-logs`{{copy}} uses the `clustered`{{copy}} preset, which deploys Alloy as a coordinated group of replicas. This is required for collecting Pod logs through the Kubernetes API, since the replicas need to coordinate which Pods each one collects logs from.

  At least one collector must be defined under `collectors:`{{copy}}, and every `collector:`{{copy}} reference elsewhere in the file (such as `clusterEvents.collector`{{copy}} and `podLogsViaKubernetesApi.collector`{{copy}}) must match a key in this map, or the install fails with an error naming the missing collector.
