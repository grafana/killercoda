# Running the Demo

1. Clone the repository:

   ```bash
   git clone https://github.com/grafana/alloy-scenarios.git
   cd alloy-scenarios
   ```{{exec}}

1. Navigate to this example:

   ```bash
   cd game-of-tracing
   ```{{exec}}

1. Run using Docker Compose:

   ```bash
   docker compose up -d
   ```{{exec}}

1. Access the components:

   - Game UI: [http://localhost:8080]({{TRAFFIC_HOST1_8080}})

   - Grafana: [http://localhost:3000]({{TRAFFIC_HOST1_3000}})

   - Prometheus: [http://localhost:9090]({{TRAFFIC_HOST1_9090}})

   - Alloy Debug: [http://localhost:12345/debug/livedebugging]({{TRAFFIC_HOST1_12345}}/debug/livedebugging)

1. Multiplayer Access:

   - The game supports multiple players simultaneously

   - Players can join using:
     - `http://localhost:8080`{{copy}} from the same machine

     - `http://<host-ip>:8080`{{copy}} from other machines on the network

   - Each player can choose either the Southern or Northern faction

   - The game prevents multiple players from selecting the same faction

1. Single-Player Mode:

   - Toggle “Enable AI Opponent” in the game interface

   - The AI will automatically control the faction not chosen by the player

   - The AI provides a balanced challenge with adaptive strategies

   - For two-player games, keep the AI toggle disabled
