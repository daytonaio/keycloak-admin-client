# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmOrganizationsGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsGet) | **Get** /admin/realms/{realm}/organizations | Returns a paginated list of organizations filtered according to the specified parameters
[**AdminRealmsRealmOrganizationsIdDelete**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdDelete) | **Delete** /admin/realms/{realm}/organizations/{id} | Deletes the organization
[**AdminRealmsRealmOrganizationsIdGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdGet) | **Get** /admin/realms/{realm}/organizations/{id} | Returns the organization representation
[**AdminRealmsRealmOrganizationsIdIdentityProvidersAliasDelete**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdIdentityProvidersAliasDelete) | **Delete** /admin/realms/{realm}/organizations/{id}/identity-providers/{alias} | Removes the identity provider with the specified alias from the organization
[**AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet) | **Get** /admin/realms/{realm}/organizations/{id}/identity-providers/{alias} | Returns the identity provider associated with the organization that has the specified alias
[**AdminRealmsRealmOrganizationsIdIdentityProvidersGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdIdentityProvidersGet) | **Get** /admin/realms/{realm}/organizations/{id}/identity-providers | Returns all identity providers associated with the organization
[**AdminRealmsRealmOrganizationsIdIdentityProvidersPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdIdentityProvidersPost) | **Post** /admin/realms/{realm}/organizations/{id}/identity-providers | Adds the identity provider with the specified id to the organization
[**AdminRealmsRealmOrganizationsIdMembersCountGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersCountGet) | **Get** /admin/realms/{realm}/organizations/{id}/members/count | Returns number of members in the organization.
[**AdminRealmsRealmOrganizationsIdMembersGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersGet) | **Get** /admin/realms/{realm}/organizations/{id}/members | Returns a paginated list of organization members filtered according to the specified parameters
[**AdminRealmsRealmOrganizationsIdMembersInviteExistingUserPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersInviteExistingUserPost) | **Post** /admin/realms/{realm}/organizations/{id}/members/invite-existing-user | Invites an existing user to the organization, using the specified user id
[**AdminRealmsRealmOrganizationsIdMembersInviteUserPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersInviteUserPost) | **Post** /admin/realms/{realm}/organizations/{id}/members/invite-user | Invites an existing user or sends a registration link to a new user, based on the provided e-mail address.
[**AdminRealmsRealmOrganizationsIdMembersPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersPost) | **Post** /admin/realms/{realm}/organizations/{id}/members | Adds the user with the specified id as a member of the organization
[**AdminRealmsRealmOrganizationsIdMembersUserIdDelete**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersUserIdDelete) | **Delete** /admin/realms/{realm}/organizations/{id}/members/{userId} | Removes the user with the specified id from the organization
[**AdminRealmsRealmOrganizationsIdMembersUserIdGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersUserIdGet) | **Get** /admin/realms/{realm}/organizations/{id}/members/{userId} | Returns the member of the organization with the specified id
[**AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet) | **Get** /admin/realms/{realm}/organizations/{id}/members/{userId}/organizations | Returns the organizations associated with the user that has the specified id
[**AdminRealmsRealmOrganizationsIdPut**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsIdPut) | **Put** /admin/realms/{realm}/organizations/{id} | Updates the organization
[**AdminRealmsRealmOrganizationsMembersIdOrganizationsGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsMembersIdOrganizationsGet) | **Get** /admin/realms/{realm}/organizations/members/{id}/organizations | Returns the organizations associated with the user that has the specified id
[**AdminRealmsRealmOrganizationsPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsPost) | **Post** /admin/realms/{realm}/organizations | Creates a new organization



## AdminRealmsRealmOrganizationsGet

> []OrganizationRepresentation AdminRealmsRealmOrganizationsGet(ctx, realm).Exact(exact).First(first).Max(max).Q(q).Search(search).Execute()

Returns a paginated list of organizations filtered according to the specified parameters

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
	exact := true // bool | Boolean which defines whether the param 'search' must match exactly or not (optional)
	first := int32(56) // int32 | The position of the first result to be processed (pagination offset) (optional)
	max := int32(56) // int32 | The maximum number of results to be returned - defaults to 10 (optional)
	q := "q_example" // string | A query to search for custom attributes, in the format 'key1:value2 key2:value2' (optional)
	search := "search_example" // string | A String representing either an organization name or domain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsGet(context.Background(), realm).Exact(exact).First(first).Max(max).Q(q).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsGet`: []OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exact** | **bool** | Boolean which defines whether the param &#39;search&#39; must match exactly or not | 
 **first** | **int32** | The position of the first result to be processed (pagination offset) | 
 **max** | **int32** | The maximum number of results to be returned - defaults to 10 | 
 **q** | **string** | A query to search for custom attributes, in the format &#39;key1:value2 key2:value2&#39; | 
 **search** | **string** | A String representing either an organization name or domain | 

### Return type

[**[]OrganizationRepresentation**](OrganizationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdDelete

> AdminRealmsRealmOrganizationsIdDelete(ctx, realm, id).Execute()

Deletes the organization

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdDelete(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsIdGet

> OrganizationRepresentation AdminRealmsRealmOrganizationsIdGet(ctx, realm, id).Execute()

Returns the organization representation

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdGet`: OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrganizationRepresentation**](OrganizationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdIdentityProvidersAliasDelete

> AdminRealmsRealmOrganizationsIdIdentityProvidersAliasDelete(ctx, realm, id, alias).Execute()

Removes the identity provider with the specified alias from the organization



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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersAliasDelete(context.Background(), realm, id, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersAliasDelete``: %v\n", err)
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
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdIdentityProvidersAliasDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet

> IdentityProviderRepresentation AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet(ctx, realm, id, alias).Execute()

Returns the identity provider associated with the organization that has the specified alias



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
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet(context.Background(), realm, id, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet`: IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersAliasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdIdentityProvidersAliasGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsIdIdentityProvidersGet

> []IdentityProviderRepresentation AdminRealmsRealmOrganizationsIdIdentityProvidersGet(ctx, realm, id).Execute()

Returns all identity providers associated with the organization

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdIdentityProvidersGet`: []IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdIdentityProvidersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AdminRealmsRealmOrganizationsIdIdentityProvidersPost

> AdminRealmsRealmOrganizationsIdIdentityProvidersPost(ctx, realm, id).Body(body).Execute()

Adds the identity provider with the specified id to the organization



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersPost(context.Background(), realm, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdIdentityProvidersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdIdentityProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## AdminRealmsRealmOrganizationsIdMembersCountGet

> int64 AdminRealmsRealmOrganizationsIdMembersCountGet(ctx, realm, id).Execute()

Returns number of members in the organization.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersCountGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdMembersCountGet`: int64
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdMembersGet

> []MemberRepresentation AdminRealmsRealmOrganizationsIdMembersGet(ctx, realm, id).Exact(exact).First(first).Max(max).Search(search).Execute()

Returns a paginated list of organization members filtered according to the specified parameters

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
	exact := true // bool | Boolean which defines whether the param 'search' must match exactly or not (optional)
	first := int32(56) // int32 | The position of the first result to be processed (pagination offset) (optional)
	max := int32(56) // int32 | The maximum number of results to be returned. Defaults to 10 (optional)
	search := "search_example" // string | A String representing either a member's username, e-mail, first name, or last name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersGet(context.Background(), realm, id).Exact(exact).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdMembersGet`: []MemberRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Boolean which defines whether the param &#39;search&#39; must match exactly or not | 
 **first** | **int32** | The position of the first result to be processed (pagination offset) | 
 **max** | **int32** | The maximum number of results to be returned. Defaults to 10 | 
 **search** | **string** | A String representing either a member&#39;s username, e-mail, first name, or last name. | 

### Return type

[**[]MemberRepresentation**](MemberRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdMembersInviteExistingUserPost

> AdminRealmsRealmOrganizationsIdMembersInviteExistingUserPost(ctx, realm, id).Id2(id2).Execute()

Invites an existing user to the organization, using the specified user id

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
	id2 := "id_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersInviteExistingUserPost(context.Background(), realm, id).Id2(id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersInviteExistingUserPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersInviteExistingUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id2** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdMembersInviteUserPost

> AdminRealmsRealmOrganizationsIdMembersInviteUserPost(ctx, realm, id).Email(email).FirstName(firstName).LastName(lastName).Execute()

Invites an existing user or sends a registration link to a new user, based on the provided e-mail address.



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
	email := "email_example" // string |  (optional)
	firstName := "firstName_example" // string |  (optional)
	lastName := "lastName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersInviteUserPost(context.Background(), realm, id).Email(email).FirstName(firstName).LastName(lastName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersInviteUserPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersInviteUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **email** | **string** |  | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdMembersPost

> AdminRealmsRealmOrganizationsIdMembersPost(ctx, realm, id).Body(body).Execute()

Adds the user with the specified id as a member of the organization



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
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersPost(context.Background(), realm, id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

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


## AdminRealmsRealmOrganizationsIdMembersUserIdDelete

> AdminRealmsRealmOrganizationsIdMembersUserIdDelete(ctx, realm, id, userId).Execute()

Removes the user with the specified id from the organization



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdDelete(context.Background(), realm, id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdDelete``: %v\n", err)
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
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersUserIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsIdMembersUserIdGet

> MemberRepresentation AdminRealmsRealmOrganizationsIdMembersUserIdGet(ctx, realm, id, userId).Execute()

Returns the member of the organization with the specified id



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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdGet(context.Background(), realm, id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdMembersUserIdGet`: MemberRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MemberRepresentation**](MemberRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet

> []OrganizationRepresentation AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet(ctx, realm, id, userId).Execute()

Returns the organizations associated with the user that has the specified id

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
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet(context.Background(), realm, id, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet`: []OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdMembersUserIdOrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]OrganizationRepresentation**](OrganizationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsIdPut

> AdminRealmsRealmOrganizationsIdPut(ctx, realm, id).OrganizationRepresentation(organizationRepresentation).Execute()

Updates the organization

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
	organizationRepresentation := *openapiclient.NewOrganizationRepresentation() // OrganizationRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsIdPut(context.Background(), realm, id).OrganizationRepresentation(organizationRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsIdPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **organizationRepresentation** | [**OrganizationRepresentation**](OrganizationRepresentation.md) |  | 

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


## AdminRealmsRealmOrganizationsMembersIdOrganizationsGet

> []OrganizationRepresentation AdminRealmsRealmOrganizationsMembersIdOrganizationsGet(ctx, realm, id).Execute()

Returns the organizations associated with the user that has the specified id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsMembersIdOrganizationsGet(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsMembersIdOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsMembersIdOrganizationsGet`: []OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsMembersIdOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsMembersIdOrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]OrganizationRepresentation**](OrganizationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmOrganizationsPost

> AdminRealmsRealmOrganizationsPost(ctx, realm).OrganizationRepresentation(organizationRepresentation).Execute()

Creates a new organization

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
	organizationRepresentation := *openapiclient.NewOrganizationRepresentation() // OrganizationRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsPost(context.Background(), realm).OrganizationRepresentation(organizationRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationRepresentation** | [**OrganizationRepresentation**](OrganizationRepresentation.md) |  | 

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

