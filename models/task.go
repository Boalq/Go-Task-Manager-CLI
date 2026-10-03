package models

type Task struct {
	ID         int    `json:"ID"`
	Todo       string `json:"ToDo"`
	InProgress bool   `json:"InProgress"`
	Done       bool   `json:"Done"`
}
