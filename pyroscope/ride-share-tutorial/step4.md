> **Tip:**
> Unfortunately, due to a bug within the Sandbox environment, the Explore Profiles app is currently unavailable. We are working on a fix and will update this tutorial once resolved. If you would like to try out Explore Profiles, you can run the example locally on your machine. Or you can try out this example in [Grafana Play](https://play.grafana.org/a/grafana-pyroscope-app/profiles-explorer?searchText=&panelType=time-series&layout=grid&hideNoData=off&explorationType=labels&var-serviceName=pyroscope-rideshare-python&var-profileMetricId=process_cpu:cpu:nanoseconds:cpu:nanoseconds&var-dataSource=grafanacloud-profiles&var-groupBy=all&var-filters=)

# Integrating Pyroscope with Grafana

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
