package edge

import (
	"github.com/ablades/letterdice/node"
)

//Edge connects two nodes
type Edge struct {
	From     *node.Node
	To       *node.Node
	Reverse  *Edge //Edge in the opposite direction
	Origin   int   // Flow Field
	Residual int   // Flow Field
	//TODO: Init reverse edge inplace
}

func New(*node.Node from, *node.Node to){
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