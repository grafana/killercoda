
# Step 2: Orchesrating our app with OpenTelemetry

In this step, we will be orchestrating our application with OpenTelemetry. We will do this in two parts:
1. Including the dependencies in our application
2. Initializing the OpenTelemetry logger and exporting logs to the OpenTelemetry Collector

**Note:** We will be making use of the vscode editor to make changes to our application. You can access the editor by clicking on the Editor tab.

## Including the dependencies in our application

The first step is to include the dependencies in our application. We will be using the OpenTelemetry API, SDK and the OTLP exporter. We will also be using the OpenTelemetry SDK logs module to set the global logger provider.

```python
# Import the function to set the global logger provider from the OpenTelemetry logs module.
from opentelemetry._logs import set_logger_provider

# Import the OTLPLogExporter class from the OpenTelemetry gRPC log exporter module.
from opentelemetry.exporter.otlp.proto.grpc._log_exporter import (
    OTLPLogExporter,
)

# Import the LoggerProvider and LoggingHandler classes from the OpenTelemetry SDK logs module.
from opentelemetry.sdk._logs import LoggerProvider, LoggingHandler

# Import the BatchLogRecordProcessor class from the OpenTelemetry SDK logs export module.
from opentelemetry.sdk._logs.export import BatchLogRecordProcessor

# Import the Resource class from the OpenTelemetry SDK resources module.
from opentelemetry.sdk.resources import Resource
```{{copy}}

Make sure to copy the above code snippet to the `app.py` file in the vscode editor. Under the section called `#### otel dependencies ####`, you can paste the code snippet.

## Initializing the OpenTelemetry logger and exporting logs to the OpenTelemetry Collector

The next step is to initialize the OpenTelemetry logger and export logs to the OpenTelemetry Collector. We will be using the OTLP exporter to export logs to the OpenTelemetry Collector. We will also be setting the global logger provider to the OpenTelemetry SDK logger provider.

```python
# Create an instance of LoggerProvider with a Resource object that includes
# service name and instance ID, identifying the source of the logs.
logger_provider = LoggerProvider(
    resource=Resource.create(
        {
            "service.name": "greenhouse-app",
            "service.instance.id": "instance-1",
        }
    ),
)

# Set the created LoggerProvider as the global logger provider.
set_logger_provider(logger_provider)

# Create an instance of OTLPLogExporter with insecure connection.
exporter = OTLPLogExporter(insecure=True)

# Add a BatchLogRecordProcessor to the logger provider with the exporter.
logger_provider.add_log_record_processor(BatchLogRecordProcessor(exporter))

# Create a LoggingHandler with the specified logger provider and log level set to NOTSET.
handler = LoggingHandler(level=logging.NOTSET, logger_provider=logger_provider)

# Attach OTLP handler to the root logger.
logging.getLogger().addHandler(handler)
```{{copy}}

Make sure to copy the above code snippet to the `app.py` file in the vscode editor. Under the section called `## Otel logger initialization ##`, you can paste the code snippet.


## Stuck?

If you are stuck, we have provided the completed `app.py` file in the `completed` directory. You can copy the file to the `~/carnivorous-greenhouse` directory by running the following command:

```bash
cp ./completed/app.py ./app.py
```{{execute}}
