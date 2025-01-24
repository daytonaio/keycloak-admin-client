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

// checks if the AuthenticationFlowRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationFlowRepresentation{}

// AuthenticationFlowRepresentation struct for AuthenticationFlowRepresentation
type AuthenticationFlowRepresentation struct {
	Id *string `json:"id,omitempty"`
	Alias *string `json:"alias,omitempty"`
	Description *string `json:"description,omitempty"`
	ProviderId *string `json:"providerId,omitempty"`
	TopLevel *bool `json:"topLevel,omitempty"`
	BuiltIn *bool `json:"builtIn,omitempty"`
	AuthenticationExecutions []AuthenticationExecutionExportRepresentation `json:"authenticationExecutions,omitempty"`
}

// NewAuthenticationFlowRepresentation instantiates a new AuthenticationFlowRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationFlowRepresentation() *AuthenticationFlowRepresentation {
	this := AuthenticationFlowRepresentation{}
	return &this
}

// NewAuthenticationFlowRepresentationWithDefaults instantiates a new AuthenticationFlowRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationFlowRepresentationWithDefaults() *AuthenticationFlowRepresentation {
	this := AuthenticationFlowRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationFlowRepresentation) SetId(v string) {
	o.Id = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *AuthenticationFlowRepresentation) SetAlias(v string) {
	o.Alias = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticationFlowRepresentation) SetDescription(v string) {
	o.Description = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *AuthenticationFlowRepresentation) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetTopLevel returns the TopLevel field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetTopLevel() bool {
	if o == nil || IsNil(o.TopLevel) {
		var ret bool
		return ret
	}
	return *o.TopLevel
}

// GetTopLevelOk returns a tuple with the TopLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetTopLevelOk() (*bool, bool) {
	if o == nil || IsNil(o.TopLevel) {
		return nil, false
	}
	return o.TopLevel, true
}

// HasTopLevel returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasTopLevel() bool {
	if o != nil && !IsNil(o.TopLevel) {
		return true
	}

	return false
}

// SetTopLevel gets a reference to the given bool and assigns it to the TopLevel field.
func (o *AuthenticationFlowRepresentation) SetTopLevel(v bool) {
	o.TopLevel = &v
}

// GetBuiltIn returns the BuiltIn field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetBuiltIn() bool {
	if o == nil || IsNil(o.BuiltIn) {
		var ret bool
		return ret
	}
	return *o.BuiltIn
}

// GetBuiltInOk returns a tuple with the BuiltIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetBuiltInOk() (*bool, bool) {
	if o == nil || IsNil(o.BuiltIn) {
		return nil, false
	}
	return o.BuiltIn, true
}

// HasBuiltIn returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasBuiltIn() bool {
	if o != nil && !IsNil(o.BuiltIn) {
		return true
	}

	return false
}

// SetBuiltIn gets a reference to the given bool and assigns it to the BuiltIn field.
func (o *AuthenticationFlowRepresentation) SetBuiltIn(v bool) {
	o.BuiltIn = &v
}

// GetAuthenticationExecutions returns the AuthenticationExecutions field value if set, zero value otherwise.
func (o *AuthenticationFlowRepresentation) GetAuthenticationExecutions() []AuthenticationExecutionExportRepresentation {
	if o == nil || IsNil(o.AuthenticationExecutions) {
		var ret []AuthenticationExecutionExportRepresentation
		return ret
	}
	return o.AuthenticationExecutions
}

// GetAuthenticationExecutionsOk returns a tuple with the AuthenticationExecutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationFlowRepresentation) GetAuthenticationExecutionsOk() ([]AuthenticationExecutionExportRepresentation, bool) {
	if o == nil || IsNil(o.AuthenticationExecutions) {
		return nil, false
	}
	return o.AuthenticationExecutions, true
}

// HasAuthenticationExecutions returns a boolean if a field has been set.
func (o *AuthenticationFlowRepresentation) HasAuthenticationExecutions() bool {
	if o != nil && !IsNil(o.AuthenticationExecutions) {
		return true
	}

	return false
}

// SetAuthenticationExecutions gets a reference to the given []AuthenticationExecutionExportRepresentation and assigns it to the AuthenticationExecutions field.
func (o *AuthenticationFlowRepresentation) SetAuthenticationExecutions(v []AuthenticationExecutionExportRepresentation) {
	o.AuthenticationExecutions = v
}

func (o AuthenticationFlowRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationFlowRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.TopLevel) {
		toSerialize["topLevel"] = o.TopLevel
	}
	if !IsNil(o.BuiltIn) {
		toSerialize["builtIn"] = o.BuiltIn
	}
	if !IsNil(o.AuthenticationExecutions) {
		toSerialize["authenticationExecutions"] = o.AuthenticationExecutions
	}
	return toSerialize, nil
}

type NullableAuthenticationFlowRepresentation struct {
	value *AuthenticationFlowRepresentation
	isSet bool
}

func (v NullableAuthenticationFlowRepresentation) Get() *AuthenticationFlowRepresentation {
	return v.value
}

func (v *NullableAuthenticationFlowRepresentation) Set(val *AuthenticationFlowRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationFlowRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationFlowRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationFlowRepresentation(val *AuthenticationFlowRepresentation) *NullableAuthenticationFlowRepresentation {
	return &NullableAuthenticationFlowRepresentation{value: val, isSet: true}
}

func (v NullableAuthenticationFlowRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationFlowRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


