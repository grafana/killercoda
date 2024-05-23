# Installing Alloy

You can install Grafana Alloy as a systemd service on Linux.

## Before you begin

Some Debian-based cloud Virtual Machines don't have GPG installed by default.
To install GPG in your Linux Virtual Machine, run the following command in a terminal window.

```shell
sudo apt install gpg
```

We also need to spin up our local Grafana stack so alloy can write data to it. 

```bash
docker-compose -f /education/docker-compose.yml up -d
```

## Install

To install Grafana Alloy on Linux, run the following commands in a terminal window.

1. Import the GPG key and add the Grafana package repository.

   ```bash
   sudo mkdir -p /etc/apt/keyrings/
   wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null
   echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee /etc/apt/sources.list.d/grafana.list
  ```{{exec}}

1. Update the repositories.

   ```bash
   sudo apt-get update
   ```{{exec}}
2. Install Grafana Alloy.

   {{< code >}}
   ```debian-ubuntu
   sudo apt-get install alloy
   ```

3. Start the Grafana Alloy service.

   ```bash
   sudo systemctl start alloy
   ```

4. After starting the Alloy service, should now be able to access Alloy UI:
   [http://localhost:12345]({{TRAFFIC_HOST1_12345}})