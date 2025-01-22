# \ClientsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientsClientUuidClientSecretGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidClientSecretGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/client-secret | Get the client secret
[**AdminRealmsRealmClientsClientUuidClientSecretPost**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidClientSecretPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/client-secret | Generate a new secret for the client
[**AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/client-secret/rotated | Invalidate the rotated secret for the client
[**AdminRealmsRealmClientsClientUuidClientSecretRotatedGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidClientSecretRotatedGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/client-secret/rotated | Get the rotated client secret
[**AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/default-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut) | **Put** /admin/realms/{realm}/clients/{client-uuid}/default-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmClientsClientUuidDefaultClientScopesGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidDefaultClientScopesGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/default-client-scopes | Get default client scopes.  Only name and ids are returned.
[**AdminRealmsRealmClientsClientUuidDelete**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid} | Delete the client
[**AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/evaluate-scopes/generate-example-access-token | Create JSON with payload of example access token
[**AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/evaluate-scopes/generate-example-id-token | Create JSON with payload of example id token
[**AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/evaluate-scopes/generate-example-userinfo | Create JSON with payload of example user info
[**AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/evaluate-scopes/protocol-mappers | Return list of all protocol mappers, which will be used when generating tokens issued for particular client.
[**AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/evaluate-scopes/scope-mappings/{roleContainerId}/granted | Get effective scope mapping of all roles of particular role container, which this client is defacto allowed to have in the accessToken issued for him.
[**AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/evaluate-scopes/scope-mappings/{roleContainerId}/not-granted | Get roles, which this client doesn&#39;t have scope for and can&#39;t have them in the accessToken issued for him.
[**AdminRealmsRealmClientsClientUuidGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidGet) | **Get** /admin/realms/{realm}/clients/{client-uuid} | Get representation of the client
[**AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/installation/providers/{providerId} | 
[**AdminRealmsRealmClientsClientUuidManagementPermissionsGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidManagementPermissionsGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/management/permissions | Return object stating whether client Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmClientsClientUuidManagementPermissionsPut**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidManagementPermissionsPut) | **Put** /admin/realms/{realm}/clients/{client-uuid}/management/permissions | Return object stating whether client Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmClientsClientUuidNodesNodeDelete**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidNodesNodeDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/nodes/{node} | Unregister a cluster node from the client
[**AdminRealmsRealmClientsClientUuidNodesPost**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidNodesPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/nodes | Register a cluster node with the client Manually register cluster node to this client - usually it’s not needed to call this directly as adapter should handle by sending registration request to Keycloak
[**AdminRealmsRealmClientsClientUuidOfflineSessionCountGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidOfflineSessionCountGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/offline-session-count | Get application offline session count Returns a number of offline user sessions associated with this client { \&quot;count\&quot;: number }
[**AdminRealmsRealmClientsClientUuidOfflineSessionsGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidOfflineSessionsGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/offline-sessions | Get offline sessions for client Returns a list of offline user sessions associated with this client
[**AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/optional-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut) | **Put** /admin/realms/{realm}/clients/{client-uuid}/optional-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmClientsClientUuidOptionalClientScopesGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidOptionalClientScopesGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/optional-client-scopes | Get optional client scopes.  Only name and ids are returned.
[**AdminRealmsRealmClientsClientUuidPushRevocationPost**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidPushRevocationPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/push-revocation | Push the client&#39;s revocation policy to its admin URL If the client has an admin URL, push revocation policy to it.
[**AdminRealmsRealmClientsClientUuidPut**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidPut) | **Put** /admin/realms/{realm}/clients/{client-uuid} | Update the client
[**AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/registration-access-token | Generate a new registration access token for the client
[**AdminRealmsRealmClientsClientUuidServiceAccountUserGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidServiceAccountUserGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/service-account-user | Get a user dedicated to the service account
[**AdminRealmsRealmClientsClientUuidSessionCountGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidSessionCountGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/session-count | Get application session count Returns a number of user sessions associated with this client { \&quot;count\&quot;: number }
[**AdminRealmsRealmClientsClientUuidTestNodesAvailableGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidTestNodesAvailableGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/test-nodes-available | Test if registered cluster nodes are available Tests availability by sending &#39;ping&#39; request to all cluster nodes.
[**AdminRealmsRealmClientsClientUuidUserSessionsGet**](ClientsAPI.md#AdminRealmsRealmClientsClientUuidUserSessionsGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/user-sessions | Get user sessions for client Returns a list of user sessions associated with this client 
[**AdminRealmsRealmClientsGet**](ClientsAPI.md#AdminRealmsRealmClientsGet) | **Get** /admin/realms/{realm}/clients | Get clients belonging to the realm.
[**AdminRealmsRealmClientsPost**](ClientsAPI.md#AdminRealmsRealmClientsPost) | **Post** /admin/realms/{realm}/clients | Create a new client Client’s client_id must be unique!



## AdminRealmsRealmClientsClientUuidClientSecretGet

> CredentialRepresentation AdminRealmsRealmClientsClientUuidClientSecretGet(ctx, realm, clientUuid).Execute()

Get the client secret

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidClientSecretGet`: CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidClientSecretGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidClientSecretPost

> CredentialRepresentation AdminRealmsRealmClientsClientUuidClientSecretPost(ctx, realm, clientUuid).Execute()

Generate a new secret for the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretPost(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidClientSecretPost`: CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidClientSecretPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete

> AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete(ctx, realm, clientUuid).Execute()

Invalidate the rotated secret for the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidClientSecretRotatedDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidClientSecretRotatedGet

> CredentialRepresentation AdminRealmsRealmClientsClientUuidClientSecretRotatedGet(ctx, realm, clientUuid).Execute()

Get the rotated client secret

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidClientSecretRotatedGet`: CredentialRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidClientSecretRotatedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidClientSecretRotatedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialRepresentation**](CredentialRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete

> AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete(ctx, realm, clientUuid, clientScopeId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete(context.Background(), realm, clientUuid, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut

> AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut(ctx, realm, clientUuid, clientScopeId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut(context.Background(), realm, clientUuid, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidDefaultClientScopesClientScopeIdPutRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidDefaultClientScopesGet

> []ClientScopeRepresentation AdminRealmsRealmClientsClientUuidDefaultClientScopesGet(ctx, realm, clientUuid).Execute()

Get default client scopes.  Only name and ids are returned.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidDefaultClientScopesGet`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidDefaultClientScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidDefaultClientScopesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ClientScopeRepresentation**](ClientScopeRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidDelete

> AdminRealmsRealmClientsClientUuidDelete(ctx, realm, clientUuid).Execute()

Delete the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidDelete(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet

> AccessToken AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet(ctx, realm, clientUuid).Scope(scope).UserId(userId).Execute()

Create JSON with payload of example access token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	scope := "scope_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet(context.Background(), realm, clientUuid).Scope(scope).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet`: AccessToken
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleAccessTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 
 **userId** | **string** |  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet

> IDToken AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet(ctx, realm, clientUuid).Scope(scope).UserId(userId).Execute()

Create JSON with payload of example id token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	scope := "scope_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet(context.Background(), realm, clientUuid).Scope(scope).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet`: IDToken
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleIdTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 
 **userId** | **string** |  | 

### Return type

[**IDToken**](IDToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet

> map[string]interface{} AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet(ctx, realm, clientUuid).Scope(scope).UserId(userId).Execute()

Create JSON with payload of example user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	scope := "scope_example" // string |  (optional)
	userId := "userId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet(context.Background(), realm, clientUuid).Scope(scope).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidEvaluateScopesGenerateExampleUserinfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 
 **userId** | **string** |  | 

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


## AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet

> []ProtocolMapperEvaluationRepresentation AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet(ctx, realm, clientUuid).Scope(scope).Execute()

Return list of all protocol mappers, which will be used when generating tokens issued for particular client.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet(context.Background(), realm, clientUuid).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet`: []ProtocolMapperEvaluationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidEvaluateScopesProtocolMappersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | **string** |  | 

### Return type

[**[]ProtocolMapperEvaluationRepresentation**](ProtocolMapperEvaluationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet(ctx, realm, clientUuid, roleContainerId).Scope(scope).Execute()

Get effective scope mapping of all roles of particular role container, which this client is defacto allowed to have in the accessToken issued for him.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	roleContainerId := "roleContainerId_example" // string | either realm name OR client UUID
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet(context.Background(), realm, clientUuid, roleContainerId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleContainerId** | **string** | either realm name OR client UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdGrantedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scope** | **string** |  | 

### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet(ctx, realm, clientUuid, roleContainerId).Scope(scope).Execute()

Get roles, which this client doesn't have scope for and can't have them in the accessToken issued for him.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	roleContainerId := "roleContainerId_example" // string | either realm name OR client UUID
	scope := "scope_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet(context.Background(), realm, clientUuid, roleContainerId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleContainerId** | **string** | either realm name OR client UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidEvaluateScopesScopeMappingsRoleContainerIdNotGrantedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **scope** | **string** |  | 

### Return type

[**[]RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidGet

> ClientRepresentation AdminRealmsRealmClientsClientUuidGet(ctx, realm, clientUuid).Execute()

Get representation of the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidGet`: ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientRepresentation**](ClientRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet

> AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet(ctx, realm, clientUuid, providerId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	providerId := "providerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet(context.Background(), realm, clientUuid, providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidInstallationProvidersProviderIdGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmClientsClientUuidManagementPermissionsGet(ctx, realm, clientUuid).Execute()

Return object stating whether client Authorization permissions have been initialized or not and a reference

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidManagementPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ManagementPermissionReference**](ManagementPermissionReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmClientsClientUuidManagementPermissionsPut(ctx, realm, clientUuid).ManagementPermissionReference(managementPermissionReference).Execute()

Return object stating whether client Authorization permissions have been initialized or not and a reference

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsPut(context.Background(), realm, clientUuid).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidManagementPermissionsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **managementPermissionReference** | [**ManagementPermissionReference**](ManagementPermissionReference.md) |  | 

### Return type

[**ManagementPermissionReference**](ManagementPermissionReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidNodesNodeDelete

> AdminRealmsRealmClientsClientUuidNodesNodeDelete(ctx, realm, clientUuid, node).Execute()

Unregister a cluster node from the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	node := "node_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidNodesNodeDelete(context.Background(), realm, clientUuid, node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidNodesNodeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**node** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidNodesNodeDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidNodesPost

> AdminRealmsRealmClientsClientUuidNodesPost(ctx, realm, clientUuid).RequestBody(requestBody).Execute()

Register a cluster node with the client Manually register cluster node to this client - usually it’s not needed to call this directly as adapter should handle by sending registration request to Keycloak

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidNodesPost(context.Background(), realm, clientUuid).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidNodesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidNodesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]string** |  | 

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


## AdminRealmsRealmClientsClientUuidOfflineSessionCountGet

> map[string]int64 AdminRealmsRealmClientsClientUuidOfflineSessionCountGet(ctx, realm, clientUuid).Execute()

Get application offline session count Returns a number of offline user sessions associated with this client { \"count\": number }

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionCountGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidOfflineSessionCountGet`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidOfflineSessionCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientsClientUuidOfflineSessionsGet

> []UserSessionRepresentation AdminRealmsRealmClientsClientUuidOfflineSessionsGet(ctx, realm, clientUuid).First(first).Max(max).Execute()

Get offline sessions for client Returns a list of offline user sessions associated with this client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	first := int32(56) // int32 | Paging offset (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionsGet(context.Background(), realm, clientUuid).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidOfflineSessionsGet`: []UserSessionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidOfflineSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidOfflineSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **int32** | Paging offset | 
 **max** | **int32** | Maximum results size (defaults to 100) | 

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


## AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete

> AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete(ctx, realm, clientUuid, clientScopeId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete(context.Background(), realm, clientUuid, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut

> AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut(ctx, realm, clientUuid, clientScopeId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut(context.Background(), realm, clientUuid, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidOptionalClientScopesClientScopeIdPutRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidOptionalClientScopesGet

> []ClientScopeRepresentation AdminRealmsRealmClientsClientUuidOptionalClientScopesGet(ctx, realm, clientUuid).Execute()

Get optional client scopes.  Only name and ids are returned.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidOptionalClientScopesGet`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidOptionalClientScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidOptionalClientScopesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ClientScopeRepresentation**](ClientScopeRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidPushRevocationPost

> GlobalRequestResult AdminRealmsRealmClientsClientUuidPushRevocationPost(ctx, realm, clientUuid).Execute()

Push the client's revocation policy to its admin URL If the client has an admin URL, push revocation policy to it.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidPushRevocationPost(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidPushRevocationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidPushRevocationPost`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidPushRevocationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidPushRevocationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GlobalRequestResult**](GlobalRequestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidPut

> AdminRealmsRealmClientsClientUuidPut(ctx, realm, clientUuid).ClientRepresentation(clientRepresentation).Execute()

Update the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	clientRepresentation := *openapiclient.NewClientRepresentation() // ClientRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidPut(context.Background(), realm, clientUuid).ClientRepresentation(clientRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientRepresentation** | [**ClientRepresentation**](ClientRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost

> ClientRepresentation AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost(ctx, realm, clientUuid).Execute()

Generate a new registration access token for the client

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost`: ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidRegistrationAccessTokenPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRegistrationAccessTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientRepresentation**](ClientRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidServiceAccountUserGet

> UserRepresentation AdminRealmsRealmClientsClientUuidServiceAccountUserGet(ctx, realm, clientUuid).Execute()

Get a user dedicated to the service account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidServiceAccountUserGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidServiceAccountUserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidServiceAccountUserGet`: UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidServiceAccountUserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidServiceAccountUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientsClientUuidSessionCountGet

> map[string]int64 AdminRealmsRealmClientsClientUuidSessionCountGet(ctx, realm, clientUuid).Execute()

Get application session count Returns a number of user sessions associated with this client { \"count\": number }

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidSessionCountGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidSessionCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidSessionCountGet`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidSessionCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidSessionCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientsClientUuidTestNodesAvailableGet

> GlobalRequestResult AdminRealmsRealmClientsClientUuidTestNodesAvailableGet(ctx, realm, clientUuid).Execute()

Test if registered cluster nodes are available Tests availability by sending 'ping' request to all cluster nodes.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidTestNodesAvailableGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidTestNodesAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidTestNodesAvailableGet`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidTestNodesAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidTestNodesAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GlobalRequestResult**](GlobalRequestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidUserSessionsGet

> []UserSessionRepresentation AdminRealmsRealmClientsClientUuidUserSessionsGet(ctx, realm, clientUuid).First(first).Max(max).Execute()

Get user sessions for client Returns a list of user sessions associated with this client 

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	first := int32(56) // int32 | Paging offset (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsClientUuidUserSessionsGet(context.Background(), realm, clientUuid).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsClientUuidUserSessionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidUserSessionsGet`: []UserSessionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsClientUuidUserSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidUserSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **int32** | Paging offset | 
 **max** | **int32** | Maximum results size (defaults to 100) | 

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


## AdminRealmsRealmClientsGet

> []ClientRepresentation AdminRealmsRealmClientsGet(ctx, realm).ClientId(clientId).First(first).Max(max).Q(q).Search(search).ViewableOnly(viewableOnly).Execute()

Get clients belonging to the realm.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientId := "clientId_example" // string | filter by clientId (optional)
	first := int32(56) // int32 | the first result (optional)
	max := int32(56) // int32 | the max results to return (optional)
	q := "q_example" // string |  (optional)
	search := true // bool | whether this is a search query or a getClientById query (optional) (default to false)
	viewableOnly := true // bool | filter clients that cannot be viewed in full by admin (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsGet(context.Background(), realm).ClientId(clientId).First(first).Max(max).Q(q).Search(search).ViewableOnly(viewableOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsGet`: []ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.AdminRealmsRealmClientsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientId** | **string** | filter by clientId | 
 **first** | **int32** | the first result | 
 **max** | **int32** | the max results to return | 
 **q** | **string** |  | 
 **search** | **bool** | whether this is a search query or a getClientById query | [default to false]
 **viewableOnly** | **bool** | filter clients that cannot be viewed in full by admin | [default to false]

### Return type

[**[]ClientRepresentation**](ClientRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsPost

> AdminRealmsRealmClientsPost(ctx, realm).ClientRepresentation(clientRepresentation).Execute()

Create a new client Client’s client_id must be unique!

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/daytona-ee-draft/pkg/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientRepresentation := *openapiclient.NewClientRepresentation() // ClientRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.AdminRealmsRealmClientsPost(context.Background(), realm).ClientRepresentation(clientRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.AdminRealmsRealmClientsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientRepresentation** | [**ClientRepresentation**](ClientRepresentation.md) |  | 

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

