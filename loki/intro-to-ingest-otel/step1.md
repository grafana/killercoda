
# Step 1: Preparing our Python application

For this demo, we will be using a simple Python application called Carnivorous Greenhouse. This application generates logs that we will be ingesting into Loki. This application will be installed locally on the sandbox environment.

## Package Installation

First we will create our Python virtual environment. We will use the following command to create our virtual environment:

```bash
python3 -m venv .venv
source ./.venv/bin/activate
```{{execute}}

The carniverous greenhouse application relies on a few Python packages. We will install these packages using the following command:

```bash
pip install -r requirements.txt
```{{execute}}

Next we will install the requried OpenTelemetry packages to instrument our application: 
* Opentelemetry-distro: This contains OpenTelemetry API, SDK
* Opentelemetry-exporter-otlp: This package allows us to export these logs within the OTLP format otherwise known as OpenTelemtry Protocol more on where we are sending our log entries later.
We will use the following command to install the required packages:

```bash
pip install opentelemetry-distro opentelemetry-exporter-otlp
```{{execute}}




