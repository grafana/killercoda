# Fourth component: Write logs to Loki

Paste this component last in your configuration file:

```alloy
loki.write "grafana_loki" {
  endpoint {
    url = "http://localhost:3100/loki/api/v1/push"

    // basic_auth {
    //  username = "admin"
    //  password = "admin"
    // }
  }
}
```{{copy}}

This last component creates a [loki.write](https://grafana.com/../../reference/components/loki.write/) component named `grafana_loki`{{copy}} that points to `http://localhost:3100/loki/api/v1/push`{{copy}}.
This completes the simple configuration pipeline.

{{< admonition type=“tip” >}}
The `basic_auth`{{copy}} is commented out because the local `docker compose`{{copy}} stack doesn’t require it.
It is included in this example to show how you can configure authorization for other environments.
For further authorization options, refer to the [loki.write](https://grafana.com/../../reference/components/loki.write/) component reference.

{{< /admonition >}}

This connects directly to the Loki instance running in the Docker container.
