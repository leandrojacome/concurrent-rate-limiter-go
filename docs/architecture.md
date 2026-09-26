# Architecture

## Domain model

The **Traffic Protection** bounded context uses bucket, capacity, tokens, refill, client, and retry as its ubiquitous language. `Decision` is an immutable value object; each internal bucket is the stateful entity identified by a client key.

## Layers

- **Domain:** the decision observed by the consumer.
- **Application:** token-bucket algorithm and `Clock`/`Limiter` ports.
- **Infrastructure:** system clock, HTTP adapter, and binary composition.

## Patterns and alternatives

- **Adapter:** the handler converts HTTP protocol details into a key and decision.
- Repository is intentionally absent: state belongs to the local algorithm. A distributed implementation would introduce its own storage port.
- Visitor was rejected because there is a single decision type and no element hierarchy.

A single mutex simplifies invariants and race safety but limits in-process parallelism. Sharding is a measured future optimization, not a preventive abstraction.
