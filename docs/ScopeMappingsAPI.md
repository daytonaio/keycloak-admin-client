# \ScopeMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/clients/{client}/available | The available client-level roles Returns the roles for the client that can be associated with the client&#39;s scope
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/clients/{client}/composite | Get effective client roles Returns the roles for the client that are associated with the client&#39;s scope.
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete) | **Delete** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/clients/{client} | Remove client-level roles from the client&#39;s scope.
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/clients/{client} | Get the roles associated with a client&#39;s scope Returns roles for the client.
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost) | **Post** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/clients/{client} | Add client-level roles to the client&#39;s scope
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings | Get all scope mappings for the client
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/realm/available | Get realm-level roles that are available to attach to this client&#39;s scope
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/realm/composite | Get effective realm-level roles associated with the client’s scope What this does is recurse any composite roles associated with the client’s scope and adds the roles to this lists.
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete) | **Delete** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/realm | Remove a set of realm-level roles from the client&#39;s scope
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/realm | Get realm-level roles associated with the client&#39;s scope
[**AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost**](ScopeMappingsAPI.md#AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost) | **Post** /admin/realms/{realm}/client-scopes/{client-scope-id}/scope-mappings/realm | Add a set of realm-level roles to the client&#39;s scope
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/clients/{client}/available | The available client-level roles Returns the roles for the client that can be associated with the client&#39;s scope
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/clients/{client}/composite | Get effective client roles Returns the roles for the client that are associated with the client&#39;s scope.
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete) | **Delete** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/clients/{client} | Remove client-level roles from the client&#39;s scope.
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/clients/{client} | Get the roles associated with a client&#39;s scope Returns roles for the client.
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost) | **Post** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/clients/{client} | Add client-level roles to the client&#39;s scope
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings | Get all scope mappings for the client
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/realm/available | Get realm-level roles that are available to attach to this client&#39;s scope
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/realm/composite | Get effective realm-level roles associated with the client’s scope What this does is recurse any composite roles associated with the client’s scope and adds the roles to this lists.
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete) | **Delete** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/realm | Remove a set of realm-level roles from the client&#39;s scope
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/realm | Get realm-level roles associated with the client&#39;s scope
[**AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost**](ScopeMappingsAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost) | **Post** /admin/realms/{realm}/client-templates/{client-scope-id}/scope-mappings/realm | Add a set of realm-level roles to the client&#39;s scope
[**AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/clients/{client}/available | The available client-level roles Returns the roles for the client that can be associated with the client&#39;s scope
[**AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/clients/{client}/composite | Get effective client roles Returns the roles for the client that are associated with the client&#39;s scope.
[**AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/clients/{client} | Remove client-level roles from the client&#39;s scope.
[**AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/clients/{client} | Get the roles associated with a client&#39;s scope Returns roles for the client.
[**AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/clients/{client} | Add client-level roles to the client&#39;s scope
[**AdminRealmsRealmClientsClientUuidScopeMappingsGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings | Get all scope mappings for the client
[**AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/realm/available | Get realm-level roles that are available to attach to this client&#39;s scope
[**AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/realm/composite | Get effective realm-level roles associated with the client’s scope What this does is recurse any composite roles associated with the client’s scope and adds the roles to this lists.
[**AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/realm | Remove a set of realm-level roles from the client&#39;s scope
[**AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/realm | Get realm-level roles associated with the client&#39;s scope
[**AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost**](ScopeMappingsAPI.md#AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/scope-mappings/realm | Add a set of realm-level roles to the client&#39;s scope



## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet

> []RoleRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet(ctx, realm, clientScopeId, client).Execute()

The available client-level roles Returns the roles for the client that can be associated with the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet(context.Background(), realm, clientScopeId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet

> []RoleRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet(ctx, realm, clientScopeId, client).BriefRepresentation(briefRepresentation).Execute()

Get effective client roles Returns the roles for the client that are associated with the client's scope.

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet(context.Background(), realm, clientScopeId, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete

> AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete(ctx, realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()

Remove client-level roles from the client's scope.

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete(context.Background(), realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet

> []RoleRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet(ctx, realm, clientScopeId, client).Execute()

Get the roles associated with a client's scope Returns roles for the client.

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet(context.Background(), realm, clientScopeId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost

> AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost(ctx, realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost(context.Background(), realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsClientsClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet

> MappingsRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet(ctx, realm, clientScopeId).Execute()

Get all scope mappings for the client

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappingsRepresentation**](MappingsRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet

> []RoleRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet(ctx, realm, clientScopeId).Execute()

Get realm-level roles that are available to attach to this client's scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet

> []RoleRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet(ctx, realm, clientScopeId).BriefRepresentation(briefRepresentation).Execute()

Get effective realm-level roles associated with the client’s scope What this does is recurse any composite roles associated with the client’s scope and adds the roles to this lists.



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
	clientScopeId := "clientScopeId_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet(context.Background(), realm, clientScopeId).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete

> AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete(ctx, realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()

Remove a set of realm-level roles from the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete(context.Background(), realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet

> []RoleRepresentation AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet(ctx, realm, clientScopeId).Execute()

Get realm-level roles associated with the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost

> AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost(ctx, realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()

Add a set of realm-level roles to the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost(context.Background(), realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdScopeMappingsRealmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet

> []RoleRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet(ctx, realm, clientScopeId, client).Execute()

The available client-level roles Returns the roles for the client that can be associated with the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet(context.Background(), realm, clientScopeId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet

> []RoleRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet(ctx, realm, clientScopeId, client).BriefRepresentation(briefRepresentation).Execute()

Get effective client roles Returns the roles for the client that are associated with the client's scope.

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet(context.Background(), realm, clientScopeId, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete

> AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete(ctx, realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()

Remove client-level roles from the client's scope.

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete(context.Background(), realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet

> []RoleRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet(ctx, realm, clientScopeId, client).Execute()

Get the roles associated with a client's scope Returns roles for the client.

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet(context.Background(), realm, clientScopeId, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost

> AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost(ctx, realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost(context.Background(), realm, clientScopeId, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsClientsClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet

> MappingsRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet(ctx, realm, clientScopeId).Execute()

Get all scope mappings for the client

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappingsRepresentation**](MappingsRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet

> []RoleRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet(ctx, realm, clientScopeId).Execute()

Get realm-level roles that are available to attach to this client's scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet

> []RoleRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet(ctx, realm, clientScopeId).BriefRepresentation(briefRepresentation).Execute()

Get effective realm-level roles associated with the client’s scope What this does is recurse any composite roles associated with the client’s scope and adds the roles to this lists.



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
	clientScopeId := "clientScopeId_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet(context.Background(), realm, clientScopeId).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete

> AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete(ctx, realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()

Remove a set of realm-level roles from the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete(context.Background(), realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet

> []RoleRepresentation AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet(ctx, realm, clientScopeId).Execute()

Get realm-level roles associated with the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost

> AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost(ctx, realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()

Add a set of realm-level roles to the client's scope

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
	clientScopeId := "clientScopeId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost(context.Background(), realm, clientScopeId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdScopeMappingsRealmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet(ctx, realm, clientUuid, client).Execute()

The available client-level roles Returns the roles for the client that can be associated with the client's scope

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet(context.Background(), realm, clientUuid, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsClientsClientAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet(ctx, realm, clientUuid, client).BriefRepresentation(briefRepresentation).Execute()

Get effective client roles Returns the roles for the client that are associated with the client's scope.

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	client := "client_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet(context.Background(), realm, clientUuid, client).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsClientsClientCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

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


## AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete

> AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete(ctx, realm, clientUuid, client).RoleRepresentation(roleRepresentation).Execute()

Remove client-level roles from the client's scope.

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete(context.Background(), realm, clientUuid, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDelete``: %v\n", err)
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
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsClientsClientDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet(ctx, realm, clientUuid, client).Execute()

Get the roles associated with a client's scope Returns roles for the client.

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	client := "client_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet(context.Background(), realm, clientUuid, client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsClientsClientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost

> AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost(ctx, realm, clientUuid, client).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the client's scope

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	client := "client_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost(context.Background(), realm, clientUuid, client).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPost``: %v\n", err)
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
**client** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsClientsClientPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidScopeMappingsGet

> MappingsRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsGet(ctx, realm, clientUuid).Execute()

Get all scope mappings for the client

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsGet`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MappingsRepresentation**](MappingsRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet(ctx, realm, clientUuid).Execute()

Get realm-level roles that are available to attach to this client's scope

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsRealmAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet(ctx, realm, clientUuid).BriefRepresentation(briefRepresentation).Execute()

Get effective realm-level roles associated with the client’s scope What this does is recurse any composite roles associated with the client’s scope and adds the roles to this lists.



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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet(context.Background(), realm, clientUuid).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsRealmCompositeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | if false, return roles with their attributes | [default to true]

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


## AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete

> AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete(ctx, realm, clientUuid).RoleRepresentation(roleRepresentation).Execute()

Remove a set of realm-level roles from the client's scope

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete(context.Background(), realm, clientUuid).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsRealmDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet(ctx, realm, clientUuid).Execute()

Get realm-level roles associated with the client's scope

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet(context.Background(), realm, clientUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsRealmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost

> AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost(ctx, realm, clientUuid).RoleRepresentation(roleRepresentation).Execute()

Add a set of realm-level roles to the client's scope

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
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost(context.Background(), realm, clientUuid).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScopeMappingsAPI.AdminRealmsRealmClientsClientUuidScopeMappingsRealmPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidScopeMappingsRealmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**[]RoleRepresentation**](RoleRepresentation.md) |  | 

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

