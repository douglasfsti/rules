# Regras Core (Seções 1–5)

## 1. Comportamento do Agente
### 1.1 Gerenciamento de Memória
- [MUST] Salvar dados importantes e revisar memória antes de agir.
- [MUST] Registrar decisões/ag acordos em memória e ADRs.
- [SHOULD] Mapear identidade, preferências e objetivos do usuário.

### 1.2 Comunicação e Contexto
- [MUST] Perguntar quando houver dúvidas e pesquisar o código antes de alterar.
- [MUST] Explicar decisões arquiteturais e gerar documentos educativos ao detectar violações.
- [MUST] Usar o template `best_practice.md` ao registrar boas práticas.

### 1.3 Planejamento e Execução
- [MUST] Apresentar plano (usar `plan.md`) e obter aprovação antes de alterar código.
- [SHOULD] Dividir tarefas grandes e manter checklist mínimo (Contexto, Objetivo, Passos, Riscos, Métrica).

### 1.4 Base de Conhecimento
- Checklist inicial: garantir `.knowledgebase/` pronta, revisar padrões e registrar como usá-los.
- Manter arquitetura, integrações e ADRs atualizados.
- Usar exemplos em `.knowledgebase/examples/` e comandos `cp` indicados.

### 1.5 Análise Contínua
- Sugerir refactors, registrar pontos fortes/fracos e atualizar `.cursorrules` com novas decisões.

## 2. Princípios Gerais
- Seguir Clean Architecture, Hexagonal, DDD tático, SOLID, KISS.
- Priorizar legibilidade, testabilidade, observabilidade.

## 3. Estrutura de Projeto
- Layout padrão disponível em `.knowledgebase/templates/project_structure.txt` (ajustar ao projeto).
- Manter documentação viva na base de conhecimento.

## 4. Estilo de Código (Go)
- Nomenclatura clara, ctx para contextos, evitar table-driven, wrapping de erros.
- Concorrência com context, canais e sync; evitar estado global.

## 5. Desenvolvimento de APIs
- Preferir chi + Go LTS.
- Validação rigorosa, JSON correto, middlewares, segurança (retries, circuit breaker, rate limit).

Consulte `.cursorrules` para visão resumida e `.knowledgebase/docs/resumo.md` para o guia rápido.
