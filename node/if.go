package node

// IFNode ...
type IFNode struct {
	Node
	PositiveNext string
	NegativeNext string

	Expression    func() bool
	ExpressionRet bool

	Positive func() interface{}
	Negative func() interface{}
}

// Name ...
func (n *IFNode) Name() string {
	return "IF"
}

// Kind ...
func (n *IFNode) Kind() Kind {
	return KIF
}

// // Expression ...
// func (n IFNode) Expression() bool {
// 	fmt.Println("IFNode .... Expression")
// 	return true
// }

// Run ...
func (n *IFNode) Run() interface{} {
	if n.Expression != nil {
		n.ExpressionRet = n.Expression()
	} else {
		n.ExpressionRet = false
	}

	if n.ExpressionRet {
		if n.Positive != nil {
			return n.Positive()
		}
	}

	if n.Negative != nil {
		return n.Negative()
	}

	return nil
}

// // Positive ...
// func (n *IFNode) Positive() interface{} {
// 	return nil
// }

// // Negative ...
// func (n *IFNode) Negative() interface{} {
// 	return nil
// }
