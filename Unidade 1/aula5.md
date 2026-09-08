# LinkedList e Double LinkedList

## Revisão - ArrayList: 
No struct do ArrayList possuia dois parametros, o vetor `v` e uma variável que falava quantos elementos foram adicionados dentro da lista, chamado de `inserted`

## LinkedList
No struct inicial ele tem declarado o endereço do primeiro elemeneto da lista ligada, chamado de `head` e também uma variável que fala quantos valores já foram inseridos, chamada de `inseted`. O `head` vai guardar o endereço do primeiro nó, no qual o nó possui, em seu struct os parametros para guardar o valor da lista naquele index, chamada `val` e outro valor, que será um ponteiro, que aponta para o próximo nó da lista, denominado de `next`. O último nó da lista, tem como valor definido em `next` como nulo (`null` ou `nil`).

<p align="center">
  <img src="https://media.geeksforgeeks.org/wp-content/uploads/20210409184741/HowtoImplementGenericLinkedListinJava.jpg" alt="Linked List em Java">
</p>

Nele, a implementação é um pouco diferente do ArrayList. Por ter essa estrutura diferenciada, com nós, a adicão no final, que antes era simple no ArrayList, agora requer mais processo. Isso também impacta na adição de elementos no meio da lista, ou seja, em um `index` específico. Para isso é necessário a presença de uma estrutura de repetição. 

Abaixo veja como é a implementação do mótodo `add` do LinkedList:
```golang
aux := list.head // aux agora é um ponteiro que aponta para o primeiro nó da lista ligada
for aux.next != nil {
  aux = aux.next // aux agora vai apontar para o próximo nó, tendo acesso ao endereço do próximo nó, se existir, e ao valor do nó desse index
}

// Aqui ele cria um outro nó, adiciona o valor e define o valor aux.next igual ao endereço desse novo nó
```
