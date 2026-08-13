# Listas

## Listas com ArrayList

<p align="center">
  <img src="https://miro.medium.com/1*DxlJoSM_VaoeHQEaPnFT5g.png" alt="Imagem">
</p>

O array list aloca memória na implementação do array e quando é necessário adicionar no inicio se torna mais custoso, pois é necessário realocar todos os elementos da lista para fazer essa operação, o que torna a operação do tipo $O(n)$. Já adicionar um elemento no final no array é mais rápido, pois não é necessário realocar todo os elementos do array, tendo uma complexidade $O(1)$.

## Listas com LinkedList

<p align="center">
  <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT_ZSTiF8PJ6JttYG_mWGzgI0pyGzKTStlrEEVY03L6jw&s=10" alt="Imagem">
</p>

No LinkedList ele aloca a memória aos poucos quando os elementos vão sendo adicionados, onde cada elemento é armazenado em um struct que guarda o ponteiro para o valor e também para o próximo elemento do LinkedList e esse elementos não fica encadeados linearmente na memória, ao contrário do array list, mas em espaços diferentes. Desse modo, adicionar no inicio desse tipo de abstração de dados é mais fácil, tendo uma complexidade $O(1)$, enquanto adicionar no final tem um complexidade $O(n)$.

Em resumo as diferenças e caracteristicas de cada uma é mostrada na imagem abaixo: 
<p align="center">
  <img src="https://programai.com.br/wp-content/uploads/2024/12/array-vs-lista.png" alt="Array vs Lista">
</p>

# Complexidade
## Notação
- Big-Oh $\rightarrow O()$
- Little-Oh $\rightarrow o()$

## Tipos de complexidade
Existe vários tipos de complexidade de algoritmos, desde os que são extremamente bons para valores de $n$ extremamente grande e outras são ruins. A **Análise Assintótica** de um algoritmo tem como objetivo justamente esse, analisar o comportamento do algoritmo a a partir do momento que o valor da entrada aumenta indefinitivamente.

<p align="center">
  <img src="https://hermes.dio.me/articles/cover/84e2daa2-c071-4e00-a9ce-4d110056be18.jpg" alt="Imagem">
</p>

## Complexidade de algoritmos
### Complexidade do tipo $O(1)$
O tempo de um algoritmo depende da entrada, no entanto alguns algoritmos, os denominados de $O(1)$ não tem relação com a entrada, e o tempo de execulção desse algoritmo para qualquer valor de entrada será igual, sempre igual, em muita vezes instantâneo.

```golang
func soma(a int, b int) int {
 return a + b
}
```

### Complexidade do tipo $O(n)$
Já existem algoritmos que não fojem dessa exceção, o tempo de execulção realmente depende do número da entrada $n$. O exemplo disso é quando existe um `for` no meio ou alguma estrutura de repetição, onde um processo vai ser feita no máximo $n$ vezes. Um exemplo desse tipo de algoritmo é:
```golang
func somaLista(a []int, n int) {
  soma := 0

  for i := 0; i < n; i++ {
    soma += a[i]
  }

  return soma
}
```
No entanto, o tipo de complexidade pode depender de como o incremento e o limite do laço é definido. Por exemplo, se o incremento do `i` dentro do `for` for feito de forma diferente, ele pode terminar mais rápido ou. até mesmo, demorar mais. Outra coisa que afeta é o valor do limite do laço ser diferente de $n$, isso impacta também na valocidade do algoritmo, ai que surge os tipos de algoritmos de complexidade $O(log_2(n))$ e outros não convencionais.

Outro ponto interessante a salientar, é que a complexidade dentro de um algoritmo, como falado antes, está diretamente ligada a quatidade de laços de repetição e quantidade de repetições que o algoritmo execulta.

### Complexidade do tipo $O(n^2)$
Um algoritmos que é desse tipo é causado por laços aninhados. Um exemplo é mostrado abaixo
```golang

func somaMatriz( matriz [][]int, n int) {
  soma := 0

  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      soma += matriz[i][j]
    }
  }
  return soma
}

```

Note que para cada $i$ o $j$ é execultado $n$ vezes em $n$ vezes que o $i$ é execultado. Veja com mais detalhe a seguir:





