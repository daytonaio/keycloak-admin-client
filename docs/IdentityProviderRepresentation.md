# IdentityProviderRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**UpdateProfileFirstLoginMode** | Pointer to **string** |  | [optional] 
**TrustEmail** | Pointer to **bool** |  | [optional] 
**StoreToken** | Pointer to **bool** |  | [optional] 
**AddReadTokenRoleOnCreate** | Pointer to **bool** |  | [optional] 
**AuthenticateByDefault** | Pointer to **bool** |  | [optional] 
**LinkOnly** | Pointer to **bool** |  | [optional] 
**HideOnLogin** | Pointer to **bool** |  | [optional] 
**FirstBrokerLoginFlowAlias** | Pointer to **string** |  | [optional] 
**PostBrokerLoginFlowAlias** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]string** |  | [optional] 
**UpdateProfileFirstLogin** | Pointer to **bool** |  | [optional] 

## Methods

### NewIdentityProviderRepresentation

`func NewIdentityProviderRepresentation() *IdentityProviderRepresentation`

NewIdentityProviderRepresentation instantiates a new IdentityProviderRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderRepresentationWithDefaults

`func NewIdentityProviderRepresentationWithDefaults() *IdentityProviderRepresentation`

NewIdentityProviderRepresentationWithDefaults instantiates a new IdentityProviderRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *IdentityProviderRepresentation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *IdentityProviderRepresentation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *IdentityProviderRepresentation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *IdentityProviderRepresentation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDisplayName

`func (o *IdentityProviderRepresentation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityProviderRepresentation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityProviderRepresentation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityProviderRepresentation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInternalId

`func (o *IdentityProviderRepresentation) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *IdentityProviderRepresentation) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *IdentityProviderRepresentation) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *IdentityProviderRepresentation) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *IdentityProviderRepresentation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *IdentityProviderRepresentation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *IdentityProviderRepresentation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *IdentityProviderRepresentation) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityProviderRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityProviderRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityProviderRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityProviderRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUpdateProfileFirstLoginMode

`func (o *IdentityProviderRepresentation) GetUpdateProfileFirstLoginMode() string`

GetUpdateProfileFirstLoginMode returns the UpdateProfileFirstLoginMode field if non-nil, zero value otherwise.

### GetUpdateProfileFirstLoginModeOk

`func (o *IdentityProviderRepresentation) GetUpdateProfileFirstLoginModeOk() (*string, bool)`

GetUpdateProfileFirstLoginModeOk returns a tuple with the UpdateProfileFirstLoginMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProfileFirstLoginMode

`func (o *IdentityProviderRepresentation) SetUpdateProfileFirstLoginMode(v string)`

SetUpdateProfileFirstLoginMode sets UpdateProfileFirstLoginMode field to given value.

### HasUpdateProfileFirstLoginMode

`func (o *IdentityProviderRepresentation) HasUpdateProfileFirstLoginMode() bool`

HasUpdateProfileFirstLoginMode returns a boolean if a field has been set.

### GetTrustEmail

`func (o *IdentityProviderRepresentation) GetTrustEmail() bool`

GetTrustEmail returns the TrustEmail field if non-nil, zero value otherwise.

### GetTrustEmailOk

`func (o *IdentityProviderRepresentation) GetTrustEmailOk() (*bool, bool)`

GetTrustEmailOk returns a tuple with the TrustEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustEmail

`func (o *IdentityProviderRepresentation) SetTrustEmail(v bool)`

SetTrustEmail sets TrustEmail field to given value.

### HasTrustEmail

`func (o *IdentityProviderRepresentation) HasTrustEmail() bool`

HasTrustEmail returns a boolean if a field has been set.

### GetStoreToken

`func (o *IdentityProviderRepresentation) GetStoreToken() bool`

GetStoreToken returns the StoreToken field if non-nil, zero value otherwise.

### GetStoreTokenOk

`func (o *IdentityProviderRepresentation) GetStoreTokenOk() (*bool, bool)`

GetStoreTokenOk returns a tuple with the StoreToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreToken

`func (o *IdentityProviderRepresentation) SetStoreToken(v bool)`

SetStoreToken sets StoreToken field to given value.

### HasStoreToken

`func (o *IdentityProviderRepresentation) HasStoreToken() bool`

HasStoreToken returns a boolean if a field has been set.

### GetAddReadTokenRoleOnCreate

`func (o *IdentityProviderRepresentation) GetAddReadTokenRoleOnCreate() bool`

GetAddReadTokenRoleOnCreate returns the AddReadTokenRoleOnCreate field if non-nil, zero value otherwise.

### GetAddReadTokenRoleOnCreateOk

`func (o *IdentityProviderRepresentation) GetAddReadTokenRoleOnCreateOk() (*bool, bool)`

GetAddReadTokenRoleOnCreateOk returns a tuple with the AddReadTokenRoleOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddReadTokenRoleOnCreate

`func (o *IdentityProviderRepresentation) SetAddReadTokenRoleOnCreate(v bool)`

SetAddReadTokenRoleOnCreate sets AddReadTokenRoleOnCreate field to given value.

### HasAddReadTokenRoleOnCreate

`func (o *IdentityProviderRepresentation) HasAddReadTokenRoleOnCreate() bool`

HasAddReadTokenRoleOnCreate returns a boolean if a field has been set.

### GetAuthenticateByDefault

`func (o *IdentityProviderRepresentation) GetAuthenticateByDefault() bool`

GetAuthenticateByDefault returns the AuthenticateByDefault field if non-nil, zero value otherwise.

### GetAuthenticateByDefaultOk

`func (o *IdentityProviderRepresentation) GetAuthenticateByDefaultOk() (*bool, bool)`

GetAuthenticateByDefaultOk returns a tuple with the AuthenticateByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticateByDefault

`func (o *IdentityProviderRepresentation) SetAuthenticateByDefault(v bool)`

SetAuthenticateByDefault sets AuthenticateByDefault field to given value.

### HasAuthenticateByDefault

`func (o *IdentityProviderRepresentation) HasAuthenticateByDefault() bool`

HasAuthenticateByDefault returns a boolean if a field has been set.

### GetLinkOnly

`func (o *IdentityProviderRepresentation) GetLinkOnly() bool`

GetLinkOnly returns the LinkOnly field if non-nil, zero value otherwise.

### GetLinkOnlyOk

`func (o *IdentityProviderRepresentation) GetLinkOnlyOk() (*bool, bool)`

GetLinkOnlyOk returns a tuple with the LinkOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkOnly

`func (o *IdentityProviderRepresentation) SetLinkOnly(v bool)`

SetLinkOnly sets LinkOnly field to given value.

### HasLinkOnly

`func (o *IdentityProviderRepresentation) HasLinkOnly() bool`

HasLinkOnly returns a boolean if a field has been set.

### GetHideOnLogin

`func (o *IdentityProviderRepresentation) GetHideOnLogin() bool`

GetHideOnLogin returns the HideOnLogin field if non-nil, zero value otherwise.

### GetHideOnLoginOk

`func (o *IdentityProviderRepresentation) GetHideOnLoginOk() (*bool, bool)`

GetHideOnLoginOk returns a tuple with the HideOnLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideOnLogin

`func (o *IdentityProviderRepresentation) SetHideOnLogin(v bool)`

SetHideOnLogin sets HideOnLogin field to given value.

### HasHideOnLogin

`func (o *IdentityProviderRepresentation) HasHideOnLogin() bool`

HasHideOnLogin returns a boolean if a field has been set.

### GetFirstBrokerLoginFlowAlias

`func (o *IdentityProviderRepresentation) GetFirstBrokerLoginFlowAlias() string`

GetFirstBrokerLoginFlowAlias returns the FirstBrokerLoginFlowAlias field if non-nil, zero value otherwise.

### GetFirstBrokerLoginFlowAliasOk

`func (o *IdentityProviderRepresentation) GetFirstBrokerLoginFlowAliasOk() (*string, bool)`

GetFirstBrokerLoginFlowAliasOk returns a tuple with the FirstBrokerLoginFlowAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstBrokerLoginFlowAlias

`func (o *IdentityProviderRepresentation) SetFirstBrokerLoginFlowAlias(v string)`

SetFirstBrokerLoginFlowAlias sets FirstBrokerLoginFlowAlias field to given value.

### HasFirstBrokerLoginFlowAlias

`func (o *IdentityProviderRepresentation) HasFirstBrokerLoginFlowAlias() bool`

HasFirstBrokerLoginFlowAlias returns a boolean if a field has been set.

### GetPostBrokerLoginFlowAlias

`func (o *IdentityProviderRepresentation) GetPostBrokerLoginFlowAlias() string`

GetPostBrokerLoginFlowAlias returns the PostBrokerLoginFlowAlias field if non-nil, zero value otherwise.

### GetPostBrokerLoginFlowAliasOk

`func (o *IdentityProviderRepresentation) GetPostBrokerLoginFlowAliasOk() (*string, bool)`

GetPostBrokerLoginFlowAliasOk returns a tuple with the PostBrokerLoginFlowAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBrokerLoginFlowAlias

`func (o *IdentityProviderRepresentation) SetPostBrokerLoginFlowAlias(v string)`

SetPostBrokerLoginFlowAlias sets PostBrokerLoginFlowAlias field to given value.

### HasPostBrokerLoginFlowAlias

`func (o *IdentityProviderRepresentation) HasPostBrokerLoginFlowAlias() bool`

HasPostBrokerLoginFlowAlias returns a boolean if a field has been set.

### GetOrganizationId

`func (o *IdentityProviderRepresentation) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *IdentityProviderRepresentation) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *IdentityProviderRepresentation) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *IdentityProviderRepresentation) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfig

`func (o *IdentityProviderRepresentation) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IdentityProviderRepresentation) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IdentityProviderRepresentation) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IdentityProviderRepresentation) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetUpdateProfileFirstLogin

`func (o *IdentityProviderRepresentation) GetUpdateProfileFirstLogin() bool`

GetUpdateProfileFirstLogin returns the UpdateProfileFirstLogin field if non-nil, zero value otherwise.

### GetUpdateProfileFirstLoginOk

`func (o *IdentityProviderRepresentation) GetUpdateProfileFirstLoginOk() (*bool, bool)`

GetUpdateProfileFirstLoginOk returns a tuple with the UpdateProfileFirstLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProfileFirstLogin

`func (o *IdentityProviderRepresentation) SetUpdateProfileFirstLogin(v bool)`

SetUpdateProfileFirstLogin sets UpdateProfileFirstLogin field to given value.

### HasUpdateProfileFirstLogin

`func (o *IdentityProviderRepresentation) HasUpdateProfileFirstLogin() bool`

HasUpdateProfileFirstLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


