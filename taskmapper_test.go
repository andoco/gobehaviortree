package gobehaviortree

import "testing"

func TestGettingTaskExecutor(t *testing.T) {
	executors := map[string]TaskExecutor{"test": JournalTaskExecutor{}}

	testCases := []struct {
		node         *Node
		expectError  bool
		taskExecutor TaskExecutor
		name         string
	}{
		{
			NewNode("test"),
			false,
			executors["test"],
			"node with name and matching executor",
		},
		{
			NewNode(""),
			true,
			nil,
			"node with no name",
		},
		{
			NewNode("missing"),
			true,
			nil,
			"node with no matching executor",
		},
	}

	mapper := DefaultTaskExecutorMapper{executors: executors}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			executor, err := mapper.Get(tc.node)

			if tc.expectError {
				if err == nil {
					t.Error("expected error but got nil")
				}
			} else if executor == nil {
				t.Error("expected executor but got nil")
			}
		})
	}
}
