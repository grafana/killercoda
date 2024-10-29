# Identifying the performance bottleneck

The first step when analyzing a profile outputted from your application, is to take note of the largest node which is where your application is spending the most resources. To discover this, you can use the `Flame Graph`{{copy}} view:

1. Open the Profile Explorer in Grafana UI, which can be accessed using the following url: [http://localhost:3000/a/grafana-pyroscope-app/profiles-explorer]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/profiles-explorer).

1. Select the `Single View`{{copy}} tab.

1. Make sure `flask-ride-sharing-app:process_cpu:cpu`{{copy}} is selected in the dropdown menu.

It should look something like this:

![Bottleneck](https://grafana.com/media/docs/pyroscope/ride-share-bottle-neck-3.png)

The flask `dispatch_request`{{copy}} function is the parent to three functions that correspond to the three endpoints of the application:

- `order_bike`{{copy}}

- `order_car`{{copy}}

- `order_scooter`{{copy}}

By tagging both `region`{{copy}} and `vehicle`{{copy}} and looking at the Tag Explorer page, you can hypothesize:

- Something is wrong with the `/car`{{copy}} endpoint code where `car`{{copy}} vehicle tag is consuming **68% of CPU**

- Something is wrong with one of our regions where `eu-north`{{copy}} region tag is consuming **54% of CPU**

From the flame graph we can see that for the `eu-north`{{copy}} tag the biggest performance impact comes from the `find_nearest_vehicle()`{{copy}} function which consumes close to **68% of cpu**. To analyze this we can go directly to the comparison page using the comparison dropdown.

## Comparing two time periods

The comparison page allows you to compare two time periods side by side. This is useful for identifying changes in performance over time. In this example we will compare the performance of the `eu-north`{{copy}} region within a given time period against the other regions.

1. Open the Profile Explorer in Grafana UI, which can be accessed using the following url: [http://localhost:3000/a/grafana-pyroscope-app/profiles-explorer]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/profiles-explorer).

1. Select on `Diff flame graph`{{copy}} in the `Exploration`{{copy}} path.

1. In `Baseline`{{copy}} filter by `region`{{copy}} and select `!= eu-north`{{copy}}.

1. In `Comparison`{{copy}} filter by `region`{{copy}} and select `== eu-north`{{copy}}.

1. In `Baseline`{{copy}} select the time period you want to compare against.

Scroll down to compare the two time periods side by side.
Note that the `eu-north`{{copy}} region (right hand side) shows an excessive amount of time spent in the `find_nearest_vehicle`{{copy}} function.
This looks to be caused by a mutex lock that is causing the function to block.

![Time Comparison](https://grafana.com/media/docs/pyroscope/ride-share-time-comparison-2.png)
