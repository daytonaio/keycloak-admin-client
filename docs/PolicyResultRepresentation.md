# PolicyResultRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**PolicyRepresentation**](PolicyRepresentation.md) |  | [optional] 
**Status** | Pointer to [**DecisionEffect**](DecisionEffect.md) |  | [optional] 
**AssociatedPolicies** | Pointer to [**[]PolicyResultRepresentation**](PolicyResultRepresentation.md) |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyResultRepresentation

`func NewPolicyResultRepresentation() *PolicyResultRepresentation`

NewPolicyResultRepresentation instantiates a new PolicyResultRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResultRepresentationWithDefaults

`func NewPolicyResultRepresentationWithDefaults() *PolicyResultRepresentation`

NewPolicyResultRepresentationWithDefaults instantiates a new PolicyResultRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *PolicyResultRepresentation) GetPolicy() PolicyRepresentation`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PolicyResultRepresentation) GetPolicyOk() (*PolicyRepresentation, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PolicyResultRepresentation) SetPolicy(v PolicyRepresentation)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PolicyResultRepresentation) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetStatus

`func (o *PolicyResultRepresentation) GetStatus() DecisionEffect`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyResultRepresentation) GetStatusOk() (*DecisionEffect, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyResultRepresentation) SetStatus(v DecisionEffect)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PolicyResultRepresentation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAssociatedPolicies

`func (o *PolicyResultRepresentation) GetAssociatedPolicies() []PolicyResultRepresentation`

GetAssociatedPolicies returns the AssociatedPolicies field if non-nil, zero value otherwise.

### GetAssociatedPoliciesOk

`func (o *PolicyResultRepresentation) GetAssociatedPoliciesOk() (*[]PolicyResultRepresentation, bool)`

GetAssociatedPoliciesOk returns a tuple with the AssociatedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPolicies

`func (o *PolicyResultRepresentation) SetAssociatedPolicies(v []PolicyResultRepresentation)`

SetAssociatedPolicies sets AssociatedPolicies field to given value.

### HasAssociatedPolicies

`func (o *PolicyResultRepresentation) HasAssociatedPolicies() bool`

HasAssociatedPolicies returns a boolean if a field has been set.

### GetScopes

`func (o *PolicyResultRepresentation) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PolicyResultRepresentation) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PolicyResultRepresentation) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PolicyResultRepresentation) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


