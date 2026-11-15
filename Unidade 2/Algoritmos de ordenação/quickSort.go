package main
import "fmt"
import "math/rand"

func partition(v[] int, inicio, fim int) int {
  // Randomizar o pivô para evitar o pior caso (vetor ordenado)
  randPivo := rand.Intn(fim - inicio + 1) + inicio
  pivo := v[randPivo]
  v[randPivo] = v[fim]
  v[fim] = pivo

  pIndex := inicio

  for i := inicio; i < fim; i++ {
    if v[i] < pivo {
      var_temp := v[pIndex]
      v[pIndex] = v[i]
      v[i] = var_temp

      pIndex++
    }
  }

  var_temp := v[pIndex]
  v[pIndex] = v[fim]
  v[fim] = var_temp

  return pIndex
}

func quickSort(v[] int, inicio, fim int) {
  if inicio < fim {
    pIndex := partition(v, inicio, fim)
    quickSort(v, inicio, pIndex - 1)
    quickSort(v, pIndex + 1, fim)
  }
}

func main() {
  v_test := []int{24, 9, 48, 10, 88, 6, 11, 15}

  fmt.Println("Vetor antes do QuickSort:")
  fmt.Println(v_test)

  // QuickSort
  inicio := 0
  fim := len(v_test) - 1
  quickSort(v_test, inicio, fim)

  fmt.Println("Vetor depois do QuickSort:")
  fmt.Println(v_test)
}
