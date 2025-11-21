# Prompt: Commit Message

**Objetivo:** Gerar uma mensagem de commit seguindo o padrão Conventional Commits.

**Instruções:**
Analise as alterações feitas no código (diff) e gere uma mensagem de commit clara e concisa.

**Padrão (Conventional Commits):**
- `feat`: Nova funcionalidade.
- `fix`: Correção de bug.
- `docs`: Alterações na documentação.
- `style`: Formatação, ponto e vírgula, etc (sem alteração de código).
- `refactor`: Refatoração de código (sem nova funcionalidade ou correção de bug).
- `perf`: Melhoria de performance.
- `test`: Adição ou correção de testes.
- `chore`: Alterações no build, ferramentas auxiliares, etc.

**Formato da Saída:**
```text
<tipo>(<escopo opcional>): <descrição curta>

[Corpo opcional: Explicação mais detalhada do "porquê" e "o quê", não "como"]

[Rodapé opcional: Breaking Changes ou referências a issues (ex: Closes #123)]
```

**Diff:**
```diff
[INSIRA O DIFF AQUI]
```
