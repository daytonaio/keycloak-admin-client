# ClientProfileRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Executors** | Pointer to [**[]ClientPolicyExecutorRepresentation**](ClientPolicyExecutorRepresentation.md) |  | [optional] 

## Methods

### NewClientProfileRepresentation

`func NewClientProfileRepresentation() *ClientProfileRepresentation`

NewClientProfileRepresentation instantiates a new ClientProfileRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientProfileRepresentationWithDefaults

`func NewClientProfileRepresentationWithDefaults() *ClientProfileRepresentation`

NewClientProfileRepresentationWithDefaults instantiates a new ClientProfileRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientProfileRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientProfileRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientProfileRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientProfileRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClientProfileRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientProfileRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientProfileRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientProfileRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutors

`func (o *ClientProfileRepresentation) GetExecutors() []ClientPolicyExecutorRepresentation`

GetExecutors returns the Executors field if non-nil, zero value otherwise.

### GetExecutorsOk

`func (o *ClientProfileRepresentation) GetExecutorsOk() (*[]ClientPolicyExecutorRepresentation, bool)`

GetExecutorsOk returns a tuple with the Executors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutors

`func (o *ClientProfileRepresentation) SetExecutors(v []ClientPolicyExecutorRepresentation)`

SetExecutors sets Executors field to given value.

### HasExecutors

`func (o *ClientProfileRepresentation) HasExecutors() bool`

HasExecutors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


