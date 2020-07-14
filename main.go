type Node struct {
	ID       int
	Letters  []bool
	Adj      []*edge.Edge //Adjacency List
	NodeType string       // type of node source, sink, die, letter
	Visited  bool
	BackEdge *edge.Edge
}

//Edge connects two nodes
type Edge struct {
	From     *node.Node
	To       *node.Node
	Reverse  *Edge //Edge in the opposite direction
	Origin   int   // Flow Field
	Residual int   // Flow Field
	//TODO: Init reverse edge inplace
}

//New creates a new node
func NewNode(id int, nodeType string) Node {
	n := Node{
		ID:      id,
		Letters: make([]bool, 26),
	}
	return n
}

func NewEdge(*node.Node from, *node.Node to){
	var edge Edge
	var edge reverseEdge
	
	edge := {
		From: &from,
		To: &to,
		Reverse: &reverseEdge,
		Origin: 1,
		Residual: 0,
	}

	reverseEdge := {
		From: &to, 
		To: &from,
		Reverse:  &edge,
		Origin: 0,
		Residual: 1,
	}
}