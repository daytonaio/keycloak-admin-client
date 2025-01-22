# AdminEventRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int64** |  | [optional] 
**RealmId** | Pointer to **string** |  | [optional] 
**AuthDetails** | Pointer to [**AuthDetailsRepresentation**](AuthDetailsRepresentation.md) |  | [optional] 
**OperationType** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResourcePath** | Pointer to **string** |  | [optional] 
**Representation** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAdminEventRepresentation

`func NewAdminEventRepresentation() *AdminEventRepresentation`

NewAdminEventRepresentation instantiates a new AdminEventRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminEventRepresentationWithDefaults

`func NewAdminEventRepresentationWithDefaults() *AdminEventRepresentation`

NewAdminEventRepresentationWithDefaults instantiates a new AdminEventRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *AdminEventRepresentation) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AdminEventRepresentation) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AdminEventRepresentation) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AdminEventRepresentation) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRealmId

`func (o *AdminEventRepresentation) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *AdminEventRepresentation) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *AdminEventRepresentation) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *AdminEventRepresentation) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetAuthDetails

`func (o *AdminEventRepresentation) GetAuthDetails() AuthDetailsRepresentation`

GetAuthDetails returns the AuthDetails field if non-nil, zero value otherwise.

### GetAuthDetailsOk

`func (o *AdminEventRepresentation) GetAuthDetailsOk() (*AuthDetailsRepresentation, bool)`

GetAuthDetailsOk returns a tuple with the AuthDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDetails

`func (o *AdminEventRepresentation) SetAuthDetails(v AuthDetailsRepresentation)`

SetAuthDetails sets AuthDetails field to given value.

### HasAuthDetails

`func (o *AdminEventRepresentation) HasAuthDetails() bool`

HasAuthDetails returns a boolean if a field has been set.

### GetOperationType

`func (o *AdminEventRepresentation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AdminEventRepresentation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AdminEventRepresentation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *AdminEventRepresentation) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetResourceType

`func (o *AdminEventRepresentation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AdminEventRepresentation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AdminEventRepresentation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AdminEventRepresentation) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourcePath

`func (o *AdminEventRepresentation) GetResourcePath() string`

GetResourcePath returns the ResourcePath field if non-nil, zero value otherwise.

### GetResourcePathOk

`func (o *AdminEventRepresentation) GetResourcePathOk() (*string, bool)`

GetResourcePathOk returns a tuple with the ResourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePath

`func (o *AdminEventRepresentation) SetResourcePath(v string)`

SetResourcePath sets ResourcePath field to given value.

### HasResourcePath

`func (o *AdminEventRepresentation) HasResourcePath() bool`

HasResourcePath returns a boolean if a field has been set.

### GetRepresentation

`func (o *AdminEventRepresentation) GetRepresentation() string`

GetRepresentation returns the Representation field if non-nil, zero value otherwise.

### GetRepresentationOk

`func (o *AdminEventRepresentation) GetRepresentationOk() (*string, bool)`

GetRepresentationOk returns a tuple with the Representation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentation

`func (o *AdminEventRepresentation) SetRepresentation(v string)`

SetRepresentation sets Representation field to given value.

### HasRepresentation

`func (o *AdminEventRepresentation) HasRepresentation() bool`

HasRepresentation returns a boolean if a field has been set.

### GetError

`func (o *AdminEventRepresentation) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AdminEventRepresentation) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AdminEventRepresentation) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AdminEventRepresentation) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetails

`func (o *AdminEventRepresentation) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdminEventRepresentation) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdminEventRepresentation) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdminEventRepresentation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


