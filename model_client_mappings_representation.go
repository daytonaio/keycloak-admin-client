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

// checks if the ClientMappingsRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientMappingsRepresentation{}

// ClientMappingsRepresentation struct for ClientMappingsRepresentation
type ClientMappingsRepresentation struct {
	Id       *string              `json:"id,omitempty"`
	Client   *string              `json:"client,omitempty"`
	Mappings []RoleRepresentation `json:"mappings,omitempty"`
}

// NewClientMappingsRepresentation instantiates a new ClientMappingsRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientMappingsRepresentation() *ClientMappingsRepresentation {
	this := ClientMappingsRepresentation{}
	return &this
}

// NewClientMappingsRepresentationWithDefaults instantiates a new ClientMappingsRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientMappingsRepresentationWithDefaults() *ClientMappingsRepresentation {
	this := ClientMappingsRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientMappingsRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMappingsRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientMappingsRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientMappingsRepresentation) SetId(v string) {
	o.Id = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ClientMappingsRepresentation) GetClient() string {
	if o == nil || IsNil(o.Client) {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMappingsRepresentation) GetClientOk() (*string, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ClientMappingsRepresentation) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *ClientMappingsRepresentation) SetClient(v string) {
	o.Client = &v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *ClientMappingsRepresentation) GetMappings() []RoleRepresentation {
	if o == nil || IsNil(o.Mappings) {
		var ret []RoleRepresentation
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientMappingsRepresentation) GetMappingsOk() ([]RoleRepresentation, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ClientMappingsRepresentation) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []RoleRepresentation and assigns it to the Mappings field.
func (o *ClientMappingsRepresentation) SetMappings(v []RoleRepresentation) {
	o.Mappings = v
}

func (o ClientMappingsRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientMappingsRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Mappings) {
		toSerialize["mappings"] = o.Mappings
	}
	return toSerialize, nil
}

type NullableClientMappingsRepresentation struct {
	value *ClientMappingsRepresentation
	isSet bool
}

func (v NullableClientMappingsRepresentation) Get() *ClientMappingsRepresentation {
	return v.value
}

func (v *NullableClientMappingsRepresentation) Set(val *ClientMappingsRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableClientMappingsRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableClientMappingsRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientMappingsRepresentation(val *ClientMappingsRepresentation) *NullableClientMappingsRepresentation {
	return &NullableClientMappingsRepresentation{value: val, isSet: true}
}

func (v NullableClientMappingsRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientMappingsRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
