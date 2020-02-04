package flow

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/copier"
	"github.com/rockli/arkfbp-go/graph"
	"github.com/rockli/arkfbp-go/node"
)

// Flow ...
type Flow struct {
	g *graph.Graph
}

func DeepFields(iface interface{}) []reflect.Value {
	fields := make([]reflect.Value, 0)
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			fields = append(fields, DeepFields(v.Interface())...)
		default:
			fields = append(fields, v)
		}
	}

	return fields
}

// Run ...
func (f *Flow) Run() {
	fmt.Println("Flow::Run")

	g := f.g

	n := g.FindEntryNode()
	var nn node.INode
	for n != nil {
		copier.Copy(&nn, &n)
		outputs := nn.Run()
		fmt.Printf("outputs: %v\n", outputs)
		n = g.FindNodeByID(n.Next())
	}

}

// SetGraph ...
func (f *Flow) SetGraph(g *graph.Graph) {
	f.g = g
}

// New ...
func New() *Flow {
	f := Flow{}

	return &f
}
