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

// checks if the ClientPolicyRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientPolicyRepresentation{}

// ClientPolicyRepresentation struct for ClientPolicyRepresentation
type ClientPolicyRepresentation struct {
	Name        *string                               `json:"name,omitempty"`
	Description *string                               `json:"description,omitempty"`
	Enabled     *bool                                 `json:"enabled,omitempty"`
	Conditions  []ClientPolicyConditionRepresentation `json:"conditions,omitempty"`
	Profiles    []string                              `json:"profiles,omitempty"`
}

// NewClientPolicyRepresentation instantiates a new ClientPolicyRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientPolicyRepresentation() *ClientPolicyRepresentation {
	this := ClientPolicyRepresentation{}
	return &this
}

// NewClientPolicyRepresentationWithDefaults instantiates a new ClientPolicyRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientPolicyRepresentationWithDefaults() *ClientPolicyRepresentation {
	this := ClientPolicyRepresentation{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientPolicyRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientPolicyRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientPolicyRepresentation) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClientPolicyRepresentation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyRepresentation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClientPolicyRepresentation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClientPolicyRepresentation) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ClientPolicyRepresentation) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyRepresentation) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClientPolicyRepresentation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClientPolicyRepresentation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ClientPolicyRepresentation) GetConditions() []ClientPolicyConditionRepresentation {
	if o == nil || IsNil(o.Conditions) {
		var ret []ClientPolicyConditionRepresentation
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyRepresentation) GetConditionsOk() ([]ClientPolicyConditionRepresentation, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ClientPolicyRepresentation) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ClientPolicyConditionRepresentation and assigns it to the Conditions field.
func (o *ClientPolicyRepresentation) SetConditions(v []ClientPolicyConditionRepresentation) {
	o.Conditions = v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *ClientPolicyRepresentation) GetProfiles() []string {
	if o == nil || IsNil(o.Profiles) {
		var ret []string
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyRepresentation) GetProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ClientPolicyRepresentation) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []string and assigns it to the Profiles field.
func (o *ClientPolicyRepresentation) SetProfiles(v []string) {
	o.Profiles = v
}

func (o ClientPolicyRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientPolicyRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	return toSerialize, nil
}

type NullableClientPolicyRepresentation struct {
	value *ClientPolicyRepresentation
	isSet bool
}

func (v NullableClientPolicyRepresentation) Get() *ClientPolicyRepresentation {
	return v.value
}

func (v *NullableClientPolicyRepresentation) Set(val *ClientPolicyRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableClientPolicyRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableClientPolicyRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientPolicyRepresentation(val *ClientPolicyRepresentation) *NullableClientPolicyRepresentation {
	return &NullableClientPolicyRepresentation{value: val, isSet: true}
}

func (v NullableClientPolicyRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientPolicyRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
