package models

type Task struct {
	ID          int `json:"id"`
	Description string
	Done        bool
}
