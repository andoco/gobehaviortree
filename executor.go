package gobehaviortree

type ScheduleFunc func(node *Node)

type TaskExecutor interface {
	Init(node *Node) error
	Execute(node *Node, schedule ScheduleFunc) (TaskResult, error)
}

func NewJournalTaskExecutor(result TaskResult) JournalTaskExecutor {
	return JournalTaskExecutor{Result: result, Counts: &TaskExecutionCount{}}
}

type TaskExecutionCount struct {
	InitCount    int
	ExecuteCount int
}

type JournalTaskExecutor struct {
	Result TaskResult
	Counts *TaskExecutionCount
}

func (e JournalTaskExecutor) Init(node *Node) error {
	e.Counts.InitCount++
	return nil
}

func (e JournalTaskExecutor) Execute(node *Node, schedule ScheduleFunc) (TaskResult, error) {
	e.Counts.ExecuteCount++
	return e.Result, nil
}
