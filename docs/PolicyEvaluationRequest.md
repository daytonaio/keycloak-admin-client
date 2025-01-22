# PolicyEvaluationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]map[string]string** |  | [optional] 
**Resources** | Pointer to [**[]ResourceRepresentation**](ResourceRepresentation.md) |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**RoleIds** | Pointer to **[]string** |  | [optional] 
**Entitlements** | Pointer to **bool** |  | [optional] 

## Methods

### NewPolicyEvaluationRequest

`func NewPolicyEvaluationRequest() *PolicyEvaluationRequest`

NewPolicyEvaluationRequest instantiates a new PolicyEvaluationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationRequestWithDefaults

`func NewPolicyEvaluationRequestWithDefaults() *PolicyEvaluationRequest`

NewPolicyEvaluationRequestWithDefaults instantiates a new PolicyEvaluationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PolicyEvaluationRequest) GetContext() map[string]map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PolicyEvaluationRequest) GetContextOk() (*map[string]map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PolicyEvaluationRequest) SetContext(v map[string]map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PolicyEvaluationRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetResources

`func (o *PolicyEvaluationRequest) GetResources() []ResourceRepresentation`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PolicyEvaluationRequest) GetResourcesOk() (*[]ResourceRepresentation, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PolicyEvaluationRequest) SetResources(v []ResourceRepresentation)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PolicyEvaluationRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetClientId

`func (o *PolicyEvaluationRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PolicyEvaluationRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PolicyEvaluationRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PolicyEvaluationRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetUserId

`func (o *PolicyEvaluationRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PolicyEvaluationRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PolicyEvaluationRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PolicyEvaluationRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRoleIds

`func (o *PolicyEvaluationRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *PolicyEvaluationRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *PolicyEvaluationRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *PolicyEvaluationRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetEntitlements

`func (o *PolicyEvaluationRequest) GetEntitlements() bool`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *PolicyEvaluationRequest) GetEntitlementsOk() (*bool, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *PolicyEvaluationRequest) SetEntitlements(v bool)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *PolicyEvaluationRequest) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


