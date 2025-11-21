# Cursor Rules - Resumo Executivo

Use este guia rápido antes de cada sessão no Cursor.

## Passos Essenciais
1. Revisar memória, `.cursorrules` e `.knowledgebase/` (anotar gaps).
2. Confirmar plano usando `.knowledgebase/templates/plan.md` e obter aprovação.
3. Dividir tarefas grandes e manter checklist atualizado em `.cursorrules`.
4. Pesquisar o código e registrar padrões antes de alterar qualquer arquivo.
5. Copiar templates oficiais (testes, boas práticas, layout, OTel) da pasta `.knowledgebase/templates/`.

## Obrigatórios (MUST)
- Salvar decisões e informações relevantes na memória + ADR.
- Explicar razões arquiteturais e gerar documento educativo em violações.
- Usar testes Given-When-Then (`test_gwt.go`) e cobertura ≥ 90%.
- Instrumentar com OpenTelemetry (snippet `otel_otlp.go`) e logs estruturados.
- Atualizar `.knowledgebase/architecture.md` e ADRs sempre que o design mudar.

## Recomendados (SHOULD)
- Adicionar links de referência quando ensinar conceitos difíceis.
- Executar testes em paralelo e manter mocks gerados.
- Usar o layout sugerido (`project_structure.txt`) ajustado ao projeto atual.
- Registrar spikes, hotfixes e exceções no checklist final.

## Onde encontrar
- Templates: `.knowledgebase/templates/`
- Exemplos iniciais: `.knowledgebase/examples/`
- Documentos educativos: `.knowledgebase/best-practices/`
- ADRs: `.knowledgebase/adr/`

Consulte `.cursorrules` para detalhes completos e mantenha este resumo à vista durante o fluxo.EOF