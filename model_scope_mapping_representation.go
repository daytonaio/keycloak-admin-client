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

// checks if the ScopeMappingRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopeMappingRepresentation{}

// ScopeMappingRepresentation struct for ScopeMappingRepresentation
type ScopeMappingRepresentation struct {
	Self   *string `json:"self,omitempty"`
	Client *string `json:"client,omitempty"`
	// Deprecated
	ClientTemplate *string  `json:"clientTemplate,omitempty"`
	ClientScope    *string  `json:"clientScope,omitempty"`
	Roles          []string `json:"roles,omitempty"`
}

// NewScopeMappingRepresentation instantiates a new ScopeMappingRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopeMappingRepresentation() *ScopeMappingRepresentation {
	this := ScopeMappingRepresentation{}
	return &this
}

// NewScopeMappingRepresentationWithDefaults instantiates a new ScopeMappingRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeMappingRepresentationWithDefaults() *ScopeMappingRepresentation {
	this := ScopeMappingRepresentation{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ScopeMappingRepresentation) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeMappingRepresentation) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ScopeMappingRepresentation) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ScopeMappingRepresentation) SetSelf(v string) {
	o.Self = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ScopeMappingRepresentation) GetClient() string {
	if o == nil || IsNil(o.Client) {
		var ret string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeMappingRepresentation) GetClientOk() (*string, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ScopeMappingRepresentation) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given string and assigns it to the Client field.
func (o *ScopeMappingRepresentation) SetClient(v string) {
	o.Client = &v
}

// GetClientTemplate returns the ClientTemplate field value if set, zero value otherwise.
// Deprecated
func (o *ScopeMappingRepresentation) GetClientTemplate() string {
	if o == nil || IsNil(o.ClientTemplate) {
		var ret string
		return ret
	}
	return *o.ClientTemplate
}

// GetClientTemplateOk returns a tuple with the ClientTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ScopeMappingRepresentation) GetClientTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientTemplate) {
		return nil, false
	}
	return o.ClientTemplate, true
}

// HasClientTemplate returns a boolean if a field has been set.
func (o *ScopeMappingRepresentation) HasClientTemplate() bool {
	if o != nil && !IsNil(o.ClientTemplate) {
		return true
	}

	return false
}

// SetClientTemplate gets a reference to the given string and assigns it to the ClientTemplate field.
// Deprecated
func (o *ScopeMappingRepresentation) SetClientTemplate(v string) {
	o.ClientTemplate = &v
}

// GetClientScope returns the ClientScope field value if set, zero value otherwise.
func (o *ScopeMappingRepresentation) GetClientScope() string {
	if o == nil || IsNil(o.ClientScope) {
		var ret string
		return ret
	}
	return *o.ClientScope
}

// GetClientScopeOk returns a tuple with the ClientScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeMappingRepresentation) GetClientScopeOk() (*string, bool) {
	if o == nil || IsNil(o.ClientScope) {
		return nil, false
	}
	return o.ClientScope, true
}

// HasClientScope returns a boolean if a field has been set.
func (o *ScopeMappingRepresentation) HasClientScope() bool {
	if o != nil && !IsNil(o.ClientScope) {
		return true
	}

	return false
}

// SetClientScope gets a reference to the given string and assigns it to the ClientScope field.
func (o *ScopeMappingRepresentation) SetClientScope(v string) {
	o.ClientScope = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ScopeMappingRepresentation) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopeMappingRepresentation) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ScopeMappingRepresentation) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *ScopeMappingRepresentation) SetRoles(v []string) {
	o.Roles = v
}

func (o ScopeMappingRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopeMappingRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.ClientTemplate) {
		toSerialize["clientTemplate"] = o.ClientTemplate
	}
	if !IsNil(o.ClientScope) {
		toSerialize["clientScope"] = o.ClientScope
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableScopeMappingRepresentation struct {
	value *ScopeMappingRepresentation
	isSet bool
}

func (v NullableScopeMappingRepresentation) Get() *ScopeMappingRepresentation {
	return v.value
}

func (v *NullableScopeMappingRepresentation) Set(val *ScopeMappingRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeMappingRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeMappingRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeMappingRepresentation(val *ScopeMappingRepresentation) *NullableScopeMappingRepresentation {
	return &NullableScopeMappingRepresentation{value: val, isSet: true}
}

func (v NullableScopeMappingRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeMappingRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
