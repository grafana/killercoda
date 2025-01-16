# Receiving grouped alert notifications

Now that the alert rules have been configured, you should receive alert notifications in the contact point(s) whenever alerts trigger.

When the configured alert rule detects CPU or memory usage higher than 75% across multiple regions, it will evaluate the metric every minute. If the condition persists, notifications will be grouped together, with a Group wait of 30 seconds before the first alert is sent. Follow-up notifications for the same alert group will be sent at intervals of 2 minutes (US-west alert instances only), increasing the frequency of the grouped alert notifications. US-east instances follow-up notifications should be sent at the default interval of 5 minutes. If the condition continues for an extended period, a Repeat interval of 4 hours ensures that the alert is only resent if the issue persists.

As a result, our notification policies should route three notifications: one grouped notification grouping both CPU and memory alert instances from the us-west region and two separate notifications with alert instances from the us-east region.

Grouped notifications example:

```json
{
  "receiver": "US-West-Alerts",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "High CPU usage - Multi-region",
        "grafana_folder": "Multi-region alerts",
        "instance": "server-05",
...
  {
    "status": "firing",
      "labels": {
        "alertname": "High Memory usage - Multi-region",
        "grafana_folder": "Multi-region alerts",
        "instance": "server-10",
      },

...}
```{{copy}}

_Detail of CPU and memory alert instances grouped into a single notification for us-west contact point._

```json
{
  "receiver": "US-East-Alerts",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "High CPU usage - Multi-region",
        "grafana_folder": "Multi-region alerts",
        "instance": "server-03",
        "region": "us-east",
        "service": "web-server-2"
...}}}
```{{copy}}

_Detail of CPU alert instances grouped into a separate notification for us-east contact point._

```json
{
  "receiver": "US-East-Alerts",
  "status": "firing",
  "alerts": [
    {
      "status": "firing",
      "labels": {
        "alertname": "High memory usage - Multi-region",
        "grafana_folder": "Multi-region memory alerts",
        "instance": "server-12",
        "region": "us-east"
...}}}
```{{copy}}

_Detail of memory alert instances grouped into a separate notification for us-east contact point._
