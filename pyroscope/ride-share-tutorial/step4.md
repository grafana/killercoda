# How was Pyroscope integrated with Grafana in this tutorial?

The `docker-compose.yml`{{copy}} file includes a Grafana container that’s pre-configured with the Pyroscope plugin:

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

Grafana is also pre-configured with the Pyroscope data source.

## Challenge

As a challenge, see if you can generate a similar comparison with the `vehicle`{{copy}} tag.
