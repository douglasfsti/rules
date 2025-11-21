# Regras de Execução (Seções 6–9)

## 6. Testes
- [MUST] Sem table-driven; usar Given-When-Then (`test_gwt.go`).
- [MUST] Cobertura ≥ 90%, mocks gerados, unit tests obrigatórios.
- [SHOULD] t.Parallel, smoke tests só com exceção aprovada.
- Prioridades: hotfix (checklist reduzido), refatoração (checklist completo), spikes (registrar descobertas).

## 7. Observabilidade
- [MUST] OpenTelemetry + spans em todos os limites.
- [MUST] Context em logs/métricas; snippet oficial em `.knowledgebase/templates/otel_otlp.go`.
- Registrar reaproveitamento de tracer em architecture.md/ADR.
- Logs estruturados, métricas críticas, SLIs/alertas.

## 8. Documentação
- Docstrings em inglês para métodos principais.
- Evitar comentários excessivos; usar READMEs e ARCHITECTURE.md.

## 9. Ferramentas e Dependências
- gofmt/goimports/golangci-lint obrigatórios; automação com CI.
- Go modules com versões travadas e libs estáveis.

Para passos concretos, abra `.knowledgebase/docs/observabilidade.md` e `.knowledgebase/docs/checklist-visual.md`.
