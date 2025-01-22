# ClientTemplateRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**FullScopeAllowed** | Pointer to **bool** |  | [optional] 
**BearerOnly** | Pointer to **bool** |  | [optional] 
**ConsentRequired** | Pointer to **bool** |  | [optional] 
**StandardFlowEnabled** | Pointer to **bool** |  | [optional] 
**ImplicitFlowEnabled** | Pointer to **bool** |  | [optional] 
**DirectAccessGrantsEnabled** | Pointer to **bool** |  | [optional] 
**ServiceAccountsEnabled** | Pointer to **bool** |  | [optional] 
**PublicClient** | Pointer to **bool** |  | [optional] 
**FrontchannelLogout** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**ProtocolMappers** | Pointer to [**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) |  | [optional] 

## Methods

### NewClientTemplateRepresentation

`func NewClientTemplateRepresentation() *ClientTemplateRepresentation`

NewClientTemplateRepresentation instantiates a new ClientTemplateRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTemplateRepresentationWithDefaults

`func NewClientTemplateRepresentationWithDefaults() *ClientTemplateRepresentation`

NewClientTemplateRepresentationWithDefaults instantiates a new ClientTemplateRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientTemplateRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientTemplateRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientTemplateRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientTemplateRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClientTemplateRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientTemplateRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientTemplateRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientTemplateRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ClientTemplateRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClientTemplateRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClientTemplateRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClientTemplateRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProtocol

`func (o *ClientTemplateRepresentation) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ClientTemplateRepresentation) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ClientTemplateRepresentation) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ClientTemplateRepresentation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetFullScopeAllowed

`func (o *ClientTemplateRepresentation) GetFullScopeAllowed() bool`

GetFullScopeAllowed returns the FullScopeAllowed field if non-nil, zero value otherwise.

### GetFullScopeAllowedOk

`func (o *ClientTemplateRepresentation) GetFullScopeAllowedOk() (*bool, bool)`

GetFullScopeAllowedOk returns a tuple with the FullScopeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScopeAllowed

`func (o *ClientTemplateRepresentation) SetFullScopeAllowed(v bool)`

SetFullScopeAllowed sets FullScopeAllowed field to given value.

### HasFullScopeAllowed

`func (o *ClientTemplateRepresentation) HasFullScopeAllowed() bool`

HasFullScopeAllowed returns a boolean if a field has been set.

### GetBearerOnly

`func (o *ClientTemplateRepresentation) GetBearerOnly() bool`

GetBearerOnly returns the BearerOnly field if non-nil, zero value otherwise.

### GetBearerOnlyOk

`func (o *ClientTemplateRepresentation) GetBearerOnlyOk() (*bool, bool)`

GetBearerOnlyOk returns a tuple with the BearerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerOnly

`func (o *ClientTemplateRepresentation) SetBearerOnly(v bool)`

SetBearerOnly sets BearerOnly field to given value.

### HasBearerOnly

`func (o *ClientTemplateRepresentation) HasBearerOnly() bool`

HasBearerOnly returns a boolean if a field has been set.

### GetConsentRequired

`func (o *ClientTemplateRepresentation) GetConsentRequired() bool`

GetConsentRequired returns the ConsentRequired field if non-nil, zero value otherwise.

### GetConsentRequiredOk

`func (o *ClientTemplateRepresentation) GetConsentRequiredOk() (*bool, bool)`

GetConsentRequiredOk returns a tuple with the ConsentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRequired

`func (o *ClientTemplateRepresentation) SetConsentRequired(v bool)`

SetConsentRequired sets ConsentRequired field to given value.

### HasConsentRequired

`func (o *ClientTemplateRepresentation) HasConsentRequired() bool`

HasConsentRequired returns a boolean if a field has been set.

### GetStandardFlowEnabled

`func (o *ClientTemplateRepresentation) GetStandardFlowEnabled() bool`

GetStandardFlowEnabled returns the StandardFlowEnabled field if non-nil, zero value otherwise.

### GetStandardFlowEnabledOk

`func (o *ClientTemplateRepresentation) GetStandardFlowEnabledOk() (*bool, bool)`

GetStandardFlowEnabledOk returns a tuple with the StandardFlowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardFlowEnabled

`func (o *ClientTemplateRepresentation) SetStandardFlowEnabled(v bool)`

SetStandardFlowEnabled sets StandardFlowEnabled field to given value.

### HasStandardFlowEnabled

`func (o *ClientTemplateRepresentation) HasStandardFlowEnabled() bool`

HasStandardFlowEnabled returns a boolean if a field has been set.

### GetImplicitFlowEnabled

`func (o *ClientTemplateRepresentation) GetImplicitFlowEnabled() bool`

GetImplicitFlowEnabled returns the ImplicitFlowEnabled field if non-nil, zero value otherwise.

### GetImplicitFlowEnabledOk

`func (o *ClientTemplateRepresentation) GetImplicitFlowEnabledOk() (*bool, bool)`

GetImplicitFlowEnabledOk returns a tuple with the ImplicitFlowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitFlowEnabled

`func (o *ClientTemplateRepresentation) SetImplicitFlowEnabled(v bool)`

SetImplicitFlowEnabled sets ImplicitFlowEnabled field to given value.

### HasImplicitFlowEnabled

`func (o *ClientTemplateRepresentation) HasImplicitFlowEnabled() bool`

HasImplicitFlowEnabled returns a boolean if a field has been set.

### GetDirectAccessGrantsEnabled

`func (o *ClientTemplateRepresentation) GetDirectAccessGrantsEnabled() bool`

GetDirectAccessGrantsEnabled returns the DirectAccessGrantsEnabled field if non-nil, zero value otherwise.

### GetDirectAccessGrantsEnabledOk

`func (o *ClientTemplateRepresentation) GetDirectAccessGrantsEnabledOk() (*bool, bool)`

GetDirectAccessGrantsEnabledOk returns a tuple with the DirectAccessGrantsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectAccessGrantsEnabled

`func (o *ClientTemplateRepresentation) SetDirectAccessGrantsEnabled(v bool)`

SetDirectAccessGrantsEnabled sets DirectAccessGrantsEnabled field to given value.

### HasDirectAccessGrantsEnabled

`func (o *ClientTemplateRepresentation) HasDirectAccessGrantsEnabled() bool`

HasDirectAccessGrantsEnabled returns a boolean if a field has been set.

### GetServiceAccountsEnabled

`func (o *ClientTemplateRepresentation) GetServiceAccountsEnabled() bool`

GetServiceAccountsEnabled returns the ServiceAccountsEnabled field if non-nil, zero value otherwise.

### GetServiceAccountsEnabledOk

`func (o *ClientTemplateRepresentation) GetServiceAccountsEnabledOk() (*bool, bool)`

GetServiceAccountsEnabledOk returns a tuple with the ServiceAccountsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountsEnabled

`func (o *ClientTemplateRepresentation) SetServiceAccountsEnabled(v bool)`

SetServiceAccountsEnabled sets ServiceAccountsEnabled field to given value.

### HasServiceAccountsEnabled

`func (o *ClientTemplateRepresentation) HasServiceAccountsEnabled() bool`

HasServiceAccountsEnabled returns a boolean if a field has been set.

### GetPublicClient

`func (o *ClientTemplateRepresentation) GetPublicClient() bool`

GetPublicClient returns the PublicClient field if non-nil, zero value otherwise.

### GetPublicClientOk

`func (o *ClientTemplateRepresentation) GetPublicClientOk() (*bool, bool)`

GetPublicClientOk returns a tuple with the PublicClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicClient

`func (o *ClientTemplateRepresentation) SetPublicClient(v bool)`

SetPublicClient sets PublicClient field to given value.

### HasPublicClient

`func (o *ClientTemplateRepresentation) HasPublicClient() bool`

HasPublicClient returns a boolean if a field has been set.

### GetFrontchannelLogout

`func (o *ClientTemplateRepresentation) GetFrontchannelLogout() bool`

GetFrontchannelLogout returns the FrontchannelLogout field if non-nil, zero value otherwise.

### GetFrontchannelLogoutOk

`func (o *ClientTemplateRepresentation) GetFrontchannelLogoutOk() (*bool, bool)`

GetFrontchannelLogoutOk returns a tuple with the FrontchannelLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontchannelLogout

`func (o *ClientTemplateRepresentation) SetFrontchannelLogout(v bool)`

SetFrontchannelLogout sets FrontchannelLogout field to given value.

### HasFrontchannelLogout

`func (o *ClientTemplateRepresentation) HasFrontchannelLogout() bool`

HasFrontchannelLogout returns a boolean if a field has been set.

### GetAttributes

`func (o *ClientTemplateRepresentation) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ClientTemplateRepresentation) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ClientTemplateRepresentation) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ClientTemplateRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetProtocolMappers

`func (o *ClientTemplateRepresentation) GetProtocolMappers() []ProtocolMapperRepresentation`

GetProtocolMappers returns the ProtocolMappers field if non-nil, zero value otherwise.

### GetProtocolMappersOk

`func (o *ClientTemplateRepresentation) GetProtocolMappersOk() (*[]ProtocolMapperRepresentation, bool)`

GetProtocolMappersOk returns a tuple with the ProtocolMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMappers

`func (o *ClientTemplateRepresentation) SetProtocolMappers(v []ProtocolMapperRepresentation)`

SetProtocolMappers sets ProtocolMappers field to given value.

### HasProtocolMappers

`func (o *ClientTemplateRepresentation) HasProtocolMappers() bool`

HasProtocolMappers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


