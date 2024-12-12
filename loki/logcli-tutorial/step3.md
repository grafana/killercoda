# Meta queries

As site managers, it’s essential to maintain good data hygiene and ensure Loki operates efficiently. Understanding the labels and log volume in your logs plays a key role in this process. Beyond querying logs, LogCLI also supports meta queries on your Loki instance. Meta queries don’t return log data but provide insights into the structure of your logs and the performance of your queries. The following examples demonstrate some of the core meta queries we run internally to better understand how a Loki instance is performing.

## Checking series cardinality

One of the most important aspects of keeping Loki healthy is to monitor the series cardinality. This is the number of unique series in your logs. A high series cardinality can lead to performance issues and increased storage costs. We can use LogCLI to check the series cardinality of our logs.

To start let’s print how many unique series we have in our logs:

```bash
logcli series '{}'
```{{exec}}

This will return a list of all the unique series in our logs. The output will look similar to the following:

```console
{package_size="Small", service_name="Delivery World", state="Florida"}
{package_size="Medium", service_name="Delivery World", state="Florida"}
{package_size="Small", service_name="Delivery World", state="California"}
{package_size="Large", service_name="Delivery World", state="New York"}
{package_size="Small", service_name="Delivery World", state="Illinois"}
{package_size="Large", service_name="Delivery World", state="Florida"}
{package_size="Medium", service_name="Delivery World", state="Illinois"}
{package_size="Large", service_name="Delivery World", state="Texas"}
{package_size="Medium", service_name="Delivery World", state="California"}
{package_size="Medium", service_name="Delivery World", state="Texas"}
{package_size="Small", service_name="Delivery World", state="Texas"}
{package_size="Large", service_name="Delivery World", state="Illinois"}
{package_size="Small", service_name="Delivery World", state="New York"}
{package_size="Medium", service_name="Delivery World", state="New York"}
{package_size="Large", service_name="Delivery World", state="California"}
```{{copy}}

We can further improve this query by adding `--analyze-labels`{{copy}}:

```bash
logcli series '{}' --analyze-labels
```{{exec}}

This will return a summary of the unique values for each label in our logs. The output will look similar to the following:

```console
Label Name    Unique Values  Found In Streams
state         5              15
package_size  3              15
service_name  1              15
```{{copy}}

## Detected fields

Another useful feature of LogCLI is the ability to detect fields in your logs. This can be useful for understanding the structure of your logs and the keys that are present. This will let us detect keys which could be promoted to labels or to structured metadata.

```bash
logcli detected-fields --since 24h '{service_name="Delivery World"}'
```{{exec}}

This will return a list of all the keys detected in our logs. The output will look similar to the following:

```console
label: city                   type: string  cardinality: 15
label: detected_level         type: string  cardinality: 3
label: note                   type: string  cardinality: 7
label: package_id             type: string  cardinality: 994
label: package_size_extracted type: string  cardinality: 3
label: package_status         type: string  cardinality: 4
label: package_type           type: string  cardinality: 5
label: receiver_address       type: string  cardinality: 991
label: receiver_name          type: string  cardinality: 100
label: sender_address         type: string  cardinality: 991
label: sender_name            type: string  cardinality: 100
label: state_extracted        type: string  cardinality: 5
label: timestamp              type: string  cardinality: 1000
```{{copy}}

You can now see why we opted to keep `package_id`{{copy}} in structured metadata and `package_size`{{copy}} as a label. Package ID has a high cardinality and is unique to each log entry, making it a good candidate for structured metadata since we potentially may need to query for it directly. Package size, on the other hand, has a low cardinality, making it a good candidate for a label.

## Checking query performance

Another important aspect of keeping Loki healthy is to monitor the query performance. We can use LogCLI to check the query performance of our logs.

> **Note:**
> The LogCLI can only return statistics for queries that touch object storage. In this example we force the Loki ingesters to flush chunks every 5 minutes which isn’t recommended for production use. When running this demo if you don’t see any statistics returned, try running the command again after a few minutes.

To start lets print the query performance of our logs:

```bash
logcli stats --since 24h '{service_name="Delivery World"}'
```{{exec}}

This will provide a JSON object containing statistics on the amount of data queried. The output will look similar to the following:

```console
http://localhost:3100/loki/api/v1/index/stats?end=1732639430272850000&query=%7Bservice_name%3D%22Delivery+World%22%7D&start=1732553030272850000
{
  bytes: 12MB
  chunks: 63
  streams: 15
  entries: 29529
}
```{{copy}}

This will return the total number of bytes queried, the number of chunks queried, the number of streams queried, and the number of entries queried. If we narrow down the query by specifying a secondary label, we can see the performance of the query:

```bash
logcli stats --since 24h '{service_name="Delivery World", package_size="Large"}'
```{{exec}}

This will return the statistics for the logs where the `service_name`{{copy}} is `Delivery World`{{copy}} and the `package_size`{{copy}} is `Large`{{copy}}. The output will look similar to the following:

```console
{
  bytes: 4.2MB
  chunks: 22
  streams: 5
  entries: 10198
}
```{{copy}}

As you can see, we touched far fewer streams and entries by narrowing down the query.

## Checking the log volume

We may also want to check the log volume in our logs. This can be useful for understanding the amount of data being ingested into Loki. We can use LogCLI to check the log volume in our logs.

```bash
logcli volume --since 24h '{service_name="Delivery World"}'
```{{exec}}

This returns the total number of logs ingested for the label `Delivery World`{{copy}} in the last 24 hours. The output will look similar to the following:

```console
[
  {
    "metric": {
      "service_name": "Delivery World"
    },
    "value": [
      1732640292.354,
      "11669299"
    ]
  }
]
```{{copy}}

The result includes the timestamp and the total number of logs ingested.

We can also return the log volume over time by using `volume_range`{{copy}}:

```bash
logcli volume_range --since 24h --step=1h '{service_name="Delivery World"}'
```{{exec}}

This will provide a JSON object containing the log volume for the label `Delivery World`{{copy}} in the last 24 hours. `--step`{{copy}} will aggregate the log volume into 1 hour buckets. Note that if there are no logs for a specific hour, the log volume for that hour will not return a value.

We can even aggregate the log volume into buckets based on a specific labels value:

```bash
logcli volume_range --since 24h --step=1h --targetLabels='state' '{service_name="Delivery World"}' 
```{{exec}}

This will provide a similar JSON object but will aggregate the log volume into buckets based on the `state`{{copy}} label value.
