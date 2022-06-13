// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package appresource_test is a generated GoMock package.
package appresource_test

import (
	context "context"
	reflect "reflect"

	resource "github.com/authgear/authgear-server/pkg/util/resource"
	gomock "github.com/golang/mock/gomock"
)

// MockTutorialService is a mock of TutorialService interface.
type MockTutorialService struct {
	ctrl     *gomock.Controller
	recorder *MockTutorialServiceMockRecorder
}

// MockTutorialServiceMockRecorder is the mock recorder for MockTutorialService.
type MockTutorialServiceMockRecorder struct {
	mock *MockTutorialService
}

// NewMockTutorialService creates a new mock instance.
func NewMockTutorialService(ctrl *gomock.Controller) *MockTutorialService {
	mock := &MockTutorialService{ctrl: ctrl}
	mock.recorder = &MockTutorialServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTutorialService) EXPECT() *MockTutorialServiceMockRecorder {
	return m.recorder
}

// OnUpdateResource mocks base method.
func (m *MockTutorialService) OnUpdateResource(ctx context.Context, appID string, resourcesInAllFss []resource.ResourceFile, resourceInTargetFs *resource.ResourceFile, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnUpdateResource", ctx, appID, resourcesInAllFss, resourceInTargetFs, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnUpdateResource indicates an expected call of OnUpdateResource.
func (mr *MockTutorialServiceMockRecorder) OnUpdateResource(ctx, appID, resourcesInAllFss, resourceInTargetFs, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnUpdateResource", reflect.TypeOf((*MockTutorialService)(nil).OnUpdateResource), ctx, appID, resourcesInAllFss, resourceInTargetFs, data)
}