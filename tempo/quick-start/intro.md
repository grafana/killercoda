# Quick start for Tempo

The Tempo repository provides [multiple examples](https://github.com/grafana/tempo/tree/main/example/docker-compose) to help you quickly get started using Tempo and distributed tracing data.

Every example has a `docker-compose.yaml`{{copy}} manifest that includes all of the options needed to explore trace data in Grafana, including resource configuration and trace data generation.

The Tempo examples running with Docker using docker-compose include a version of Tempo and a storage configuration suitable for testing or development.

This quick start guide uses the `local`{{copy}} example running Tempo as a single binary (monolithic). Any data is stored locally in the `tempo-data`{{copy}} folder.
To learn more, read the [local storage example README](https://github.com/grafana/tempo/blob/main/example/docker-compose/local).
