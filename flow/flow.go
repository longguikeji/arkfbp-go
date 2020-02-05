package flow

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
	"github.com/rockli/arkfbp-go/graph"
	"github.com/rockli/arkfbp-go/node"
	"github.com/rockli/arkfbp-go/state"
)

// Flow ...
type Flow struct {
	BeforeInitialize func()
	Initialized      func()
	BeforeExecute    func()
	Executed         func()
	BeforeDestroy    func()

	g     *graph.Graph
	state *state.FlowState
}

// Run ...
func (f *Flow) Run() {
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
	)

	for n != nil {
		copier.Copy(&nn, &n)
		fmt.Printf(">>> Executing Node %s\n", nn.ID())

		var outputs interface{}

		nn.SetInputs(lastOutputs)

		switch n.Kind() {
		case node.KIF:
			outputs = f.executeIFNode(nn)
		default:
			outputs = f.executeNode(nn)
		}

		lastOutputs = outputs
		nn.SetOutputs(lastOutputs)

		f.state.Push(nn)

		next := f.getNextNodeID(nn)
		fmt.Printf(">>> Next node: %s\n", next)
		if next == "" {
			break
		}

		n = g.FindNodeByID(next)
	}
}

// SetGraph ...
func (f *Flow) SetGraph(g *graph.Graph) {
	f.g = g
}

func (f *Flow) init() {
	f.state = state.NewFlowState()
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
