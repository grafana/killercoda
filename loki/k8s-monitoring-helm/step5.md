# Accessing Grafana

To access Grafana, you will need to port-forward the Grafana service to your local machine. To do this, run the following command:

```bash
export POD_NAME=$(kubectl get pods --namespace meta -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}") && \
kubectl --namespace meta port-forward $POD_NAME 3000 --address 0.0.0.0
```{{exec}}

> **Tip:**
> This will make your terminal unusable until you stop the port-forwarding process. To stop the process, press `Ctrl + C`{{copy}}.

This command will port-forward the Grafana service to your local machine on port `3000`{{copy}}.

You can now access Grafana by navigating to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}) in your browser. The default credentials are `admin`{{copy}} and `adminadminadmin`{{copy}}.

One of the first places you should visit is Explore Logs which lets you automatically visualize and explore your logs without having to write queries:
[http://localhost:3000/a/grafana-lokiexplore-app]({{TRAFFIC_HOST1_3000}}/a/grafana-lokiexplore-app)

![Explore Logs view of K8s logs](https://grafana.com/media/docs/loki/k8s-logs-explore-logs.png)

# (Optional): View the Alloy UI

The Kubernetes Monitoring Helm chart deploys Grafana Alloy to collect and forward telemetry data from the Kubernetes cluster. The Helm is designed to abstract you from creating an Alloy configuration file. However if you would like to understand the pipeline you can view the Alloy UI. To access the Alloy UI, you will need to port-forward the Alloy service to your local machine. To do this, run the following command:

```bash
export POD_NAME=$(kubectl get pods --namespace meta -l "app.kubernetes.io/name=alloy-logs,app.kubernetes.io/instance=k8s" -o jsonpath="{.items[0].metadata.name}") && \
kubectl --namespace meta port-forward $POD_NAME 12345 --address 0.0.0.0
```{{exec}}

> **Tip:**
> This will make your terminal unusable until you stop the port-forwarding process. To stop the process, press `Ctrl + C`{{copy}}.

This command will port-forward the Alloy service to your local machine on port `12345`{{copy}}. You can access the Alloy UI by navigating to [http://localhost:12345]({{TRAFFIC_HOST1_12345}}) in your browser.

![Grafana Alloy UI](https://grafana.com/media/docs/loki/k8s-logs-alloy-ui.png)
