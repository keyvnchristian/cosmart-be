// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go

// Package mocks_openlibrary is a generated GoMock package.
package mocks_openlibrary

import (
	context "context"
	openlibrary "cosmart/routing/openlibrary"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLibraryService is a mock of LibraryService interface.
type MockLibraryService struct {
	ctrl     *gomock.Controller
	recorder *MockLibraryServiceMockRecorder
}

// MockLibraryServiceMockRecorder is the mock recorder for MockLibraryService.
type MockLibraryServiceMockRecorder struct {
	mock *MockLibraryService
}

// NewMockLibraryService creates a new mock instance.
func NewMockLibraryService(ctrl *gomock.Controller) *MockLibraryService {
	mock := &MockLibraryService{ctrl: ctrl}
	mock.recorder = &MockLibraryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLibraryService) EXPECT() *MockLibraryServiceMockRecorder {
	return m.recorder
}

// FetchBook mocks base method.
func (m *MockLibraryService) FetchBook(ctx context.Context, req openlibrary.LibraryRequest) (openlibrary.LibraryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBook", ctx, req)
	ret0, _ := ret[0].(openlibrary.LibraryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBook indicates an expected call of FetchBook.
func (mr *MockLibraryServiceMockRecorder) FetchBook(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBook", reflect.TypeOf((*MockLibraryService)(nil).FetchBook), ctx, req)
}
