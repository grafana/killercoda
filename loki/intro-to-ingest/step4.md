# Step 4: Alloy Migration Tool

Lastly we are going to look at an example of how you can use the Alloy Migration tool to migrate your existing:
* Promtail
* OpenTelemetry Collector
* Grafana Agent
Configurations to Alloy..

## Alloy Migration Tool (Promtail)

We have an example promtail config located in `promtail/config.yml`. We can use the Alloy Migration tool to convert this to an Alloy config.

Lets first take a look at the promtail config:
```bash
cat ./promtail/config.yml
```{{exec}}

Next lets run the Alloy Migration tool to convert this to an Alloy config:
```bash
alloy convert --source-format=promtail --output=./promtail/config.alloy ./promtail/config.yml
```{{exec}}

Now we can take a look at the converted Alloy config:
```bash
cat ./promtail/config.alloy
```{{exec}}



