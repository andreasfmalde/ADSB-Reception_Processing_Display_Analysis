// Code generated by MockGen. DO NOT EDIT.
// Source: .\backend\internal\service\sbsService\sbsService.go

// Package mockData is a generated GoMock package.
package mock

import (
	models "adsb-api/internal/global/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSbsService is a mockData of SbsService interface.
type MockSbsService struct {
	ctrl     *gomock.Controller
	recorder *MockSbsServiceMockRecorder
}

// MockSbsServiceMockRecorder is the mockData recorder for MockSbsService.
type MockSbsServiceMockRecorder struct {
	mock *MockSbsService
}

// NewMockSbsService creates a new mockData instance.
func NewMockSbsService(ctrl *gomock.Controller) *MockSbsService {
	mock := &MockSbsService{ctrl: ctrl}
	mock.recorder = &MockSbsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSbsService) EXPECT() *MockSbsServiceMockRecorder {
	return m.recorder
}

// CreateAdsbTables mocks base method.
func (m *MockSbsService) CreateAdsbTables() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdsbTables")
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdsbTables indicates an expected call of CreateAdsbTables.
func (mr *MockSbsServiceMockRecorder) CreateAdsbTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdsbTables", reflect.TypeOf((*MockSbsService)(nil).CreateAdsbTables))
}

// InitAndStartCleanUpJob mocks base method.
func (m *MockSbsService) InitAndStartCleanUpJob() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleCleanUpJob")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitAndStartCleanUpJob indicates an expected call of InitAndStartCleanUpJob.
func (mr *MockSbsServiceMockRecorder) InitAndStartCleanUpJob() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleCleanUpJob", reflect.TypeOf((*MockSbsService)(nil).InitAndStartCleanUpJob))
}

// InsertNewSbsData mocks base method.
func (m *MockSbsService) InsertNewSbsData(aircraft []models.AircraftCurrentModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNewSbsData", aircraft)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNewSbsData indicates an expected call of InsertNewSbsData.
func (mr *MockSbsServiceMockRecorder) InsertNewSbsData(aircraft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNewSbsData", reflect.TypeOf((*MockSbsService)(nil).InsertNewSbsData), aircraft)
}

// StartScheduler mocks base method.
func (m *MockSbsService) StartScheduler() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartScheduler")
	ret0, _ := ret[0].(error)
	return ret0
}

// StartScheduler indicates an expected call of StartScheduler.
func (mr *MockSbsServiceMockRecorder) StartScheduler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartScheduler", reflect.TypeOf((*MockSbsService)(nil).StartScheduler))
}
