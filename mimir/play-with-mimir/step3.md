# Explore Grafana Mimir dashboards

Open Grafana on your local host [`http://localhost:9000`{{copy}}]({{TRAFFIC_HOST1_9000}}) and view dashboards showing the status
and health of your Grafana Mimir cluster. The dashboards query Grafana Mimir for the metrics they display.

To start, we recommend looking at these dashboards:

- [Writes]({{TRAFFIC_HOST1_9000}}/d/8280707b8f16e7b87b840fc1cc92d4c5/mimir-writes)

- [Reads]({{TRAFFIC_HOST1_9000}}/d/e327503188913dc38ad571c647eef643/mimir-reads)

- [Queries]({{TRAFFIC_HOST1_9000}}/d/b3abe8d5c040395cc36615cb4334c92d/mimir-queries)

- [Object Store]({{TRAFFIC_HOST1_9000}}/d/e1324ee2a434f4158c00a9ee279d3292/mimir-object-store)

A couple of caveats:

- It typically takes a few minutes after Grafana Mimir starts to display meaningful metrics in the dashboards.

- Because this tutorial runs Grafana Mimir without any query-scheduler, or memcached, the related panels are expected to be empty.

The dashboards installed in the Grafana are taken from the Grafana Mimir mixin which packages up Grafana Labs’ best practice dashboards, recording rules, and alerts for monitoring Grafana Mimir. To learn more about the mixin, check out the Grafana Mimir mixin documentation. To learn more about how Grafana is connecting to Grafana Mimir, review the [Mimir datasource]({{TRAFFIC_HOST1_9000}}/datasources).
