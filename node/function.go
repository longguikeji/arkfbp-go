package node

// FunctionNode ...
type FunctionNode struct {
	Node
}

// Name ...
func (n FunctionNode) Name() string {
	return "Function"
}

// Kind ...
func (n FunctionNode) Kind() Kind {
	return KFunction
}

// Run ...
func (n FunctionNode) Run() interface{} {
	return nil
}
