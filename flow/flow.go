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

// Run ...
func (f *Flow) Run() {
	g := f.g

	n := g.FindEntryNode()
	var nn node.INode
	for n != nil {
		copier.Copy(&nn, &n)
		fmt.Printf(">>> Executing Node %s\n", nn.ID())

		var outputs interface{}

		switch n.Kind() {
		case node.KIF:
			outputs = f.executeIFNode(nn)
			fmt.Println(n)
			fmt.Println(nn)
		default:
			outputs = f.executeNode(nn)
		}

		_ = outputs
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

func (f *Flow) executeNode(n node.INode) interface{} {
	return n.Run()
}

func (f *Flow) executeIFNode(n node.INode) interface{} {
	// t := reflect.TypeOf(n)
	// v := reflect.ValueOf(n)
	// for i := 0; i < t.NumField(); i++ {
	// 	f := t.Field(i)
	// 	val := v.Field(i).Interface()
	// 	if f.Type.Name() == "IFNode" {
	// 		val := val.(node.IFNode)
	// 		val.ExpressionRet = val.Expression()
	// 	}
	// 	break
	// }

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
