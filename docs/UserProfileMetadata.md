# UserProfileMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]UserProfileAttributeMetadata**](UserProfileAttributeMetadata.md) |  | [optional] 
**Groups** | Pointer to [**[]UserProfileAttributeGroupMetadata**](UserProfileAttributeGroupMetadata.md) |  | [optional] 

## Methods

### NewUserProfileMetadata

`func NewUserProfileMetadata() *UserProfileMetadata`

NewUserProfileMetadata instantiates a new UserProfileMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileMetadataWithDefaults

`func NewUserProfileMetadataWithDefaults() *UserProfileMetadata`

NewUserProfileMetadataWithDefaults instantiates a new UserProfileMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UserProfileMetadata) GetAttributes() []UserProfileAttributeMetadata`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserProfileMetadata) GetAttributesOk() (*[]UserProfileAttributeMetadata, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserProfileMetadata) SetAttributes(v []UserProfileAttributeMetadata)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserProfileMetadata) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetGroups

`func (o *UserProfileMetadata) GetGroups() []UserProfileAttributeGroupMetadata`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserProfileMetadata) GetGroupsOk() (*[]UserProfileAttributeGroupMetadata, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserProfileMetadata) SetGroups(v []UserProfileAttributeGroupMetadata)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserProfileMetadata) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


