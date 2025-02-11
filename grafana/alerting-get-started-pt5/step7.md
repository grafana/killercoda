# Create mute timings

Now that we’ve set up notification policies, we can demonstrate how to mute alerts for recurring periods of time. You can mute notifications for either the production or staging policies, depending on your needs.

Mute timings are useful for suppressing alerts with certain labels during maintenance windows or weekends.

1. Navigate to **Alerts & IRM > Alerting > Notification Policies**.
   - Enter a name. E.g., `Planned downtime`{{copy}} , or `Non-business hours`{{copy}}.

   - Select **Sat** and **Sun**, to apply the mute timing to all Saturdays and Sundays.

   - Click **Save mute timing**.

1. Add mute timing to the desired policy:
   - Go to the notification policy that routes instances with the `staging`{{copy}} label.

   - Select **More > Edit**.

   - Choose the mute timing from the drop-down menu

   - Click **Update policy**.

This mute timing will apply to any alerts from the staging environment that trigger on Saturdays and Sundays.
