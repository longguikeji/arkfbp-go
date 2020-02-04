package main

import (
	"fmt"

	"github.com/rockli/arkfbp-go/flow"
	"github.com/rockli/arkfbp-go/graph"
	"github.com/rockli/arkfbp-go/node"
)

// MyFlow ...
type MyFlow struct {
	flow.Flow
}

// Node1 ...
type Node1 struct {
	node.StartNode
}

// ID ...
func (n Node1) ID() string {
	return "Node1"
}

// Next ...
func (n Node1) Next() string {
	return "Node2"
}

// Run ...
func (n Node1) Run() interface{} {
	fmt.Println("Node1 Run...")
	return nil
}

// Node2 ...
type Node2 struct {
	node.NopNode
}

// ID ...
func (n Node2) ID() string {
	return "Node2"
}

// Next ...
func (n Node2) Next() string {
	return "Node3"
}

// Run ...
func (n Node2) Run() interface{} {
	fmt.Println("Node2 Run...")
	return nil
}

type Node3 struct {
	node.APINode
}

// ID ...
func (n Node3) ID() string {
	return "Node3"
}

func main() {
	flow := MyFlow{}

	g := graph.New()
	g.Add(Node2{})
	g.Add(Node1{})

	nn := Node3{
		node.APINode{
			//Mode:   "direct",
			Method: "GET",
			URL:    "https://api.github.com/repos/longguikeji/arkid-core/stargazers",
		},
	}
	nn.Mode = "direct"
	g.Add(nn)

	flow.SetGraph(g)
	flow.Run()
}
