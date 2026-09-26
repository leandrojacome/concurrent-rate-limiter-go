# ADR-001: Local token bucket

## Status

Accepted for the portfolio scope.

## Decision

Use an in-memory token bucket protected by a mutex and explicitly document the distributed limitation.

## Consequences

This provides low latency and deterministic tests. There is no coordination across replicas or cleanup of inactive keys; both are production-evolution requirements.
