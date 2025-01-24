# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmOrganizationsGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsGet) | **Get** /admin/realms/{realm}/organizations | Returns a paginated list of organizations filtered according to the specified parameters
[**AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet) | **Get** /admin/realms/{realm}/organizations/members/{member-id}/organizations | Returns the organizations associated with the user that has the specified id
[**AdminRealmsRealmOrganizationsOrgIdDelete**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdDelete) | **Delete** /admin/realms/{realm}/organizations/{org-id} | Deletes the organization
[**AdminRealmsRealmOrganizationsOrgIdGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdGet) | **Get** /admin/realms/{realm}/organizations/{org-id} | Returns the organization representation
[**AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDelete**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDelete) | **Delete** /admin/realms/{realm}/organizations/{org-id}/identity-providers/{alias} | Removes the identity provider with the specified alias from the organization
[**AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet) | **Get** /admin/realms/{realm}/organizations/{org-id}/identity-providers/{alias} | Returns the identity provider associated with the organization that has the specified alias
[**AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet) | **Get** /admin/realms/{realm}/organizations/{org-id}/identity-providers | Returns all identity providers associated with the organization
[**AdminRealmsRealmOrganizationsOrgIdIdentityProvidersPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdIdentityProvidersPost) | **Post** /admin/realms/{realm}/organizations/{org-id}/identity-providers | Adds the identity provider with the specified id to the organization
[**AdminRealmsRealmOrganizationsOrgIdMembersCountGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersCountGet) | **Get** /admin/realms/{realm}/organizations/{org-id}/members/count | Returns number of members in the organization.
[**AdminRealmsRealmOrganizationsOrgIdMembersGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersGet) | **Get** /admin/realms/{realm}/organizations/{org-id}/members | Returns a paginated list of organization members filtered according to the specified parameters
[**AdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPost) | **Post** /admin/realms/{realm}/organizations/{org-id}/members/invite-existing-user | Invites an existing user to the organization, using the specified user id
[**AdminRealmsRealmOrganizationsOrgIdMembersInviteUserPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersInviteUserPost) | **Post** /admin/realms/{realm}/organizations/{org-id}/members/invite-user | Invites an existing user or sends a registration link to a new user, based on the provided e-mail address.
[**AdminRealmsRealmOrganizationsOrgIdMembersMemberIdDelete**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersMemberIdDelete) | **Delete** /admin/realms/{realm}/organizations/{org-id}/members/{member-id} | Removes the user with the specified id from the organization
[**AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet) | **Get** /admin/realms/{realm}/organizations/{org-id}/members/{member-id} | Returns the member of the organization with the specified id
[**AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet) | **Get** /admin/realms/{realm}/organizations/{org-id}/members/{member-id}/organizations | Returns the organizations associated with the user that has the specified id
[**AdminRealmsRealmOrganizationsOrgIdMembersPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdMembersPost) | **Post** /admin/realms/{realm}/organizations/{org-id}/members | Adds the user with the specified id as a member of the organization
[**AdminRealmsRealmOrganizationsOrgIdPut**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsOrgIdPut) | **Put** /admin/realms/{realm}/organizations/{org-id} | Updates the organization
[**AdminRealmsRealmOrganizationsPost**](OrganizationsAPI.md#AdminRealmsRealmOrganizationsPost) | **Post** /admin/realms/{realm}/organizations | Creates a new organization



## AdminRealmsRealmOrganizationsGet

> []OrganizationRepresentation AdminRealmsRealmOrganizationsGet(ctx, realm).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).Q(q).Search(search).Execute()

Returns a paginated list of organizations filtered according to the specified parameters

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
	briefRepresentation := true // bool | if true, return the full representation. Otherwise, only the basic fields are returned. (optional) (default to false)
	exact := true // bool | Boolean which defines whether the param 'search' must match exactly or not (optional)
	first := int32(56) // int32 | The position of the first result to be processed (pagination offset) (optional)
	max := int32(56) // int32 | The maximum number of results to be returned - defaults to 10 (optional)
	q := "q_example" // string | A query to search for custom attributes, in the format 'key1:value2 key2:value2' (optional)
	search := "search_example" // string | A String representing either an organization name or domain (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsGet(context.Background(), realm).BriefRepresentation(briefRepresentation).Exact(exact).First(first).Max(max).Q(q).Search(search).Execute()
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

 **briefRepresentation** | **bool** | if true, return the full representation. Otherwise, only the basic fields are returned. | [default to false]
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


## AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet

> []OrganizationRepresentation AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet(ctx, realm, memberId).Execute()

Returns the organizations associated with the user that has the specified id

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
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet(context.Background(), realm, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet`: []OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsMembersMemberIdOrganizationsGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdDelete

> AdminRealmsRealmOrganizationsOrgIdDelete(ctx, realm, orgId).Execute()

Deletes the organization

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdDelete(context.Background(), realm, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdGet

> OrganizationRepresentation AdminRealmsRealmOrganizationsOrgIdGet(ctx, realm, orgId).Execute()

Returns the organization representation

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdGet(context.Background(), realm, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdGet`: OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDelete

> AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDelete(ctx, realm, orgId, alias).Execute()

Removes the identity provider with the specified alias from the organization



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
	orgId := "orgId_example" // string | 
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDelete(context.Background(), realm, orgId, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet

> IdentityProviderRepresentation AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet(ctx, realm, orgId, alias).Execute()

Returns the identity provider associated with the organization that has the specified alias



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
	orgId := "orgId_example" // string | 
	alias := "alias_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet(context.Background(), realm, orgId, alias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet`: IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 
**alias** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdIdentityProvidersAliasGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet

> []IdentityProviderRepresentation AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet(ctx, realm, orgId).Execute()

Returns all identity providers associated with the organization

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet(context.Background(), realm, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet`: []IdentityProviderRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdIdentityProvidersGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdIdentityProvidersPost

> AdminRealmsRealmOrganizationsOrgIdIdentityProvidersPost(ctx, realm, orgId).Body(body).Execute()

Adds the identity provider with the specified id to the organization



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
	orgId := "orgId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersPost(context.Background(), realm, orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdIdentityProvidersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdIdentityProvidersPostRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdMembersCountGet

> int64 AdminRealmsRealmOrganizationsOrgIdMembersCountGet(ctx, realm, orgId).Execute()

Returns number of members in the organization.

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersCountGet(context.Background(), realm, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdMembersCountGet`: int64
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersCountGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdMembersGet

> []MemberRepresentation AdminRealmsRealmOrganizationsOrgIdMembersGet(ctx, realm, orgId).Exact(exact).First(first).Max(max).MembershipType(membershipType).Search(search).Execute()

Returns a paginated list of organization members filtered according to the specified parameters

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
	orgId := "orgId_example" // string | 
	exact := true // bool | Boolean which defines whether the param 'search' must match exactly or not (optional)
	first := int32(56) // int32 | The position of the first result to be processed (pagination offset) (optional)
	max := int32(56) // int32 | The maximum number of results to be returned. Defaults to 10 (optional)
	membershipType := "membershipType_example" // string | The membership type (optional)
	search := "search_example" // string | A String representing either a member's username, e-mail, first name, or last name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersGet(context.Background(), realm, orgId).Exact(exact).First(first).Max(max).MembershipType(membershipType).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdMembersGet`: []MemberRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Boolean which defines whether the param &#39;search&#39; must match exactly or not | 
 **first** | **int32** | The position of the first result to be processed (pagination offset) | 
 **max** | **int32** | The maximum number of results to be returned. Defaults to 10 | 
 **membershipType** | **string** | The membership type | 
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


## AdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPost

> AdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPost(ctx, realm, orgId).Id(id).Execute()

Invites an existing user to the organization, using the specified user id

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
	orgId := "orgId_example" // string | 
	id := "id_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPost(context.Background(), realm, orgId).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersInviteExistingUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **string** |  | 

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


## AdminRealmsRealmOrganizationsOrgIdMembersInviteUserPost

> AdminRealmsRealmOrganizationsOrgIdMembersInviteUserPost(ctx, realm, orgId).Email(email).FirstName(firstName).LastName(lastName).Execute()

Invites an existing user or sends a registration link to a new user, based on the provided e-mail address.



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
	orgId := "orgId_example" // string | 
	email := "email_example" // string |  (optional)
	firstName := "firstName_example" // string |  (optional)
	lastName := "lastName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersInviteUserPost(context.Background(), realm, orgId).Email(email).FirstName(firstName).LastName(lastName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersInviteUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersInviteUserPostRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdMembersMemberIdDelete

> AdminRealmsRealmOrganizationsOrgIdMembersMemberIdDelete(ctx, realm, orgId, memberId).Execute()

Removes the user with the specified id from the organization



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
	orgId := "orgId_example" // string | 
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdDelete(context.Background(), realm, orgId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersMemberIdDeleteRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet

> MemberRepresentation AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet(ctx, realm, orgId, memberId).Execute()

Returns the member of the organization with the specified id



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
	orgId := "orgId_example" // string | 
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet(context.Background(), realm, orgId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet`: MemberRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersMemberIdGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet

> []OrganizationRepresentation AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet(ctx, realm, orgId, memberId).Execute()

Returns the organizations associated with the user that has the specified id

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
	orgId := "orgId_example" // string | 
	memberId := "memberId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet(context.Background(), realm, orgId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet`: []OrganizationRepresentation
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersMemberIdOrganizationsGetRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdMembersPost

> AdminRealmsRealmOrganizationsOrgIdMembersPost(ctx, realm, orgId).Body(body).Execute()

Adds the user with the specified id as a member of the organization



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
	orgId := "orgId_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersPost(context.Background(), realm, orgId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdMembersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdMembersPostRequest struct via the builder pattern


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


## AdminRealmsRealmOrganizationsOrgIdPut

> AdminRealmsRealmOrganizationsOrgIdPut(ctx, realm, orgId).OrganizationRepresentation(organizationRepresentation).Execute()

Updates the organization

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
	orgId := "orgId_example" // string | 
	organizationRepresentation := *openapiclient.NewOrganizationRepresentation() // OrganizationRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdPut(context.Background(), realm, orgId).OrganizationRepresentation(organizationRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AdminRealmsRealmOrganizationsOrgIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmOrganizationsOrgIdPutRequest struct via the builder pattern


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
	openapiclient "github.com/daytonaio/keycloak-admin-client"
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

