# \RolesByIDAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet) | **Get** /admin/realms/{realm}/roles-by-id/{role-id}/composites/clients/{clientUuid} | Get client-level roles for the client that are in the role&#39;s composite
[**AdminRealmsRealmRolesByIdRoleIdCompositesDelete**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdCompositesDelete) | **Delete** /admin/realms/{realm}/roles-by-id/{role-id}/composites | Remove a set of roles from the role&#39;s composite
[**AdminRealmsRealmRolesByIdRoleIdCompositesGet**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdCompositesGet) | **Get** /admin/realms/{realm}/roles-by-id/{role-id}/composites | Get role&#39;s children Returns a set of role&#39;s children provided the role is a composite.
[**AdminRealmsRealmRolesByIdRoleIdCompositesPost**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdCompositesPost) | **Post** /admin/realms/{realm}/roles-by-id/{role-id}/composites | Make the role a composite role by associating some child roles
[**AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet) | **Get** /admin/realms/{realm}/roles-by-id/{role-id}/composites/realm | Get realm-level roles that are in the role&#39;s composite
[**AdminRealmsRealmRolesByIdRoleIdDelete**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdDelete) | **Delete** /admin/realms/{realm}/roles-by-id/{role-id} | Delete the role
[**AdminRealmsRealmRolesByIdRoleIdGet**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdGet) | **Get** /admin/realms/{realm}/roles-by-id/{role-id} | Get a specific role&#39;s representation
[**AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet) | **Get** /admin/realms/{realm}/roles-by-id/{role-id}/management/permissions | Return object stating whether role Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut) | **Put** /admin/realms/{realm}/roles-by-id/{role-id}/management/permissions | Return object stating whether role Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmRolesByIdRoleIdPut**](RolesByIDAPI.md#AdminRealmsRealmRolesByIdRoleIdPut) | **Put** /admin/realms/{realm}/roles-by-id/{role-id} | Update the role



## AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet

> []RoleRepresentation AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet(ctx, realm, clientUuid, roleId).Execute()

Get client-level roles for the client that are in the role's composite

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
	clientUuid := "clientUuid_example" // string | 
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet(context.Background(), realm, clientUuid, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** |  | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdCompositesClientsClientUuidGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdCompositesDelete

> AdminRealmsRealmRolesByIdRoleIdCompositesDelete(ctx, realm, roleId).RoleRepresentation(roleRepresentation).Execute()

Remove a set of roles from the role's composite

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
	roleId := "roleId_example" // string | Role id
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesDelete(context.Background(), realm, roleId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | Role id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdCompositesDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdCompositesGet

> []RoleRepresentation AdminRealmsRealmRolesByIdRoleIdCompositesGet(ctx, realm, roleId).First(first).Max(max).Search(search).Execute()

Get role's children Returns a set of role's children provided the role is a composite.

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
	roleId := "roleId_example" // string | 
	first := int32(56) // int32 |  (optional)
	max := int32(56) // int32 |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesGet(context.Background(), realm, roleId).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesByIdRoleIdCompositesGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdCompositesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **first** | **int32** |  | 
 **max** | **int32** |  | 
 **search** | **string** |  | 

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


## AdminRealmsRealmRolesByIdRoleIdCompositesPost

> AdminRealmsRealmRolesByIdRoleIdCompositesPost(ctx, realm, roleId).RoleRepresentation(roleRepresentation).Execute()

Make the role a composite role by associating some child roles

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
	roleId := "roleId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesPost(context.Background(), realm, roleId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdCompositesPostRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet

> []RoleRepresentation AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet(ctx, realm, roleId).Execute()

Get realm-level roles that are in the role's composite

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
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdCompositesRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdCompositesRealmGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdDelete

> AdminRealmsRealmRolesByIdRoleIdDelete(ctx, realm, roleId).Execute()

Delete the role

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
	roleId := "roleId_example" // string | id of role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdDelete(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | id of role | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdGet

> RoleRepresentation AdminRealmsRealmRolesByIdRoleIdGet(ctx, realm, roleId).Execute()

Get a specific role's representation

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
	roleId := "roleId_example" // string | id of role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdGet(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesByIdRoleIdGet`: RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | id of role | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet(ctx, realm, roleId).Execute()

Return object stating whether role Authorization permissions have been initialized or not and a reference

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
	roleId := "roleId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet(context.Background(), realm, roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdManagementPermissionsGetRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut(ctx, realm, roleId).ManagementPermissionReference(managementPermissionReference).Execute()

Return object stating whether role Authorization permissions have been initialized or not and a reference

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
	roleId := "roleId_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut(context.Background(), realm, roleId).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdManagementPermissionsPutRequest struct via the builder pattern


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


## AdminRealmsRealmRolesByIdRoleIdPut

> AdminRealmsRealmRolesByIdRoleIdPut(ctx, realm, roleId).RoleRepresentation(roleRepresentation).Execute()

Update the role

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
	roleId := "roleId_example" // string | id of role
	roleRepresentation := *openapiclient.NewRoleRepresentation() // RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdPut(context.Background(), realm, roleId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RolesByIDAPI.AdminRealmsRealmRolesByIdRoleIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**roleId** | **string** | id of role | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmRolesByIdRoleIdPutRequest struct via the builder pattern


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

