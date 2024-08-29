# Set up the sample application

This tutorial uses a sample application to demonstrate some of the features in Grafana. To complete the exercises in this tutorial, you need to download the files to your local machine.

In this step, you’ll set up the sample application, as well as supporting services, such as [Loki](https://grafana.com/oss/loki/).

> **Note:** [Prometheus](https://prometheus.io/), a popular time series database (TSDB), has already been configured as a data source as part of this tutorial.
1. Clone the [github.com/grafana/tutorial-environment](https://github.com/grafana/tutorial-environment) repository.

   ```bash
   git clone https://github.com/grafana/tutorial-environment.git
   ```{{exec}}

1. Change to the directory where you cloned this repository:

   ```bash
   cd tutorial-environment
   ```{{exec}}

1. Make sure Docker is running:

   ```bash
   docker ps
   ```{{exec}}

   No errors means it is running. If you get an error, then start Docker and then run the command again.

1. Start the sample application:

   ```bash
   docker-compose up -d
   ```{{exec}}

   The first time you run `docker-compose up -d`{{copy}}, Docker downloads all the necessary resources for the tutorial. This might take a few minutes, depending on your internet connection.

   > **Note:** If you already have Grafana, Loki, or Prometheus running on your system, then you might see errors because the Docker image is trying to use ports that your local installations are already using. Stop the services, then run the command again.

1. Ensure all services are up-and-running:

   ```bash
   docker-compose ps
   ```{{exec}}

   In the **State** column, it should say `Up`{{copy}} for all services.

1. Browse to the sample application on [http://localhost:8081]({{TRAFFIC_HOST1_8081}}).

## Grafana News

The sample application, Grafana News, lets you post links and vote for the ones you like.

To add a link:

1. In **Title**, enter **Example**.

1. In **URL**, enter **<https://example.com>**.

1. Click **Submit** to add the link.

   The link appears in the list under the Grafana News heading.

To vote for a link, click the triangle icon next to the name of the link.
