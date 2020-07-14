package main

//NodeTypes
const (
	SOURCE int = 0
	DIE    int = 1
	LETTER int = 2
	SINK   int = 3
)

//Node component of graph
type Node struct {
	ID       int
	Letters  []bool
	Adj      []*Edge //Adjacency List
	Type     int     // type of node source, sink, die, letter
	Visited  bool
	BackEdge *Edge
}

//NewNode creates a new node
func NewNode(id int, nodeType int) Node {
	n := Node{
		ID:      id,
		Letters: make([]bool, 26),
		Type:    nodeType,
	}
	return n
}

//Edge connects two nodes
type Edge struct {
	From     *Node
	To       *Node
	Reverse  *Edge //Edge in the opposite direction
	Original int   // Flow Field
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
		Original: 1,
		Residual: 0,
	}

	reverseEdge = Edge{
		From:     to,
		To:       from,
		Reverse:  &edge,
		Original: 0,
		Residual: 1,
	}
}

//Graph contains nodes and all helper functions
type Graph struct {
	Nodes    []*Node
	NodeIDS  map[int]int
	minNodes int
}

//bfs breath first search on graph
func (g Graph) bfs() bool {
	queue := make([]int, 10)
	queue = append(queue, SOURCE)

	//Continue until queue is empty
	for len(queue) > 0 {
		//dequeue item
		nodeID := queue[0]
		queue = queue[1:]

		//mark node as visited
		g.Nodes[nodeID].Visited = true
		var nextNode *Node
		var edge *Edge

		//Iterate over adj list
		for _, node := range g.Nodes[nodeID].Adj {
			nextNode = node.To
			edge = node

			//Add unvisited nodes flowing in the right direction to queue
			if nextNode.Visited == false && edge.Original == 1 {
				//Set the backedge for the node
				nextNode.BackEdge = edge.Reverse
				//Add to queue
				queue = append(queue, nextNode.ID)

				//Path has been found
				if nextNode.Type == SINK {
					return true
				}
			}
		}
	}
	return false
}

func (g Graph) canSpell() bool {

}

//resetNodes adj list resid/origi
func (g Graph) resetNodes() {

}
