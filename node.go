package gobehaviortree

// Node holds information and state for a task in the behaviour tree.
type Node struct {
	Id       int
	Name     string
	Parent   *Node
	Children []*Node
	State    *NodeState
}

// NodeState holds the mutable state associated with a task in the behaviour tree.
type NodeState struct {
	Status TaskStatus
	Result TaskResult
}

func NewNode(name string) *Node {
	return &Node{Name: name, State: &NodeState{}}
}
