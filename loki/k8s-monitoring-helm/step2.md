# Add the Grafana Helm repository

All three Helm charts (Loki, Grafana, and the Kubernetes Monitoring Helm) are available in the Grafana Helm repository. Add the Grafana Helm repository by running the following command:

```bash
helm repo add grafana https://grafana.github.io/helm-charts && helm repo update
```{{exec}}

As well as adding the repo to our local helm list, we also run `helm repo update`{{copy}} to ensure you have the latest version of the charts.

# Clone the tutorial repository

Clone the tutorial repository by running the following command:

```bash
git clone https://github.com/grafana/alloy-scenarios.git
```{{exec}}

Then change directories to the `alloy-scenarios/k8s/logs`{{copy}} directory:

```bash
cd alloy-scenarios/k8s/logs
```{{exec}}

**The rest of this tutorial assumes you are in this directory.**
