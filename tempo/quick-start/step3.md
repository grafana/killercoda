# Explore Traces Plugin

Another method to explore traces within Grafana is via the [Explore Traces](https://grafana.com/docs/grafana/latest/explore/simplified-exploration/traces/) plugin. This plugin offers an opinionated non query-based approach to exploring traces. Lets take a look at some of its key features and panels.

1. Open a browser and navigate to [http://localhost:3000/a/grafana-exploretraces-app]({{TRAFFIC_HOST1_3000}}/a/grafana-exploretraces-app).

1. Within the filter bar there is a dropdown menu set to `Rate`{{copy}} of `Full traces`{{copy}}. Change this to `Duration`{{copy}} and `All spans`{{copy}}.

This should provide an updated panel view like this.

![Explore Traces panel](https://grafana.com/media/docs/tempo/explore-spans-error-view.png)

Breakdown of the view:

- The histogram at the top shows the distribution of span durations. The lighter the color, the more spans in that duration bucket. In this example most spans fall within `537ms`{{copy}} which is considered the average duration for the system.

- There are some high peaks in the histogram, which indicate spans that are taking longer than the average (As high as `2.15s`{{copy}}). These are likely to be the spans that are causing performance issues. Lets investigate further to see if we can identify the root cause.

Select `Slow traces`{{copy}} tab in the navigation bar to view the slowest traces in the system.

![Slow traces panel](https://grafana.com/media/docs/tempo/slow-trace-view.png)

`shop-backend`{{copy}} appears to be the primary culprit for the slow traces. This happens when a user initiates the `article-to-cart`{{copy}} operation. From here we can select the Trace Name which will open the Trace View panel.

![Trace View panel](https://grafana.com/media/docs/tempo/slow-trace-trace-view.png)

The Trace View panel provides a detailed view of the trace. The panel is divided into three sections:

- The top section shows the trace ID, duration, and the service that generated the trace.

- The middle section shows the trace timeline. Each span is represented as a horizontal bar. The color of the bar represents the span’s status. The width of the bar represents the duration of the span.

- The bottom section shows the details of the selected span. This includes the span name, duration, and tags.

Drilling into the `shop-backend`{{copy}} span, we can see that the `place-articles`{{copy}} operation has an exception event tied to it. This is likely the root cause of the slow trace.

![Span View panel](https://grafana.com/media/docs/tempo/slow-trace-root-cause-2.png)

If you would like to dive deeper into the Explore Traces plugin and its panel concepts, refer to the [Explore Traces Concepts](https://grafana.com/docs/grafana/latest/explore/simplified-exploration/traces/concepts/).
