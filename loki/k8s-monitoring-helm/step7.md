# (Optional) View the Alloy UI

The Kubernetes Monitoring Helm chart deploys Grafana Alloy to collect and forward telemetry data from the Kubernetes cluster. The Helm chart is designed to abstract you away from creating an Alloy configuration file. However if you would like to understand the pipeline you can view the Alloy UI. To access the Alloy UI, you will need to port-forward the Alloy service to your local machine. To do this, run the following command:

```bash
export POD_NAME=$(kubectl get pods --namespace meta -l "app.kubernetes.io/name=alloy-logs,app.kubernetes.io/instance=k8s" -o jsonpath="{.items[0].metadata.name}") && \
kubectl --namespace meta port-forward $POD_NAME 12345 --address 0.0.0.0
```{{exec}}

> **Tip:**
> This will make your terminal unusable until you stop the port-forwarding process. To stop the process, press `Ctrl + C`{{copy}}.

This command will port-forward the Alloy service to your local machine on port `12345`{{copy}}. You can access the Alloy UI by navigating to [http://localhost:12345]({{TRAFFIC_HOST1_12345}}) in your browser.

![Grafana Alloy UI](https://grafana.com/media/docs/loki/k8s-logs-alloy-ui.png)
