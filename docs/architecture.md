# Arquitetura

## Contexto e linguagem

Bounded context **Proteção de Tráfego**. Bucket, capacidade, tokens, reposição, cliente e retry compõem a linguagem ubíqua. `Decision` é um value object imutável; cada bucket interno é a entidade de estado identificada pela chave do cliente.

## Fronteiras

- **Domínio:** decisão observável pelo consumidor.
- **Aplicação:** algoritmo token bucket e portas `Clock`/`Limiter`.
- **Infraestrutura:** relógio do sistema, HTTP e composição do binário.

## Padrões e alternativas

- **Strategy:** `Limiter` permite trocar token bucket por janela deslizante. Condicionais por algoritmo foram descartadas.
- **Adapter:** handler converte protocolo HTTP em chave/decisão.
- Repository não foi introduzido: o estado pertence ao algoritmo local; uma versão distribuída criaria uma porta de armazenamento própria.
- Visitor foi descartado porque há um único tipo de decisão e nenhuma árvore de elementos.

Mutex único simplifica invariantes e race safety, mas limita paralelismo por processo. Sharding seria uma evolução medida, não uma abstração preventiva.
