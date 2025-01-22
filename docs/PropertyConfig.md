# PropertyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applicable** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPropertyConfig

`func NewPropertyConfig() *PropertyConfig`

NewPropertyConfig instantiates a new PropertyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyConfigWithDefaults

`func NewPropertyConfigWithDefaults() *PropertyConfig`

NewPropertyConfigWithDefaults instantiates a new PropertyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicable

`func (o *PropertyConfig) GetApplicable() bool`

GetApplicable returns the Applicable field if non-nil, zero value otherwise.

### GetApplicableOk

`func (o *PropertyConfig) GetApplicableOk() (*bool, bool)`

GetApplicableOk returns a tuple with the Applicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicable

`func (o *PropertyConfig) SetApplicable(v bool)`

SetApplicable sets Applicable field to given value.

### HasApplicable

`func (o *PropertyConfig) HasApplicable() bool`

HasApplicable returns a boolean if a field has been set.

### GetValue

`func (o *PropertyConfig) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyConfig) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyConfig) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PropertyConfig) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PropertyConfig) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PropertyConfig) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


