package gobehaviortree

type TaskExecutor interface {
	Init(node *Node) error
	Execute(node *Node) (TaskResult, error)
}

type JournalTaskExecutor struct {
	InitCount    int
	ExecuteCount int
	Result       TaskResult
}

func (e *JournalTaskExecutor) Init(node *Node) error {
	e.InitCount++
	return nil
}

func (e *JournalTaskExecutor) Execute(node *Node) (TaskResult, error) {
	e.ExecuteCount++
	return e.Result, nil
}
