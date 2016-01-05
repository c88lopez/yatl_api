package yatl

import "time"

// Todo struct
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos type
type Todos []Todo
