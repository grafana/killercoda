# Extracting attributes from logs

Loki by design does not force log lines into a specific schema format. Whether you are using JSON, key-value pairs, plain text, Logfmt, or any other format, Loki ingests these logs lines as a stream of characters. The sample application we are using stores logs in [Logfmt](https://brandur.org/logfmt) format:

```bash
ts=2025-02-21 16:09:42,176 level=INFO line=97 msg="192.168.65.1 - - [21/Feb/2025 16:09:42] "GET /static/style.css HTTP/1.1" 304 -"
```{{copy}}

To break this down:

- `ts=2025-02-21 16:09:42,176`{{copy}} is the timestamp of the log line.

- `level=INFO`{{copy}} is the log level.

- `line=97`{{copy}} is the line number in the code.

- `msg="192.168.65.1 - - [21/Feb/2025 16:09:42] "GET /static/style.css HTTP/1.1" 304 -"`{{copy}} is the log message.

When querying Loki, you can pipe the result of the label selector through a parser. This extracts attributes from the log line for further processing. For example, lets pipe `{container="greenhouse-main_app-1"}`{{copy}} through the `logfmt`{{copy}} parser to extract the `level`{{copy}} and `line`{{copy}} attributes:

```logql
{container="greenhouse-main_app-1"} | logfmt
```{{copy}}

When you now expand a log line in the query result, you will see the extracted attributes.

> **Tip:**
> Before we move on to the next section, let’s generate some error logs. To do this, enable the bug service in the sample application. This is done by setting the `Toggle Error Mode`{{copy}} to `On`{{copy}} in the Carnivorous Greenhouse application. This will cause the bug service to randomly cause services to fail.

# Advanced and Metrics Queries

With Error Mode enabled the bug service will start causing services to fail, in these next few LogQL examples we will track down some of these errors.  Lets start by parsing the logs to extract the `level`{{copy}} attribute and then filter for logs with a `level`{{copy}} of `ERROR`{{copy}}:

```logql
{container="greenhouse-plant_service-1"} | logfmt | level="ERROR"
```{{copy}}

This query will return all the logs from the `greenhouse-plant_service-1`{{copy}} container that have a `level`{{copy}} attribute of `ERROR`{{copy}}. You can further refine this query by filtering for a specific code line:

```logql
{container="greenhouse-plant_service-1"} | logfmt | level="ERROR", line="58"
```{{copy}}

This query will return all the logs from the `greenhouse-plant_service-1`{{copy}} container that have a `level`{{copy}} attribute of `ERROR`{{copy}} and a `line`{{copy}} attribute of `58`{{copy}}.

LogQL also supports metrics queries. Metrics are useful for abstracting the raw log data aggregating attributes into numeric values. This allows you to utilise more visualization options in Grafana as well as generate alerts on your logs.

For example, you can use a metric query to count the number of logs per second that have a specific attribute:

```logql
sum(rate({container="greenhouse-plant_service-1"} | logfmt | level="ERROR" [$__auto]))
```{{copy}}

It worth changing the visualization from `lines`{{copy}} to `bars`{{copy}} to visualize the error rate over time since the error count is quite low.

Another example is to get the top 10 services producing the highest rate of errors:

```logql
topk(10,sum(rate({level="error"} | logfmt [5m])) by (service_name))
```{{copy}}

> **Note:**
> `service_name`{{copy}} is a label created by Loki when no service name is provided in the log line. It will use the container name as the service name. A list of all automatically generated labels can be found in [Labels](https://grafana.com/docs/loki/latest/get-started/labels/#default-labels-for-all-users).

Finally, lets take a look at the total log throughput of each container in our production environment:

```logql
sum by (service_name) (rate({env="production"} | logfmt [$__auto]))
```{{copy}}

This is made possible by the `service_name`{{copy}} label and the `env`{{copy}} label that we have added to our log lines. Note that `env`{{copy}} is a static label that we added to all log lines as they are processed by Alloy.
