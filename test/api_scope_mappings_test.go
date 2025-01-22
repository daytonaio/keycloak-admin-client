/*
Keycloak Admin REST API

Testing ScopeMappingsAPIService

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

func Test_keycloak_admin_client_ScopeMappingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string
		var client string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost(context.Background(), realm, clientScopeId, client).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet(context.Background(), realm, clientUuid, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet(context.Background(), realm, clientUuid, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var client string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete(context.Background(), realm, clientUuid, client).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var client string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet(context.Background(), realm, clientUuid, client).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string
		var client string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost(context.Background(), realm, clientUuid, client).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		resp, httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScopeMappingsAPIService AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientUuid string

		httpRes, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost(context.Background(), realm, clientUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
