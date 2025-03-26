Since Grafana Alloy is configured to tail logs from all Docker containers, Loki should already be receiving logs. The best place to verify log collection is using the Grafana Logs Drilldown feature. To do this, navigate to [http://localhost:3000/drilldown]({{TRAFFIC_HOST1_3000}}/drilldown). Select **Logs**. You should see the Grafana Logs Drilldown page.

![Grafana Logs Drilldown](https://grafana.com/media/docs/loki/get-started-drill-down.png)

If you have only the getting started demo deployed in your Docker environment, you should see three containers and their logs; `loki-fundamentals-alloy-1`{{copy}}, `loki-fundamentals-grafana-1`{{copy}} and `loki-fundamentals-loki-1`{{copy}}.  In the `loki-fundamentals-loki-1`{{copy}} container, click **Show Logs**  to drill down into the logs for that container.

![Grafana Drilldown Service View](https://grafana.com/media/docs/loki/get-started-drill-down-container.png)

We will not cover the rest of the Grafana Logs Drilldown features in this quickstart guide. For more information on how to use the Grafana Logs Drilldown feature, see [Get started with Grafana Logs Drilldown](https://grafana.com/docs/grafana/latest/explore/simplified-exploration/logs/get-started/) page.
