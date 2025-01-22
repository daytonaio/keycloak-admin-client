# \AttackDetectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmAttackDetectionBruteForceUsersDelete**](AttackDetectionAPI.md#AdminRealmsRealmAttackDetectionBruteForceUsersDelete) | **Delete** /admin/realms/{realm}/attack-detection/brute-force/users | Clear any user login failures for all users This can release temporary disabled users
[**AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete**](AttackDetectionAPI.md#AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete) | **Delete** /admin/realms/{realm}/attack-detection/brute-force/users/{userId} | Clear any user login failures for the user This can release temporary disabled user
[**AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet**](AttackDetectionAPI.md#AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet) | **Get** /admin/realms/{realm}/attack-detection/brute-force/users/{userId} | Get status of a username in brute force detection



## AdminRealmsRealmAttackDetectionBruteForceUsersDelete

> AdminRealmsRealmAttackDetectionBruteForceUsersDelete(ctx, realm).Execute()

Clear any user login failures for all users This can release temporary disabled users

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
	r, err := apiClient.AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersDelete(context.Background(), realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAttackDetectionBruteForceUsersDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete

> AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete(ctx, realm, userId).Execute()

Clear any user login failures for the user This can release temporary disabled user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet

> map[string]interface{} AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet(ctx, realm, userId).Execute()

Get status of a username in brute force detection

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet(context.Background(), realm, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AttackDetectionAPI.AdminRealmsRealmAttackDetectionBruteForceUsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmAttackDetectionBruteForceUsersUserIdGetRequest struct via the builder pattern


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

