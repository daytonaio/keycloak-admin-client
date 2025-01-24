# \RolesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientsClientUuidRolesGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles | Get all roles for the realm or client
[**AdminRealmsRealmClientsClientUuidRolesPost**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/roles | Create a new role for the realm or client
[**AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/composites/clients/{client-in-composite-uuid} | Get client-level roles for the client that are in the role&#39;s composite
[**AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDelete**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/composites | Remove roles from the role&#39;s composite
[**AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/composites | Get composites of the role
[**AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPost**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/composites | Add a composite to the role
[**AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/composites/realm | Get realm-level roles of the role&#39;s composite
[**AdminRealmsRealmClientsClientUuidRolesRoleNameDelete**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameDelete) | **Delete** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name} | Delete a role by name
[**AdminRealmsRealmClientsClientUuidRolesRoleNameGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name} | Get a role by name
[**AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/groups | Returns a stream of groups that have the specified role name
[**AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/management/permissions | Return object stating whether role Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut) | **Put** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/management/permissions | Return object stating whether role Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmClientsClientUuidRolesRoleNamePut**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNamePut) | **Put** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name} | Update a role by name
[**AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet**](RolesAPI.md#AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/roles/{role-name}/users | Returns a stream of users that have the specified role name.
[**AdminRealmsRealmRolesGet**](RolesAPI.md#AdminRealmsRealmRolesGet) | **Get** /admin/realms/{realm}/roles | Get all roles for the realm or client
[**AdminRealmsRealmRolesPost**](RolesAPI.md#AdminRealmsRealmRolesPost) | **Post** /admin/realms/{realm}/roles | Create a new role for the realm or client
[**AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet) | **Get** /admin/realms/{realm}/roles/{role-name}/composites/clients/{client-uuid} | Get client-level roles for the client that are in the role&#39;s composite
[**AdminRealmsRealmRolesRoleNameCompositesDelete**](RolesAPI.md#AdminRealmsRealmRolesRoleNameCompositesDelete) | **Delete** /admin/realms/{realm}/roles/{role-name}/composites | Remove roles from the role&#39;s composite
[**AdminRealmsRealmRolesRoleNameCompositesGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameCompositesGet) | **Get** /admin/realms/{realm}/roles/{role-name}/composites | Get composites of the role
[**AdminRealmsRealmRolesRoleNameCompositesPost**](RolesAPI.md#AdminRealmsRealmRolesRoleNameCompositesPost) | **Post** /admin/realms/{realm}/roles/{role-name}/composites | Add a composite to the role
[**AdminRealmsRealmRolesRoleNameCompositesRealmGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameCompositesRealmGet) | **Get** /admin/realms/{realm}/roles/{role-name}/composites/realm | Get realm-level roles of the role&#39;s composite
[**AdminRealmsRealmRolesRoleNameDelete**](RolesAPI.md#AdminRealmsRealmRolesRoleNameDelete) | **Delete** /admin/realms/{realm}/roles/{role-name} | Delete a role by name
[**AdminRealmsRealmRolesRoleNameGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameGet) | **Get** /admin/realms/{realm}/roles/{role-name} | Get a role by name
[**AdminRealmsRealmRolesRoleNameGroupsGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameGroupsGet) | **Get** /admin/realms/{realm}/roles/{role-name}/groups | Returns a stream of groups that have the specified role name
[**AdminRealmsRealmRolesRoleNameManagementPermissionsGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameManagementPermissionsGet) | **Get** /admin/realms/{realm}/roles/{role-name}/management/permissions | Return object stating whether role Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmRolesRoleNameManagementPermissionsPut**](RolesAPI.md#AdminRealmsRealmRolesRoleNameManagementPermissionsPut) | **Put** /admin/realms/{realm}/roles/{role-name}/management/permissions | Return object stating whether role Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmRolesRoleNamePut**](RolesAPI.md#AdminRealmsRealmRolesRoleNamePut) | **Put** /admin/realms/{realm}/roles/{role-name} | Update a role by name
[**AdminRealmsRealmRolesRoleNameUsersGet**](RolesAPI.md#AdminRealmsRealmRolesRoleNameUsersGet) | **Get** /admin/realms/{realm}/roles/{role-name}/users | Returns a stream of users that have the specified role name.



## AdminRealmsRealmClientsClientUuidRolesGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidRolesGet(ctx, realm, clientUuid).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()

Get all roles for the realm or client

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
	briefRepresentation := true // bool |  (optional) (default to true)
	first := int32(56) // int32 |  (optional)
	max := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesGet(context.Background(), realm, clientUuid).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** |  | [default to true]
 **first** | **int32** |  | 
 **max** | **int32** |  | 
 **search** | **string** |  | [default to &quot;&quot;]

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


## AdminRealmsRealmClientsClientUuidRolesPost

> AdminRealmsRealmClientsClientUuidRolesPost(ctx, realm, clientUuid).RoleRepresentation(roleRepresentation).Execute()

Create a new role for the realm or client

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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesPost(context.Background(), realm, clientUuid).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet(ctx, realm, clientUuid, roleName, clientInCompositeUuid).Execute()

Get client-level roles for the client that are in the role's composite

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
	clientUuid := "clientUuid_example" // string | 
	roleName := "roleName_example" // string | role's name (not id!)
	clientInCompositeUuid := "clientInCompositeUuid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet(context.Background(), realm, clientUuid, roleName, clientInCompositeUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** |  | 
**roleName** | **string** | role&#39;s name (not id!) | 
**clientInCompositeUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameCompositesClientsClientInCompositeUuidGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDelete

> AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDelete(ctx, realm, clientUuid, roleName).RoleRepresentation(roleRepresentation).Execute()

Remove roles from the role's composite

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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDelete(context.Background(), realm, clientUuid, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDelete``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameCompositesDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet(ctx, realm, clientUuid, roleName).Execute()

Get composites of the role

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet(context.Background(), realm, clientUuid, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameCompositesGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPost

> AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPost(ctx, realm, clientUuid, roleName).RoleRepresentation(roleRepresentation).Execute()

Add a composite to the role

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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPost(context.Background(), realm, clientUuid, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPost``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameCompositesPostRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet

> []RoleRepresentation AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet(ctx, realm, clientUuid, roleName).Execute()

Get realm-level roles of the role's composite

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet(context.Background(), realm, clientUuid, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameCompositesRealmGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameDelete

> AdminRealmsRealmClientsClientUuidRolesRoleNameDelete(ctx, realm, clientUuid, roleName).Execute()

Delete a role by name

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameDelete(context.Background(), realm, clientUuid, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameDelete``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameGet

> RoleRepresentation AdminRealmsRealmClientsClientUuidRolesRoleNameGet(ctx, realm, clientUuid, roleName).Execute()

Get a role by name

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameGet(context.Background(), realm, clientUuid, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameGet`: RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet

> []GroupRepresentation AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet(ctx, realm, clientUuid, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()

Returns a stream of groups that have the specified role name

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
	roleName := "roleName_example" // string | the role name.
	briefRepresentation := true // bool | if false, return a full representation of the {@code GroupRepresentation} objects. (optional) (default to true)
	first := int32(56) // int32 | first result to return. Ignored if negative or {@code null}. (optional)
	max := int32(56) // int32 | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet(context.Background(), realm, clientUuid, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | if false, return a full representation of the {@code GroupRepresentation} objects. | [default to true]
 **first** | **int32** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **int32** | maximum number of results to return. Ignored if negative or {@code null}. | 

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


## AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet(ctx, realm, clientUuid, roleName).Execute()

Return object stating whether role Authorization permissions have been initialized or not and a reference

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
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet(context.Background(), realm, clientUuid, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut(ctx, realm, clientUuid, roleName).ManagementPermissionReference(managementPermissionReference).Execute()

Return object stating whether role Authorization permissions have been initialized or not and a reference

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
	roleName := "roleName_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut(context.Background(), realm, clientUuid, roleName).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameManagementPermissionsPutRequest struct via the builder pattern


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


## AdminRealmsRealmClientsClientUuidRolesRoleNamePut

> AdminRealmsRealmClientsClientUuidRolesRoleNamePut(ctx, realm, clientUuid, roleName).RoleRepresentation(roleRepresentation).Execute()

Update a role by name

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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNamePut(context.Background(), realm, clientUuid, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNamePut``: %v\n", err)
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
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet

> []UserRepresentation AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet(ctx, realm, clientUuid, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()

Returns a stream of users that have the specified role name.

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
	roleName := "roleName_example" // string | the role name.
	briefRepresentation := true // bool | Boolean which defines whether brief representations are returned (default: false) (optional)
	first := int32(56) // int32 | first result to return. Ignored if negative or {@code null}. (optional)
	max := int32(56) // int32 | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet(context.Background(), realm, clientUuid, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmClientsClientUuidRolesRoleNameUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidRolesRoleNameUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **briefRepresentation** | **bool** | Boolean which defines whether brief representations are returned (default: false) | 
 **first** | **int32** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **int32** | maximum number of results to return. Ignored if negative or {@code null}. | 

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


## AdminRealmsRealmRolesGet

> []RoleRepresentation AdminRealmsRealmRolesGet(ctx, realm).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()

Get all roles for the realm or client

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
	briefRepresentation := true // bool |  (optional) (default to true)
	first := int32(56) // int32 |  (optional)
	max := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesGet(context.Background(), realm).BriefRepresentation(briefRepresentation).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **bool** |  | [default to true]
 **first** | **int32** |  | 
 **max** | **int32** |  | 
 **search** | **string** |  | [default to &quot;&quot;]

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


## AdminRealmsRealmRolesPost

> AdminRealmsRealmRolesPost(ctx, realm).RoleRepresentation(roleRepresentation).Execute()

Create a new role for the realm or client

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
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmRolesPost(context.Background(), realm).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet

> []RoleRepresentation AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet(ctx, realm, clientUuid, roleName).Execute()

Get client-level roles for the client that are in the role's composite

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
	clientUuid := "clientUuid_example" // string | 
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet(context.Background(), realm, clientUuid, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** |  | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameCompositesClientsClientUuidGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameCompositesDelete

> AdminRealmsRealmRolesRoleNameCompositesDelete(ctx, realm, roleName).RoleRepresentation(roleRepresentation).Execute()

Remove roles from the role's composite

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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameCompositesDelete(context.Background(), realm, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameCompositesDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameCompositesGet

> []RoleRepresentation AdminRealmsRealmRolesRoleNameCompositesGet(ctx, realm, roleName).Execute()

Get composites of the role

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameCompositesGet(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameCompositesGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameCompositesGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameCompositesPost

> AdminRealmsRealmRolesRoleNameCompositesPost(ctx, realm, roleName).RoleRepresentation(roleRepresentation).Execute()

Add a composite to the role

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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameCompositesPost(context.Background(), realm, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameCompositesPostRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameCompositesRealmGet

> []RoleRepresentation AdminRealmsRealmRolesRoleNameCompositesRealmGet(ctx, realm, roleName).Execute()

Get realm-level roles of the role's composite

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameCompositesRealmGet(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameCompositesRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameCompositesRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameCompositesRealmGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameDelete

> AdminRealmsRealmRolesRoleNameDelete(ctx, realm, roleName).Execute()

Delete a role by name

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameDelete(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameGet

> RoleRepresentation AdminRealmsRealmRolesRoleNameGet(ctx, realm, roleName).Execute()

Get a role by name

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
	roleName := "roleName_example" // string | role's name (not id!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameGet(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameGet`: RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleRepresentation**](RoleRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmRolesRoleNameGroupsGet

> []GroupRepresentation AdminRealmsRealmRolesRoleNameGroupsGet(ctx, realm, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()

Returns a stream of groups that have the specified role name

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
	roleName := "roleName_example" // string | the role name.
	briefRepresentation := true // bool | if false, return a full representation of the {@code GroupRepresentation} objects. (optional) (default to true)
	first := int32(56) // int32 | first result to return. Ignored if negative or {@code null}. (optional)
	max := int32(56) // int32 | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameGroupsGet(context.Background(), realm, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameGroupsGet`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | if false, return a full representation of the {@code GroupRepresentation} objects. | [default to true]
 **first** | **int32** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **int32** | maximum number of results to return. Ignored if negative or {@code null}. | 

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


## AdminRealmsRealmRolesRoleNameManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmRolesRoleNameManagementPermissionsGet(ctx, realm, roleName).Execute()

Return object stating whether role Authorization permissions have been initialized or not and a reference

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
	roleName := "roleName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameManagementPermissionsGet(context.Background(), realm, roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameManagementPermissionsGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNameManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmRolesRoleNameManagementPermissionsPut(ctx, realm, roleName).ManagementPermissionReference(managementPermissionReference).Execute()

Return object stating whether role Authorization permissions have been initialized or not and a reference

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
	roleName := "roleName_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameManagementPermissionsPut(context.Background(), realm, roleName).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameManagementPermissionsPutRequest struct via the builder pattern


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


## AdminRealmsRealmRolesRoleNamePut

> AdminRealmsRealmRolesRoleNamePut(ctx, realm, roleName).RoleRepresentation(roleRepresentation).Execute()

Update a role by name

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
	roleName := "roleName_example" // string | role's name (not id!)
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNamePut(context.Background(), realm, roleName).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | role&#39;s name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleRepresentation** | [**RoleRepresentation**](RoleRepresentation.md) |  | 

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


## AdminRealmsRealmRolesRoleNameUsersGet

> []UserRepresentation AdminRealmsRealmRolesRoleNameUsersGet(ctx, realm, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()

Returns a stream of users that have the specified role name.

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
	roleName := "roleName_example" // string | the role name.
	briefRepresentation := true // bool | Boolean which defines whether brief representations are returned (default: false) (optional)
	first := int32(56) // int32 | first result to return. Ignored if negative or {@code null}. (optional)
	max := int32(56) // int32 | maximum number of results to return. Ignored if negative or {@code null}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesAPI.AdminRealmsRealmRolesRoleNameUsersGet(context.Background(), realm, roleName).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.AdminRealmsRealmRolesRoleNameUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesRoleNameUsersGet`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesAPI.AdminRealmsRealmRolesRoleNameUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleName** | **string** | the role name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesRoleNameUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | Boolean which defines whether brief representations are returned (default: false) | 
 **first** | **int32** | first result to return. Ignored if negative or {@code null}. | 
 **max** | **int32** | maximum number of results to return. Ignored if negative or {@code null}. | 

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

