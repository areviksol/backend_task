package model

// Model struct represents the data model
type Model struct {
	data string
}

// NewModel creates a new instance of Model
func NewModel(initialData string) *Model {
	return &Model{data: initialData}
}

// SetData updates the model's data
func (m *Model) SetData(newData string) {
	m.data = newData
}

// GetData returns the model's data
func (m *Model) GetData() string {
	return m.data
}
