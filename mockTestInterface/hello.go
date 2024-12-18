package mockTestInterface

type Result struct {
	Result int
}

type TestModule interface {
	Add(a int, b int)
}

func (res *Result) Add(a, b int) {
	res.Result = a + b
}
