package tree

type ITree interface {
	Search(value int) bool
	Add(value int)
	Min() int
	Max() int
	PreOrderNav() // Pré-Ordem: RED
	InOrderNav() // Em-Ordem: ERD
	PostOrderNav() // Pós-Ordem: EDR
	LevelOrderNav()
	Height() int
	Remove(value int) error
	Size() int
}