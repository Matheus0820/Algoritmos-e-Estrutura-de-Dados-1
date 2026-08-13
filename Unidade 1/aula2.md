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
