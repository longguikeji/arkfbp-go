package node

import "git.intra.longguikeji.com/longguikeji/arkfbp-go/intr"

// TestNode ...
type TestNode struct {
	Node

	Flow  intr.IFlow
	Start interface{}
	Stop  interface{}
}

// Name ...
func (n *TestNode) Name() string {
	return "Test"
}

// Kind ...
func (n *TestNode) Kind() Kind {
	return KTest
}

// Run ...
func (n *TestNode) Run() interface{} {
	return nil
}
