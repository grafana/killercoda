# Learning Through Play

## 1. Trace Context Propagation

Watch how actions propagate through the system:

- Resource collection triggers spans across services

- Army movements create trace chains

- Battle events generate nested spans

## 2. Sampling Strategies

The game demonstrates different sampling approaches:

- Error-based sampling (captures failed battles)

- Latency-based sampling (slow resource transfers)

- Attribute-based sampling (specific game events)

## 3. Service Graph Analysis

Learn how services interact:

- Village-to-capital resource flows

- Army movement paths

- Battle resolution chains

# Observability Features

## 1. Resource Movement Tracing

```console
{span.resource.movement = true}
```{{copy}}

Track resource transfers between locations with detailed timing and amounts.

## 2. Battle Analysis

```console
{span.battle.occurred = true}
```{{copy}}

Analyze combat events, outcomes, and participating forces.

## 3. Player Actions

```console
{span.player.action = true}
```{{copy}}

Monitor player interactions and their impact on the game state.
