// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pivotalservices/cf-mgmt/cloudcontroller (interfaces: Manager)

// Package mock_cloudcontroller is a generated GoMock package.
package mock_cloudcontroller

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cloudcontroller "github.com/pivotalservices/cf-mgmt/cloudcontroller"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// AddUserToOrg mocks base method
func (m *MockManager) AddUserToOrg(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AddUserToOrg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToOrg indicates an expected call of AddUserToOrg
func (mr *MockManagerMockRecorder) AddUserToOrg(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToOrg", reflect.TypeOf((*MockManager)(nil).AddUserToOrg), arg0, arg1)
}

// AddUserToOrgRole mocks base method
func (m *MockManager) AddUserToOrgRole(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "AddUserToOrgRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToOrgRole indicates an expected call of AddUserToOrgRole
func (mr *MockManagerMockRecorder) AddUserToOrgRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToOrgRole", reflect.TypeOf((*MockManager)(nil).AddUserToOrgRole), arg0, arg1, arg2)
}

// AddUserToSpaceRole mocks base method
func (m *MockManager) AddUserToSpaceRole(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "AddUserToSpaceRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToSpaceRole indicates an expected call of AddUserToSpaceRole
func (mr *MockManagerMockRecorder) AddUserToSpaceRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToSpaceRole", reflect.TypeOf((*MockManager)(nil).AddUserToSpaceRole), arg0, arg1, arg2)
}

// AssignQuotaToOrg mocks base method
func (m *MockManager) AssignQuotaToOrg(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AssignQuotaToOrg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignQuotaToOrg indicates an expected call of AssignQuotaToOrg
func (mr *MockManagerMockRecorder) AssignQuotaToOrg(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignQuotaToOrg", reflect.TypeOf((*MockManager)(nil).AssignQuotaToOrg), arg0, arg1)
}

// AssignQuotaToSpace mocks base method
func (m *MockManager) AssignQuotaToSpace(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AssignQuotaToSpace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignQuotaToSpace indicates an expected call of AssignQuotaToSpace
func (mr *MockManagerMockRecorder) AssignQuotaToSpace(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignQuotaToSpace", reflect.TypeOf((*MockManager)(nil).AssignQuotaToSpace), arg0, arg1)
}

// AssignSecurityGroupToSpace mocks base method
func (m *MockManager) AssignSecurityGroupToSpace(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "AssignSecurityGroupToSpace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AssignSecurityGroupToSpace indicates an expected call of AssignSecurityGroupToSpace
func (mr *MockManagerMockRecorder) AssignSecurityGroupToSpace(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignSecurityGroupToSpace", reflect.TypeOf((*MockManager)(nil).AssignSecurityGroupToSpace), arg0, arg1)
}

// CreateOrg mocks base method
func (m *MockManager) CreateOrg(arg0 string) error {
	ret := m.ctrl.Call(m, "CreateOrg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrg indicates an expected call of CreateOrg
func (mr *MockManagerMockRecorder) CreateOrg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrg", reflect.TypeOf((*MockManager)(nil).CreateOrg), arg0)
}

// CreatePrivateDomain mocks base method
func (m *MockManager) CreatePrivateDomain(arg0, arg1 string) (string, error) {
	ret := m.ctrl.Call(m, "CreatePrivateDomain", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePrivateDomain indicates an expected call of CreatePrivateDomain
func (mr *MockManagerMockRecorder) CreatePrivateDomain(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrivateDomain", reflect.TypeOf((*MockManager)(nil).CreatePrivateDomain), arg0, arg1)
}

// CreateQuota mocks base method
func (m *MockManager) CreateQuota(arg0 cloudcontroller.QuotaEntity) (string, error) {
	ret := m.ctrl.Call(m, "CreateQuota", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuota indicates an expected call of CreateQuota
func (mr *MockManagerMockRecorder) CreateQuota(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuota", reflect.TypeOf((*MockManager)(nil).CreateQuota), arg0)
}

// CreateSecurityGroup mocks base method
func (m *MockManager) CreateSecurityGroup(arg0, arg1 string) (string, error) {
	ret := m.ctrl.Call(m, "CreateSecurityGroup", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecurityGroup indicates an expected call of CreateSecurityGroup
func (mr *MockManagerMockRecorder) CreateSecurityGroup(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecurityGroup", reflect.TypeOf((*MockManager)(nil).CreateSecurityGroup), arg0, arg1)
}

// CreateSpace mocks base method
func (m *MockManager) CreateSpace(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "CreateSpace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSpace indicates an expected call of CreateSpace
func (mr *MockManagerMockRecorder) CreateSpace(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpace", reflect.TypeOf((*MockManager)(nil).CreateSpace), arg0, arg1)
}

// CreateSpaceQuota mocks base method
func (m *MockManager) CreateSpaceQuota(arg0 cloudcontroller.SpaceQuotaEntity) (string, error) {
	ret := m.ctrl.Call(m, "CreateSpaceQuota", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSpaceQuota indicates an expected call of CreateSpaceQuota
func (mr *MockManagerMockRecorder) CreateSpaceQuota(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpaceQuota", reflect.TypeOf((*MockManager)(nil).CreateSpaceQuota), arg0)
}

// DeleteOrg mocks base method
func (m *MockManager) DeleteOrg(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteOrg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrg indicates an expected call of DeleteOrg
func (mr *MockManagerMockRecorder) DeleteOrg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrg", reflect.TypeOf((*MockManager)(nil).DeleteOrg), arg0)
}

// DeleteOrgByName mocks base method
func (m *MockManager) DeleteOrgByName(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteOrgByName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrgByName indicates an expected call of DeleteOrgByName
func (mr *MockManagerMockRecorder) DeleteOrgByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrgByName", reflect.TypeOf((*MockManager)(nil).DeleteOrgByName), arg0)
}

// DeletePrivateDomain mocks base method
func (m *MockManager) DeletePrivateDomain(arg0 string) error {
	ret := m.ctrl.Call(m, "DeletePrivateDomain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePrivateDomain indicates an expected call of DeletePrivateDomain
func (mr *MockManagerMockRecorder) DeletePrivateDomain(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrivateDomain", reflect.TypeOf((*MockManager)(nil).DeletePrivateDomain), arg0)
}

// DeleteSpace mocks base method
func (m *MockManager) DeleteSpace(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteSpace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSpace indicates an expected call of DeleteSpace
func (mr *MockManagerMockRecorder) DeleteSpace(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSpace", reflect.TypeOf((*MockManager)(nil).DeleteSpace), arg0)
}

// GetCFUsers mocks base method
func (m *MockManager) GetCFUsers(arg0, arg1, arg2 string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "GetCFUsers", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCFUsers indicates an expected call of GetCFUsers
func (mr *MockManagerMockRecorder) GetCFUsers(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCFUsers", reflect.TypeOf((*MockManager)(nil).GetCFUsers), arg0, arg1, arg2)
}

// GetSecurityGroupRules mocks base method
func (m *MockManager) GetSecurityGroupRules(arg0 string) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetSecurityGroupRules", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecurityGroupRules indicates an expected call of GetSecurityGroupRules
func (mr *MockManagerMockRecorder) GetSecurityGroupRules(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityGroupRules", reflect.TypeOf((*MockManager)(nil).GetSecurityGroupRules), arg0)
}

// ListAllOrgQuotas mocks base method
func (m *MockManager) ListAllOrgQuotas() (map[string]string, error) {
	ret := m.ctrl.Call(m, "ListAllOrgQuotas")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllOrgQuotas indicates an expected call of ListAllOrgQuotas
func (mr *MockManagerMockRecorder) ListAllOrgQuotas() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllOrgQuotas", reflect.TypeOf((*MockManager)(nil).ListAllOrgQuotas))
}

// ListAllPrivateDomains mocks base method
func (m *MockManager) ListAllPrivateDomains() (map[string]cloudcontroller.PrivateDomainInfo, error) {
	ret := m.ctrl.Call(m, "ListAllPrivateDomains")
	ret0, _ := ret[0].(map[string]cloudcontroller.PrivateDomainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllPrivateDomains indicates an expected call of ListAllPrivateDomains
func (mr *MockManagerMockRecorder) ListAllPrivateDomains() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllPrivateDomains", reflect.TypeOf((*MockManager)(nil).ListAllPrivateDomains))
}

// ListAllSpaceQuotasForOrg mocks base method
func (m *MockManager) ListAllSpaceQuotasForOrg(arg0 string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "ListAllSpaceQuotasForOrg", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllSpaceQuotasForOrg indicates an expected call of ListAllSpaceQuotasForOrg
func (mr *MockManagerMockRecorder) ListAllSpaceQuotasForOrg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllSpaceQuotasForOrg", reflect.TypeOf((*MockManager)(nil).ListAllSpaceQuotasForOrg), arg0)
}

// ListNonDefaultSecurityGroups mocks base method
func (m *MockManager) ListNonDefaultSecurityGroups() (map[string]cloudcontroller.SecurityGroupInfo, error) {
	ret := m.ctrl.Call(m, "ListNonDefaultSecurityGroups")
	ret0, _ := ret[0].(map[string]cloudcontroller.SecurityGroupInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNonDefaultSecurityGroups indicates an expected call of ListNonDefaultSecurityGroups
func (mr *MockManagerMockRecorder) ListNonDefaultSecurityGroups() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNonDefaultSecurityGroups", reflect.TypeOf((*MockManager)(nil).ListNonDefaultSecurityGroups))
}

// ListOrgOwnedPrivateDomains mocks base method
func (m *MockManager) ListOrgOwnedPrivateDomains(arg0 string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "ListOrgOwnedPrivateDomains", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrgOwnedPrivateDomains indicates an expected call of ListOrgOwnedPrivateDomains
func (mr *MockManagerMockRecorder) ListOrgOwnedPrivateDomains(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrgOwnedPrivateDomains", reflect.TypeOf((*MockManager)(nil).ListOrgOwnedPrivateDomains), arg0)
}

// ListOrgSharedPrivateDomains mocks base method
func (m *MockManager) ListOrgSharedPrivateDomains(arg0 string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "ListOrgSharedPrivateDomains", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrgSharedPrivateDomains indicates an expected call of ListOrgSharedPrivateDomains
func (mr *MockManagerMockRecorder) ListOrgSharedPrivateDomains(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrgSharedPrivateDomains", reflect.TypeOf((*MockManager)(nil).ListOrgSharedPrivateDomains), arg0)
}

// ListOrgs mocks base method
func (m *MockManager) ListOrgs() ([]*cloudcontroller.Org, error) {
	ret := m.ctrl.Call(m, "ListOrgs")
	ret0, _ := ret[0].([]*cloudcontroller.Org)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrgs indicates an expected call of ListOrgs
func (mr *MockManagerMockRecorder) ListOrgs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrgs", reflect.TypeOf((*MockManager)(nil).ListOrgs))
}

// ListSpaceSecurityGroups mocks base method
func (m *MockManager) ListSpaceSecurityGroups(arg0 string) (map[string]string, error) {
	ret := m.ctrl.Call(m, "ListSpaceSecurityGroups", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSpaceSecurityGroups indicates an expected call of ListSpaceSecurityGroups
func (mr *MockManagerMockRecorder) ListSpaceSecurityGroups(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSpaceSecurityGroups", reflect.TypeOf((*MockManager)(nil).ListSpaceSecurityGroups), arg0)
}

// ListSpaces mocks base method
func (m *MockManager) ListSpaces(arg0 string) ([]*cloudcontroller.Space, error) {
	ret := m.ctrl.Call(m, "ListSpaces", arg0)
	ret0, _ := ret[0].([]*cloudcontroller.Space)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSpaces indicates an expected call of ListSpaces
func (mr *MockManagerMockRecorder) ListSpaces(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSpaces", reflect.TypeOf((*MockManager)(nil).ListSpaces), arg0)
}

// QuotaDef mocks base method
func (m *MockManager) QuotaDef(arg0, arg1 string) (*cloudcontroller.Quota, error) {
	ret := m.ctrl.Call(m, "QuotaDef", arg0, arg1)
	ret0, _ := ret[0].(*cloudcontroller.Quota)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuotaDef indicates an expected call of QuotaDef
func (mr *MockManagerMockRecorder) QuotaDef(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuotaDef", reflect.TypeOf((*MockManager)(nil).QuotaDef), arg0, arg1)
}

// RemoveCFUser mocks base method
func (m *MockManager) RemoveCFUser(arg0, arg1, arg2, arg3 string) error {
	ret := m.ctrl.Call(m, "RemoveCFUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCFUser indicates an expected call of RemoveCFUser
func (mr *MockManagerMockRecorder) RemoveCFUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCFUser", reflect.TypeOf((*MockManager)(nil).RemoveCFUser), arg0, arg1, arg2, arg3)
}

// RemoveSharedPrivateDomain mocks base method
func (m *MockManager) RemoveSharedPrivateDomain(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "RemoveSharedPrivateDomain", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSharedPrivateDomain indicates an expected call of RemoveSharedPrivateDomain
func (mr *MockManagerMockRecorder) RemoveSharedPrivateDomain(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSharedPrivateDomain", reflect.TypeOf((*MockManager)(nil).RemoveSharedPrivateDomain), arg0, arg1)
}

// SharePrivateDomain mocks base method
func (m *MockManager) SharePrivateDomain(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SharePrivateDomain", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SharePrivateDomain indicates an expected call of SharePrivateDomain
func (mr *MockManagerMockRecorder) SharePrivateDomain(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SharePrivateDomain", reflect.TypeOf((*MockManager)(nil).SharePrivateDomain), arg0, arg1)
}

// UpdateQuota mocks base method
func (m *MockManager) UpdateQuota(arg0 string, arg1 cloudcontroller.QuotaEntity) error {
	ret := m.ctrl.Call(m, "UpdateQuota", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuota indicates an expected call of UpdateQuota
func (mr *MockManagerMockRecorder) UpdateQuota(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuota", reflect.TypeOf((*MockManager)(nil).UpdateQuota), arg0, arg1)
}

// UpdateSecurityGroup mocks base method
func (m *MockManager) UpdateSecurityGroup(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "UpdateSecurityGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecurityGroup indicates an expected call of UpdateSecurityGroup
func (mr *MockManagerMockRecorder) UpdateSecurityGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecurityGroup", reflect.TypeOf((*MockManager)(nil).UpdateSecurityGroup), arg0, arg1, arg2)
}

// UpdateSpaceQuota mocks base method
func (m *MockManager) UpdateSpaceQuota(arg0 string, arg1 cloudcontroller.SpaceQuotaEntity) error {
	ret := m.ctrl.Call(m, "UpdateSpaceQuota", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSpaceQuota indicates an expected call of UpdateSpaceQuota
func (mr *MockManagerMockRecorder) UpdateSpaceQuota(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSpaceQuota", reflect.TypeOf((*MockManager)(nil).UpdateSpaceQuota), arg0, arg1)
}

// UpdateSpaceSSH mocks base method
func (m *MockManager) UpdateSpaceSSH(arg0 bool, arg1 string) error {
	ret := m.ctrl.Call(m, "UpdateSpaceSSH", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSpaceSSH indicates an expected call of UpdateSpaceSSH
func (mr *MockManagerMockRecorder) UpdateSpaceSSH(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSpaceSSH", reflect.TypeOf((*MockManager)(nil).UpdateSpaceSSH), arg0, arg1)
}
