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

// checks if the AddressClaimSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressClaimSet{}

// AddressClaimSet struct for AddressClaimSet
type AddressClaimSet struct {
	Formatted *string `json:"formatted,omitempty"`
	StreetAddress *string `json:"street_address,omitempty"`
	Locality *string `json:"locality,omitempty"`
	Region *string `json:"region,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Country *string `json:"country,omitempty"`
}

// NewAddressClaimSet instantiates a new AddressClaimSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressClaimSet() *AddressClaimSet {
	this := AddressClaimSet{}
	return &this
}

// NewAddressClaimSetWithDefaults instantiates a new AddressClaimSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressClaimSetWithDefaults() *AddressClaimSet {
	this := AddressClaimSet{}
	return &this
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *AddressClaimSet) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressClaimSet) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *AddressClaimSet) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *AddressClaimSet) SetFormatted(v string) {
	o.Formatted = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *AddressClaimSet) GetStreetAddress() string {
	if o == nil || IsNil(o.StreetAddress) {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressClaimSet) GetStreetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *AddressClaimSet) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *AddressClaimSet) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *AddressClaimSet) GetLocality() string {
	if o == nil || IsNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressClaimSet) GetLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.Locality) {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *AddressClaimSet) HasLocality() bool {
	if o != nil && !IsNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *AddressClaimSet) SetLocality(v string) {
	o.Locality = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AddressClaimSet) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressClaimSet) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AddressClaimSet) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AddressClaimSet) SetRegion(v string) {
	o.Region = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AddressClaimSet) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressClaimSet) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressClaimSet) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AddressClaimSet) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AddressClaimSet) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressClaimSet) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AddressClaimSet) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AddressClaimSet) SetCountry(v string) {
	o.Country = &v
}

func (o AddressClaimSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressClaimSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["street_address"] = o.StreetAddress
	}
	if !IsNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableAddressClaimSet struct {
	value *AddressClaimSet
	isSet bool
}

func (v NullableAddressClaimSet) Get() *AddressClaimSet {
	return v.value
}

func (v *NullableAddressClaimSet) Set(val *AddressClaimSet) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressClaimSet) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressClaimSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressClaimSet(val *AddressClaimSet) *NullableAddressClaimSet {
	return &NullableAddressClaimSet{value: val, isSet: true}
}

func (v NullableAddressClaimSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressClaimSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


