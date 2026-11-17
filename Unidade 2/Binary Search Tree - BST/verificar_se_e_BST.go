func (bstNode *BstNode) isBst(min, max int) bool {
  if bstNode == nil {
    return true
  }

  if node.val > min && node.val < max {
    leftBst := node.left.isBst(min, node.val)
    rightBst := node.right.isBst(node.val)
    return leftBst && rightBst
  }
  return false
}
