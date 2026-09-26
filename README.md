# Concurrent Rate Limiter in Go

A concurrency-safe token-bucket limiter with an HTTP handler and injectable clock. It demonstrates a common API-protection decision without hiding the rule in framework middleware.

## Architecture and patterns

- **Adapter:** HTTP translates headers and status codes without entering the core.
- **Dependency Inversion:** clock and limiter are small interfaces defined by their consumer.
- A mutex protects invariants; tests also run with the Go race detector.
- **Pragmatic DDD:** the Traffic Protection bounded context uses bucket, tokens, refill, and decision as its ubiquitous language.

## Run

```bash
go test -race ./...
```

See [Architecture](docs/architecture.md) and [ADR-001](docs/adr/001-local-token-bucket.md).

## Trade-offs

Local state avoids a network dependency and offers low latency, but each replica maintains independent limits. Global coordination requires Redis, a database, or consistent routing, with availability and latency implications.

Repository and Visitor were rejected because they do not solve a current variation: state belongs to the local algorithm and there is only one decision type.

## License

MIT
