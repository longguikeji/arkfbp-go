package node

// INode ...
type INode interface {
	ID() string
	Name() string
	Kind() Kind
	Run() interface{}
	Next() string

	Inputs() interface{}
	SetInputs(interface{})

	Outputs() interface{}
	SetOutputs(interface{})
}

// Kind ...
type Kind int

const (
	KStart = iota
	KStop
	KNop
	KFunction
	KAPI
	KIF
	KLoop
)

// Node ...
type Node struct {
	inputs  interface{}
	outputs interface{}
}

// ID ...
func (n *Node) ID() string {
	return ""
}

// Inputs ...
func (n *Node) Inputs() interface{} {
	return n.inputs
}

// SetInputs ...
func (n *Node) SetInputs(v interface{}) {
	n.inputs = v
}

// Outputs ...
func (n *Node) Outputs() interface{} {
	return n.outputs
}

// SetOutputs ...
func (n *Node) SetOutputs(v interface{}) {
	n.outputs = v
}

// Name ...
func (n *Node) Name() string {
	return ""
}

// Kind ...
func (n *Node) Kind() Kind {
	return KStart
}

// Next ...
func (n *Node) Next() string {
	return ""
}

// Init ...
func (n *Node) Init() {

}

// Run ...
func (n *Node) Run() interface{} {
	return nil
}
