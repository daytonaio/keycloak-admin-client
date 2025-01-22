# ConfigPropertyRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**HelpText** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **interface{}** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**Secret** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewConfigPropertyRepresentation

`func NewConfigPropertyRepresentation() *ConfigPropertyRepresentation`

NewConfigPropertyRepresentation instantiates a new ConfigPropertyRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPropertyRepresentationWithDefaults

`func NewConfigPropertyRepresentationWithDefaults() *ConfigPropertyRepresentation`

NewConfigPropertyRepresentationWithDefaults instantiates a new ConfigPropertyRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigPropertyRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigPropertyRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigPropertyRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigPropertyRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *ConfigPropertyRepresentation) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConfigPropertyRepresentation) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConfigPropertyRepresentation) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConfigPropertyRepresentation) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetHelpText

`func (o *ConfigPropertyRepresentation) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *ConfigPropertyRepresentation) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *ConfigPropertyRepresentation) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *ConfigPropertyRepresentation) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetType

`func (o *ConfigPropertyRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigPropertyRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigPropertyRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigPropertyRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ConfigPropertyRepresentation) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ConfigPropertyRepresentation) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ConfigPropertyRepresentation) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ConfigPropertyRepresentation) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ConfigPropertyRepresentation) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ConfigPropertyRepresentation) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptions

`func (o *ConfigPropertyRepresentation) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ConfigPropertyRepresentation) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ConfigPropertyRepresentation) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ConfigPropertyRepresentation) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSecret

`func (o *ConfigPropertyRepresentation) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ConfigPropertyRepresentation) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ConfigPropertyRepresentation) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ConfigPropertyRepresentation) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRequired

`func (o *ConfigPropertyRepresentation) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ConfigPropertyRepresentation) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ConfigPropertyRepresentation) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ConfigPropertyRepresentation) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetReadOnly

`func (o *ConfigPropertyRepresentation) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ConfigPropertyRepresentation) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ConfigPropertyRepresentation) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ConfigPropertyRepresentation) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


