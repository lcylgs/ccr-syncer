// Code generated by MockGen. DO NOT EDIT.
// Source: rpc/fe.go
//
// Generated by this command:
//
//	mockgen -source=rpc/fe.go -destination=ccr/fe_mock.go -package=ccr
//
// Package ccr is a generated GoMock package.
package ccr

import (
	reflect "reflect"

	base "github.com/selectdb/ccr_syncer/pkg/ccr/base"
	frontendservice "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/frontendservice"
	types "github.com/selectdb/ccr_syncer/pkg/rpc/kitex_gen/types"
	gomock "go.uber.org/mock/gomock"
)

// MockIFeRpc is a mock of IFeRpc interface.
type MockIFeRpc struct {
	ctrl     *gomock.Controller
	recorder *MockIFeRpcMockRecorder
}

// MockIFeRpcMockRecorder is the mock recorder for MockIFeRpc.
type MockIFeRpcMockRecorder struct {
	mock *MockIFeRpc
}

// NewMockIFeRpc creates a new mock instance.
func NewMockIFeRpc(ctrl *gomock.Controller) *MockIFeRpc {
	mock := &MockIFeRpc{ctrl: ctrl}
	mock.recorder = &MockIFeRpcMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFeRpc) EXPECT() *MockIFeRpcMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method.
func (m *MockIFeRpc) BeginTransaction(arg0 *base.Spec, arg1 string, arg2 []int64) (*frontendservice.TBeginTxnResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(*frontendservice.TBeginTxnResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction.
func (mr *MockIFeRpcMockRecorder) BeginTransaction(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockIFeRpc)(nil).BeginTransaction), arg0, arg1, arg2)
}

// CommitTransaction mocks base method.
func (m *MockIFeRpc) CommitTransaction(arg0 *base.Spec, arg1 int64, arg2 []*types.TTabletCommitInfo) (*frontendservice.TCommitTxnResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(*frontendservice.TCommitTxnResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitTransaction indicates an expected call of CommitTransaction.
func (mr *MockIFeRpcMockRecorder) CommitTransaction(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitTransaction", reflect.TypeOf((*MockIFeRpc)(nil).CommitTransaction), arg0, arg1, arg2)
}

// GetBinlog mocks base method.
func (m *MockIFeRpc) GetBinlog(arg0 *base.Spec, arg1 int64) (*frontendservice.TGetBinlogResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinlog", arg0, arg1)
	ret0, _ := ret[0].(*frontendservice.TGetBinlogResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinlog indicates an expected call of GetBinlog.
func (mr *MockIFeRpcMockRecorder) GetBinlog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinlog", reflect.TypeOf((*MockIFeRpc)(nil).GetBinlog), arg0, arg1)
}

// GetBinlogLag mocks base method.
func (m *MockIFeRpc) GetBinlogLag(arg0 *base.Spec, arg1 int64) (*frontendservice.TGetBinlogLagResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinlogLag", arg0, arg1)
	ret0, _ := ret[0].(*frontendservice.TGetBinlogLagResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinlogLag indicates an expected call of GetBinlogLag.
func (mr *MockIFeRpcMockRecorder) GetBinlogLag(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinlogLag", reflect.TypeOf((*MockIFeRpc)(nil).GetBinlogLag), arg0, arg1)
}

// GetMasterToken mocks base method.
func (m *MockIFeRpc) GetMasterToken(arg0 *base.Spec) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMasterToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterToken indicates an expected call of GetMasterToken.
func (mr *MockIFeRpcMockRecorder) GetMasterToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterToken", reflect.TypeOf((*MockIFeRpc)(nil).GetMasterToken), arg0)
}

// GetSnapshot mocks base method.
func (m *MockIFeRpc) GetSnapshot(arg0 *base.Spec, arg1 string) (*frontendservice.TGetSnapshotResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", arg0, arg1)
	ret0, _ := ret[0].(*frontendservice.TGetSnapshotResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshot indicates an expected call of GetSnapshot.
func (mr *MockIFeRpcMockRecorder) GetSnapshot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockIFeRpc)(nil).GetSnapshot), arg0, arg1)
}

// RestoreSnapshot mocks base method.
func (m *MockIFeRpc) RestoreSnapshot(arg0 *base.Spec, arg1 []*frontendservice.TTableRef, arg2 string, arg3 *frontendservice.TGetSnapshotResult_) (*frontendservice.TRestoreSnapshotResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSnapshot", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*frontendservice.TRestoreSnapshotResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreSnapshot indicates an expected call of RestoreSnapshot.
func (mr *MockIFeRpcMockRecorder) RestoreSnapshot(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSnapshot", reflect.TypeOf((*MockIFeRpc)(nil).RestoreSnapshot), arg0, arg1, arg2, arg3)
}

// RollbackTransaction mocks base method.
func (m *MockIFeRpc) RollbackTransaction(spec *base.Spec, txnId int64) (*frontendservice.TRollbackTxnResult_, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTransaction", spec, txnId)
	ret0, _ := ret[0].(*frontendservice.TRollbackTxnResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RollbackTransaction indicates an expected call of RollbackTransaction.
func (mr *MockIFeRpcMockRecorder) RollbackTransaction(spec, txnId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTransaction", reflect.TypeOf((*MockIFeRpc)(nil).RollbackTransaction), spec, txnId)
}

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// SetDb mocks base method.
func (m *MockRequest) SetDb(arg0 *string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDb", arg0)
}

// SetDb indicates an expected call of SetDb.
func (mr *MockRequestMockRecorder) SetDb(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDb", reflect.TypeOf((*MockRequest)(nil).SetDb), arg0)
}

// SetPasswd mocks base method.
func (m *MockRequest) SetPasswd(arg0 *string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPasswd", arg0)
}

// SetPasswd indicates an expected call of SetPasswd.
func (mr *MockRequestMockRecorder) SetPasswd(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPasswd", reflect.TypeOf((*MockRequest)(nil).SetPasswd), arg0)
}

// SetUser mocks base method.
func (m *MockRequest) SetUser(arg0 *string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUser", arg0)
}

// SetUser indicates an expected call of SetUser.
func (mr *MockRequestMockRecorder) SetUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUser", reflect.TypeOf((*MockRequest)(nil).SetUser), arg0)
}