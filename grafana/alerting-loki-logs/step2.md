# Generate sample logs

To demonstrate how to create alert rules based on logs, you’ll use a script that generates realistic log entries to simulate typical monitoring data in Grafana. Running this script outputs logs continuously, each containing a timestamp, HTTP method (either GET or POST), status code (200 for success or 500 for failures), and request duration in milliseconds.

1. Download and save a Python file that generates logs.

   ```bash
   wget https://raw.githubusercontent.com/grafana/tutorial-environment/master/app/loki/web-server-logs-simulator.py
   ```{{exec}}

1. Execute the log-generating Python script.

   ```bash
   python3 ./web-server-logs-simulator.py | sudo tee -a /var/log/web_requests.log
   ```{{exec}}

## Troubleshooting the script

If you don’t see the sample logs in Explore:

- Does the output file exist, check `/var/log/web_requests.log`{{copy}} to see if it contains logs.

- If the file is empty, check that you followed the steps above to create the file.

- If the file exists, verify that promtail container is running.

- In Grafana Explore, check that the time range is only for the last 5 minutes.
