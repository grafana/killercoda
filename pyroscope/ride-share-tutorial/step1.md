# Clone the repository

Clone the repository to your local machine:

```bash
git clone https://github.com/grafana/pyroscope.git 
```{{exec}}

Navigate to the tutorial directory:

```bash
cd examples/language-sdk-instrumentation/python/rideshare/flask
```{{exec}}

# Start the application

Start the application using Docker Compose:

```bash
docker compose up -d
```{{exec}}

This may take a few minutes to download the required images and build the demo application. Once ready, you will see the following output:

```console
 ✔ Network flask_default  Created
 ✔ Container flask-ap-south-1  Started
 ✔ Container flask-grafana-1  Started
 ✔ Container flask-pyroscope-1  Started     
 ✔ Container flask-load-generator-1 Started 
 ✔ Container flask-eu-north-1 Started       
 ✔ Container flask-us-east-1 Started  
```{{copy}}

(Optional) To verify the containers are running run:

```bash
docker ps -a
```{{exec}}
