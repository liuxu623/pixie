// Code generated by MockGen. DO NOT EDIT.
// Source: probe.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	gomock "github.com/golang/mock/gomock"
	storepb "pixielabs.ai/pixielabs/src/vizier/services/metadata/storepb"
	reflect "reflect"
)

// MockProbeStore is a mock of ProbeStore interface
type MockProbeStore struct {
	ctrl     *gomock.Controller
	recorder *MockProbeStoreMockRecorder
}

// MockProbeStoreMockRecorder is the mock recorder for MockProbeStore
type MockProbeStoreMockRecorder struct {
	mock *MockProbeStore
}

// NewMockProbeStore creates a new mock instance
func NewMockProbeStore(ctrl *gomock.Controller) *MockProbeStore {
	mock := &MockProbeStore{ctrl: ctrl}
	mock.recorder = &MockProbeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProbeStore) EXPECT() *MockProbeStoreMockRecorder {
	return m.recorder
}

// UpsertProbe mocks base method
func (m *MockProbeStore) UpsertProbe(arg0 string, arg1 *storepb.ProbeInfo) error {
	ret := m.ctrl.Call(m, "UpsertProbe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertProbe indicates an expected call of UpsertProbe
func (mr *MockProbeStoreMockRecorder) UpsertProbe(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertProbe", reflect.TypeOf((*MockProbeStore)(nil).UpsertProbe), arg0, arg1)
}

// GetProbe mocks base method
func (m *MockProbeStore) GetProbe(arg0 string) (*storepb.ProbeInfo, error) {
	ret := m.ctrl.Call(m, "GetProbe", arg0)
	ret0, _ := ret[0].(*storepb.ProbeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProbe indicates an expected call of GetProbe
func (mr *MockProbeStoreMockRecorder) GetProbe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProbe", reflect.TypeOf((*MockProbeStore)(nil).GetProbe), arg0)
}

// GetProbes mocks base method
func (m *MockProbeStore) GetProbes() ([]*storepb.ProbeInfo, error) {
	ret := m.ctrl.Call(m, "GetProbes")
	ret0, _ := ret[0].([]*storepb.ProbeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProbes indicates an expected call of GetProbes
func (mr *MockProbeStoreMockRecorder) GetProbes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProbes", reflect.TypeOf((*MockProbeStore)(nil).GetProbes))
}

// UpdateProbeState mocks base method
func (m *MockProbeStore) UpdateProbeState(arg0 *storepb.AgentProbeStatus) error {
	ret := m.ctrl.Call(m, "UpdateProbeState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProbeState indicates an expected call of UpdateProbeState
func (mr *MockProbeStoreMockRecorder) UpdateProbeState(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProbeState", reflect.TypeOf((*MockProbeStore)(nil).UpdateProbeState), arg0)
}

// GetProbeStates mocks base method
func (m *MockProbeStore) GetProbeStates(arg0 string) ([]*storepb.AgentProbeStatus, error) {
	ret := m.ctrl.Call(m, "GetProbeStates", arg0)
	ret0, _ := ret[0].([]*storepb.AgentProbeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProbeStates indicates an expected call of GetProbeStates
func (mr *MockProbeStoreMockRecorder) GetProbeStates(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProbeStates", reflect.TypeOf((*MockProbeStore)(nil).GetProbeStates), arg0)
}
