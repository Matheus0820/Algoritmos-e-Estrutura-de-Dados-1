package main
import "fmt"

func bubbleSortCrec(v[] int) {
  n := len(v)

  for varredura := 0; varredura < n - 1; varredura++ {
    flag_trocou := false

    for i := 0; i < n - 1 - varredura; i++ {
      if v[i] > v[i + 1] {
        var_temp := v[i]
        v[i] = v[i + 1]
        v[i + 1] = var_temp
        flag_trocou = true
      }
    }
    if !flag_trocou {
      return // Se flag_trocou == false => Não houve troca => Vetor ordenado
    }
  }
}


func bubbleSortDec(v[] int) {
  n := len(v)

  for varredura := 0; varredura < n - 1; varredura++ {
    flag_trocou := false

    for i := 0; i < n - 1 - varredura; i++ {
      if v[i] < v[i + 1] {
        var_temp := v[i]
        v[i] = v[i + 1]
        v[i + 1] = var_temp

        flag_trocou = true
      }
    }
    if !flag_trocou {
      return // Se flag_trocou == false => Não houve troca => Vetor ordenado
    }
  }
}

func main() {
  v_test := []int{24, 9, 48, 10, 88, 6, 11, 15}

  fmt.Println("Vator antes da ordenação por BubbleSort:")
  fmt.Println(v_test)
  fmt.Println("")

  // Ordenação por BubbleSort Crescente
  bubbleSortCrec(v_test)

  fmt.Println("Vetor ordenado por BubbleSort Crescente:")
  fmt.Println(v_test)
  fmt.Println("")

  // Ordenação por BubbleSort Decrescente
  bubbleSortDec(v_test)
  
  fmt.Println("Vetor ordenado por BubbleSort Decrescente")
  fmt.Println(v_test)
  fmt.Println("")
  
}