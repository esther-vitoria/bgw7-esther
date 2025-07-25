package simulator

// MockCatchSimulator é um mock para a interface CatchSimulator.
type MockCatchSimulator struct {
	// Valor a ser retornado por CanCatch
	ReturnValue bool

	// Indica se CanCatch foi chamado
	Called bool

	// Opcional: guarda os argumentos da última chamada
	HunterArg *Subject
	PreyArg   *Subject
}

func (m *MockCatchSimulator) CanCatch(hunter, prey *Subject) bool {
	m.Called = true
	m.HunterArg = hunter
	m.PreyArg = prey
	return m.ReturnValue
}
