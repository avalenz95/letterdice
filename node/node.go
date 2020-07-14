package node

import (
	"github.com/ablades/letterdice/edge"
)

type Node struct {
	ID       int
	Letters  []bool
	Adj      []*edge.Edge //Adjacency List
	NodeType string       // type of node source, sink, die, letter
	Visited  bool
	BackEdge *edge.Edge
}

//New creates a new node
func New(id int, nodeType string) Node {
	n := Node{
		ID:      id,
		Letters: make([]bool, 26),
	}
	return n
}
