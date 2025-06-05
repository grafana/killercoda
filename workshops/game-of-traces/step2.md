# Setting Up the Dashboard

1. Log into Grafana at <http://localhost:3000> (default credentials: admin/admin)

1. Import the dashboard:

   - Click the “+” icon in the left sidebar

   - Select “Import dashboard”

   - Click “Upload JSON file”

   - Navigate to `grafana/dashboards/War of Kingdoms-1747821967780.json`{{copy}}

   - Click “Import”

1. Configure data sources:

   - The dashboard requires Prometheus, Loki, and Tempo data sources

   - These should be automatically configured if you’re using the provided Docker setup

   - If not, ensure the following URLs are set:
     - Prometheus: <http://prometheus:9090>

     - Loki: <http://loki:3100>

     - Tempo: <http://tempo:3200>

1. The dashboard provides:

   - Real-time army and resource metrics

   - Battle analytics

   - Territory control visualization

   - Service dependency mapping

   - Trace analytics for game events
