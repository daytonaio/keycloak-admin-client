# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmUsersCountGet**](UsersAPI.md#AdminRealmsRealmUsersCountGet) | **Get** /admin/realms/{realm}/users/count | Returns the number of users that match the given criteria.
[**AdminRealmsRealmUsersGet**](UsersAPI.md#AdminRealmsRealmUsersGet) | **Get** /admin/realms/{realm}/users | Get users Returns a stream of users, filtered according to query parameters.
[**AdminRealmsRealmUsersPost**](UsersAPI.md#AdminRealmsRealmUsersPost) | **Post** /admin/realms/{realm}/users | Create a new user Username must be unique.
[**AdminRealmsRealmUsersProfileGet**](UsersAPI.md#AdminRealmsRealmUsersProfileGet) | **Get** /admin/realms/{realm}/users/profile | 
[**AdminRealmsRealmUsersProfileMetadataGet**](UsersAPI.md#AdminRealmsRealmUsersProfileMetadataGet) | **Get** /admin/realms/{realm}/users/profile/metadata | 
[**AdminRealmsRealmUsersProfilePut**](UsersAPI.md#AdminRealmsRealmUsersProfilePut) | **Put** /admin/realms/{realm}/users/profile | 
[**AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet) | **Get** /admin/realms/{realm}/users/{user-id}/configured-user-storage-credential-types | Return credential types, which are provided by the user storage where user is stored.
[**AdminRealmsRealmUsersUserIdConsentsClientDelete**](UsersAPI.md#AdminRealmsRealmUsersUserIdConsentsClientDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/consents/{client} | Revoke consent and offline tokens for particular client from user
[**AdminRealmsRealmUsersUserIdConsentsGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdConsentsGet) | **Get** /admin/realms/{realm}/users/{user-id}/consents | Get consents granted by the user
[**AdminRealmsRealmUsersUserIdCredentialsCredentialIdDelete**](UsersAPI.md#AdminRealmsRealmUsersUserIdCredentialsCredentialIdDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/credentials/{credentialId} | Remove a credential for a user
[**AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPost**](UsersAPI.md#AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPost) | **Post** /admin/realms/{realm}/users/{user-id}/credentials/{credentialId}/moveAfter/{newPreviousCredentialId} | Move a credential to a position behind another credential
[**AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPost**](UsersAPI.md#AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPost) | **Post** /admin/realms/{realm}/users/{user-id}/credentials/{credentialId}/moveToFirst | Move a credential to a first position in the credentials list of the user
[**AdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPut) | **Put** /admin/realms/{realm}/users/{user-id}/credentials/{credentialId}/userLabel | Update a credential label for a user
[**AdminRealmsRealmUsersUserIdCredentialsGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdCredentialsGet) | **Get** /admin/realms/{realm}/users/{user-id}/credentials | 
[**AdminRealmsRealmUsersUserIdDelete**](UsersAPI.md#AdminRealmsRealmUsersUserIdDelete) | **Delete** /admin/realms/{realm}/users/{user-id} | Delete the user
[**AdminRealmsRealmUsersUserIdDisableCredentialTypesPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdDisableCredentialTypesPut) | **Put** /admin/realms/{realm}/users/{user-id}/disable-credential-types | Disable all credentials for a user of a specific type
[**AdminRealmsRealmUsersUserIdExecuteActionsEmailPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdExecuteActionsEmailPut) | **Put** /admin/realms/{realm}/users/{user-id}/execute-actions-email | Send an email to the user with a link they can click to execute particular actions.
[**AdminRealmsRealmUsersUserIdFederatedIdentityGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdFederatedIdentityGet) | **Get** /admin/realms/{realm}/users/{user-id}/federated-identity | Get social logins associated with the user
[**AdminRealmsRealmUsersUserIdFederatedIdentityProviderDelete**](UsersAPI.md#AdminRealmsRealmUsersUserIdFederatedIdentityProviderDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/federated-identity/{provider} | Remove a social login provider from user
[**AdminRealmsRealmUsersUserIdFederatedIdentityProviderPost**](UsersAPI.md#AdminRealmsRealmUsersUserIdFederatedIdentityProviderPost) | **Post** /admin/realms/{realm}/users/{user-id}/federated-identity/{provider} | Add a social login provider to the user
[**AdminRealmsRealmUsersUserIdGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdGet) | **Get** /admin/realms/{realm}/users/{user-id} | Get representation of the user
[**AdminRealmsRealmUsersUserIdGroupsCountGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdGroupsCountGet) | **Get** /admin/realms/{realm}/users/{user-id}/groups/count | 
[**AdminRealmsRealmUsersUserIdGroupsGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdGroupsGet) | **Get** /admin/realms/{realm}/users/{user-id}/groups | 
[**AdminRealmsRealmUsersUserIdGroupsGroupIdDelete**](UsersAPI.md#AdminRealmsRealmUsersUserIdGroupsGroupIdDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/groups/{groupId} | 
[**AdminRealmsRealmUsersUserIdGroupsGroupIdPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdGroupsGroupIdPut) | **Put** /admin/realms/{realm}/users/{user-id}/groups/{groupId} | 
[**AdminRealmsRealmUsersUserIdImpersonationPost**](UsersAPI.md#AdminRealmsRealmUsersUserIdImpersonationPost) | **Post** /admin/realms/{realm}/users/{user-id}/impersonation | Impersonate the user
[**AdminRealmsRealmUsersUserIdLogoutPost**](UsersAPI.md#AdminRealmsRealmUsersUserIdLogoutPost) | **Post** /admin/realms/{realm}/users/{user-id}/logout | Remove all user sessions associated with the user Also send notification to all clients that have an admin URL to invalidate the sessions for the particular user.
[**AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet) | **Get** /admin/realms/{realm}/users/{user-id}/offline-sessions/{clientUuid} | Get offline sessions associated with the user and client
[**AdminRealmsRealmUsersUserIdPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdPut) | **Put** /admin/realms/{realm}/users/{user-id} | Update the user
[**AdminRealmsRealmUsersUserIdResetPasswordEmailPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdResetPasswordEmailPut) | **Put** /admin/realms/{realm}/users/{user-id}/reset-password-email | Send an email to the user with a link they can click to reset their password.
[**AdminRealmsRealmUsersUserIdResetPasswordPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdResetPasswordPut) | **Put** /admin/realms/{realm}/users/{user-id}/reset-password | Set up a new password for the user.
[**AdminRealmsRealmUsersUserIdSendVerifyEmailPut**](UsersAPI.md#AdminRealmsRealmUsersUserIdSendVerifyEmailPut) | **Put** /admin/realms/{realm}/users/{user-id}/send-verify-email | Send an email-verification email to the user An email contains a link the user can click to verify their email address.
[**AdminRealmsRealmUsersUserIdSessionsGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdSessionsGet) | **Get** /admin/realms/{realm}/users/{user-id}/sessions | Get sessions associated with the user
[**AdminRealmsRealmUsersUserIdUnmanagedAttributesGet**](UsersAPI.md#AdminRealmsRealmUsersUserIdUnmanagedAttributesGet) | **Get** /admin/realms/{realm}/users/{user-id}/unmanagedAttributes | 



## AdminRealmsRealmUsersCountGet

> int32 AdminRealmsRealmUsersCountGet(ctx, realm).Email(email).EmailVerified(emailVerified).Enabled(enabled).FirstName(firstName).LastName(lastName).Q(q).Search(search).Username(username).Execute()

Returns the number of users that match the given criteria.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	email := "email_example" // string | email filter (optional)
	emailVerified := true // bool |  (optional)
	enabled := true // bool | Boolean representing if user is enabled or not (optional)
	firstName := "firstName_example" // string | first name filter (optional)
	lastName := "lastName_example" // string | last name filter (optional)
	q := "q_example" // string |  (optional)
	search := "search_example" // string | arbitrary search string for all the fields below. Default search behavior is prefix-based (e.g., foo or foo*). Use *foo* for infix search and \"foo\" for exact search. (optional)
	username := "username_example" // string | username filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersCountGet(context.Background(), realm).Email(email).EmailVerified(emailVerified).Enabled(enabled).FirstName(firstName).LastName(lastName).Q(q).Search(search).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersCountGet`: int32
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **string** | email filter | 
 **emailVerified** | **bool** |  | 
 **enabled** | **bool** | Boolean representing if user is enabled or not | 
 **firstName** | **string** | first name filter | 
 **lastName** | **string** | last name filter | 
 **q** | **string** |  | 
 **search** | **string** | arbitrary search string for all the fields below. Default search behavior is prefix-based (e.g., foo or foo*). Use *foo* for infix search and \&quot;foo\&quot; for exact search. | 
 **username** | **string** | username filter | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersGet

> []UserRepresentation AdminRealmsRealmUsersGet(ctx, realm).BriefRepresentation(briefRepresentation).Email(email).EmailVerified(emailVerified).Enabled(enabled).Exact(exact).First(first).FirstName(firstName).IdpAlias(idpAlias).IdpUserId(idpUserId).LastName(lastName).Max(max).Q(q).Search(search).Username(username).Execute()

Get users Returns a stream of users, filtered according to query parameters.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	briefRepresentation := true // bool | Boolean which defines whether brief representations are returned (default: false) (optional)
	email := "email_example" // string | A String contained in email, or the complete email, if param \"exact\" is true (optional)
	emailVerified := true // bool | whether the email has been verified (optional)
	enabled := true // bool | Boolean representing if user is enabled or not (optional)
	exact := true // bool | Boolean which defines whether the params \"last\", \"first\", \"email\" and \"username\" must match exactly (optional)
	first := int32(56) // int32 | Pagination offset (optional)
	firstName := "firstName_example" // string | A String contained in firstName, or the complete firstName, if param \"exact\" is true (optional)
	idpAlias := "idpAlias_example" // string | The alias of an Identity Provider linked to the user (optional)
	idpUserId := "idpUserId_example" // string | The userId at an Identity Provider linked to the user (optional)
	lastName := "lastName_example" // string | A String contained in lastName, or the complete lastName, if param \"exact\" is true (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)
	q := "q_example" // string | A query to search for custom attributes, in the format 'key1:value2 key2:value2' (optional)
	search := "search_example" // string | A String contained in username, first or last name, or email. Default search behavior is prefix-based (e.g., foo or foo*). Use *foo* for infix search and \"foo\" for exact search. (optional)
	username := "username_example" // string | A String contained in username, or the complete username, if param \"exact\" is true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersGet(context.Background(), realm).BriefRepresentation(briefRepresentation).Email(email).EmailVerified(emailVerified).Enabled(enabled).Exact(exact).First(first).FirstName(firstName).IdpAlias(idpAlias).IdpUserId(idpUserId).LastName(lastName).Max(max).Q(q).Search(search).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersGet`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **bool** | Boolean which defines whether brief representations are returned (default: false) | 
 **email** | **string** | A String contained in email, or the complete email, if param \&quot;exact\&quot; is true | 
 **emailVerified** | **bool** | whether the email has been verified | 
 **enabled** | **bool** | Boolean representing if user is enabled or not | 
 **exact** | **bool** | Boolean which defines whether the params \&quot;last\&quot;, \&quot;first\&quot;, \&quot;email\&quot; and \&quot;username\&quot; must match exactly | 
 **first** | **int32** | Pagination offset | 
 **firstName** | **string** | A String contained in firstName, or the complete firstName, if param \&quot;exact\&quot; is true | 
 **idpAlias** | **string** | The alias of an Identity Provider linked to the user | 
 **idpUserId** | **string** | The userId at an Identity Provider linked to the user | 
 **lastName** | **string** | A String contained in lastName, or the complete lastName, if param \&quot;exact\&quot; is true | 
 **max** | **int32** | Maximum results size (defaults to 100) | 
 **q** | **string** | A query to search for custom attributes, in the format &#39;key1:value2 key2:value2&#39; | 
 **search** | **string** | A String contained in username, first or last name, or email. Default search behavior is prefix-based (e.g., foo or foo*). Use *foo* for infix search and \&quot;foo\&quot; for exact search. | 
 **username** | **string** | A String contained in username, or the complete username, if param \&quot;exact\&quot; is true | 

### Return type

[**[]UserRepresentation**](UserRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersPost

> AdminRealmsRealmUsersPost(ctx, realm).UserRepresentation(userRepresentation).Execute()

Create a new user Username must be unique.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userRepresentation := *openapiclient.NewUserRepresentation() // UserRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersPost(context.Background(), realm).UserRepresentation(userRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRepresentation** | [**UserRepresentation**](UserRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersProfileGet

> UPConfig AdminRealmsRealmUsersProfileGet(ctx, realm).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersProfileGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersProfileGet`: UPConfig
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersProfileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UPConfig**](UPConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersProfileMetadataGet

> UserProfileMetadata AdminRealmsRealmUsersProfileMetadataGet(ctx, realm).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersProfileMetadataGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersProfileMetadataGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersProfileMetadataGet`: UserProfileMetadata
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersProfileMetadataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersProfileMetadataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserProfileMetadata**](UserProfileMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersProfilePut

> UPConfig AdminRealmsRealmUsersProfilePut(ctx, realm).UPConfig(uPConfig).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	uPConfig := *openapiclient.NewUPConfig() // UPConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersProfilePut(context.Background(), realm).UPConfig(uPConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersProfilePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersProfilePut`: UPConfig
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersProfilePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersProfilePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uPConfig** | [**UPConfig**](UPConfig.md) |  | 

### Return type

[**UPConfig**](UPConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet

> []string AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet(ctx, realm, userId).Execute()

Return credential types, which are provided by the user storage where user is stored.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdConfiguredUserStorageCredentialTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdConsentsClientDelete

> AdminRealmsRealmUsersUserIdConsentsClientDelete(ctx, realm, userId, client).Execute()

Revoke consent and offline tokens for particular client from user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	client := "client_example" // string | Client id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdConsentsClientDelete(context.Background(), realm, userId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdConsentsClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**client** | **string** | Client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdConsentsClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdConsentsGet

> []map[string]interface{} AdminRealmsRealmUsersUserIdConsentsGet(ctx, realm, userId).Execute()

Get consents granted by the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdConsentsGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdConsentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdConsentsGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdConsentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdConsentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdCredentialsCredentialIdDelete

> AdminRealmsRealmUsersUserIdCredentialsCredentialIdDelete(ctx, realm, userId, credentialId).Execute()

Remove a credential for a user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	credentialId := "credentialId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdDelete(context.Background(), realm, userId, credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdCredentialsCredentialIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPost

> AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPost(ctx, realm, userId, credentialId, newPreviousCredentialId).Execute()

Move a credential to a position behind another credential

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	credentialId := "credentialId_example" // string | The credential to move
	newPreviousCredentialId := "newPreviousCredentialId_example" // string | The credential that will be the previous element in the list. If set to null, the moved credential will be the first element in the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPost(context.Background(), realm, userId, credentialId, newPreviousCredentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**credentialId** | **string** | The credential to move | 
**newPreviousCredentialId** | **string** | The credential that will be the previous element in the list. If set to null, the moved credential will be the first element in the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveAfterNewPreviousCredentialIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPost

> AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPost(ctx, realm, userId, credentialId).Execute()

Move a credential to a first position in the credentials list of the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	credentialId := "credentialId_example" // string | The credential to move

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPost(context.Background(), realm, userId, credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**credentialId** | **string** | The credential to move | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdCredentialsCredentialIdMoveToFirstPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPut

> AdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPut(ctx, realm, userId, credentialId).Body(body).Execute()

Update a credential label for a user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	credentialId := "credentialId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPut(context.Background(), realm, userId, credentialId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**credentialId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdCredentialsCredentialIdUserLabelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdCredentialsGet

> []CredentialRepresentation AdminRealmsRealmUsersUserIdCredentialsGet(ctx, realm, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdCredentialsGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdCredentialsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdCredentialsGet`: []CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdCredentialsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdCredentialsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdDelete

> AdminRealmsRealmUsersUserIdDelete(ctx, realm, userId).Execute()

Delete the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdDelete(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdDisableCredentialTypesPut

> AdminRealmsRealmUsersUserIdDisableCredentialTypesPut(ctx, realm, userId).RequestBody(requestBody).Execute()

Disable all credentials for a user of a specific type

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdDisableCredentialTypesPut(context.Background(), realm, userId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdDisableCredentialTypesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdDisableCredentialTypesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdExecuteActionsEmailPut

> AdminRealmsRealmUsersUserIdExecuteActionsEmailPut(ctx, realm, userId).ClientId(clientId).Lifespan(lifespan).RedirectUri(redirectUri).RequestBody(requestBody).Execute()

Send an email to the user with a link they can click to execute particular actions.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | Client id (optional)
	lifespan := int32(56) // int32 | Number of seconds after which the generated token expires (optional)
	redirectUri := "redirectUri_example" // string | Redirect uri (optional)
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdExecuteActionsEmailPut(context.Background(), realm, userId).ClientId(clientId).Lifespan(lifespan).RedirectUri(redirectUri).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdExecuteActionsEmailPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdExecuteActionsEmailPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** | Client id | 
 **lifespan** | **int32** | Number of seconds after which the generated token expires | 
 **redirectUri** | **string** | Redirect uri | 
 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdFederatedIdentityGet

> []FederatedIdentityRepresentation AdminRealmsRealmUsersUserIdFederatedIdentityGet(ctx, realm, userId).Execute()

Get social logins associated with the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdFederatedIdentityGet`: []FederatedIdentityRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdFederatedIdentityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]FederatedIdentityRepresentation**](FederatedIdentityRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdFederatedIdentityProviderDelete

> AdminRealmsRealmUsersUserIdFederatedIdentityProviderDelete(ctx, realm, userId, provider).Execute()

Remove a social login provider from user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	provider := "provider_example" // string | Social login provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityProviderDelete(context.Background(), realm, userId, provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityProviderDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**provider** | **string** | Social login provider id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdFederatedIdentityProviderDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdFederatedIdentityProviderPost

> AdminRealmsRealmUsersUserIdFederatedIdentityProviderPost(ctx, realm, userId, provider).Execute()

Add a social login provider to the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	provider := "provider_example" // string | Social login provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityProviderPost(context.Background(), realm, userId, provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdFederatedIdentityProviderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**provider** | **string** | Social login provider id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdFederatedIdentityProviderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdGet

> UserRepresentation AdminRealmsRealmUsersUserIdGet(ctx, realm, userId).UserProfileMetadata(userProfileMetadata).Execute()

Get representation of the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	userProfileMetadata := true // bool | Indicates if the user profile metadata should be added to the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdGet(context.Background(), realm, userId).UserProfileMetadata(userProfileMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdGet`: UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userProfileMetadata** | **bool** | Indicates if the user profile metadata should be added to the response | 

### Return type

[**UserRepresentation**](UserRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdGroupsCountGet

> map[string]int64 AdminRealmsRealmUsersUserIdGroupsCountGet(ctx, realm, userId).Search(search).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdGroupsCountGet(context.Background(), realm, userId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdGroupsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdGroupsCountGet`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdGroupsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdGroupsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **search** | **string** |  | 

### Return type

**map[string]int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdGroupsGet

> []GroupRepresentation AdminRealmsRealmUsersUserIdGroupsGet(ctx, realm, userId).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	briefRepresentation := true // bool |  (optional) (default to true)
	first := int32(56) // int32 |  (optional)
	max := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdGroupsGet(context.Background(), realm, userId).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdGroupsGet`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** |  | [default to true]
 **first** | **int32** |  | 
 **max** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]GroupRepresentation**](GroupRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdGroupsGroupIdDelete

> AdminRealmsRealmUsersUserIdGroupsGroupIdDelete(ctx, realm, userId, groupId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdGroupsGroupIdDelete(context.Background(), realm, userId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdGroupsGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdGroupsGroupIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdGroupsGroupIdPut

> AdminRealmsRealmUsersUserIdGroupsGroupIdPut(ctx, realm, userId, groupId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdGroupsGroupIdPut(context.Background(), realm, userId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdGroupsGroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdGroupsGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdImpersonationPost

> map[string]interface{} AdminRealmsRealmUsersUserIdImpersonationPost(ctx, realm, userId).Execute()

Impersonate the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdImpersonationPost(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdImpersonationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdImpersonationPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdImpersonationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdImpersonationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdLogoutPost

> AdminRealmsRealmUsersUserIdLogoutPost(ctx, realm, userId).Execute()

Remove all user sessions associated with the user Also send notification to all clients that have an admin URL to invalidate the sessions for the particular user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdLogoutPost(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdLogoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet

> []UserSessionRepresentation AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet(ctx, realm, userId, clientUuid).Execute()

Get offline sessions associated with the user and client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	clientUuid := "clientUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet(context.Background(), realm, userId, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet`: []UserSessionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**clientUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdOfflineSessionsClientUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]UserSessionRepresentation**](UserSessionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdPut

> AdminRealmsRealmUsersUserIdPut(ctx, realm, userId).UserRepresentation(userRepresentation).Execute()

Update the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	userRepresentation := *openapiclient.NewUserRepresentation() // UserRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdPut(context.Background(), realm, userId).UserRepresentation(userRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userRepresentation** | [**UserRepresentation**](UserRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdResetPasswordEmailPut

> AdminRealmsRealmUsersUserIdResetPasswordEmailPut(ctx, realm, userId).ClientId(clientId).RedirectUri(redirectUri).Execute()

Send an email to the user with a link they can click to reset their password.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | client id (optional)
	redirectUri := "redirectUri_example" // string | redirect uri (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdResetPasswordEmailPut(context.Background(), realm, userId).ClientId(clientId).RedirectUri(redirectUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdResetPasswordEmailPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdResetPasswordEmailPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** | client id | 
 **redirectUri** | **string** | redirect uri | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdResetPasswordPut

> AdminRealmsRealmUsersUserIdResetPasswordPut(ctx, realm, userId).CredentialRepresentation(credentialRepresentation).Execute()

Set up a new password for the user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	credentialRepresentation := *openapiclient.NewCredentialRepresentation() // CredentialRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdResetPasswordPut(context.Background(), realm, userId).CredentialRepresentation(credentialRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdResetPasswordPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdResetPasswordPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialRepresentation** | [**CredentialRepresentation**](CredentialRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdSendVerifyEmailPut

> AdminRealmsRealmUsersUserIdSendVerifyEmailPut(ctx, realm, userId).ClientId(clientId).Lifespan(lifespan).RedirectUri(redirectUri).Execute()

Send an email-verification email to the user An email contains a link the user can click to verify their email address.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | Client id (optional)
	lifespan := int32(56) // int32 | Number of seconds after which the generated token expires (optional)
	redirectUri := "redirectUri_example" // string | Redirect uri (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdSendVerifyEmailPut(context.Background(), realm, userId).ClientId(clientId).Lifespan(lifespan).RedirectUri(redirectUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdSendVerifyEmailPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdSendVerifyEmailPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientId** | **string** | Client id | 
 **lifespan** | **int32** | Number of seconds after which the generated token expires | 
 **redirectUri** | **string** | Redirect uri | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdSessionsGet

> []UserSessionRepresentation AdminRealmsRealmUsersUserIdSessionsGet(ctx, realm, userId).Execute()

Get sessions associated with the user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdSessionsGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdSessionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdSessionsGet`: []UserSessionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]UserSessionRepresentation**](UserSessionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersUserIdUnmanagedAttributesGet

> map[string][]string AdminRealmsRealmUsersUserIdUnmanagedAttributesGet(ctx, realm, userId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AdminRealmsRealmUsersUserIdUnmanagedAttributesGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AdminRealmsRealmUsersUserIdUnmanagedAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdUnmanagedAttributesGet`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AdminRealmsRealmUsersUserIdUnmanagedAttributesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdUnmanagedAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

