# \RealmsAdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsGet**](RealmsAdminAPI.md#AdminRealmsGet) | **Get** /admin/realms | Get accessible realms Returns a list of accessible realms. The list is filtered based on what realms the caller is allowed to view.
[**AdminRealmsPost**](RealmsAdminAPI.md#AdminRealmsPost) | **Post** /admin/realms | Import a realm. Imports a realm from a full representation of that realm.
[**AdminRealmsRealmAdminEventsDelete**](RealmsAdminAPI.md#AdminRealmsRealmAdminEventsDelete) | **Delete** /admin/realms/{realm}/admin-events | Delete all admin events
[**AdminRealmsRealmAdminEventsGet**](RealmsAdminAPI.md#AdminRealmsRealmAdminEventsGet) | **Get** /admin/realms/{realm}/admin-events | Get admin events Returns all admin events, or filters events based on URL query parameters listed here
[**AdminRealmsRealmClientDescriptionConverterPost**](RealmsAdminAPI.md#AdminRealmsRealmClientDescriptionConverterPost) | **Post** /admin/realms/{realm}/client-description-converter | Base path for importing clients under this realm.
[**AdminRealmsRealmClientPoliciesPoliciesGet**](RealmsAdminAPI.md#AdminRealmsRealmClientPoliciesPoliciesGet) | **Get** /admin/realms/{realm}/client-policies/policies | 
[**AdminRealmsRealmClientPoliciesPoliciesPut**](RealmsAdminAPI.md#AdminRealmsRealmClientPoliciesPoliciesPut) | **Put** /admin/realms/{realm}/client-policies/policies | 
[**AdminRealmsRealmClientPoliciesProfilesGet**](RealmsAdminAPI.md#AdminRealmsRealmClientPoliciesProfilesGet) | **Get** /admin/realms/{realm}/client-policies/profiles | 
[**AdminRealmsRealmClientPoliciesProfilesPut**](RealmsAdminAPI.md#AdminRealmsRealmClientPoliciesProfilesPut) | **Put** /admin/realms/{realm}/client-policies/profiles | 
[**AdminRealmsRealmClientSessionStatsGet**](RealmsAdminAPI.md#AdminRealmsRealmClientSessionStatsGet) | **Get** /admin/realms/{realm}/client-session-stats | Get client session stats Returns a JSON map.
[**AdminRealmsRealmClientTypesGet**](RealmsAdminAPI.md#AdminRealmsRealmClientTypesGet) | **Get** /admin/realms/{realm}/client-types | List all client types available in the current realm
[**AdminRealmsRealmClientTypesPut**](RealmsAdminAPI.md#AdminRealmsRealmClientTypesPut) | **Put** /admin/realms/{realm}/client-types | Update a client type
[**AdminRealmsRealmCredentialRegistratorsGet**](RealmsAdminAPI.md#AdminRealmsRealmCredentialRegistratorsGet) | **Get** /admin/realms/{realm}/credential-registrators | 
[**AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete**](RealmsAdminAPI.md#AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete) | **Delete** /admin/realms/{realm}/default-default-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut**](RealmsAdminAPI.md#AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut) | **Put** /admin/realms/{realm}/default-default-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmDefaultDefaultClientScopesGet**](RealmsAdminAPI.md#AdminRealmsRealmDefaultDefaultClientScopesGet) | **Get** /admin/realms/{realm}/default-default-client-scopes | Get realm default client scopes.  Only name and ids are returned.
[**AdminRealmsRealmDefaultGroupsGet**](RealmsAdminAPI.md#AdminRealmsRealmDefaultGroupsGet) | **Get** /admin/realms/{realm}/default-groups | Get group hierarchy.  Only name and ids are returned.
[**AdminRealmsRealmDefaultGroupsGroupIdDelete**](RealmsAdminAPI.md#AdminRealmsRealmDefaultGroupsGroupIdDelete) | **Delete** /admin/realms/{realm}/default-groups/{groupId} | 
[**AdminRealmsRealmDefaultGroupsGroupIdPut**](RealmsAdminAPI.md#AdminRealmsRealmDefaultGroupsGroupIdPut) | **Put** /admin/realms/{realm}/default-groups/{groupId} | 
[**AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete**](RealmsAdminAPI.md#AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete) | **Delete** /admin/realms/{realm}/default-optional-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut**](RealmsAdminAPI.md#AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut) | **Put** /admin/realms/{realm}/default-optional-client-scopes/{clientScopeId} | 
[**AdminRealmsRealmDefaultOptionalClientScopesGet**](RealmsAdminAPI.md#AdminRealmsRealmDefaultOptionalClientScopesGet) | **Get** /admin/realms/{realm}/default-optional-client-scopes | Get realm optional client scopes.  Only name and ids are returned.
[**AdminRealmsRealmDelete**](RealmsAdminAPI.md#AdminRealmsRealmDelete) | **Delete** /admin/realms/{realm} | Delete the realm
[**AdminRealmsRealmEventsConfigGet**](RealmsAdminAPI.md#AdminRealmsRealmEventsConfigGet) | **Get** /admin/realms/{realm}/events/config | Get the events provider configuration Returns JSON object with events provider configuration
[**AdminRealmsRealmEventsConfigPut**](RealmsAdminAPI.md#AdminRealmsRealmEventsConfigPut) | **Put** /admin/realms/{realm}/events/config | 
[**AdminRealmsRealmEventsDelete**](RealmsAdminAPI.md#AdminRealmsRealmEventsDelete) | **Delete** /admin/realms/{realm}/events | Delete all events
[**AdminRealmsRealmEventsGet**](RealmsAdminAPI.md#AdminRealmsRealmEventsGet) | **Get** /admin/realms/{realm}/events | Get events Returns all events, or filters them based on URL query parameters listed here
[**AdminRealmsRealmGet**](RealmsAdminAPI.md#AdminRealmsRealmGet) | **Get** /admin/realms/{realm} | Get the top-level representation of the realm It will not include nested information like User and Client representations.
[**AdminRealmsRealmGroupByPathPathGet**](RealmsAdminAPI.md#AdminRealmsRealmGroupByPathPathGet) | **Get** /admin/realms/{realm}/group-by-path/{path} | 
[**AdminRealmsRealmLocalizationGet**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationGet) | **Get** /admin/realms/{realm}/localization | 
[**AdminRealmsRealmLocalizationLocaleDelete**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationLocaleDelete) | **Delete** /admin/realms/{realm}/localization/{locale} | 
[**AdminRealmsRealmLocalizationLocaleGet**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationLocaleGet) | **Get** /admin/realms/{realm}/localization/{locale} | 
[**AdminRealmsRealmLocalizationLocaleKeyDelete**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationLocaleKeyDelete) | **Delete** /admin/realms/{realm}/localization/{locale}/{key} | 
[**AdminRealmsRealmLocalizationLocaleKeyGet**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationLocaleKeyGet) | **Get** /admin/realms/{realm}/localization/{locale}/{key} | 
[**AdminRealmsRealmLocalizationLocaleKeyPut**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationLocaleKeyPut) | **Put** /admin/realms/{realm}/localization/{locale}/{key} | 
[**AdminRealmsRealmLocalizationLocalePost**](RealmsAdminAPI.md#AdminRealmsRealmLocalizationLocalePost) | **Post** /admin/realms/{realm}/localization/{locale} | Import localization from uploaded JSON file
[**AdminRealmsRealmLogoutAllPost**](RealmsAdminAPI.md#AdminRealmsRealmLogoutAllPost) | **Post** /admin/realms/{realm}/logout-all | Removes all user sessions.
[**AdminRealmsRealmPartialExportPost**](RealmsAdminAPI.md#AdminRealmsRealmPartialExportPost) | **Post** /admin/realms/{realm}/partial-export | Partial export of existing realm into a JSON file.
[**AdminRealmsRealmPartialImportPost**](RealmsAdminAPI.md#AdminRealmsRealmPartialImportPost) | **Post** /admin/realms/{realm}/partialImport | Partial import from a JSON file to an existing realm.
[**AdminRealmsRealmPushRevocationPost**](RealmsAdminAPI.md#AdminRealmsRealmPushRevocationPost) | **Post** /admin/realms/{realm}/push-revocation | Push the realm&#39;s revocation policy to any client that has an admin url associated with it.
[**AdminRealmsRealmPut**](RealmsAdminAPI.md#AdminRealmsRealmPut) | **Put** /admin/realms/{realm} | Update the top-level information of the realm Any user, roles or client information in the representation will be ignored.
[**AdminRealmsRealmSessionsSessionDelete**](RealmsAdminAPI.md#AdminRealmsRealmSessionsSessionDelete) | **Delete** /admin/realms/{realm}/sessions/{session} | Remove a specific user session.
[**AdminRealmsRealmTestSMTPConnectionPost**](RealmsAdminAPI.md#AdminRealmsRealmTestSMTPConnectionPost) | **Post** /admin/realms/{realm}/testSMTPConnection | Test SMTP connection with current logged in user
[**AdminRealmsRealmUsersManagementPermissionsGet**](RealmsAdminAPI.md#AdminRealmsRealmUsersManagementPermissionsGet) | **Get** /admin/realms/{realm}/users-management-permissions | 
[**AdminRealmsRealmUsersManagementPermissionsPut**](RealmsAdminAPI.md#AdminRealmsRealmUsersManagementPermissionsPut) | **Put** /admin/realms/{realm}/users-management-permissions | 



## AdminRealmsGet

> []RealmRepresentation AdminRealmsGet(ctx).BriefRepresentation(briefRepresentation).Execute()

Get accessible realms Returns a list of accessible realms. The list is filtered based on what realms the caller is allowed to view.

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
	briefRepresentation := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsGet(context.Background()).BriefRepresentation(briefRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsGet`: []RealmRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **briefRepresentation** | **bool** |  | [default to false]

### Return type

[**[]RealmRepresentation**](RealmRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsPost

> AdminRealmsPost(ctx).Body(body).Execute()

Import a realm. Imports a realm from a full representation of that realm.



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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

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


## AdminRealmsRealmAdminEventsDelete

> AdminRealmsRealmAdminEventsDelete(ctx, realm).Execute()

Delete all admin events

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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmAdminEventsDelete(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmAdminEventsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAdminEventsDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAdminEventsGet

> []AdminEventRepresentation AdminRealmsRealmAdminEventsGet(ctx, realm).AuthClient(authClient).AuthIpAddress(authIpAddress).AuthRealm(authRealm).AuthUser(authUser).DateFrom(dateFrom).DateTo(dateTo).First(first).Max(max).OperationTypes(operationTypes).ResourcePath(resourcePath).ResourceTypes(resourceTypes).Execute()

Get admin events Returns all admin events, or filters events based on URL query parameters listed here

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
	authClient := "authClient_example" // string |  (optional)
	authIpAddress := "authIpAddress_example" // string |  (optional)
	authRealm := "authRealm_example" // string |  (optional)
	authUser := "authUser_example" // string | user id (optional)
	dateFrom := "dateFrom_example" // string |  (optional)
	dateTo := "dateTo_example" // string |  (optional)
	first := int32(56) // int32 |  (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)
	operationTypes := []string{"Inner_example"} // []string |  (optional)
	resourcePath := "resourcePath_example" // string |  (optional)
	resourceTypes := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmAdminEventsGet(context.Background(), realm).AuthClient(authClient).AuthIpAddress(authIpAddress).AuthRealm(authRealm).AuthUser(authUser).DateFrom(dateFrom).DateTo(dateTo).First(first).Max(max).OperationTypes(operationTypes).ResourcePath(resourcePath).ResourceTypes(resourceTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmAdminEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAdminEventsGet`: []AdminEventRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmAdminEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAdminEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authClient** | **string** |  | 
 **authIpAddress** | **string** |  | 
 **authRealm** | **string** |  | 
 **authUser** | **string** | user id | 
 **dateFrom** | **string** |  | 
 **dateTo** | **string** |  | 
 **first** | **int32** |  | 
 **max** | **int32** | Maximum results size (defaults to 100) | 
 **operationTypes** | **[]string** |  | 
 **resourcePath** | **string** |  | 
 **resourceTypes** | **[]string** |  | 

### Return type

[**[]AdminEventRepresentation**](AdminEventRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientDescriptionConverterPost

> ClientRepresentation AdminRealmsRealmClientDescriptionConverterPost(ctx, realm).Body(body).Execute()

Base path for importing clients under this realm.

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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientDescriptionConverterPost(context.Background(), realm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientDescriptionConverterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientDescriptionConverterPost`: ClientRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmClientDescriptionConverterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientDescriptionConverterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** |  | 

### Return type

[**ClientRepresentation**](ClientRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientPoliciesPoliciesGet

> ClientPoliciesRepresentation AdminRealmsRealmClientPoliciesPoliciesGet(ctx, realm).IncludeGlobalPolicies(includeGlobalPolicies).Execute()



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
	includeGlobalPolicies := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesGet(context.Background(), realm).IncludeGlobalPolicies(includeGlobalPolicies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientPoliciesPoliciesGet`: ClientPoliciesRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientPoliciesPoliciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGlobalPolicies** | **bool** |  | 

### Return type

[**ClientPoliciesRepresentation**](ClientPoliciesRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientPoliciesPoliciesPut

> AdminRealmsRealmClientPoliciesPoliciesPut(ctx, realm).ClientPoliciesRepresentation(clientPoliciesRepresentation).Execute()



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
	clientPoliciesRepresentation := *openapiclient.NewClientPoliciesRepresentation() // ClientPoliciesRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesPut(context.Background(), realm).ClientPoliciesRepresentation(clientPoliciesRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientPoliciesPoliciesPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientPoliciesPoliciesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientPoliciesRepresentation** | [**ClientPoliciesRepresentation**](ClientPoliciesRepresentation.md) |  | 

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


## AdminRealmsRealmClientPoliciesProfilesGet

> ClientProfilesRepresentation AdminRealmsRealmClientPoliciesProfilesGet(ctx, realm).IncludeGlobalProfiles(includeGlobalProfiles).Execute()



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
	includeGlobalProfiles := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesGet(context.Background(), realm).IncludeGlobalProfiles(includeGlobalProfiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientPoliciesProfilesGet`: ClientProfilesRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientPoliciesProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGlobalProfiles** | **bool** |  | 

### Return type

[**ClientProfilesRepresentation**](ClientProfilesRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientPoliciesProfilesPut

> AdminRealmsRealmClientPoliciesProfilesPut(ctx, realm).ClientProfilesRepresentation(clientProfilesRepresentation).Execute()



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
	clientProfilesRepresentation := *openapiclient.NewClientProfilesRepresentation() // ClientProfilesRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesPut(context.Background(), realm).ClientProfilesRepresentation(clientProfilesRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientPoliciesProfilesPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientPoliciesProfilesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientProfilesRepresentation** | [**ClientProfilesRepresentation**](ClientProfilesRepresentation.md) |  | 

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


## AdminRealmsRealmClientSessionStatsGet

> []map[string]string AdminRealmsRealmClientSessionStatsGet(ctx, realm).Execute()

Get client session stats Returns a JSON map.



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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientSessionStatsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientSessionStatsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientSessionStatsGet`: []map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmClientSessionStatsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientSessionStatsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]string**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientTypesGet

> ClientTypesRepresentation AdminRealmsRealmClientTypesGet(ctx, realm).Execute()

List all client types available in the current realm



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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientTypesGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTypesGet`: ClientTypesRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmClientTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientTypesRepresentation**](ClientTypesRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientTypesPut

> AdminRealmsRealmClientTypesPut(ctx, realm).ClientTypesRepresentation(clientTypesRepresentation).Execute()

Update a client type



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
	clientTypesRepresentation := *openapiclient.NewClientTypesRepresentation() // ClientTypesRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmClientTypesPut(context.Background(), realm).ClientTypesRepresentation(clientTypesRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmClientTypesPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTypesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientTypesRepresentation** | [**ClientTypesRepresentation**](ClientTypesRepresentation.md) |  | 

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


## AdminRealmsRealmCredentialRegistratorsGet

> []string AdminRealmsRealmCredentialRegistratorsGet(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmCredentialRegistratorsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmCredentialRegistratorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmCredentialRegistratorsGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmCredentialRegistratorsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmCredentialRegistratorsGetRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete

> AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete(ctx, realm, clientScopeId).Execute()



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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultDefaultClientScopesClientScopeIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut

> AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut(ctx, realm, clientScopeId).Execute()



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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultDefaultClientScopesClientScopeIdPutRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultDefaultClientScopesGet

> []ClientScopeRepresentation AdminRealmsRealmDefaultDefaultClientScopesGet(ctx, realm).Execute()

Get realm default client scopes.  Only name and ids are returned.

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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmDefaultDefaultClientScopesGet`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmDefaultDefaultClientScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultDefaultClientScopesGetRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultGroupsGet

> []GroupRepresentation AdminRealmsRealmDefaultGroupsGet(ctx, realm).Execute()

Get group hierarchy.  Only name and ids are returned.

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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmDefaultGroupsGet`: []GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## AdminRealmsRealmDefaultGroupsGroupIdDelete

> AdminRealmsRealmDefaultGroupsGroupIdDelete(ctx, realm, groupId).Execute()



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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGroupIdDelete(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGroupIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultGroupsGroupIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultGroupsGroupIdPut

> AdminRealmsRealmDefaultGroupsGroupIdPut(ctx, realm, groupId).Execute()



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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGroupIdPut(context.Background(), realm, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultGroupsGroupIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultGroupsGroupIdPutRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete

> AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete(ctx, realm, clientScopeId).Execute()



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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultOptionalClientScopesClientScopeIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut

> AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut(ctx, realm, clientScopeId).Execute()



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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultOptionalClientScopesClientScopeIdPutRequest struct via the builder pattern


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


## AdminRealmsRealmDefaultOptionalClientScopesGet

> []ClientScopeRepresentation AdminRealmsRealmDefaultOptionalClientScopesGet(ctx, realm).Execute()

Get realm optional client scopes.  Only name and ids are returned.

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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmDefaultOptionalClientScopesGet`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmDefaultOptionalClientScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmDefaultOptionalClientScopesGetRequest struct via the builder pattern


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


## AdminRealmsRealmDelete

> AdminRealmsRealmDelete(ctx, realm).Execute()

Delete the realm

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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmDelete(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmEventsConfigGet

> RealmEventsConfigRepresentation AdminRealmsRealmEventsConfigGet(ctx, realm).Execute()

Get the events provider configuration Returns JSON object with events provider configuration

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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsConfigGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmEventsConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmEventsConfigGet`: RealmEventsConfigRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmEventsConfigGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmEventsConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RealmEventsConfigRepresentation**](RealmEventsConfigRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmEventsConfigPut

> AdminRealmsRealmEventsConfigPut(ctx, realm).RealmEventsConfigRepresentation(realmEventsConfigRepresentation).Execute()





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
	realmEventsConfigRepresentation := *openapiclient.NewRealmEventsConfigRepresentation() // RealmEventsConfigRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsConfigPut(context.Background(), realm).RealmEventsConfigRepresentation(realmEventsConfigRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmEventsConfigPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmEventsConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realmEventsConfigRepresentation** | [**RealmEventsConfigRepresentation**](RealmEventsConfigRepresentation.md) |  | 

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


## AdminRealmsRealmEventsDelete

> AdminRealmsRealmEventsDelete(ctx, realm).Execute()

Delete all events

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
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsDelete(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmEventsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmEventsDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmEventsGet

> []EventRepresentation AdminRealmsRealmEventsGet(ctx, realm).Client(client).DateFrom(dateFrom).DateTo(dateTo).First(first).IpAddress(ipAddress).Max(max).Type_(type_).User(user).Execute()

Get events Returns all events, or filters them based on URL query parameters listed here

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
	client := "client_example" // string | App or oauth client name (optional)
	dateFrom := "dateFrom_example" // string | From date (optional)
	dateTo := "dateTo_example" // string | To date (optional)
	first := int32(56) // int32 | Paging offset (optional)
	ipAddress := "ipAddress_example" // string | IP Address (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)
	type_ := []string{"Inner_example"} // []string | The types of events to return (optional)
	user := "user_example" // string | User id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmEventsGet(context.Background(), realm).Client(client).DateFrom(dateFrom).DateTo(dateTo).First(first).IpAddress(ipAddress).Max(max).Type_(type_).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmEventsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmEventsGet`: []EventRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | **string** | App or oauth client name | 
 **dateFrom** | **string** | From date | 
 **dateTo** | **string** | To date | 
 **first** | **int32** | Paging offset | 
 **ipAddress** | **string** | IP Address | 
 **max** | **int32** | Maximum results size (defaults to 100) | 
 **type_** | **[]string** | The types of events to return | 
 **user** | **string** | User id | 

### Return type

[**[]EventRepresentation**](EventRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGet

> RealmRepresentation AdminRealmsRealmGet(ctx, realm).Execute()

Get the top-level representation of the realm It will not include nested information like User and Client representations.

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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGet`: RealmRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RealmRepresentation**](RealmRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmGroupByPathPathGet

> GroupRepresentation AdminRealmsRealmGroupByPathPathGet(ctx, realm, path).Execute()



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmGroupByPathPathGet(context.Background(), realm, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmGroupByPathPathGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmGroupByPathPathGet`: GroupRepresentation
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmGroupByPathPathGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmGroupByPathPathGetRequest struct via the builder pattern


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


## AdminRealmsRealmLocalizationGet

> []string AdminRealmsRealmLocalizationGet(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmLocalizationGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmLocalizationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationGetRequest struct via the builder pattern


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


## AdminRealmsRealmLocalizationLocaleDelete

> AdminRealmsRealmLocalizationLocaleDelete(ctx, realm, locale).Execute()



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
	locale := "locale_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleDelete(context.Background(), realm, locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationLocaleDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmLocalizationLocaleGet

> map[string]string AdminRealmsRealmLocalizationLocaleGet(ctx, realm, locale).UseRealmDefaultLocaleFallback(useRealmDefaultLocaleFallback).Execute()



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
	locale := "locale_example" // string | 
	useRealmDefaultLocaleFallback := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleGet(context.Background(), realm, locale).UseRealmDefaultLocaleFallback(useRealmDefaultLocaleFallback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmLocalizationLocaleGet`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationLocaleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **useRealmDefaultLocaleFallback** | **bool** |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmLocalizationLocaleKeyDelete

> AdminRealmsRealmLocalizationLocaleKeyDelete(ctx, realm, key, locale).Execute()



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
	key := "key_example" // string | 
	locale := "locale_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyDelete(context.Background(), realm, key, locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**key** | **string** |  | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationLocaleKeyDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmLocalizationLocaleKeyGet

> string AdminRealmsRealmLocalizationLocaleKeyGet(ctx, realm, key, locale).Execute()



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
	key := "key_example" // string | 
	locale := "locale_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyGet(context.Background(), realm, key, locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmLocalizationLocaleKeyGet`: string
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**key** | **string** |  | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationLocaleKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmLocalizationLocaleKeyPut

> AdminRealmsRealmLocalizationLocaleKeyPut(ctx, realm, key, locale).Body(body).Execute()



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
	key := "key_example" // string | 
	locale := "locale_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyPut(context.Background(), realm, key, locale).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationLocaleKeyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**key** | **string** |  | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationLocaleKeyPutRequest struct via the builder pattern


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


## AdminRealmsRealmLocalizationLocalePost

> AdminRealmsRealmLocalizationLocalePost(ctx, realm, locale).RequestBody(requestBody).Execute()

Import localization from uploaded JSON file

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
	locale := "locale_example" // string | 
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLocalizationLocalePost(context.Background(), realm, locale).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLocalizationLocalePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**locale** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLocalizationLocalePostRequest struct via the builder pattern


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


## AdminRealmsRealmLogoutAllPost

> GlobalRequestResult AdminRealmsRealmLogoutAllPost(ctx, realm).Execute()

Removes all user sessions.



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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmLogoutAllPost(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmLogoutAllPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmLogoutAllPost`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmLogoutAllPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmLogoutAllPostRequest struct via the builder pattern


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


## AdminRealmsRealmPartialExportPost

> AdminRealmsRealmPartialExportPost(ctx, realm).ExportClients(exportClients).ExportGroupsAndRoles(exportGroupsAndRoles).Execute()

Partial export of existing realm into a JSON file.

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
	exportClients := true // bool |  (optional)
	exportGroupsAndRoles := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPartialExportPost(context.Background(), realm).ExportClients(exportClients).ExportGroupsAndRoles(exportGroupsAndRoles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmPartialExportPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmPartialExportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportClients** | **bool** |  | 
 **exportGroupsAndRoles** | **bool** |  | 

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


## AdminRealmsRealmPartialImportPost

> AdminRealmsRealmPartialImportPost(ctx, realm).Body(body).Execute()

Partial import from a JSON file to an existing realm.

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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPartialImportPost(context.Background(), realm).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmPartialImportPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmPartialImportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

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


## AdminRealmsRealmPushRevocationPost

> GlobalRequestResult AdminRealmsRealmPushRevocationPost(ctx, realm).Execute()

Push the realm's revocation policy to any client that has an admin url associated with it.

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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPushRevocationPost(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmPushRevocationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmPushRevocationPost`: GlobalRequestResult
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmPushRevocationPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmPushRevocationPostRequest struct via the builder pattern


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


## AdminRealmsRealmPut

> AdminRealmsRealmPut(ctx, realm).RealmRepresentation(realmRepresentation).Execute()

Update the top-level information of the realm Any user, roles or client information in the representation will be ignored.



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
	realmRepresentation := *openapiclient.NewRealmRepresentation() // RealmRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmPut(context.Background(), realm).RealmRepresentation(realmRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realmRepresentation** | [**RealmRepresentation**](RealmRepresentation.md) |  | 

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


## AdminRealmsRealmSessionsSessionDelete

> AdminRealmsRealmSessionsSessionDelete(ctx, realm, session).IsOffline(isOffline).Execute()

Remove a specific user session.



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
	session := "session_example" // string | 
	isOffline := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmSessionsSessionDelete(context.Background(), realm, session).IsOffline(isOffline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmSessionsSessionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**session** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmSessionsSessionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isOffline** | **bool** |  | [default to false]

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


## AdminRealmsRealmTestSMTPConnectionPost

> AdminRealmsRealmTestSMTPConnectionPost(ctx, realm).RequestBody(requestBody).Execute()

Test SMTP connection with current logged in user

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
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmTestSMTPConnectionPost(context.Background(), realm).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmTestSMTPConnectionPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmTestSMTPConnectionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmUsersManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmUsersManagementPermissionsGet(ctx, realm).Execute()



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
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersManagementPermissionsGetRequest struct via the builder pattern


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


## AdminRealmsRealmUsersManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmUsersManagementPermissionsPut(ctx, realm).ManagementPermissionReference(managementPermissionReference).Execute()



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
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsPut(context.Background(), realm).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmUsersManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `RealmsAdminAPI.AdminRealmsRealmUsersManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmUsersManagementPermissionsPutRequest struct via the builder pattern


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

