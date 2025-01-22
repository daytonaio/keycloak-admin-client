/*
Keycloak Admin REST API

Testing RealmsAdminAPIService

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

func Test_keycloak_admin_client_RealmsAdminAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RealmsAdminAPIService AdminRealmsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmAdminEventsDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmAdminEventsDelete(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmAdminEventsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmAdminEventsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientDescriptionConverterPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientDescriptionConverterPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientPoliciesPoliciesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientPoliciesPoliciesPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesPut(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientPoliciesProfilesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientPoliciesProfilesPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesPut(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientSessionStatsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientSessionStatsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientTypesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientTypesGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmClientTypesPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientTypesPut(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmCredentialRegistratorsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmCredentialRegistratorsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultDefaultClientScopesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultGroupsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultGroupsGroupIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var groupId string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGroupIdDelete(context.Background(), realm, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultGroupsGroupIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var groupId string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGroupIdPut(context.Background(), realm, groupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var clientScopeId string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut(context.Background(), realm, clientScopeId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDefaultOptionalClientScopesGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDelete(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmEventsConfigGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsConfigGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmEventsConfigPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsConfigPut(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmEventsDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsDelete(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmEventsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmGroupByPathPathGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var path string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmGroupByPathPathGet(context.Background(), realm, path).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationLocaleDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var locale string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleDelete(context.Background(), realm, locale).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationLocaleGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var locale string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleGet(context.Background(), realm, locale).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationLocaleKeyDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var key string
		var locale string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyDelete(context.Background(), realm, key, locale).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationLocaleKeyGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var key string
		var locale string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyGet(context.Background(), realm, key, locale).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationLocaleKeyPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var key string
		var locale string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyPut(context.Background(), realm, key, locale).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLocalizationLocalePost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var locale string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocalePost(context.Background(), realm, locale).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmLogoutAllPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLogoutAllPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmPartialExportPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPartialExportPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmPartialImportPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPartialImportPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmPushRevocationPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPushRevocationPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPut(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmSessionsSessionDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string
		var session string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmSessionsSessionDelete(context.Background(), realm, session).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmTestSMTPConnectionPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmTestSMTPConnectionPost(context.Background(), realm).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmUsersManagementPermissionsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsGet(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RealmsAdminAPIService AdminRealmsRealmUsersManagementPermissionsPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var realm string

		resp, httpRes, err := apiClient.RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsPut(context.Background(), realm).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
