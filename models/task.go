package models

type Task struct {
	ID          int             `json:"ID"`
	Description string          `json:"Description"`
	Status      StatusComponent `json:"Status"`
	CreatedAt   string          `json:"Created-At"`
	UpdatedAt   string          `json:"Updated-At"`
}

type StatusComponent struct {
	ToDo       bool `json:"To Do"`
	InProgress bool `json:"In Progress"`
	Done       bool `json:"Done"`
}
