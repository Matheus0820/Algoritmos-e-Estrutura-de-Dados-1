packege main

type BSTNode Struct {
	var left *BSTNode
	var val int
	var right *BSTNode

}

func createBSTNode(val int) *BSTNode {
	return &BSTNode {
		val: val,
	}
}

func (node *BSTNode) Search(val int) bool{
	if node == nil {
		return false
	} else {
		if node.val == val {
			return true
		} else if val < node.val { // deve estar na esquerda
			return node.left.Search(val)
		} else {
			return node.right.Search(val)
		}
	}
}

func (node *BSTNode) Add(val int) {
	if val <= node.val {
		if node.left == nil {
			node.left == createBSTNode(val)
		} else {
			node.left.Add(val)
		}
	} else {
		if node.right == nil {
			node.right == createBSTNode(val)
		} else {
			node.right.Add(val)
		}
	}
}

func (node *BSTNode) Min() int {
	if node.left == nil {
		return node.val
	} else {
		return node.left.Min()
	}
}

func (node *BSTNode) Max() int {
	if node.right == nil {
		return node.val
	} else {
		return node.right.Max()
	}
}

func (node *BSTNode) PreOrderNav() {

}

func (node *BSTNode) InOrderNav() {

}

func main() {
	root := createBSTNode(10)
}