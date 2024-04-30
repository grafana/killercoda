
# Step 1: Installing Loki and collecting sample logs

**To install Loki locally, follow these steps:**

1. Create a directory called `evaluate-loki` for the demo environment. Make `evaluate-loki` your current working directory:

    ```bash
    mkdir evaluate-loki
    cd evaluate-loki
    ```{{execute}}

2. Download `loki-config.yaml`, `alloy-local-config.yaml`, and `docker-compose.yaml`:

    ```bash
    wget https://raw.githubusercontent.com/grafana/loki/main/examples/getting-started/loki-config.yaml -O loki-config.yaml
    wget https://raw.githubusercontent.com/grafana/loki/main/examples/getting-started/alloy-local-config.yaml -O alloy-local-config.yaml
    wget https://raw.githubusercontent.com/grafana/loki/main/examples/getting-started/docker-compose.yaml -O docker-compose.yaml
    ```{{execute}}

3. Deploy the sample Docker image.

    With `evaluate-loki` as the current working directory, start the demo environment using `docker compose`:

    ```bash
    docker-compose up -d
    ```{{execute}}

    You should see something similar to the following:

    ```bash
    ✔ Network evaluate-loki_loki          Created      0.1s 
    ✔ Container evaluate-loki-minio-1     Started      0.6s 
    ✔ Container evaluate-loki-flog-1      Started      0.6s 
    ✔ Container evaluate-loki-backend-1   Started      0.8s 
    ✔ Container evaluate-loki-write-1     Started      0.8s 
    ✔ Container evaluate-loki-read-1      Started      0.8s 
    ✔ Container evaluate-loki-gateway-1   Started      1.1s 
    ✔ Container evaluate-loki-grafana-1   Started      1.4s 
    ✔ Container evaluate-loki-alloy-1     Started      1.4s
    ```

4. (Optional) To check the status of the containers, run the following command:

    ```bash
    docker ps -a 
    ```{{execute}}

All containers should be showing the following:

```bash
  CREATED         STATUS                       
   5 seconds ago   Up 3 seconds    
```

5. (Optional) Verify that the Loki cluster is up and running.
    - The read component returns `ready` when you point a web browser at [{{TRAFFIC_HOST1_3101}}/ready](http://localhost:3101/ready). The message `Query Frontend not ready: not ready: number of schedulers this worker is connected to is 0` will show prior to the read component being ready.
    - The write component returns `ready` when you point a web browser at [{{TRAFFIC_HOST1_3102}}/ready](http://localhost:3102/ready). The message `Ingester not ready: waiting for 15s after being ready` will show prior to the write component being ready.

6. (Optional) Verify that Grafana Alloy is running.
    - Grafana Alloy's UI can be accessed at [http://localhost:12345]({{TRAFFIC_HOST1_12345}}).  

