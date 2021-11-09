package task

import (
	"time"
)

// Task defines the properties of a task
type Task struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Created     time.Time `json:"created"`
}
