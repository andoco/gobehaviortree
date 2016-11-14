package gobehaviortree

type SequenceTaskExecutor struct {
}

func (e SequenceTaskExecutor) Init(node *Node) error {
	node.State.ExecutorState = newSequenceState()
	return nil
}

func (e SequenceTaskExecutor) Execute(node *Node, scheduler ScheduleFunc) (TaskResult, error) {
	state := node.State.ExecutorState.(sequenceState)

	if state.activeChild == -1 {
		state.activeChild = 0
		scheduler(node.Children[state.activeChild])
		return Pending, nil
	}

	if node.Children[state.activeChild].State.Result == Failure {
		state.activeChild = -1
		return Failure, nil
	} else {
		state.activeChild++
		if state.activeChild == len(node.Children) {
			state.activeChild = -1
			return Success, nil
		}
		scheduler(node.Children[state.activeChild])
		return Pending, nil
	}
}

type sequenceState struct {
	activeChild int
}

func newSequenceState() *sequenceState {
	return &sequenceState{activeChild: -1}
}
