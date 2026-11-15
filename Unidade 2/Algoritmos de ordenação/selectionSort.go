// Online Go compiler (editor)
// Write and run Go online using this editor.

package main
import "fmt"

func selectionSortCrec(v[] int) {
  n := len(v)

  for varredura := 0; varredura < n - 1; varredura++ {
    index_menor := varredura

    for i := varredura + 1; i < n; i++ {
      if v[i] < v[index_menor] {
        index_menor = i
      }
    }
    temp := v[varredura]
    v[varredura] = v[index_menor]
    v[index_menor] = temp
  }
}

func selectionSortDec(v[] int) {
  n := len(v)

  for varredura := 0; varredura < n - 1; varredura++ {
    index_maior := varredura;

    for i := varredura; i < n; i++ {
      if v[i] > v[index_maior] {
        index_maior = i
      }
    }
    temp := v[varredura]
    v[varredura] = v[index_maior]
    v[index_maior] = temp
  }
}

func main() {
  v_test := []int{24, 9, 48, 10, 88, 6, 11, 15}

  fmt.Println("Vetor antes do Selection Sort")
  fmt.Println(v_test)

  // Ordenação por Selection Sort Crescente
  selectionSortCrec(v_test)

  fmt.Println("Vetor ordenado com o Selection Sort Crescente")
  fmt.Println(v_test)

  // Ordenação por Selection Sort Crescente
  selectionSortDec(v_test)

  fmt.Println("Vetor ordenado com o Selection Sort Decrescente")
  fmt.Println(v_test)
  
}