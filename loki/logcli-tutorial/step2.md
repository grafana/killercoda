# Querying logs

As part of our role within the logistics company, we need to build a report on the overall health of the shipments. Unfortunately, we only have access to a console and cannot use Grafana to visualize the data. We can use LogCLI to query the logs and build the report.

## Find all critical packages

To find all critical packages in the last hour (default lookback time), we can run the following query:

```bash
logcli query '{service_name="Delivery World"} | package_status="critical"'
```{{exec}}

This will return all logs where the `service_name`{{copy}} is `Delivery World`{{copy}} and the `package_status`{{copy}} is `critical`{{copy}}. The output will look similar to the following:

```console
http://localhost:3100/loki/api/v1/query_range?direction=BACKWARD&end=1732617594381712000&limit=30&query=%7Bservice_name%3D%22Delivery+World%22%7D+%7C+package_status%3D%22critical%22&start=1732613994381712000
Common labels: {package_status="critical", service_name="Delivery World"}
2024-11-26T10:39:52Z {package_id="PKG79755", package_size="Small", state="Texas"}       {"timestamp": "2024-11-26T10:39:52.521602Z", "state": "Texas", "city": "Dallas", "package_id": "PKG79755", "package_type": "Clothing", "package_size": "Small", "package_status": "critical", "note": "In transit", "sender": {"name": "Sender38", "address": "906 Maple Ave, Dallas, Texas"}, "receiver": {"name": "Receiver41", "address": "455 Pine Rd, Dallas, Texas"}}
2024-11-26T10:39:50Z {package_id="PKG34018", package_size="Large", state="Illinois"}    {"timestamp": "2024-11-26T10:39:50.510841Z", "state": "Illinois", "city": "Chicago", "package_id": "PKG34018", "package_type": "Clothing", "package_size": "Large", "package_status": "critical", "note": "Delayed due to weather", "sender": {"name": "Sender22", "address": "758 Elm St, Chicago, Illinois"}, "receiver": {"name": "Receiver10", "address": "441 Cedar Blvd, Naperville, Illinois"}}
```{{copy}}

Lets suppose we want to look back for the last 24 hours, we can use the `--since`{{copy}} flag to specify the time range:

```bash
logcli query --since 24h '{service_name="Delivery World"} | package_status="critical"' 
```{{exec}}

This will query all logs for the `package_status`{{copy}} `critical`{{copy}} in the last 24 hours. However it will not return all of the logs, but only the first 30 logs. We can use the `--limit`{{copy}} flag to specify the number of logs to return:

```bash
logcli query --since 24h --limit 100 '{service_name="Delivery World"} | package_status="critical"' 
```{{exec}}

## Metric Queries

We can also use LogCLI to query logs based on metrics. For instance as part of the site report we want to count the total number of packages sent from California in the last 24 hours in 1 hour intervals. We can use the following query:

```bash
logcli query --since 24h 'sum(count_over_time({state="California"}[1h]))'
```{{exec}}

This will return a JSON object containing a list of timestamps (Unix format) and the number of packages sent from California in 1 hour intervals. Since we summing the count of logs over time, we will see the total number of logs steadily increase over time. The output will look similar to the following:

```console
[
  {
    "metric": {},
    "values": [
      [
        1733913765,
        "46"
      ],
      [
        1733914110,
        "114"
      ],
      [
        1733914455,
        "179"
      ],
      [
        1733914800,
        "250"
      ],
      [
        1733915145,
        "318"
      ],
      [
        1733915490,
        "392"
      ],
      [
        1733915835,
        "396"
      ]
    ]
  }
]
```{{copy}}

We can take this a step further and filter the logs based on the `package_type`{{copy}} label. For instance, we can count the number of documents sent from California in the last 24 hours in 1 hour intervals:

```bash
logcli query --since 24h  'sum(count_over_time({state="California"}| json | package_type= "Documents" [1h]))'
```{{exec}}

This will return a similar JSON object above but will only show a trend of the number of documents sent from California in 1 hour intervals.

## Instant Metric Queries

Instant metric queries are a subset of metric queries that return the value of the metric at a specific point in time. This can be useful for quickly understanding an aggregate state of the stored logs.

For instance, we can use the following query to get the number of packages sent from California in the last 5 minutes:

```bash
logcli instant-query 'sum(count_over_time({state="California"}[5m]))'
```{{exec}}

This will return a result similar to the following:

```console
[
  {
    "metric": {},
    "value": [
      1732702998.725,
      "58"
    ]
  }
]
```{{copy}}

## Writing query results to a file

Another useful feature of LogCLI is the ability to write the query results to a file. This can be useful for downloading the results of our inventory report:

First we need to create a directory to store the logs:

```bash
mkdir -p ./inventory
```{{exec}}

Next we can run the following query to write the logs to the `./inventory`{{copy}} directory:

```bash
  logcli query \
     --timezone=UTC \
     --output=jsonl \
     --parallel-duration="12h" \
     --parallel-max-workers="4" \
     --part-path-prefix="./inventory/inv" \
     --since=24h \
     '{service_name="Delivery World"}'
```{{exec}}

This will write all logs for the `service_name`{{copy}} `Delivery World`{{copy}} in the last 24 hours to the `./inventory`{{copy}} directory. The logs will be split into two files, each containing 12 hours of logs. Note that we do not need to specify `--limit`{{copy}} as this is overridden by the `--parallel-duration`{{copy}} flag.
