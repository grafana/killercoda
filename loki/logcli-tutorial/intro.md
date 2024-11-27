# LogCLI tutorial

This [LogCLI](https://grafana.com/docs/loki/%3CVERSION%3E/query/logcli/) tutorial will walk you through the following concepts:

- Querying logs

- Meta Queries against your Loki instance

- Queries against static log files

## Scenario

You are a site manager for a new logistics company. The company uses structured logs to keep track of every shipment sent and received. The payload format looks like this:

```json
{"timestamp": "2024-11-22T13:22:56.377884", "state": "New York", "city": "Buffalo", "package_id": "PKG34245", "package_type": "Documents", "package_size": "Medium", "package_status": "error", "note": "Out for delivery", "sender": {"name": "Sender27", "address": "144 Elm St, Buffalo, New York"}, "receiver": {"name": "Receiver4", "address": "260 Cedar Blvd, New York City, New York"}}
```{{copy}}

The logs are processed from Grafana Alloy to extract labels and structured metadata before they’re stored in Loki. You have been tasked with monitoring the logs using the LokiCLI and build a report on the overall health of the shipments.
