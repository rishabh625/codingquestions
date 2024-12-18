package mockTestInterface

type TestModuleMock struct {
	Result int
}

func (s *TestModuleMock) Add(a, b int) {
	s.Result = 30
}
