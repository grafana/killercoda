# Configure your first recording rule

Recording rules allow you to precompute frequently needed or computationally expensive expressions and save their result
as a new set of time series. In this section you’re going to configure a recording rule in Grafana Mimir using tooling
offered by Grafana.

1. Open [Grafana Alerting]({{TRAFFIC_HOST1_9000}}/alerting/list).

1. Click **New recording rule**, which also allows you to configure recording rules.

1. Configure the recording rule:
   1. Give the rule a name, such as `sum:up`{{copy}}.

   1. Choose **Mimir** in the **Select data source** field.

   1. Choose **Code** in the **Builder | Code** field on the right.

   1. Type `sum(up)`{{copy}} in the **Metrics browser** query field.

   1. Type `example-namespace`{{copy}} in the **Namespace** field.

   1. Type `example-group`{{copy}} in the **Group** field.

   1. From the upper-right corner, click the **Save and exit** button.

Your `sum:up`{{copy}} recording rule will show the number of Mimir instances that are `up`{{copy}}, meaning reachable to be scraped. The
rule is now being created in Grafana Mimir ruler and will be soon available for querying:

1. Open [Grafana Explore]({{TRAFFIC_HOST1_9000}}/explore)
   and query the resulting series from the recording rule, which may require up to one minute to display after configuration:
   ```
   sum:up
   ```{{copy}}

1. Confirm the query returns a value of `3`{{copy}} which is the number of Mimir instances currently running in your local setup.
