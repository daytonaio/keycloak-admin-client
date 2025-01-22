# PolicyRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Logic** | Pointer to [**Logic**](Logic.md) |  | [optional] 
**DecisionStrategy** | Pointer to [**DecisionStrategy**](DecisionStrategy.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**ResourcesData** | Pointer to [**[]ResourceRepresentation**](ResourceRepresentation.md) |  | [optional] 
**ScopesData** | Pointer to [**[]ScopeRepresentation**](ScopeRepresentation.md) |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPolicyRepresentation

`func NewPolicyRepresentation() *PolicyRepresentation`

NewPolicyRepresentation instantiates a new PolicyRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRepresentationWithDefaults

`func NewPolicyRepresentationWithDefaults() *PolicyRepresentation`

NewPolicyRepresentationWithDefaults instantiates a new PolicyRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PolicyRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PolicyRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPolicies

`func (o *PolicyRepresentation) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PolicyRepresentation) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PolicyRepresentation) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *PolicyRepresentation) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetResources

`func (o *PolicyRepresentation) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PolicyRepresentation) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PolicyRepresentation) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PolicyRepresentation) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetScopes

`func (o *PolicyRepresentation) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PolicyRepresentation) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PolicyRepresentation) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PolicyRepresentation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetLogic

`func (o *PolicyRepresentation) GetLogic() Logic`

GetLogic returns the Logic field if non-nil, zero value otherwise.

### GetLogicOk

`func (o *PolicyRepresentation) GetLogicOk() (*Logic, bool)`

GetLogicOk returns a tuple with the Logic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogic

`func (o *PolicyRepresentation) SetLogic(v Logic)`

SetLogic sets Logic field to given value.

### HasLogic

`func (o *PolicyRepresentation) HasLogic() bool`

HasLogic returns a boolean if a field has been set.

### GetDecisionStrategy

`func (o *PolicyRepresentation) GetDecisionStrategy() DecisionStrategy`

GetDecisionStrategy returns the DecisionStrategy field if non-nil, zero value otherwise.

### GetDecisionStrategyOk

`func (o *PolicyRepresentation) GetDecisionStrategyOk() (*DecisionStrategy, bool)`

GetDecisionStrategyOk returns a tuple with the DecisionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionStrategy

`func (o *PolicyRepresentation) SetDecisionStrategy(v DecisionStrategy)`

SetDecisionStrategy sets DecisionStrategy field to given value.

### HasDecisionStrategy

`func (o *PolicyRepresentation) HasDecisionStrategy() bool`

HasDecisionStrategy returns a boolean if a field has been set.

### GetOwner

`func (o *PolicyRepresentation) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PolicyRepresentation) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PolicyRepresentation) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PolicyRepresentation) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResourcesData

`func (o *PolicyRepresentation) GetResourcesData() []ResourceRepresentation`

GetResourcesData returns the ResourcesData field if non-nil, zero value otherwise.

### GetResourcesDataOk

`func (o *PolicyRepresentation) GetResourcesDataOk() (*[]ResourceRepresentation, bool)`

GetResourcesDataOk returns a tuple with the ResourcesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesData

`func (o *PolicyRepresentation) SetResourcesData(v []ResourceRepresentation)`

SetResourcesData sets ResourcesData field to given value.

### HasResourcesData

`func (o *PolicyRepresentation) HasResourcesData() bool`

HasResourcesData returns a boolean if a field has been set.

### GetScopesData

`func (o *PolicyRepresentation) GetScopesData() []ScopeRepresentation`

GetScopesData returns the ScopesData field if non-nil, zero value otherwise.

### GetScopesDataOk

`func (o *PolicyRepresentation) GetScopesDataOk() (*[]ScopeRepresentation, bool)`

GetScopesDataOk returns a tuple with the ScopesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesData

`func (o *PolicyRepresentation) SetScopesData(v []ScopeRepresentation)`

SetScopesData sets ScopesData field to given value.

### HasScopesData

`func (o *PolicyRepresentation) HasScopesData() bool`

HasScopesData returns a boolean if a field has been set.

### GetConfig

`func (o *PolicyRepresentation) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PolicyRepresentation) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PolicyRepresentation) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PolicyRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


