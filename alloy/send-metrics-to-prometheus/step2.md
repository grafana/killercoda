# Reload the configuration

Copy your local `config.alloy`{{copy}} file into the default Alloy configuration file location.

```bash
sudo cp config.alloy /etc/alloy/config.alloy
```{{exec}}

Call the `/-/reload`{{copy}} endpoint to tell Alloy to reload the configuration file without a system service restart.

```bash
    curl -X POST http://localhost:12345/-/reload
```{{exec}}

> This step uses the Alloy UI on `localhost` port `12345`. If you chose to run Alloy in a Docker container, make sure you use the `--server.http.listen-addr=` argument. If you don’t use this argument, the [debugging UI](https://grafana.com/docs/alloy/latest/troubleshoot/debug/#alloy-ui) won’t be available outside of the Docker container.
1. Optional: You can do a system service restart Alloy and load the configuration file:

   ```bash
   sudo systemctl reload alloy
   ```{{exec}}
