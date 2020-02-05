package node

// INode ...
type INode interface {
	ID() string
	Name() string
	Kind() Kind
	Run() interface{}
	Next() string
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
)

// Node ...
type Node struct {
}

// ID ...
func (n Node) ID() string {
	return ""
}

// Name ...
func (n Node) Name() string {
	return ""
}

// Kind ...
func (n Node) Kind() Kind {
	return KStart
}

// Next ...
func (n Node) Next() string {
	return ""
}

// Init ...
func (n Node) Init() {

}

// Run ...
func (n Node) Run() interface{} {
	return nil
}
