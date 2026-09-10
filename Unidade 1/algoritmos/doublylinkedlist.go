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


func (list *DoublyLinkedList) addOnIndex(val int, index int) error {
  if index >= 0 && index <= list.inserted {
    newNode := &node2p(val: val)

    if list.inserted == 0 {
      list.head = newNode
      list.tail = newNode
      
    } else {
      if index == 0 {
        newNode.next = list.head
        list.head.prev = newNode
        list.head = newNode
        
      } else {
        aux := list.head
        for cont := 0; cont < index; cont++ {
          aux = aux.next;
        }
        newNode.prev = aux
        newNode.next = aux.next
        aux.next = newNode
        if newNode.next != nil {
          newNode.next.prev = newNode
        }
        
      }
      if index == list.inserted {
        list.tail = newNode
      }
    }
    return nil
    
  } else {
    return errors.New("Index out of bound")
  }
}
