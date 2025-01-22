/*
Keycloak Admin REST API

Testing ClientsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package keycloak_admin_client

import (
	"context"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_keycloak_admin_client_ClientsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidClientSecretGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidClientSecretPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidClientSecretRotatedGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var clientScopeId string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete(context.Background(), realm, clientUuid, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var clientScopeId string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut(context.Background(), realm, clientUuid, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidDefaultClientScopesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDelete(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var roleContainerId string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet(context.Background(), realm, clientUuid, roleContainerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var roleContainerId string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet(context.Background(), realm, clientUuid, roleContainerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var providerId string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet(context.Background(), realm, clientUuid, providerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidManagementPermissionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidManagementPermissionsPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsPut(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidNodesNodeDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var node string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidNodesNodeDelete(context.Background(), realm, clientUuid, node).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidNodesPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidNodesPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidOfflineSessionCountGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionCountGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidOfflineSessionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionsGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var clientScopeId string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete(context.Background(), realm, clientUuid, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var clientScopeId string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut(context.Background(), realm, clientUuid, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidOptionalClientScopesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidPushRevocationPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidPushRevocationPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidPut(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidServiceAccountUserGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidServiceAccountUserGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidSessionCountGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidSessionCountGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidTestNodesAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidTestNodesAvailableGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsClientUuidUserSessionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidUserSessionsGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientsAPIService AdminRealmsRealmClientsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.ClientsAPI.AdminRealmsRealmClientsPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
