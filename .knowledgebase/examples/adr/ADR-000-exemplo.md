# ADR-000 - Preferência por Router chi

## Contexto
Precisamos definir um roteador HTTP padrão para novos serviços Go, priorizando simplicidade, performance e comunidade ativa.

## Decisão
Adotar o framework `github.com/go-chi/chi/v5` como router padrão para novos serviços HTTP.

## Consequências
- **Positivas**: Middleware robusto, rotas composáveis, comunidade ativa, alinhamento com requisitos de Clean Architecture.
- **Negativas**: Dependência adicional no projeto, necessidade de treinar a equipe nos recursos do chi.

## Referências
- https://github.com/go-chi/chi
- https://pkg.go.dev/github.com/go-chi/chi/v5
