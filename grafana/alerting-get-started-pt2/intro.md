This tutorial is a continuation of the [Grafana Alerting - Create and receive your first alert](http://www.grafana.com/tutorials/alerting-get-started/) tutorial.

In this guide, we dig into more complex yet equally fundamental elements of Grafana Alerting: **alert instances** and **notification policies**.

{{< youtube id="nI-_MEnFBQs" >}}

After introducing each component, you will learn how to:

- Configure an alert rule that returns more than one alert instance
- Create notification policies that route firing alert instances to different contact points
- Use labels to match alert instances and notification policies

Learning about alert instances and notification policies is useful if you have more than one contact point in your organization, or if your alert rule returns a number of metrics that you want to handle separately by routing each alert instance to a specific contact point. The tutorial will introduce each concept, followed by how to apply both concepts in a real-world scenario.
