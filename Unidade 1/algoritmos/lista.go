package main 

import "fmt"
import "errors"

type list interface { // Tipo abstrato de dado
	add(val int) // Adicionar elemento no final
	addOnIndex(val int, index int) error // Adicionar elemento em uma posição específica sem remover o antes existente naquela posição
	get(index int) (int, error) // Pegar o elemento de uma posição específica
	set(val int, index int) (int, error) // Adicionar elemento em uma posição específica
	remove() (int, error) // Remover elemento do final da lista
	removeOnIndex(index int) (int, error) // Remover elemento em uma posição específica
	size() int // Tamanho da lista
}

type arraylist struct {
	v []int
	inserted int
}

func (list *arraylist) inceraseCapacity() {
	newVector := make([]int, 2*len(list.v))

	for i := 0; i < len(list.v); i++ {
		newVector[i] = list.v[i]
	}

	list.v = newVector
}

func (list *arraylist) moveVectorRight(start int) {
	for i := list.inserted; i > start; i-- {
		list.v[i] = list.v[i - 1]
	}
}

func (list *arraylist) moveVectorLeft(start int) {
	for i := i < list.inserted - 1; i++ {
		list.v[i] = list.v[i + 1]
	}
}

func (list *arraylist) add(val int) {
	if list.inserted == len(list.v) {
		list.inceraseCapacity()
	}
	list.v[list.inserted] = val
	list.inserted++
}

func (list *arraylist) addOnIndex(val int, index int) {
	if list.inserted == len(list.v) {
		list.inceraseCapacity()
	}

	list.moveVectorRight(index)
	list.v[index] = val
	list.inserted++
}

func (list *arraylist) get(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		return list.v[index], nil
	} else {
		return -1, errors.New("Index out of bounds")
	}
}

func (list *arraylist) set(val int, index int) error {
	if index >= 0 && index < list.inserted {
		list.v[index] = val
		return nil

	} else {
		return errors.New("Index out of bounds")
	}
}

func (list *arraylist) remove() (int, error) {
	if list.inserted == 0 {
		return -1, errors.New("Can't remove from empty list")
	} else {
		temp := list.v[list.inserted - 1]
		list.inserted--
		
		return temp, nil
	}
}

func (list *arraylist) removeOnIndex(index int) (int, error) {
	if index >= 0 && index < list.inserted {
		valRemoved := list.v[index]
		moveVectorLeft(index)
		list.inserted--
		
		return valRemoved, nil
	}

	return -1, errors.New("Index out of bounds")
}

func main() {
	// list := &arraylist{v: make([]int, 10), inserted: 0}
	vector := make([]int, 10)
	list := &arraylist{v: vector, inserted: 0}

	for i := 0; i < 15; i++ {
		list.add(i)
	}

	for i := 0; i < 15; i++ {
		fmt.Printf("%d, ", list.v[i])
	}
}