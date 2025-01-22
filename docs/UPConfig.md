# UPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]UPAttribute**](UPAttribute.md) |  | [optional] 
**Groups** | Pointer to [**[]UPGroup**](UPGroup.md) |  | [optional] 
**UnmanagedAttributePolicy** | Pointer to [**UnmanagedAttributePolicy**](UnmanagedAttributePolicy.md) |  | [optional] 

## Methods

### NewUPConfig

`func NewUPConfig() *UPConfig`

NewUPConfig instantiates a new UPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUPConfigWithDefaults

`func NewUPConfigWithDefaults() *UPConfig`

NewUPConfigWithDefaults instantiates a new UPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UPConfig) GetAttributes() []UPAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UPConfig) GetAttributesOk() (*[]UPAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UPConfig) SetAttributes(v []UPAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UPConfig) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGroups

`func (o *UPConfig) GetGroups() []UPGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UPConfig) GetGroupsOk() (*[]UPGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UPConfig) SetGroups(v []UPGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UPConfig) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUnmanagedAttributePolicy

`func (o *UPConfig) GetUnmanagedAttributePolicy() UnmanagedAttributePolicy`

GetUnmanagedAttributePolicy returns the UnmanagedAttributePolicy field if non-nil, zero value otherwise.

### GetUnmanagedAttributePolicyOk

`func (o *UPConfig) GetUnmanagedAttributePolicyOk() (*UnmanagedAttributePolicy, bool)`

GetUnmanagedAttributePolicyOk returns a tuple with the UnmanagedAttributePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmanagedAttributePolicy

`func (o *UPConfig) SetUnmanagedAttributePolicy(v UnmanagedAttributePolicy)`

SetUnmanagedAttributePolicy sets UnmanagedAttributePolicy field to given value.

### HasUnmanagedAttributePolicy

`func (o *UPConfig) HasUnmanagedAttributePolicy() bool`

HasUnmanagedAttributePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


