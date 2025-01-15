1. Run Pyroscope.

   In a terminal, run one of the following commands:

   - Using Docker:

     ```bash
     docker network create pyroscope-demo
     docker run --rm --name pyroscope --network=pyroscope-demo -p 4040:4040 grafana/pyroscope:latest
     ```{{exec}}

   - Using a local binary:

     ```bash
     ./pyroscope
     ```{{exec}}

1. Verify that Pyroscope is ready. Pyroscope listens on port `4040`{{copy}}.

   ```bash
   curl localhost:4040/ready
   ```{{exec}}

1. Configure Pyroscope to scrape profiles.

   By default, Pyroscope is configured to scrape itself.
   To collect more profiles, you must either instrument your application with an SDK or use Grafana Alloy.

   To learn more about language integrations and the Pyroscope agent, refer to [Pyroscope Agent]({{< relref “../configure-client/_index.md” >}}).
