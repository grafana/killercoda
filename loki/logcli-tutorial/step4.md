# Queries against static log files

In addition to querying logs from Loki, LogCLI also supports querying static log files. This can be useful for querying logs that are not stored in Loki. Earlier in the tutorial we stored the logs in the `./inventory`{{copy}} directory. Lets run a similar query but pipe it into a log file:

```bash
  logcli query \
     --timezone=UTC \
     --parallel-duration="12h" \
     --parallel-max-workers="4" \
     --part-path-prefix="./inventory/inv" \
     --since=24h \
     --merge-parts \
     --output=raw \
     '{service_name="Delivery World"}' > ./inventory/complete.log
```{{exec}}

Next lets run a query against the static log file:

```bash
cat ./inventory/complete.log |  logcli --stdin query '{service_name="Delivery World"} | json | package_status="critical"'
```{{exec}}

Note that since we are querying a static log file, labels are not automatically detected:

- `{service_name="Delivery World"}`{{copy}} is optional in this case but is recommended for clarity.
- `json`{{copy}} is required to parse the log file as JSON. This lets us extract the `package_status`{{copy}} field.

For example, suppose we try to query the log file without the `json`{{copy}} filter:

```bash
cat ./inventory/complete.log | logcli --stdin query '{service_name="Delivery World"} | package_status="critical"'
```{{exec}}

This will return no results as the `package_status`{{copy}} field is not detected.
