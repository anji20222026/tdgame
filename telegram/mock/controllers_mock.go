// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aoyako/telegram_2ch_res_bot/controller (interfaces: User,Subscription,Info)

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	logic "github.com/aoyako/telegram_2ch_res_bot/logic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetUsersByPublication mocks base method
func (m *MockUser) GetUsersByPublication(arg0 *logic.Publication) ([]logic.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByPublication", arg0)
	ret0, _ := ret[0].([]logic.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByPublication indicates an expected call of GetUsersByPublication
func (mr *MockUserMockRecorder) GetUsersByPublication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByPublication", reflect.TypeOf((*MockUser)(nil).GetUsersByPublication), arg0)
}

// Register mocks base method
func (m *MockUser) Register(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockUserMockRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUser)(nil).Register), arg0)
}

// Unregister mocks base method
func (m *MockUser) Unregister(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister
func (mr *MockUserMockRecorder) Unregister(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockUser)(nil).Unregister), arg0)
}

// MockSubscription is a mock of Subscription interface
type MockSubscription struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionMockRecorder
}

// MockSubscriptionMockRecorder is the mock recorder for MockSubscription
type MockSubscriptionMockRecorder struct {
	mock *MockSubscription
}

// NewMockSubscription creates a new mock instance
func NewMockSubscription(ctrl *gomock.Controller) *MockSubscription {
	mock := &MockSubscription{ctrl: ctrl}
	mock.recorder = &MockSubscriptionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubscription) EXPECT() *MockSubscriptionMockRecorder {
	return m.recorder
}

// AddNew mocks base method
func (m *MockSubscription) AddNew(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNew", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNew indicates an expected call of AddNew
func (mr *MockSubscriptionMockRecorder) AddNew(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNew", reflect.TypeOf((*MockSubscription)(nil).AddNew), arg0, arg1)
}

// Create mocks base method
func (m *MockSubscription) Create(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockSubscriptionMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSubscription)(nil).Create), arg0, arg1)
}

// GetAllDefaultSubs mocks base method
func (m *MockSubscription) GetAllDefaultSubs() []logic.Publication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDefaultSubs")
	ret0, _ := ret[0].([]logic.Publication)
	return ret0
}

// GetAllDefaultSubs indicates an expected call of GetAllDefaultSubs
func (mr *MockSubscriptionMockRecorder) GetAllDefaultSubs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDefaultSubs", reflect.TypeOf((*MockSubscription)(nil).GetAllDefaultSubs))
}

// GetAllSubs mocks base method
func (m *MockSubscription) GetAllSubs() []logic.Publication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubs")
	ret0, _ := ret[0].([]logic.Publication)
	return ret0
}

// GetAllSubs indicates an expected call of GetAllSubs
func (mr *MockSubscriptionMockRecorder) GetAllSubs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubs", reflect.TypeOf((*MockSubscription)(nil).GetAllSubs))
}

// GetSubsByChatID mocks base method
func (m *MockSubscription) GetSubsByChatID(arg0 int64) ([]logic.Publication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubsByChatID", arg0)
	ret0, _ := ret[0].([]logic.Publication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubsByChatID indicates an expected call of GetSubsByChatID
func (mr *MockSubscriptionMockRecorder) GetSubsByChatID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubsByChatID", reflect.TypeOf((*MockSubscription)(nil).GetSubsByChatID), arg0)
}

// Remove mocks base method
func (m *MockSubscription) Remove(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockSubscriptionMockRecorder) Remove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockSubscription)(nil).Remove), arg0, arg1)
}

// RemoveDefault mocks base method
func (m *MockSubscription) RemoveDefault(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDefault", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDefault indicates an expected call of RemoveDefault
func (mr *MockSubscriptionMockRecorder) RemoveDefault(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDefault", reflect.TypeOf((*MockSubscription)(nil).RemoveDefault), arg0, arg1)
}

// Subscribe mocks base method
func (m *MockSubscription) Subscribe(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockSubscriptionMockRecorder) Subscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscription)(nil).Subscribe), arg0, arg1)
}

// Update mocks base method
func (m *MockSubscription) Update(arg0 int64, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockSubscriptionMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSubscription)(nil).Update), arg0, arg1)
}

// MockInfo is a mock of Info interface
type MockInfo struct {
	ctrl     *gomock.Controller
	recorder *MockInfoMockRecorder
}

// MockInfoMockRecorder is the mock recorder for MockInfo
type MockInfoMockRecorder struct {
	mock *MockInfo
}

// NewMockInfo creates a new mock instance
func NewMockInfo(ctrl *gomock.Controller) *MockInfo {
	mock := &MockInfo{ctrl: ctrl}
	mock.recorder = &MockInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfo) EXPECT() *MockInfoMockRecorder {
	return m.recorder
}

// GetLastTimestamp mocks base method
func (m *MockInfo) GetLastTimestamp() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastTimestamp")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLastTimestamp indicates an expected call of GetLastTimestamp
func (mr *MockInfoMockRecorder) GetLastTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastTimestamp", reflect.TypeOf((*MockInfo)(nil).GetLastTimestamp))
}

// SetLastTimestamp mocks base method
func (m *MockInfo) SetLastTimestamp(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastTimestamp", arg0)
}

// SetLastTimestamp indicates an expected call of SetLastTimestamp
func (mr *MockInfoMockRecorder) SetLastTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastTimestamp", reflect.TypeOf((*MockInfo)(nil).SetLastTimestamp), arg0)
}
