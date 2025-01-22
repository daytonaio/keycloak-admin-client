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

// checks if the AuthenticationExecutionInfoRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationExecutionInfoRepresentation{}

// AuthenticationExecutionInfoRepresentation struct for AuthenticationExecutionInfoRepresentation
type AuthenticationExecutionInfoRepresentation struct {
	Id                   *string  `json:"id,omitempty"`
	Requirement          *string  `json:"requirement,omitempty"`
	DisplayName          *string  `json:"displayName,omitempty"`
	Alias                *string  `json:"alias,omitempty"`
	Description          *string  `json:"description,omitempty"`
	RequirementChoices   []string `json:"requirementChoices,omitempty"`
	Configurable         *bool    `json:"configurable,omitempty"`
	AuthenticationFlow   *bool    `json:"authenticationFlow,omitempty"`
	ProviderId           *string  `json:"providerId,omitempty"`
	AuthenticationConfig *string  `json:"authenticationConfig,omitempty"`
	FlowId               *string  `json:"flowId,omitempty"`
	Level                *int32   `json:"level,omitempty"`
	Index                *int32   `json:"index,omitempty"`
	Priority             *int32   `json:"priority,omitempty"`
}

// NewAuthenticationExecutionInfoRepresentation instantiates a new AuthenticationExecutionInfoRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationExecutionInfoRepresentation() *AuthenticationExecutionInfoRepresentation {
	this := AuthenticationExecutionInfoRepresentation{}
	return &this
}

// NewAuthenticationExecutionInfoRepresentationWithDefaults instantiates a new AuthenticationExecutionInfoRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationExecutionInfoRepresentationWithDefaults() *AuthenticationExecutionInfoRepresentation {
	this := AuthenticationExecutionInfoRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthenticationExecutionInfoRepresentation) SetId(v string) {
	o.Id = &v
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetRequirement() string {
	if o == nil || IsNil(o.Requirement) {
		var ret string
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetRequirementOk() (*string, bool) {
	if o == nil || IsNil(o.Requirement) {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasRequirement() bool {
	if o != nil && !IsNil(o.Requirement) {
		return true
	}

	return false
}

// SetRequirement gets a reference to the given string and assigns it to the Requirement field.
func (o *AuthenticationExecutionInfoRepresentation) SetRequirement(v string) {
	o.Requirement = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AuthenticationExecutionInfoRepresentation) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *AuthenticationExecutionInfoRepresentation) SetAlias(v string) {
	o.Alias = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticationExecutionInfoRepresentation) SetDescription(v string) {
	o.Description = &v
}

// GetRequirementChoices returns the RequirementChoices field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetRequirementChoices() []string {
	if o == nil || IsNil(o.RequirementChoices) {
		var ret []string
		return ret
	}
	return o.RequirementChoices
}

// GetRequirementChoicesOk returns a tuple with the RequirementChoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetRequirementChoicesOk() ([]string, bool) {
	if o == nil || IsNil(o.RequirementChoices) {
		return nil, false
	}
	return o.RequirementChoices, true
}

// HasRequirementChoices returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasRequirementChoices() bool {
	if o != nil && !IsNil(o.RequirementChoices) {
		return true
	}

	return false
}

// SetRequirementChoices gets a reference to the given []string and assigns it to the RequirementChoices field.
func (o *AuthenticationExecutionInfoRepresentation) SetRequirementChoices(v []string) {
	o.RequirementChoices = v
}

// GetConfigurable returns the Configurable field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetConfigurable() bool {
	if o == nil || IsNil(o.Configurable) {
		var ret bool
		return ret
	}
	return *o.Configurable
}

// GetConfigurableOk returns a tuple with the Configurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetConfigurableOk() (*bool, bool) {
	if o == nil || IsNil(o.Configurable) {
		return nil, false
	}
	return o.Configurable, true
}

// HasConfigurable returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasConfigurable() bool {
	if o != nil && !IsNil(o.Configurable) {
		return true
	}

	return false
}

// SetConfigurable gets a reference to the given bool and assigns it to the Configurable field.
func (o *AuthenticationExecutionInfoRepresentation) SetConfigurable(v bool) {
	o.Configurable = &v
}

// GetAuthenticationFlow returns the AuthenticationFlow field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationFlow() bool {
	if o == nil || IsNil(o.AuthenticationFlow) {
		var ret bool
		return ret
	}
	return *o.AuthenticationFlow
}

// GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationFlowOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthenticationFlow) {
		return nil, false
	}
	return o.AuthenticationFlow, true
}

// HasAuthenticationFlow returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasAuthenticationFlow() bool {
	if o != nil && !IsNil(o.AuthenticationFlow) {
		return true
	}

	return false
}

// SetAuthenticationFlow gets a reference to the given bool and assigns it to the AuthenticationFlow field.
func (o *AuthenticationExecutionInfoRepresentation) SetAuthenticationFlow(v bool) {
	o.AuthenticationFlow = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *AuthenticationExecutionInfoRepresentation) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetAuthenticationConfig returns the AuthenticationConfig field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationConfig() string {
	if o == nil || IsNil(o.AuthenticationConfig) {
		var ret string
		return ret
	}
	return *o.AuthenticationConfig
}

// GetAuthenticationConfigOk returns a tuple with the AuthenticationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationConfigOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationConfig) {
		return nil, false
	}
	return o.AuthenticationConfig, true
}

// HasAuthenticationConfig returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasAuthenticationConfig() bool {
	if o != nil && !IsNil(o.AuthenticationConfig) {
		return true
	}

	return false
}

// SetAuthenticationConfig gets a reference to the given string and assigns it to the AuthenticationConfig field.
func (o *AuthenticationExecutionInfoRepresentation) SetAuthenticationConfig(v string) {
	o.AuthenticationConfig = &v
}

// GetFlowId returns the FlowId field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetFlowId() string {
	if o == nil || IsNil(o.FlowId) {
		var ret string
		return ret
	}
	return *o.FlowId
}

// GetFlowIdOk returns a tuple with the FlowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetFlowIdOk() (*string, bool) {
	if o == nil || IsNil(o.FlowId) {
		return nil, false
	}
	return o.FlowId, true
}

// HasFlowId returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasFlowId() bool {
	if o != nil && !IsNil(o.FlowId) {
		return true
	}

	return false
}

// SetFlowId gets a reference to the given string and assigns it to the FlowId field.
func (o *AuthenticationExecutionInfoRepresentation) SetFlowId(v string) {
	o.FlowId = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetLevel() int32 {
	if o == nil || IsNil(o.Level) {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *AuthenticationExecutionInfoRepresentation) SetLevel(v int32) {
	o.Level = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *AuthenticationExecutionInfoRepresentation) SetIndex(v int32) {
	o.Index = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AuthenticationExecutionInfoRepresentation) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationExecutionInfoRepresentation) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AuthenticationExecutionInfoRepresentation) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *AuthenticationExecutionInfoRepresentation) SetPriority(v int32) {
	o.Priority = &v
}

func (o AuthenticationExecutionInfoRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationExecutionInfoRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Requirement) {
		toSerialize["requirement"] = o.Requirement
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RequirementChoices) {
		toSerialize["requirementChoices"] = o.RequirementChoices
	}
	if !IsNil(o.Configurable) {
		toSerialize["configurable"] = o.Configurable
	}
	if !IsNil(o.AuthenticationFlow) {
		toSerialize["authenticationFlow"] = o.AuthenticationFlow
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.AuthenticationConfig) {
		toSerialize["authenticationConfig"] = o.AuthenticationConfig
	}
	if !IsNil(o.FlowId) {
		toSerialize["flowId"] = o.FlowId
	}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

type NullableAuthenticationExecutionInfoRepresentation struct {
	value *AuthenticationExecutionInfoRepresentation
	isSet bool
}

func (v NullableAuthenticationExecutionInfoRepresentation) Get() *AuthenticationExecutionInfoRepresentation {
	return v.value
}

func (v *NullableAuthenticationExecutionInfoRepresentation) Set(val *AuthenticationExecutionInfoRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationExecutionInfoRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationExecutionInfoRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationExecutionInfoRepresentation(val *AuthenticationExecutionInfoRepresentation) *NullableAuthenticationExecutionInfoRepresentation {
	return &NullableAuthenticationExecutionInfoRepresentation{value: val, isSet: true}
}

func (v NullableAuthenticationExecutionInfoRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationExecutionInfoRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
