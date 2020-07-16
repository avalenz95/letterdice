package main

import "fmt"

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
func NewNode(id int, nodeType int) *Node {
	node := Node{
		ID:      id,
		Letters: make([]bool, 26),
		Type:    nodeType,
	}
	return &node
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
func NewEdge(from *Node, to *Node) *Edge {
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

	return &edge
}

//Graph contains nodes and all helper functions
type Graph struct {
	Nodes    []*Node
	NodeIDS  map[int]int
	minNodes int
}

//bfs breath first search on graph
func (g Graph) bfs() bool {

	fmt.Println("BFS")
	//Reset backedges on every bfs call
	for _, node := range g.Nodes {
		node.Visited = false
		node.BackEdge = nil
	}

	var queue []int
	queue = append(queue, SOURCE)

	//Continue until queue is empty
	for len(queue) > 0 {
		fmt.Printf("queue: %+v \n", queue)
		//dequeue item
		nodeID := queue[0]
		queue = queue[1:]
		//mark node as visited
		g.Nodes[nodeID].Visited = true

		//Iterate over adj list
		for _, edge := range g.Nodes[nodeID].Adj {
			nextNode := edge.To
			e := edge

			//Add unvisited nodes flowing in the right direction to queue
			if nextNode.Visited == false && e.Original == 1 {
				//Set the backedge for the node
				nextNode.BackEdge = e.Reverse
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

	//
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

//Inputs passed into the program
type Inputs struct {
	Word string
	Die  []string
}

func main() {

	input := Inputs{
		Word: "RAGE",
		Die: []string{
			"ENG",
			"SAA",
			"PRR",
			"EAE",
		},
	}

	var graph Graph
	graph.NodeIDS = make(map[int]int)
	nextID := 0
	//Create Source add to graph
	source := NewNode(0, SOURCE)
	graph.Nodes = append(graph.Nodes, source)
	//Create Sink Node
	sink := NewNode(0, SINK)

	for _, dice := range input.Die {
		nextID++
		//Connect to source
		node := NewNode(nextID, DIE)
		edge := NewEdge(source, node)
		source.Adj = append(source.Adj, edge)
		//Init dice letters slice
		for _, letter := range dice {
			pos := letter - 'A'
			node.Letters[pos] = true
		}
		node.Type = DIE
		graph.Nodes = append(graph.Nodes, node)
	}
	graph.minNodes = nextID

	//Create Letter Nodes
	for _, letter := range input.Word {
		nextID++
		node := NewNode(nextID, LETTER)
		pos := letter - 'A'
		node.Letters[pos] = true

		//Set Edges for each letter
		for index := 1; index < graph.minNodes; index++ {
			if graph.Nodes[index].Letters[pos] == true {
				edge := NewEdge(graph.Nodes[index], node)
				graph.Nodes[index].Adj = append(graph.Nodes[index].Adj, edge)
				node.Adj = append(node.Adj, edge.Reverse)
			}
		}
		edge := NewEdge(node, sink)
		node.Adj = append(node.Adj, edge)
		graph.Nodes = append(graph.Nodes, node)
	}
	//Add sink to end of graph
	sink.ID = len(graph.Nodes)

	graph.Nodes = append(graph.Nodes, sink)
	for _, node := range graph.Nodes {
		fmt.Printf("ID: %v, Type: %v \n", node.ID, node.Type)
		for _, edge := range node.Adj {
			fmt.Printf("edge for %v , %v -->  %v \n", node.ID, edge.From.ID, edge.To.ID)
		}
	}
	fmt.Println("-----------------")
	if graph.canSpell() == true {
		fmt.Println("Can Spell")
	} else {
		fmt.Println("Can't Spell")
	}

	fmt.Println("-----------------")
	for _, node := range graph.Nodes {
		fmt.Printf("ID: %v, Type: %v \n", node.ID, node.Type)
		for _, edge := range node.Adj {
			fmt.Printf("edge for %v , %v -->  %v \n", node.ID, edge.From.ID, edge.To.ID)
		}
	}

}
