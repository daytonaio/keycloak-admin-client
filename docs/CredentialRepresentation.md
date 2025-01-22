# CredentialRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserLabel** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **int64** |  | [optional] 
**SecretData** | Pointer to **string** |  | [optional] 
**CredentialData** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Temporary** | Pointer to **bool** |  | [optional] 
**Device** | Pointer to **string** |  | [optional] 
**HashedSaltedValue** | Pointer to **string** |  | [optional] 
**Salt** | Pointer to **string** |  | [optional] 
**HashIterations** | Pointer to **int32** |  | [optional] 
**Counter** | Pointer to **int32** |  | [optional] 
**Algorithm** | Pointer to **string** |  | [optional] 
**Digits** | Pointer to **int32** |  | [optional] 
**Period** | Pointer to **int32** |  | [optional] 
**Config** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCredentialRepresentation

`func NewCredentialRepresentation() *CredentialRepresentation`

NewCredentialRepresentation instantiates a new CredentialRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialRepresentationWithDefaults

`func NewCredentialRepresentationWithDefaults() *CredentialRepresentation`

NewCredentialRepresentationWithDefaults instantiates a new CredentialRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CredentialRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserLabel

`func (o *CredentialRepresentation) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *CredentialRepresentation) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *CredentialRepresentation) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *CredentialRepresentation) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetCreatedDate

`func (o *CredentialRepresentation) GetCreatedDate() int64`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *CredentialRepresentation) GetCreatedDateOk() (*int64, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *CredentialRepresentation) SetCreatedDate(v int64)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *CredentialRepresentation) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetSecretData

`func (o *CredentialRepresentation) GetSecretData() string`

GetSecretData returns the SecretData field if non-nil, zero value otherwise.

### GetSecretDataOk

`func (o *CredentialRepresentation) GetSecretDataOk() (*string, bool)`

GetSecretDataOk returns a tuple with the SecretData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretData

`func (o *CredentialRepresentation) SetSecretData(v string)`

SetSecretData sets SecretData field to given value.

### HasSecretData

`func (o *CredentialRepresentation) HasSecretData() bool`

HasSecretData returns a boolean if a field has been set.

### GetCredentialData

`func (o *CredentialRepresentation) GetCredentialData() string`

GetCredentialData returns the CredentialData field if non-nil, zero value otherwise.

### GetCredentialDataOk

`func (o *CredentialRepresentation) GetCredentialDataOk() (*string, bool)`

GetCredentialDataOk returns a tuple with the CredentialData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialData

`func (o *CredentialRepresentation) SetCredentialData(v string)`

SetCredentialData sets CredentialData field to given value.

### HasCredentialData

`func (o *CredentialRepresentation) HasCredentialData() bool`

HasCredentialData returns a boolean if a field has been set.

### GetPriority

`func (o *CredentialRepresentation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CredentialRepresentation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CredentialRepresentation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CredentialRepresentation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetValue

`func (o *CredentialRepresentation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CredentialRepresentation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CredentialRepresentation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CredentialRepresentation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTemporary

`func (o *CredentialRepresentation) GetTemporary() bool`

GetTemporary returns the Temporary field if non-nil, zero value otherwise.

### GetTemporaryOk

`func (o *CredentialRepresentation) GetTemporaryOk() (*bool, bool)`

GetTemporaryOk returns a tuple with the Temporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporary

`func (o *CredentialRepresentation) SetTemporary(v bool)`

SetTemporary sets Temporary field to given value.

### HasTemporary

`func (o *CredentialRepresentation) HasTemporary() bool`

HasTemporary returns a boolean if a field has been set.

### GetDevice

`func (o *CredentialRepresentation) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CredentialRepresentation) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CredentialRepresentation) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CredentialRepresentation) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHashedSaltedValue

`func (o *CredentialRepresentation) GetHashedSaltedValue() string`

GetHashedSaltedValue returns the HashedSaltedValue field if non-nil, zero value otherwise.

### GetHashedSaltedValueOk

`func (o *CredentialRepresentation) GetHashedSaltedValueOk() (*string, bool)`

GetHashedSaltedValueOk returns a tuple with the HashedSaltedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashedSaltedValue

`func (o *CredentialRepresentation) SetHashedSaltedValue(v string)`

SetHashedSaltedValue sets HashedSaltedValue field to given value.

### HasHashedSaltedValue

`func (o *CredentialRepresentation) HasHashedSaltedValue() bool`

HasHashedSaltedValue returns a boolean if a field has been set.

### GetSalt

`func (o *CredentialRepresentation) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *CredentialRepresentation) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *CredentialRepresentation) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *CredentialRepresentation) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetHashIterations

`func (o *CredentialRepresentation) GetHashIterations() int32`

GetHashIterations returns the HashIterations field if non-nil, zero value otherwise.

### GetHashIterationsOk

`func (o *CredentialRepresentation) GetHashIterationsOk() (*int32, bool)`

GetHashIterationsOk returns a tuple with the HashIterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashIterations

`func (o *CredentialRepresentation) SetHashIterations(v int32)`

SetHashIterations sets HashIterations field to given value.

### HasHashIterations

`func (o *CredentialRepresentation) HasHashIterations() bool`

HasHashIterations returns a boolean if a field has been set.

### GetCounter

`func (o *CredentialRepresentation) GetCounter() int32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *CredentialRepresentation) GetCounterOk() (*int32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *CredentialRepresentation) SetCounter(v int32)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *CredentialRepresentation) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetAlgorithm

`func (o *CredentialRepresentation) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CredentialRepresentation) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CredentialRepresentation) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *CredentialRepresentation) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetDigits

`func (o *CredentialRepresentation) GetDigits() int32`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *CredentialRepresentation) GetDigitsOk() (*int32, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *CredentialRepresentation) SetDigits(v int32)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *CredentialRepresentation) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetPeriod

`func (o *CredentialRepresentation) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CredentialRepresentation) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CredentialRepresentation) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CredentialRepresentation) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetConfig

`func (o *CredentialRepresentation) GetConfig() map[string][]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CredentialRepresentation) GetConfigOk() (*map[string][]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CredentialRepresentation) SetConfig(v map[string][]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CredentialRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


