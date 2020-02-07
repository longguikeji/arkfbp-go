package flow

import (
	"fmt"
	"reflect"
	"strings"

	"git.intra.longguikeji.com/longguikeji/arkfbp-go/graph"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/node"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/request"
	"git.intra.longguikeji.com/longguikeji/arkfbp-go/response"
	"github.com/jinzhu/copier"
)

// Flow ...
type Flow struct {
	// CreateGraph is called during the flow executing to generate the graph dependency
	CreateGraph func() *graph.Graph

	// Hooks
	BeforeInitialize func()
	Initialized      func()
	BeforeExecute    func()
	Executed         func()
	BeforeDestroy    func()

	g        *graph.Graph
	state    *node.FlowState
	appState *node.AppState

	request  *request.Request
	response *response.Response
}

// Run ...
func (f *Flow) Run() interface{} {
	if f.BeforeInitialize != nil {
		f.BeforeInitialize()
	}
	f.init()
	if f.Initialized != nil {
		f.Initialized()
	}

	g := f.g

	n := g.FindEntryNode()
	var (
		nn          node.INode
		lastOutputs interface{}
		outputs     interface{}
	)

	for n != nil {
		copier.Copy(&nn, &n)
		fmt.Printf(">>> Executing Node %s\n", nn.ID())

		nn.SetInputs(lastOutputs)
		nn.SetRequest(f.request)
		nn.SetResponse(f.response)
		nn.SetState(f.state)
		nn.SetAppState(f.appState)

		switch n.Kind() {
		case node.KIF:
			outputs = f.executeIFNode(nn)
		case node.KTest:
			outputs = f.executeTestNode(nn)
		default:
			outputs = f.executeNode(nn)
		}

		lastOutputs = outputs
		nn.SetOutputs(lastOutputs)

		f.state.Push(nn)

		next := f.getNextNodeID(nn)
		if next == "" {
			break
		}

		n = g.FindNodeByID(next)
	}

	return lastOutputs
}

// SetGraph ...
func (f *Flow) SetGraph(g *graph.Graph) {
	f.g = g
}

// SetRequest ...
func (f *Flow) SetRequest(r *request.Request) {
	f.request = r
}

// Request ...
func (f *Flow) Request() *request.Request {
	return f.request
}

// SetResponse ...
func (f *Flow) SetResponse(r *response.Response) {
	f.response = r
}

// Response ...
func (f *Flow) Response() *response.Response {
	return f.response
}

func (f *Flow) init() {
	f.state = node.NewFlowState()

	if f.CreateGraph != nil {
		f.SetGraph(f.CreateGraph())
	}

	if f.response == nil {
		f.response = new(response.Response)
	}
}

func (f *Flow) executeTestNode(n node.INode) interface{} {
	var (
		tn node.INode
	)
	copier.Copy(&tn, &n)

	nodeType := reflect.TypeOf(tn)
	nodeValue := reflect.ValueOf(tn).Elem()
	for i := 0; i < nodeType.NumMethod(); i++ {
		method := nodeType.Method(i)
		if strings.HasPrefix(method.Name, "Test") {
			// Execute the flow

			for j := 0; j < nodeValue.NumField(); j++ {
				t := nodeValue.Field(j)
				if t.Type().Name() == "TestNode" {
					tt := t.Addr().Interface().(*node.TestNode)
					fmt.Printf("> Execute test flow: %s\n", reflect.ValueOf(tt.Flow).Elem().Type().Name())
					tt.Flow.Run()
					fmt.Printf("> End the execution\n")
					break
				}
			}

			// Invoke TestCase
			method.Func.Call([]reflect.Value{
				reflect.ValueOf(tn),
			})
		}
	}

	return nil
}

func (f *Flow) executeNode(n node.INode) interface{} {
	return n.Run()
}

func (f *Flow) executeIFNode(n node.INode) interface{} {
	return n.Run()
}

func (f *Flow) getNextNodeID(n node.INode) string {
	switch n.Kind() {
	case node.KIF:
		v := reflect.ValueOf(n)
		v = reflect.Indirect(v)
		t := reflect.TypeOf(reflect.Indirect(v))
		for i := 0; i < t.NumField(); i++ {
			val := v.Field(i).Interface()
			fmt.Println(reflect.Indirect(v.Field(i)).Type().Name())
			if reflect.Indirect(v.Field(i)).Type().Name() == "IFNode" {
				val := val.(node.IFNode)
				if val.ExpressionRet {
					return val.PositiveNext
				}

				return val.NegativeNext
			}
			break
		}

		return ""
	default:
		return n.Next()
	}

}

// New ...
func New() *Flow {
	f := Flow{}

	return &f
}
