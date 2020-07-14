package main

type Node struct {
	ID       int
	Letters  []bool
	Adj      []*Edge //Adjacency List
	NodeType string  // type of node source, sink, die, letter
	Visited  bool
	BackEdge *Edge
}

//NewNode creates a new node
func NewNode(id int, nodeType string) Node {
	n := Node{
		ID:      id,
		Letters: make([]bool, 26),
	}
	return n
}

//Edge connects two nodes
type Edge struct {
	From     *Node
	To       *Node
	Reverse  *Edge //Edge in the opposite direction
	Origin   int   // Flow Field
	Residual int   // Flow Field
	//TODO: Init reverse edge inplace
}

//NewEdge Creates edges from -> to and to -> from
func NewEdge(from *Node, to *Node) {
	var edge Edge
	var reverseEdge Edge

	edge = Edge{
		From:     from,
		To:       to,
		Reverse:  &reverseEdge,
		Origin:   1,
		Residual: 0,
	}

	reverseEdge = Edge{
		From:     to,
		To:       from,
		Reverse:  &edge,
		Origin:   0,
		Residual: 1,
	}
}
