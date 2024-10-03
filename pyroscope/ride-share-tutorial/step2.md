# Accessing the Pyroscope UI

Pyroscope includes a web-based UI that you can use to view the profile data. To access the Pyroscope UI, open a browser and navigate to [http://localhost:4040]({{TRAFFIC_HOST1_4040}}).

## How tagging works

In this example, the application is instrumented with Pyroscope using the Python SDK. The SDK allows you to tag functions with metadata that can be used to filter and group the profile data in the Pyroscope UI. In this example we have used two forms of tagging; static and dynamic.

To start lets take a look at a static tag use case. Within the `server.py`{{copy}} file we can find the Pyroscope configuration:

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

The reason this tag is considered static is due to the fact that the tag is set at the start of the application and doesn’t change. In our case this is useful for grouping profiles on a per region basis. Allowing us to see the performance of the application per region.

Lets take a look within the Pyroscope UI to see how this tag is used:

1. Open the Pyroscope UI in your browser at [http://localhost:4040]({{TRAFFIC_HOST1_4040}}).

1. Click on `Tag Explorer`{{copy}} in the left-hand menu.

1. Select the `region`{{copy}} tag from the dropdown menu.

You should now see a list of regions that the application is running in. You can see that `eu-north`{{copy}} is experiencing the most load.

![Region Tag](https://grafana.com/media/docs/pyroscope/ride-share-tag-region.png)

Next lets take a look at a dynamic tag use case. Within the `utils.py`{{copy}} file we can find the following function:

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

In this example we are `tag_wrapper`{{copy}} to tag the function with the vehicle type. Notice that the tag is dynamic as it changes based on the vehicle type. This is useful for grouping profiles on a per vehicle basis. Allowing us to see the performance of the application per vehicle type being requested.

Lets take a look within the Pyroscope UI to see how this tag is used:

1. Open the Pyroscope UI in your browser at [http://localhost:4040]({{TRAFFIC_HOST1_4040}}).

1. Click on `Tag Explorer`{{copy}} in the left-hand menu.

1. Select the `vehicle`{{copy}} tag from the dropdown menu.

You should now see a list of vehicle types that the application is using. You can see that `car`{{copy}} is experiencing the most load.
