# Setup

To get started, we need to clone the [Alloy Scenario](https://github.com/grafana/alloy-scenarios) repository and start the mail-house example:

1. Clone the repository:
   ```bash
   git clone https://github.com/grafana/alloy-scenarios.git
   ```{{exec}}
1. Start the mail-house example:
   ```bash
   docker compose -f alloy-scenarios/mail-house/docker-compose.yml up -d
   ```{{exec}}

This will start the mail-house example and expose the Loki instance at [`http://localhost:3100`{{copy}}]({{TRAFFIC_HOST1_3100}}). We have also included a Grafana instance to verify the LogCLI results which can be accessed at [http://localhost:3000]({{TRAFFIC_HOST1_3000}}).

## Connecting LogCLI to Loki

To connect LogCLI to the Loki instance, you need to set the `LOKI_ADDR`{{copy}} environment variable:

> **Tip:**
> If you are running this example against your own Loki instance and have configured authentication, you will also need to set the `LOKI_USERNAME`{{copy}} and `LOKI_PASSWORD`{{copy}} environment variables.

```bash
export LOKI_ADDR=http://localhost:3100
```{{exec}}

Now let's verify the connection by running the following command:

```bash
logcli labels
```{{exec}}

This should return an output similar to the following:

```console
http://localhost:3100/loki/api/v1/labels?end=1732282703894072000&start=1732279103894072000
package_size
service_name
state
```{{copy}}

This confirms that LogCLI is connected to the Loki instance and we now know that the logs contain the following labels: `package_size`{{copy}}, `service_name`{{copy}}, and `state`{{copy}}. Let's run some queries against Loki to better understand our package logistics.
