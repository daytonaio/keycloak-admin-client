# EvaluationResultRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**ResourceRepresentation**](ResourceRepresentation.md) |  | [optional] 
**Scopes** | Pointer to [**[]ScopeRepresentation**](ScopeRepresentation.md) |  | [optional] 
**Policies** | Pointer to [**[]PolicyResultRepresentation**](PolicyResultRepresentation.md) |  | [optional] 
**Status** | Pointer to [**DecisionEffect**](DecisionEffect.md) |  | [optional] 
**AllowedScopes** | Pointer to [**[]ScopeRepresentation**](ScopeRepresentation.md) |  | [optional] 

## Methods

### NewEvaluationResultRepresentation

`func NewEvaluationResultRepresentation() *EvaluationResultRepresentation`

NewEvaluationResultRepresentation instantiates a new EvaluationResultRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationResultRepresentationWithDefaults

`func NewEvaluationResultRepresentationWithDefaults() *EvaluationResultRepresentation`

NewEvaluationResultRepresentationWithDefaults instantiates a new EvaluationResultRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *EvaluationResultRepresentation) GetResource() ResourceRepresentation`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EvaluationResultRepresentation) GetResourceOk() (*ResourceRepresentation, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EvaluationResultRepresentation) SetResource(v ResourceRepresentation)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EvaluationResultRepresentation) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetScopes

`func (o *EvaluationResultRepresentation) GetScopes() []ScopeRepresentation`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EvaluationResultRepresentation) GetScopesOk() (*[]ScopeRepresentation, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EvaluationResultRepresentation) SetScopes(v []ScopeRepresentation)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *EvaluationResultRepresentation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetPolicies

`func (o *EvaluationResultRepresentation) GetPolicies() []PolicyResultRepresentation`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *EvaluationResultRepresentation) GetPoliciesOk() (*[]PolicyResultRepresentation, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *EvaluationResultRepresentation) SetPolicies(v []PolicyResultRepresentation)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *EvaluationResultRepresentation) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetStatus

`func (o *EvaluationResultRepresentation) GetStatus() DecisionEffect`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationResultRepresentation) GetStatusOk() (*DecisionEffect, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationResultRepresentation) SetStatus(v DecisionEffect)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EvaluationResultRepresentation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAllowedScopes

`func (o *EvaluationResultRepresentation) GetAllowedScopes() []ScopeRepresentation`

GetAllowedScopes returns the AllowedScopes field if non-nil, zero value otherwise.

### GetAllowedScopesOk

`func (o *EvaluationResultRepresentation) GetAllowedScopesOk() (*[]ScopeRepresentation, bool)`

GetAllowedScopesOk returns a tuple with the AllowedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScopes

`func (o *EvaluationResultRepresentation) SetAllowedScopes(v []ScopeRepresentation)`

SetAllowedScopes sets AllowedScopes field to given value.

### HasAllowedScopes

`func (o *EvaluationResultRepresentation) HasAllowedScopes() bool`

HasAllowedScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


