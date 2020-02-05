package node

// LoopNode ...
type LoopNode struct {
	Node

	Init func()
	Cond func() bool
	Post func()

	Process func()
}

// Name ...
func (n *LoopNode) Name() string {
	return "Function"
}

// Kind ...
func (n *LoopNode) Kind() Kind {
	return KLoop
}

// Run ...
func (n *LoopNode) Run() interface{} {
	if n.Init != nil {
		n.Init()
	}

	for {
		if n.Cond != nil {
			if !n.Cond() {
				break
			}
		}

		if n.Process != nil {
			n.Process()
		}

		if n.Post != nil {
			n.Post()
		}
	}

	return nil
}
