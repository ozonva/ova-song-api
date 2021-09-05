// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-song-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozonva/ova-song-api/internal/models"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddSong mocks base method.
func (m *MockRepo) AddSong(arg0 context.Context, arg1 models.Song) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSong", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSong indicates an expected call of AddSong.
func (mr *MockRepoMockRecorder) AddSong(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSong", reflect.TypeOf((*MockRepo)(nil).AddSong), arg0, arg1)
}

// AddSongs mocks base method.
func (m *MockRepo) AddSongs(arg0 context.Context, arg1 []models.Song) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSongs", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSongs indicates an expected call of AddSongs.
func (mr *MockRepoMockRecorder) AddSongs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSongs", reflect.TypeOf((*MockRepo)(nil).AddSongs), arg0, arg1)
}

// DescribeSong mocks base method.
func (m *MockRepo) DescribeSong(arg0 context.Context, arg1 uint64) (*models.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSong", arg0, arg1)
	ret0, _ := ret[0].(*models.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSong indicates an expected call of DescribeSong.
func (mr *MockRepoMockRecorder) DescribeSong(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSong", reflect.TypeOf((*MockRepo)(nil).DescribeSong), arg0, arg1)
}

// ListSongs mocks base method.
func (m *MockRepo) ListSongs(arg0 context.Context, arg1, arg2 uint64) ([]models.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSongs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSongs indicates an expected call of ListSongs.
func (mr *MockRepoMockRecorder) ListSongs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSongs", reflect.TypeOf((*MockRepo)(nil).ListSongs), arg0, arg1, arg2)
}

// RemoveSong mocks base method.
func (m *MockRepo) RemoveSong(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSong", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSong indicates an expected call of RemoveSong.
func (mr *MockRepoMockRecorder) RemoveSong(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSong", reflect.TypeOf((*MockRepo)(nil).RemoveSong), arg0, arg1)
}

// UpdateSong mocks base method.
func (m *MockRepo) UpdateSong(arg0 context.Context, arg1 models.Song) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSong", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSong indicates an expected call of UpdateSong.
func (mr *MockRepoMockRecorder) UpdateSong(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSong", reflect.TypeOf((*MockRepo)(nil).UpdateSong), arg0, arg1)
}
