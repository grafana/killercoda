# Step 2: Add the Grafana Helm repository

All three helm charts (Loki, Grafana, and the Kubernetes Monitoring Helm) are available in the Grafana Helm repository. Add the Grafana Helm repository by running the following command:

```bash
helm repo add grafana https://grafana.github.io/helm-charts && helm repo update
```{{exec}}

It’s recommended to also run `helm repo update`{{copy}} to ensure you have the latest version of the charts.

# Step 3: Clone the tutorial repository

Clone the tutorial repository by running the following command:

```bash
git clone https://github.com/grafana/alloy-scenarios.git && cd alloy-scenarios/k8s-logs
```{{exec}}

As well as cloning the repository, we have also changed directories to `alloy-scenarios/k8s-logs`{{copy}}. **The rest of this tutorial assumes you are in this directory.**
