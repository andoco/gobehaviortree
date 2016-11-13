package gobehaviortree

import "fmt"

func NewDefaultScheduler() DefaultScheduler {
	return DefaultScheduler{nodeExecutor: DefaultNodeExecutor{}}
}

// Scheduler handled the scheduling of nodes in the tree for execution.
type Scheduler interface {
	Schedule(node *Node) error
	Execute(executors map[string]TaskExecutor) error
}

type DefaultScheduler struct {
	scheduled    []*Node
	nodeExecutor NodeExecutor
}

func (s *DefaultScheduler) Schedule(node *Node) error {
	if node == nil {
		return fmt.Errorf("cannot schedule a nil node")
	}

	s.scheduled = append(s.scheduled, node)
	return nil
}

func (s *DefaultScheduler) Execute(executors map[string]TaskExecutor) error {
	for _, node := range s.scheduled {
		taskExecutor, ok := executors[node.Name]
		if !ok {
			return fmt.Errorf("no task executor could be found for the task name %d", node.Name)
		}
		s.nodeExecutor.Execute(node, taskExecutor)
	}

	s.scheduled = nil

	return nil
}
