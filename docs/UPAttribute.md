# UPAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Validations** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Annotations** | Pointer to **map[string]interface{}** |  | [optional] 
**Required** | Pointer to [**UPAttributeRequired**](UPAttributeRequired.md) |  | [optional] 
**Permissions** | Pointer to [**UPAttributePermissions**](UPAttributePermissions.md) |  | [optional] 
**Selector** | Pointer to [**UPAttributeSelector**](UPAttributeSelector.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Multivalued** | Pointer to **bool** |  | [optional] 

## Methods

### NewUPAttribute

`func NewUPAttribute() *UPAttribute`

NewUPAttribute instantiates a new UPAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUPAttributeWithDefaults

`func NewUPAttributeWithDefaults() *UPAttribute`

NewUPAttributeWithDefaults instantiates a new UPAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UPAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UPAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UPAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UPAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UPAttribute) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UPAttribute) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UPAttribute) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UPAttribute) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetValidations

`func (o *UPAttribute) GetValidations() map[string]map[string]interface{}`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *UPAttribute) GetValidationsOk() (*map[string]map[string]interface{}, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *UPAttribute) SetValidations(v map[string]map[string]interface{})`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *UPAttribute) HasValidations() bool`

HasValidations returns a boolean if a field has been set.

### GetAnnotations

`func (o *UPAttribute) GetAnnotations() map[string]interface{}`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *UPAttribute) GetAnnotationsOk() (*map[string]interface{}, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *UPAttribute) SetAnnotations(v map[string]interface{})`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *UPAttribute) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetRequired

`func (o *UPAttribute) GetRequired() UPAttributeRequired`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UPAttribute) GetRequiredOk() (*UPAttributeRequired, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UPAttribute) SetRequired(v UPAttributeRequired)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UPAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPermissions

`func (o *UPAttribute) GetPermissions() UPAttributePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UPAttribute) GetPermissionsOk() (*UPAttributePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UPAttribute) SetPermissions(v UPAttributePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UPAttribute) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSelector

`func (o *UPAttribute) GetSelector() UPAttributeSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *UPAttribute) GetSelectorOk() (*UPAttributeSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *UPAttribute) SetSelector(v UPAttributeSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *UPAttribute) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetGroup

`func (o *UPAttribute) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UPAttribute) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UPAttribute) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UPAttribute) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMultivalued

`func (o *UPAttribute) GetMultivalued() bool`

GetMultivalued returns the Multivalued field if non-nil, zero value otherwise.

### GetMultivaluedOk

`func (o *UPAttribute) GetMultivaluedOk() (*bool, bool)`

GetMultivaluedOk returns a tuple with the Multivalued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultivalued

`func (o *UPAttribute) SetMultivalued(v bool)`

SetMultivalued sets Multivalued field to given value.

### HasMultivalued

`func (o *UPAttribute) HasMultivalued() bool`

HasMultivalued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


