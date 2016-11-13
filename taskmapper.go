package gobehaviortree

import (
	"errors"
	"fmt"
)

type TaskExecutorMapper interface {
	Get(node *Node) (TaskExecutor, error)
}

type DefaultTaskExecutorMapper struct {
	executors map[string]TaskExecutor
}

func (m DefaultTaskExecutorMapper) Get(node *Node) (TaskExecutor, error) {
	if node.Name == "" {
		return nil, errors.New("error getting the executor for a node with no name")
	}

	for executor, ok := m.executors[node.Name]; ok; {
		return executor, nil
	}

	return nil, fmt.Errorf("task executor not found for node with name %s", node.Name)
}
