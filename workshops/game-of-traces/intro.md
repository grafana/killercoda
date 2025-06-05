# War of Kingdoms: A Distributed Tracing Tutorial Game

This educational game demonstrates distributed tracing concepts through an interactive strategy game built with OpenTelemetry and Grafana Alloy. Players learn about trace sampling, service graphs, and observability while competing for territory control.

## Educational Goals

This game teaches several key concepts in distributed tracing:

1. **Distributed System Architecture**

   - Multiple microservices (locations) communicating via HTTP

   - Shared state management

   - Event-driven updates

   - Real-time data propagation

1. **OpenTelemetry Concepts**

   - Trace context propagation

   - Span creation and attributes

   - Service naming and resource attributes

   - Manual instrumentation techniques

1. **Observability Patterns**

   - Trace sampling strategies

   - Error tracking and monitoring

   - Performance measurement

   - Service dependencies visualization

## Game Overview

The game simulates a war between two kingdoms, each starting from their capital city. Players must:

- Collect resources from their territories

- Build armies to expand their influence

- Capture neutral villages

- Send resources back to their capital

- Launch strategic attacks on enemy territories

The game supports both single-player (with AI opponent) and two-player modes.

Each action in the game generates traces that can be analyzed in Grafana Tempo, demonstrating how distributed tracing works in a real application.

## Technical Components

The application consists of:

- **Location Servers**: Python Flask microservices representing different map locations

- **War Map UI**: Web interface for game interaction

- **AI Opponent**: Intelligent computer player for single-player mode

- **Telemetry Pipeline**:
  - OpenTelemetry SDK for instrumentation

  - Grafana Alloy for trace processing

  - Tempo for trace storage

  - Prometheus for metrics

  - Loki for logs

  - Grafana for visualization
