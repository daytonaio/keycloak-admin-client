/*
Keycloak Admin REST API

Testing ProtocolMappersAPIService

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

func Test_keycloak_admin_client_ProtocolMappersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersAddModelsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersAddModelsPost(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var id string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsIdDelete(context.Background(), realm, clientScopeId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var id string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsIdGet(context.Background(), realm, clientScopeId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var id string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsIdPut(context.Background(), realm, clientScopeId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersModelsPost(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientScopesClientScopeIdProtocolMappersProtocolProtocolGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var protocol string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientScopesClientScopeIdProtocolMappersProtocolProtocolGet(context.Background(), realm, clientScopeId, protocol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersAddModelsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersAddModelsPost(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var id string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsIdDelete(context.Background(), realm, clientScopeId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var id string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsIdGet(context.Background(), realm, clientScopeId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var id string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsIdPut(context.Background(), realm, clientScopeId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersModelsPost(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersProtocolProtocolGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var protocol string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientTemplatesClientScopeIdProtocolMappersProtocolProtocolGet(context.Background(), realm, clientScopeId, protocol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersAddModelsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersAddModelsPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersModelsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersModelsGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersModelsIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var id string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersModelsIdDelete(context.Background(), realm, clientUuid, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersModelsIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var id string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersModelsIdGet(context.Background(), realm, clientUuid, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersModelsIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var id string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersModelsIdPut(context.Background(), realm, clientUuid, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersModelsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersModelsPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProtocolMappersAPIService AdminRealmsRealmClientsClientUuidProtocolMappersProtocolProtocolGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var protocol string

		resp, httpRes, err := apiClient.ProtocolMappersAPI.AdminRealmsRealmClientsClientUuidProtocolMappersProtocolProtocolGet(context.Background(), realm, clientUuid, protocol).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
