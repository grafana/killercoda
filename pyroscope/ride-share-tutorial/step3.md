# Identifying the performance bottleneck

The first step when analyzing a profile outputted from your application, is to take note of the largest node which is where your application is spending the most resources.
To discover this, you can use the **Flame graph** view:

1. Open Profiles Drilldown using the following url: [http://localhost:3000/a/grafana-pyroscope-app/profiles-explorer]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/profiles-explorer).

1. Select **Flame graph** from the **Exploration** path.

1. Verify that  `ride-sharing-app`{{copy}} is selected in the **Service** drop-down menu and `process_cpu/cpu`{{copy}} in the **Profile type** drop-down menu.

It should look something like this:

![Bottleneck](https://grafana.com/media/docs/pyroscope/ride-share-bottle-neck-3.png)

The flask `dispatch_request`{{copy}} function is the parent to three functions that correspond to the three endpoints of the application:

- `order_bike`{{copy}}

- `order_car`{{copy}}

- `order_scooter`{{copy}}

By tagging both `region`{{copy}} and `vehicle`{{copy}} and looking at the [**Labels** view](https://grafana.com/docs/grafana/latest/explore/simplified-exploration/profiles/choose-a-view/#labels), you can hypothesize:

- Something is wrong with the `/car`{{copy}} endpoint code where `car`{{copy}} vehicle tag is consuming **68% of CPU**

- Something is wrong with one of our regions where `eu-north`{{copy}} region tag is consuming **54% of CPU**

From the flame graph, you can see that for the `eu-north`{{copy}} tag the biggest performance impact comes from the `find_nearest_vehicle()`{{copy}} function which consumes close to **68% of cpu**.
To analyze this, go directly to the comparison page using the comparison dropdown.

## Comparing two time periods

The **Diff flame graph** view lets you compare two time periods side by side.
This is useful for identifying changes in performance over time.
This example compares the performance of the `eu-north`{{copy}} region within a given time period against the other regions.

1. Open Profiles Drilldown in Grafana using the following url: [http://localhost:3000/a/grafana-pyroscope-app/profiles-explorer]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/profiles-explorer).

1. Select **Diff flame graph** in the **Exploration** path.

1. In **Baseline**, filter by `region`{{copy}} and select `!= eu-north`{{copy}}.

1. In **Comparison**, filter by `region`{{copy}} and select `== eu-north`{{copy}}.

1. In **Baseline**, select the time period you want to compare against.

Scroll down to compare the two time periods side by side.
Note that the `eu-north`{{copy}} region (right side) shows an excessive amount of time spent in the `find_nearest_vehicle`{{copy}} function.
This looks to be caused by a mutex lock that is causing the function to block.

![Time Comparison](https://grafana.com/media/docs/pyroscope/ride-share-time-comparison-2.png)
