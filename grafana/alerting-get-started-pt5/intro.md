The Get started with Grafana Alerting tutorial Part 5 is a continuation of [Get started with Grafana Alerting tutorial Part 4](http://www.grafana.com/tutorials/alerting-get-started-pt4/).

In this tutorial, we focus on optimizing your alerting strategy using Grafana for monitoring system health, particularly when working with [Prometheus](https://grafana.com/docs/grafana/latest/datasources/prometheus/). Imagine you are managing a web application or a fleet of servers, tracking critical metrics such as CPU, memory, and disk usage. While monitoring is essential, managing alerts allows your team to act on issues without necessarily feeling overwhelmed by the noise.

In this tutorial you will learn how to:

- Leverage notification policies for **dynamic routing based on query values**: Use notification policies to route alerts based on dynamically generated labels, in a way that critical alerts reach the on-call team and less urgent ones go to a general monitoring channel.

- Set **mute timings** to suppress certain alerts during maintenance, or weekends.

- **Link alerts to dashboards** to provide more context to resolve issues.
