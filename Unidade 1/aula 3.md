# Análise de Complexidade Assintótica
## Motivação
- Algoritmos de podem possuir diferentes complexidade a por meio de diferentes implementações.
- Exemplo: Calculo de números primos
- De $0$ até $n$:
```golang
for i := 2; i < n; i++ {
  if n % i == 0 {
    return false
  }
} return true
  
```
- De $0$ até $\sqrt{n}$:
```golang
for i := 2; i < n**(1/2); i++ {
  if n % i == 0 {
    return false
  }
} return true
```

## Como comparar?
### Modelo Hipotético de maquina
- Fator que não depende da máquina $\rightarrow$ Tamanho da entrada ($n$)
- Operações lógicas aritméticas - Tempo $\rightarrow 1$ unidade de tempo ($ut$)
- Atribuições, retornos - Tempo $\rightarrow 1 ut$

## Exércícios simples
### Exemplo 1:
```golang
func soma (a int, b int) int { // 1 ut (atribuição)
  return a + b // 1 ut
} // Total: 2 ut
```
- Note que para valores grandes de $a$ e $b$ não muda nada no tempo de execulção do algoritmo, ele sempre vai execultar em $2 ut$. Portanto, dizemos que esse tipo de algoritmo tem complexidade $O(1)$

### Exemplo 2:
```golang
func somaLista(v []int) int {
  tot := 0 // menos importante para a análise

  for i := 0; i < len(v); i ++ { // IMPORTANTE PARA A ANÁLISE
    tot += v[i] // menos importante para análise
  }

  return tot // menos importante para a análise
}
```
- Note que nesse exemplo possui um `for`, no qual faz com que o algoritmo depende de mais tempo para ser execultado, pois o `for` ele execulta $n$ vezes, onde, nesse exemplo $n$ é o tamanho da lista `v`. Desse modo, dizemos que a complexidade desse algoritmo é $O(n)$.

### Exemplo 3: 
```golang
func somaMatriz(mat [][]int) int {
  tot := 0; // menos importante para a análise

  for i := 0; i < len(mat); i++ { // IMPORTANTE PARA A ANÁLISE
    for j := 0; j < len(mat[i]); j++ { // IMPORTANTE PARA A ANÁLISE
      tot += mat[i][j] // menos importante para análise
    }
  }
  return tot // menos importante para análise
}
```
- Note que em uma execulção do primeiro `for`, de $n$ execulções que ele vai fazer, o segundo `for` é execultado $n$ vezes, onde $n$ é o tamanho da matriz. Portanto, deizemos que um algoritmo desse tipo possui uma complexidade de $O(n^2)$.

## Famílias de complexidade
Quando é feito a análise de um algoritmo costumamos encontrar complexidades diferentes dos padrões citados, mas isso não faz diferença de fato, falamos que ele, a depender da sua forma, faz parte de uma família de complexidade expecífica.

### Família $O(n)$:
- $T = n$
- $T = 90 \cdot n$
- $T = 10493 \cdot n + 10$
- $T = 10^{99} \cdot n$

Não importa o que ta multiplicando o $n$, a complexidade de algoritmos assim sempre será $O(n)$

### Família $O(n^2)$:
- $T = n^2$
- $T = 900 \cdot n^2$
- $T = n^2 + 3906 \cdot n + 102$

Não importa os outros argumentos multiplicando ou somando, o que importa é o dominante, que nesse caso é o $n^2$ que cresce mais rápido. Desse modo, todos possuem complexidade $O(n^2)$
