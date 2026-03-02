# Step 2: Template notifications

In this step, we use a built-in notification template to format alert notifications in a clear and organized way. Notification templates allow us to customize the structure of alert messages, making them easier to read and more relevant.

Without a notification template, the alert messages would include the default Grafana formatting (`default.message`{{copy}}, see image above).

## Adding a notification template:

1. Navigate to **Alerts & IRM** > **Alerting** > **Contact point**s.
1. Select the **Notification Templates** tab.
1. Click **+ New notification template**.
1. Enter a name. E.g `instance-cpu-summary`{{copy}}.
1. From the **Add example** dropdown menu, choose `Print firing and resolved alerts`{{copy}}.

This template prints out alert instances into two sections: **firing alerts** and **resolved alerts**, and includes only the key details for each. In addition, it adds our summary and description annotations.

```
{{- /* Example displaying firing and resolved alerts separately in the notification. */ -}}
{{- /* Edit the template name and template content as needed. */ -}}
{{ define "custom.firing_and_resolved_alerts" -}}
{{ len .Alerts.Resolved }} resolved alert(s)
{{ range .Alerts.Resolved -}}
  {{ template "alert.summary_and_description" . -}}
{{ end }}
{{ len .Alerts.Firing }} firing alert(s)
{{ range .Alerts.Firing -}}
  {{ template "alert.summary_and_description" . -}}
{{ end -}}
{{ end -}}
{{ define "alert.summary_and_description" }}
  Summary: {{.Annotations.summary}}
  Status: {{ .Status }}
  Description: {{.Annotations.description}}
{{ end -}}
```{{copy}}

Note: Your notification template name (`{{define "<NAME>"}}`{{copy}}) must be unique. You cannot have two templates with the same name in the same notification template group or in different notification template groups.

Here’s a breakdown of the template:

- `{{ define "custom.firing_and_resolved_alerts" -}}`{{copy}} section: Displays the number of resolved alerts and their summaries, using the `alert.summary_and_description`{{copy}} template to include the summary, status, and description for each alert.
- `.Alerts.Firing`{{copy}} section: Similarly lists the number of firing alert instances and their details.
- `alert.summary_and_description`{{copy}}: This sub-template pulls the summary annotation you configured earlier.

In the **Preview** area, you can see a sample of how the notification would look. Since we’ve already created our alert rule, you can take it a step further by previewing how an actual alert instance from your rule would appear in the notification.

1. Click **Edit payload**.
1. Click **Use existing alert instance**.

   You should see our alert rule listed on the left.
1. Click the alert rule.
1. Select an instance.
1. Click **Add alert data to payload**.

   The alert instance is added to the bottom of the preview.

   ![Preview of an alert instance in a notification template](https://grafana.com/media/docs/alerting/alert-instance-preview-in-template.png)
1. Click **Save**.

With the notification template ready, the next step is to apply it to your contact point to see it in action.
