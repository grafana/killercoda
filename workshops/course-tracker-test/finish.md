# What next?

## Back to docs

Head back to where you started from to continue with the Loki documentation: [Loki documentation](https://grafana.com/docs/loki/latest/get-started/quick-start/).

You have completed the Loki Quickstart demo. So where to go next? Here are a few suggestions:

- **Deploy:** Loki can be deployed in multiple ways. For production use cases we recommend deploying Loki via the [Helm chart](https://grafana.com/docs/loki/latest/setup/install/helm/).

- **Send Logs:** In this example we used Grafana Alloy to collect and send logs to Loki. However there are many other methods you can use depending upon your needs. For more information see [send data](https://grafana.com/docs/loki/latest/send-data/).

- **Query Logs:** LogQL is an extensive query language for logs and contains many tools to improve log retrival and generate insights. For more information see the [Query section](https://grafana.com/docs/loki/latest/query/).

- **Alert:** Lastly you can use the ruler component of Loki to create alerts based on log queries. For more information see [Alerting](https://grafana.com/docs/loki/latest/alert/).

## Complete metrics, logs, traces, and profiling example

If you would like to run a demonstration environment that includes Mimir, Loki, Tempo, and Grafana, you can use [Introduction to Metrics, Logs, Traces, and Profiling in Grafana](https://github.com/grafana/intro-to-mlt).
It’s a self-contained environment for learning about Mimir, Loki, Tempo, and Grafana.

The project includes detailed explanations of each component and annotated configurations for a single-instance deployment.
You can also push the data from the environment to [Grafana Cloud](https://grafana.com/cloud/).
