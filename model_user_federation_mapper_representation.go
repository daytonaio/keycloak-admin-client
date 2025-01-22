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

// checks if the UserFederationMapperRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFederationMapperRepresentation{}

// UserFederationMapperRepresentation struct for UserFederationMapperRepresentation
type UserFederationMapperRepresentation struct {
	Id                            *string            `json:"id,omitempty"`
	Name                          *string            `json:"name,omitempty"`
	FederationProviderDisplayName *string            `json:"federationProviderDisplayName,omitempty"`
	FederationMapperType          *string            `json:"federationMapperType,omitempty"`
	Config                        *map[string]string `json:"config,omitempty"`
}

// NewUserFederationMapperRepresentation instantiates a new UserFederationMapperRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFederationMapperRepresentation() *UserFederationMapperRepresentation {
	this := UserFederationMapperRepresentation{}
	return &this
}

// NewUserFederationMapperRepresentationWithDefaults instantiates a new UserFederationMapperRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFederationMapperRepresentationWithDefaults() *UserFederationMapperRepresentation {
	this := UserFederationMapperRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserFederationMapperRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFederationMapperRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserFederationMapperRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserFederationMapperRepresentation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserFederationMapperRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFederationMapperRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserFederationMapperRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserFederationMapperRepresentation) SetName(v string) {
	o.Name = &v
}

// GetFederationProviderDisplayName returns the FederationProviderDisplayName field value if set, zero value otherwise.
func (o *UserFederationMapperRepresentation) GetFederationProviderDisplayName() string {
	if o == nil || IsNil(o.FederationProviderDisplayName) {
		var ret string
		return ret
	}
	return *o.FederationProviderDisplayName
}

// GetFederationProviderDisplayNameOk returns a tuple with the FederationProviderDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFederationMapperRepresentation) GetFederationProviderDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.FederationProviderDisplayName) {
		return nil, false
	}
	return o.FederationProviderDisplayName, true
}

// HasFederationProviderDisplayName returns a boolean if a field has been set.
func (o *UserFederationMapperRepresentation) HasFederationProviderDisplayName() bool {
	if o != nil && !IsNil(o.FederationProviderDisplayName) {
		return true
	}

	return false
}

// SetFederationProviderDisplayName gets a reference to the given string and assigns it to the FederationProviderDisplayName field.
func (o *UserFederationMapperRepresentation) SetFederationProviderDisplayName(v string) {
	o.FederationProviderDisplayName = &v
}

// GetFederationMapperType returns the FederationMapperType field value if set, zero value otherwise.
func (o *UserFederationMapperRepresentation) GetFederationMapperType() string {
	if o == nil || IsNil(o.FederationMapperType) {
		var ret string
		return ret
	}
	return *o.FederationMapperType
}

// GetFederationMapperTypeOk returns a tuple with the FederationMapperType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFederationMapperRepresentation) GetFederationMapperTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FederationMapperType) {
		return nil, false
	}
	return o.FederationMapperType, true
}

// HasFederationMapperType returns a boolean if a field has been set.
func (o *UserFederationMapperRepresentation) HasFederationMapperType() bool {
	if o != nil && !IsNil(o.FederationMapperType) {
		return true
	}

	return false
}

// SetFederationMapperType gets a reference to the given string and assigns it to the FederationMapperType field.
func (o *UserFederationMapperRepresentation) SetFederationMapperType(v string) {
	o.FederationMapperType = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *UserFederationMapperRepresentation) GetConfig() map[string]string {
	if o == nil || IsNil(o.Config) {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFederationMapperRepresentation) GetConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *UserFederationMapperRepresentation) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *UserFederationMapperRepresentation) SetConfig(v map[string]string) {
	o.Config = &v
}

func (o UserFederationMapperRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFederationMapperRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FederationProviderDisplayName) {
		toSerialize["federationProviderDisplayName"] = o.FederationProviderDisplayName
	}
	if !IsNil(o.FederationMapperType) {
		toSerialize["federationMapperType"] = o.FederationMapperType
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableUserFederationMapperRepresentation struct {
	value *UserFederationMapperRepresentation
	isSet bool
}

func (v NullableUserFederationMapperRepresentation) Get() *UserFederationMapperRepresentation {
	return v.value
}

func (v *NullableUserFederationMapperRepresentation) Set(val *UserFederationMapperRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFederationMapperRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFederationMapperRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFederationMapperRepresentation(val *UserFederationMapperRepresentation) *NullableUserFederationMapperRepresentation {
	return &NullableUserFederationMapperRepresentation{value: val, isSet: true}
}

func (v NullableUserFederationMapperRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFederationMapperRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
