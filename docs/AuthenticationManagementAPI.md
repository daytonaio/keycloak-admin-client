# \AuthenticationManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmAuthenticationAuthenticatorProvidersGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationAuthenticatorProvidersGet) | **Get** /admin/realms/{realm}/authentication/authenticator-providers | Get authenticator providers Returns a stream of authenticator providers.
[**AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet) | **Get** /admin/realms/{realm}/authentication/client-authenticator-providers | Get client authenticator providers Returns a stream of client authenticator providers.
[**AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet) | **Get** /admin/realms/{realm}/authentication/config-description/{providerId} | Get authenticator provider&#39;s configuration description
[**AdminRealmsRealmAuthenticationConfigIdDelete**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationConfigIdDelete) | **Delete** /admin/realms/{realm}/authentication/config/{id} | Delete authenticator configuration
[**AdminRealmsRealmAuthenticationConfigIdGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationConfigIdGet) | **Get** /admin/realms/{realm}/authentication/config/{id} | Get authenticator configuration
[**AdminRealmsRealmAuthenticationConfigIdPut**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationConfigIdPut) | **Put** /admin/realms/{realm}/authentication/config/{id} | Update authenticator configuration
[**AdminRealmsRealmAuthenticationConfigPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationConfigPost) | **Post** /admin/realms/{realm}/authentication/config | Create new authenticator configuration
[**AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet) | **Get** /admin/realms/{realm}/authentication/executions/{executionId}/config/{id} | Get execution&#39;s configuration
[**AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost) | **Post** /admin/realms/{realm}/authentication/executions/{executionId}/config | Update execution with new configuration
[**AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete) | **Delete** /admin/realms/{realm}/authentication/executions/{executionId} | Delete execution
[**AdminRealmsRealmAuthenticationExecutionsExecutionIdGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsExecutionIdGet) | **Get** /admin/realms/{realm}/authentication/executions/{executionId} | Get Single Execution
[**AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost) | **Post** /admin/realms/{realm}/authentication/executions/{executionId}/lower-priority | Lower execution&#39;s priority
[**AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost) | **Post** /admin/realms/{realm}/authentication/executions/{executionId}/raise-priority | Raise execution&#39;s priority
[**AdminRealmsRealmAuthenticationExecutionsPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationExecutionsPost) | **Post** /admin/realms/{realm}/authentication/executions | Add new authentication execution
[**AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost) | **Post** /admin/realms/{realm}/authentication/flows/{flowAlias}/copy | Copy existing authentication flow under a new name The new name is given as &#39;newName&#39; attribute of the passed JSON object
[**AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost) | **Post** /admin/realms/{realm}/authentication/flows/{flowAlias}/executions/execution | Add new authentication execution to a flow
[**AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost) | **Post** /admin/realms/{realm}/authentication/flows/{flowAlias}/executions/flow | Add new flow with new execution to existing flow
[**AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet) | **Get** /admin/realms/{realm}/authentication/flows/{flowAlias}/executions | Get authentication executions for a flow
[**AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut) | **Put** /admin/realms/{realm}/authentication/flows/{flowAlias}/executions | Update authentication executions of a Flow
[**AdminRealmsRealmAuthenticationFlowsGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsGet) | **Get** /admin/realms/{realm}/authentication/flows | Get authentication flows Returns a stream of authentication flows.
[**AdminRealmsRealmAuthenticationFlowsIdDelete**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsIdDelete) | **Delete** /admin/realms/{realm}/authentication/flows/{id} | Delete an authentication flow
[**AdminRealmsRealmAuthenticationFlowsIdGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsIdGet) | **Get** /admin/realms/{realm}/authentication/flows/{id} | Get authentication flow for id
[**AdminRealmsRealmAuthenticationFlowsIdPut**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsIdPut) | **Put** /admin/realms/{realm}/authentication/flows/{id} | Update an authentication flow
[**AdminRealmsRealmAuthenticationFlowsPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFlowsPost) | **Post** /admin/realms/{realm}/authentication/flows | Create a new authentication flow
[**AdminRealmsRealmAuthenticationFormActionProvidersGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFormActionProvidersGet) | **Get** /admin/realms/{realm}/authentication/form-action-providers | Get form action providers Returns a stream of form action providers.
[**AdminRealmsRealmAuthenticationFormProvidersGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationFormProvidersGet) | **Get** /admin/realms/{realm}/authentication/form-providers | Get form providers Returns a stream of form providers.
[**AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet) | **Get** /admin/realms/{realm}/authentication/per-client-config-description | Get configuration descriptions for all clients
[**AdminRealmsRealmAuthenticationRegisterRequiredActionPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRegisterRequiredActionPost) | **Post** /admin/realms/{realm}/authentication/register-required-action | Register a new required actions
[**AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete) | **Delete** /admin/realms/{realm}/authentication/required-actions/{alias}/config | Delete RequiredAction configuration
[**AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet) | **Get** /admin/realms/{realm}/authentication/required-actions/{alias}/config-description | Get RequiredAction provider configuration description
[**AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet) | **Get** /admin/realms/{realm}/authentication/required-actions/{alias}/config | Get RequiredAction configuration
[**AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut) | **Put** /admin/realms/{realm}/authentication/required-actions/{alias}/config | Update RequiredAction configuration
[**AdminRealmsRealmAuthenticationRequiredActionsAliasDelete**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasDelete) | **Delete** /admin/realms/{realm}/authentication/required-actions/{alias} | Delete required action
[**AdminRealmsRealmAuthenticationRequiredActionsAliasGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasGet) | **Get** /admin/realms/{realm}/authentication/required-actions/{alias} | Get required action for alias
[**AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost) | **Post** /admin/realms/{realm}/authentication/required-actions/{alias}/lower-priority | Lower required action&#39;s priority
[**AdminRealmsRealmAuthenticationRequiredActionsAliasPut**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasPut) | **Put** /admin/realms/{realm}/authentication/required-actions/{alias} | Update required action
[**AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost) | **Post** /admin/realms/{realm}/authentication/required-actions/{alias}/raise-priority | Raise required action&#39;s priority
[**AdminRealmsRealmAuthenticationRequiredActionsGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationRequiredActionsGet) | **Get** /admin/realms/{realm}/authentication/required-actions | Get required actions Returns a stream of required actions.
[**AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet**](AuthenticationManagementAPI.md#AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet) | **Get** /admin/realms/{realm}/authentication/unregistered-required-actions | Get unregistered required actions Returns a stream of unregistered required actions.



## AdminRealmsRealmAuthenticationAuthenticatorProvidersGet

> []map[string]interface{} AdminRealmsRealmAuthenticationAuthenticatorProvidersGet(ctx, realm).Execute()

Get authenticator providers Returns a stream of authenticator providers.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationAuthenticatorProvidersGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationAuthenticatorProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationAuthenticatorProvidersGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationAuthenticatorProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationAuthenticatorProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet

> []map[string]interface{} AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet(ctx, realm).Execute()

Get client authenticator providers Returns a stream of client authenticator providers.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationClientAuthenticatorProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationClientAuthenticatorProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet

> AuthenticatorConfigInfoRepresentation AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet(ctx, realm, providerId).Execute()

Get authenticator provider's configuration description

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
	providerId := "providerId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet(context.Background(), realm, providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet`: AuthenticatorConfigInfoRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigDescriptionProviderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**providerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationConfigDescriptionProviderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthenticatorConfigInfoRepresentation**](AuthenticatorConfigInfoRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationConfigIdDelete

> AdminRealmsRealmAuthenticationConfigIdDelete(ctx, realm, id).Execute()

Delete authenticator configuration

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
	id := "id_example" // string | Configuration id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdDelete(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationConfigIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationConfigIdGet

> AuthenticatorConfigRepresentation AdminRealmsRealmAuthenticationConfigIdGet(ctx, realm, id).Execute()

Get authenticator configuration

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
	id := "id_example" // string | Configuration id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationConfigIdGet`: AuthenticatorConfigRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationConfigIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthenticatorConfigRepresentation**](AuthenticatorConfigRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationConfigIdPut

> AdminRealmsRealmAuthenticationConfigIdPut(ctx, realm, id).AuthenticatorConfigRepresentation(authenticatorConfigRepresentation).Execute()

Update authenticator configuration

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
	id := "id_example" // string | Configuration id
	authenticatorConfigRepresentation := *openapiclient.NewAuthenticatorConfigRepresentation() // AuthenticatorConfigRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdPut(context.Background(), realm, id).AuthenticatorConfigRepresentation(authenticatorConfigRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationConfigIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authenticatorConfigRepresentation** | [**AuthenticatorConfigRepresentation**](AuthenticatorConfigRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationConfigPost

> AdminRealmsRealmAuthenticationConfigPost(ctx, realm).AuthenticatorConfigRepresentation(authenticatorConfigRepresentation).Execute()

Create new authenticator configuration

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
	authenticatorConfigRepresentation := *openapiclient.NewAuthenticatorConfigRepresentation() // AuthenticatorConfigRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigPost(context.Background(), realm).AuthenticatorConfigRepresentation(authenticatorConfigRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationConfigPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticatorConfigRepresentation** | [**AuthenticatorConfigRepresentation**](AuthenticatorConfigRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet

> AuthenticatorConfigRepresentation AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet(ctx, realm, executionId, id).Execute()

Get execution's configuration

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
	executionId := "executionId_example" // string | Execution id
	id := "id_example" // string | Configuration id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet(context.Background(), realm, executionId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet`: AuthenticatorConfigRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**executionId** | **string** | Execution id | 
**id** | **string** | Configuration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsExecutionIdConfigIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AuthenticatorConfigRepresentation**](AuthenticatorConfigRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost

> AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost(ctx, realm, executionId).AuthenticatorConfigRepresentation(authenticatorConfigRepresentation).Execute()

Update execution with new configuration

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
	executionId := "executionId_example" // string | Execution id
	authenticatorConfigRepresentation := *openapiclient.NewAuthenticatorConfigRepresentation() // AuthenticatorConfigRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost(context.Background(), realm, executionId).AuthenticatorConfigRepresentation(authenticatorConfigRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**executionId** | **string** | Execution id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsExecutionIdConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authenticatorConfigRepresentation** | [**AuthenticatorConfigRepresentation**](AuthenticatorConfigRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete

> AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete(ctx, realm, executionId).Execute()

Delete execution

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
	executionId := "executionId_example" // string | Execution id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete(context.Background(), realm, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**executionId** | **string** | Execution id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsExecutionIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationExecutionsExecutionIdGet

> AuthenticationExecutionRepresentation AdminRealmsRealmAuthenticationExecutionsExecutionIdGet(ctx, realm, executionId).Execute()

Get Single Execution

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
	executionId := "executionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdGet(context.Background(), realm, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationExecutionsExecutionIdGet`: AuthenticationExecutionRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**executionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsExecutionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthenticationExecutionRepresentation**](AuthenticationExecutionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost

> AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost(ctx, realm, executionId).Execute()

Lower execution's priority

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
	executionId := "executionId_example" // string | Execution id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost(context.Background(), realm, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**executionId** | **string** | Execution id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsExecutionIdLowerPriorityPostRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost

> AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost(ctx, realm, executionId).Execute()

Raise execution's priority

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
	executionId := "executionId_example" // string | Execution id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost(context.Background(), realm, executionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**executionId** | **string** | Execution id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsExecutionIdRaisePriorityPostRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationExecutionsPost

> AdminRealmsRealmAuthenticationExecutionsPost(ctx, realm).AuthenticationExecutionRepresentation(authenticationExecutionRepresentation).Execute()

Add new authentication execution

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
	authenticationExecutionRepresentation := *openapiclient.NewAuthenticationExecutionRepresentation() // AuthenticationExecutionRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsPost(context.Background(), realm).AuthenticationExecutionRepresentation(authenticationExecutionRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationExecutionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationExecutionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationExecutionRepresentation** | [**AuthenticationExecutionRepresentation**](AuthenticationExecutionRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost

> AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost(ctx, realm, flowAlias).RequestBody(requestBody).Execute()

Copy existing authentication flow under a new name The new name is given as 'newName' attribute of the passed JSON object

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
	flowAlias := "flowAlias_example" // string | name of the existing authentication flow
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost(context.Background(), realm, flowAlias).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**flowAlias** | **string** | name of the existing authentication flow | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsFlowAliasCopyPostRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost

> AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost(ctx, realm, flowAlias).RequestBody(requestBody).Execute()

Add new authentication execution to a flow

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
	flowAlias := "flowAlias_example" // string | Alias of parent flow
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost(context.Background(), realm, flowAlias).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**flowAlias** | **string** | Alias of parent flow | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsExecutionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 

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


## AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost

> AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost(ctx, realm, flowAlias).RequestBody(requestBody).Execute()

Add new flow with new execution to existing flow

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
	flowAlias := "flowAlias_example" // string | Alias of parent authentication flow
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost(context.Background(), realm, flowAlias).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**flowAlias** | **string** | Alias of parent authentication flow | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsFlowPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 

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


## AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet

> []AuthenticationExecutionInfoRepresentation AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet(ctx, realm, flowAlias).Execute()

Get authentication executions for a flow

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
	flowAlias := "flowAlias_example" // string | Flow alias

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet(context.Background(), realm, flowAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet`: []AuthenticationExecutionInfoRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**flowAlias** | **string** | Flow alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AuthenticationExecutionInfoRepresentation**](AuthenticationExecutionInfoRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut

> AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut(ctx, realm, flowAlias).AuthenticationExecutionInfoRepresentation(authenticationExecutionInfoRepresentation).Execute()

Update authentication executions of a Flow

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
	flowAlias := "flowAlias_example" // string | Flow alias
	authenticationExecutionInfoRepresentation := *openapiclient.NewAuthenticationExecutionInfoRepresentation() // AuthenticationExecutionInfoRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut(context.Background(), realm, flowAlias).AuthenticationExecutionInfoRepresentation(authenticationExecutionInfoRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**flowAlias** | **string** | Flow alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsFlowAliasExecutionsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authenticationExecutionInfoRepresentation** | [**AuthenticationExecutionInfoRepresentation**](AuthenticationExecutionInfoRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationFlowsGet

> []AuthenticationFlowRepresentation AdminRealmsRealmAuthenticationFlowsGet(ctx, realm).Execute()

Get authentication flows Returns a stream of authentication flows.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationFlowsGet`: []AuthenticationFlowRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthenticationFlowRepresentation**](AuthenticationFlowRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationFlowsIdDelete

> AdminRealmsRealmAuthenticationFlowsIdDelete(ctx, realm, id).Execute()

Delete an authentication flow

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
	id := "id_example" // string | Flow id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdDelete(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Flow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationFlowsIdGet

> AuthenticationFlowRepresentation AdminRealmsRealmAuthenticationFlowsIdGet(ctx, realm, id).Execute()

Get authentication flow for id

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
	id := "id_example" // string | Flow id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationFlowsIdGet`: AuthenticationFlowRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Flow id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthenticationFlowRepresentation**](AuthenticationFlowRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationFlowsIdPut

> AdminRealmsRealmAuthenticationFlowsIdPut(ctx, realm, id).AuthenticationFlowRepresentation(authenticationFlowRepresentation).Execute()

Update an authentication flow

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
	id := "id_example" // string | 
	authenticationFlowRepresentation := *openapiclient.NewAuthenticationFlowRepresentation() // AuthenticationFlowRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdPut(context.Background(), realm, id).AuthenticationFlowRepresentation(authenticationFlowRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authenticationFlowRepresentation** | [**AuthenticationFlowRepresentation**](AuthenticationFlowRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationFlowsPost

> AdminRealmsRealmAuthenticationFlowsPost(ctx, realm).AuthenticationFlowRepresentation(authenticationFlowRepresentation).Execute()

Create a new authentication flow

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
	authenticationFlowRepresentation := *openapiclient.NewAuthenticationFlowRepresentation() // AuthenticationFlowRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsPost(context.Background(), realm).AuthenticationFlowRepresentation(authenticationFlowRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFlowsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFlowsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationFlowRepresentation** | [**AuthenticationFlowRepresentation**](AuthenticationFlowRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationFormActionProvidersGet

> []map[string]interface{} AdminRealmsRealmAuthenticationFormActionProvidersGet(ctx, realm).Execute()

Get form action providers Returns a stream of form action providers.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormActionProvidersGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormActionProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationFormActionProvidersGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormActionProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFormActionProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationFormProvidersGet

> []map[string]interface{} AdminRealmsRealmAuthenticationFormProvidersGet(ctx, realm).Execute()

Get form providers Returns a stream of form providers.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormProvidersGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationFormProvidersGet`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationFormProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationFormProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet

> map[string][]ConfigPropertyRepresentation AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet(ctx, realm).Execute()

Get configuration descriptions for all clients

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet`: map[string][]ConfigPropertyRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationPerClientConfigDescriptionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationPerClientConfigDescriptionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string][]ConfigPropertyRepresentation**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationRegisterRequiredActionPost

> AdminRealmsRealmAuthenticationRegisterRequiredActionPost(ctx, realm).RequestBody(requestBody).Execute()

Register a new required actions

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
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRegisterRequiredActionPost(context.Background(), realm).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRegisterRequiredActionPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRegisterRequiredActionPostRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete

> AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete(ctx, realm, alias).Execute()

Delete RequiredAction configuration

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasConfigDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet

> RequiredActionConfigInfoRepresentation AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet(ctx, realm, alias).Execute()

Get RequiredAction provider configuration description

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet`: RequiredActionConfigInfoRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasConfigDescriptionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequiredActionConfigInfoRepresentation**](RequiredActionConfigInfoRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet

> RequiredActionConfigRepresentation AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet(ctx, realm, alias).Execute()

Get RequiredAction configuration

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet`: RequiredActionConfigRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequiredActionConfigRepresentation**](RequiredActionConfigRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut

> AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut(ctx, realm, alias).RequiredActionConfigRepresentation(requiredActionConfigRepresentation).Execute()

Update RequiredAction configuration

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
	alias := "alias_example" // string | Alias of required action
	requiredActionConfigRepresentation := *openapiclient.NewRequiredActionConfigRepresentation() // RequiredActionConfigRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut(context.Background(), realm, alias).RequiredActionConfigRepresentation(requiredActionConfigRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requiredActionConfigRepresentation** | [**RequiredActionConfigRepresentation**](RequiredActionConfigRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationRequiredActionsAliasDelete

> AdminRealmsRealmAuthenticationRequiredActionsAliasDelete(ctx, realm, alias).Execute()

Delete required action

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasDelete(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationRequiredActionsAliasGet

> RequiredActionProviderRepresentation AdminRealmsRealmAuthenticationRequiredActionsAliasGet(ctx, realm, alias).Execute()

Get required action for alias

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasGet(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationRequiredActionsAliasGet`: RequiredActionProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequiredActionProviderRepresentation**](RequiredActionProviderRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost

> AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost(ctx, realm, alias).Execute()

Lower required action's priority

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasLowerPriorityPostRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationRequiredActionsAliasPut

> AdminRealmsRealmAuthenticationRequiredActionsAliasPut(ctx, realm, alias).RequiredActionProviderRepresentation(requiredActionProviderRepresentation).Execute()

Update required action

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
	alias := "alias_example" // string | Alias of required action
	requiredActionProviderRepresentation := *openapiclient.NewRequiredActionProviderRepresentation() // RequiredActionProviderRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasPut(context.Background(), realm, alias).RequiredActionProviderRepresentation(requiredActionProviderRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requiredActionProviderRepresentation** | [**RequiredActionProviderRepresentation**](RequiredActionProviderRepresentation.md) |  | 

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


## AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost

> AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost(ctx, realm, alias).Execute()

Raise required action's priority

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
	alias := "alias_example" // string | Alias of required action

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost(context.Background(), realm, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**alias** | **string** | Alias of required action | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsAliasRaisePriorityPostRequest struct via the builder pattern


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


## AdminRealmsRealmAuthenticationRequiredActionsGet

> []RequiredActionProviderRepresentation AdminRealmsRealmAuthenticationRequiredActionsGet(ctx, realm).Execute()

Get required actions Returns a stream of required actions.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationRequiredActionsGet`: []RequiredActionProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationRequiredActionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationRequiredActionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RequiredActionProviderRepresentation**](RequiredActionProviderRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet

> []map[string]string AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet(ctx, realm).Execute()

Get unregistered required actions Returns a stream of unregistered required actions.

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
	resp, r, err := apiClient.AuthenticationManagementAPI.AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet`: []map[string]string
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationManagementAPI.AdminRealmsRealmAuthenticationUnregisteredRequiredActionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAuthenticationUnregisteredRequiredActionsGetRequest struct via the builder pattern


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

