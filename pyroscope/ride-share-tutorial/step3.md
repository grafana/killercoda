# Identifying the performance bottleneck

The first step when analyzing a profile outputted from your application, is to take note of the largest node which is where your application is spending the most resources. To discover this, you can use the `Flame Graph`{{copy}} view within the Pyroscope UI:

1. Open the Pyroscope UI in your browser at [http://localhost:4040]({{TRAFFIC_HOST1_4040}}).

1. Select the `Single View`{{copy}} tab.

1. Make sure `flask-ride-sharing-app:process_cpu:cpu`{{copy}} is selected in the dropdown menu.

It should look something like this:

![Bottleneck](https://grafana.com/media/docs/pyroscope/ride-share-bottle-neck.jpg)

The flask `dispatch_request`{{copy}} function is the parent to three functions that correspond to the three endpoints of the application:

- `order_bike`{{copy}}

- `order_car`{{copy}}

- `order_scooter`{{copy}}

The benefit of using Pyroscope, is that by tagging both `region`{{copy}} and `vehicle`{{copy}} and looking at the Tag Explorer page we can hypothesize:

- Something is wrong with the `/car`{{copy}} endpoint code where `car`{{copy}} vehicle tag is consuming **68% of CPU**

- Something is wrong with one of our regions where `eu-north`{{copy}} region tag is consuming **54% of CPU**

From the flame graph we can see that for the `eu-north`{{copy}} tag the biggest performance impact comes from the `find_nearest_vehicle()`{{copy}} function which consumes close to **68% of cpu**. To analyze this we can go directly to the comparison page using the comparison dropdown.

## Comparing two time periods

The comparison page allows you to compare two time periods side by side. This is useful for identifying changes in performance over time. In this example we will compare the performance of the `eu-north`{{copy}} region within a given time period against the other regions.

1. Open the Pyroscope UI in your browser at [http://localhost:4040]({{TRAFFIC_HOST1_4040}}).

1. Click on `Comparison`{{copy}} in the left-hand menu.

1. Within `Baseline time range`{{copy}} copy and paste the following query:
   ```console
   process_cpu:cpu:nanoseconds:cpu:nanoseconds{service_name="flask-ride-sharing-app", vehicle="car", region!="eu-north"}
   ```{{copy}}

1. Within `Comparison time range`{{copy}} copy and paste the following query:
   ```console
   process_cpu:cpu:nanoseconds:cpu:nanoseconds{service_name="flask-ride-sharing-app", vehicle="car", region="eu-north"}
   ```{{copy}}

1. Execute both queries by clicking the `Execute`{{copy}} buttons.

If we scroll down to compare the two time periods side by side we can see that the `eu-north`{{copy}} region (right hand side) we can see an excessive amount of time spent in the `find_nearest_vehicle`{{copy}} function. This looks to be caused by a mutex lock that is causing the function to block.

![Time Comparison](https://grafana.com/media/docs/pyroscope/ride-share-time-comparison.png)

To confirm our suspicions we can use the `Diff`{{copy}} view to see the difference between the two time periods.

## Viewing the difference between two time periods

The `Diff`{{copy}} view allows you to see the difference between two time periods. This is useful for identifying changes in performance over time. In this example we will compare the performance of the `eu-north`{{copy}} region within a given time period against the other regions.

1. Open the Pyroscope UI in your browser at [http://localhost:4040]({{TRAFFIC_HOST1_4040}}).

1. Click on `Diff`{{copy}} in the left-hand menu.

1. Make sure to have set the `Baseline time range`{{copy}} and `Comparison time range`{{copy}} queries as per the previous step.

1. Click on the `Execute`{{copy}} buttons.

If we scroll down to compare the two time periods side by side we can see that the `eu-north`{{copy}} region (right hand side) we can see an excessive amount of time spent in the `find_nearest_vehicle`{{copy}} function. This confirms our suspicions that the mutex lock that is causing the function to block.

![Diff](https://grafana.com/media/docs/pyroscope/ride-share-diff-page.png)
