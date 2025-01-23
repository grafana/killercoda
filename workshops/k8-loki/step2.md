# Step 2: Add the Grafana Helm repository

All three helm charts (Loki, Grafana, and the Kubernetes Monitoring Helm) are available in the Grafana Helm repository. Add the Grafana Helm repository by running the following command:

```bash
helm repo add grafana https://grafana.github.io/helm-charts && helm repo update
```{{exec}}

It’s recommended to also run `helm repo update`{{copy}} to ensure you have the latest version of the charts.
