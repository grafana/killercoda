# Install Alloy and start the service

> This online sandbox environment is based on an Ubuntu image and has Docker pre-installed. To install Alloy in the sandbox, perform the following steps.
## Linux

Install and run Alloy on Linux.

1. [Install Alloy](https://grafana.com/docs/alloy/latest/set-up/install/linux/).

1. To view the Alloy UI within the sandbox, Alloy must run on all interfaces. Run the following command before you start the Alloy service.
   ```bash
   sed -i -e 's/CUSTOM_ARGS=""/CUSTOM_ARGS="--server.http.listen-addr=0.0.0.0:12345"/' /etc/default/alloy
   ```{{exec}}

1. [Run Alloy](https://grafana.com/docs/alloy/latest/set-up/run/linux/).

You should now be able to access the Alloy UI at [http://localhost:12345]({{TRAFFIC_HOST1_12345}}).
