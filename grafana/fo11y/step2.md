# Creating the Frontend Observability App in Grafana Cloud

Now that you have both the frontend and backend applications running, the next thing we need to do is create the Frontend Observability app in your Grafana Cloud instance.
If you haven't done so already, open your Grafana Cloud instance.

*Note: If you haven't created a Grafana Cloud instance yet, do that now by [signing in or registering](https://grafana.com/auth/sign-in/) and creating a new stack.*

## Adding the Frontend Observability application in Grafana Cloud

On the left hand side navigation, expand the **Frontend** section and click on **Frontend Apps**.

Next, in the right hand panel, click on the **Create New** button.

Fill in the following information:

- **App Name:** (e.g. fo11y-demo)
- **CORS Allowed Origins:** {{TRAFFIC_HOST1_3000}}
- **Default Attributes:** Add an attribute with a name of `application_name` with a value of `fo11y-demo`
- **Acknowledge cloud costs:** Tick the box

Then click on **Create**.

In the next section, we'll instrument our React application to begin pushing data to our Grafana Cloud instance.