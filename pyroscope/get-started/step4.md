# Add a Pyroscope data source and query data

1. In a terminal, run a local Grafana server using Docker:

   ```bash
   docker run --rm --name=grafana \
     --network=pyroscope-demo \
     -p 3000:3000 \
     -e "GF_INSTALL_PLUGINS=grafana-pyroscope-app"\
     -e "GF_AUTH_ANONYMOUS_ENABLED=true" \
     -e "GF_AUTH_ANONYMOUS_ORG_ROLE=Admin" \
     -e "GF_AUTH_DISABLE_LOGIN_FORM=true" \
     grafana/grafana:main
   ```{{exec}}

1. In a browser, go to the Grafana server at [http://localhost:3000/datasources]({{TRAFFIC_HOST1_3000}}/datasources).

1. Use the following settings to configure a Pyroscope data source to query the local Pyroscope server:

|    Field | Value |
| --- | --- |
| Name | Pyroscope |
| URL | [http://pyroscope:4040/](http://pyroscope:4040/) OR [http://host.docker.internal:4040/](http://host.docker.internal:4040/) if using Docker |


To learn more about adding data sources, refer to [Add a data source](https://grafana.com/docs/grafana/latest/datasources/add-a-data-source/).

1. In a browser, go to [Explore Profiles](https://grafana.com/docs/grafana/latest/explore/simplified-exploration/profiles/) in your Grafana instance at [https://localhost:3000/a/grafana-pyroscope-app/profiles-explorer]({{TRAFFIC_HOST1_3000}}/a/grafana-pyroscope-app/profiles-explorer). This will let you use an intuitive interface for exploring your profile data.
