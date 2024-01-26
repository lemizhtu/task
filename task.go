package task

import (
	"strings"
	"time"
)

// task struct represents a task item.
type task struct {
	Name        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// Tasks represents a list of task items.
type Tasks []task

// Add creates a new task and appends it to the list.
func (ts *Tasks) Add(name string) {
	t := task{
		Name:        strings.TrimSpace(name),
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*ts = append(*ts, t)
}
