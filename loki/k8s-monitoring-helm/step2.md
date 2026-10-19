# Add the Grafana Helm repository

The Grafana and Kubernetes Monitoring Helm charts are available in the Grafana Helm repository. The Loki Helm chart is maintained separately in the Grafana Community Helm Charts repository. Add both Helm repositories by running the following commands:

```bash
helm repo add grafana https://grafana.github.io/helm-charts && helm repo update
helm repo add grafana-community https://grafana-community.github.io/helm-charts && helm repo update
```{{exec}}

As well as adding the repos to your local helm list, you should also run `helm repo update`{{copy}} to ensure you have the latest version of the charts.

# Clone the tutorial repository

Clone the tutorial repository by running the following command:

```bash
git clone https://github.com/grafana/alloy-scenarios.git
```{{exec}}

Then change directories to the `alloy-scenarios/k8s/logs`{{copy}} directory:

```bash
cd alloy-scenarios/k8s/logs
```{{exec}}

**The rest of this tutorial assumes you are in the `alloy-scenarios/k8s/logs`{{copy}} directory.**
