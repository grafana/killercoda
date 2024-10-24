# Create an alert rule that returns alert instances

The alert rule that you are about to create is meant to monitor web traffic page views. The objective is to explore what an alert instance is and how to leverage routing individual alert instances by using label matchers and notification policies.

## Add a data source

Grafana includes a [test data source](https://grafana.com/docs/grafana/latest/datasources/testdata/) that creates simulated time series data.

1. In Grafana navigate to **Connections > Add new connection**.

1. Search for **TestData**.

1. Click **Add new data source**.

1. Click **Save & test**.

   You should see a message confirming that the data source is working.
