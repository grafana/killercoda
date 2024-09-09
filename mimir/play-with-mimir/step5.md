# Configure your first alert rule

Alerting rules allow you to define alert conditions based on PromQL expressions and to send notifications about firing
alerts to Grafana Mimir Alertmanager. In this section you’re going to configure an alerting rule in Grafana Mimir using
tooling offered by Grafana.

1. Open [Grafana Alerting]({{TRAFFIC_HOST1_9000}}/alerting/list).

1. Click **New alert rule**.

1. Configure the alert rule:
   1. Select **Mimir or Loki alert** in the top selector.

   1. Choose **Mimir** in the **Select data source** field.

   1. Type `up == 0`{{copy}} in the **Metrics browser** query field.

   1. Type `MimirNotRunning`{{copy}} in the **Rule name** field.

   1. Select `example-namespace`{{copy}} in the **Namespace** field.

   1. Select `example-group`{{copy}} in the **Group** field.

   1. From the upper-right corner, click the **Save and exit** button.

Your `MimirNotRunning`{{copy}} alert rule is now being created in Grafana Mimir ruler and is expected to fire when the number of
Grafana Mimir instances is less than three. You can check its status by opening the [Grafana Alerting]({{TRAFFIC_HOST1_9000}}/alerting/list)
page and expanding the “example-namespace > example-group” row. The status should be “Normal” since all three instances are currently running.

To see the alert firing we can introduce an outage in the Grafana Mimir cluster:

1. Abruptly terminate one of the three Grafana Mimir instances:
   ```bash
   docker compose kill mimir-3
   ```{{exec}}

1. Open [Grafana Alerting]({{TRAFFIC_HOST1_9000}}/alerting/list) and check out the state of the alert `MimirNotRunning`{{copy}},
   which should switch to “Pending” state in about one minute and to “Firing” state after another minute. _Note: since we abruptly
   terminated a Mimir instance, Grafana Alerting UI may temporarily show an error when querying rules: the error will
   auto resolve shortly, as soon as Grafana Mimir internal health checking detects the terminated instance as unhealthy._

Grafana Mimir Alertmanager has not been configured yet to notify alerts through a notification channel. To configure the
Alertmanager you can open the [Contact points]({{TRAFFIC_HOST1_9000}}/alerting/notifications) page in Grafana and
set your preferred notification channel. Note the email receiver doesn’t work in this example because there’s no
SMTP server running.

Before adding back our terminated Mimir instance to resolve the alert, go into the Grafana Explore page and query your `sum:up`{{copy}}
recording rule. You should see that value of `sum:up`{{copy}} should have dropped to `2`{{copy}}, now that one instance is down. You’ll also notice
that querying for this rule and all other metrics continues to work even though one instance is down. This demonstrates that highly
available Grafana Mimir setups like the three instance deployment in this demo are resilient to outages of individual nodes.

To resolve the alert and recover from the outage, restart the Grafana Mimir instance that was abruptly terminated:

1. Start the Grafana Mimir instances:
   ```bash
   docker-compose start mimir-3
   ```{{exec}}

1. Open [Grafana Alerting]({{TRAFFIC_HOST1_9000}}/alerting/list) and check out the state of the alert `MimirNotRunning`{{copy}},
   which should switch to “Normal” state in about one minute.
