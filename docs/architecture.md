# Arquitetura

O domínio contém apenas a decisão. A aplicação implementa token bucket e depende de `Clock`; adaptadores fornecem relógio e HTTP. O binário faz a composição explícita.
