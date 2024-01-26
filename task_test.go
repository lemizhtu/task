package task_test

import (
	"testing"

	"github.com/lemizhtu/task"
	"github.com/stretchr/testify/assert"
)

func TestTasks_Add(t *testing.T) {
	var (
		tasks = task.Tasks{}
	)

	testCases := []struct {
		name     string
		taskName string
		index    int
		length   int
	}{
		{"Empty", "", 0, 1},
		{"Normal", "Do homework", 1, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tasks.Add(tc.taskName)

			assert.Len(t, tasks, tc.length)
			assert.Equal(t, tc.taskName, tasks[tc.index].Name)
			assert.False(t, tasks[tc.index].Done)
			assert.Nil(t, tasks[tc.index].CompletedAt)
		})
	}
}

func TestTasks_Complete(t *testing.T) {

}
