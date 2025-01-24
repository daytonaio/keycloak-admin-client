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

// DecisionStrategy the model 'DecisionStrategy'
type DecisionStrategy string

// List of DecisionStrategy
const (
	AFFIRMATIVE DecisionStrategy = "AFFIRMATIVE"
	UNANIMOUS DecisionStrategy = "UNANIMOUS"
	CONSENSUS DecisionStrategy = "CONSENSUS"
)

// All allowed values of DecisionStrategy enum
var AllowedDecisionStrategyEnumValues = []DecisionStrategy{
	"AFFIRMATIVE",
	"UNANIMOUS",
	"CONSENSUS",
}

func (v *DecisionStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DecisionStrategy(value)
	for _, existing := range AllowedDecisionStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DecisionStrategy", value)
}

// NewDecisionStrategyFromValue returns a pointer to a valid DecisionStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDecisionStrategyFromValue(v string) (*DecisionStrategy, error) {
	ev := DecisionStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DecisionStrategy: valid values are %v", v, AllowedDecisionStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DecisionStrategy) IsValid() bool {
	for _, existing := range AllowedDecisionStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DecisionStrategy value
func (v DecisionStrategy) Ptr() *DecisionStrategy {
	return &v
}

type NullableDecisionStrategy struct {
	value *DecisionStrategy
	isSet bool
}

func (v NullableDecisionStrategy) Get() *DecisionStrategy {
	return v.value
}

func (v *NullableDecisionStrategy) Set(val *DecisionStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionStrategy(val *DecisionStrategy) *NullableDecisionStrategy {
	return &NullableDecisionStrategy{value: val, isSet: true}
}

func (v NullableDecisionStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

