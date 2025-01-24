/*
Keycloak Admin REST API

Testing ClientRoleMappingsAPIService

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

func Test_keycloak_admin_client_ClientRoleMappingsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var groupId string
		var clientId string

		resp, httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet(context.Background(), realm, groupId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var groupId string
		var clientId string

		resp, httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet(context.Background(), realm, groupId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var groupId string
		var clientId string

		httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete(context.Background(), realm, groupId, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var groupId string
		var clientId string

		resp, httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet(context.Background(), realm, groupId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var groupId string
		var clientId string

		httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost(context.Background(), realm, groupId, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string
		var clientId string

		resp, httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet(context.Background(), realm, userId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string
		var clientId string

		resp, httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet(context.Background(), realm, userId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string
		var clientId string

		httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete(context.Background(), realm, userId, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string
		var clientId string

		resp, httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet(context.Background(), realm, userId, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ClientRoleMappingsAPIService AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var realm string
		var userId string
		var clientId string

		httpRes, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost(context.Background(), realm, userId, clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
