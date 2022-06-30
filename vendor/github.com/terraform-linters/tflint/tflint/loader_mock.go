// Code generated by MockGen. DO NOT EDIT.
// Source: loader.go

// Package tflint is a generated GoMock package.
package tflint

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v2 "github.com/hashicorp/hcl/v2"
	configs "github.com/terraform-linters/tflint/terraform/configs"
	terraform "github.com/terraform-linters/tflint/terraform/terraform"
)

// MockAbstractLoader is a mock of AbstractLoader interface.
type MockAbstractLoader struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractLoaderMockRecorder
}

// MockAbstractLoaderMockRecorder is the mock recorder for MockAbstractLoader.
type MockAbstractLoaderMockRecorder struct {
	mock *MockAbstractLoader
}

// NewMockAbstractLoader creates a new mock instance.
func NewMockAbstractLoader(ctrl *gomock.Controller) *MockAbstractLoader {
	mock := &MockAbstractLoader{ctrl: ctrl}
	mock.recorder = &MockAbstractLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAbstractLoader) EXPECT() *MockAbstractLoaderMockRecorder {
	return m.recorder
}

// Files mocks base method.
func (m *MockAbstractLoader) Files() (map[string]*v2.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Files")
	ret0, _ := ret[0].(map[string]*v2.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Files indicates an expected call of Files.
func (mr *MockAbstractLoaderMockRecorder) Files() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Files", reflect.TypeOf((*MockAbstractLoader)(nil).Files))
}

// LoadAnnotations mocks base method.
func (m *MockAbstractLoader) LoadAnnotations(arg0 string) (map[string]Annotations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAnnotations", arg0)
	ret0, _ := ret[0].(map[string]Annotations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAnnotations indicates an expected call of LoadAnnotations.
func (mr *MockAbstractLoaderMockRecorder) LoadAnnotations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAnnotations", reflect.TypeOf((*MockAbstractLoader)(nil).LoadAnnotations), arg0)
}

// LoadConfig mocks base method.
func (m *MockAbstractLoader) LoadConfig(arg0 string) (*configs.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadConfig", arg0)
	ret0, _ := ret[0].(*configs.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadConfig indicates an expected call of LoadConfig.
func (mr *MockAbstractLoaderMockRecorder) LoadConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadConfig", reflect.TypeOf((*MockAbstractLoader)(nil).LoadConfig), arg0)
}

// LoadValuesFiles mocks base method.
func (m *MockAbstractLoader) LoadValuesFiles(arg0 ...string) ([]terraform.InputValues, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LoadValuesFiles", varargs...)
	ret0, _ := ret[0].([]terraform.InputValues)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadValuesFiles indicates an expected call of LoadValuesFiles.
func (mr *MockAbstractLoaderMockRecorder) LoadValuesFiles(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadValuesFiles", reflect.TypeOf((*MockAbstractLoader)(nil).LoadValuesFiles), arg0...)
}

// Sources mocks base method.
func (m *MockAbstractLoader) Sources() map[string][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sources")
	ret0, _ := ret[0].(map[string][]byte)
	return ret0
}

// Sources indicates an expected call of Sources.
func (mr *MockAbstractLoaderMockRecorder) Sources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sources", reflect.TypeOf((*MockAbstractLoader)(nil).Sources))
}
