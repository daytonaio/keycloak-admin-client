# UserSessionRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 
**LastAccess** | Pointer to **int64** |  | [optional] 
**RememberMe** | Pointer to **bool** |  | [optional] 
**Clients** | Pointer to **map[string]string** |  | [optional] 
**TransientUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserSessionRepresentation

`func NewUserSessionRepresentation() *UserSessionRepresentation`

NewUserSessionRepresentation instantiates a new UserSessionRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionRepresentationWithDefaults

`func NewUserSessionRepresentationWithDefaults() *UserSessionRepresentation`

NewUserSessionRepresentationWithDefaults instantiates a new UserSessionRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSessionRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSessionRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSessionRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserSessionRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UserSessionRepresentation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSessionRepresentation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSessionRepresentation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSessionRepresentation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUserId

`func (o *UserSessionRepresentation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSessionRepresentation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSessionRepresentation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserSessionRepresentation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIpAddress

`func (o *UserSessionRepresentation) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserSessionRepresentation) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserSessionRepresentation) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserSessionRepresentation) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetStart

`func (o *UserSessionRepresentation) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UserSessionRepresentation) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UserSessionRepresentation) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *UserSessionRepresentation) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetLastAccess

`func (o *UserSessionRepresentation) GetLastAccess() int64`

GetLastAccess returns the LastAccess field if non-nil, zero value otherwise.

### GetLastAccessOk

`func (o *UserSessionRepresentation) GetLastAccessOk() (*int64, bool)`

GetLastAccessOk returns a tuple with the LastAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccess

`func (o *UserSessionRepresentation) SetLastAccess(v int64)`

SetLastAccess sets LastAccess field to given value.

### HasLastAccess

`func (o *UserSessionRepresentation) HasLastAccess() bool`

HasLastAccess returns a boolean if a field has been set.

### GetRememberMe

`func (o *UserSessionRepresentation) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *UserSessionRepresentation) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *UserSessionRepresentation) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *UserSessionRepresentation) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetClients

`func (o *UserSessionRepresentation) GetClients() map[string]string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *UserSessionRepresentation) GetClientsOk() (*map[string]string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *UserSessionRepresentation) SetClients(v map[string]string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *UserSessionRepresentation) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetTransientUser

`func (o *UserSessionRepresentation) GetTransientUser() bool`

GetTransientUser returns the TransientUser field if non-nil, zero value otherwise.

### GetTransientUserOk

`func (o *UserSessionRepresentation) GetTransientUserOk() (*bool, bool)`

GetTransientUserOk returns a tuple with the TransientUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransientUser

`func (o *UserSessionRepresentation) SetTransientUser(v bool)`

SetTransientUser sets TransientUser field to given value.

### HasTransientUser

`func (o *UserSessionRepresentation) HasTransientUser() bool`

HasTransientUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


