# Prompt: Benchmark Tests

**Objetivo:** Criar testes de benchmark em Go para avaliar a performance de funções críticas.

**Instruções:**
Escreva benchmarks para o código fornecido, focando em medir tempo de execução e alocação de memória.

**Requisitos:**
1.  Use o padrão `func BenchmarkNome(b *testing.B)`.
2.  Inclua `b.ReportAllocs()` para monitorar alocações.
3.  Certifique-se de resetar o timer (`b.ResetTimer()`) se houver setup pesado.
4.  Teste diferentes tamanhos de input se relevante.

**Exemplo:**
```go
func BenchmarkFuncao(b *testing.B) {
    // Setup (se necessário)
    input := "dados de teste"
    
    b.ResetTimer()
    b.ReportAllocs()
    
    for i := 0; i < b.N; i++ {
        Funcao(input)
    }
}
```

**Código para Benchmark:**
```go
[INSIRA O CÓDIGO AQUI]
```
