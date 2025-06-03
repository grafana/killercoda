# (Optional) Step 3: Create a second alert rule for memory usage

1. Duplicate the existing alert rule (**More > Duplicate**), or create a new alert rule for memory usage, defining a threshold condition (e.g., memory usage exceeding `60%`{{copy}}).
1. Give it a name. For example: `memory-usage`{{copy}}
1. Query: `flask_app_memory_usage{instance="flask-prod:5000"}`{{copy}}
1. Link to the same visualization to obtain memory usage annotations

Check how your dashboard looks now that both alerts have been linked to your dashboard panel.
