package gobehaviortree

// TaskStatus is the current execution status of the task node.
type TaskStatus int

// TaskResult is the result returned from a task when executed.
type TaskResult int

const (
	// None indicates the task has not been initialized yet.
	None TaskStatus = iota

	// Ready indicates the task has been initialized and is ready to be executed.
	Ready

	// Running indicates that the task is currently running.
	Running
)

const (
	// Success indicates the task completed successfully.
	Success TaskResult = iota

	// Failure indicates the task did not complete successfully.
	Failure

	// Pending indicates the task is still running.
	Pending
)
