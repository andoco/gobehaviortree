package gobehaviortree

import (
	"fmt"
	"testing"
)

func TestSequenceExecution(t *testing.T) {
	testCases := []struct {
		numChildren   int
		activeChild   int
		childResult   TaskResult
		scheduleChild int
		result        TaskResult
		name          string
	}{
		{
			3,
			-1,
			-1,
			0,
			Pending,
			"no active child",
		},
		{
			3,
			0,
			Success,
			1,
			Pending,
			"first child active and returns success",
		},
		{
			3,
			0,
			Failure,
			-1,
			Failure,
			"first child active and returns failure",
		},
		{
			3,
			2,
			Success,
			-1,
			Success,
			"last child active and returns success",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			seq := SequenceTaskExecutor{}
			seqNode := NewNode("seq")

			for i := 0; i < tc.numChildren; i++ {
				seqNode.Children = append(seqNode.Children, NewNode(fmt.Sprintf("test%d", i)))
			}

			if tc.activeChild > -1 {
				seqNode.Children[tc.activeChild].State.Result = tc.childResult
			}

			var scheduledNode *Node = nil
			scheduler := func(node *Node) { scheduledNode = node }

			seqNode.State.ExecutorState = sequenceState{activeChild: tc.activeChild}

			result, err := seq.Execute(seqNode, scheduler)

			if err != nil {
				t.Fatalf("received error %v", err)
			}

			if tc.scheduleChild > -1 {
				if scheduledNode != seqNode.Children[tc.scheduleChild] {
					t.Errorf("expected child %v to be scheduled, got %v", seqNode.Children[tc.scheduleChild], scheduledNode)
				}
			} else {
				if scheduledNode != nil {
					t.Fatalf("expected no child to be scheduled, got %v", scheduledNode)
				}
			}

			if result != tc.result {
				t.Errorf("expected result %d, got %d", tc.result, result)
			}
		})
	}
}
