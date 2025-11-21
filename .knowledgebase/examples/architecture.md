# Visão Geral
Este documento descreve o estado atual da arquitetura de referência utilizada pelos serviços Go.

## Contexto Atual
- Serviços escritos em Go 1.22+
- Clean Architecture + Ports & Adapters
- Observabilidade com OpenTelemetry (OTLP gRPC)

## Componentes
1. `cmd/api`: entrypoint HTTP
2. `internal/core/domain`: entidades e regras de negócio
3. `internal/core/ports`: interfaces de entrada/saída
4. `internal/adapters`: HTTP handlers, repositórios, integrações externas

## Fluxos Principais
- Requisições HTTP entram pelo handler do chi
- Handler chama caso de uso no pacote `services`
- Casos de uso conversam com repositórios via interfaces
- Respostas retornam validadas e observáveis

## Alertas/Observações
- Registrar novas integrações externas nesta seção
- Atualizar ADR correspondente para decisões estruturais
- Garantir que os diagramas e fluxos fiquem sincronizados com este arquivo
