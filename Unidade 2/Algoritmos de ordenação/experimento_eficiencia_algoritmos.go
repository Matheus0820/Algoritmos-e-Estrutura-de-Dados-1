package main

import "fmt"
import "math/rand"
import "time"

func insertionSort(v []int) {
	for i := 1; i < len(v); i++ {
		for j := i; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}

func Merge(v []int, esq []int, dir []int) {
	ie, id, iv := 0, 0, 0
	for ie < len(esq) && id < len(dir) {
		if esq[ie] <= dir[id] {
			v[iv] = esq[ie]
			ie++
		} else {
			v[iv] = dir[id]
			id++
		}
		iv++
	}
	for ie < len(esq) {
		v[iv] = esq[ie]
		ie++
		iv++
	}
	for id < len(dir) {
		v[iv] = dir[id]
		id++
		iv++
	}
}

func MergeSort(v []int) {
	if len(v) > 1 {
		mid := len(v) / 2
		esq := make([]int, mid)
		dir := make([]int, len(v)-mid)
		copy(esq, v[:mid])
		copy(dir, v[mid:])
		MergeSort(esq)
		MergeSort(dir)
		Merge(v, esq, dir)
	}
}

func partition(v []int, ini int, fim int) int {
	pivot := v[fim]
	pIndex := ini
	for i := ini; i < fim; i++ {
		if v[i] <= pivot {
			v[pIndex], v[i] = v[i], v[pIndex]
			pIndex++
		}
	}
	v[fim], v[pIndex] = v[pIndex], v[fim]
	return pIndex
}

func quickSort(v []int, ini int, fim int, randomize bool) {
	if ini < fim {
		if randomize {
			p := rand.Intn(fim-ini+1) + ini
			v[p], v[fim] = v[fim], v[p]
		}
		idxPivot := partition(v, ini, fim)
		quickSort(v, ini, idxPivot-1, randomize)
		quickSort(v, idxPivot+1, fim, randomize)
	}
}

func countingSort(v []int, k int) {
	count := make([]int, k+1)
	saida := make([]int, len(v))
	for i := 0; i < len(v); i++ {
		count[v[i]]++
	}
	for i := 1; i <= k; i++ {
		count[i] += count[i-1]
	}
	for i := len(v) - 1; i >= 0; i-- {
		saida[count[v[i]]-1] = v[i]
		count[v[i]]--
	}
	copy(v, saida)
}

func createArray(size int, order string, k int) []int {
	v := make([]int, size)
	for i := 0; i < size; i++ {
		switch order {
		case "crescente":
			v[i] = i
		case "decrescente":
			v[i] = size - 1 - i
		case "aleatorio":
			v[i] = rand.Intn(k + 1)
		}
	}
	return v
}

func runTest(name string, v []int, sortFunc func([]int)) {
	arrCopy := make([]int, len(v))
	copy(arrCopy, v)
	start := time.Now()
	sortFunc(arrCopy)
	duration := time.Since(start)
	fmt.Printf("%-25s - Tempo: %v\n", name, duration)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("==> Fato 1: Impacto nulo em vetores pequenos (N = 10) ===")
	smallV := createArray(10, "aleatorio", 10)
	runTest("InsertionSort", smallV, insertionSort)
	runTest("MergeSort", smallV, MergeSort)
	runTest("CountingSort", smallV, func(v []int) { countingSort(v, 10) })
	fmt.Println()

	fmt.Println("==> Fato 2 e 3: Impacto em vetores grandes e disposicao (N = 100000) ===")
	largeRand := createArray(100000, "aleatorio", 100000)
	largeAsc := createArray(100000, "crescente", 100000)
	largeDesc := createArray(100000, "decrescente", 100000)

	fmt.Println("--> Aleatório")
	runTest("InsertionSort (O(n^2))", largeRand, insertionSort)
	runTest("MergeSort (O(n log n))", largeRand, MergeSort)
	runTest("CountingSort (O(n+k))", largeRand, func(v []int) { countingSort(v, 100000) })

	fmt.Println("\n--> Crescente (Ordenado)")
	runTest("InsertionSort (Omega(n))", largeAsc, insertionSort)
	runTest("MergeSort (O(n log n))", largeAsc, MergeSort)
	runTest("CountingSort (O(n+k))", largeAsc, func(v []int) { countingSort(v, 100000) })

	fmt.Println("\n--> Decrescente")
	runTest("InsertionSort (O(n^2))", largeDesc, insertionSort)
	runTest("MergeSort (O(n log n))", largeDesc, MergeSort)
	runTest("CountingSort (O(n+k))", largeDesc, func(v []int) { countingSort(v, 100000) })
	fmt.Println()

	fmt.Println("==> Fato 4: QuickSort em vetor ordenado (N = 10000) ===")
	mediumAsc := createArray(10000, "crescente", 10000)
	runTest("QuickSort (Pivô Fim)", mediumAsc, func(v []int) { quickSort(v, 0, len(v)-1, false) })
	runTest("QuickSort (Pivô Random)", mediumAsc, func(v []int) { quickSort(v, 0, len(v)-1, true) })
	fmt.Println()

	fmt.Println("==> Fato 5: Desempenho do CountingSort ===")
	kSmall := 1000
	kLarge := 1000000000
	vSmallK := createArray(10000, "aleatorio", kSmall)
	
	fmt.Printf("--> Experimento A: K proporcional (%d)\n", kSmall)
	runTest("CountingSort", vSmallK, func(v []int) { countingSort(v, kSmall) })
	
	fmt.Printf("\n--> Experimento B: K outlier (%d)\n", kLarge)
	vLargeK := createArray(10000, "aleatorio", kSmall) 
	vLargeK[len(vLargeK)-1] = kLarge 
	runTest("CountingSort", vLargeK, func(v []int) { countingSort(v, kLarge) })
}
