/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keycloak_admin_client

import (
	"encoding/json"
	"fmt"
)

// ScopeEnforcementMode the model 'ScopeEnforcementMode'
type ScopeEnforcementMode string

// List of ScopeEnforcementMode
const (
	SCOPE_ALL      ScopeEnforcementMode = "ALL"
	SCOPE_ANY      ScopeEnforcementMode = "ANY"
	SCOPE_DISABLED ScopeEnforcementMode = "DISABLED"
)

// All allowed values of ScopeEnforcementMode enum
var AllowedScopeEnforcementModeEnumValues = []ScopeEnforcementMode{
	"ALL",
	"ANY",
	"DISABLED",
}

func (v *ScopeEnforcementMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScopeEnforcementMode(value)
	for _, existing := range AllowedScopeEnforcementModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScopeEnforcementMode", value)
}

// NewScopeEnforcementModeFromValue returns a pointer to a valid ScopeEnforcementMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScopeEnforcementModeFromValue(v string) (*ScopeEnforcementMode, error) {
	ev := ScopeEnforcementMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScopeEnforcementMode: valid values are %v", v, AllowedScopeEnforcementModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScopeEnforcementMode) IsValid() bool {
	for _, existing := range AllowedScopeEnforcementModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScopeEnforcementMode value
func (v ScopeEnforcementMode) Ptr() *ScopeEnforcementMode {
	return &v
}

type NullableScopeEnforcementMode struct {
	value *ScopeEnforcementMode
	isSet bool
}

func (v NullableScopeEnforcementMode) Get() *ScopeEnforcementMode {
	return v.value
}

func (v *NullableScopeEnforcementMode) Set(val *ScopeEnforcementMode) {
	v.value = val
	v.isSet = true
}

func (v NullableScopeEnforcementMode) IsSet() bool {
	return v.isSet
}

func (v *NullableScopeEnforcementMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopeEnforcementMode(val *ScopeEnforcementMode) *NullableScopeEnforcementMode {
	return &NullableScopeEnforcementMode{value: val, isSet: true}
}

func (v NullableScopeEnforcementMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopeEnforcementMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
