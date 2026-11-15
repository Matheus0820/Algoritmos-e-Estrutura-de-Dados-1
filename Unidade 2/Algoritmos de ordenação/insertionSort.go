package main
import "fmt"

func insertionSort(v[] int) {
  for i := 1; i < len(v); i++ {
    for j := i; j > 0 && v[j] < v[j - 1]; j-- {
      var_temp := v[j]
      v[j] = v[j - 1]
      v[j - 1] = var_temp
    }
  }
}

func main() {
  v_test := []int{24, 9, 48, 10, 88, 6, 11, 15}

  fmt.Println("Vetor antes do InsertionSort:")
  fmt.Println(v_test)

  // InsertionSort
  insertionSort(v_test)

  fmt.Println("Vetor depois do InsertionSort:")
  fmt.Println(v_test)

}
