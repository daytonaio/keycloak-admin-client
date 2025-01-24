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

// checks if the ComponentTypeRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentTypeRepresentation{}

// ComponentTypeRepresentation struct for ComponentTypeRepresentation
type ComponentTypeRepresentation struct {
	Id *string `json:"id,omitempty"`
	HelpText *string `json:"helpText,omitempty"`
	Properties []ConfigPropertyRepresentation `json:"properties,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewComponentTypeRepresentation instantiates a new ComponentTypeRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentTypeRepresentation() *ComponentTypeRepresentation {
	this := ComponentTypeRepresentation{}
	return &this
}

// NewComponentTypeRepresentationWithDefaults instantiates a new ComponentTypeRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentTypeRepresentationWithDefaults() *ComponentTypeRepresentation {
	this := ComponentTypeRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ComponentTypeRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentTypeRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ComponentTypeRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ComponentTypeRepresentation) SetId(v string) {
	o.Id = &v
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *ComponentTypeRepresentation) GetHelpText() string {
	if o == nil || IsNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentTypeRepresentation) GetHelpTextOk() (*string, bool) {
	if o == nil || IsNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *ComponentTypeRepresentation) HasHelpText() bool {
	if o != nil && !IsNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *ComponentTypeRepresentation) SetHelpText(v string) {
	o.HelpText = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ComponentTypeRepresentation) GetProperties() []ConfigPropertyRepresentation {
	if o == nil || IsNil(o.Properties) {
		var ret []ConfigPropertyRepresentation
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentTypeRepresentation) GetPropertiesOk() ([]ConfigPropertyRepresentation, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ComponentTypeRepresentation) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ConfigPropertyRepresentation and assigns it to the Properties field.
func (o *ComponentTypeRepresentation) SetProperties(v []ConfigPropertyRepresentation) {
	o.Properties = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ComponentTypeRepresentation) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentTypeRepresentation) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ComponentTypeRepresentation) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ComponentTypeRepresentation) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ComponentTypeRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentTypeRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.HelpText) {
		toSerialize["helpText"] = o.HelpText
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableComponentTypeRepresentation struct {
	value *ComponentTypeRepresentation
	isSet bool
}

func (v NullableComponentTypeRepresentation) Get() *ComponentTypeRepresentation {
	return v.value
}

func (v *NullableComponentTypeRepresentation) Set(val *ComponentTypeRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentTypeRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentTypeRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentTypeRepresentation(val *ComponentTypeRepresentation) *NullableComponentTypeRepresentation {
	return &NullableComponentTypeRepresentation{value: val, isSet: true}
}

func (v NullableComponentTypeRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentTypeRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


