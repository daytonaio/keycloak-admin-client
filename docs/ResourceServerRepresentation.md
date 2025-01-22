# ResourceServerRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AllowRemoteResourceManagement** | Pointer to **bool** |  | [optional] 
**PolicyEnforcementMode** | Pointer to [**PolicyEnforcementMode**](PolicyEnforcementMode.md) |  | [optional] 
**Resources** | Pointer to [**[]ResourceRepresentation**](ResourceRepresentation.md) |  | [optional] 
**Policies** | Pointer to [**[]PolicyRepresentation**](PolicyRepresentation.md) |  | [optional] 
**Scopes** | Pointer to [**[]ScopeRepresentation**](ScopeRepresentation.md) |  | [optional] 
**DecisionStrategy** | Pointer to [**DecisionStrategy**](DecisionStrategy.md) |  | [optional] 

## Methods

### NewResourceServerRepresentation

`func NewResourceServerRepresentation() *ResourceServerRepresentation`

NewResourceServerRepresentation instantiates a new ResourceServerRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceServerRepresentationWithDefaults

`func NewResourceServerRepresentationWithDefaults() *ResourceServerRepresentation`

NewResourceServerRepresentationWithDefaults instantiates a new ResourceServerRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceServerRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceServerRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceServerRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceServerRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *ResourceServerRepresentation) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ResourceServerRepresentation) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ResourceServerRepresentation) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ResourceServerRepresentation) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetName

`func (o *ResourceServerRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceServerRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceServerRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceServerRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAllowRemoteResourceManagement

`func (o *ResourceServerRepresentation) GetAllowRemoteResourceManagement() bool`

GetAllowRemoteResourceManagement returns the AllowRemoteResourceManagement field if non-nil, zero value otherwise.

### GetAllowRemoteResourceManagementOk

`func (o *ResourceServerRepresentation) GetAllowRemoteResourceManagementOk() (*bool, bool)`

GetAllowRemoteResourceManagementOk returns a tuple with the AllowRemoteResourceManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteResourceManagement

`func (o *ResourceServerRepresentation) SetAllowRemoteResourceManagement(v bool)`

SetAllowRemoteResourceManagement sets AllowRemoteResourceManagement field to given value.

### HasAllowRemoteResourceManagement

`func (o *ResourceServerRepresentation) HasAllowRemoteResourceManagement() bool`

HasAllowRemoteResourceManagement returns a boolean if a field has been set.

### GetPolicyEnforcementMode

`func (o *ResourceServerRepresentation) GetPolicyEnforcementMode() PolicyEnforcementMode`

GetPolicyEnforcementMode returns the PolicyEnforcementMode field if non-nil, zero value otherwise.

### GetPolicyEnforcementModeOk

`func (o *ResourceServerRepresentation) GetPolicyEnforcementModeOk() (*PolicyEnforcementMode, bool)`

GetPolicyEnforcementModeOk returns a tuple with the PolicyEnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnforcementMode

`func (o *ResourceServerRepresentation) SetPolicyEnforcementMode(v PolicyEnforcementMode)`

SetPolicyEnforcementMode sets PolicyEnforcementMode field to given value.

### HasPolicyEnforcementMode

`func (o *ResourceServerRepresentation) HasPolicyEnforcementMode() bool`

HasPolicyEnforcementMode returns a boolean if a field has been set.

### GetResources

`func (o *ResourceServerRepresentation) GetResources() []ResourceRepresentation`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ResourceServerRepresentation) GetResourcesOk() (*[]ResourceRepresentation, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ResourceServerRepresentation) SetResources(v []ResourceRepresentation)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ResourceServerRepresentation) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetPolicies

`func (o *ResourceServerRepresentation) GetPolicies() []PolicyRepresentation`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ResourceServerRepresentation) GetPoliciesOk() (*[]PolicyRepresentation, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ResourceServerRepresentation) SetPolicies(v []PolicyRepresentation)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ResourceServerRepresentation) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetScopes

`func (o *ResourceServerRepresentation) GetScopes() []ScopeRepresentation`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ResourceServerRepresentation) GetScopesOk() (*[]ScopeRepresentation, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ResourceServerRepresentation) SetScopes(v []ScopeRepresentation)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ResourceServerRepresentation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetDecisionStrategy

`func (o *ResourceServerRepresentation) GetDecisionStrategy() DecisionStrategy`

GetDecisionStrategy returns the DecisionStrategy field if non-nil, zero value otherwise.

### GetDecisionStrategyOk

`func (o *ResourceServerRepresentation) GetDecisionStrategyOk() (*DecisionStrategy, bool)`

GetDecisionStrategyOk returns a tuple with the DecisionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionStrategy

`func (o *ResourceServerRepresentation) SetDecisionStrategy(v DecisionStrategy)`

SetDecisionStrategy sets DecisionStrategy field to given value.

### HasDecisionStrategy

`func (o *ResourceServerRepresentation) HasDecisionStrategy() bool`

HasDecisionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


