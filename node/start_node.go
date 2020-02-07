package node

import "fmt"

// StartNode ...
type StartNode struct {
	Node
}

// Name ...
func (n *StartNode) Name() string {
	return "Start"
}

// Kind ...
func (n *StartNode) Kind() Kind {
	return KStart
}

// Run ...
func (n *StartNode) Run() interface{} {
	fmt.Println("StartNode...")
	return nil
}
