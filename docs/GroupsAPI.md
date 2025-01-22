# \GroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmGroupsCountGet**](GroupsAPI.md#AdminRealmsRealmGroupsCountGet) | **Get** /admin/realms/{realm}/groups/count | Returns the groups counts.
[**AdminRealmsRealmGroupsGet**](GroupsAPI.md#AdminRealmsRealmGroupsGet) | **Get** /admin/realms/{realm}/groups | Get group hierarchy.  Only name and ids are returned.
[**AdminRealmsRealmGroupsGroupIdChildrenGet**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdChildrenGet) | **Get** /admin/realms/{realm}/groups/{group-id}/children | Return a paginated list of subgroups that have a parent group corresponding to the group on the URL
[**AdminRealmsRealmGroupsGroupIdChildrenPost**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdChildrenPost) | **Post** /admin/realms/{realm}/groups/{group-id}/children | Set or create child.
[**AdminRealmsRealmGroupsGroupIdDelete**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdDelete) | **Delete** /admin/realms/{realm}/groups/{group-id} | 
[**AdminRealmsRealmGroupsGroupIdGet**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdGet) | **Get** /admin/realms/{realm}/groups/{group-id} | 
[**AdminRealmsRealmGroupsGroupIdManagementPermissionsGet**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdManagementPermissionsGet) | **Get** /admin/realms/{realm}/groups/{group-id}/management/permissions | Return object stating whether client Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmGroupsGroupIdManagementPermissionsPut**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdManagementPermissionsPut) | **Put** /admin/realms/{realm}/groups/{group-id}/management/permissions | Return object stating whether client Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmGroupsGroupIdMembersGet**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdMembersGet) | **Get** /admin/realms/{realm}/groups/{group-id}/members | Get users Returns a stream of users, filtered according to query parameters
[**AdminRealmsRealmGroupsGroupIdPut**](GroupsAPI.md#AdminRealmsRealmGroupsGroupIdPut) | **Put** /admin/realms/{realm}/groups/{group-id} | Update group, ignores subgroups.
[**AdminRealmsRealmGroupsPost**](GroupsAPI.md#AdminRealmsRealmGroupsPost) | **Post** /admin/realms/{realm}/groups | create or add a top level realm groupSet or create child.



## AdminRealmsRealmGroupsCountGet

> map[string]int64 AdminRealmsRealmGroupsCountGet(ctx, realm).Search(search).Top(top).Execute()

Returns the groups counts.

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
	search := "search_example" // string |  (optional)
	top := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsCountGet(context.Background(), realm).Search(search).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsCountGet`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 
 **top** | **bool** |  | [default to false]

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


## AdminRealmsRealmGroupsGet

> []GroupRepresentation AdminRealmsRealmGroupsGet(ctx, realm).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).PopulateHierarchy(populateHierarchy).Q(q).Search(search).Execute()

Get group hierarchy.  Only name and ids are returned.

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
	briefRepresentation := true // bool |  (optional) (default to true)
	exact := true // bool |  (optional) (default to false)
	first := int32(56) // int32 |  (optional)
	max := int32(56) // int32 |  (optional)
	populateHierarchy := true // bool |  (optional) (default to true)
	q := "q_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGet(context.Background(), realm).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).PopulateHierarchy(populateHierarchy).Q(q).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGet`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **bool** |  | [default to true]
 **exact** | **bool** |  | [default to false]
 **first** | **int32** |  | 
 **max** | **int32** |  | 
 **populateHierarchy** | **bool** |  | [default to true]
 **q** | **string** |  | 
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


## AdminRealmsRealmGroupsGroupIdChildrenGet

> []GroupRepresentation AdminRealmsRealmGroupsGroupIdChildrenGet(ctx, realm, groupId).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).Search(search).Execute()

Return a paginated list of subgroups that have a parent group corresponding to the group on the URL

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
	briefRepresentation := true // bool | Boolean which defines whether brief groups representations are returned or not (default: false) (optional) (default to false)
	exact := true // bool | Boolean which defines whether the params \"search\" must match exactly or not (optional)
	first := int32(56) // int32 | The position of the first result to be returned (pagination offset). (optional)
	max := int32(56) // int32 | The maximum number of results that are to be returned. Defaults to 10 (optional)
	search := "search_example" // string | A String representing either an exact group name or a partial name (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdChildrenGet(context.Background(), realm, groupId).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdChildrenGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdChildrenGet`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsGroupIdChildrenGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdChildrenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | Boolean which defines whether brief groups representations are returned or not (default: false) | [default to false]
 **exact** | **bool** | Boolean which defines whether the params \&quot;search\&quot; must match exactly or not | 
 **first** | **int32** | The position of the first result to be returned (pagination offset). | 
 **max** | **int32** | The maximum number of results that are to be returned. Defaults to 10 | 
 **search** | **string** | A String representing either an exact group name or a partial name | 

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


## AdminRealmsRealmGroupsGroupIdChildrenPost

> AdminRealmsRealmGroupsGroupIdChildrenPost(ctx, realm, groupId).GroupRepresentation(groupRepresentation).Execute()

Set or create child.



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
	groupRepresentation := *openapiclient.NewGroupRepresentation() // GroupRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdChildrenPost(context.Background(), realm, groupId).GroupRepresentation(groupRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdChildrenPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdChildrenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupRepresentation** | [**GroupRepresentation**](GroupRepresentation.md) |  | 

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


## AdminRealmsRealmGroupsGroupIdDelete

> AdminRealmsRealmGroupsGroupIdDelete(ctx, realm, groupId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdDelete(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdGet

> GroupRepresentation AdminRealmsRealmGroupsGroupIdGet(ctx, realm, groupId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdGet(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdGet`: GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupRepresentation**](GroupRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGroupsGroupIdManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmGroupsGroupIdManagementPermissionsGet(ctx, realm, groupId).Execute()

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdManagementPermissionsGet(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsGroupIdManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdManagementPermissionsGetRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmGroupsGroupIdManagementPermissionsPut(ctx, realm, groupId).ManagementPermissionReference(managementPermissionReference).Execute()

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
	groupId := "groupId_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdManagementPermissionsPut(context.Background(), realm, groupId).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsGroupIdManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdManagementPermissionsPutRequest struct via the builder pattern


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


## AdminRealmsRealmGroupsGroupIdMembersGet

> []UserRepresentation AdminRealmsRealmGroupsGroupIdMembersGet(ctx, realm, groupId).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()

Get users Returns a stream of users, filtered according to query parameters

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
	briefRepresentation := true // bool | Only return basic information (only guaranteed to return id, username, created, first and last name, email, enabled state, email verification state, federation link, and access. Note that it means that namely user attributes, required actions, and not before are not returned.) (optional)
	first := int32(56) // int32 | Pagination offset (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdMembersGet(context.Background(), realm, groupId).BriefRepresentation(briefRepresentation).First(first).Max(max).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupsGroupIdMembersGet`: []UserRepresentation
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AdminRealmsRealmGroupsGroupIdMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **briefRepresentation** | **bool** | Only return basic information (only guaranteed to return id, username, created, first and last name, email, enabled state, email verification state, federation link, and access. Note that it means that namely user attributes, required actions, and not before are not returned.) | 
 **first** | **int32** | Pagination offset | 
 **max** | **int32** | Maximum results size (defaults to 100) | 

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


## AdminRealmsRealmGroupsGroupIdPut

> AdminRealmsRealmGroupsGroupIdPut(ctx, realm, groupId).GroupRepresentation(groupRepresentation).Execute()

Update group, ignores subgroups.

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
	groupRepresentation := *openapiclient.NewGroupRepresentation() // GroupRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsGroupIdPut(context.Background(), realm, groupId).GroupRepresentation(groupRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsGroupIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupRepresentation** | [**GroupRepresentation**](GroupRepresentation.md) |  | 

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


## AdminRealmsRealmGroupsPost

> AdminRealmsRealmGroupsPost(ctx, realm).GroupRepresentation(groupRepresentation).Execute()

create or add a top level realm groupSet or create child.



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
	groupRepresentation := *openapiclient.NewGroupRepresentation() // GroupRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.AdminRealmsRealmGroupsPost(context.Background(), realm).GroupRepresentation(groupRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AdminRealmsRealmGroupsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupRepresentation** | [**GroupRepresentation**](GroupRepresentation.md) |  | 

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

