# AuthenticatorConfigInfoRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**HelpText** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]ConfigPropertyRepresentation**](ConfigPropertyRepresentation.md) |  | [optional] 

## Methods

### NewAuthenticatorConfigInfoRepresentation

`func NewAuthenticatorConfigInfoRepresentation() *AuthenticatorConfigInfoRepresentation`

NewAuthenticatorConfigInfoRepresentation instantiates a new AuthenticatorConfigInfoRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorConfigInfoRepresentationWithDefaults

`func NewAuthenticatorConfigInfoRepresentationWithDefaults() *AuthenticatorConfigInfoRepresentation`

NewAuthenticatorConfigInfoRepresentationWithDefaults instantiates a new AuthenticatorConfigInfoRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticatorConfigInfoRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticatorConfigInfoRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticatorConfigInfoRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticatorConfigInfoRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderId

`func (o *AuthenticatorConfigInfoRepresentation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AuthenticatorConfigInfoRepresentation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AuthenticatorConfigInfoRepresentation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AuthenticatorConfigInfoRepresentation) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetHelpText

`func (o *AuthenticatorConfigInfoRepresentation) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *AuthenticatorConfigInfoRepresentation) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *AuthenticatorConfigInfoRepresentation) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *AuthenticatorConfigInfoRepresentation) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetProperties

`func (o *AuthenticatorConfigInfoRepresentation) GetProperties() []ConfigPropertyRepresentation`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AuthenticatorConfigInfoRepresentation) GetPropertiesOk() (*[]ConfigPropertyRepresentation, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AuthenticatorConfigInfoRepresentation) SetProperties(v []ConfigPropertyRepresentation)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AuthenticatorConfigInfoRepresentation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


