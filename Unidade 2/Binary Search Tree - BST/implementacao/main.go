package tree

import "errors"
import "fmt"

type ITree interface {
	Search(value int) bool
	Add(value int)
	Min() int
	Max() int
	PreOrderNav()  // Pré-Ordem: RED
	InOrderNav()   // Em-Ordem: ERD
	PostOrderNav() // Pós-Ordem: EDR
	LevelOrderNav()
	Height() int
	Remove(value int) error
	Size() int
}

type BSTNode struct {
	left  *BSTNode
	val   int
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func NewBST() *BST {
	return &BST{}
}

func createBSTNode(val int) *BSTNode {
	return &BSTNode{
		val: val,
	}
}


func (t *BST) Search(value int) bool {
	return t.root.search(value)
}

func (t *BST) Add(value int) {
	if t.root == nil {
		t.root = createBSTNode(value)
		return
	}
	t.root.add(value)
}

func (t *BST) Min() int {
	if t.root == nil {
		return -1
	}
	return t.root.min()
}

func (t *BST) Max() int {
	if t.root == nil {
		return -1
	}
	return t.root.max()
}

func (t *BST) PreOrderNav() {
	t.root.preOrderNav()
	fmt.Println()
}

func (t *BST) InOrderNav() {
	t.root.inOrderNav()
	fmt.Println()
}

func (t *BST) PostOrderNav() {
	t.root.postOrderNav()
	fmt.Println()
}

func (t *BST) LevelOrderNav() {
	if t.root == nil {
		return
	}
	queue := []*BSTNode{t.root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%d ", current.val)

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
	fmt.Println()
}

func (t *BST) Height() int {
	return t.root.height()
}

func (t *BST) Remove(value int) error {
	if t.root == nil {
		return errors.New("árvore vazia")
	}
	var removed bool
	t.root, removed = t.root.remove(value)
	if !removed {
		return errors.New("valor não encontrado na árvore")
	}
	return nil
}

func (t *BST) Size() int {
	return t.root.size()
}


func (node *BSTNode) search(val int) bool {
	if node == nil {
		return false
	}
	if node.val == val {
		return true
	} else if val < node.val {
		return node.left.search(val)
	}
	return node.right.search(val)
}

func (node *BSTNode) add(val int) {
	if val <= node.val {
		if node.left == nil {
			node.left = createBSTNode(val)
		} else {
			node.left.add(val)
		}
	} else {
		if node.right == nil {
			node.right = createBSTNode(val)
		} else {
			node.right.add(val)
		}
	}
}

func (node *BSTNode) min() int {
	if node.left == nil {
		return node.val
	}
	return node.left.min()
}

func (node *BSTNode) max() int {
	if node.right == nil {
		return node.val
	}
	return node.right.max()
}

func (node *BSTNode) preOrderNav() {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.val)
	node.left.preOrderNav()
	node.right.preOrderNav()
}

func (node *BSTNode) inOrderNav() {
	if node == nil {
		return
	}
	node.left.inOrderNav()
	fmt.Printf("%d ", node.val)
	node.right.inOrderNav()
}

func (node *BSTNode) postOrderNav() {
	if node == nil {
		return
	}
	node.left.postOrderNav()
	node.right.postOrderNav()
	fmt.Printf("%d ", node.val)
}

func (node *BSTNode) height() int {
	if node == nil {
		return -1
	}
	leftH := node.left.height()
	rightH := node.right.height()
	if leftH > rightH {
		return leftH + 1
	}
	return rightH + 1
}

func (node *BSTNode) size() int {
	if node == nil {
		return 0
	}
	return 1 + node.left.size() + node.right.size()
}

func (node *BSTNode) remove(val int) (*BSTNode, bool) {
	if node == nil {
		return nil, false
	}

	var removed bool
	if val < node.val {
		node.left, removed = node.left.remove(val)
	} else if val > node.val {
		node.right, removed = node.right.remove(val)
	} else {
		// Caso 1 (folha) e Caso 2 (um filho)
		if node.left == nil {
			return node.right, true
		}
		if node.right == nil {
			return node.left, true
		}

		// Caso 3 (dois filhos): substitui pelo menor valor da subárvore direita
		minRight := node.right.min()
		node.val = minRight
		node.right, _ = node.right.remove(minRight)
		return node, true
	}

	return node, removed
}
