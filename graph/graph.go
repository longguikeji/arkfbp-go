package graph

import (
	"github.com/longguikeji/arkfbp-go/node"
)

// Node ...
type Node struct {
	ID   string
	Next string
}

// Graph ...
type Graph struct {
	nodes []node.INode
}

// FindEntryNode ...
func (g *Graph) FindEntryNode() node.INode {
	if len(g.nodes) == 0 {
		return nil
	}

	for _, n := range g.nodes {

		if n.Kind() == node.KStart {
			return n
		}

	}

	return g.nodes[0]
}

// FindNodeByID ...
func (g *Graph) FindNodeByID(id string) node.INode {
	for _, n := range g.nodes {
		if n.ID() == id {
			return n
		}
	}

	return nil
}

// Add ...
func (g *Graph) Add(node node.INode) {
	g.nodes = append(g.nodes, node)
}

// New ...
func New() *Graph {
	g := Graph{}
	return &g
}
