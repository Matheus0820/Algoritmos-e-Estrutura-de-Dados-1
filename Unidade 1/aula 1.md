# Aula 1 de AED1
## Análise Assintótica de Algoritmos
A complexidade de um algoritmo é denominada de $O$, no qual indica qual é a quantidade de esforço feito para fazer certa tarefa com $n$ elementos.

### Exemplos de complexidade
#### $O(1)$ $\rightarrow$ Big. Oh (pior caso):
- Retorno de uma função simples:
```golang
func size() {
  return size;
}
```

#### $O(log_2 (n))$
- Busca binária
- Todos os algoritmos de busca que partam a estrutura no meio para fazer busca

#### $O(n)$
- Busca sequêncial

#### $O(n \cdot log_2 (n))$
- Merge Sort

#### $O(n^2)$
- Laços de repetição dentro de laços de repetição
```golang
for i := 0; i <= n; i++ {
  for j := 0; j <= n; j++ {
        fmt.Println(i);
        fmt.Println(j);
    }
}
```
## Tarefa para trazer na próxima aula
Fazer um algoritmo na linguagem que preferir para adicionar o numero de 1 a 1 milhão pelo inicio e pelo final e medir o tempo dos dois. Analisar se o tempo é diferente e porquê é diferente
### Adicionando atrás:
```golang
package main

import "fmt"
import "time"

func main() {
  var vetor []int

  inicio := time.Now()

  for i := 0; i < 1000000; i++ {
    vetor = append(vetor, i)
  }

  fim := time.Since(inicio)

  tempo_formated := fmt.Sprintf("%f segundos", fim.Seconds())
  fmt.Println(tempo_formated)

  
}
```


### Adicionando na frente:
```golang
package main

import "fmt"
import "time"

func main() {
  var vetor []int

  inicio := time.Now()

  for i := 0; i < 1000000; i++ {
    vetor = append([]int{i}, vetor...)
  }

  fim := time.Since(inicio)

  tempo_formated := fmt.Sprintf("%f segundos", fim.Seconds())
  fmt.Println(tempo_formated)

  
}
```
