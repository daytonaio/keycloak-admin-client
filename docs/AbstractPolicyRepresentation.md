# AbstractPolicyRepresentation

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

## Methods

### NewAbstractPolicyRepresentation

`func NewAbstractPolicyRepresentation() *AbstractPolicyRepresentation`

NewAbstractPolicyRepresentation instantiates a new AbstractPolicyRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractPolicyRepresentationWithDefaults

`func NewAbstractPolicyRepresentationWithDefaults() *AbstractPolicyRepresentation`

NewAbstractPolicyRepresentationWithDefaults instantiates a new AbstractPolicyRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AbstractPolicyRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractPolicyRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractPolicyRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractPolicyRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AbstractPolicyRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractPolicyRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractPolicyRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractPolicyRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AbstractPolicyRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AbstractPolicyRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AbstractPolicyRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AbstractPolicyRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AbstractPolicyRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbstractPolicyRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbstractPolicyRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AbstractPolicyRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPolicies

`func (o *AbstractPolicyRepresentation) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AbstractPolicyRepresentation) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AbstractPolicyRepresentation) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AbstractPolicyRepresentation) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetResources

`func (o *AbstractPolicyRepresentation) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AbstractPolicyRepresentation) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AbstractPolicyRepresentation) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AbstractPolicyRepresentation) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetScopes

`func (o *AbstractPolicyRepresentation) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AbstractPolicyRepresentation) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AbstractPolicyRepresentation) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AbstractPolicyRepresentation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetLogic

`func (o *AbstractPolicyRepresentation) GetLogic() Logic`

GetLogic returns the Logic field if non-nil, zero value otherwise.

### GetLogicOk

`func (o *AbstractPolicyRepresentation) GetLogicOk() (*Logic, bool)`

GetLogicOk returns a tuple with the Logic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogic

`func (o *AbstractPolicyRepresentation) SetLogic(v Logic)`

SetLogic sets Logic field to given value.

### HasLogic

`func (o *AbstractPolicyRepresentation) HasLogic() bool`

HasLogic returns a boolean if a field has been set.

### GetDecisionStrategy

`func (o *AbstractPolicyRepresentation) GetDecisionStrategy() DecisionStrategy`

GetDecisionStrategy returns the DecisionStrategy field if non-nil, zero value otherwise.

### GetDecisionStrategyOk

`func (o *AbstractPolicyRepresentation) GetDecisionStrategyOk() (*DecisionStrategy, bool)`

GetDecisionStrategyOk returns a tuple with the DecisionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionStrategy

`func (o *AbstractPolicyRepresentation) SetDecisionStrategy(v DecisionStrategy)`

SetDecisionStrategy sets DecisionStrategy field to given value.

### HasDecisionStrategy

`func (o *AbstractPolicyRepresentation) HasDecisionStrategy() bool`

HasDecisionStrategy returns a boolean if a field has been set.

### GetOwner

`func (o *AbstractPolicyRepresentation) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AbstractPolicyRepresentation) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AbstractPolicyRepresentation) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AbstractPolicyRepresentation) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetResourcesData

`func (o *AbstractPolicyRepresentation) GetResourcesData() []ResourceRepresentation`

GetResourcesData returns the ResourcesData field if non-nil, zero value otherwise.

### GetResourcesDataOk

`func (o *AbstractPolicyRepresentation) GetResourcesDataOk() (*[]ResourceRepresentation, bool)`

GetResourcesDataOk returns a tuple with the ResourcesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesData

`func (o *AbstractPolicyRepresentation) SetResourcesData(v []ResourceRepresentation)`

SetResourcesData sets ResourcesData field to given value.

### HasResourcesData

`func (o *AbstractPolicyRepresentation) HasResourcesData() bool`

HasResourcesData returns a boolean if a field has been set.

### GetScopesData

`func (o *AbstractPolicyRepresentation) GetScopesData() []ScopeRepresentation`

GetScopesData returns the ScopesData field if non-nil, zero value otherwise.

### GetScopesDataOk

`func (o *AbstractPolicyRepresentation) GetScopesDataOk() (*[]ScopeRepresentation, bool)`

GetScopesDataOk returns a tuple with the ScopesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesData

`func (o *AbstractPolicyRepresentation) SetScopesData(v []ScopeRepresentation)`

SetScopesData sets ScopesData field to given value.

### HasScopesData

`func (o *AbstractPolicyRepresentation) HasScopesData() bool`

HasScopesData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


