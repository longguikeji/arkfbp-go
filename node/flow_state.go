package node

// FlowState ...
type FlowState struct {
	data  interface{}
	steps []INode
	nodes map[string][]INode
}

// Push ...
func (st *FlowState) Push(n INode) {
	st.nodes[n.ID()] = append(st.nodes[n.ID()])
	st.steps = append(st.steps, n)
}

// Commit ...
func (st *FlowState) Commit(fn func(interface{}) interface{}) {
	st.data = fn(st.data)
}

// Query ...
func (st *FlowState) Query() interface{} {
	return st.data
}

// Inputs ...
func (st *FlowState) Inputs(nodeID string) interface{} {
	nodes, ok := st.nodes[nodeID]
	if !ok || len(nodes) == 0 {
		return nil
	}

	return nodes[len(nodes)-1].Inputs()
}

// Outputs ...
func (st *FlowState) Outputs(nodeID string) interface{} {
	nodes, ok := st.nodes[nodeID]
	if !ok || len(nodes) == 0 {
		return nil
	}

	return nodes[len(nodes)-1].Outputs()
}

// NewFlowState ...
func NewFlowState() *FlowState {
	flowState := &FlowState{
		nodes: make(map[string][]INode),
	}
	return flowState
}
