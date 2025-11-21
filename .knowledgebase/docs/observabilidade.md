# Guia Rápido de Observabilidade

## Objetivo
Explicar rapidamente como aplicar OpenTelemetry, logging estruturado e monitoramento no projeto.

## Passos Essenciais
1. **Tracer**
   - Copiar `.knowledgebase/templates/otel_otlp.go` para o módulo principal.
   - Ajustar endpoint `OTEL_EXPORTER_OTLP_ENDPOINT`, recursos e atributos.
   - Registrar no plano + `.knowledgebase/architecture.md`.
2. **Spans & Contexto**
   - Criar spans em handlers, casos de uso e integrações (`otel.Tracer("svc")`).
   - Propagar `context.Context` em todas as funções públicas.
3. **Logs Estruturados**
   - Utilizar logger com JSON (`log/slog`, `zerolog` ou similar).
   - Injetar `trace_id` e `span_id` em cada log.
4. **Métricas**
   - Criar contadores/ histogramas críticos (latência, throughput, erros).
   - Exportar para Prometheus/Grafana ou collector equivalente.
5. **Alertas & SLIs**
   - Definir metas (p. ex.: p95 < 300ms, erro < 1%).
   - Configurar alertas para 5xx, timeouts de DB, filas acumuladas.

## Checklist Rápido
- [ ] Snippet OTLP aplicado e documentado.
- [ ] Handlers instrumentados com spans.
- [ ] Logs estruturados com IDs correlacionados.
- [ ] Métricas de latência/erro expostas.
- [ ] SLIs/alertas registrados no ADR ou doc de arquitetura.

Consulte `.cursorrules` seção 7 para detalhes.
