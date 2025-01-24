# \ClientRegistrationPolicyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientRegistrationPolicyProvidersGet**](ClientRegistrationPolicyAPI.md#AdminRealmsRealmClientRegistrationPolicyProvidersGet) | **Get** /admin/realms/{realm}/client-registration-policy/providers | Base path for retrieve providers with the configProperties properly filled



## AdminRealmsRealmClientRegistrationPolicyProvidersGet

> []ComponentTypeRepresentation AdminRealmsRealmClientRegistrationPolicyProvidersGet(ctx, realm).Execute()

Base path for retrieve providers with the configProperties properly filled

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
	resp, r, err := apiClient.ClientRegistrationPolicyAPI.AdminRealmsRealmClientRegistrationPolicyProvidersGet(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientRegistrationPolicyAPI.AdminRealmsRealmClientRegistrationPolicyProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientRegistrationPolicyProvidersGet`: []ComponentTypeRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientRegistrationPolicyAPI.AdminRealmsRealmClientRegistrationPolicyProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientRegistrationPolicyProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

