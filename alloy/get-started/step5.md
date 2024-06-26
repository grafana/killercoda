# Reload the configuration

Copy your local `config.alloy`{{copy}} file into the default configuration file location, which depends on your OS.

{{< code >}}

```macos
sudo cp config.alloy $(brew --prefix)/etc/alloy/config.alloy
```{{copy}}

```linux
sudo cp config.alloy /etc/alloy/config.alloy
```{{copy}}

{{< /code >}}

Finally, call the reload endpoint to notify {{< param “PRODUCT_NAME” >}} to the configuration change without the need
for restarting the system service.

```bash
    curl -X POST http://localhost:12345/-/reload
```{{copy}}

{{< admonition type=“tip” >}}
This step uses the Alloy UI, which is exposed on `localhost`{{copy}} port `12345`{{copy}}.
If you chose to run Alloy in a Docker container, make sure you use the `--server.http.listen-addr=0.0.0.0:12345`{{copy}} argument.
If you don’t use this argument, the [debugging UI](https://grafana.com/../../tasks/debug/#alloy-ui) won’t be available outside of the Docker container.

{{< /admonition >}}

The alternative to using this endpoint is to reload the {{< param “PRODUCT_NAME” >}} configuration, which can
be done as follows:

{{< code >}}

```macos
brew services restart alloy
```{{copy}}

```linux
sudo systemctl reload alloy
```{{copy}}

{{< /code >}}
