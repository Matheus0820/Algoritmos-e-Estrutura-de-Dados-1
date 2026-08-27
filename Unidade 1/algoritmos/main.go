package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func divisao(a int, b int) float64 {
	// denominador := 10.0
	var denominador float64
	denominador = 10.0

	return float64(a + b) / denominador
}

func contaAte100() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	// Go não tem while, mas o for funcina igualmente da seguinte forma
	// i := 0
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // Terceira forma
	// for i := range 10 {
	// 	fmt.Println(i)
	// }
}

func eMaior(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println("Olá, Mundo!")
	fmt.Println(soma(1, 3))
	contaAte100()
	fmt.Println(eMaior(10, 20))
}