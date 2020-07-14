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
	//Reset backedges on every bfs call
	for _, node := range g.Nodes {
		node.Visited = false
		node.BackEdge = nil
	}

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

//canSpell solve with backwards approach
func (g Graph) canSpell() bool {

	//Find a path
	for g.bfs() == true {
		//Get Last node
		node := g.Nodes[len(g.Nodes)-1]
		var letter int
		//Traverse Backedges
		for node.Type != SOURCE {
			//Change Flow
			node.BackEdge.Original = 1
			node.BackEdge.Residual = 0
			node.BackEdge.Reverse.Original = 0
			node.BackEdge.Reverse.Residual = 1

			//Store letter
			if node.Type == LETTER {
				letter = node.ID
			} else if node.Type == DIE {
				g.NodeIDS[letter] = node.ID //Add letter to solution
			}

			//Advance
			node = node.BackEdge.To
		}
	}

	for index := g.minNodes + 1; index < len(g.Nodes); index++ {
		for _, n := range g.Nodes[index].Adj {

			//Look at sink
			if n.To.Type == SINK {
				//Not connected, cant spell
				if n.Residual != 1 {
					return false
				}
			}
		}
	}
	return true
}

//resetNodes adj list resid/origi
func (g Graph) resetNodes() {

}
