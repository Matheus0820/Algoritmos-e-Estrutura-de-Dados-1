package main
import "fmt"

func Merge (v, esq, dir []int) {
  i_esquerda, i_direita, i_vetor := 0, 0, 0

  for i_esquerda < len(esq) && i_direita < len(dir) {
    if esq[i_esquerda] <= dir[i_direita] {
      v[i_vetor] = esq[i_esquerda]
      i_esquerda++
    } else {
      v[i_vetor] = dir[i_direita]
      i_direita++
    }
    i_vetor++
  }
  
  for i_esquerda < len(esq) {
    v[i_vetor] = esq[i_esquerda]
    i_esquerda++
    i_vetor++
  }

  for i_direita < len(dir) {
    v[i_vetor] = dir[i_direita]
    i_direita++
    i_vetor++
  }
}

func MergeSort(v[] int) {
  if len(v) > 1{
    tam_esquerda := len(v) / 2
    esq := make([]int, tam_esquerda)
    dir := make([]int, len(v) - tam_esquerda)
    copy(esq, v[:tam_esquerda])
    copy(dir, v[tam_esquerda:])

    MergeSort(esq)
    MergeSort(dir)
    Merge(v, esq, dir)
  }
}

func main() {
  v_test := []int{24, 9, 48, 10, 88, 6, 11, 15}

  fmt.Println("Vetor antes do MergeSort:")
  fmt.Println(v_test)

  // MergeSort
  MergeSort(v_test)

  fmt.Println("Vetor depois do MergeSort:")
  fmt.Println(v_test)
}
