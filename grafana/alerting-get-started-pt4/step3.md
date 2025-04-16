# Step 1: Template labels and annotations

Now that we've introduced how templating works, let’s move on to the next step. We guide you through creating an alert rule with a summary and description annotation. In doing so, we incorporate CPU usage and instance names, which we later use in our notification template.

## Create an alert rule

1. Sign in to Grafana:

   - **Grafana Cloud** users: Log in via Grafana Cloud.
   - **OSS users**: Go to [http://localhost:3000]({{TRAFFIC_HOST1_3000}}).
1. Create an alert rule that includes a summary and description annotation:

   - Navigate to **Alerts & IRM > Alerting > Alert rules**.
   - Click **+ New alert rule**.
   - Enter an **alert rule name**. Name it `High CPU usage`{{copy}}
1. **Define query an alert condition** section:

   - Select TestData data source from the drop-down menu.

     [TestData](https://grafana.com/docs/grafana/latest/datasources/testdata/) is included in the demo environment. If you’re working in Grafana Cloud or your own local Grafana instance, you can add the data source through the Connections menu.
   - From **Scenario** select **CSV Content**.
   - Copy in the following CSV data:

     ```
     region,cpu-usage,service,instance
     us-west,88,web-server-1,server-01
     us-west,81,web-server-1,server-02
     us-east,79,web-server-2,server-03
     us-east,52,web-server-2,server-04
     ```{{copy}}

     This dataset simulates a data source returning multiple time series, with each time series generating a separate alert instance.
1. **Alert condition** section:

   - Keep Last as the value for the reducer function (`WHEN`{{copy}}), and `75`{{copy}} as the threshold value, representing CPU usage above 75% .This is the value above which the alert rule should trigger.
   - Click **Preview alert rule condition** to run the queries.

   It should return 3 series in Firing state, and 1 in Normal state.

   ![Preview of a query returning alert instances](https://grafana.com/media/docs/alerting/part-4-firing-instances-preview.png)
1. Add folders and labels section:

   - In **Folder**, click **+ New folder** and enter a name. For example: `System metrics`{{copy}} . This folder contains our alert rules.

     Note: while it's possible to template labels here, in this tutorial, we focus on templating the summary and annotations fields instead.
1. **Set evaluation behaviour** section:

   - In the **Evaluation group and interval**, repeat the above step to create a new evaluation group. Name it `High usage`{{copy}}.
   - Choose an **Evaluation interval** (how often the alert will be evaluated). Choose `1m`{{copy}}.
   - Set the **pending period** to 0s (zero seconds), so the alert rule fires the moment the condition is met (this minimizes the waiting time for the demonstration.).
1. **Configure notifications** section:

   Select who should receive a notification when an alert rule fires.

   - Select a **Contact point**. If you don’t have any contact points, click _View or create contact points_.
1. **Configure notification message** section:

   In this step, you’ll configure the **summary** and **description** annotations to make your alert notifications informative and easy to understand. These annotations use templates to dynamically include key information about the alert.

   - **Summary** annotation: Enter the following code as the value for the annotation.:

     ```go
     {{- "\n" -}}
     Instance: {{ index $labels "instance" }}
     {{- "\t" -}} Usage: {{ index $values "A"}}%{{- "\n" -}}
     ```{{copy}}

     This template automatically adds the instance name (from the [$labels](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/templates/reference/#labels) data) and its current CPU usage (from [$values["A"]](https://grafana.com/docs/grafana/latest/alerting/alerting-rules/templates/reference/#values)) into the alert summary. `\t`{{copy}}: Adds a tab space between the instance name and the value. And, `\n`{{copy}}: Inserts a new line after the value.

     Output example:

     ```
     server-01	88
     ```{{copy}}

     This output helps you quickly see which instance is affected and its usage level.
1. Optional: Add a description to help the on-call engineer to better understand what the alert rule is about. Eg. This alert monitors CPU usage across instances and triggers if any instance exceeds a usage threshold of 75%.
1. Click **Save rule and exit**.

Now that we’ve configured an alert rule with dynamic templates for the **summary** annotation, the next step is to customize the alert notifications themselves. While the default notification message includes the summary annotation and works well, it can often be too verbose.

![Default email alert notification with templated annotation](https://grafana.com/media/docs/alerting/templated-annotation-alert.png)

To make our alert notifications more concise and tailored to our needs, we’ll create a custom **notification template** that references the summary annotation we just set up. Notification templates are especially useful because they can be reused across multiple contact points, ensuring consistent alert messages.
