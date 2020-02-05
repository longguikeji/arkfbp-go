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

// // Node1 ...
// type Node1 struct {
// 	node.StartNode
// }

// // ID ...
// func (n Node1) ID() string {
// 	return "Node1"
// }

// // Next ...
// func (n Node1) Next() string {
// 	return "Node2"
// }

// // Run ...
// func (n Node1) Run() interface{} {
// 	fmt.Println("Node1 Run...")
// 	return nil
// }

// // Node2 ...
// type Node2 struct {
// 	node.NopNode
// }

// // ID ...
// func (n Node2) ID() string {
// 	return "Node2"
// }

// // Next ...
// func (n Node2) Next() string {
// 	return "Node3"
// }

// // Run ...
// func (n Node2) Run() interface{} {
// 	fmt.Println("Node2 Run...")
// 	return nil
// }

// Node3 ...
type Node3 struct {
	node.APINode
}

// ID ...
func (n *Node3) ID() string {
	return "Node3"
}

// Next ...
func (n *Node3) Next() string {
	return "Node4"
}

// Node4 ...
type Node4 struct {
	node.IFNode
}

// ID ...
func (n Node4) ID() string {
	return "Node4"
}

// Positive ...
func (n Node4) Positive() interface{} {
	fmt.Println("Node4 IF Postive")
	return nil
}

// Negative ...
func (n Node4) Negative() interface{} {
	fmt.Println("Node4 IF Negative")
	return nil
}

// Node5 ...
type Node5 struct {
	node.FunctionNode
}

// ID ...
func (n *Node5) ID() string {
	return "Node5"
}

// Run ...
func (n *Node5) Run() interface{} {
	fmt.Println("Node5 Run...")
	return nil
}

// Node6 ...
type Node6 struct {
	node.FunctionNode
}

// ID ...
func (n *Node6) ID() string {
	return "Node6"
}

// Run ...
func (n *Node6) Run() interface{} {
	fmt.Println("Node6 Run...")
	return nil
}

// TestLoopNode ...
type TestLoopNode struct {
	node.LoopNode
}

// ID ...
func (n *TestLoopNode) ID() string {
	return "TestLoopNode"
}

func main() {
	flow := MyFlow{}

	g := graph.New()
	// // g.Add(Node2{})
	// // g.Add(Node1{})

	// nn := &Node3{
	// 	node.APINode{
	// 		Mode:   "direct",
	// 		Method: "GET",
	// 		URL:    "https://api.github.com/repos/longguikeji/arkid-core/stargazers",
	// 	},
	// }
	// g.Add(nn)

	// g.Add(&Node4{
	// 	node.IFNode{
	// 		Expression: func() bool {
	// 			fmt.Println("Node4.Expression")
	// 			return true
	// 		},
	// 		PositiveNext: "Node5",
	// 		NegativeNext: "Node6",
	// 	},
	// })

	// g.Add(&Node5{})
	// g.Add(&Node6{})

	n := TestLoopNode{
		node.LoopNode{
			Init:    func() {},
			Cond:    func() bool { return true },
			Post:    func() {},
			Process: func() { fmt.Println("Loop.Process") },
		},
	}

	g.Add(&n)

	flow.SetGraph(g)
	flow.Run()
}
