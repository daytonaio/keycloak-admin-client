# \ComponentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmComponentsGet**](ComponentAPI.md#AdminRealmsRealmComponentsGet) | **Get** /admin/realms/{realm}/components | 
[**AdminRealmsRealmComponentsIdDelete**](ComponentAPI.md#AdminRealmsRealmComponentsIdDelete) | **Delete** /admin/realms/{realm}/components/{id} | 
[**AdminRealmsRealmComponentsIdGet**](ComponentAPI.md#AdminRealmsRealmComponentsIdGet) | **Get** /admin/realms/{realm}/components/{id} | 
[**AdminRealmsRealmComponentsIdPut**](ComponentAPI.md#AdminRealmsRealmComponentsIdPut) | **Put** /admin/realms/{realm}/components/{id} | 
[**AdminRealmsRealmComponentsIdSubComponentTypesGet**](ComponentAPI.md#AdminRealmsRealmComponentsIdSubComponentTypesGet) | **Get** /admin/realms/{realm}/components/{id}/sub-component-types | List of subcomponent types that are available to configure for a particular parent component.
[**AdminRealmsRealmComponentsPost**](ComponentAPI.md#AdminRealmsRealmComponentsPost) | **Post** /admin/realms/{realm}/components | 



## AdminRealmsRealmComponentsGet

> []ComponentRepresentation AdminRealmsRealmComponentsGet(ctx, realm).Name(name).Parent(parent).Type_(type_).Execute()



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
	name := "name_example" // string |  (optional)
	parent := "parent_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentAPI.AdminRealmsRealmComponentsGet(context.Background(), realm).Name(name).Parent(parent).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.AdminRealmsRealmComponentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmComponentsGet`: []ComponentRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.AdminRealmsRealmComponentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmComponentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **parent** | **string** |  | 
 **type_** | **string** |  | 

### Return type

[**[]ComponentRepresentation**](ComponentRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmComponentsIdDelete

> AdminRealmsRealmComponentsIdDelete(ctx, realm, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentAPI.AdminRealmsRealmComponentsIdDelete(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.AdminRealmsRealmComponentsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmComponentsIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmComponentsIdGet

> ComponentRepresentation AdminRealmsRealmComponentsIdGet(ctx, realm, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentAPI.AdminRealmsRealmComponentsIdGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.AdminRealmsRealmComponentsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmComponentsIdGet`: ComponentRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.AdminRealmsRealmComponentsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmComponentsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ComponentRepresentation**](ComponentRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmComponentsIdPut

> AdminRealmsRealmComponentsIdPut(ctx, realm, id).ComponentRepresentation(componentRepresentation).Execute()



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
	id := "id_example" // string | 
	componentRepresentation := *openapiclient.NewComponentRepresentation() // ComponentRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentAPI.AdminRealmsRealmComponentsIdPut(context.Background(), realm, id).ComponentRepresentation(componentRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.AdminRealmsRealmComponentsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmComponentsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentRepresentation** | [**ComponentRepresentation**](ComponentRepresentation.md) |  | 

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


## AdminRealmsRealmComponentsIdSubComponentTypesGet

> []ComponentTypeRepresentation AdminRealmsRealmComponentsIdSubComponentTypesGet(ctx, realm, id).Type_(type_).Execute()

List of subcomponent types that are available to configure for a particular parent component.

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
	id := "id_example" // string | 
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentAPI.AdminRealmsRealmComponentsIdSubComponentTypesGet(context.Background(), realm, id).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.AdminRealmsRealmComponentsIdSubComponentTypesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmComponentsIdSubComponentTypesGet`: []ComponentTypeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ComponentAPI.AdminRealmsRealmComponentsIdSubComponentTypesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** |  | 

### Return type

[**[]ComponentTypeRepresentation**](ComponentTypeRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmComponentsPost

> AdminRealmsRealmComponentsPost(ctx, realm).ComponentRepresentation(componentRepresentation).Execute()



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
	componentRepresentation := *openapiclient.NewComponentRepresentation() // ComponentRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentAPI.AdminRealmsRealmComponentsPost(context.Background(), realm).ComponentRepresentation(componentRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentAPI.AdminRealmsRealmComponentsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmComponentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentRepresentation** | [**ComponentRepresentation**](ComponentRepresentation.md) |  | 

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

