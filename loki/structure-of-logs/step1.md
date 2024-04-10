# Running the demo

## Step 1: Check out the application

In this demo, we will be working with a simple application that generates logs. We have created a small messaging board application similar to Hacker News or Reddit. The idea is users can post messages and upvote posts they like.

To access the application click this link to Grafana News: **[http://localhost:8081]({{TRAFFIC_HOST1_8081}})**

## Step 2: Generate some logs

Start to interact with the application by posting messages and upvoting posts. This will generate logs that we can explore in the next steps. 

**Top Tip:** *Try to post a message without a URL.*

## Step 3: Explore the logs

Our application generates logs in a specific format (a hybrid between plain text and structured since it utilises key-value pairs). The log is located in the `logs` directory of the application.

To print the logs, run the following command:

```bash
cat ./logs/tns-app.log
```{{exec}}



