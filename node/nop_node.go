package node

import "fmt"

// NopNode ...
type NopNode struct {
	Node
}

// Name ...
func (n *NopNode) Name() string {
	return "Nop"
}

// Kind ...
func (n *NopNode) Kind() Kind {
	return KNop
}

// Run ...
func (n *NopNode) Run() interface{} {
	fmt.Println("NopNode...")
	return nil
}
