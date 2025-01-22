# ApplicationRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**RootUrl** | Pointer to **string** |  | [optional] 
**AdminUrl** | Pointer to **string** |  | [optional] 
**BaseUrl** | Pointer to **string** |  | [optional] 
**SurrogateAuthRequired** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AlwaysDisplayInConsole** | Pointer to **bool** |  | [optional] 
**ClientAuthenticatorType** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**RegistrationAccessToken** | Pointer to **string** |  | [optional] 
**DefaultRoles** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**WebOrigins** | Pointer to **[]string** |  | [optional] 
**NotBefore** | Pointer to **int32** |  | [optional] 
**BearerOnly** | Pointer to **bool** |  | [optional] 
**ConsentRequired** | Pointer to **bool** |  | [optional] 
**StandardFlowEnabled** | Pointer to **bool** |  | [optional] 
**ImplicitFlowEnabled** | Pointer to **bool** |  | [optional] 
**DirectAccessGrantsEnabled** | Pointer to **bool** |  | [optional] 
**ServiceAccountsEnabled** | Pointer to **bool** |  | [optional] 
**AuthorizationServicesEnabled** | Pointer to **bool** |  | [optional] 
**DirectGrantsOnly** | Pointer to **bool** |  | [optional] 
**PublicClient** | Pointer to **bool** |  | [optional] 
**FrontchannelLogout** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**AuthenticationFlowBindingOverrides** | Pointer to **map[string]string** |  | [optional] 
**FullScopeAllowed** | Pointer to **bool** |  | [optional] 
**NodeReRegistrationTimeout** | Pointer to **int32** |  | [optional] 
**RegisteredNodes** | Pointer to **map[string]int32** |  | [optional] 
**ProtocolMappers** | Pointer to [**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) |  | [optional] 
**ClientTemplate** | Pointer to **string** |  | [optional] 
**UseTemplateConfig** | Pointer to **bool** |  | [optional] 
**UseTemplateScope** | Pointer to **bool** |  | [optional] 
**UseTemplateMappers** | Pointer to **bool** |  | [optional] 
**DefaultClientScopes** | Pointer to **[]string** |  | [optional] 
**OptionalClientScopes** | Pointer to **[]string** |  | [optional] 
**AuthorizationSettings** | Pointer to [**ResourceServerRepresentation**](ResourceServerRepresentation.md) |  | [optional] 
**Access** | Pointer to **map[string]bool** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Claims** | Pointer to [**ClaimRepresentation**](ClaimRepresentation.md) |  | [optional] 

## Methods

### NewApplicationRepresentation

`func NewApplicationRepresentation() *ApplicationRepresentation`

NewApplicationRepresentation instantiates a new ApplicationRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRepresentationWithDefaults

`func NewApplicationRepresentationWithDefaults() *ApplicationRepresentation`

NewApplicationRepresentationWithDefaults instantiates a new ApplicationRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *ApplicationRepresentation) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApplicationRepresentation) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApplicationRepresentation) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ApplicationRepresentation) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ApplicationRepresentation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRepresentation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRepresentation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApplicationRepresentation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRootUrl

`func (o *ApplicationRepresentation) GetRootUrl() string`

GetRootUrl returns the RootUrl field if non-nil, zero value otherwise.

### GetRootUrlOk

`func (o *ApplicationRepresentation) GetRootUrlOk() (*string, bool)`

GetRootUrlOk returns a tuple with the RootUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUrl

`func (o *ApplicationRepresentation) SetRootUrl(v string)`

SetRootUrl sets RootUrl field to given value.

### HasRootUrl

`func (o *ApplicationRepresentation) HasRootUrl() bool`

HasRootUrl returns a boolean if a field has been set.

### GetAdminUrl

`func (o *ApplicationRepresentation) GetAdminUrl() string`

GetAdminUrl returns the AdminUrl field if non-nil, zero value otherwise.

### GetAdminUrlOk

`func (o *ApplicationRepresentation) GetAdminUrlOk() (*string, bool)`

GetAdminUrlOk returns a tuple with the AdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUrl

`func (o *ApplicationRepresentation) SetAdminUrl(v string)`

SetAdminUrl sets AdminUrl field to given value.

### HasAdminUrl

`func (o *ApplicationRepresentation) HasAdminUrl() bool`

HasAdminUrl returns a boolean if a field has been set.

### GetBaseUrl

`func (o *ApplicationRepresentation) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ApplicationRepresentation) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ApplicationRepresentation) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *ApplicationRepresentation) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetSurrogateAuthRequired

`func (o *ApplicationRepresentation) GetSurrogateAuthRequired() bool`

GetSurrogateAuthRequired returns the SurrogateAuthRequired field if non-nil, zero value otherwise.

### GetSurrogateAuthRequiredOk

`func (o *ApplicationRepresentation) GetSurrogateAuthRequiredOk() (*bool, bool)`

GetSurrogateAuthRequiredOk returns a tuple with the SurrogateAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurrogateAuthRequired

`func (o *ApplicationRepresentation) SetSurrogateAuthRequired(v bool)`

SetSurrogateAuthRequired sets SurrogateAuthRequired field to given value.

### HasSurrogateAuthRequired

`func (o *ApplicationRepresentation) HasSurrogateAuthRequired() bool`

HasSurrogateAuthRequired returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlwaysDisplayInConsole

`func (o *ApplicationRepresentation) GetAlwaysDisplayInConsole() bool`

GetAlwaysDisplayInConsole returns the AlwaysDisplayInConsole field if non-nil, zero value otherwise.

### GetAlwaysDisplayInConsoleOk

`func (o *ApplicationRepresentation) GetAlwaysDisplayInConsoleOk() (*bool, bool)`

GetAlwaysDisplayInConsoleOk returns a tuple with the AlwaysDisplayInConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysDisplayInConsole

`func (o *ApplicationRepresentation) SetAlwaysDisplayInConsole(v bool)`

SetAlwaysDisplayInConsole sets AlwaysDisplayInConsole field to given value.

### HasAlwaysDisplayInConsole

`func (o *ApplicationRepresentation) HasAlwaysDisplayInConsole() bool`

HasAlwaysDisplayInConsole returns a boolean if a field has been set.

### GetClientAuthenticatorType

`func (o *ApplicationRepresentation) GetClientAuthenticatorType() string`

GetClientAuthenticatorType returns the ClientAuthenticatorType field if non-nil, zero value otherwise.

### GetClientAuthenticatorTypeOk

`func (o *ApplicationRepresentation) GetClientAuthenticatorTypeOk() (*string, bool)`

GetClientAuthenticatorTypeOk returns a tuple with the ClientAuthenticatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticatorType

`func (o *ApplicationRepresentation) SetClientAuthenticatorType(v string)`

SetClientAuthenticatorType sets ClientAuthenticatorType field to given value.

### HasClientAuthenticatorType

`func (o *ApplicationRepresentation) HasClientAuthenticatorType() bool`

HasClientAuthenticatorType returns a boolean if a field has been set.

### GetSecret

`func (o *ApplicationRepresentation) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApplicationRepresentation) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApplicationRepresentation) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ApplicationRepresentation) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRegistrationAccessToken

`func (o *ApplicationRepresentation) GetRegistrationAccessToken() string`

GetRegistrationAccessToken returns the RegistrationAccessToken field if non-nil, zero value otherwise.

### GetRegistrationAccessTokenOk

`func (o *ApplicationRepresentation) GetRegistrationAccessTokenOk() (*string, bool)`

GetRegistrationAccessTokenOk returns a tuple with the RegistrationAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationAccessToken

`func (o *ApplicationRepresentation) SetRegistrationAccessToken(v string)`

SetRegistrationAccessToken sets RegistrationAccessToken field to given value.

### HasRegistrationAccessToken

`func (o *ApplicationRepresentation) HasRegistrationAccessToken() bool`

HasRegistrationAccessToken returns a boolean if a field has been set.

### GetDefaultRoles

`func (o *ApplicationRepresentation) GetDefaultRoles() []string`

GetDefaultRoles returns the DefaultRoles field if non-nil, zero value otherwise.

### GetDefaultRolesOk

`func (o *ApplicationRepresentation) GetDefaultRolesOk() (*[]string, bool)`

GetDefaultRolesOk returns a tuple with the DefaultRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoles

`func (o *ApplicationRepresentation) SetDefaultRoles(v []string)`

SetDefaultRoles sets DefaultRoles field to given value.

### HasDefaultRoles

`func (o *ApplicationRepresentation) HasDefaultRoles() bool`

HasDefaultRoles returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ApplicationRepresentation) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ApplicationRepresentation) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ApplicationRepresentation) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ApplicationRepresentation) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetWebOrigins

`func (o *ApplicationRepresentation) GetWebOrigins() []string`

GetWebOrigins returns the WebOrigins field if non-nil, zero value otherwise.

### GetWebOriginsOk

`func (o *ApplicationRepresentation) GetWebOriginsOk() (*[]string, bool)`

GetWebOriginsOk returns a tuple with the WebOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebOrigins

`func (o *ApplicationRepresentation) SetWebOrigins(v []string)`

SetWebOrigins sets WebOrigins field to given value.

### HasWebOrigins

`func (o *ApplicationRepresentation) HasWebOrigins() bool`

HasWebOrigins returns a boolean if a field has been set.

### GetNotBefore

`func (o *ApplicationRepresentation) GetNotBefore() int32`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ApplicationRepresentation) GetNotBeforeOk() (*int32, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ApplicationRepresentation) SetNotBefore(v int32)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *ApplicationRepresentation) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetBearerOnly

`func (o *ApplicationRepresentation) GetBearerOnly() bool`

GetBearerOnly returns the BearerOnly field if non-nil, zero value otherwise.

### GetBearerOnlyOk

`func (o *ApplicationRepresentation) GetBearerOnlyOk() (*bool, bool)`

GetBearerOnlyOk returns a tuple with the BearerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerOnly

`func (o *ApplicationRepresentation) SetBearerOnly(v bool)`

SetBearerOnly sets BearerOnly field to given value.

### HasBearerOnly

`func (o *ApplicationRepresentation) HasBearerOnly() bool`

HasBearerOnly returns a boolean if a field has been set.

### GetConsentRequired

`func (o *ApplicationRepresentation) GetConsentRequired() bool`

GetConsentRequired returns the ConsentRequired field if non-nil, zero value otherwise.

### GetConsentRequiredOk

`func (o *ApplicationRepresentation) GetConsentRequiredOk() (*bool, bool)`

GetConsentRequiredOk returns a tuple with the ConsentRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentRequired

`func (o *ApplicationRepresentation) SetConsentRequired(v bool)`

SetConsentRequired sets ConsentRequired field to given value.

### HasConsentRequired

`func (o *ApplicationRepresentation) HasConsentRequired() bool`

HasConsentRequired returns a boolean if a field has been set.

### GetStandardFlowEnabled

`func (o *ApplicationRepresentation) GetStandardFlowEnabled() bool`

GetStandardFlowEnabled returns the StandardFlowEnabled field if non-nil, zero value otherwise.

### GetStandardFlowEnabledOk

`func (o *ApplicationRepresentation) GetStandardFlowEnabledOk() (*bool, bool)`

GetStandardFlowEnabledOk returns a tuple with the StandardFlowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardFlowEnabled

`func (o *ApplicationRepresentation) SetStandardFlowEnabled(v bool)`

SetStandardFlowEnabled sets StandardFlowEnabled field to given value.

### HasStandardFlowEnabled

`func (o *ApplicationRepresentation) HasStandardFlowEnabled() bool`

HasStandardFlowEnabled returns a boolean if a field has been set.

### GetImplicitFlowEnabled

`func (o *ApplicationRepresentation) GetImplicitFlowEnabled() bool`

GetImplicitFlowEnabled returns the ImplicitFlowEnabled field if non-nil, zero value otherwise.

### GetImplicitFlowEnabledOk

`func (o *ApplicationRepresentation) GetImplicitFlowEnabledOk() (*bool, bool)`

GetImplicitFlowEnabledOk returns a tuple with the ImplicitFlowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitFlowEnabled

`func (o *ApplicationRepresentation) SetImplicitFlowEnabled(v bool)`

SetImplicitFlowEnabled sets ImplicitFlowEnabled field to given value.

### HasImplicitFlowEnabled

`func (o *ApplicationRepresentation) HasImplicitFlowEnabled() bool`

HasImplicitFlowEnabled returns a boolean if a field has been set.

### GetDirectAccessGrantsEnabled

`func (o *ApplicationRepresentation) GetDirectAccessGrantsEnabled() bool`

GetDirectAccessGrantsEnabled returns the DirectAccessGrantsEnabled field if non-nil, zero value otherwise.

### GetDirectAccessGrantsEnabledOk

`func (o *ApplicationRepresentation) GetDirectAccessGrantsEnabledOk() (*bool, bool)`

GetDirectAccessGrantsEnabledOk returns a tuple with the DirectAccessGrantsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectAccessGrantsEnabled

`func (o *ApplicationRepresentation) SetDirectAccessGrantsEnabled(v bool)`

SetDirectAccessGrantsEnabled sets DirectAccessGrantsEnabled field to given value.

### HasDirectAccessGrantsEnabled

`func (o *ApplicationRepresentation) HasDirectAccessGrantsEnabled() bool`

HasDirectAccessGrantsEnabled returns a boolean if a field has been set.

### GetServiceAccountsEnabled

`func (o *ApplicationRepresentation) GetServiceAccountsEnabled() bool`

GetServiceAccountsEnabled returns the ServiceAccountsEnabled field if non-nil, zero value otherwise.

### GetServiceAccountsEnabledOk

`func (o *ApplicationRepresentation) GetServiceAccountsEnabledOk() (*bool, bool)`

GetServiceAccountsEnabledOk returns a tuple with the ServiceAccountsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountsEnabled

`func (o *ApplicationRepresentation) SetServiceAccountsEnabled(v bool)`

SetServiceAccountsEnabled sets ServiceAccountsEnabled field to given value.

### HasServiceAccountsEnabled

`func (o *ApplicationRepresentation) HasServiceAccountsEnabled() bool`

HasServiceAccountsEnabled returns a boolean if a field has been set.

### GetAuthorizationServicesEnabled

`func (o *ApplicationRepresentation) GetAuthorizationServicesEnabled() bool`

GetAuthorizationServicesEnabled returns the AuthorizationServicesEnabled field if non-nil, zero value otherwise.

### GetAuthorizationServicesEnabledOk

`func (o *ApplicationRepresentation) GetAuthorizationServicesEnabledOk() (*bool, bool)`

GetAuthorizationServicesEnabledOk returns a tuple with the AuthorizationServicesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServicesEnabled

`func (o *ApplicationRepresentation) SetAuthorizationServicesEnabled(v bool)`

SetAuthorizationServicesEnabled sets AuthorizationServicesEnabled field to given value.

### HasAuthorizationServicesEnabled

`func (o *ApplicationRepresentation) HasAuthorizationServicesEnabled() bool`

HasAuthorizationServicesEnabled returns a boolean if a field has been set.

### GetDirectGrantsOnly

`func (o *ApplicationRepresentation) GetDirectGrantsOnly() bool`

GetDirectGrantsOnly returns the DirectGrantsOnly field if non-nil, zero value otherwise.

### GetDirectGrantsOnlyOk

`func (o *ApplicationRepresentation) GetDirectGrantsOnlyOk() (*bool, bool)`

GetDirectGrantsOnlyOk returns a tuple with the DirectGrantsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectGrantsOnly

`func (o *ApplicationRepresentation) SetDirectGrantsOnly(v bool)`

SetDirectGrantsOnly sets DirectGrantsOnly field to given value.

### HasDirectGrantsOnly

`func (o *ApplicationRepresentation) HasDirectGrantsOnly() bool`

HasDirectGrantsOnly returns a boolean if a field has been set.

### GetPublicClient

`func (o *ApplicationRepresentation) GetPublicClient() bool`

GetPublicClient returns the PublicClient field if non-nil, zero value otherwise.

### GetPublicClientOk

`func (o *ApplicationRepresentation) GetPublicClientOk() (*bool, bool)`

GetPublicClientOk returns a tuple with the PublicClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicClient

`func (o *ApplicationRepresentation) SetPublicClient(v bool)`

SetPublicClient sets PublicClient field to given value.

### HasPublicClient

`func (o *ApplicationRepresentation) HasPublicClient() bool`

HasPublicClient returns a boolean if a field has been set.

### GetFrontchannelLogout

`func (o *ApplicationRepresentation) GetFrontchannelLogout() bool`

GetFrontchannelLogout returns the FrontchannelLogout field if non-nil, zero value otherwise.

### GetFrontchannelLogoutOk

`func (o *ApplicationRepresentation) GetFrontchannelLogoutOk() (*bool, bool)`

GetFrontchannelLogoutOk returns a tuple with the FrontchannelLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontchannelLogout

`func (o *ApplicationRepresentation) SetFrontchannelLogout(v bool)`

SetFrontchannelLogout sets FrontchannelLogout field to given value.

### HasFrontchannelLogout

`func (o *ApplicationRepresentation) HasFrontchannelLogout() bool`

HasFrontchannelLogout returns a boolean if a field has been set.

### GetProtocol

`func (o *ApplicationRepresentation) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApplicationRepresentation) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApplicationRepresentation) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApplicationRepresentation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAttributes

`func (o *ApplicationRepresentation) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ApplicationRepresentation) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ApplicationRepresentation) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ApplicationRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAuthenticationFlowBindingOverrides

`func (o *ApplicationRepresentation) GetAuthenticationFlowBindingOverrides() map[string]string`

GetAuthenticationFlowBindingOverrides returns the AuthenticationFlowBindingOverrides field if non-nil, zero value otherwise.

### GetAuthenticationFlowBindingOverridesOk

`func (o *ApplicationRepresentation) GetAuthenticationFlowBindingOverridesOk() (*map[string]string, bool)`

GetAuthenticationFlowBindingOverridesOk returns a tuple with the AuthenticationFlowBindingOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlowBindingOverrides

`func (o *ApplicationRepresentation) SetAuthenticationFlowBindingOverrides(v map[string]string)`

SetAuthenticationFlowBindingOverrides sets AuthenticationFlowBindingOverrides field to given value.

### HasAuthenticationFlowBindingOverrides

`func (o *ApplicationRepresentation) HasAuthenticationFlowBindingOverrides() bool`

HasAuthenticationFlowBindingOverrides returns a boolean if a field has been set.

### GetFullScopeAllowed

`func (o *ApplicationRepresentation) GetFullScopeAllowed() bool`

GetFullScopeAllowed returns the FullScopeAllowed field if non-nil, zero value otherwise.

### GetFullScopeAllowedOk

`func (o *ApplicationRepresentation) GetFullScopeAllowedOk() (*bool, bool)`

GetFullScopeAllowedOk returns a tuple with the FullScopeAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullScopeAllowed

`func (o *ApplicationRepresentation) SetFullScopeAllowed(v bool)`

SetFullScopeAllowed sets FullScopeAllowed field to given value.

### HasFullScopeAllowed

`func (o *ApplicationRepresentation) HasFullScopeAllowed() bool`

HasFullScopeAllowed returns a boolean if a field has been set.

### GetNodeReRegistrationTimeout

`func (o *ApplicationRepresentation) GetNodeReRegistrationTimeout() int32`

GetNodeReRegistrationTimeout returns the NodeReRegistrationTimeout field if non-nil, zero value otherwise.

### GetNodeReRegistrationTimeoutOk

`func (o *ApplicationRepresentation) GetNodeReRegistrationTimeoutOk() (*int32, bool)`

GetNodeReRegistrationTimeoutOk returns a tuple with the NodeReRegistrationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeReRegistrationTimeout

`func (o *ApplicationRepresentation) SetNodeReRegistrationTimeout(v int32)`

SetNodeReRegistrationTimeout sets NodeReRegistrationTimeout field to given value.

### HasNodeReRegistrationTimeout

`func (o *ApplicationRepresentation) HasNodeReRegistrationTimeout() bool`

HasNodeReRegistrationTimeout returns a boolean if a field has been set.

### GetRegisteredNodes

`func (o *ApplicationRepresentation) GetRegisteredNodes() map[string]int32`

GetRegisteredNodes returns the RegisteredNodes field if non-nil, zero value otherwise.

### GetRegisteredNodesOk

`func (o *ApplicationRepresentation) GetRegisteredNodesOk() (*map[string]int32, bool)`

GetRegisteredNodesOk returns a tuple with the RegisteredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredNodes

`func (o *ApplicationRepresentation) SetRegisteredNodes(v map[string]int32)`

SetRegisteredNodes sets RegisteredNodes field to given value.

### HasRegisteredNodes

`func (o *ApplicationRepresentation) HasRegisteredNodes() bool`

HasRegisteredNodes returns a boolean if a field has been set.

### GetProtocolMappers

`func (o *ApplicationRepresentation) GetProtocolMappers() []ProtocolMapperRepresentation`

GetProtocolMappers returns the ProtocolMappers field if non-nil, zero value otherwise.

### GetProtocolMappersOk

`func (o *ApplicationRepresentation) GetProtocolMappersOk() (*[]ProtocolMapperRepresentation, bool)`

GetProtocolMappersOk returns a tuple with the ProtocolMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMappers

`func (o *ApplicationRepresentation) SetProtocolMappers(v []ProtocolMapperRepresentation)`

SetProtocolMappers sets ProtocolMappers field to given value.

### HasProtocolMappers

`func (o *ApplicationRepresentation) HasProtocolMappers() bool`

HasProtocolMappers returns a boolean if a field has been set.

### GetClientTemplate

`func (o *ApplicationRepresentation) GetClientTemplate() string`

GetClientTemplate returns the ClientTemplate field if non-nil, zero value otherwise.

### GetClientTemplateOk

`func (o *ApplicationRepresentation) GetClientTemplateOk() (*string, bool)`

GetClientTemplateOk returns a tuple with the ClientTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTemplate

`func (o *ApplicationRepresentation) SetClientTemplate(v string)`

SetClientTemplate sets ClientTemplate field to given value.

### HasClientTemplate

`func (o *ApplicationRepresentation) HasClientTemplate() bool`

HasClientTemplate returns a boolean if a field has been set.

### GetUseTemplateConfig

`func (o *ApplicationRepresentation) GetUseTemplateConfig() bool`

GetUseTemplateConfig returns the UseTemplateConfig field if non-nil, zero value otherwise.

### GetUseTemplateConfigOk

`func (o *ApplicationRepresentation) GetUseTemplateConfigOk() (*bool, bool)`

GetUseTemplateConfigOk returns a tuple with the UseTemplateConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTemplateConfig

`func (o *ApplicationRepresentation) SetUseTemplateConfig(v bool)`

SetUseTemplateConfig sets UseTemplateConfig field to given value.

### HasUseTemplateConfig

`func (o *ApplicationRepresentation) HasUseTemplateConfig() bool`

HasUseTemplateConfig returns a boolean if a field has been set.

### GetUseTemplateScope

`func (o *ApplicationRepresentation) GetUseTemplateScope() bool`

GetUseTemplateScope returns the UseTemplateScope field if non-nil, zero value otherwise.

### GetUseTemplateScopeOk

`func (o *ApplicationRepresentation) GetUseTemplateScopeOk() (*bool, bool)`

GetUseTemplateScopeOk returns a tuple with the UseTemplateScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTemplateScope

`func (o *ApplicationRepresentation) SetUseTemplateScope(v bool)`

SetUseTemplateScope sets UseTemplateScope field to given value.

### HasUseTemplateScope

`func (o *ApplicationRepresentation) HasUseTemplateScope() bool`

HasUseTemplateScope returns a boolean if a field has been set.

### GetUseTemplateMappers

`func (o *ApplicationRepresentation) GetUseTemplateMappers() bool`

GetUseTemplateMappers returns the UseTemplateMappers field if non-nil, zero value otherwise.

### GetUseTemplateMappersOk

`func (o *ApplicationRepresentation) GetUseTemplateMappersOk() (*bool, bool)`

GetUseTemplateMappersOk returns a tuple with the UseTemplateMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTemplateMappers

`func (o *ApplicationRepresentation) SetUseTemplateMappers(v bool)`

SetUseTemplateMappers sets UseTemplateMappers field to given value.

### HasUseTemplateMappers

`func (o *ApplicationRepresentation) HasUseTemplateMappers() bool`

HasUseTemplateMappers returns a boolean if a field has been set.

### GetDefaultClientScopes

`func (o *ApplicationRepresentation) GetDefaultClientScopes() []string`

GetDefaultClientScopes returns the DefaultClientScopes field if non-nil, zero value otherwise.

### GetDefaultClientScopesOk

`func (o *ApplicationRepresentation) GetDefaultClientScopesOk() (*[]string, bool)`

GetDefaultClientScopesOk returns a tuple with the DefaultClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientScopes

`func (o *ApplicationRepresentation) SetDefaultClientScopes(v []string)`

SetDefaultClientScopes sets DefaultClientScopes field to given value.

### HasDefaultClientScopes

`func (o *ApplicationRepresentation) HasDefaultClientScopes() bool`

HasDefaultClientScopes returns a boolean if a field has been set.

### GetOptionalClientScopes

`func (o *ApplicationRepresentation) GetOptionalClientScopes() []string`

GetOptionalClientScopes returns the OptionalClientScopes field if non-nil, zero value otherwise.

### GetOptionalClientScopesOk

`func (o *ApplicationRepresentation) GetOptionalClientScopesOk() (*[]string, bool)`

GetOptionalClientScopesOk returns a tuple with the OptionalClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalClientScopes

`func (o *ApplicationRepresentation) SetOptionalClientScopes(v []string)`

SetOptionalClientScopes sets OptionalClientScopes field to given value.

### HasOptionalClientScopes

`func (o *ApplicationRepresentation) HasOptionalClientScopes() bool`

HasOptionalClientScopes returns a boolean if a field has been set.

### GetAuthorizationSettings

`func (o *ApplicationRepresentation) GetAuthorizationSettings() ResourceServerRepresentation`

GetAuthorizationSettings returns the AuthorizationSettings field if non-nil, zero value otherwise.

### GetAuthorizationSettingsOk

`func (o *ApplicationRepresentation) GetAuthorizationSettingsOk() (*ResourceServerRepresentation, bool)`

GetAuthorizationSettingsOk returns a tuple with the AuthorizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationSettings

`func (o *ApplicationRepresentation) SetAuthorizationSettings(v ResourceServerRepresentation)`

SetAuthorizationSettings sets AuthorizationSettings field to given value.

### HasAuthorizationSettings

`func (o *ApplicationRepresentation) HasAuthorizationSettings() bool`

HasAuthorizationSettings returns a boolean if a field has been set.

### GetAccess

`func (o *ApplicationRepresentation) GetAccess() map[string]bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ApplicationRepresentation) GetAccessOk() (*map[string]bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ApplicationRepresentation) SetAccess(v map[string]bool)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ApplicationRepresentation) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetOrigin

`func (o *ApplicationRepresentation) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ApplicationRepresentation) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ApplicationRepresentation) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ApplicationRepresentation) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClaims

`func (o *ApplicationRepresentation) GetClaims() ClaimRepresentation`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *ApplicationRepresentation) GetClaimsOk() (*ClaimRepresentation, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *ApplicationRepresentation) SetClaims(v ClaimRepresentation)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *ApplicationRepresentation) HasClaims() bool`

HasClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


