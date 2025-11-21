# Prompt: Unit Tests (GWT)

**Objetivo:** Criar testes unitários em Go seguindo o padrão Given-When-Then (GWT), conforme as regras do projeto.

**Instruções:**
Atue como um especialista em testes em Go. Escreva testes unitários para a função/método fornecido.

**Requisitos:**
1.  Use a biblioteca padrão `testing`.
2.  Adote a estrutura Given-When-Then (GWT) dentro dos testes.
3.  Evite Table-Driven Tests complexos; prefira subtestes explícitos (`t.Run`) ou funções de teste separadas para clareza.
4.  Garanta alta cobertura de código.
5.  Use `assert` ou `require` (se a lib `testify` estiver disponível no projeto) ou verificações manuais claras.

**Exemplo de Estrutura:**
```go
func TestNomeDaFuncao(t *testing.T) {
    t.Run("Deve retornar sucesso quando input for válido", func(t *testing.T) {
        // Given (Setup)
        input := "dados válidos"
        expected := "resultado esperado"

        // When (Execução)
        got, err := NomeDaFuncao(input)

        // Then (Verificação)
        if err != nil {
            t.Fatalf("erro inesperado: %v", err)
        }
        if got != expected {
            t.Errorf("got = %v, want %v", got, expected)
        }
    })

    t.Run("Deve retornar erro quando input for inválido", func(t *testing.T) {
        // Given
        input := "dados inválidos"

        // When
        _, err := NomeDaFuncao(input)

        // Then
        if err == nil {
            t.Error("esperava erro, mas não ocorreu")
        }
    })
}
```

**Código para Testar:**
```go
[INSIRA O CÓDIGO AQUI]
```
