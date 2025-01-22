# \ClientScopesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientScopesClientScopeIdDelete**](ClientScopesAPI.md#AdminRealmsRealmClientScopesClientScopeIdDelete) | **Delete** /admin/realms/{realm}/client-scopes/{client-scope-id} | Delete the client scope
[**AdminRealmsRealmClientScopesClientScopeIdGet**](ClientScopesAPI.md#AdminRealmsRealmClientScopesClientScopeIdGet) | **Get** /admin/realms/{realm}/client-scopes/{client-scope-id} | Get representation of the client scope
[**AdminRealmsRealmClientScopesClientScopeIdPut**](ClientScopesAPI.md#AdminRealmsRealmClientScopesClientScopeIdPut) | **Put** /admin/realms/{realm}/client-scopes/{client-scope-id} | Update the client scope
[**AdminRealmsRealmClientScopesGet**](ClientScopesAPI.md#AdminRealmsRealmClientScopesGet) | **Get** /admin/realms/{realm}/client-scopes | Get client scopes belonging to the realm Returns a list of client scopes belonging to the realm
[**AdminRealmsRealmClientScopesPost**](ClientScopesAPI.md#AdminRealmsRealmClientScopesPost) | **Post** /admin/realms/{realm}/client-scopes | Create a new client scope Client Scope’s name must be unique!
[**AdminRealmsRealmClientTemplatesClientScopeIdDelete**](ClientScopesAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdDelete) | **Delete** /admin/realms/{realm}/client-templates/{client-scope-id} | Delete the client scope
[**AdminRealmsRealmClientTemplatesClientScopeIdGet**](ClientScopesAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdGet) | **Get** /admin/realms/{realm}/client-templates/{client-scope-id} | Get representation of the client scope
[**AdminRealmsRealmClientTemplatesClientScopeIdPut**](ClientScopesAPI.md#AdminRealmsRealmClientTemplatesClientScopeIdPut) | **Put** /admin/realms/{realm}/client-templates/{client-scope-id} | Update the client scope
[**AdminRealmsRealmClientTemplatesGet**](ClientScopesAPI.md#AdminRealmsRealmClientTemplatesGet) | **Get** /admin/realms/{realm}/client-templates | Get client scopes belonging to the realm Returns a list of client scopes belonging to the realm
[**AdminRealmsRealmClientTemplatesPost**](ClientScopesAPI.md#AdminRealmsRealmClientTemplatesPost) | **Post** /admin/realms/{realm}/client-templates | Create a new client scope Client Scope’s name must be unique!



## AdminRealmsRealmClientScopesClientScopeIdDelete

> AdminRealmsRealmClientScopesClientScopeIdDelete(ctx, realm, clientScopeId).Execute()

Delete the client scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdDelete(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientScopesClientScopeIdGet

> ClientScopeRepresentation AdminRealmsRealmClientScopesClientScopeIdGet(ctx, realm, clientScopeId).Execute()

Get representation of the client scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesClientScopeIdGet`: ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientScopeRepresentation**](ClientScopeRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientScopesClientScopeIdPut

> AdminRealmsRealmClientScopesClientScopeIdPut(ctx, realm, clientScopeId).ClientScopeRepresentation(clientScopeRepresentation).Execute()

Update the client scope

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
	clientScopeId := "clientScopeId_example" // string | 
	clientScopeRepresentation := *openapiclient.NewClientScopeRepresentation() // ClientScopeRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdPut(context.Background(), realm, clientScopeId).ClientScopeRepresentation(clientScopeRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientScopesClientScopeIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesClientScopeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientScopeRepresentation** | [**ClientScopeRepresentation**](ClientScopeRepresentation.md) |  | 

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


## AdminRealmsRealmClientScopesGet

> []ClientScopeRepresentation AdminRealmsRealmClientScopesGet(ctx, realm).Execute()

Get client scopes belonging to the realm Returns a list of client scopes belonging to the realm

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientScopesGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientScopesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientScopesGet`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientScopesAPI.AdminRealmsRealmClientScopesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientScopesPost

> AdminRealmsRealmClientScopesPost(ctx, realm).ClientScopeRepresentation(clientScopeRepresentation).Execute()

Create a new client scope Client Scope’s name must be unique!

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
	clientScopeRepresentation := *openapiclient.NewClientScopeRepresentation() // ClientScopeRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientScopesPost(context.Background(), realm).ClientScopeRepresentation(clientScopeRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientScopesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientScopesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientScopeRepresentation** | [**ClientScopeRepresentation**](ClientScopeRepresentation.md) |  | 

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


## AdminRealmsRealmClientTemplatesClientScopeIdDelete

> AdminRealmsRealmClientTemplatesClientScopeIdDelete(ctx, realm, clientScopeId).Execute()

Delete the client scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdDelete(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmClientTemplatesClientScopeIdGet

> ClientScopeRepresentation AdminRealmsRealmClientTemplatesClientScopeIdGet(ctx, realm, clientScopeId).Execute()

Get representation of the client scope

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
	clientScopeId := "clientScopeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdGet(context.Background(), realm, clientScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesClientScopeIdGet`: ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientScopeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ClientScopeRepresentation**](ClientScopeRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientTemplatesClientScopeIdPut

> AdminRealmsRealmClientTemplatesClientScopeIdPut(ctx, realm, clientScopeId).ClientScopeRepresentation(clientScopeRepresentation).Execute()

Update the client scope

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
	clientScopeId := "clientScopeId_example" // string | 
	clientScopeRepresentation := *openapiclient.NewClientScopeRepresentation() // ClientScopeRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdPut(context.Background(), realm, clientScopeId).ClientScopeRepresentation(clientScopeRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientTemplatesClientScopeIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesClientScopeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientScopeRepresentation** | [**ClientScopeRepresentation**](ClientScopeRepresentation.md) |  | 

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


## AdminRealmsRealmClientTemplatesGet

> []ClientScopeRepresentation AdminRealmsRealmClientTemplatesGet(ctx, realm).Execute()

Get client scopes belonging to the realm Returns a list of client scopes belonging to the realm

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientTemplatesGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientTemplatesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientTemplatesGet`: []ClientScopeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientScopesAPI.AdminRealmsRealmClientTemplatesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesGetRequest struct via the builder pattern


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


## AdminRealmsRealmClientTemplatesPost

> AdminRealmsRealmClientTemplatesPost(ctx, realm).ClientScopeRepresentation(clientScopeRepresentation).Execute()

Create a new client scope Client Scope’s name must be unique!

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
	clientScopeRepresentation := *openapiclient.NewClientScopeRepresentation() // ClientScopeRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientScopesAPI.AdminRealmsRealmClientTemplatesPost(context.Background(), realm).ClientScopeRepresentation(clientScopeRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientScopesAPI.AdminRealmsRealmClientTemplatesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientTemplatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientScopeRepresentation** | [**ClientScopeRepresentation**](ClientScopeRepresentation.md) |  | 

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

