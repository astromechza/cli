// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package client is a generated GoMock package.
package client

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	functions "github.com/oracle/oci-go-sdk/v28/functions"
	reflect "reflect"
)

// MockFunctionsManagementClient is a mock of FunctionsManagementClient interface
type MockFunctionsManagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockFunctionsManagementClientMockRecorder
}

// MockFunctionsManagementClientMockRecorder is the mock recorder for MockFunctionsManagementClient
type MockFunctionsManagementClientMockRecorder struct {
	mock *MockFunctionsManagementClient
}

// NewMockFunctionsManagementClient creates a new mock instance
func NewMockFunctionsManagementClient(ctrl *gomock.Controller) *MockFunctionsManagementClient {
	mock := &MockFunctionsManagementClient{ctrl: ctrl}
	mock.recorder = &MockFunctionsManagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFunctionsManagementClient) EXPECT() *MockFunctionsManagementClientMockRecorder {
	return m.recorder
}

// CreateApplication mocks base method
func (m *MockFunctionsManagementClient) CreateApplication(ctx context.Context, request functions.CreateApplicationRequest) (functions.CreateApplicationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", ctx, request)
	ret0, _ := ret[0].(functions.CreateApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication
func (mr *MockFunctionsManagementClientMockRecorder) CreateApplication(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockFunctionsManagementClient)(nil).CreateApplication), ctx, request)
}

// CreateFunction mocks base method
func (m *MockFunctionsManagementClient) CreateFunction(ctx context.Context, request functions.CreateFunctionRequest) (functions.CreateFunctionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFunction", ctx, request)
	ret0, _ := ret[0].(functions.CreateFunctionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFunction indicates an expected call of CreateFunction
func (mr *MockFunctionsManagementClientMockRecorder) CreateFunction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFunction", reflect.TypeOf((*MockFunctionsManagementClient)(nil).CreateFunction), ctx, request)
}

// DeleteApplication mocks base method
func (m *MockFunctionsManagementClient) DeleteApplication(ctx context.Context, request functions.DeleteApplicationRequest) (functions.DeleteApplicationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteApplication", ctx, request)
	ret0, _ := ret[0].(functions.DeleteApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApplication indicates an expected call of DeleteApplication
func (mr *MockFunctionsManagementClientMockRecorder) DeleteApplication(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockFunctionsManagementClient)(nil).DeleteApplication), ctx, request)
}

// DeleteFunction mocks base method
func (m *MockFunctionsManagementClient) DeleteFunction(ctx context.Context, request functions.DeleteFunctionRequest) (functions.DeleteFunctionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFunction", ctx, request)
	ret0, _ := ret[0].(functions.DeleteFunctionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFunction indicates an expected call of DeleteFunction
func (mr *MockFunctionsManagementClientMockRecorder) DeleteFunction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFunction", reflect.TypeOf((*MockFunctionsManagementClient)(nil).DeleteFunction), ctx, request)
}

// GetApplication mocks base method
func (m *MockFunctionsManagementClient) GetApplication(ctx context.Context, request functions.GetApplicationRequest) (functions.GetApplicationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApplication", ctx, request)
	ret0, _ := ret[0].(functions.GetApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication
func (mr *MockFunctionsManagementClientMockRecorder) GetApplication(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockFunctionsManagementClient)(nil).GetApplication), ctx, request)
}

// GetFunction mocks base method
func (m *MockFunctionsManagementClient) GetFunction(ctx context.Context, request functions.GetFunctionRequest) (functions.GetFunctionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunction", ctx, request)
	ret0, _ := ret[0].(functions.GetFunctionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunction indicates an expected call of GetFunction
func (mr *MockFunctionsManagementClientMockRecorder) GetFunction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockFunctionsManagementClient)(nil).GetFunction), ctx, request)
}

// ListApplications mocks base method
func (m *MockFunctionsManagementClient) ListApplications(ctx context.Context, request functions.ListApplicationsRequest) (functions.ListApplicationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", ctx, request)
	ret0, _ := ret[0].(functions.ListApplicationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications
func (mr *MockFunctionsManagementClientMockRecorder) ListApplications(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockFunctionsManagementClient)(nil).ListApplications), ctx, request)
}

// ListFunctions mocks base method
func (m *MockFunctionsManagementClient) ListFunctions(ctx context.Context, request functions.ListFunctionsRequest) (functions.ListFunctionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFunctions", ctx, request)
	ret0, _ := ret[0].(functions.ListFunctionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctions indicates an expected call of ListFunctions
func (mr *MockFunctionsManagementClientMockRecorder) ListFunctions(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockFunctionsManagementClient)(nil).ListFunctions), ctx, request)
}

// UpdateApplication mocks base method
func (m *MockFunctionsManagementClient) UpdateApplication(ctx context.Context, request functions.UpdateApplicationRequest) (functions.UpdateApplicationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", ctx, request)
	ret0, _ := ret[0].(functions.UpdateApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApplication indicates an expected call of UpdateApplication
func (mr *MockFunctionsManagementClientMockRecorder) UpdateApplication(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockFunctionsManagementClient)(nil).UpdateApplication), ctx, request)
}

// UpdateFunction mocks base method
func (m *MockFunctionsManagementClient) UpdateFunction(ctx context.Context, request functions.UpdateFunctionRequest) (functions.UpdateFunctionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFunction", ctx, request)
	ret0, _ := ret[0].(functions.UpdateFunctionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFunction indicates an expected call of UpdateFunction
func (mr *MockFunctionsManagementClientMockRecorder) UpdateFunction(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFunction", reflect.TypeOf((*MockFunctionsManagementClient)(nil).UpdateFunction), ctx, request)
}