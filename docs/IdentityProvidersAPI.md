# \IdentityProvidersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmIdentityProviderImportConfigPost**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderImportConfigPost) | **Post** /admin/realms/{realm}/identity-provider/import-config | Import identity provider from JSON body
[**AdminRealmsRealmIdentityProviderInstancesAliasDelete**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasDelete) | **Delete** /admin/realms/{realm}/identity-provider/instances/{alias} | Delete the identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasExportGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasExportGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias}/export | Export public broker configuration for identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias} | Get the identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias}/management/permissions | Return object stating whether client Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut) | **Put** /admin/realms/{realm}/identity-provider/instances/{alias}/management/permissions | Return object stating whether client Authorization permissions have been initialized or not and a reference
[**AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias}/mapper-types | Get mapper types for identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasMappersGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasMappersGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias}/mappers | Get mappers for identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasMappersIdDelete**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasMappersIdDelete) | **Delete** /admin/realms/{realm}/identity-provider/instances/{alias}/mappers/{id} | Delete a mapper for the identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias}/mappers/{id} | Get mapper by id for the identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasMappersIdPut**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasMappersIdPut) | **Put** /admin/realms/{realm}/identity-provider/instances/{alias}/mappers/{id} | Update a mapper for the identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasMappersPost**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasMappersPost) | **Post** /admin/realms/{realm}/identity-provider/instances/{alias}/mappers | Add a mapper to identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasPut**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasPut) | **Put** /admin/realms/{realm}/identity-provider/instances/{alias} | Update the identity provider
[**AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet) | **Get** /admin/realms/{realm}/identity-provider/instances/{alias}/reload-keys | Reaload keys for the identity provider if the provider supports it, \&quot;true\&quot; is returned if reload was performed, \&quot;false\&quot; if not.
[**AdminRealmsRealmIdentityProviderInstancesGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesGet) | **Get** /admin/realms/{realm}/identity-provider/instances | List identity providers
[**AdminRealmsRealmIdentityProviderInstancesPost**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderInstancesPost) | **Post** /admin/realms/{realm}/identity-provider/instances | Create a new identity provider
[**AdminRealmsRealmIdentityProviderProvidersProviderIdGet**](IdentityProvidersAPI.md#AdminRealmsRealmIdentityProviderProvidersProviderIdGet) | **Get** /admin/realms/{realm}/identity-provider/providers/{provider_id} | Get the identity provider factory for that provider id



## AdminRealmsRealmIdentityProviderImportConfigPost

> map[string]string AdminRealmsRealmIdentityProviderImportConfigPost(ctx, realm).RequestBody(requestBody).Execute()

Import identity provider from JSON body



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderImportConfigPost(context.Background(), realm).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderImportConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderImportConfigPost`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderImportConfigPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderImportConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesAliasDelete

> AdminRealmsRealmIdentityProviderInstancesAliasDelete(ctx, realm, alias).Execute()

Delete the identity provider

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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasDelete(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmIdentityProviderInstancesAliasExportGet

> AdminRealmsRealmIdentityProviderInstancesAliasExportGet(ctx, realm, alias).Format(format).Execute()

Export public broker configuration for identity provider

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
	alias := "alias_example" // string | 
	format := "format_example" // string | Format to use (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasExportGet(context.Background(), realm, alias).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasExportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasExportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Format to use | 

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


## AdminRealmsRealmIdentityProviderInstancesAliasGet

> IdentityProviderRepresentation AdminRealmsRealmIdentityProviderInstancesAliasGet(ctx, realm, alias).Execute()

Get the identity provider

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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasGet`: IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityProviderRepresentation**](IdentityProviderRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet

> ManagementPermissionReference AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet(ctx, realm, alias).Execute()

Return object stating whether client Authorization permissions have been initialized or not and a reference

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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsGetRequest struct via the builder pattern


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


## AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut

> ManagementPermissionReference AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut(ctx, realm, alias).ManagementPermissionReference(managementPermissionReference).Execute()

Return object stating whether client Authorization permissions have been initialized or not and a reference

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
	alias := "alias_example" // string | 
	managementPermissionReference := *openapiclient.NewManagementPermissionReference() // ManagementPermissionReference |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut(context.Background(), realm, alias).ManagementPermissionReference(managementPermissionReference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut`: ManagementPermissionReference
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasManagementPermissionsPutRequest struct via the builder pattern


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


## AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet

> map[string]IdentityProviderMapperTypeRepresentation AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet(ctx, realm, alias).Execute()

Get mapper types for identity provider

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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet`: map[string]IdentityProviderMapperTypeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasMapperTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]IdentityProviderMapperTypeRepresentation**](IdentityProviderMapperTypeRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesAliasMappersGet

> []IdentityProviderMapperRepresentation AdminRealmsRealmIdentityProviderInstancesAliasMappersGet(ctx, realm, alias).Execute()

Get mappers for identity provider

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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasMappersGet`: []IdentityProviderMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasMappersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesAliasMappersIdDelete

> AdminRealmsRealmIdentityProviderInstancesAliasMappersIdDelete(ctx, realm, alias, id).Execute()

Delete a mapper for the identity provider

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
	alias := "alias_example" // string | 
	id := "id_example" // string | Mapper id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdDelete(context.Background(), realm, alias, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 
**id** | **string** | Mapper id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasMappersIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet

> IdentityProviderMapperRepresentation AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet(ctx, realm, alias, id).Execute()

Get mapper by id for the identity provider

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
	alias := "alias_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet(context.Background(), realm, alias, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet`: IdentityProviderMapperRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasMappersIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesAliasMappersIdPut

> AdminRealmsRealmIdentityProviderInstancesAliasMappersIdPut(ctx, realm, alias, id).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()

Update a mapper for the identity provider

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
	alias := "alias_example" // string | 
	id := "id_example" // string | Mapper id
	identityProviderMapperRepresentation := *openapiclient.NewIdentityProviderMapperRepresentation() // IdentityProviderMapperRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdPut(context.Background(), realm, alias, id).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 
**id** | **string** | Mapper id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasMappersIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **identityProviderMapperRepresentation** | [**IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md) |  | 

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


## AdminRealmsRealmIdentityProviderInstancesAliasMappersPost

> AdminRealmsRealmIdentityProviderInstancesAliasMappersPost(ctx, realm, alias).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()

Add a mapper to identity provider

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
	alias := "alias_example" // string | 
	identityProviderMapperRepresentation := *openapiclient.NewIdentityProviderMapperRepresentation() // IdentityProviderMapperRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersPost(context.Background(), realm, alias).IdentityProviderMapperRepresentation(identityProviderMapperRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasMappersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasMappersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProviderMapperRepresentation** | [**IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md) |  | 

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


## AdminRealmsRealmIdentityProviderInstancesAliasPut

> AdminRealmsRealmIdentityProviderInstancesAliasPut(ctx, realm, alias).IdentityProviderRepresentation(identityProviderRepresentation).Execute()

Update the identity provider

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
	alias := "alias_example" // string | 
	identityProviderRepresentation := *openapiclient.NewIdentityProviderRepresentation() // IdentityProviderRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasPut(context.Background(), realm, alias).IdentityProviderRepresentation(identityProviderRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProviderRepresentation** | [**IdentityProviderRepresentation**](IdentityProviderRepresentation.md) |  | 

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


## AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet

> bool AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet(ctx, realm, alias).Execute()

Reaload keys for the identity provider if the provider supports it, \"true\" is returned if reload was performed, \"false\" if not.

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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet`: bool
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesAliasReloadKeysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesGet

> []IdentityProviderRepresentation AdminRealmsRealmIdentityProviderInstancesGet(ctx, realm).BriefRepresentation(briefRepresentation).First(first).Max(max).RealmOnly(realmOnly).Search(search).Execute()

List identity providers

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
	first := int32(56) // int32 | Pagination offset (optional)
	max := int32(56) // int32 | Maximum results size (defaults to 100) (optional)
	realmOnly := true // bool | Boolean which defines if only realm-level IDPs (not associated with orgs) should be returned (default: false) (optional)
	search := "search_example" // string | Filter specific providers by name. Search can be prefix (name*), contains (*name*) or exact (\"name\"). Default prefixed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesGet(context.Background(), realm).BriefRepresentation(briefRepresentation).First(first).Max(max).RealmOnly(realmOnly).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderInstancesGet`: []IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **briefRepresentation** | **bool** | Boolean which defines whether brief representations are returned (default: false) | 
 **first** | **int32** | Pagination offset | 
 **max** | **int32** | Maximum results size (defaults to 100) | 
 **realmOnly** | **bool** | Boolean which defines if only realm-level IDPs (not associated with orgs) should be returned (default: false) | 
 **search** | **string** | Filter specific providers by name. Search can be prefix (name*), contains (*name*) or exact (\&quot;name\&quot;). Default prefixed. | 

### Return type

[**[]IdentityProviderRepresentation**](IdentityProviderRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmIdentityProviderInstancesPost

> AdminRealmsRealmIdentityProviderInstancesPost(ctx, realm).IdentityProviderRepresentation(identityProviderRepresentation).Execute()

Create a new identity provider

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
	identityProviderRepresentation := *openapiclient.NewIdentityProviderRepresentation() // IdentityProviderRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesPost(context.Background(), realm).IdentityProviderRepresentation(identityProviderRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderInstancesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderInstancesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityProviderRepresentation** | [**IdentityProviderRepresentation**](IdentityProviderRepresentation.md) |  | 

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


## AdminRealmsRealmIdentityProviderProvidersProviderIdGet

> map[string]interface{} AdminRealmsRealmIdentityProviderProvidersProviderIdGet(ctx, realm, providerId).Execute()

Get the identity provider factory for that provider id

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
	providerId := "providerId_example" // string | The provider id to get the factory

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentityProvidersAPI.AdminRealmsRealmIdentityProviderProvidersProviderIdGet(context.Background(), realm, providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderProvidersProviderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmIdentityProviderProvidersProviderIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersAPI.AdminRealmsRealmIdentityProviderProvidersProviderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**providerId** | **string** | The provider id to get the factory | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmIdentityProviderProvidersProviderIdGetRequest struct via the builder pattern


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

