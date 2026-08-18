# Accessing Profiles Drilldown in Grafana

Grafana includes the [Profiles Drilldown](https://grafana.com/docs/grafana/latest/explore/simplified-exploration/profiles/) app that you can use to view profile data. To access Profiles Drilldown, open a browser and navigate to [http://localhost:3000/a/grafana-pyroscope-app/explore]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/explore).

## How tagging works

In this example, the application is instrumented with Pyroscope using the Python SDK.
The SDK allows you to tag functions with metadata that can be used to filter and group the profile data in the Profiles Drilldown.
This example uses static and dynamic tagging.

To start, let's take a look at a static tag use case. Within the `lib/server.py`{{copy}} file, find the Pyroscope configuration:

```python
 pyroscope.configure(
     application_name = app_name,
     server_address   = server_addr,
     basic_auth_username = basic_auth_username, # for grafana cloud
     basic_auth_password = basic_auth_password, # for grafana cloud
     tags             = {
         "region":   f'{os.getenv("REGION")}',
     }
 )
```{{copy}}

This tag is considered static because the tag is set at the start of the application and doesn't change.
In this case, it's useful for grouping profiles on a per region basis, which lets you see the performance of the application per region.

1. Open Grafana using the following url: [http://localhost:3000/a/grafana-pyroscope-app/explore]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/explore).
1. In the main menu, select **Drilldown** > **Profiles**.
1. Select  **Labels** in the **Exploration** path.
1. Select  **ride-sharing-app** in the **Service** drop-down menu.
1. Select the **region** tab in the **Group by labels** section.

You should now see a list of regions that the application is running in. You can see that `eu-north`{{copy}} is experiencing the most load.

![Region Tag](https://grafana.com/media/docs/pyroscope/ride-share-tag-region-2.png)

Next, look at a dynamic tag use case. Within the `lib/utility/utility.py`{{copy}} file,  find the following function:

```python
 def find_nearest_vehicle(n, vehicle):
     with pyroscope.tag_wrapper({ "vehicle": vehicle}):
         i = 0
         start_time = time.time()
         while time.time() - start_time < n:
             i += 1
         if vehicle == "car":
             check_driver_availability(n)
```{{copy}}

This example uses `tag_wrapper`{{copy}} to tag the function with the vehicle type.
Notice that the tag is dynamic as it changes based on the vehicle type.
This is useful for grouping profiles on a per vehicle basis, allowing us to see the performance of the application per vehicle type being requested.

Use Profiles Drilldown to see how this tag is used:

1. Open Profiles Drilldown using the following url: [http://localhost:3000/a/grafana-pyroscope-app/explore]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/explore).
1. Select on **Labels** in the **Exploration** path.
1. In the **Group by labels** section, select the **vehicle** tab.

You should now see a list of vehicle types that the application is using. You can see that `car`{{copy}} is experiencing the most load.
