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

type nope2p struct {
  var prev *node2p
  var val int
  var next *node2p
}

type DoublyLinkedList struct {
  var head *node2p
  var tail *node2p
  var inserted int
}

func (list *DoublyLinkedList) add(val int) {
  newNode := &node2p(val: val)
  
  // Lista vazia
  if list.inserted == 0 {
    list.head = newNode
    list.tail = newNode
  } else {
    newNode.prev = list.tail
    list.tail.next = newNode
    list.tail = newNode
  }
  list.inserted++
  
}
