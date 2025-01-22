# EventRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RealmId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEventRepresentation

`func NewEventRepresentation() *EventRepresentation`

NewEventRepresentation instantiates a new EventRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventRepresentationWithDefaults

`func NewEventRepresentationWithDefaults() *EventRepresentation`

NewEventRepresentationWithDefaults instantiates a new EventRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *EventRepresentation) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventRepresentation) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EventRepresentation) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *EventRepresentation) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *EventRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRealmId

`func (o *EventRepresentation) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *EventRepresentation) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *EventRepresentation) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *EventRepresentation) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetClientId

`func (o *EventRepresentation) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EventRepresentation) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EventRepresentation) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EventRepresentation) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetUserId

`func (o *EventRepresentation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EventRepresentation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EventRepresentation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EventRepresentation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSessionId

`func (o *EventRepresentation) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EventRepresentation) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EventRepresentation) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EventRepresentation) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetIpAddress

`func (o *EventRepresentation) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *EventRepresentation) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *EventRepresentation) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *EventRepresentation) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetError

`func (o *EventRepresentation) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventRepresentation) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventRepresentation) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EventRepresentation) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetails

`func (o *EventRepresentation) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EventRepresentation) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EventRepresentation) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EventRepresentation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


