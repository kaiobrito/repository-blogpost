package data

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func (t Todo) GetID() string {
	return t.ID
}
