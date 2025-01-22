# OrganizationRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 
**Domains** | Pointer to [**[]OrganizationDomainRepresentation**](OrganizationDomainRepresentation.md) |  | [optional] 
**Members** | Pointer to [**[]MemberRepresentation**](MemberRepresentation.md) |  | [optional] 
**IdentityProviders** | Pointer to [**[]IdentityProviderRepresentation**](IdentityProviderRepresentation.md) |  | [optional] 

## Methods

### NewOrganizationRepresentation

`func NewOrganizationRepresentation() *OrganizationRepresentation`

NewOrganizationRepresentation instantiates a new OrganizationRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationRepresentationWithDefaults

`func NewOrganizationRepresentationWithDefaults() *OrganizationRepresentation`

NewOrganizationRepresentationWithDefaults instantiates a new OrganizationRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAlias

`func (o *OrganizationRepresentation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *OrganizationRepresentation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *OrganizationRepresentation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *OrganizationRepresentation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEnabled

`func (o *OrganizationRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *OrganizationRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *OrganizationRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *OrganizationRepresentation) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *OrganizationRepresentation) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *OrganizationRepresentation) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *OrganizationRepresentation) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetAttributes

`func (o *OrganizationRepresentation) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *OrganizationRepresentation) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *OrganizationRepresentation) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *OrganizationRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDomains

`func (o *OrganizationRepresentation) GetDomains() []OrganizationDomainRepresentation`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *OrganizationRepresentation) GetDomainsOk() (*[]OrganizationDomainRepresentation, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *OrganizationRepresentation) SetDomains(v []OrganizationDomainRepresentation)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *OrganizationRepresentation) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetMembers

`func (o *OrganizationRepresentation) GetMembers() []MemberRepresentation`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationRepresentation) GetMembersOk() (*[]MemberRepresentation, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationRepresentation) SetMembers(v []MemberRepresentation)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OrganizationRepresentation) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetIdentityProviders

`func (o *OrganizationRepresentation) GetIdentityProviders() []IdentityProviderRepresentation`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *OrganizationRepresentation) GetIdentityProvidersOk() (*[]IdentityProviderRepresentation, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *OrganizationRepresentation) SetIdentityProviders(v []IdentityProviderRepresentation)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *OrganizationRepresentation) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


