package main 

import "fmt"
import "errors"

type list interface {
  add(val int)
  addOnIndex(val int, index int) error
  get(index int) (int, error)
  set(val int, index int) (int, error)
  remove() (int, error)
  removeOnIndex(index int) (int, error)
  size() int
}

type linkedlist struct {
  var head *node1p;
  var inserted int;
}

type node1p struct {
  var val int;
  var next *node1p;
}

func (list *linkedlist) add(val int) {
  if list.head == nil {
    newNode := &node1p{val: val, next: nil}
    list.head = newNode
    list.inserted++
  } else {
  
    aux := list.head;
  
    for aux.next != nil {
      aux = aux.next
    }
  
    newNode := &node1p{val: val, next: nil}
    aux.next = newNode
    list.inserted++
  }
}
