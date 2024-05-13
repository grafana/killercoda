
# Step 1: Generating some logs

In this step, we will generate some logs using the Carnivorous Greenhouse application. This application generates logs when you create an account, login, and collect metrics from a series of hungry plants. The logs are sent to Loki. 

To access the application, click this link to Carnivorous Greenhouse: **[http://localhost:5005]({{TRAFFIC_HOST1_5005}})**

## Generate some logs

1. Create an account by clicking the **Sign Up** button.
2. Create a user by entering a username and password.
3. Log in by clicking the **Login** button.
4. Create your first plant by:
   * Adding a name
   * Selecting a plant type
   * Clicking the **Add Plant** button.

### Optional: Generate errors

The plants do a good job eating all the bugs, but sometimes they get a little too hungry and cause errors. To generate an error, toggle the **Toggle Error mode**. This will cause a variety of errors to occure such as; sign up errors, login errors, and plant creation errors. These errors can be investigated in the logs.


