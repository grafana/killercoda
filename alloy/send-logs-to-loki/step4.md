# Reload the configuration

1. Copy your local `config.alloy`{{copy}} file into the default Alloy configuration file location.

   ```bash
    sudo cp config.alloy /etc/alloy/config.alloy
   ```{{exec}}

1. Call the `/-/reload`{{copy}} endpoint to tell Alloy to reload the configuration file without a system service restart.

   ```bash
    curl -X POST http://localhost:12345/-/reload
   ```{{exec}}

> This step uses the Alloy UI on `localhost` port `12345`. If you chose to run Alloy in a Docker container, make sure you use the `--server.http.listen-addr=` argument. If you don’t use this argument, the [debugging UI][debug] won’t be available outside of the Docker container.
1. Optional: You can do a system service restart Alloy and load the configuration file:

   ```bash
    sudo systemctl reload alloy
   ```{{exec}}

# Inspect your configuration in the Alloy UI

Open [http://localhost:12345]({{TRAFFIC_HOST1_12345}}) and click the **Graph** tab at the top.
The graph should look similar to the following:

![Your configuration in the Alloy UI](https://grafana.com/media/docs/alloy/tutorial/Inspect-your-config-in-the-Alloy-UI-image.png)

The Alloy UI shows you a visual representation of the pipeline you built with your Alloy component configuration.

You can see that the components are healthy, and you are ready to explore the logs in Grafana.
