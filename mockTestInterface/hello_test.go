package mockTestInterface_test

import (
	"github.com/stretchr/testify/suite"
	"mockTestInterface"
	"testing"
)

type ModuleTestSuiteStruct struct {
	suite.Suite
}

func ModuleTestSuite(t *testing.T) *ModuleTestSuiteStruct {
	return &ModuleTestSuiteStruct{}
}

func TestMouduleTestSuite(t *testing.T) {
	suite.Run(t, ModuleTestSuite(t))
}

func (s *ModuleTestSuiteStruct) TestAdd() {
	result := &mockTestInterface.Result{}
	result.Add(10, 20)
	s.Equal(30, result.Result)
}

func (s *ModuleTestSuiteStruct) TestAddViaMock() {
	result := &mockTestInterface.TestModuleMock{}
	result.Add(10, 20)
	s.Equal(30, result.Result)
}
