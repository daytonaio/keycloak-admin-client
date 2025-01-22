# KeyMetadataRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | Pointer to **string** |  | [optional] 
**ProviderPriority** | Pointer to **int64** |  | [optional] 
**Kid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Algorithm** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**Use** | Pointer to [**KeyUse**](KeyUse.md) |  | [optional] 
**ValidTo** | Pointer to **int64** |  | [optional] 

## Methods

### NewKeyMetadataRepresentation

`func NewKeyMetadataRepresentation() *KeyMetadataRepresentation`

NewKeyMetadataRepresentation instantiates a new KeyMetadataRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyMetadataRepresentationWithDefaults

`func NewKeyMetadataRepresentationWithDefaults() *KeyMetadataRepresentation`

NewKeyMetadataRepresentationWithDefaults instantiates a new KeyMetadataRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *KeyMetadataRepresentation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *KeyMetadataRepresentation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *KeyMetadataRepresentation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *KeyMetadataRepresentation) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetProviderPriority

`func (o *KeyMetadataRepresentation) GetProviderPriority() int64`

GetProviderPriority returns the ProviderPriority field if non-nil, zero value otherwise.

### GetProviderPriorityOk

`func (o *KeyMetadataRepresentation) GetProviderPriorityOk() (*int64, bool)`

GetProviderPriorityOk returns a tuple with the ProviderPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderPriority

`func (o *KeyMetadataRepresentation) SetProviderPriority(v int64)`

SetProviderPriority sets ProviderPriority field to given value.

### HasProviderPriority

`func (o *KeyMetadataRepresentation) HasProviderPriority() bool`

HasProviderPriority returns a boolean if a field has been set.

### GetKid

`func (o *KeyMetadataRepresentation) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *KeyMetadataRepresentation) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *KeyMetadataRepresentation) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *KeyMetadataRepresentation) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetStatus

`func (o *KeyMetadataRepresentation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyMetadataRepresentation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyMetadataRepresentation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyMetadataRepresentation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *KeyMetadataRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyMetadataRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyMetadataRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyMetadataRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAlgorithm

`func (o *KeyMetadataRepresentation) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *KeyMetadataRepresentation) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *KeyMetadataRepresentation) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *KeyMetadataRepresentation) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetPublicKey

`func (o *KeyMetadataRepresentation) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KeyMetadataRepresentation) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KeyMetadataRepresentation) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KeyMetadataRepresentation) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetCertificate

`func (o *KeyMetadataRepresentation) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *KeyMetadataRepresentation) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *KeyMetadataRepresentation) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *KeyMetadataRepresentation) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetUse

`func (o *KeyMetadataRepresentation) GetUse() KeyUse`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *KeyMetadataRepresentation) GetUseOk() (*KeyUse, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *KeyMetadataRepresentation) SetUse(v KeyUse)`

SetUse sets Use field to given value.

### HasUse

`func (o *KeyMetadataRepresentation) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetValidTo

`func (o *KeyMetadataRepresentation) GetValidTo() int64`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *KeyMetadataRepresentation) GetValidToOk() (*int64, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *KeyMetadataRepresentation) SetValidTo(v int64)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *KeyMetadataRepresentation) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


