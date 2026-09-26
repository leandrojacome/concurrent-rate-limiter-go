# Concurrent Rate Limiter — Go

Limitador token-bucket seguro para concorrência, com handler HTTP e relógio injetável. Demonstra uma decisão típica de proteção de APIs sem esconder a regra em middleware de framework.

## Arquitetura e padrões

- **Strategy (GoF):** a porta `Limiter` permite trocar token bucket por janela deslizante.
- **Adapter:** HTTP traduz cabeçalhos/status sem entrar no núcleo.
- **Dependency Inversion:** relógio e limitador são interfaces pequenas, definidas pelo consumidor.
- Mutex protege invariantes; os testes executam também com detector de corrida.
- **DDD proporcional:** bounded context de Proteção de Tráfego, linguagem de bucket/tokens/reposição e `Decision` como value object.

```bash
go test -race ./...
go vet ./...
govulncheck ./...
go run ./cmd/server
```

Veja [Arquitetura](docs/architecture.md) e [ADR-001](docs/adr/001-local-token-bucket.md).

## Trade-offs

O estado local evita dependência de rede e oferece baixa latência, mas cada réplica mantém limites independentes. Coordenação global requer Redis, banco ou roteamento consistente, com impacto de disponibilidade e latência.

Repository e Visitor foram descartados por não resolverem uma variação presente: o estado é parte do algoritmo e existe apenas um tipo de decisão.

## Licença

MIT.
