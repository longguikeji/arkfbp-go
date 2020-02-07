package node

import (
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/request"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/response"
)

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

	Request() *request.Request
	SetRequest(*request.Request)
	Response() *response.Response
	SetResponse(*response.Response)

	State() *FlowState
	SetState(*FlowState)

	AppState() *AppState
	SetAppState(*AppState)
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
	KTest
)

// Node ...
type Node struct {
	// Hooks
	BeforeInitialize func()
	Initialized      func()
	BeforeExecute    func()
	Executed         func()
	BeforeDestroy    func()

	inputs  interface{}
	outputs interface{}

	request  *request.Request
	response *response.Response

	state    *FlowState
	appState *AppState
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

// SetRequest ...
func (n *Node) SetRequest(r *request.Request) {
	n.request = r
}

// Request ...
func (n *Node) Request() *request.Request {
	return n.request
}

// SetResponse ...
func (n *Node) SetResponse(r *response.Response) {
	n.response = r
}

// Response ...
func (n *Node) Response() *response.Response {
	return n.response
}

// SetState ...
func (n *Node) SetState(s *FlowState) {
	n.state = s
}

// State ...
func (n *Node) State() *FlowState {
	return n.state
}

// SetAppState ...
func (n *Node) SetAppState(s *AppState) {
	n.appState = s
}

// AppState ...
func (n *Node) AppState() *AppState {
	return n.appState
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
