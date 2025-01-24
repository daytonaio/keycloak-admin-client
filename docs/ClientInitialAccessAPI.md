# \ClientInitialAccessAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientsInitialAccessGet**](ClientInitialAccessAPI.md#AdminRealmsRealmClientsInitialAccessGet) | **Get** /admin/realms/{realm}/clients-initial-access | 
[**AdminRealmsRealmClientsInitialAccessIdDelete**](ClientInitialAccessAPI.md#AdminRealmsRealmClientsInitialAccessIdDelete) | **Delete** /admin/realms/{realm}/clients-initial-access/{id} | 
[**AdminRealmsRealmClientsInitialAccessPost**](ClientInitialAccessAPI.md#AdminRealmsRealmClientsInitialAccessPost) | **Post** /admin/realms/{realm}/clients-initial-access | Create a new initial access token.



## AdminRealmsRealmClientsInitialAccessGet

> []ClientInitialAccessPresentation AdminRealmsRealmClientsInitialAccessGet(ctx, realm).Execute()



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
	resp, r, err := apiClient.ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsInitialAccessGet`: []ClientInitialAccessPresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsInitialAccessGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ClientInitialAccessPresentation**](ClientInitialAccessPresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsInitialAccessIdDelete

> AdminRealmsRealmClientsInitialAccessIdDelete(ctx, realm, id).Execute()



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
	r, err := apiClient.ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessIdDelete(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsInitialAccessIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientsInitialAccessPost

> ClientInitialAccessCreatePresentation AdminRealmsRealmClientsInitialAccessPost(ctx, realm).ClientInitialAccessCreatePresentation(clientInitialAccessCreatePresentation).Execute()

Create a new initial access token.

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
	clientInitialAccessCreatePresentation := *openapiclient.NewClientInitialAccessCreatePresentation() // ClientInitialAccessCreatePresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessPost(context.Background(), realm).ClientInitialAccessCreatePresentation(clientInitialAccessCreatePresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsInitialAccessPost`: ClientInitialAccessCreatePresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientInitialAccessAPI.AdminRealmsRealmClientsInitialAccessPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsInitialAccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientInitialAccessCreatePresentation** | [**ClientInitialAccessCreatePresentation**](ClientInitialAccessCreatePresentation.md) |  | 

### Return type

[**ClientInitialAccessCreatePresentation**](ClientInitialAccessCreatePresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

