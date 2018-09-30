package models

type Assignee struct {
	ID string `json:"id"`
}

type Assignment struct {
	Assignee `json:"assignee"`
}

func (ass Assignment) AssignedTo() string {
	return ass.Assignee.ID
}
