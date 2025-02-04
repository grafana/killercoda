# Step 1: Create the `meta`{{copy}} and `prod`{{copy}} namespaces

The K8s Monitoring Helm chart will monitor two namespaces: `meta`{{copy}} and `prod`{{copy}}:

- `meta`{{copy}} namespace: This namespace will be used to deploy Loki, Grafana, and Alloy.

- `prod`{{copy}} namespace: This namespace will be used to deploy the sample application that will generate logs.

Create the `meta`{{copy}} and `prod`{{copy}} namespaces by running the following command:

```bash
kubectl create namespace meta && kubectl create namespace prod
```{{exec}}
