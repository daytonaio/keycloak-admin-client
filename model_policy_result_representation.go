/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keycloak_admin_client

import (
	"encoding/json"
)

// checks if the PolicyResultRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyResultRepresentation{}

// PolicyResultRepresentation struct for PolicyResultRepresentation
type PolicyResultRepresentation struct {
	Policy *PolicyRepresentation `json:"policy,omitempty"`
	Status *DecisionEffect `json:"status,omitempty"`
	AssociatedPolicies []PolicyResultRepresentation `json:"associatedPolicies,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}

// NewPolicyResultRepresentation instantiates a new PolicyResultRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyResultRepresentation() *PolicyResultRepresentation {
	this := PolicyResultRepresentation{}
	return &this
}

// NewPolicyResultRepresentationWithDefaults instantiates a new PolicyResultRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyResultRepresentationWithDefaults() *PolicyResultRepresentation {
	this := PolicyResultRepresentation{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PolicyResultRepresentation) GetPolicy() PolicyRepresentation {
	if o == nil || IsNil(o.Policy) {
		var ret PolicyRepresentation
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResultRepresentation) GetPolicyOk() (*PolicyRepresentation, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PolicyResultRepresentation) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given PolicyRepresentation and assigns it to the Policy field.
func (o *PolicyResultRepresentation) SetPolicy(v PolicyRepresentation) {
	o.Policy = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PolicyResultRepresentation) GetStatus() DecisionEffect {
	if o == nil || IsNil(o.Status) {
		var ret DecisionEffect
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResultRepresentation) GetStatusOk() (*DecisionEffect, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PolicyResultRepresentation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DecisionEffect and assigns it to the Status field.
func (o *PolicyResultRepresentation) SetStatus(v DecisionEffect) {
	o.Status = &v
}

// GetAssociatedPolicies returns the AssociatedPolicies field value if set, zero value otherwise.
func (o *PolicyResultRepresentation) GetAssociatedPolicies() []PolicyResultRepresentation {
	if o == nil || IsNil(o.AssociatedPolicies) {
		var ret []PolicyResultRepresentation
		return ret
	}
	return o.AssociatedPolicies
}

// GetAssociatedPoliciesOk returns a tuple with the AssociatedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResultRepresentation) GetAssociatedPoliciesOk() ([]PolicyResultRepresentation, bool) {
	if o == nil || IsNil(o.AssociatedPolicies) {
		return nil, false
	}
	return o.AssociatedPolicies, true
}

// HasAssociatedPolicies returns a boolean if a field has been set.
func (o *PolicyResultRepresentation) HasAssociatedPolicies() bool {
	if o != nil && !IsNil(o.AssociatedPolicies) {
		return true
	}

	return false
}

// SetAssociatedPolicies gets a reference to the given []PolicyResultRepresentation and assigns it to the AssociatedPolicies field.
func (o *PolicyResultRepresentation) SetAssociatedPolicies(v []PolicyResultRepresentation) {
	o.AssociatedPolicies = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *PolicyResultRepresentation) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResultRepresentation) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *PolicyResultRepresentation) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *PolicyResultRepresentation) SetScopes(v []string) {
	o.Scopes = v
}

func (o PolicyResultRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyResultRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.AssociatedPolicies) {
		toSerialize["associatedPolicies"] = o.AssociatedPolicies
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullablePolicyResultRepresentation struct {
	value *PolicyResultRepresentation
	isSet bool
}

func (v NullablePolicyResultRepresentation) Get() *PolicyResultRepresentation {
	return v.value
}

func (v *NullablePolicyResultRepresentation) Set(val *PolicyResultRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyResultRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyResultRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyResultRepresentation(val *PolicyResultRepresentation) *NullablePolicyResultRepresentation {
	return &NullablePolicyResultRepresentation{value: val, isSet: true}
}

func (v NullablePolicyResultRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyResultRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


