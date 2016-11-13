package gobehaviortree

type NodeExecutor interface {
	Execute(node *Node, task TaskExecutor) error
}

type DefaultNodeExecutor struct {
}

func (e DefaultNodeExecutor) Execute(node *Node, task TaskExecutor) error {
	if node.State.Status == None {
		task.Init(node)
	}

	node.State.Status = Running

	result, _ := task.Execute(node)

	switch result {
	case Success, Failure:
		node.State.Status = Ready
	}

	return nil
}

type FakeNodeExecutor struct {
	Executed bool
}

func (e *FakeNodeExecutor) Execute(node *Node, task TaskExecutor) error {
	e.Executed = true
	return nil
}
