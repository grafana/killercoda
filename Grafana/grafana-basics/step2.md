# Step 2: Installing Grafana
In this step, we will install Grafana on within our virtual environment. Grafana is an open-source platform for monitoring and observability that allows you to create, explore, and share dashboards and data visualizations.

## Installing Grafana
1. Lets install Grafana via apt install.
    ```
    sudo apt-get install -y apt-transport-https software-properties-common wget &&
    sudo mkdir -p /etc/apt/keyrings/ &&
    wget -q -O - https://apt.grafana.com/gpg.key | gpg --dearmor | sudo tee /etc/apt/keyrings/grafana.gpg > /dev/null &&
    echo "deb [signed-by=/etc/apt/keyrings/grafana.gpg] https://apt.grafana.com stable main" | sudo tee -a /etc/apt/sources.list.d/grafana.list &&
    sudo apt-get update && sudo apt-get install -y grafana
    ```{{execute}}

2. Lets make sure our Grafana service is running.
    ```
    sudo systemctl start grafana-server && sudo systemctl status grafana-server
    ```{{execute}}

We should now be able to access Grafana by visiting the following URL in your browser: [http://localhost:3000](TRAFFIC_HOST1_3000). If you see a page with a login prompt, then Grafana is running correctly.
