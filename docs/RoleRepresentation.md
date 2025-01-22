# RoleRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ScopeParamRequired** | Pointer to **bool** |  | [optional] 
**Composite** | Pointer to **bool** |  | [optional] 
**Composites** | Pointer to [**Composites**](Composites.md) |  | [optional] 
**ClientRole** | Pointer to **bool** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewRoleRepresentation

`func NewRoleRepresentation() *RoleRepresentation`

NewRoleRepresentation instantiates a new RoleRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleRepresentationWithDefaults

`func NewRoleRepresentationWithDefaults() *RoleRepresentation`

NewRoleRepresentationWithDefaults instantiates a new RoleRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RoleRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RoleRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScopeParamRequired

`func (o *RoleRepresentation) GetScopeParamRequired() bool`

GetScopeParamRequired returns the ScopeParamRequired field if non-nil, zero value otherwise.

### GetScopeParamRequiredOk

`func (o *RoleRepresentation) GetScopeParamRequiredOk() (*bool, bool)`

GetScopeParamRequiredOk returns a tuple with the ScopeParamRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeParamRequired

`func (o *RoleRepresentation) SetScopeParamRequired(v bool)`

SetScopeParamRequired sets ScopeParamRequired field to given value.

### HasScopeParamRequired

`func (o *RoleRepresentation) HasScopeParamRequired() bool`

HasScopeParamRequired returns a boolean if a field has been set.

### GetComposite

`func (o *RoleRepresentation) GetComposite() bool`

GetComposite returns the Composite field if non-nil, zero value otherwise.

### GetCompositeOk

`func (o *RoleRepresentation) GetCompositeOk() (*bool, bool)`

GetCompositeOk returns a tuple with the Composite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposite

`func (o *RoleRepresentation) SetComposite(v bool)`

SetComposite sets Composite field to given value.

### HasComposite

`func (o *RoleRepresentation) HasComposite() bool`

HasComposite returns a boolean if a field has been set.

### GetComposites

`func (o *RoleRepresentation) GetComposites() Composites`

GetComposites returns the Composites field if non-nil, zero value otherwise.

### GetCompositesOk

`func (o *RoleRepresentation) GetCompositesOk() (*Composites, bool)`

GetCompositesOk returns a tuple with the Composites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposites

`func (o *RoleRepresentation) SetComposites(v Composites)`

SetComposites sets Composites field to given value.

### HasComposites

`func (o *RoleRepresentation) HasComposites() bool`

HasComposites returns a boolean if a field has been set.

### GetClientRole

`func (o *RoleRepresentation) GetClientRole() bool`

GetClientRole returns the ClientRole field if non-nil, zero value otherwise.

### GetClientRoleOk

`func (o *RoleRepresentation) GetClientRoleOk() (*bool, bool)`

GetClientRoleOk returns a tuple with the ClientRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRole

`func (o *RoleRepresentation) SetClientRole(v bool)`

SetClientRole sets ClientRole field to given value.

### HasClientRole

`func (o *RoleRepresentation) HasClientRole() bool`

HasClientRole returns a boolean if a field has been set.

### GetContainerId

`func (o *RoleRepresentation) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *RoleRepresentation) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *RoleRepresentation) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *RoleRepresentation) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetAttributes

`func (o *RoleRepresentation) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoleRepresentation) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoleRepresentation) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RoleRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


