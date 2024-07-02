// Code generated by MockGen. DO NOT EDIT.
// Source: handler_token.go

// Package handler_test is a generated GoMock package.
package handler_test

import (
	reflect "reflect"
	time "time"

	event "github.com/authgear/authgear-server/pkg/api/event"
	app2app "github.com/authgear/authgear-server/pkg/lib/app2app"
	challenge "github.com/authgear/authgear-server/pkg/lib/authn/challenge"
	user "github.com/authgear/authgear-server/pkg/lib/authn/user"
	config "github.com/authgear/authgear-server/pkg/lib/config"
	oauth "github.com/authgear/authgear-server/pkg/lib/oauth"
	handler "github.com/authgear/authgear-server/pkg/lib/oauth/handler"
	oidc "github.com/authgear/authgear-server/pkg/lib/oauth/oidc"
	protocol "github.com/authgear/authgear-server/pkg/lib/oauth/protocol"
	access "github.com/authgear/authgear-server/pkg/lib/session/access"
	gomock "github.com/golang/mock/gomock"
	jwk "github.com/lestrrat-go/jwx/v2/jwk"
	jwt "github.com/lestrrat-go/jwx/v2/jwt"
)

// MockIDTokenIssuer is a mock of IDTokenIssuer interface.
type MockIDTokenIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockIDTokenIssuerMockRecorder
}

// MockIDTokenIssuerMockRecorder is the mock recorder for MockIDTokenIssuer.
type MockIDTokenIssuerMockRecorder struct {
	mock *MockIDTokenIssuer
}

// NewMockIDTokenIssuer creates a new mock instance.
func NewMockIDTokenIssuer(ctrl *gomock.Controller) *MockIDTokenIssuer {
	mock := &MockIDTokenIssuer{ctrl: ctrl}
	mock.recorder = &MockIDTokenIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDTokenIssuer) EXPECT() *MockIDTokenIssuerMockRecorder {
	return m.recorder
}

// Iss mocks base method.
func (m *MockIDTokenIssuer) Iss() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iss")
	ret0, _ := ret[0].(string)
	return ret0
}

// Iss indicates an expected call of Iss.
func (mr *MockIDTokenIssuerMockRecorder) Iss() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iss", reflect.TypeOf((*MockIDTokenIssuer)(nil).Iss))
}

// IssueIDToken mocks base method.
func (m *MockIDTokenIssuer) IssueIDToken(opts oidc.IssueIDTokenOptions) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueIDToken", opts)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueIDToken indicates an expected call of IssueIDToken.
func (mr *MockIDTokenIssuerMockRecorder) IssueIDToken(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueIDToken", reflect.TypeOf((*MockIDTokenIssuer)(nil).IssueIDToken), opts)
}

// VerifyIDTokenWithoutClient mocks base method.
func (m *MockIDTokenIssuer) VerifyIDTokenWithoutClient(idToken string) (jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIDTokenWithoutClient", idToken)
	ret0, _ := ret[0].(jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyIDTokenWithoutClient indicates an expected call of VerifyIDTokenWithoutClient.
func (mr *MockIDTokenIssuerMockRecorder) VerifyIDTokenWithoutClient(idToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIDTokenWithoutClient", reflect.TypeOf((*MockIDTokenIssuer)(nil).VerifyIDTokenWithoutClient), idToken)
}

// MockAccessTokenIssuer is a mock of AccessTokenIssuer interface.
type MockAccessTokenIssuer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenIssuerMockRecorder
}

// MockAccessTokenIssuerMockRecorder is the mock recorder for MockAccessTokenIssuer.
type MockAccessTokenIssuerMockRecorder struct {
	mock *MockAccessTokenIssuer
}

// NewMockAccessTokenIssuer creates a new mock instance.
func NewMockAccessTokenIssuer(ctrl *gomock.Controller) *MockAccessTokenIssuer {
	mock := &MockAccessTokenIssuer{ctrl: ctrl}
	mock.recorder = &MockAccessTokenIssuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessTokenIssuer) EXPECT() *MockAccessTokenIssuerMockRecorder {
	return m.recorder
}

// EncodeAccessToken mocks base method.
func (m *MockAccessTokenIssuer) EncodeAccessToken(client *config.OAuthClientConfig, grant *oauth.AccessGrant, userID, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EncodeAccessToken", client, grant, userID, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EncodeAccessToken indicates an expected call of EncodeAccessToken.
func (mr *MockAccessTokenIssuerMockRecorder) EncodeAccessToken(client, grant, userID, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EncodeAccessToken", reflect.TypeOf((*MockAccessTokenIssuer)(nil).EncodeAccessToken), client, grant, userID, token)
}

// MockEventService is a mock of EventService interface.
type MockEventService struct {
	ctrl     *gomock.Controller
	recorder *MockEventServiceMockRecorder
}

// MockEventServiceMockRecorder is the mock recorder for MockEventService.
type MockEventServiceMockRecorder struct {
	mock *MockEventService
}

// NewMockEventService creates a new mock instance.
func NewMockEventService(ctrl *gomock.Controller) *MockEventService {
	mock := &MockEventService{ctrl: ctrl}
	mock.recorder = &MockEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventService) EXPECT() *MockEventServiceMockRecorder {
	return m.recorder
}

// DispatchEventOnCommit mocks base method.
func (m *MockEventService) DispatchEventOnCommit(payload event.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DispatchEventOnCommit", payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// DispatchEventOnCommit indicates an expected call of DispatchEventOnCommit.
func (mr *MockEventServiceMockRecorder) DispatchEventOnCommit(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchEventOnCommit", reflect.TypeOf((*MockEventService)(nil).DispatchEventOnCommit), payload)
}

// MockTokenHandlerUserFacade is a mock of TokenHandlerUserFacade interface.
type MockTokenHandlerUserFacade struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerUserFacadeMockRecorder
}

// MockTokenHandlerUserFacadeMockRecorder is the mock recorder for MockTokenHandlerUserFacade.
type MockTokenHandlerUserFacadeMockRecorder struct {
	mock *MockTokenHandlerUserFacade
}

// NewMockTokenHandlerUserFacade creates a new mock instance.
func NewMockTokenHandlerUserFacade(ctrl *gomock.Controller) *MockTokenHandlerUserFacade {
	mock := &MockTokenHandlerUserFacade{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerUserFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerUserFacade) EXPECT() *MockTokenHandlerUserFacadeMockRecorder {
	return m.recorder
}

// GetRaw mocks base method.
func (m *MockTokenHandlerUserFacade) GetRaw(id string) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRaw", id)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRaw indicates an expected call of GetRaw.
func (mr *MockTokenHandlerUserFacadeMockRecorder) GetRaw(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRaw", reflect.TypeOf((*MockTokenHandlerUserFacade)(nil).GetRaw), id)
}

// MockApp2AppService is a mock of App2AppService interface.
type MockApp2AppService struct {
	ctrl     *gomock.Controller
	recorder *MockApp2AppServiceMockRecorder
}

// MockApp2AppServiceMockRecorder is the mock recorder for MockApp2AppService.
type MockApp2AppServiceMockRecorder struct {
	mock *MockApp2AppService
}

// NewMockApp2AppService creates a new mock instance.
func NewMockApp2AppService(ctrl *gomock.Controller) *MockApp2AppService {
	mock := &MockApp2AppService{ctrl: ctrl}
	mock.recorder = &MockApp2AppServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApp2AppService) EXPECT() *MockApp2AppServiceMockRecorder {
	return m.recorder
}

// ParseToken mocks base method.
func (m *MockApp2AppService) ParseToken(requestJWT string, key jwk.Key) (*app2app.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", requestJWT, key)
	ret0, _ := ret[0].(*app2app.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockApp2AppServiceMockRecorder) ParseToken(requestJWT, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockApp2AppService)(nil).ParseToken), requestJWT, key)
}

// ParseTokenUnverified mocks base method.
func (m *MockApp2AppService) ParseTokenUnverified(requestJWT string) (*app2app.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseTokenUnverified", requestJWT)
	ret0, _ := ret[0].(*app2app.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseTokenUnverified indicates an expected call of ParseTokenUnverified.
func (mr *MockApp2AppServiceMockRecorder) ParseTokenUnverified(requestJWT interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseTokenUnverified", reflect.TypeOf((*MockApp2AppService)(nil).ParseTokenUnverified), requestJWT)
}

// MockChallengeProvider is a mock of ChallengeProvider interface.
type MockChallengeProvider struct {
	ctrl     *gomock.Controller
	recorder *MockChallengeProviderMockRecorder
}

// MockChallengeProviderMockRecorder is the mock recorder for MockChallengeProvider.
type MockChallengeProviderMockRecorder struct {
	mock *MockChallengeProvider
}

// NewMockChallengeProvider creates a new mock instance.
func NewMockChallengeProvider(ctrl *gomock.Controller) *MockChallengeProvider {
	mock := &MockChallengeProvider{ctrl: ctrl}
	mock.recorder = &MockChallengeProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChallengeProvider) EXPECT() *MockChallengeProviderMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockChallengeProvider) Consume(token string) (*challenge.Purpose, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", token)
	ret0, _ := ret[0].(*challenge.Purpose)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockChallengeProviderMockRecorder) Consume(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockChallengeProvider)(nil).Consume), token)
}

// MockTokenHandlerCodeGrantStore is a mock of TokenHandlerCodeGrantStore interface.
type MockTokenHandlerCodeGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerCodeGrantStoreMockRecorder
}

// MockTokenHandlerCodeGrantStoreMockRecorder is the mock recorder for MockTokenHandlerCodeGrantStore.
type MockTokenHandlerCodeGrantStoreMockRecorder struct {
	mock *MockTokenHandlerCodeGrantStore
}

// NewMockTokenHandlerCodeGrantStore creates a new mock instance.
func NewMockTokenHandlerCodeGrantStore(ctrl *gomock.Controller) *MockTokenHandlerCodeGrantStore {
	mock := &MockTokenHandlerCodeGrantStore{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerCodeGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerCodeGrantStore) EXPECT() *MockTokenHandlerCodeGrantStoreMockRecorder {
	return m.recorder
}

// DeleteCodeGrant mocks base method.
func (m *MockTokenHandlerCodeGrantStore) DeleteCodeGrant(arg0 *oauth.CodeGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCodeGrant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCodeGrant indicates an expected call of DeleteCodeGrant.
func (mr *MockTokenHandlerCodeGrantStoreMockRecorder) DeleteCodeGrant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCodeGrant", reflect.TypeOf((*MockTokenHandlerCodeGrantStore)(nil).DeleteCodeGrant), arg0)
}

// GetCodeGrant mocks base method.
func (m *MockTokenHandlerCodeGrantStore) GetCodeGrant(codeHash string) (*oauth.CodeGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodeGrant", codeHash)
	ret0, _ := ret[0].(*oauth.CodeGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCodeGrant indicates an expected call of GetCodeGrant.
func (mr *MockTokenHandlerCodeGrantStoreMockRecorder) GetCodeGrant(codeHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeGrant", reflect.TypeOf((*MockTokenHandlerCodeGrantStore)(nil).GetCodeGrant), codeHash)
}

// MockTokenHandlerSettingsActionGrantStore is a mock of TokenHandlerSettingsActionGrantStore interface.
type MockTokenHandlerSettingsActionGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerSettingsActionGrantStoreMockRecorder
}

// MockTokenHandlerSettingsActionGrantStoreMockRecorder is the mock recorder for MockTokenHandlerSettingsActionGrantStore.
type MockTokenHandlerSettingsActionGrantStoreMockRecorder struct {
	mock *MockTokenHandlerSettingsActionGrantStore
}

// NewMockTokenHandlerSettingsActionGrantStore creates a new mock instance.
func NewMockTokenHandlerSettingsActionGrantStore(ctrl *gomock.Controller) *MockTokenHandlerSettingsActionGrantStore {
	mock := &MockTokenHandlerSettingsActionGrantStore{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerSettingsActionGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerSettingsActionGrantStore) EXPECT() *MockTokenHandlerSettingsActionGrantStoreMockRecorder {
	return m.recorder
}

// DeleteSettingsActionGrant mocks base method.
func (m *MockTokenHandlerSettingsActionGrantStore) DeleteSettingsActionGrant(arg0 *oauth.SettingsActionGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSettingsActionGrant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettingsActionGrant indicates an expected call of DeleteSettingsActionGrant.
func (mr *MockTokenHandlerSettingsActionGrantStoreMockRecorder) DeleteSettingsActionGrant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettingsActionGrant", reflect.TypeOf((*MockTokenHandlerSettingsActionGrantStore)(nil).DeleteSettingsActionGrant), arg0)
}

// GetSettingsActionGrant mocks base method.
func (m *MockTokenHandlerSettingsActionGrantStore) GetSettingsActionGrant(codeHash string) (*oauth.SettingsActionGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettingsActionGrant", codeHash)
	ret0, _ := ret[0].(*oauth.SettingsActionGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettingsActionGrant indicates an expected call of GetSettingsActionGrant.
func (mr *MockTokenHandlerSettingsActionGrantStoreMockRecorder) GetSettingsActionGrant(codeHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettingsActionGrant", reflect.TypeOf((*MockTokenHandlerSettingsActionGrantStore)(nil).GetSettingsActionGrant), codeHash)
}

// MockTokenHandlerOfflineGrantStore is a mock of TokenHandlerOfflineGrantStore interface.
type MockTokenHandlerOfflineGrantStore struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerOfflineGrantStoreMockRecorder
}

// MockTokenHandlerOfflineGrantStoreMockRecorder is the mock recorder for MockTokenHandlerOfflineGrantStore.
type MockTokenHandlerOfflineGrantStoreMockRecorder struct {
	mock *MockTokenHandlerOfflineGrantStore
}

// NewMockTokenHandlerOfflineGrantStore creates a new mock instance.
func NewMockTokenHandlerOfflineGrantStore(ctrl *gomock.Controller) *MockTokenHandlerOfflineGrantStore {
	mock := &MockTokenHandlerOfflineGrantStore{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerOfflineGrantStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerOfflineGrantStore) EXPECT() *MockTokenHandlerOfflineGrantStoreMockRecorder {
	return m.recorder
}

// AccessOfflineGrantAndUpdateDeviceInfo mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) AccessOfflineGrantAndUpdateDeviceInfo(id string, accessEvent access.Event, deviceInfo map[string]interface{}, expireAt time.Time) (*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessOfflineGrantAndUpdateDeviceInfo", id, accessEvent, deviceInfo, expireAt)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccessOfflineGrantAndUpdateDeviceInfo indicates an expected call of AccessOfflineGrantAndUpdateDeviceInfo.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) AccessOfflineGrantAndUpdateDeviceInfo(id, accessEvent, deviceInfo, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessOfflineGrantAndUpdateDeviceInfo", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).AccessOfflineGrantAndUpdateDeviceInfo), id, accessEvent, deviceInfo, expireAt)
}

// DeleteOfflineGrant mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) DeleteOfflineGrant(arg0 *oauth.OfflineGrant) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOfflineGrant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOfflineGrant indicates an expected call of DeleteOfflineGrant.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) DeleteOfflineGrant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOfflineGrant", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).DeleteOfflineGrant), arg0)
}

// GetOfflineGrant mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) GetOfflineGrant(id string) (*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOfflineGrant", id)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOfflineGrant indicates an expected call of GetOfflineGrant.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) GetOfflineGrant(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOfflineGrant", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).GetOfflineGrant), id)
}

// ListClientOfflineGrants mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) ListClientOfflineGrants(clientID, userID string) ([]*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClientOfflineGrants", clientID, userID)
	ret0, _ := ret[0].([]*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClientOfflineGrants indicates an expected call of ListClientOfflineGrants.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) ListClientOfflineGrants(clientID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClientOfflineGrants", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).ListClientOfflineGrants), clientID, userID)
}

// ListOfflineGrants mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) ListOfflineGrants(userID string) ([]*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOfflineGrants", userID)
	ret0, _ := ret[0].([]*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOfflineGrants indicates an expected call of ListOfflineGrants.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) ListOfflineGrants(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOfflineGrants", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).ListOfflineGrants), userID)
}

// UpdateOfflineGrantApp2AppDeviceKey mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) UpdateOfflineGrantApp2AppDeviceKey(id, newKey string, expireAt time.Time) (*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantApp2AppDeviceKey", id, newKey, expireAt)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantApp2AppDeviceKey indicates an expected call of UpdateOfflineGrantApp2AppDeviceKey.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) UpdateOfflineGrantApp2AppDeviceKey(id, newKey, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantApp2AppDeviceKey", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).UpdateOfflineGrantApp2AppDeviceKey), id, newKey, expireAt)
}

// UpdateOfflineGrantAuthenticatedAt mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) UpdateOfflineGrantAuthenticatedAt(id string, authenticatedAt, expireAt time.Time) (*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantAuthenticatedAt", id, authenticatedAt, expireAt)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantAuthenticatedAt indicates an expected call of UpdateOfflineGrantAuthenticatedAt.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) UpdateOfflineGrantAuthenticatedAt(id, authenticatedAt, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantAuthenticatedAt", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).UpdateOfflineGrantAuthenticatedAt), id, authenticatedAt, expireAt)
}

// UpdateOfflineGrantDeviceSecretHash mocks base method.
func (m *MockTokenHandlerOfflineGrantStore) UpdateOfflineGrantDeviceSecretHash(grantID, newDeviceSecretHash string, expireAt time.Time) (*oauth.OfflineGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOfflineGrantDeviceSecretHash", grantID, newDeviceSecretHash, expireAt)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOfflineGrantDeviceSecretHash indicates an expected call of UpdateOfflineGrantDeviceSecretHash.
func (mr *MockTokenHandlerOfflineGrantStoreMockRecorder) UpdateOfflineGrantDeviceSecretHash(grantID, newDeviceSecretHash, expireAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOfflineGrantDeviceSecretHash", reflect.TypeOf((*MockTokenHandlerOfflineGrantStore)(nil).UpdateOfflineGrantDeviceSecretHash), grantID, newDeviceSecretHash, expireAt)
}

// MockTokenHandlerAppSessionTokenStore is a mock of TokenHandlerAppSessionTokenStore interface.
type MockTokenHandlerAppSessionTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerAppSessionTokenStoreMockRecorder
}

// MockTokenHandlerAppSessionTokenStoreMockRecorder is the mock recorder for MockTokenHandlerAppSessionTokenStore.
type MockTokenHandlerAppSessionTokenStoreMockRecorder struct {
	mock *MockTokenHandlerAppSessionTokenStore
}

// NewMockTokenHandlerAppSessionTokenStore creates a new mock instance.
func NewMockTokenHandlerAppSessionTokenStore(ctrl *gomock.Controller) *MockTokenHandlerAppSessionTokenStore {
	mock := &MockTokenHandlerAppSessionTokenStore{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerAppSessionTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerAppSessionTokenStore) EXPECT() *MockTokenHandlerAppSessionTokenStoreMockRecorder {
	return m.recorder
}

// CreateAppSessionToken mocks base method.
func (m *MockTokenHandlerAppSessionTokenStore) CreateAppSessionToken(arg0 *oauth.AppSessionToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAppSessionToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAppSessionToken indicates an expected call of CreateAppSessionToken.
func (mr *MockTokenHandlerAppSessionTokenStoreMockRecorder) CreateAppSessionToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAppSessionToken", reflect.TypeOf((*MockTokenHandlerAppSessionTokenStore)(nil).CreateAppSessionToken), arg0)
}

// MockTokenHandlerOfflineGrantService is a mock of TokenHandlerOfflineGrantService interface.
type MockTokenHandlerOfflineGrantService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerOfflineGrantServiceMockRecorder
}

// MockTokenHandlerOfflineGrantServiceMockRecorder is the mock recorder for MockTokenHandlerOfflineGrantService.
type MockTokenHandlerOfflineGrantServiceMockRecorder struct {
	mock *MockTokenHandlerOfflineGrantService
}

// NewMockTokenHandlerOfflineGrantService creates a new mock instance.
func NewMockTokenHandlerOfflineGrantService(ctrl *gomock.Controller) *MockTokenHandlerOfflineGrantService {
	mock := &MockTokenHandlerOfflineGrantService{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerOfflineGrantServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerOfflineGrantService) EXPECT() *MockTokenHandlerOfflineGrantServiceMockRecorder {
	return m.recorder
}

// ComputeOfflineGrantExpiry mocks base method.
func (m *MockTokenHandlerOfflineGrantService) ComputeOfflineGrantExpiry(session *oauth.OfflineGrant) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeOfflineGrantExpiry", session)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeOfflineGrantExpiry indicates an expected call of ComputeOfflineGrantExpiry.
func (mr *MockTokenHandlerOfflineGrantServiceMockRecorder) ComputeOfflineGrantExpiry(session interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeOfflineGrantExpiry", reflect.TypeOf((*MockTokenHandlerOfflineGrantService)(nil).ComputeOfflineGrantExpiry), session)
}

// MockTokenHandlerTokenService is a mock of TokenHandlerTokenService interface.
type MockTokenHandlerTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenHandlerTokenServiceMockRecorder
}

// MockTokenHandlerTokenServiceMockRecorder is the mock recorder for MockTokenHandlerTokenService.
type MockTokenHandlerTokenServiceMockRecorder struct {
	mock *MockTokenHandlerTokenService
}

// NewMockTokenHandlerTokenService creates a new mock instance.
func NewMockTokenHandlerTokenService(ctrl *gomock.Controller) *MockTokenHandlerTokenService {
	mock := &MockTokenHandlerTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenHandlerTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenHandlerTokenService) EXPECT() *MockTokenHandlerTokenServiceMockRecorder {
	return m.recorder
}

// IssueAccessGrant mocks base method.
func (m *MockTokenHandlerTokenService) IssueAccessGrant(client *config.OAuthClientConfig, scopes []string, authzID, userID, sessionID string, sessionKind oauth.GrantSessionKind, refreshTokenHash string, resp protocol.TokenResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueAccessGrant", client, scopes, authzID, userID, sessionID, sessionKind, refreshTokenHash, resp)
	ret0, _ := ret[0].(error)
	return ret0
}

// IssueAccessGrant indicates an expected call of IssueAccessGrant.
func (mr *MockTokenHandlerTokenServiceMockRecorder) IssueAccessGrant(client, scopes, authzID, userID, sessionID, sessionKind, refreshTokenHash, resp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueAccessGrant", reflect.TypeOf((*MockTokenHandlerTokenService)(nil).IssueAccessGrant), client, scopes, authzID, userID, sessionID, sessionKind, refreshTokenHash, resp)
}

// IssueDeviceSecret mocks base method.
func (m *MockTokenHandlerTokenService) IssueDeviceSecret(resp protocol.TokenResponse) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueDeviceSecret", resp)
	ret0, _ := ret[0].(string)
	return ret0
}

// IssueDeviceSecret indicates an expected call of IssueDeviceSecret.
func (mr *MockTokenHandlerTokenServiceMockRecorder) IssueDeviceSecret(resp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueDeviceSecret", reflect.TypeOf((*MockTokenHandlerTokenService)(nil).IssueDeviceSecret), resp)
}

// IssueOfflineGrant mocks base method.
func (m *MockTokenHandlerTokenService) IssueOfflineGrant(client *config.OAuthClientConfig, opts handler.IssueOfflineGrantOptions, resp protocol.TokenResponse) (*oauth.OfflineGrant, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueOfflineGrant", client, opts, resp)
	ret0, _ := ret[0].(*oauth.OfflineGrant)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IssueOfflineGrant indicates an expected call of IssueOfflineGrant.
func (mr *MockTokenHandlerTokenServiceMockRecorder) IssueOfflineGrant(client, opts, resp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueOfflineGrant", reflect.TypeOf((*MockTokenHandlerTokenService)(nil).IssueOfflineGrant), client, opts, resp)
}

// ParseRefreshToken mocks base method.
func (m *MockTokenHandlerTokenService) ParseRefreshToken(token string) (*oauth.Authorization, *oauth.OfflineGrant, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseRefreshToken", token)
	ret0, _ := ret[0].(*oauth.Authorization)
	ret1, _ := ret[1].(*oauth.OfflineGrant)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ParseRefreshToken indicates an expected call of ParseRefreshToken.
func (mr *MockTokenHandlerTokenServiceMockRecorder) ParseRefreshToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRefreshToken", reflect.TypeOf((*MockTokenHandlerTokenService)(nil).ParseRefreshToken), token)
}

// MockAppInitiatedSSOToWebTokenService is a mock of AppInitiatedSSOToWebTokenService interface.
type MockAppInitiatedSSOToWebTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockAppInitiatedSSOToWebTokenServiceMockRecorder
}

// MockAppInitiatedSSOToWebTokenServiceMockRecorder is the mock recorder for MockAppInitiatedSSOToWebTokenService.
type MockAppInitiatedSSOToWebTokenServiceMockRecorder struct {
	mock *MockAppInitiatedSSOToWebTokenService
}

// NewMockAppInitiatedSSOToWebTokenService creates a new mock instance.
func NewMockAppInitiatedSSOToWebTokenService(ctrl *gomock.Controller) *MockAppInitiatedSSOToWebTokenService {
	mock := &MockAppInitiatedSSOToWebTokenService{ctrl: ctrl}
	mock.recorder = &MockAppInitiatedSSOToWebTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppInitiatedSSOToWebTokenService) EXPECT() *MockAppInitiatedSSOToWebTokenServiceMockRecorder {
	return m.recorder
}

// IssueAppInitiatedSSOToWebToken mocks base method.
func (m *MockAppInitiatedSSOToWebTokenService) IssueAppInitiatedSSOToWebToken(options *oauth.IssueAppInitiatedSSOToWebTokenOptions) (*oauth.IssueAppInitiatedSSOToWebTokenResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueAppInitiatedSSOToWebToken", options)
	ret0, _ := ret[0].(*oauth.IssueAppInitiatedSSOToWebTokenResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueAppInitiatedSSOToWebToken indicates an expected call of IssueAppInitiatedSSOToWebToken.
func (mr *MockAppInitiatedSSOToWebTokenServiceMockRecorder) IssueAppInitiatedSSOToWebToken(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueAppInitiatedSSOToWebToken", reflect.TypeOf((*MockAppInitiatedSSOToWebTokenService)(nil).IssueAppInitiatedSSOToWebToken), options)
}
