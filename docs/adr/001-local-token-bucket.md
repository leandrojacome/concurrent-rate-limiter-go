# ADR-001: token bucket local

## Status

Aceito para o escopo do portfólio.

## Decisão

Usar token bucket em memória protegido por mutex e explicitar a limitação distribuída.

## Consequências

Baixa latência e teste determinístico. Não há coordenação entre réplicas nem limpeza de chaves inativas; ambos são requisitos de uma evolução produtiva.
