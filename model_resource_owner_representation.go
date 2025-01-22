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

// checks if the ResourceOwnerRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceOwnerRepresentation{}

// ResourceOwnerRepresentation struct for ResourceOwnerRepresentation
type ResourceOwnerRepresentation struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewResourceOwnerRepresentation instantiates a new ResourceOwnerRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnerRepresentation() *ResourceOwnerRepresentation {
	this := ResourceOwnerRepresentation{}
	return &this
}

// NewResourceOwnerRepresentationWithDefaults instantiates a new ResourceOwnerRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnerRepresentationWithDefaults() *ResourceOwnerRepresentation {
	this := ResourceOwnerRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceOwnerRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnerRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceOwnerRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceOwnerRepresentation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceOwnerRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceOwnerRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceOwnerRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceOwnerRepresentation) SetName(v string) {
	o.Name = &v
}

func (o ResourceOwnerRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceOwnerRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableResourceOwnerRepresentation struct {
	value *ResourceOwnerRepresentation
	isSet bool
}

func (v NullableResourceOwnerRepresentation) Get() *ResourceOwnerRepresentation {
	return v.value
}

func (v *NullableResourceOwnerRepresentation) Set(val *ResourceOwnerRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnerRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnerRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnerRepresentation(val *ResourceOwnerRepresentation) *NullableResourceOwnerRepresentation {
	return &NullableResourceOwnerRepresentation{value: val, isSet: true}
}

func (v NullableResourceOwnerRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnerRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
