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

// checks if the RequiredActionConfigInfoRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequiredActionConfigInfoRepresentation{}

// RequiredActionConfigInfoRepresentation struct for RequiredActionConfigInfoRepresentation
type RequiredActionConfigInfoRepresentation struct {
	Properties []ConfigPropertyRepresentation `json:"properties,omitempty"`
}

// NewRequiredActionConfigInfoRepresentation instantiates a new RequiredActionConfigInfoRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequiredActionConfigInfoRepresentation() *RequiredActionConfigInfoRepresentation {
	this := RequiredActionConfigInfoRepresentation{}
	return &this
}

// NewRequiredActionConfigInfoRepresentationWithDefaults instantiates a new RequiredActionConfigInfoRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequiredActionConfigInfoRepresentationWithDefaults() *RequiredActionConfigInfoRepresentation {
	this := RequiredActionConfigInfoRepresentation{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *RequiredActionConfigInfoRepresentation) GetProperties() []ConfigPropertyRepresentation {
	if o == nil || IsNil(o.Properties) {
		var ret []ConfigPropertyRepresentation
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredActionConfigInfoRepresentation) GetPropertiesOk() ([]ConfigPropertyRepresentation, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *RequiredActionConfigInfoRepresentation) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ConfigPropertyRepresentation and assigns it to the Properties field.
func (o *RequiredActionConfigInfoRepresentation) SetProperties(v []ConfigPropertyRepresentation) {
	o.Properties = v
}

func (o RequiredActionConfigInfoRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequiredActionConfigInfoRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableRequiredActionConfigInfoRepresentation struct {
	value *RequiredActionConfigInfoRepresentation
	isSet bool
}

func (v NullableRequiredActionConfigInfoRepresentation) Get() *RequiredActionConfigInfoRepresentation {
	return v.value
}

func (v *NullableRequiredActionConfigInfoRepresentation) Set(val *RequiredActionConfigInfoRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredActionConfigInfoRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredActionConfigInfoRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredActionConfigInfoRepresentation(val *RequiredActionConfigInfoRepresentation) *NullableRequiredActionConfigInfoRepresentation {
	return &NullableRequiredActionConfigInfoRepresentation{value: val, isSet: true}
}

func (v NullableRequiredActionConfigInfoRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequiredActionConfigInfoRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
