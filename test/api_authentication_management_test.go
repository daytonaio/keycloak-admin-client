/*
Keycloak Admin REST API

Testing AuthenticationManagementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package keycloak_admin_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func Test_keycloak_admin_client_AuthenticationManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationAuthenticatorProvidersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationAuthenticatorProvidersGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var providerId string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet(context.Background(), realm, providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationConfigIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdDelete(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationConfigIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdGet(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationConfigIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdPut(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationConfigPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var executionId string
		var id string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet(context.Background(), realm, executionId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var executionId string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost(context.Background(), realm, executionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var executionId string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete(context.Background(), realm, executionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsExecutionIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var executionId string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdGet(context.Background(), realm, executionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var executionId string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost(context.Background(), realm, executionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var executionId string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost(context.Background(), realm, executionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationExecutionsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var flowAlias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost(context.Background(), realm, flowAlias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var flowAlias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost(context.Background(), realm, flowAlias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var flowAlias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost(context.Background(), realm, flowAlias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var flowAlias string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet(context.Background(), realm, flowAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var flowAlias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut(context.Background(), realm, flowAlias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdDelete(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdGet(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var id string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdPut(context.Background(), realm, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFlowsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFormActionProvidersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormActionProvidersGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationFormProvidersGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormProvidersGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRegisterRequiredActionPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRegisterRequiredActionPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasDelete(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasGet(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasPut(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var alias string

		httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost(context.Background(), realm, alias).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationRequiredActionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuthenticationManagementAPIService AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string

		resp, httpRes, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
