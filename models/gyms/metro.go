package gyms

type Metro struct {
	ID   int
	Name string
}

func (m *Metro) TableName() string {
	return "metros"
}