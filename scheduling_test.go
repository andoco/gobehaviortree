package gobehaviortree

import (
	"reflect"
	"testing"
)

func TestNewScheduler(t *testing.T) {
	scheduler := NewDefaultScheduler()
	schedulerType := reflect.TypeOf(scheduler)
	expectedType := reflect.TypeOf(&DefaultScheduler{})

	if schedulerType != expectedType {
		t.Errorf("expected scheduler of type %v, got %v", expectedType, schedulerType)
	}

	if scheduler.nodeExecutor == nil {
		t.Errorf("expected node executor, got nil")
	}
}

func TestScheduleNode(t *testing.T) {
	scheduler := NewDefaultScheduler()

	testCases := []struct {
		node        *Node
		expectError bool
		msg         string
	}{
		{
			NewNode(),
			false,
			"valid node",
		},
		{
			nil,
			true,
			"nil node",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			err := scheduler.Schedule(tc.node)

			if tc.expectError && err == nil {
				t.Error("expected error but received none")
			}

			if len(scheduler.scheduled) == 0 {
				t.Errorf("expected node to be scheduled, but slice was empty")
			}
		})
	}
}

func TestExecute(t *testing.T) {
	nodeExecutor := &FakeNodeExecutor{}
	scheduler := NewDefaultScheduler()
	scheduler.nodeExecutor = nodeExecutor
	node := NewNode()
	node.Name = "test"

	scheduler.Schedule(node)

	t.Run("returns error if task executor not found", func(t *testing.T) {
		err := scheduler.Execute(nil)
		if err == nil {
			t.Fatal("expected error to be returned")
		}
	})

	t.Run("executes scheduled node", func(t *testing.T) {
		scheduler.Execute(map[string]TaskExecutor{"test": &JournalTaskExecutor{}})
		if !nodeExecutor.Executed {
			t.Fatal("expected node to be executed, but it wasn't")
		}
	})

	t.Run("clears scheduled nodes", func(t *testing.T) {
		scheduler.Execute(map[string]TaskExecutor{"test": &JournalTaskExecutor{}})
		if len(scheduler.scheduled) > 0 {
			t.Fatalf("expect no nodes scheduled, got %d", len(scheduler.scheduled))
		}
	})
}
