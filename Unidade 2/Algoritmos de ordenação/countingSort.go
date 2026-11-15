package main
import "fmt"

func countingSort(v []int) []int {
  if len(v) == 0 {
    return v
  }

  max := v[0]

  for _, val := range v {
    if val > max {
      max = val
    }
  }

  count := make([]int, max + 1)
  out := make([]int, len(v))

  for _, val := range v {
    count[val]++
  }

  for i := 1; i <= max; i++ {
    count[i] += count[i - 1]
  }

  for i := len(v) - 1; i >= 0; i-- {
    out[count[v[i]] - 1] = v[i]
    count[v[i]]--
  }

  return out
}

func main() {
  v_test := []int{24, 9, 48, 10, 88, 6, 11, 15}

  fmt.Println("Vetor antes do CountingSort:")
  fmt.Println(v_test)

  // CountingSort
  v_test_ordenado := countingSort(v_test)

  fmt.Println("Vetor depois do CountingSort:")
  fmt.Println(v_test_ordenado)
}
