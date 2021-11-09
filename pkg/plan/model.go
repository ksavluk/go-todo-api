package plan

import (
	"time"
)

// Plan defines the properties of a plan
type Plan struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}
