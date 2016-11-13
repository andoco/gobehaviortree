package gobehaviortree

import "testing"

func TestNodeInitOnFirstRun(t *testing.T) {
	node := NewNode("test")
	nodeExecutor := DefaultNodeExecutor{}
	taskExecutor := NewJournalTaskExecutor(Success)

	nodeExecutor.Execute(node, &taskExecutor)

	if taskExecutor.Counts.InitCount != 1 {
		t.Fatalf("expected init count of %d, got %d", 1, taskExecutor.Counts.InitCount)
	}
}

func TestChangingNodeStatus(t *testing.T) {
	testCases := []struct {
		name           string
		existingStatus TaskStatus
		result         TaskResult
		expectedStatus TaskStatus
	}{
		{"success result", Ready, Success, Ready},
		{"failure result", Ready, Failure, Ready},
		{"pending result", Ready, Pending, Running},
	}

	nodeExecutor := DefaultNodeExecutor{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node := NewNode("test")
			node.State.Status = tc.existingStatus
			journalExecutor := NewJournalTaskExecutor(tc.result)

			nodeExecutor.Execute(node, journalExecutor)

			if node.State.Status != tc.expectedStatus {
				t.Fatalf("expect status %v, got %v", tc.expectedStatus, node.State.Status)
			}
		})
	}
}
