# \ClientRoleMappingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client-id}/available | Get available client-level roles that can be mapped to the user or group
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client-id}/composite | Get effective client-level role mappings This recurses any composite roles
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete) | **Delete** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client-id} | Delete client-level roles from user or group role mapping
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client-id} | Get client-level role mappings for the user or group, and the app
[**AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost**](ClientRoleMappingsAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost) | **Post** /admin/realms/{realm}/groups/{group-id}/role-mappings/clients/{client-id} | Add client-level roles to the user or group role mapping
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client-id}/available | Get available client-level roles that can be mapped to the user or group
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client-id}/composite | Get effective client-level role mappings This recurses any composite roles
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client-id} | Delete client-level roles from user or group role mapping
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client-id} | Get client-level role mappings for the user or group, and the app
[**AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost**](ClientRoleMappingsAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost) | **Post** /admin/realms/{realm}/users/{user-id}/role-mappings/clients/{client-id} | Add client-level roles to the user or group role mapping



## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet(ctx, realm, groupId, clientId).Execute()

Get available client-level roles that can be mapped to the user or group

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
	groupId := "groupId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet(context.Background(), realm, groupId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdAvailableGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet(ctx, realm, groupId, clientId).BriefRepresentation(briefRepresentation).Execute()

Get effective client-level role mappings This recurses any composite roles

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
	groupId := "groupId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet(context.Background(), realm, groupId, clientId).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdCompositeGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete

> AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete(ctx, realm, groupId, clientId).RoleRepresentation(roleRepresentation).Execute()

Delete client-level roles from user or group role mapping

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
	groupId := "groupId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete(context.Background(), realm, groupId, clientId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet(ctx, realm, groupId, clientId).Execute()

Get client-level role mappings for the user or group, and the app

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
	groupId := "groupId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet(context.Background(), realm, groupId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost

> AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost(ctx, realm, groupId, clientId).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the user or group role mapping

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
	groupId := "groupId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost(context.Background(), realm, groupId, clientId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsClientsClientIdPostRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet(ctx, realm, userId, clientId).Execute()

Get available client-level roles that can be mapped to the user or group

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
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet(context.Background(), realm, userId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdAvailableGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet(ctx, realm, userId, clientId).BriefRepresentation(briefRepresentation).Execute()

Get effective client-level role mappings This recurses any composite roles

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
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet(context.Background(), realm, userId, clientId).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdCompositeGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete

> AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete(ctx, realm, userId, clientId).RoleRepresentation(roleRepresentation).Execute()

Delete client-level roles from user or group role mapping

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
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete(context.Background(), realm, userId, clientId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDelete``: %v\n", err)
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
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet(ctx, realm, userId, clientId).Execute()

Get client-level role mappings for the user or group, and the app

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
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet(context.Background(), realm, userId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost

> AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost(ctx, realm, userId, clientId).RoleRepresentation(roleRepresentation).Execute()

Add client-level roles to the user or group role mapping

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
	userId := "userId_example" // string | 
	clientId := "clientId_example" // string | client id (not clientId!)
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost(context.Background(), realm, userId, clientId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRoleMappingsAPI.AdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPost``: %v\n", err)
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
**clientId** | **string** | client id (not clientId!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsClientsClientIdPostRequest struct via the builder pattern


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

