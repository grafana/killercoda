# Install Alloy and start the service

This tutorial requires a Linux or macOS environment with Docker installed.

> This online sandbox enviroment is based on an Ubuntu image and has Docker pre-installed. To install Alloy simply follow the links below and copy and paste the `Ubuntu/Debian` commands in the terminal.
## Linux

Install and run Alloy on Linux.

1. [Install Alloy](https://grafana.com/docs/alloy/latest/set-up/install/linux/).

1. To view the Alloy UI within the sandbox, Alloy must run on all interfaces. Run the following command before you start the alloy service:
   ```bash
   sed -i -e 's/CUSTOM_ARGS=""/CUSTOM_ARGS="--server.http.listen-addr=0.0.0.0:12345"/' /etc/default/alloy
   ```{{exec}}

1. [Run Alloy](https://grafana.com/docs/alloy/latest/set-up/run/linux/).
