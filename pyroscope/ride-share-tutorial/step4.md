> **Tip:**
> Unfortunately, due to a bug within the Sandbox environment, the profile explorer app is currently unavailable. We are working on a fix and will update this tutorial once resolved. If you would like to try out the profile explorer app, you can run the example locally on your machine.

# Integrating Pyroscope with Grafana

As part of the `docker-compose.yml`{{copy}} file, we have included a Grafana container that’s pre-configured with the Pyroscope plugin:

```yaml
  grafana:
    image: grafana/grafana:latest
    environment:
    - GF_INSTALL_PLUGINS=grafana-pyroscope-app
    - GF_AUTH_ANONYMOUS_ENABLED=true
    - GF_AUTH_ANONYMOUS_ORG_ROLE=Admin
    - GF_AUTH_DISABLE_LOGIN_FORM=true
    volumes:
    - ./grafana-provisioning:/etc/grafana/provisioning
    ports:
    - 3000:3000
```{{copy}}

We’ve also pre-configured the Pyroscope data source in Grafana.

To access the Pyroscope app in Grafana, navigate to [http://localhost:3000/a/grafana-pyroscope-app]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app).

## Challenge

As a challenge see if you can generate the same comparison we achieved in the Pyroscope UI within Grafana. It should look something like this:

![Grafana](https://grafana.com/media/docs/pyroscope/ride-share-grafana.png)
