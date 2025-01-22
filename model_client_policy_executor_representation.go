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

// checks if the ClientPolicyExecutorRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientPolicyExecutorRepresentation{}

// ClientPolicyExecutorRepresentation struct for ClientPolicyExecutorRepresentation
type ClientPolicyExecutorRepresentation struct {
	Executor      *string  `json:"executor,omitempty"`
	Configuration []string `json:"configuration,omitempty"`
}

// NewClientPolicyExecutorRepresentation instantiates a new ClientPolicyExecutorRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientPolicyExecutorRepresentation() *ClientPolicyExecutorRepresentation {
	this := ClientPolicyExecutorRepresentation{}
	return &this
}

// NewClientPolicyExecutorRepresentationWithDefaults instantiates a new ClientPolicyExecutorRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientPolicyExecutorRepresentationWithDefaults() *ClientPolicyExecutorRepresentation {
	this := ClientPolicyExecutorRepresentation{}
	return &this
}

// GetExecutor returns the Executor field value if set, zero value otherwise.
func (o *ClientPolicyExecutorRepresentation) GetExecutor() string {
	if o == nil || IsNil(o.Executor) {
		var ret string
		return ret
	}
	return *o.Executor
}

// GetExecutorOk returns a tuple with the Executor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyExecutorRepresentation) GetExecutorOk() (*string, bool) {
	if o == nil || IsNil(o.Executor) {
		return nil, false
	}
	return o.Executor, true
}

// HasExecutor returns a boolean if a field has been set.
func (o *ClientPolicyExecutorRepresentation) HasExecutor() bool {
	if o != nil && !IsNil(o.Executor) {
		return true
	}

	return false
}

// SetExecutor gets a reference to the given string and assigns it to the Executor field.
func (o *ClientPolicyExecutorRepresentation) SetExecutor(v string) {
	o.Executor = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ClientPolicyExecutorRepresentation) GetConfiguration() []string {
	if o == nil || IsNil(o.Configuration) {
		var ret []string
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPolicyExecutorRepresentation) GetConfigurationOk() ([]string, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ClientPolicyExecutorRepresentation) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given []string and assigns it to the Configuration field.
func (o *ClientPolicyExecutorRepresentation) SetConfiguration(v []string) {
	o.Configuration = v
}

func (o ClientPolicyExecutorRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientPolicyExecutorRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Executor) {
		toSerialize["executor"] = o.Executor
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableClientPolicyExecutorRepresentation struct {
	value *ClientPolicyExecutorRepresentation
	isSet bool
}

func (v NullableClientPolicyExecutorRepresentation) Get() *ClientPolicyExecutorRepresentation {
	return v.value
}

func (v *NullableClientPolicyExecutorRepresentation) Set(val *ClientPolicyExecutorRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableClientPolicyExecutorRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableClientPolicyExecutorRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientPolicyExecutorRepresentation(val *ClientPolicyExecutorRepresentation) *NullableClientPolicyExecutorRepresentation {
	return &NullableClientPolicyExecutorRepresentation{value: val, isSet: true}
}

func (v NullableClientPolicyExecutorRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientPolicyExecutorRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
