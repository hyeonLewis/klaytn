// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/klaytn/klaytn/consensus/istanbul (interfaces: Backend)

// Package mock_istanbul is a generated GoMock package.
package mock_istanbul

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/klaytn/klaytn/common"
	istanbul "github.com/klaytn/klaytn/consensus/istanbul"
	event "github.com/klaytn/klaytn/event"
	big "math/big"
	reflect "reflect"
	time "time"
)

// MockBackend is a mock of Backend interface
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockBackend) Address() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockBackendMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockBackend)(nil).Address))
}

// Broadcast mocks base method
func (m *MockBackend) Broadcast(arg0 common.Hash, arg1 istanbul.ValidatorSet, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockBackendMockRecorder) Broadcast(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockBackend)(nil).Broadcast), arg0, arg1, arg2)
}

// CheckSignature mocks base method
func (m *MockBackend) CheckSignature(arg0 []byte, arg1 common.Address, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSignature", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSignature indicates an expected call of CheckSignature
func (mr *MockBackendMockRecorder) CheckSignature(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSignature", reflect.TypeOf((*MockBackend)(nil).CheckSignature), arg0, arg1, arg2)
}

// Commit mocks base method
func (m *MockBackend) Commit(arg0 istanbul.Proposal, arg1 [][]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockBackendMockRecorder) Commit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockBackend)(nil).Commit), arg0, arg1)
}

// EventMux mocks base method
func (m *MockBackend) EventMux() *event.TypeMux {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventMux")
	ret0, _ := ret[0].(*event.TypeMux)
	return ret0
}

// EventMux indicates an expected call of EventMux
func (mr *MockBackendMockRecorder) EventMux() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventMux", reflect.TypeOf((*MockBackend)(nil).EventMux))
}

// GetProposer mocks base method
func (m *MockBackend) GetProposer(arg0 uint64) common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposer", arg0)
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// GetProposer indicates an expected call of GetProposer
func (mr *MockBackendMockRecorder) GetProposer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposer", reflect.TypeOf((*MockBackend)(nil).GetProposer), arg0)
}

// GetRewardBase mocks base method
func (m *MockBackend) GetRewardBase() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardBase")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// GetRewardBase indicates an expected call of GetRewardBase
func (mr *MockBackendMockRecorder) GetRewardBase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardBase", reflect.TypeOf((*MockBackend)(nil).GetRewardBase))
}

// GetSubGroupSize mocks base method
func (m *MockBackend) GetSubGroupSize() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubGroupSize")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetSubGroupSize indicates an expected call of GetSubGroupSize
func (mr *MockBackendMockRecorder) GetSubGroupSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubGroupSize", reflect.TypeOf((*MockBackend)(nil).GetSubGroupSize))
}

// Gossip mocks base method
func (m *MockBackend) Gossip(arg0 istanbul.ValidatorSet, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gossip", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gossip indicates an expected call of Gossip
func (mr *MockBackendMockRecorder) Gossip(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gossip", reflect.TypeOf((*MockBackend)(nil).Gossip), arg0, arg1)
}

// GossipSubPeer mocks base method
func (m *MockBackend) GossipSubPeer(arg0 common.Hash, arg1 istanbul.ValidatorSet, arg2 []byte) map[common.Address]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GossipSubPeer", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[common.Address]bool)
	return ret0
}

// GossipSubPeer indicates an expected call of GossipSubPeer
func (mr *MockBackendMockRecorder) GossipSubPeer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GossipSubPeer", reflect.TypeOf((*MockBackend)(nil).GossipSubPeer), arg0, arg1, arg2)
}

// HasBadProposal mocks base method
func (m *MockBackend) HasBadProposal(arg0 common.Hash) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBadProposal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBadProposal indicates an expected call of HasBadProposal
func (mr *MockBackendMockRecorder) HasBadProposal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBadProposal", reflect.TypeOf((*MockBackend)(nil).HasBadProposal), arg0)
}

// HasPropsal mocks base method
func (m *MockBackend) HasPropsal(arg0 common.Hash, arg1 *big.Int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPropsal", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPropsal indicates an expected call of HasPropsal
func (mr *MockBackendMockRecorder) HasPropsal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPropsal", reflect.TypeOf((*MockBackend)(nil).HasPropsal), arg0, arg1)
}

// LastProposal mocks base method
func (m *MockBackend) LastProposal() (istanbul.Proposal, common.Address) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastProposal")
	ret0, _ := ret[0].(istanbul.Proposal)
	ret1, _ := ret[1].(common.Address)
	return ret0, ret1
}

// LastProposal indicates an expected call of LastProposal
func (mr *MockBackendMockRecorder) LastProposal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastProposal", reflect.TypeOf((*MockBackend)(nil).LastProposal))
}

// NodeType mocks base method
func (m *MockBackend) NodeType() common.ConnType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeType")
	ret0, _ := ret[0].(common.ConnType)
	return ret0
}

// NodeType indicates an expected call of NodeType
func (mr *MockBackendMockRecorder) NodeType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeType", reflect.TypeOf((*MockBackend)(nil).NodeType))
}

// ParentValidators mocks base method
func (m *MockBackend) ParentValidators(arg0 istanbul.Proposal) istanbul.ValidatorSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentValidators", arg0)
	ret0, _ := ret[0].(istanbul.ValidatorSet)
	return ret0
}

// ParentValidators indicates an expected call of ParentValidators
func (mr *MockBackendMockRecorder) ParentValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentValidators", reflect.TypeOf((*MockBackend)(nil).ParentValidators), arg0)
}

// SetCurrentView mocks base method
func (m *MockBackend) SetCurrentView(arg0 *istanbul.View) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentView", arg0)
}

// SetCurrentView indicates an expected call of SetCurrentView
func (mr *MockBackendMockRecorder) SetCurrentView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentView", reflect.TypeOf((*MockBackend)(nil).SetCurrentView), arg0)
}

// Sign mocks base method
func (m *MockBackend) Sign(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockBackendMockRecorder) Sign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockBackend)(nil).Sign), arg0)
}

// Validators mocks base method
func (m *MockBackend) Validators(arg0 istanbul.Proposal) istanbul.ValidatorSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", arg0)
	ret0, _ := ret[0].(istanbul.ValidatorSet)
	return ret0
}

// Validators indicates an expected call of Validators
func (mr *MockBackendMockRecorder) Validators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockBackend)(nil).Validators), arg0)
}

// Verify mocks base method
func (m *MockBackend) Verify(arg0 istanbul.Proposal) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify
func (mr *MockBackendMockRecorder) Verify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockBackend)(nil).Verify), arg0)
}
