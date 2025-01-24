# \RoleMapperAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmGroupsGroupIdRoleMappingsGet**](RoleMapperAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings | Get role mappings
[**AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet**](RoleMapperAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/realm/available | Get realm-level roles that can be mapped
[**AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet**](RoleMapperAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/realm/composite | Get effective realm-level role mappings This will recurse all composite roles to get the result.
[**AdminRealmsRealmGroupsGroupIdRoleMappingsRealmDelete**](RoleMapperAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsRealmDelete) | **Delete** /admin/realms/{realm}/groups/{group-id}/role-mappings/realm | Delete realm-level role mappings
[**AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet**](RoleMapperAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet) | **Get** /admin/realms/{realm}/groups/{group-id}/role-mappings/realm | Get realm-level role mappings
[**AdminRealmsRealmGroupsGroupIdRoleMappingsRealmPost**](RoleMapperAPI.md#AdminRealmsRealmGroupsGroupIdRoleMappingsRealmPost) | **Post** /admin/realms/{realm}/groups/{group-id}/role-mappings/realm | Add realm-level role mappings to the user
[**AdminRealmsRealmUsersUserIdRoleMappingsGet**](RoleMapperAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings | Get role mappings
[**AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet**](RoleMapperAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/realm/available | Get realm-level roles that can be mapped
[**AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet**](RoleMapperAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/realm/composite | Get effective realm-level role mappings This will recurse all composite roles to get the result.
[**AdminRealmsRealmUsersUserIdRoleMappingsRealmDelete**](RoleMapperAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsRealmDelete) | **Delete** /admin/realms/{realm}/users/{user-id}/role-mappings/realm | Delete realm-level role mappings
[**AdminRealmsRealmUsersUserIdRoleMappingsRealmGet**](RoleMapperAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsRealmGet) | **Get** /admin/realms/{realm}/users/{user-id}/role-mappings/realm | Get realm-level role mappings
[**AdminRealmsRealmUsersUserIdRoleMappingsRealmPost**](RoleMapperAPI.md#AdminRealmsRealmUsersUserIdRoleMappingsRealmPost) | **Post** /admin/realms/{realm}/users/{user-id}/role-mappings/realm | Add realm-level role mappings to the user



## AdminRealmsRealmGroupsGroupIdRoleMappingsGet

> MappingsRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsGet(ctx, realm, groupId).Execute()

Get role mappings

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsGet(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsGet`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet(ctx, realm, groupId).Execute()

Get realm-level roles that can be mapped

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsRealmAvailableGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet(ctx, realm, groupId).BriefRepresentation(briefRepresentation).Execute()

Get effective realm-level role mappings This will recurse all composite roles to get the result.

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
	groupId := "groupId_example" // string | 
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet(context.Background(), realm, groupId).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsRealmCompositeGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsRealmDelete

> AdminRealmsRealmGroupsGroupIdRoleMappingsRealmDelete(ctx, realm, groupId).RoleRepresentation(roleRepresentation).Execute()

Delete realm-level role mappings

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
	groupId := "groupId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmDelete(context.Background(), realm, groupId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmDelete``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsRealmDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet

> []RoleRepresentation AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet(ctx, realm, groupId).Execute()

Get realm-level role mappings

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsRealmGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdRoleMappingsRealmPost

> AdminRealmsRealmGroupsGroupIdRoleMappingsRealmPost(ctx, realm, groupId).RoleRepresentation(roleRepresentation).Execute()

Add realm-level role mappings to the user

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
	groupId := "groupId_example" // string | 
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmPost(context.Background(), realm, groupId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmGroupsGroupIdRoleMappingsRealmPost``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdRoleMappingsRealmPostRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsGet

> MappingsRepresentation AdminRealmsRealmUsersUserIdRoleMappingsGet(ctx, realm, userId).Execute()

Get role mappings

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
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsGet`: MappingsRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet(ctx, realm, userId).Execute()

Get realm-level roles that can be mapped

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
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsRealmAvailableGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet(ctx, realm, userId).BriefRepresentation(briefRepresentation).Execute()

Get effective realm-level role mappings This will recurse all composite roles to get the result.

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
	briefRepresentation := true // bool | if false, return roles with their attributes (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet(context.Background(), realm, userId).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsRealmCompositeGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsRealmDelete

> AdminRealmsRealmUsersUserIdRoleMappingsRealmDelete(ctx, realm, userId).RoleRepresentation(roleRepresentation).Execute()

Delete realm-level role mappings

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
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmDelete(context.Background(), realm, userId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsRealmDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsRealmGet

> []RoleRepresentation AdminRealmsRealmUsersUserIdRoleMappingsRealmGet(ctx, realm, userId).Execute()

Get realm-level role mappings

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
	resp, r, err := apiClient.RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersUserIdRoleMappingsRealmGet`: []RoleRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsRealmGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersUserIdRoleMappingsRealmPost

> AdminRealmsRealmUsersUserIdRoleMappingsRealmPost(ctx, realm, userId).RoleRepresentation(roleRepresentation).Execute()

Add realm-level role mappings to the user

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
	roleRepresentation := []openapiclient.RoleRepresentation{*openapiclient.NewRoleRepresentation()} // []RoleRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmPost(context.Background(), realm, userId).RoleRepresentation(roleRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoleMapperAPI.AdminRealmsRealmUsersUserIdRoleMappingsRealmPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersUserIdRoleMappingsRealmPostRequest struct via the builder pattern


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

