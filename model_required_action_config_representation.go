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

// checks if the RequiredActionConfigRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequiredActionConfigRepresentation{}

// RequiredActionConfigRepresentation struct for RequiredActionConfigRepresentation
type RequiredActionConfigRepresentation struct {
	Config *map[string]string `json:"config,omitempty"`
}

// NewRequiredActionConfigRepresentation instantiates a new RequiredActionConfigRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequiredActionConfigRepresentation() *RequiredActionConfigRepresentation {
	this := RequiredActionConfigRepresentation{}
	return &this
}

// NewRequiredActionConfigRepresentationWithDefaults instantiates a new RequiredActionConfigRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequiredActionConfigRepresentationWithDefaults() *RequiredActionConfigRepresentation {
	this := RequiredActionConfigRepresentation{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *RequiredActionConfigRepresentation) GetConfig() map[string]string {
	if o == nil || IsNil(o.Config) {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequiredActionConfigRepresentation) GetConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *RequiredActionConfigRepresentation) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *RequiredActionConfigRepresentation) SetConfig(v map[string]string) {
	o.Config = &v
}

func (o RequiredActionConfigRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequiredActionConfigRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableRequiredActionConfigRepresentation struct {
	value *RequiredActionConfigRepresentation
	isSet bool
}

func (v NullableRequiredActionConfigRepresentation) Get() *RequiredActionConfigRepresentation {
	return v.value
}

func (v *NullableRequiredActionConfigRepresentation) Set(val *RequiredActionConfigRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableRequiredActionConfigRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableRequiredActionConfigRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequiredActionConfigRepresentation(val *RequiredActionConfigRepresentation) *NullableRequiredActionConfigRepresentation {
	return &NullableRequiredActionConfigRepresentation{value: val, isSet: true}
}

func (v NullableRequiredActionConfigRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequiredActionConfigRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


