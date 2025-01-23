# Step 9: Adding a sample application to `prod`{{copy}}

Lastly, lets deploy a sample application to the `prod`{{copy}} namespace that will generate some logs. To deploy the sample application run the following command:

```bash
helm install tempo grafana/tempo-distributed -n prod
```{{exec}}

This will deploy a default version of Grafana Tempo to the `prod`{{copy}} namespace. Tempo is a distributed tracing backend that is used to store and query traces. Normally Tempo would sit alongside Loki and Grafana in the `meta`{{copy}} namespace, but for the purpose of this tutorial, we will pretend this is the primary application generating logs.

Once deployed lets expose Grafana once more:

```bash
export POD_NAME=$(kubectl get pods --namespace meta -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}") && \
kubectl --namespace meta port-forward $POD_NAME 3000 --address 0.0.0.0
```{{exec}}

and navigate to [http://localhost:3000/a/grafana-lokiexplore-app]({{TRAFFIC_HOST1_3000}}/a/grafana-lokiexplore-app) to view Grafana Tempo logs.

![Label view of Tempo logs](https://grafana.com/media/docs/loki/k8s-logs-tempo.png)
