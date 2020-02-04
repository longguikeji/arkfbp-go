package node

// StopNode ...
type StopNode struct {
	Node
}

// Name ...
func (n StopNode) Name() string {
	return "Stop"
}

// Kind ...
func (n StopNode) Kind() Kind {
	return KStop
}

// Run ...
func (n StopNode) Run() interface{} {
	return nil
}
