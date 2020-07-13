package node

type Node struct {
	id      int
	letters []rune
	adj     []*Edge
}
