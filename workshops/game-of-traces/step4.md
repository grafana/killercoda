# Architecture Deep Dive

## Trace Flow Example: Army Movement

1. Player initiates move (UI span)

1. Source location processes request (source span)

1. Movement calculation (path span)

1. Target location receives army (target span)

1. Battle resolution if needed (battle span)

1. State updates propagate (update spans)

Each step generates spans with relevant attributes, demonstrating trace context propagation in a distributed system.

# Educational Use

This project is designed for educational purposes to teach:

- Distributed systems concepts

- Observability practices

- Microservice architecture

- Real-time data flow

- System instrumentation
