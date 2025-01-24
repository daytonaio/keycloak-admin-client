# RealmRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DisplayNameHtml** | Pointer to **string** |  | [optional] 
**NotBefore** | Pointer to **int32** |  | [optional] 
**DefaultSignatureAlgorithm** | Pointer to **string** |  | [optional] 
**RevokeRefreshToken** | Pointer to **bool** |  | [optional] 
**RefreshTokenMaxReuse** | Pointer to **int32** |  | [optional] 
**AccessTokenLifespan** | Pointer to **int32** |  | [optional] 
**AccessTokenLifespanForImplicitFlow** | Pointer to **int32** |  | [optional] 
**SsoSessionIdleTimeout** | Pointer to **int32** |  | [optional] 
**SsoSessionMaxLifespan** | Pointer to **int32** |  | [optional] 
**SsoSessionIdleTimeoutRememberMe** | Pointer to **int32** |  | [optional] 
**SsoSessionMaxLifespanRememberMe** | Pointer to **int32** |  | [optional] 
**OfflineSessionIdleTimeout** | Pointer to **int32** |  | [optional] 
**OfflineSessionMaxLifespanEnabled** | Pointer to **bool** |  | [optional] 
**OfflineSessionMaxLifespan** | Pointer to **int32** |  | [optional] 
**ClientSessionIdleTimeout** | Pointer to **int32** |  | [optional] 
**ClientSessionMaxLifespan** | Pointer to **int32** |  | [optional] 
**ClientOfflineSessionIdleTimeout** | Pointer to **int32** |  | [optional] 
**ClientOfflineSessionMaxLifespan** | Pointer to **int32** |  | [optional] 
**AccessCodeLifespan** | Pointer to **int32** |  | [optional] 
**AccessCodeLifespanUserAction** | Pointer to **int32** |  | [optional] 
**AccessCodeLifespanLogin** | Pointer to **int32** |  | [optional] 
**ActionTokenGeneratedByAdminLifespan** | Pointer to **int32** |  | [optional] 
**ActionTokenGeneratedByUserLifespan** | Pointer to **int32** |  | [optional] 
**Oauth2DeviceCodeLifespan** | Pointer to **int32** |  | [optional] 
**Oauth2DevicePollingInterval** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SslRequired** | Pointer to **string** |  | [optional] 
**PasswordCredentialGrantAllowed** | Pointer to **bool** |  | [optional] 
**RegistrationAllowed** | Pointer to **bool** |  | [optional] 
**RegistrationEmailAsUsername** | Pointer to **bool** |  | [optional] 
**RememberMe** | Pointer to **bool** |  | [optional] 
**VerifyEmail** | Pointer to **bool** |  | [optional] 
**LoginWithEmailAllowed** | Pointer to **bool** |  | [optional] 
**DuplicateEmailsAllowed** | Pointer to **bool** |  | [optional] 
**ResetPasswordAllowed** | Pointer to **bool** |  | [optional] 
**EditUsernameAllowed** | Pointer to **bool** |  | [optional] 
**UserCacheEnabled** | Pointer to **bool** |  | [optional] 
**RealmCacheEnabled** | Pointer to **bool** |  | [optional] 
**BruteForceProtected** | Pointer to **bool** |  | [optional] 
**PermanentLockout** | Pointer to **bool** |  | [optional] 
**MaxTemporaryLockouts** | Pointer to **int32** |  | [optional] 
**BruteForceStrategy** | Pointer to [**BruteForceStrategy**](BruteForceStrategy.md) |  | [optional] 
**MaxFailureWaitSeconds** | Pointer to **int32** |  | [optional] 
**MinimumQuickLoginWaitSeconds** | Pointer to **int32** |  | [optional] 
**WaitIncrementSeconds** | Pointer to **int32** |  | [optional] 
**QuickLoginCheckMilliSeconds** | Pointer to **int64** |  | [optional] 
**MaxDeltaTimeSeconds** | Pointer to **int32** |  | [optional] 
**FailureFactor** | Pointer to **int32** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CodeSecret** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to [**RolesRepresentation**](RolesRepresentation.md) |  | [optional] 
**Groups** | Pointer to [**[]GroupRepresentation**](GroupRepresentation.md) |  | [optional] 
**DefaultRoles** | Pointer to **[]string** |  | [optional] 
**DefaultRole** | Pointer to [**RoleRepresentation**](RoleRepresentation.md) |  | [optional] 
**AdminPermissionsClient** | Pointer to [**ClientRepresentation**](ClientRepresentation.md) |  | [optional] 
**DefaultGroups** | Pointer to **[]string** |  | [optional] 
**RequiredCredentials** | Pointer to **[]string** |  | [optional] 
**PasswordPolicy** | Pointer to **string** |  | [optional] 
**OtpPolicyType** | Pointer to **string** |  | [optional] 
**OtpPolicyAlgorithm** | Pointer to **string** |  | [optional] 
**OtpPolicyInitialCounter** | Pointer to **int32** |  | [optional] 
**OtpPolicyDigits** | Pointer to **int32** |  | [optional] 
**OtpPolicyLookAheadWindow** | Pointer to **int32** |  | [optional] 
**OtpPolicyPeriod** | Pointer to **int32** |  | [optional] 
**OtpPolicyCodeReusable** | Pointer to **bool** |  | [optional] 
**OtpSupportedApplications** | Pointer to **[]string** |  | [optional] 
**LocalizationTexts** | Pointer to **map[string]map[string]string** |  | [optional] 
**WebAuthnPolicyRpEntityName** | Pointer to **string** |  | [optional] 
**WebAuthnPolicySignatureAlgorithms** | Pointer to **[]string** |  | [optional] 
**WebAuthnPolicyRpId** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyAttestationConveyancePreference** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyAuthenticatorAttachment** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyRequireResidentKey** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyUserVerificationRequirement** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyCreateTimeout** | Pointer to **int32** |  | [optional] 
**WebAuthnPolicyAvoidSameAuthenticatorRegister** | Pointer to **bool** |  | [optional] 
**WebAuthnPolicyAcceptableAaguids** | Pointer to **[]string** |  | [optional] 
**WebAuthnPolicyExtraOrigins** | Pointer to **[]string** |  | [optional] 
**WebAuthnPolicyPasswordlessRpEntityName** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyPasswordlessSignatureAlgorithms** | Pointer to **[]string** |  | [optional] 
**WebAuthnPolicyPasswordlessRpId** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyPasswordlessAttestationConveyancePreference** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyPasswordlessAuthenticatorAttachment** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyPasswordlessRequireResidentKey** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyPasswordlessUserVerificationRequirement** | Pointer to **string** |  | [optional] 
**WebAuthnPolicyPasswordlessCreateTimeout** | Pointer to **int32** |  | [optional] 
**WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister** | Pointer to **bool** |  | [optional] 
**WebAuthnPolicyPasswordlessAcceptableAaguids** | Pointer to **[]string** |  | [optional] 
**WebAuthnPolicyPasswordlessExtraOrigins** | Pointer to **[]string** |  | [optional] 
**ClientProfiles** | Pointer to [**ClientProfilesRepresentation**](ClientProfilesRepresentation.md) |  | [optional] 
**ClientPolicies** | Pointer to [**ClientPoliciesRepresentation**](ClientPoliciesRepresentation.md) |  | [optional] 
**Users** | Pointer to [**[]UserRepresentation**](UserRepresentation.md) |  | [optional] 
**FederatedUsers** | Pointer to [**[]UserRepresentation**](UserRepresentation.md) |  | [optional] 
**ScopeMappings** | Pointer to [**[]ScopeMappingRepresentation**](ScopeMappingRepresentation.md) |  | [optional] 
**ClientScopeMappings** | Pointer to [**map[string][]ScopeMappingRepresentation**](array.md) |  | [optional] 
**Clients** | Pointer to [**[]ClientRepresentation**](ClientRepresentation.md) |  | [optional] 
**ClientScopes** | Pointer to [**[]ClientScopeRepresentation**](ClientScopeRepresentation.md) |  | [optional] 
**DefaultDefaultClientScopes** | Pointer to **[]string** |  | [optional] 
**DefaultOptionalClientScopes** | Pointer to **[]string** |  | [optional] 
**BrowserSecurityHeaders** | Pointer to **map[string]string** |  | [optional] 
**SmtpServer** | Pointer to **map[string]string** |  | [optional] 
**UserFederationProviders** | Pointer to [**[]UserFederationProviderRepresentation**](UserFederationProviderRepresentation.md) |  | [optional] 
**UserFederationMappers** | Pointer to [**[]UserFederationMapperRepresentation**](UserFederationMapperRepresentation.md) |  | [optional] 
**LoginTheme** | Pointer to **string** |  | [optional] 
**AccountTheme** | Pointer to **string** |  | [optional] 
**AdminTheme** | Pointer to **string** |  | [optional] 
**EmailTheme** | Pointer to **string** |  | [optional] 
**EventsEnabled** | Pointer to **bool** |  | [optional] 
**EventsExpiration** | Pointer to **int64** |  | [optional] 
**EventsListeners** | Pointer to **[]string** |  | [optional] 
**EnabledEventTypes** | Pointer to **[]string** |  | [optional] 
**AdminEventsEnabled** | Pointer to **bool** |  | [optional] 
**AdminEventsDetailsEnabled** | Pointer to **bool** |  | [optional] 
**IdentityProviders** | Pointer to [**[]IdentityProviderRepresentation**](IdentityProviderRepresentation.md) |  | [optional] 
**IdentityProviderMappers** | Pointer to [**[]IdentityProviderMapperRepresentation**](IdentityProviderMapperRepresentation.md) |  | [optional] 
**ProtocolMappers** | Pointer to [**[]ProtocolMapperRepresentation**](ProtocolMapperRepresentation.md) |  | [optional] 
**Components** | Pointer to [**map[string][]ComponentExportRepresentation**](array.md) |  | [optional] 
**InternationalizationEnabled** | Pointer to **bool** |  | [optional] 
**SupportedLocales** | Pointer to **[]string** |  | [optional] 
**DefaultLocale** | Pointer to **string** |  | [optional] 
**AuthenticationFlows** | Pointer to [**[]AuthenticationFlowRepresentation**](AuthenticationFlowRepresentation.md) |  | [optional] 
**AuthenticatorConfig** | Pointer to [**[]AuthenticatorConfigRepresentation**](AuthenticatorConfigRepresentation.md) |  | [optional] 
**RequiredActions** | Pointer to [**[]RequiredActionProviderRepresentation**](RequiredActionProviderRepresentation.md) |  | [optional] 
**BrowserFlow** | Pointer to **string** |  | [optional] 
**RegistrationFlow** | Pointer to **string** |  | [optional] 
**DirectGrantFlow** | Pointer to **string** |  | [optional] 
**ResetCredentialsFlow** | Pointer to **string** |  | [optional] 
**ClientAuthenticationFlow** | Pointer to **string** |  | [optional] 
**DockerAuthenticationFlow** | Pointer to **string** |  | [optional] 
**FirstBrokerLoginFlow** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**KeycloakVersion** | Pointer to **string** |  | [optional] 
**UserManagedAccessAllowed** | Pointer to **bool** |  | [optional] 
**OrganizationsEnabled** | Pointer to **bool** |  | [optional] 
**Organizations** | Pointer to [**[]OrganizationRepresentation**](OrganizationRepresentation.md) |  | [optional] 
**VerifiableCredentialsEnabled** | Pointer to **bool** |  | [optional] 
**AdminPermissionsEnabled** | Pointer to **bool** |  | [optional] 
**Social** | Pointer to **bool** |  | [optional] 
**UpdateProfileOnInitialSocialLogin** | Pointer to **bool** |  | [optional] 
**SocialProviders** | Pointer to **map[string]string** |  | [optional] 
**ApplicationScopeMappings** | Pointer to [**map[string][]ScopeMappingRepresentation**](array.md) |  | [optional] 
**Applications** | Pointer to [**[]ApplicationRepresentation**](ApplicationRepresentation.md) |  | [optional] 
**OauthClients** | Pointer to [**[]OAuthClientRepresentation**](OAuthClientRepresentation.md) |  | [optional] 
**ClientTemplates** | Pointer to [**[]ClientTemplateRepresentation**](ClientTemplateRepresentation.md) |  | [optional] 
**OAuth2DeviceCodeLifespan** | Pointer to **int32** |  | [optional] 
**OAuth2DevicePollingInterval** | Pointer to **int32** |  | [optional] 

## Methods

### NewRealmRepresentation

`func NewRealmRepresentation() *RealmRepresentation`

NewRealmRepresentation instantiates a new RealmRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmRepresentationWithDefaults

`func NewRealmRepresentationWithDefaults() *RealmRepresentation`

NewRealmRepresentationWithDefaults instantiates a new RealmRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RealmRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RealmRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRealm

`func (o *RealmRepresentation) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *RealmRepresentation) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *RealmRepresentation) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *RealmRepresentation) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetDisplayName

`func (o *RealmRepresentation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RealmRepresentation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RealmRepresentation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RealmRepresentation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNameHtml

`func (o *RealmRepresentation) GetDisplayNameHtml() string`

GetDisplayNameHtml returns the DisplayNameHtml field if non-nil, zero value otherwise.

### GetDisplayNameHtmlOk

`func (o *RealmRepresentation) GetDisplayNameHtmlOk() (*string, bool)`

GetDisplayNameHtmlOk returns a tuple with the DisplayNameHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameHtml

`func (o *RealmRepresentation) SetDisplayNameHtml(v string)`

SetDisplayNameHtml sets DisplayNameHtml field to given value.

### HasDisplayNameHtml

`func (o *RealmRepresentation) HasDisplayNameHtml() bool`

HasDisplayNameHtml returns a boolean if a field has been set.

### GetNotBefore

`func (o *RealmRepresentation) GetNotBefore() int32`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *RealmRepresentation) GetNotBeforeOk() (*int32, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *RealmRepresentation) SetNotBefore(v int32)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *RealmRepresentation) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetDefaultSignatureAlgorithm

`func (o *RealmRepresentation) GetDefaultSignatureAlgorithm() string`

GetDefaultSignatureAlgorithm returns the DefaultSignatureAlgorithm field if non-nil, zero value otherwise.

### GetDefaultSignatureAlgorithmOk

`func (o *RealmRepresentation) GetDefaultSignatureAlgorithmOk() (*string, bool)`

GetDefaultSignatureAlgorithmOk returns a tuple with the DefaultSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSignatureAlgorithm

`func (o *RealmRepresentation) SetDefaultSignatureAlgorithm(v string)`

SetDefaultSignatureAlgorithm sets DefaultSignatureAlgorithm field to given value.

### HasDefaultSignatureAlgorithm

`func (o *RealmRepresentation) HasDefaultSignatureAlgorithm() bool`

HasDefaultSignatureAlgorithm returns a boolean if a field has been set.

### GetRevokeRefreshToken

`func (o *RealmRepresentation) GetRevokeRefreshToken() bool`

GetRevokeRefreshToken returns the RevokeRefreshToken field if non-nil, zero value otherwise.

### GetRevokeRefreshTokenOk

`func (o *RealmRepresentation) GetRevokeRefreshTokenOk() (*bool, bool)`

GetRevokeRefreshTokenOk returns a tuple with the RevokeRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeRefreshToken

`func (o *RealmRepresentation) SetRevokeRefreshToken(v bool)`

SetRevokeRefreshToken sets RevokeRefreshToken field to given value.

### HasRevokeRefreshToken

`func (o *RealmRepresentation) HasRevokeRefreshToken() bool`

HasRevokeRefreshToken returns a boolean if a field has been set.

### GetRefreshTokenMaxReuse

`func (o *RealmRepresentation) GetRefreshTokenMaxReuse() int32`

GetRefreshTokenMaxReuse returns the RefreshTokenMaxReuse field if non-nil, zero value otherwise.

### GetRefreshTokenMaxReuseOk

`func (o *RealmRepresentation) GetRefreshTokenMaxReuseOk() (*int32, bool)`

GetRefreshTokenMaxReuseOk returns a tuple with the RefreshTokenMaxReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenMaxReuse

`func (o *RealmRepresentation) SetRefreshTokenMaxReuse(v int32)`

SetRefreshTokenMaxReuse sets RefreshTokenMaxReuse field to given value.

### HasRefreshTokenMaxReuse

`func (o *RealmRepresentation) HasRefreshTokenMaxReuse() bool`

HasRefreshTokenMaxReuse returns a boolean if a field has been set.

### GetAccessTokenLifespan

`func (o *RealmRepresentation) GetAccessTokenLifespan() int32`

GetAccessTokenLifespan returns the AccessTokenLifespan field if non-nil, zero value otherwise.

### GetAccessTokenLifespanOk

`func (o *RealmRepresentation) GetAccessTokenLifespanOk() (*int32, bool)`

GetAccessTokenLifespanOk returns a tuple with the AccessTokenLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenLifespan

`func (o *RealmRepresentation) SetAccessTokenLifespan(v int32)`

SetAccessTokenLifespan sets AccessTokenLifespan field to given value.

### HasAccessTokenLifespan

`func (o *RealmRepresentation) HasAccessTokenLifespan() bool`

HasAccessTokenLifespan returns a boolean if a field has been set.

### GetAccessTokenLifespanForImplicitFlow

`func (o *RealmRepresentation) GetAccessTokenLifespanForImplicitFlow() int32`

GetAccessTokenLifespanForImplicitFlow returns the AccessTokenLifespanForImplicitFlow field if non-nil, zero value otherwise.

### GetAccessTokenLifespanForImplicitFlowOk

`func (o *RealmRepresentation) GetAccessTokenLifespanForImplicitFlowOk() (*int32, bool)`

GetAccessTokenLifespanForImplicitFlowOk returns a tuple with the AccessTokenLifespanForImplicitFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenLifespanForImplicitFlow

`func (o *RealmRepresentation) SetAccessTokenLifespanForImplicitFlow(v int32)`

SetAccessTokenLifespanForImplicitFlow sets AccessTokenLifespanForImplicitFlow field to given value.

### HasAccessTokenLifespanForImplicitFlow

`func (o *RealmRepresentation) HasAccessTokenLifespanForImplicitFlow() bool`

HasAccessTokenLifespanForImplicitFlow returns a boolean if a field has been set.

### GetSsoSessionIdleTimeout

`func (o *RealmRepresentation) GetSsoSessionIdleTimeout() int32`

GetSsoSessionIdleTimeout returns the SsoSessionIdleTimeout field if non-nil, zero value otherwise.

### GetSsoSessionIdleTimeoutOk

`func (o *RealmRepresentation) GetSsoSessionIdleTimeoutOk() (*int32, bool)`

GetSsoSessionIdleTimeoutOk returns a tuple with the SsoSessionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionIdleTimeout

`func (o *RealmRepresentation) SetSsoSessionIdleTimeout(v int32)`

SetSsoSessionIdleTimeout sets SsoSessionIdleTimeout field to given value.

### HasSsoSessionIdleTimeout

`func (o *RealmRepresentation) HasSsoSessionIdleTimeout() bool`

HasSsoSessionIdleTimeout returns a boolean if a field has been set.

### GetSsoSessionMaxLifespan

`func (o *RealmRepresentation) GetSsoSessionMaxLifespan() int32`

GetSsoSessionMaxLifespan returns the SsoSessionMaxLifespan field if non-nil, zero value otherwise.

### GetSsoSessionMaxLifespanOk

`func (o *RealmRepresentation) GetSsoSessionMaxLifespanOk() (*int32, bool)`

GetSsoSessionMaxLifespanOk returns a tuple with the SsoSessionMaxLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionMaxLifespan

`func (o *RealmRepresentation) SetSsoSessionMaxLifespan(v int32)`

SetSsoSessionMaxLifespan sets SsoSessionMaxLifespan field to given value.

### HasSsoSessionMaxLifespan

`func (o *RealmRepresentation) HasSsoSessionMaxLifespan() bool`

HasSsoSessionMaxLifespan returns a boolean if a field has been set.

### GetSsoSessionIdleTimeoutRememberMe

`func (o *RealmRepresentation) GetSsoSessionIdleTimeoutRememberMe() int32`

GetSsoSessionIdleTimeoutRememberMe returns the SsoSessionIdleTimeoutRememberMe field if non-nil, zero value otherwise.

### GetSsoSessionIdleTimeoutRememberMeOk

`func (o *RealmRepresentation) GetSsoSessionIdleTimeoutRememberMeOk() (*int32, bool)`

GetSsoSessionIdleTimeoutRememberMeOk returns a tuple with the SsoSessionIdleTimeoutRememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionIdleTimeoutRememberMe

`func (o *RealmRepresentation) SetSsoSessionIdleTimeoutRememberMe(v int32)`

SetSsoSessionIdleTimeoutRememberMe sets SsoSessionIdleTimeoutRememberMe field to given value.

### HasSsoSessionIdleTimeoutRememberMe

`func (o *RealmRepresentation) HasSsoSessionIdleTimeoutRememberMe() bool`

HasSsoSessionIdleTimeoutRememberMe returns a boolean if a field has been set.

### GetSsoSessionMaxLifespanRememberMe

`func (o *RealmRepresentation) GetSsoSessionMaxLifespanRememberMe() int32`

GetSsoSessionMaxLifespanRememberMe returns the SsoSessionMaxLifespanRememberMe field if non-nil, zero value otherwise.

### GetSsoSessionMaxLifespanRememberMeOk

`func (o *RealmRepresentation) GetSsoSessionMaxLifespanRememberMeOk() (*int32, bool)`

GetSsoSessionMaxLifespanRememberMeOk returns a tuple with the SsoSessionMaxLifespanRememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionMaxLifespanRememberMe

`func (o *RealmRepresentation) SetSsoSessionMaxLifespanRememberMe(v int32)`

SetSsoSessionMaxLifespanRememberMe sets SsoSessionMaxLifespanRememberMe field to given value.

### HasSsoSessionMaxLifespanRememberMe

`func (o *RealmRepresentation) HasSsoSessionMaxLifespanRememberMe() bool`

HasSsoSessionMaxLifespanRememberMe returns a boolean if a field has been set.

### GetOfflineSessionIdleTimeout

`func (o *RealmRepresentation) GetOfflineSessionIdleTimeout() int32`

GetOfflineSessionIdleTimeout returns the OfflineSessionIdleTimeout field if non-nil, zero value otherwise.

### GetOfflineSessionIdleTimeoutOk

`func (o *RealmRepresentation) GetOfflineSessionIdleTimeoutOk() (*int32, bool)`

GetOfflineSessionIdleTimeoutOk returns a tuple with the OfflineSessionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineSessionIdleTimeout

`func (o *RealmRepresentation) SetOfflineSessionIdleTimeout(v int32)`

SetOfflineSessionIdleTimeout sets OfflineSessionIdleTimeout field to given value.

### HasOfflineSessionIdleTimeout

`func (o *RealmRepresentation) HasOfflineSessionIdleTimeout() bool`

HasOfflineSessionIdleTimeout returns a boolean if a field has been set.

### GetOfflineSessionMaxLifespanEnabled

`func (o *RealmRepresentation) GetOfflineSessionMaxLifespanEnabled() bool`

GetOfflineSessionMaxLifespanEnabled returns the OfflineSessionMaxLifespanEnabled field if non-nil, zero value otherwise.

### GetOfflineSessionMaxLifespanEnabledOk

`func (o *RealmRepresentation) GetOfflineSessionMaxLifespanEnabledOk() (*bool, bool)`

GetOfflineSessionMaxLifespanEnabledOk returns a tuple with the OfflineSessionMaxLifespanEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineSessionMaxLifespanEnabled

`func (o *RealmRepresentation) SetOfflineSessionMaxLifespanEnabled(v bool)`

SetOfflineSessionMaxLifespanEnabled sets OfflineSessionMaxLifespanEnabled field to given value.

### HasOfflineSessionMaxLifespanEnabled

`func (o *RealmRepresentation) HasOfflineSessionMaxLifespanEnabled() bool`

HasOfflineSessionMaxLifespanEnabled returns a boolean if a field has been set.

### GetOfflineSessionMaxLifespan

`func (o *RealmRepresentation) GetOfflineSessionMaxLifespan() int32`

GetOfflineSessionMaxLifespan returns the OfflineSessionMaxLifespan field if non-nil, zero value otherwise.

### GetOfflineSessionMaxLifespanOk

`func (o *RealmRepresentation) GetOfflineSessionMaxLifespanOk() (*int32, bool)`

GetOfflineSessionMaxLifespanOk returns a tuple with the OfflineSessionMaxLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineSessionMaxLifespan

`func (o *RealmRepresentation) SetOfflineSessionMaxLifespan(v int32)`

SetOfflineSessionMaxLifespan sets OfflineSessionMaxLifespan field to given value.

### HasOfflineSessionMaxLifespan

`func (o *RealmRepresentation) HasOfflineSessionMaxLifespan() bool`

HasOfflineSessionMaxLifespan returns a boolean if a field has been set.

### GetClientSessionIdleTimeout

`func (o *RealmRepresentation) GetClientSessionIdleTimeout() int32`

GetClientSessionIdleTimeout returns the ClientSessionIdleTimeout field if non-nil, zero value otherwise.

### GetClientSessionIdleTimeoutOk

`func (o *RealmRepresentation) GetClientSessionIdleTimeoutOk() (*int32, bool)`

GetClientSessionIdleTimeoutOk returns a tuple with the ClientSessionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSessionIdleTimeout

`func (o *RealmRepresentation) SetClientSessionIdleTimeout(v int32)`

SetClientSessionIdleTimeout sets ClientSessionIdleTimeout field to given value.

### HasClientSessionIdleTimeout

`func (o *RealmRepresentation) HasClientSessionIdleTimeout() bool`

HasClientSessionIdleTimeout returns a boolean if a field has been set.

### GetClientSessionMaxLifespan

`func (o *RealmRepresentation) GetClientSessionMaxLifespan() int32`

GetClientSessionMaxLifespan returns the ClientSessionMaxLifespan field if non-nil, zero value otherwise.

### GetClientSessionMaxLifespanOk

`func (o *RealmRepresentation) GetClientSessionMaxLifespanOk() (*int32, bool)`

GetClientSessionMaxLifespanOk returns a tuple with the ClientSessionMaxLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSessionMaxLifespan

`func (o *RealmRepresentation) SetClientSessionMaxLifespan(v int32)`

SetClientSessionMaxLifespan sets ClientSessionMaxLifespan field to given value.

### HasClientSessionMaxLifespan

`func (o *RealmRepresentation) HasClientSessionMaxLifespan() bool`

HasClientSessionMaxLifespan returns a boolean if a field has been set.

### GetClientOfflineSessionIdleTimeout

`func (o *RealmRepresentation) GetClientOfflineSessionIdleTimeout() int32`

GetClientOfflineSessionIdleTimeout returns the ClientOfflineSessionIdleTimeout field if non-nil, zero value otherwise.

### GetClientOfflineSessionIdleTimeoutOk

`func (o *RealmRepresentation) GetClientOfflineSessionIdleTimeoutOk() (*int32, bool)`

GetClientOfflineSessionIdleTimeoutOk returns a tuple with the ClientOfflineSessionIdleTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOfflineSessionIdleTimeout

`func (o *RealmRepresentation) SetClientOfflineSessionIdleTimeout(v int32)`

SetClientOfflineSessionIdleTimeout sets ClientOfflineSessionIdleTimeout field to given value.

### HasClientOfflineSessionIdleTimeout

`func (o *RealmRepresentation) HasClientOfflineSessionIdleTimeout() bool`

HasClientOfflineSessionIdleTimeout returns a boolean if a field has been set.

### GetClientOfflineSessionMaxLifespan

`func (o *RealmRepresentation) GetClientOfflineSessionMaxLifespan() int32`

GetClientOfflineSessionMaxLifespan returns the ClientOfflineSessionMaxLifespan field if non-nil, zero value otherwise.

### GetClientOfflineSessionMaxLifespanOk

`func (o *RealmRepresentation) GetClientOfflineSessionMaxLifespanOk() (*int32, bool)`

GetClientOfflineSessionMaxLifespanOk returns a tuple with the ClientOfflineSessionMaxLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOfflineSessionMaxLifespan

`func (o *RealmRepresentation) SetClientOfflineSessionMaxLifespan(v int32)`

SetClientOfflineSessionMaxLifespan sets ClientOfflineSessionMaxLifespan field to given value.

### HasClientOfflineSessionMaxLifespan

`func (o *RealmRepresentation) HasClientOfflineSessionMaxLifespan() bool`

HasClientOfflineSessionMaxLifespan returns a boolean if a field has been set.

### GetAccessCodeLifespan

`func (o *RealmRepresentation) GetAccessCodeLifespan() int32`

GetAccessCodeLifespan returns the AccessCodeLifespan field if non-nil, zero value otherwise.

### GetAccessCodeLifespanOk

`func (o *RealmRepresentation) GetAccessCodeLifespanOk() (*int32, bool)`

GetAccessCodeLifespanOk returns a tuple with the AccessCodeLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeLifespan

`func (o *RealmRepresentation) SetAccessCodeLifespan(v int32)`

SetAccessCodeLifespan sets AccessCodeLifespan field to given value.

### HasAccessCodeLifespan

`func (o *RealmRepresentation) HasAccessCodeLifespan() bool`

HasAccessCodeLifespan returns a boolean if a field has been set.

### GetAccessCodeLifespanUserAction

`func (o *RealmRepresentation) GetAccessCodeLifespanUserAction() int32`

GetAccessCodeLifespanUserAction returns the AccessCodeLifespanUserAction field if non-nil, zero value otherwise.

### GetAccessCodeLifespanUserActionOk

`func (o *RealmRepresentation) GetAccessCodeLifespanUserActionOk() (*int32, bool)`

GetAccessCodeLifespanUserActionOk returns a tuple with the AccessCodeLifespanUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeLifespanUserAction

`func (o *RealmRepresentation) SetAccessCodeLifespanUserAction(v int32)`

SetAccessCodeLifespanUserAction sets AccessCodeLifespanUserAction field to given value.

### HasAccessCodeLifespanUserAction

`func (o *RealmRepresentation) HasAccessCodeLifespanUserAction() bool`

HasAccessCodeLifespanUserAction returns a boolean if a field has been set.

### GetAccessCodeLifespanLogin

`func (o *RealmRepresentation) GetAccessCodeLifespanLogin() int32`

GetAccessCodeLifespanLogin returns the AccessCodeLifespanLogin field if non-nil, zero value otherwise.

### GetAccessCodeLifespanLoginOk

`func (o *RealmRepresentation) GetAccessCodeLifespanLoginOk() (*int32, bool)`

GetAccessCodeLifespanLoginOk returns a tuple with the AccessCodeLifespanLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCodeLifespanLogin

`func (o *RealmRepresentation) SetAccessCodeLifespanLogin(v int32)`

SetAccessCodeLifespanLogin sets AccessCodeLifespanLogin field to given value.

### HasAccessCodeLifespanLogin

`func (o *RealmRepresentation) HasAccessCodeLifespanLogin() bool`

HasAccessCodeLifespanLogin returns a boolean if a field has been set.

### GetActionTokenGeneratedByAdminLifespan

`func (o *RealmRepresentation) GetActionTokenGeneratedByAdminLifespan() int32`

GetActionTokenGeneratedByAdminLifespan returns the ActionTokenGeneratedByAdminLifespan field if non-nil, zero value otherwise.

### GetActionTokenGeneratedByAdminLifespanOk

`func (o *RealmRepresentation) GetActionTokenGeneratedByAdminLifespanOk() (*int32, bool)`

GetActionTokenGeneratedByAdminLifespanOk returns a tuple with the ActionTokenGeneratedByAdminLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTokenGeneratedByAdminLifespan

`func (o *RealmRepresentation) SetActionTokenGeneratedByAdminLifespan(v int32)`

SetActionTokenGeneratedByAdminLifespan sets ActionTokenGeneratedByAdminLifespan field to given value.

### HasActionTokenGeneratedByAdminLifespan

`func (o *RealmRepresentation) HasActionTokenGeneratedByAdminLifespan() bool`

HasActionTokenGeneratedByAdminLifespan returns a boolean if a field has been set.

### GetActionTokenGeneratedByUserLifespan

`func (o *RealmRepresentation) GetActionTokenGeneratedByUserLifespan() int32`

GetActionTokenGeneratedByUserLifespan returns the ActionTokenGeneratedByUserLifespan field if non-nil, zero value otherwise.

### GetActionTokenGeneratedByUserLifespanOk

`func (o *RealmRepresentation) GetActionTokenGeneratedByUserLifespanOk() (*int32, bool)`

GetActionTokenGeneratedByUserLifespanOk returns a tuple with the ActionTokenGeneratedByUserLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTokenGeneratedByUserLifespan

`func (o *RealmRepresentation) SetActionTokenGeneratedByUserLifespan(v int32)`

SetActionTokenGeneratedByUserLifespan sets ActionTokenGeneratedByUserLifespan field to given value.

### HasActionTokenGeneratedByUserLifespan

`func (o *RealmRepresentation) HasActionTokenGeneratedByUserLifespan() bool`

HasActionTokenGeneratedByUserLifespan returns a boolean if a field has been set.

### GetOauth2DeviceCodeLifespan

`func (o *RealmRepresentation) GetOauth2DeviceCodeLifespan() int32`

GetOauth2DeviceCodeLifespan returns the Oauth2DeviceCodeLifespan field if non-nil, zero value otherwise.

### GetOauth2DeviceCodeLifespanOk

`func (o *RealmRepresentation) GetOauth2DeviceCodeLifespanOk() (*int32, bool)`

GetOauth2DeviceCodeLifespanOk returns a tuple with the Oauth2DeviceCodeLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2DeviceCodeLifespan

`func (o *RealmRepresentation) SetOauth2DeviceCodeLifespan(v int32)`

SetOauth2DeviceCodeLifespan sets Oauth2DeviceCodeLifespan field to given value.

### HasOauth2DeviceCodeLifespan

`func (o *RealmRepresentation) HasOauth2DeviceCodeLifespan() bool`

HasOauth2DeviceCodeLifespan returns a boolean if a field has been set.

### GetOauth2DevicePollingInterval

`func (o *RealmRepresentation) GetOauth2DevicePollingInterval() int32`

GetOauth2DevicePollingInterval returns the Oauth2DevicePollingInterval field if non-nil, zero value otherwise.

### GetOauth2DevicePollingIntervalOk

`func (o *RealmRepresentation) GetOauth2DevicePollingIntervalOk() (*int32, bool)`

GetOauth2DevicePollingIntervalOk returns a tuple with the Oauth2DevicePollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2DevicePollingInterval

`func (o *RealmRepresentation) SetOauth2DevicePollingInterval(v int32)`

SetOauth2DevicePollingInterval sets Oauth2DevicePollingInterval field to given value.

### HasOauth2DevicePollingInterval

`func (o *RealmRepresentation) HasOauth2DevicePollingInterval() bool`

HasOauth2DevicePollingInterval returns a boolean if a field has been set.

### GetEnabled

`func (o *RealmRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RealmRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RealmRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RealmRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSslRequired

`func (o *RealmRepresentation) GetSslRequired() string`

GetSslRequired returns the SslRequired field if non-nil, zero value otherwise.

### GetSslRequiredOk

`func (o *RealmRepresentation) GetSslRequiredOk() (*string, bool)`

GetSslRequiredOk returns a tuple with the SslRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRequired

`func (o *RealmRepresentation) SetSslRequired(v string)`

SetSslRequired sets SslRequired field to given value.

### HasSslRequired

`func (o *RealmRepresentation) HasSslRequired() bool`

HasSslRequired returns a boolean if a field has been set.

### GetPasswordCredentialGrantAllowed

`func (o *RealmRepresentation) GetPasswordCredentialGrantAllowed() bool`

GetPasswordCredentialGrantAllowed returns the PasswordCredentialGrantAllowed field if non-nil, zero value otherwise.

### GetPasswordCredentialGrantAllowedOk

`func (o *RealmRepresentation) GetPasswordCredentialGrantAllowedOk() (*bool, bool)`

GetPasswordCredentialGrantAllowedOk returns a tuple with the PasswordCredentialGrantAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordCredentialGrantAllowed

`func (o *RealmRepresentation) SetPasswordCredentialGrantAllowed(v bool)`

SetPasswordCredentialGrantAllowed sets PasswordCredentialGrantAllowed field to given value.

### HasPasswordCredentialGrantAllowed

`func (o *RealmRepresentation) HasPasswordCredentialGrantAllowed() bool`

HasPasswordCredentialGrantAllowed returns a boolean if a field has been set.

### GetRegistrationAllowed

`func (o *RealmRepresentation) GetRegistrationAllowed() bool`

GetRegistrationAllowed returns the RegistrationAllowed field if non-nil, zero value otherwise.

### GetRegistrationAllowedOk

`func (o *RealmRepresentation) GetRegistrationAllowedOk() (*bool, bool)`

GetRegistrationAllowedOk returns a tuple with the RegistrationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationAllowed

`func (o *RealmRepresentation) SetRegistrationAllowed(v bool)`

SetRegistrationAllowed sets RegistrationAllowed field to given value.

### HasRegistrationAllowed

`func (o *RealmRepresentation) HasRegistrationAllowed() bool`

HasRegistrationAllowed returns a boolean if a field has been set.

### GetRegistrationEmailAsUsername

`func (o *RealmRepresentation) GetRegistrationEmailAsUsername() bool`

GetRegistrationEmailAsUsername returns the RegistrationEmailAsUsername field if non-nil, zero value otherwise.

### GetRegistrationEmailAsUsernameOk

`func (o *RealmRepresentation) GetRegistrationEmailAsUsernameOk() (*bool, bool)`

GetRegistrationEmailAsUsernameOk returns a tuple with the RegistrationEmailAsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEmailAsUsername

`func (o *RealmRepresentation) SetRegistrationEmailAsUsername(v bool)`

SetRegistrationEmailAsUsername sets RegistrationEmailAsUsername field to given value.

### HasRegistrationEmailAsUsername

`func (o *RealmRepresentation) HasRegistrationEmailAsUsername() bool`

HasRegistrationEmailAsUsername returns a boolean if a field has been set.

### GetRememberMe

`func (o *RealmRepresentation) GetRememberMe() bool`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *RealmRepresentation) GetRememberMeOk() (*bool, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *RealmRepresentation) SetRememberMe(v bool)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *RealmRepresentation) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetVerifyEmail

`func (o *RealmRepresentation) GetVerifyEmail() bool`

GetVerifyEmail returns the VerifyEmail field if non-nil, zero value otherwise.

### GetVerifyEmailOk

`func (o *RealmRepresentation) GetVerifyEmailOk() (*bool, bool)`

GetVerifyEmailOk returns a tuple with the VerifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyEmail

`func (o *RealmRepresentation) SetVerifyEmail(v bool)`

SetVerifyEmail sets VerifyEmail field to given value.

### HasVerifyEmail

`func (o *RealmRepresentation) HasVerifyEmail() bool`

HasVerifyEmail returns a boolean if a field has been set.

### GetLoginWithEmailAllowed

`func (o *RealmRepresentation) GetLoginWithEmailAllowed() bool`

GetLoginWithEmailAllowed returns the LoginWithEmailAllowed field if non-nil, zero value otherwise.

### GetLoginWithEmailAllowedOk

`func (o *RealmRepresentation) GetLoginWithEmailAllowedOk() (*bool, bool)`

GetLoginWithEmailAllowedOk returns a tuple with the LoginWithEmailAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginWithEmailAllowed

`func (o *RealmRepresentation) SetLoginWithEmailAllowed(v bool)`

SetLoginWithEmailAllowed sets LoginWithEmailAllowed field to given value.

### HasLoginWithEmailAllowed

`func (o *RealmRepresentation) HasLoginWithEmailAllowed() bool`

HasLoginWithEmailAllowed returns a boolean if a field has been set.

### GetDuplicateEmailsAllowed

`func (o *RealmRepresentation) GetDuplicateEmailsAllowed() bool`

GetDuplicateEmailsAllowed returns the DuplicateEmailsAllowed field if non-nil, zero value otherwise.

### GetDuplicateEmailsAllowedOk

`func (o *RealmRepresentation) GetDuplicateEmailsAllowedOk() (*bool, bool)`

GetDuplicateEmailsAllowedOk returns a tuple with the DuplicateEmailsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateEmailsAllowed

`func (o *RealmRepresentation) SetDuplicateEmailsAllowed(v bool)`

SetDuplicateEmailsAllowed sets DuplicateEmailsAllowed field to given value.

### HasDuplicateEmailsAllowed

`func (o *RealmRepresentation) HasDuplicateEmailsAllowed() bool`

HasDuplicateEmailsAllowed returns a boolean if a field has been set.

### GetResetPasswordAllowed

`func (o *RealmRepresentation) GetResetPasswordAllowed() bool`

GetResetPasswordAllowed returns the ResetPasswordAllowed field if non-nil, zero value otherwise.

### GetResetPasswordAllowedOk

`func (o *RealmRepresentation) GetResetPasswordAllowedOk() (*bool, bool)`

GetResetPasswordAllowedOk returns a tuple with the ResetPasswordAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordAllowed

`func (o *RealmRepresentation) SetResetPasswordAllowed(v bool)`

SetResetPasswordAllowed sets ResetPasswordAllowed field to given value.

### HasResetPasswordAllowed

`func (o *RealmRepresentation) HasResetPasswordAllowed() bool`

HasResetPasswordAllowed returns a boolean if a field has been set.

### GetEditUsernameAllowed

`func (o *RealmRepresentation) GetEditUsernameAllowed() bool`

GetEditUsernameAllowed returns the EditUsernameAllowed field if non-nil, zero value otherwise.

### GetEditUsernameAllowedOk

`func (o *RealmRepresentation) GetEditUsernameAllowedOk() (*bool, bool)`

GetEditUsernameAllowedOk returns a tuple with the EditUsernameAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditUsernameAllowed

`func (o *RealmRepresentation) SetEditUsernameAllowed(v bool)`

SetEditUsernameAllowed sets EditUsernameAllowed field to given value.

### HasEditUsernameAllowed

`func (o *RealmRepresentation) HasEditUsernameAllowed() bool`

HasEditUsernameAllowed returns a boolean if a field has been set.

### GetUserCacheEnabled

`func (o *RealmRepresentation) GetUserCacheEnabled() bool`

GetUserCacheEnabled returns the UserCacheEnabled field if non-nil, zero value otherwise.

### GetUserCacheEnabledOk

`func (o *RealmRepresentation) GetUserCacheEnabledOk() (*bool, bool)`

GetUserCacheEnabledOk returns a tuple with the UserCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCacheEnabled

`func (o *RealmRepresentation) SetUserCacheEnabled(v bool)`

SetUserCacheEnabled sets UserCacheEnabled field to given value.

### HasUserCacheEnabled

`func (o *RealmRepresentation) HasUserCacheEnabled() bool`

HasUserCacheEnabled returns a boolean if a field has been set.

### GetRealmCacheEnabled

`func (o *RealmRepresentation) GetRealmCacheEnabled() bool`

GetRealmCacheEnabled returns the RealmCacheEnabled field if non-nil, zero value otherwise.

### GetRealmCacheEnabledOk

`func (o *RealmRepresentation) GetRealmCacheEnabledOk() (*bool, bool)`

GetRealmCacheEnabledOk returns a tuple with the RealmCacheEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmCacheEnabled

`func (o *RealmRepresentation) SetRealmCacheEnabled(v bool)`

SetRealmCacheEnabled sets RealmCacheEnabled field to given value.

### HasRealmCacheEnabled

`func (o *RealmRepresentation) HasRealmCacheEnabled() bool`

HasRealmCacheEnabled returns a boolean if a field has been set.

### GetBruteForceProtected

`func (o *RealmRepresentation) GetBruteForceProtected() bool`

GetBruteForceProtected returns the BruteForceProtected field if non-nil, zero value otherwise.

### GetBruteForceProtectedOk

`func (o *RealmRepresentation) GetBruteForceProtectedOk() (*bool, bool)`

GetBruteForceProtectedOk returns a tuple with the BruteForceProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBruteForceProtected

`func (o *RealmRepresentation) SetBruteForceProtected(v bool)`

SetBruteForceProtected sets BruteForceProtected field to given value.

### HasBruteForceProtected

`func (o *RealmRepresentation) HasBruteForceProtected() bool`

HasBruteForceProtected returns a boolean if a field has been set.

### GetPermanentLockout

`func (o *RealmRepresentation) GetPermanentLockout() bool`

GetPermanentLockout returns the PermanentLockout field if non-nil, zero value otherwise.

### GetPermanentLockoutOk

`func (o *RealmRepresentation) GetPermanentLockoutOk() (*bool, bool)`

GetPermanentLockoutOk returns a tuple with the PermanentLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentLockout

`func (o *RealmRepresentation) SetPermanentLockout(v bool)`

SetPermanentLockout sets PermanentLockout field to given value.

### HasPermanentLockout

`func (o *RealmRepresentation) HasPermanentLockout() bool`

HasPermanentLockout returns a boolean if a field has been set.

### GetMaxTemporaryLockouts

`func (o *RealmRepresentation) GetMaxTemporaryLockouts() int32`

GetMaxTemporaryLockouts returns the MaxTemporaryLockouts field if non-nil, zero value otherwise.

### GetMaxTemporaryLockoutsOk

`func (o *RealmRepresentation) GetMaxTemporaryLockoutsOk() (*int32, bool)`

GetMaxTemporaryLockoutsOk returns a tuple with the MaxTemporaryLockouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTemporaryLockouts

`func (o *RealmRepresentation) SetMaxTemporaryLockouts(v int32)`

SetMaxTemporaryLockouts sets MaxTemporaryLockouts field to given value.

### HasMaxTemporaryLockouts

`func (o *RealmRepresentation) HasMaxTemporaryLockouts() bool`

HasMaxTemporaryLockouts returns a boolean if a field has been set.

### GetBruteForceStrategy

`func (o *RealmRepresentation) GetBruteForceStrategy() BruteForceStrategy`

GetBruteForceStrategy returns the BruteForceStrategy field if non-nil, zero value otherwise.

### GetBruteForceStrategyOk

`func (o *RealmRepresentation) GetBruteForceStrategyOk() (*BruteForceStrategy, bool)`

GetBruteForceStrategyOk returns a tuple with the BruteForceStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBruteForceStrategy

`func (o *RealmRepresentation) SetBruteForceStrategy(v BruteForceStrategy)`

SetBruteForceStrategy sets BruteForceStrategy field to given value.

### HasBruteForceStrategy

`func (o *RealmRepresentation) HasBruteForceStrategy() bool`

HasBruteForceStrategy returns a boolean if a field has been set.

### GetMaxFailureWaitSeconds

`func (o *RealmRepresentation) GetMaxFailureWaitSeconds() int32`

GetMaxFailureWaitSeconds returns the MaxFailureWaitSeconds field if non-nil, zero value otherwise.

### GetMaxFailureWaitSecondsOk

`func (o *RealmRepresentation) GetMaxFailureWaitSecondsOk() (*int32, bool)`

GetMaxFailureWaitSecondsOk returns a tuple with the MaxFailureWaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailureWaitSeconds

`func (o *RealmRepresentation) SetMaxFailureWaitSeconds(v int32)`

SetMaxFailureWaitSeconds sets MaxFailureWaitSeconds field to given value.

### HasMaxFailureWaitSeconds

`func (o *RealmRepresentation) HasMaxFailureWaitSeconds() bool`

HasMaxFailureWaitSeconds returns a boolean if a field has been set.

### GetMinimumQuickLoginWaitSeconds

`func (o *RealmRepresentation) GetMinimumQuickLoginWaitSeconds() int32`

GetMinimumQuickLoginWaitSeconds returns the MinimumQuickLoginWaitSeconds field if non-nil, zero value otherwise.

### GetMinimumQuickLoginWaitSecondsOk

`func (o *RealmRepresentation) GetMinimumQuickLoginWaitSecondsOk() (*int32, bool)`

GetMinimumQuickLoginWaitSecondsOk returns a tuple with the MinimumQuickLoginWaitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumQuickLoginWaitSeconds

`func (o *RealmRepresentation) SetMinimumQuickLoginWaitSeconds(v int32)`

SetMinimumQuickLoginWaitSeconds sets MinimumQuickLoginWaitSeconds field to given value.

### HasMinimumQuickLoginWaitSeconds

`func (o *RealmRepresentation) HasMinimumQuickLoginWaitSeconds() bool`

HasMinimumQuickLoginWaitSeconds returns a boolean if a field has been set.

### GetWaitIncrementSeconds

`func (o *RealmRepresentation) GetWaitIncrementSeconds() int32`

GetWaitIncrementSeconds returns the WaitIncrementSeconds field if non-nil, zero value otherwise.

### GetWaitIncrementSecondsOk

`func (o *RealmRepresentation) GetWaitIncrementSecondsOk() (*int32, bool)`

GetWaitIncrementSecondsOk returns a tuple with the WaitIncrementSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitIncrementSeconds

`func (o *RealmRepresentation) SetWaitIncrementSeconds(v int32)`

SetWaitIncrementSeconds sets WaitIncrementSeconds field to given value.

### HasWaitIncrementSeconds

`func (o *RealmRepresentation) HasWaitIncrementSeconds() bool`

HasWaitIncrementSeconds returns a boolean if a field has been set.

### GetQuickLoginCheckMilliSeconds

`func (o *RealmRepresentation) GetQuickLoginCheckMilliSeconds() int64`

GetQuickLoginCheckMilliSeconds returns the QuickLoginCheckMilliSeconds field if non-nil, zero value otherwise.

### GetQuickLoginCheckMilliSecondsOk

`func (o *RealmRepresentation) GetQuickLoginCheckMilliSecondsOk() (*int64, bool)`

GetQuickLoginCheckMilliSecondsOk returns a tuple with the QuickLoginCheckMilliSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickLoginCheckMilliSeconds

`func (o *RealmRepresentation) SetQuickLoginCheckMilliSeconds(v int64)`

SetQuickLoginCheckMilliSeconds sets QuickLoginCheckMilliSeconds field to given value.

### HasQuickLoginCheckMilliSeconds

`func (o *RealmRepresentation) HasQuickLoginCheckMilliSeconds() bool`

HasQuickLoginCheckMilliSeconds returns a boolean if a field has been set.

### GetMaxDeltaTimeSeconds

`func (o *RealmRepresentation) GetMaxDeltaTimeSeconds() int32`

GetMaxDeltaTimeSeconds returns the MaxDeltaTimeSeconds field if non-nil, zero value otherwise.

### GetMaxDeltaTimeSecondsOk

`func (o *RealmRepresentation) GetMaxDeltaTimeSecondsOk() (*int32, bool)`

GetMaxDeltaTimeSecondsOk returns a tuple with the MaxDeltaTimeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeltaTimeSeconds

`func (o *RealmRepresentation) SetMaxDeltaTimeSeconds(v int32)`

SetMaxDeltaTimeSeconds sets MaxDeltaTimeSeconds field to given value.

### HasMaxDeltaTimeSeconds

`func (o *RealmRepresentation) HasMaxDeltaTimeSeconds() bool`

HasMaxDeltaTimeSeconds returns a boolean if a field has been set.

### GetFailureFactor

`func (o *RealmRepresentation) GetFailureFactor() int32`

GetFailureFactor returns the FailureFactor field if non-nil, zero value otherwise.

### GetFailureFactorOk

`func (o *RealmRepresentation) GetFailureFactorOk() (*int32, bool)`

GetFailureFactorOk returns a tuple with the FailureFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureFactor

`func (o *RealmRepresentation) SetFailureFactor(v int32)`

SetFailureFactor sets FailureFactor field to given value.

### HasFailureFactor

`func (o *RealmRepresentation) HasFailureFactor() bool`

HasFailureFactor returns a boolean if a field has been set.

### GetPrivateKey

`func (o *RealmRepresentation) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *RealmRepresentation) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *RealmRepresentation) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *RealmRepresentation) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *RealmRepresentation) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RealmRepresentation) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RealmRepresentation) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RealmRepresentation) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetCertificate

`func (o *RealmRepresentation) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *RealmRepresentation) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *RealmRepresentation) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *RealmRepresentation) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCodeSecret

`func (o *RealmRepresentation) GetCodeSecret() string`

GetCodeSecret returns the CodeSecret field if non-nil, zero value otherwise.

### GetCodeSecretOk

`func (o *RealmRepresentation) GetCodeSecretOk() (*string, bool)`

GetCodeSecretOk returns a tuple with the CodeSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSecret

`func (o *RealmRepresentation) SetCodeSecret(v string)`

SetCodeSecret sets CodeSecret field to given value.

### HasCodeSecret

`func (o *RealmRepresentation) HasCodeSecret() bool`

HasCodeSecret returns a boolean if a field has been set.

### GetRoles

`func (o *RealmRepresentation) GetRoles() RolesRepresentation`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RealmRepresentation) GetRolesOk() (*RolesRepresentation, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RealmRepresentation) SetRoles(v RolesRepresentation)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RealmRepresentation) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetGroups

`func (o *RealmRepresentation) GetGroups() []GroupRepresentation`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RealmRepresentation) GetGroupsOk() (*[]GroupRepresentation, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RealmRepresentation) SetGroups(v []GroupRepresentation)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RealmRepresentation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetDefaultRoles

`func (o *RealmRepresentation) GetDefaultRoles() []string`

GetDefaultRoles returns the DefaultRoles field if non-nil, zero value otherwise.

### GetDefaultRolesOk

`func (o *RealmRepresentation) GetDefaultRolesOk() (*[]string, bool)`

GetDefaultRolesOk returns a tuple with the DefaultRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoles

`func (o *RealmRepresentation) SetDefaultRoles(v []string)`

SetDefaultRoles sets DefaultRoles field to given value.

### HasDefaultRoles

`func (o *RealmRepresentation) HasDefaultRoles() bool`

HasDefaultRoles returns a boolean if a field has been set.

### GetDefaultRole

`func (o *RealmRepresentation) GetDefaultRole() RoleRepresentation`

GetDefaultRole returns the DefaultRole field if non-nil, zero value otherwise.

### GetDefaultRoleOk

`func (o *RealmRepresentation) GetDefaultRoleOk() (*RoleRepresentation, bool)`

GetDefaultRoleOk returns a tuple with the DefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRole

`func (o *RealmRepresentation) SetDefaultRole(v RoleRepresentation)`

SetDefaultRole sets DefaultRole field to given value.

### HasDefaultRole

`func (o *RealmRepresentation) HasDefaultRole() bool`

HasDefaultRole returns a boolean if a field has been set.

### GetAdminPermissionsClient

`func (o *RealmRepresentation) GetAdminPermissionsClient() ClientRepresentation`

GetAdminPermissionsClient returns the AdminPermissionsClient field if non-nil, zero value otherwise.

### GetAdminPermissionsClientOk

`func (o *RealmRepresentation) GetAdminPermissionsClientOk() (*ClientRepresentation, bool)`

GetAdminPermissionsClientOk returns a tuple with the AdminPermissionsClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermissionsClient

`func (o *RealmRepresentation) SetAdminPermissionsClient(v ClientRepresentation)`

SetAdminPermissionsClient sets AdminPermissionsClient field to given value.

### HasAdminPermissionsClient

`func (o *RealmRepresentation) HasAdminPermissionsClient() bool`

HasAdminPermissionsClient returns a boolean if a field has been set.

### GetDefaultGroups

`func (o *RealmRepresentation) GetDefaultGroups() []string`

GetDefaultGroups returns the DefaultGroups field if non-nil, zero value otherwise.

### GetDefaultGroupsOk

`func (o *RealmRepresentation) GetDefaultGroupsOk() (*[]string, bool)`

GetDefaultGroupsOk returns a tuple with the DefaultGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroups

`func (o *RealmRepresentation) SetDefaultGroups(v []string)`

SetDefaultGroups sets DefaultGroups field to given value.

### HasDefaultGroups

`func (o *RealmRepresentation) HasDefaultGroups() bool`

HasDefaultGroups returns a boolean if a field has been set.

### GetRequiredCredentials

`func (o *RealmRepresentation) GetRequiredCredentials() []string`

GetRequiredCredentials returns the RequiredCredentials field if non-nil, zero value otherwise.

### GetRequiredCredentialsOk

`func (o *RealmRepresentation) GetRequiredCredentialsOk() (*[]string, bool)`

GetRequiredCredentialsOk returns a tuple with the RequiredCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCredentials

`func (o *RealmRepresentation) SetRequiredCredentials(v []string)`

SetRequiredCredentials sets RequiredCredentials field to given value.

### HasRequiredCredentials

`func (o *RealmRepresentation) HasRequiredCredentials() bool`

HasRequiredCredentials returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *RealmRepresentation) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *RealmRepresentation) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *RealmRepresentation) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *RealmRepresentation) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetOtpPolicyType

`func (o *RealmRepresentation) GetOtpPolicyType() string`

GetOtpPolicyType returns the OtpPolicyType field if non-nil, zero value otherwise.

### GetOtpPolicyTypeOk

`func (o *RealmRepresentation) GetOtpPolicyTypeOk() (*string, bool)`

GetOtpPolicyTypeOk returns a tuple with the OtpPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyType

`func (o *RealmRepresentation) SetOtpPolicyType(v string)`

SetOtpPolicyType sets OtpPolicyType field to given value.

### HasOtpPolicyType

`func (o *RealmRepresentation) HasOtpPolicyType() bool`

HasOtpPolicyType returns a boolean if a field has been set.

### GetOtpPolicyAlgorithm

`func (o *RealmRepresentation) GetOtpPolicyAlgorithm() string`

GetOtpPolicyAlgorithm returns the OtpPolicyAlgorithm field if non-nil, zero value otherwise.

### GetOtpPolicyAlgorithmOk

`func (o *RealmRepresentation) GetOtpPolicyAlgorithmOk() (*string, bool)`

GetOtpPolicyAlgorithmOk returns a tuple with the OtpPolicyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyAlgorithm

`func (o *RealmRepresentation) SetOtpPolicyAlgorithm(v string)`

SetOtpPolicyAlgorithm sets OtpPolicyAlgorithm field to given value.

### HasOtpPolicyAlgorithm

`func (o *RealmRepresentation) HasOtpPolicyAlgorithm() bool`

HasOtpPolicyAlgorithm returns a boolean if a field has been set.

### GetOtpPolicyInitialCounter

`func (o *RealmRepresentation) GetOtpPolicyInitialCounter() int32`

GetOtpPolicyInitialCounter returns the OtpPolicyInitialCounter field if non-nil, zero value otherwise.

### GetOtpPolicyInitialCounterOk

`func (o *RealmRepresentation) GetOtpPolicyInitialCounterOk() (*int32, bool)`

GetOtpPolicyInitialCounterOk returns a tuple with the OtpPolicyInitialCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyInitialCounter

`func (o *RealmRepresentation) SetOtpPolicyInitialCounter(v int32)`

SetOtpPolicyInitialCounter sets OtpPolicyInitialCounter field to given value.

### HasOtpPolicyInitialCounter

`func (o *RealmRepresentation) HasOtpPolicyInitialCounter() bool`

HasOtpPolicyInitialCounter returns a boolean if a field has been set.

### GetOtpPolicyDigits

`func (o *RealmRepresentation) GetOtpPolicyDigits() int32`

GetOtpPolicyDigits returns the OtpPolicyDigits field if non-nil, zero value otherwise.

### GetOtpPolicyDigitsOk

`func (o *RealmRepresentation) GetOtpPolicyDigitsOk() (*int32, bool)`

GetOtpPolicyDigitsOk returns a tuple with the OtpPolicyDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyDigits

`func (o *RealmRepresentation) SetOtpPolicyDigits(v int32)`

SetOtpPolicyDigits sets OtpPolicyDigits field to given value.

### HasOtpPolicyDigits

`func (o *RealmRepresentation) HasOtpPolicyDigits() bool`

HasOtpPolicyDigits returns a boolean if a field has been set.

### GetOtpPolicyLookAheadWindow

`func (o *RealmRepresentation) GetOtpPolicyLookAheadWindow() int32`

GetOtpPolicyLookAheadWindow returns the OtpPolicyLookAheadWindow field if non-nil, zero value otherwise.

### GetOtpPolicyLookAheadWindowOk

`func (o *RealmRepresentation) GetOtpPolicyLookAheadWindowOk() (*int32, bool)`

GetOtpPolicyLookAheadWindowOk returns a tuple with the OtpPolicyLookAheadWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyLookAheadWindow

`func (o *RealmRepresentation) SetOtpPolicyLookAheadWindow(v int32)`

SetOtpPolicyLookAheadWindow sets OtpPolicyLookAheadWindow field to given value.

### HasOtpPolicyLookAheadWindow

`func (o *RealmRepresentation) HasOtpPolicyLookAheadWindow() bool`

HasOtpPolicyLookAheadWindow returns a boolean if a field has been set.

### GetOtpPolicyPeriod

`func (o *RealmRepresentation) GetOtpPolicyPeriod() int32`

GetOtpPolicyPeriod returns the OtpPolicyPeriod field if non-nil, zero value otherwise.

### GetOtpPolicyPeriodOk

`func (o *RealmRepresentation) GetOtpPolicyPeriodOk() (*int32, bool)`

GetOtpPolicyPeriodOk returns a tuple with the OtpPolicyPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyPeriod

`func (o *RealmRepresentation) SetOtpPolicyPeriod(v int32)`

SetOtpPolicyPeriod sets OtpPolicyPeriod field to given value.

### HasOtpPolicyPeriod

`func (o *RealmRepresentation) HasOtpPolicyPeriod() bool`

HasOtpPolicyPeriod returns a boolean if a field has been set.

### GetOtpPolicyCodeReusable

`func (o *RealmRepresentation) GetOtpPolicyCodeReusable() bool`

GetOtpPolicyCodeReusable returns the OtpPolicyCodeReusable field if non-nil, zero value otherwise.

### GetOtpPolicyCodeReusableOk

`func (o *RealmRepresentation) GetOtpPolicyCodeReusableOk() (*bool, bool)`

GetOtpPolicyCodeReusableOk returns a tuple with the OtpPolicyCodeReusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpPolicyCodeReusable

`func (o *RealmRepresentation) SetOtpPolicyCodeReusable(v bool)`

SetOtpPolicyCodeReusable sets OtpPolicyCodeReusable field to given value.

### HasOtpPolicyCodeReusable

`func (o *RealmRepresentation) HasOtpPolicyCodeReusable() bool`

HasOtpPolicyCodeReusable returns a boolean if a field has been set.

### GetOtpSupportedApplications

`func (o *RealmRepresentation) GetOtpSupportedApplications() []string`

GetOtpSupportedApplications returns the OtpSupportedApplications field if non-nil, zero value otherwise.

### GetOtpSupportedApplicationsOk

`func (o *RealmRepresentation) GetOtpSupportedApplicationsOk() (*[]string, bool)`

GetOtpSupportedApplicationsOk returns a tuple with the OtpSupportedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpSupportedApplications

`func (o *RealmRepresentation) SetOtpSupportedApplications(v []string)`

SetOtpSupportedApplications sets OtpSupportedApplications field to given value.

### HasOtpSupportedApplications

`func (o *RealmRepresentation) HasOtpSupportedApplications() bool`

HasOtpSupportedApplications returns a boolean if a field has been set.

### GetLocalizationTexts

`func (o *RealmRepresentation) GetLocalizationTexts() map[string]map[string]string`

GetLocalizationTexts returns the LocalizationTexts field if non-nil, zero value otherwise.

### GetLocalizationTextsOk

`func (o *RealmRepresentation) GetLocalizationTextsOk() (*map[string]map[string]string, bool)`

GetLocalizationTextsOk returns a tuple with the LocalizationTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationTexts

`func (o *RealmRepresentation) SetLocalizationTexts(v map[string]map[string]string)`

SetLocalizationTexts sets LocalizationTexts field to given value.

### HasLocalizationTexts

`func (o *RealmRepresentation) HasLocalizationTexts() bool`

HasLocalizationTexts returns a boolean if a field has been set.

### GetWebAuthnPolicyRpEntityName

`func (o *RealmRepresentation) GetWebAuthnPolicyRpEntityName() string`

GetWebAuthnPolicyRpEntityName returns the WebAuthnPolicyRpEntityName field if non-nil, zero value otherwise.

### GetWebAuthnPolicyRpEntityNameOk

`func (o *RealmRepresentation) GetWebAuthnPolicyRpEntityNameOk() (*string, bool)`

GetWebAuthnPolicyRpEntityNameOk returns a tuple with the WebAuthnPolicyRpEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyRpEntityName

`func (o *RealmRepresentation) SetWebAuthnPolicyRpEntityName(v string)`

SetWebAuthnPolicyRpEntityName sets WebAuthnPolicyRpEntityName field to given value.

### HasWebAuthnPolicyRpEntityName

`func (o *RealmRepresentation) HasWebAuthnPolicyRpEntityName() bool`

HasWebAuthnPolicyRpEntityName returns a boolean if a field has been set.

### GetWebAuthnPolicySignatureAlgorithms

`func (o *RealmRepresentation) GetWebAuthnPolicySignatureAlgorithms() []string`

GetWebAuthnPolicySignatureAlgorithms returns the WebAuthnPolicySignatureAlgorithms field if non-nil, zero value otherwise.

### GetWebAuthnPolicySignatureAlgorithmsOk

`func (o *RealmRepresentation) GetWebAuthnPolicySignatureAlgorithmsOk() (*[]string, bool)`

GetWebAuthnPolicySignatureAlgorithmsOk returns a tuple with the WebAuthnPolicySignatureAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicySignatureAlgorithms

`func (o *RealmRepresentation) SetWebAuthnPolicySignatureAlgorithms(v []string)`

SetWebAuthnPolicySignatureAlgorithms sets WebAuthnPolicySignatureAlgorithms field to given value.

### HasWebAuthnPolicySignatureAlgorithms

`func (o *RealmRepresentation) HasWebAuthnPolicySignatureAlgorithms() bool`

HasWebAuthnPolicySignatureAlgorithms returns a boolean if a field has been set.

### GetWebAuthnPolicyRpId

`func (o *RealmRepresentation) GetWebAuthnPolicyRpId() string`

GetWebAuthnPolicyRpId returns the WebAuthnPolicyRpId field if non-nil, zero value otherwise.

### GetWebAuthnPolicyRpIdOk

`func (o *RealmRepresentation) GetWebAuthnPolicyRpIdOk() (*string, bool)`

GetWebAuthnPolicyRpIdOk returns a tuple with the WebAuthnPolicyRpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyRpId

`func (o *RealmRepresentation) SetWebAuthnPolicyRpId(v string)`

SetWebAuthnPolicyRpId sets WebAuthnPolicyRpId field to given value.

### HasWebAuthnPolicyRpId

`func (o *RealmRepresentation) HasWebAuthnPolicyRpId() bool`

HasWebAuthnPolicyRpId returns a boolean if a field has been set.

### GetWebAuthnPolicyAttestationConveyancePreference

`func (o *RealmRepresentation) GetWebAuthnPolicyAttestationConveyancePreference() string`

GetWebAuthnPolicyAttestationConveyancePreference returns the WebAuthnPolicyAttestationConveyancePreference field if non-nil, zero value otherwise.

### GetWebAuthnPolicyAttestationConveyancePreferenceOk

`func (o *RealmRepresentation) GetWebAuthnPolicyAttestationConveyancePreferenceOk() (*string, bool)`

GetWebAuthnPolicyAttestationConveyancePreferenceOk returns a tuple with the WebAuthnPolicyAttestationConveyancePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyAttestationConveyancePreference

`func (o *RealmRepresentation) SetWebAuthnPolicyAttestationConveyancePreference(v string)`

SetWebAuthnPolicyAttestationConveyancePreference sets WebAuthnPolicyAttestationConveyancePreference field to given value.

### HasWebAuthnPolicyAttestationConveyancePreference

`func (o *RealmRepresentation) HasWebAuthnPolicyAttestationConveyancePreference() bool`

HasWebAuthnPolicyAttestationConveyancePreference returns a boolean if a field has been set.

### GetWebAuthnPolicyAuthenticatorAttachment

`func (o *RealmRepresentation) GetWebAuthnPolicyAuthenticatorAttachment() string`

GetWebAuthnPolicyAuthenticatorAttachment returns the WebAuthnPolicyAuthenticatorAttachment field if non-nil, zero value otherwise.

### GetWebAuthnPolicyAuthenticatorAttachmentOk

`func (o *RealmRepresentation) GetWebAuthnPolicyAuthenticatorAttachmentOk() (*string, bool)`

GetWebAuthnPolicyAuthenticatorAttachmentOk returns a tuple with the WebAuthnPolicyAuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyAuthenticatorAttachment

`func (o *RealmRepresentation) SetWebAuthnPolicyAuthenticatorAttachment(v string)`

SetWebAuthnPolicyAuthenticatorAttachment sets WebAuthnPolicyAuthenticatorAttachment field to given value.

### HasWebAuthnPolicyAuthenticatorAttachment

`func (o *RealmRepresentation) HasWebAuthnPolicyAuthenticatorAttachment() bool`

HasWebAuthnPolicyAuthenticatorAttachment returns a boolean if a field has been set.

### GetWebAuthnPolicyRequireResidentKey

`func (o *RealmRepresentation) GetWebAuthnPolicyRequireResidentKey() string`

GetWebAuthnPolicyRequireResidentKey returns the WebAuthnPolicyRequireResidentKey field if non-nil, zero value otherwise.

### GetWebAuthnPolicyRequireResidentKeyOk

`func (o *RealmRepresentation) GetWebAuthnPolicyRequireResidentKeyOk() (*string, bool)`

GetWebAuthnPolicyRequireResidentKeyOk returns a tuple with the WebAuthnPolicyRequireResidentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyRequireResidentKey

`func (o *RealmRepresentation) SetWebAuthnPolicyRequireResidentKey(v string)`

SetWebAuthnPolicyRequireResidentKey sets WebAuthnPolicyRequireResidentKey field to given value.

### HasWebAuthnPolicyRequireResidentKey

`func (o *RealmRepresentation) HasWebAuthnPolicyRequireResidentKey() bool`

HasWebAuthnPolicyRequireResidentKey returns a boolean if a field has been set.

### GetWebAuthnPolicyUserVerificationRequirement

`func (o *RealmRepresentation) GetWebAuthnPolicyUserVerificationRequirement() string`

GetWebAuthnPolicyUserVerificationRequirement returns the WebAuthnPolicyUserVerificationRequirement field if non-nil, zero value otherwise.

### GetWebAuthnPolicyUserVerificationRequirementOk

`func (o *RealmRepresentation) GetWebAuthnPolicyUserVerificationRequirementOk() (*string, bool)`

GetWebAuthnPolicyUserVerificationRequirementOk returns a tuple with the WebAuthnPolicyUserVerificationRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyUserVerificationRequirement

`func (o *RealmRepresentation) SetWebAuthnPolicyUserVerificationRequirement(v string)`

SetWebAuthnPolicyUserVerificationRequirement sets WebAuthnPolicyUserVerificationRequirement field to given value.

### HasWebAuthnPolicyUserVerificationRequirement

`func (o *RealmRepresentation) HasWebAuthnPolicyUserVerificationRequirement() bool`

HasWebAuthnPolicyUserVerificationRequirement returns a boolean if a field has been set.

### GetWebAuthnPolicyCreateTimeout

`func (o *RealmRepresentation) GetWebAuthnPolicyCreateTimeout() int32`

GetWebAuthnPolicyCreateTimeout returns the WebAuthnPolicyCreateTimeout field if non-nil, zero value otherwise.

### GetWebAuthnPolicyCreateTimeoutOk

`func (o *RealmRepresentation) GetWebAuthnPolicyCreateTimeoutOk() (*int32, bool)`

GetWebAuthnPolicyCreateTimeoutOk returns a tuple with the WebAuthnPolicyCreateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyCreateTimeout

`func (o *RealmRepresentation) SetWebAuthnPolicyCreateTimeout(v int32)`

SetWebAuthnPolicyCreateTimeout sets WebAuthnPolicyCreateTimeout field to given value.

### HasWebAuthnPolicyCreateTimeout

`func (o *RealmRepresentation) HasWebAuthnPolicyCreateTimeout() bool`

HasWebAuthnPolicyCreateTimeout returns a boolean if a field has been set.

### GetWebAuthnPolicyAvoidSameAuthenticatorRegister

`func (o *RealmRepresentation) GetWebAuthnPolicyAvoidSameAuthenticatorRegister() bool`

GetWebAuthnPolicyAvoidSameAuthenticatorRegister returns the WebAuthnPolicyAvoidSameAuthenticatorRegister field if non-nil, zero value otherwise.

### GetWebAuthnPolicyAvoidSameAuthenticatorRegisterOk

`func (o *RealmRepresentation) GetWebAuthnPolicyAvoidSameAuthenticatorRegisterOk() (*bool, bool)`

GetWebAuthnPolicyAvoidSameAuthenticatorRegisterOk returns a tuple with the WebAuthnPolicyAvoidSameAuthenticatorRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyAvoidSameAuthenticatorRegister

`func (o *RealmRepresentation) SetWebAuthnPolicyAvoidSameAuthenticatorRegister(v bool)`

SetWebAuthnPolicyAvoidSameAuthenticatorRegister sets WebAuthnPolicyAvoidSameAuthenticatorRegister field to given value.

### HasWebAuthnPolicyAvoidSameAuthenticatorRegister

`func (o *RealmRepresentation) HasWebAuthnPolicyAvoidSameAuthenticatorRegister() bool`

HasWebAuthnPolicyAvoidSameAuthenticatorRegister returns a boolean if a field has been set.

### GetWebAuthnPolicyAcceptableAaguids

`func (o *RealmRepresentation) GetWebAuthnPolicyAcceptableAaguids() []string`

GetWebAuthnPolicyAcceptableAaguids returns the WebAuthnPolicyAcceptableAaguids field if non-nil, zero value otherwise.

### GetWebAuthnPolicyAcceptableAaguidsOk

`func (o *RealmRepresentation) GetWebAuthnPolicyAcceptableAaguidsOk() (*[]string, bool)`

GetWebAuthnPolicyAcceptableAaguidsOk returns a tuple with the WebAuthnPolicyAcceptableAaguids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyAcceptableAaguids

`func (o *RealmRepresentation) SetWebAuthnPolicyAcceptableAaguids(v []string)`

SetWebAuthnPolicyAcceptableAaguids sets WebAuthnPolicyAcceptableAaguids field to given value.

### HasWebAuthnPolicyAcceptableAaguids

`func (o *RealmRepresentation) HasWebAuthnPolicyAcceptableAaguids() bool`

HasWebAuthnPolicyAcceptableAaguids returns a boolean if a field has been set.

### GetWebAuthnPolicyExtraOrigins

`func (o *RealmRepresentation) GetWebAuthnPolicyExtraOrigins() []string`

GetWebAuthnPolicyExtraOrigins returns the WebAuthnPolicyExtraOrigins field if non-nil, zero value otherwise.

### GetWebAuthnPolicyExtraOriginsOk

`func (o *RealmRepresentation) GetWebAuthnPolicyExtraOriginsOk() (*[]string, bool)`

GetWebAuthnPolicyExtraOriginsOk returns a tuple with the WebAuthnPolicyExtraOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyExtraOrigins

`func (o *RealmRepresentation) SetWebAuthnPolicyExtraOrigins(v []string)`

SetWebAuthnPolicyExtraOrigins sets WebAuthnPolicyExtraOrigins field to given value.

### HasWebAuthnPolicyExtraOrigins

`func (o *RealmRepresentation) HasWebAuthnPolicyExtraOrigins() bool`

HasWebAuthnPolicyExtraOrigins returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessRpEntityName

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessRpEntityName() string`

GetWebAuthnPolicyPasswordlessRpEntityName returns the WebAuthnPolicyPasswordlessRpEntityName field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessRpEntityNameOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessRpEntityNameOk() (*string, bool)`

GetWebAuthnPolicyPasswordlessRpEntityNameOk returns a tuple with the WebAuthnPolicyPasswordlessRpEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessRpEntityName

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessRpEntityName(v string)`

SetWebAuthnPolicyPasswordlessRpEntityName sets WebAuthnPolicyPasswordlessRpEntityName field to given value.

### HasWebAuthnPolicyPasswordlessRpEntityName

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessRpEntityName() bool`

HasWebAuthnPolicyPasswordlessRpEntityName returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessSignatureAlgorithms

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessSignatureAlgorithms() []string`

GetWebAuthnPolicyPasswordlessSignatureAlgorithms returns the WebAuthnPolicyPasswordlessSignatureAlgorithms field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessSignatureAlgorithmsOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessSignatureAlgorithmsOk() (*[]string, bool)`

GetWebAuthnPolicyPasswordlessSignatureAlgorithmsOk returns a tuple with the WebAuthnPolicyPasswordlessSignatureAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessSignatureAlgorithms

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessSignatureAlgorithms(v []string)`

SetWebAuthnPolicyPasswordlessSignatureAlgorithms sets WebAuthnPolicyPasswordlessSignatureAlgorithms field to given value.

### HasWebAuthnPolicyPasswordlessSignatureAlgorithms

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessSignatureAlgorithms() bool`

HasWebAuthnPolicyPasswordlessSignatureAlgorithms returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessRpId

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessRpId() string`

GetWebAuthnPolicyPasswordlessRpId returns the WebAuthnPolicyPasswordlessRpId field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessRpIdOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessRpIdOk() (*string, bool)`

GetWebAuthnPolicyPasswordlessRpIdOk returns a tuple with the WebAuthnPolicyPasswordlessRpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessRpId

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessRpId(v string)`

SetWebAuthnPolicyPasswordlessRpId sets WebAuthnPolicyPasswordlessRpId field to given value.

### HasWebAuthnPolicyPasswordlessRpId

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessRpId() bool`

HasWebAuthnPolicyPasswordlessRpId returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessAttestationConveyancePreference

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAttestationConveyancePreference() string`

GetWebAuthnPolicyPasswordlessAttestationConveyancePreference returns the WebAuthnPolicyPasswordlessAttestationConveyancePreference field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessAttestationConveyancePreferenceOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAttestationConveyancePreferenceOk() (*string, bool)`

GetWebAuthnPolicyPasswordlessAttestationConveyancePreferenceOk returns a tuple with the WebAuthnPolicyPasswordlessAttestationConveyancePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessAttestationConveyancePreference

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessAttestationConveyancePreference(v string)`

SetWebAuthnPolicyPasswordlessAttestationConveyancePreference sets WebAuthnPolicyPasswordlessAttestationConveyancePreference field to given value.

### HasWebAuthnPolicyPasswordlessAttestationConveyancePreference

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessAttestationConveyancePreference() bool`

HasWebAuthnPolicyPasswordlessAttestationConveyancePreference returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessAuthenticatorAttachment

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAuthenticatorAttachment() string`

GetWebAuthnPolicyPasswordlessAuthenticatorAttachment returns the WebAuthnPolicyPasswordlessAuthenticatorAttachment field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessAuthenticatorAttachmentOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAuthenticatorAttachmentOk() (*string, bool)`

GetWebAuthnPolicyPasswordlessAuthenticatorAttachmentOk returns a tuple with the WebAuthnPolicyPasswordlessAuthenticatorAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessAuthenticatorAttachment

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessAuthenticatorAttachment(v string)`

SetWebAuthnPolicyPasswordlessAuthenticatorAttachment sets WebAuthnPolicyPasswordlessAuthenticatorAttachment field to given value.

### HasWebAuthnPolicyPasswordlessAuthenticatorAttachment

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessAuthenticatorAttachment() bool`

HasWebAuthnPolicyPasswordlessAuthenticatorAttachment returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessRequireResidentKey

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessRequireResidentKey() string`

GetWebAuthnPolicyPasswordlessRequireResidentKey returns the WebAuthnPolicyPasswordlessRequireResidentKey field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessRequireResidentKeyOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessRequireResidentKeyOk() (*string, bool)`

GetWebAuthnPolicyPasswordlessRequireResidentKeyOk returns a tuple with the WebAuthnPolicyPasswordlessRequireResidentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessRequireResidentKey

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessRequireResidentKey(v string)`

SetWebAuthnPolicyPasswordlessRequireResidentKey sets WebAuthnPolicyPasswordlessRequireResidentKey field to given value.

### HasWebAuthnPolicyPasswordlessRequireResidentKey

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessRequireResidentKey() bool`

HasWebAuthnPolicyPasswordlessRequireResidentKey returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessUserVerificationRequirement

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessUserVerificationRequirement() string`

GetWebAuthnPolicyPasswordlessUserVerificationRequirement returns the WebAuthnPolicyPasswordlessUserVerificationRequirement field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessUserVerificationRequirementOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessUserVerificationRequirementOk() (*string, bool)`

GetWebAuthnPolicyPasswordlessUserVerificationRequirementOk returns a tuple with the WebAuthnPolicyPasswordlessUserVerificationRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessUserVerificationRequirement

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessUserVerificationRequirement(v string)`

SetWebAuthnPolicyPasswordlessUserVerificationRequirement sets WebAuthnPolicyPasswordlessUserVerificationRequirement field to given value.

### HasWebAuthnPolicyPasswordlessUserVerificationRequirement

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessUserVerificationRequirement() bool`

HasWebAuthnPolicyPasswordlessUserVerificationRequirement returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessCreateTimeout

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessCreateTimeout() int32`

GetWebAuthnPolicyPasswordlessCreateTimeout returns the WebAuthnPolicyPasswordlessCreateTimeout field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessCreateTimeoutOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessCreateTimeoutOk() (*int32, bool)`

GetWebAuthnPolicyPasswordlessCreateTimeoutOk returns a tuple with the WebAuthnPolicyPasswordlessCreateTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessCreateTimeout

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessCreateTimeout(v int32)`

SetWebAuthnPolicyPasswordlessCreateTimeout sets WebAuthnPolicyPasswordlessCreateTimeout field to given value.

### HasWebAuthnPolicyPasswordlessCreateTimeout

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessCreateTimeout() bool`

HasWebAuthnPolicyPasswordlessCreateTimeout returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister() bool`

GetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister returns the WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegisterOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegisterOk() (*bool, bool)`

GetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegisterOk returns a tuple with the WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister(v bool)`

SetWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister sets WebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister field to given value.

### HasWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister() bool`

HasWebAuthnPolicyPasswordlessAvoidSameAuthenticatorRegister returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessAcceptableAaguids

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAcceptableAaguids() []string`

GetWebAuthnPolicyPasswordlessAcceptableAaguids returns the WebAuthnPolicyPasswordlessAcceptableAaguids field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessAcceptableAaguidsOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessAcceptableAaguidsOk() (*[]string, bool)`

GetWebAuthnPolicyPasswordlessAcceptableAaguidsOk returns a tuple with the WebAuthnPolicyPasswordlessAcceptableAaguids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessAcceptableAaguids

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessAcceptableAaguids(v []string)`

SetWebAuthnPolicyPasswordlessAcceptableAaguids sets WebAuthnPolicyPasswordlessAcceptableAaguids field to given value.

### HasWebAuthnPolicyPasswordlessAcceptableAaguids

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessAcceptableAaguids() bool`

HasWebAuthnPolicyPasswordlessAcceptableAaguids returns a boolean if a field has been set.

### GetWebAuthnPolicyPasswordlessExtraOrigins

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessExtraOrigins() []string`

GetWebAuthnPolicyPasswordlessExtraOrigins returns the WebAuthnPolicyPasswordlessExtraOrigins field if non-nil, zero value otherwise.

### GetWebAuthnPolicyPasswordlessExtraOriginsOk

`func (o *RealmRepresentation) GetWebAuthnPolicyPasswordlessExtraOriginsOk() (*[]string, bool)`

GetWebAuthnPolicyPasswordlessExtraOriginsOk returns a tuple with the WebAuthnPolicyPasswordlessExtraOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebAuthnPolicyPasswordlessExtraOrigins

`func (o *RealmRepresentation) SetWebAuthnPolicyPasswordlessExtraOrigins(v []string)`

SetWebAuthnPolicyPasswordlessExtraOrigins sets WebAuthnPolicyPasswordlessExtraOrigins field to given value.

### HasWebAuthnPolicyPasswordlessExtraOrigins

`func (o *RealmRepresentation) HasWebAuthnPolicyPasswordlessExtraOrigins() bool`

HasWebAuthnPolicyPasswordlessExtraOrigins returns a boolean if a field has been set.

### GetClientProfiles

`func (o *RealmRepresentation) GetClientProfiles() ClientProfilesRepresentation`

GetClientProfiles returns the ClientProfiles field if non-nil, zero value otherwise.

### GetClientProfilesOk

`func (o *RealmRepresentation) GetClientProfilesOk() (*ClientProfilesRepresentation, bool)`

GetClientProfilesOk returns a tuple with the ClientProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfiles

`func (o *RealmRepresentation) SetClientProfiles(v ClientProfilesRepresentation)`

SetClientProfiles sets ClientProfiles field to given value.

### HasClientProfiles

`func (o *RealmRepresentation) HasClientProfiles() bool`

HasClientProfiles returns a boolean if a field has been set.

### GetClientPolicies

`func (o *RealmRepresentation) GetClientPolicies() ClientPoliciesRepresentation`

GetClientPolicies returns the ClientPolicies field if non-nil, zero value otherwise.

### GetClientPoliciesOk

`func (o *RealmRepresentation) GetClientPoliciesOk() (*ClientPoliciesRepresentation, bool)`

GetClientPoliciesOk returns a tuple with the ClientPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPolicies

`func (o *RealmRepresentation) SetClientPolicies(v ClientPoliciesRepresentation)`

SetClientPolicies sets ClientPolicies field to given value.

### HasClientPolicies

`func (o *RealmRepresentation) HasClientPolicies() bool`

HasClientPolicies returns a boolean if a field has been set.

### GetUsers

`func (o *RealmRepresentation) GetUsers() []UserRepresentation`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RealmRepresentation) GetUsersOk() (*[]UserRepresentation, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RealmRepresentation) SetUsers(v []UserRepresentation)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *RealmRepresentation) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetFederatedUsers

`func (o *RealmRepresentation) GetFederatedUsers() []UserRepresentation`

GetFederatedUsers returns the FederatedUsers field if non-nil, zero value otherwise.

### GetFederatedUsersOk

`func (o *RealmRepresentation) GetFederatedUsersOk() (*[]UserRepresentation, bool)`

GetFederatedUsersOk returns a tuple with the FederatedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedUsers

`func (o *RealmRepresentation) SetFederatedUsers(v []UserRepresentation)`

SetFederatedUsers sets FederatedUsers field to given value.

### HasFederatedUsers

`func (o *RealmRepresentation) HasFederatedUsers() bool`

HasFederatedUsers returns a boolean if a field has been set.

### GetScopeMappings

`func (o *RealmRepresentation) GetScopeMappings() []ScopeMappingRepresentation`

GetScopeMappings returns the ScopeMappings field if non-nil, zero value otherwise.

### GetScopeMappingsOk

`func (o *RealmRepresentation) GetScopeMappingsOk() (*[]ScopeMappingRepresentation, bool)`

GetScopeMappingsOk returns a tuple with the ScopeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeMappings

`func (o *RealmRepresentation) SetScopeMappings(v []ScopeMappingRepresentation)`

SetScopeMappings sets ScopeMappings field to given value.

### HasScopeMappings

`func (o *RealmRepresentation) HasScopeMappings() bool`

HasScopeMappings returns a boolean if a field has been set.

### GetClientScopeMappings

`func (o *RealmRepresentation) GetClientScopeMappings() map[string][]ScopeMappingRepresentation`

GetClientScopeMappings returns the ClientScopeMappings field if non-nil, zero value otherwise.

### GetClientScopeMappingsOk

`func (o *RealmRepresentation) GetClientScopeMappingsOk() (*map[string][]ScopeMappingRepresentation, bool)`

GetClientScopeMappingsOk returns a tuple with the ClientScopeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientScopeMappings

`func (o *RealmRepresentation) SetClientScopeMappings(v map[string][]ScopeMappingRepresentation)`

SetClientScopeMappings sets ClientScopeMappings field to given value.

### HasClientScopeMappings

`func (o *RealmRepresentation) HasClientScopeMappings() bool`

HasClientScopeMappings returns a boolean if a field has been set.

### GetClients

`func (o *RealmRepresentation) GetClients() []ClientRepresentation`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *RealmRepresentation) GetClientsOk() (*[]ClientRepresentation, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *RealmRepresentation) SetClients(v []ClientRepresentation)`

SetClients sets Clients field to given value.

### HasClients

`func (o *RealmRepresentation) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetClientScopes

`func (o *RealmRepresentation) GetClientScopes() []ClientScopeRepresentation`

GetClientScopes returns the ClientScopes field if non-nil, zero value otherwise.

### GetClientScopesOk

`func (o *RealmRepresentation) GetClientScopesOk() (*[]ClientScopeRepresentation, bool)`

GetClientScopesOk returns a tuple with the ClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientScopes

`func (o *RealmRepresentation) SetClientScopes(v []ClientScopeRepresentation)`

SetClientScopes sets ClientScopes field to given value.

### HasClientScopes

`func (o *RealmRepresentation) HasClientScopes() bool`

HasClientScopes returns a boolean if a field has been set.

### GetDefaultDefaultClientScopes

`func (o *RealmRepresentation) GetDefaultDefaultClientScopes() []string`

GetDefaultDefaultClientScopes returns the DefaultDefaultClientScopes field if non-nil, zero value otherwise.

### GetDefaultDefaultClientScopesOk

`func (o *RealmRepresentation) GetDefaultDefaultClientScopesOk() (*[]string, bool)`

GetDefaultDefaultClientScopesOk returns a tuple with the DefaultDefaultClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDefaultClientScopes

`func (o *RealmRepresentation) SetDefaultDefaultClientScopes(v []string)`

SetDefaultDefaultClientScopes sets DefaultDefaultClientScopes field to given value.

### HasDefaultDefaultClientScopes

`func (o *RealmRepresentation) HasDefaultDefaultClientScopes() bool`

HasDefaultDefaultClientScopes returns a boolean if a field has been set.

### GetDefaultOptionalClientScopes

`func (o *RealmRepresentation) GetDefaultOptionalClientScopes() []string`

GetDefaultOptionalClientScopes returns the DefaultOptionalClientScopes field if non-nil, zero value otherwise.

### GetDefaultOptionalClientScopesOk

`func (o *RealmRepresentation) GetDefaultOptionalClientScopesOk() (*[]string, bool)`

GetDefaultOptionalClientScopesOk returns a tuple with the DefaultOptionalClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOptionalClientScopes

`func (o *RealmRepresentation) SetDefaultOptionalClientScopes(v []string)`

SetDefaultOptionalClientScopes sets DefaultOptionalClientScopes field to given value.

### HasDefaultOptionalClientScopes

`func (o *RealmRepresentation) HasDefaultOptionalClientScopes() bool`

HasDefaultOptionalClientScopes returns a boolean if a field has been set.

### GetBrowserSecurityHeaders

`func (o *RealmRepresentation) GetBrowserSecurityHeaders() map[string]string`

GetBrowserSecurityHeaders returns the BrowserSecurityHeaders field if non-nil, zero value otherwise.

### GetBrowserSecurityHeadersOk

`func (o *RealmRepresentation) GetBrowserSecurityHeadersOk() (*map[string]string, bool)`

GetBrowserSecurityHeadersOk returns a tuple with the BrowserSecurityHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserSecurityHeaders

`func (o *RealmRepresentation) SetBrowserSecurityHeaders(v map[string]string)`

SetBrowserSecurityHeaders sets BrowserSecurityHeaders field to given value.

### HasBrowserSecurityHeaders

`func (o *RealmRepresentation) HasBrowserSecurityHeaders() bool`

HasBrowserSecurityHeaders returns a boolean if a field has been set.

### GetSmtpServer

`func (o *RealmRepresentation) GetSmtpServer() map[string]string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *RealmRepresentation) GetSmtpServerOk() (*map[string]string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *RealmRepresentation) SetSmtpServer(v map[string]string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *RealmRepresentation) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetUserFederationProviders

`func (o *RealmRepresentation) GetUserFederationProviders() []UserFederationProviderRepresentation`

GetUserFederationProviders returns the UserFederationProviders field if non-nil, zero value otherwise.

### GetUserFederationProvidersOk

`func (o *RealmRepresentation) GetUserFederationProvidersOk() (*[]UserFederationProviderRepresentation, bool)`

GetUserFederationProvidersOk returns a tuple with the UserFederationProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFederationProviders

`func (o *RealmRepresentation) SetUserFederationProviders(v []UserFederationProviderRepresentation)`

SetUserFederationProviders sets UserFederationProviders field to given value.

### HasUserFederationProviders

`func (o *RealmRepresentation) HasUserFederationProviders() bool`

HasUserFederationProviders returns a boolean if a field has been set.

### GetUserFederationMappers

`func (o *RealmRepresentation) GetUserFederationMappers() []UserFederationMapperRepresentation`

GetUserFederationMappers returns the UserFederationMappers field if non-nil, zero value otherwise.

### GetUserFederationMappersOk

`func (o *RealmRepresentation) GetUserFederationMappersOk() (*[]UserFederationMapperRepresentation, bool)`

GetUserFederationMappersOk returns a tuple with the UserFederationMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFederationMappers

`func (o *RealmRepresentation) SetUserFederationMappers(v []UserFederationMapperRepresentation)`

SetUserFederationMappers sets UserFederationMappers field to given value.

### HasUserFederationMappers

`func (o *RealmRepresentation) HasUserFederationMappers() bool`

HasUserFederationMappers returns a boolean if a field has been set.

### GetLoginTheme

`func (o *RealmRepresentation) GetLoginTheme() string`

GetLoginTheme returns the LoginTheme field if non-nil, zero value otherwise.

### GetLoginThemeOk

`func (o *RealmRepresentation) GetLoginThemeOk() (*string, bool)`

GetLoginThemeOk returns a tuple with the LoginTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTheme

`func (o *RealmRepresentation) SetLoginTheme(v string)`

SetLoginTheme sets LoginTheme field to given value.

### HasLoginTheme

`func (o *RealmRepresentation) HasLoginTheme() bool`

HasLoginTheme returns a boolean if a field has been set.

### GetAccountTheme

`func (o *RealmRepresentation) GetAccountTheme() string`

GetAccountTheme returns the AccountTheme field if non-nil, zero value otherwise.

### GetAccountThemeOk

`func (o *RealmRepresentation) GetAccountThemeOk() (*string, bool)`

GetAccountThemeOk returns a tuple with the AccountTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTheme

`func (o *RealmRepresentation) SetAccountTheme(v string)`

SetAccountTheme sets AccountTheme field to given value.

### HasAccountTheme

`func (o *RealmRepresentation) HasAccountTheme() bool`

HasAccountTheme returns a boolean if a field has been set.

### GetAdminTheme

`func (o *RealmRepresentation) GetAdminTheme() string`

GetAdminTheme returns the AdminTheme field if non-nil, zero value otherwise.

### GetAdminThemeOk

`func (o *RealmRepresentation) GetAdminThemeOk() (*string, bool)`

GetAdminThemeOk returns a tuple with the AdminTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTheme

`func (o *RealmRepresentation) SetAdminTheme(v string)`

SetAdminTheme sets AdminTheme field to given value.

### HasAdminTheme

`func (o *RealmRepresentation) HasAdminTheme() bool`

HasAdminTheme returns a boolean if a field has been set.

### GetEmailTheme

`func (o *RealmRepresentation) GetEmailTheme() string`

GetEmailTheme returns the EmailTheme field if non-nil, zero value otherwise.

### GetEmailThemeOk

`func (o *RealmRepresentation) GetEmailThemeOk() (*string, bool)`

GetEmailThemeOk returns a tuple with the EmailTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTheme

`func (o *RealmRepresentation) SetEmailTheme(v string)`

SetEmailTheme sets EmailTheme field to given value.

### HasEmailTheme

`func (o *RealmRepresentation) HasEmailTheme() bool`

HasEmailTheme returns a boolean if a field has been set.

### GetEventsEnabled

`func (o *RealmRepresentation) GetEventsEnabled() bool`

GetEventsEnabled returns the EventsEnabled field if non-nil, zero value otherwise.

### GetEventsEnabledOk

`func (o *RealmRepresentation) GetEventsEnabledOk() (*bool, bool)`

GetEventsEnabledOk returns a tuple with the EventsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsEnabled

`func (o *RealmRepresentation) SetEventsEnabled(v bool)`

SetEventsEnabled sets EventsEnabled field to given value.

### HasEventsEnabled

`func (o *RealmRepresentation) HasEventsEnabled() bool`

HasEventsEnabled returns a boolean if a field has been set.

### GetEventsExpiration

`func (o *RealmRepresentation) GetEventsExpiration() int64`

GetEventsExpiration returns the EventsExpiration field if non-nil, zero value otherwise.

### GetEventsExpirationOk

`func (o *RealmRepresentation) GetEventsExpirationOk() (*int64, bool)`

GetEventsExpirationOk returns a tuple with the EventsExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsExpiration

`func (o *RealmRepresentation) SetEventsExpiration(v int64)`

SetEventsExpiration sets EventsExpiration field to given value.

### HasEventsExpiration

`func (o *RealmRepresentation) HasEventsExpiration() bool`

HasEventsExpiration returns a boolean if a field has been set.

### GetEventsListeners

`func (o *RealmRepresentation) GetEventsListeners() []string`

GetEventsListeners returns the EventsListeners field if non-nil, zero value otherwise.

### GetEventsListenersOk

`func (o *RealmRepresentation) GetEventsListenersOk() (*[]string, bool)`

GetEventsListenersOk returns a tuple with the EventsListeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsListeners

`func (o *RealmRepresentation) SetEventsListeners(v []string)`

SetEventsListeners sets EventsListeners field to given value.

### HasEventsListeners

`func (o *RealmRepresentation) HasEventsListeners() bool`

HasEventsListeners returns a boolean if a field has been set.

### GetEnabledEventTypes

`func (o *RealmRepresentation) GetEnabledEventTypes() []string`

GetEnabledEventTypes returns the EnabledEventTypes field if non-nil, zero value otherwise.

### GetEnabledEventTypesOk

`func (o *RealmRepresentation) GetEnabledEventTypesOk() (*[]string, bool)`

GetEnabledEventTypesOk returns a tuple with the EnabledEventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledEventTypes

`func (o *RealmRepresentation) SetEnabledEventTypes(v []string)`

SetEnabledEventTypes sets EnabledEventTypes field to given value.

### HasEnabledEventTypes

`func (o *RealmRepresentation) HasEnabledEventTypes() bool`

HasEnabledEventTypes returns a boolean if a field has been set.

### GetAdminEventsEnabled

`func (o *RealmRepresentation) GetAdminEventsEnabled() bool`

GetAdminEventsEnabled returns the AdminEventsEnabled field if non-nil, zero value otherwise.

### GetAdminEventsEnabledOk

`func (o *RealmRepresentation) GetAdminEventsEnabledOk() (*bool, bool)`

GetAdminEventsEnabledOk returns a tuple with the AdminEventsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEventsEnabled

`func (o *RealmRepresentation) SetAdminEventsEnabled(v bool)`

SetAdminEventsEnabled sets AdminEventsEnabled field to given value.

### HasAdminEventsEnabled

`func (o *RealmRepresentation) HasAdminEventsEnabled() bool`

HasAdminEventsEnabled returns a boolean if a field has been set.

### GetAdminEventsDetailsEnabled

`func (o *RealmRepresentation) GetAdminEventsDetailsEnabled() bool`

GetAdminEventsDetailsEnabled returns the AdminEventsDetailsEnabled field if non-nil, zero value otherwise.

### GetAdminEventsDetailsEnabledOk

`func (o *RealmRepresentation) GetAdminEventsDetailsEnabledOk() (*bool, bool)`

GetAdminEventsDetailsEnabledOk returns a tuple with the AdminEventsDetailsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEventsDetailsEnabled

`func (o *RealmRepresentation) SetAdminEventsDetailsEnabled(v bool)`

SetAdminEventsDetailsEnabled sets AdminEventsDetailsEnabled field to given value.

### HasAdminEventsDetailsEnabled

`func (o *RealmRepresentation) HasAdminEventsDetailsEnabled() bool`

HasAdminEventsDetailsEnabled returns a boolean if a field has been set.

### GetIdentityProviders

`func (o *RealmRepresentation) GetIdentityProviders() []IdentityProviderRepresentation`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *RealmRepresentation) GetIdentityProvidersOk() (*[]IdentityProviderRepresentation, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *RealmRepresentation) SetIdentityProviders(v []IdentityProviderRepresentation)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *RealmRepresentation) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.

### GetIdentityProviderMappers

`func (o *RealmRepresentation) GetIdentityProviderMappers() []IdentityProviderMapperRepresentation`

GetIdentityProviderMappers returns the IdentityProviderMappers field if non-nil, zero value otherwise.

### GetIdentityProviderMappersOk

`func (o *RealmRepresentation) GetIdentityProviderMappersOk() (*[]IdentityProviderMapperRepresentation, bool)`

GetIdentityProviderMappersOk returns a tuple with the IdentityProviderMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderMappers

`func (o *RealmRepresentation) SetIdentityProviderMappers(v []IdentityProviderMapperRepresentation)`

SetIdentityProviderMappers sets IdentityProviderMappers field to given value.

### HasIdentityProviderMappers

`func (o *RealmRepresentation) HasIdentityProviderMappers() bool`

HasIdentityProviderMappers returns a boolean if a field has been set.

### GetProtocolMappers

`func (o *RealmRepresentation) GetProtocolMappers() []ProtocolMapperRepresentation`

GetProtocolMappers returns the ProtocolMappers field if non-nil, zero value otherwise.

### GetProtocolMappersOk

`func (o *RealmRepresentation) GetProtocolMappersOk() (*[]ProtocolMapperRepresentation, bool)`

GetProtocolMappersOk returns a tuple with the ProtocolMappers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMappers

`func (o *RealmRepresentation) SetProtocolMappers(v []ProtocolMapperRepresentation)`

SetProtocolMappers sets ProtocolMappers field to given value.

### HasProtocolMappers

`func (o *RealmRepresentation) HasProtocolMappers() bool`

HasProtocolMappers returns a boolean if a field has been set.

### GetComponents

`func (o *RealmRepresentation) GetComponents() map[string][]ComponentExportRepresentation`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *RealmRepresentation) GetComponentsOk() (*map[string][]ComponentExportRepresentation, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *RealmRepresentation) SetComponents(v map[string][]ComponentExportRepresentation)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *RealmRepresentation) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetInternationalizationEnabled

`func (o *RealmRepresentation) GetInternationalizationEnabled() bool`

GetInternationalizationEnabled returns the InternationalizationEnabled field if non-nil, zero value otherwise.

### GetInternationalizationEnabledOk

`func (o *RealmRepresentation) GetInternationalizationEnabledOk() (*bool, bool)`

GetInternationalizationEnabledOk returns a tuple with the InternationalizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalizationEnabled

`func (o *RealmRepresentation) SetInternationalizationEnabled(v bool)`

SetInternationalizationEnabled sets InternationalizationEnabled field to given value.

### HasInternationalizationEnabled

`func (o *RealmRepresentation) HasInternationalizationEnabled() bool`

HasInternationalizationEnabled returns a boolean if a field has been set.

### GetSupportedLocales

`func (o *RealmRepresentation) GetSupportedLocales() []string`

GetSupportedLocales returns the SupportedLocales field if non-nil, zero value otherwise.

### GetSupportedLocalesOk

`func (o *RealmRepresentation) GetSupportedLocalesOk() (*[]string, bool)`

GetSupportedLocalesOk returns a tuple with the SupportedLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedLocales

`func (o *RealmRepresentation) SetSupportedLocales(v []string)`

SetSupportedLocales sets SupportedLocales field to given value.

### HasSupportedLocales

`func (o *RealmRepresentation) HasSupportedLocales() bool`

HasSupportedLocales returns a boolean if a field has been set.

### GetDefaultLocale

`func (o *RealmRepresentation) GetDefaultLocale() string`

GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.

### GetDefaultLocaleOk

`func (o *RealmRepresentation) GetDefaultLocaleOk() (*string, bool)`

GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocale

`func (o *RealmRepresentation) SetDefaultLocale(v string)`

SetDefaultLocale sets DefaultLocale field to given value.

### HasDefaultLocale

`func (o *RealmRepresentation) HasDefaultLocale() bool`

HasDefaultLocale returns a boolean if a field has been set.

### GetAuthenticationFlows

`func (o *RealmRepresentation) GetAuthenticationFlows() []AuthenticationFlowRepresentation`

GetAuthenticationFlows returns the AuthenticationFlows field if non-nil, zero value otherwise.

### GetAuthenticationFlowsOk

`func (o *RealmRepresentation) GetAuthenticationFlowsOk() (*[]AuthenticationFlowRepresentation, bool)`

GetAuthenticationFlowsOk returns a tuple with the AuthenticationFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlows

`func (o *RealmRepresentation) SetAuthenticationFlows(v []AuthenticationFlowRepresentation)`

SetAuthenticationFlows sets AuthenticationFlows field to given value.

### HasAuthenticationFlows

`func (o *RealmRepresentation) HasAuthenticationFlows() bool`

HasAuthenticationFlows returns a boolean if a field has been set.

### GetAuthenticatorConfig

`func (o *RealmRepresentation) GetAuthenticatorConfig() []AuthenticatorConfigRepresentation`

GetAuthenticatorConfig returns the AuthenticatorConfig field if non-nil, zero value otherwise.

### GetAuthenticatorConfigOk

`func (o *RealmRepresentation) GetAuthenticatorConfigOk() (*[]AuthenticatorConfigRepresentation, bool)`

GetAuthenticatorConfigOk returns a tuple with the AuthenticatorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorConfig

`func (o *RealmRepresentation) SetAuthenticatorConfig(v []AuthenticatorConfigRepresentation)`

SetAuthenticatorConfig sets AuthenticatorConfig field to given value.

### HasAuthenticatorConfig

`func (o *RealmRepresentation) HasAuthenticatorConfig() bool`

HasAuthenticatorConfig returns a boolean if a field has been set.

### GetRequiredActions

`func (o *RealmRepresentation) GetRequiredActions() []RequiredActionProviderRepresentation`

GetRequiredActions returns the RequiredActions field if non-nil, zero value otherwise.

### GetRequiredActionsOk

`func (o *RealmRepresentation) GetRequiredActionsOk() (*[]RequiredActionProviderRepresentation, bool)`

GetRequiredActionsOk returns a tuple with the RequiredActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredActions

`func (o *RealmRepresentation) SetRequiredActions(v []RequiredActionProviderRepresentation)`

SetRequiredActions sets RequiredActions field to given value.

### HasRequiredActions

`func (o *RealmRepresentation) HasRequiredActions() bool`

HasRequiredActions returns a boolean if a field has been set.

### GetBrowserFlow

`func (o *RealmRepresentation) GetBrowserFlow() string`

GetBrowserFlow returns the BrowserFlow field if non-nil, zero value otherwise.

### GetBrowserFlowOk

`func (o *RealmRepresentation) GetBrowserFlowOk() (*string, bool)`

GetBrowserFlowOk returns a tuple with the BrowserFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserFlow

`func (o *RealmRepresentation) SetBrowserFlow(v string)`

SetBrowserFlow sets BrowserFlow field to given value.

### HasBrowserFlow

`func (o *RealmRepresentation) HasBrowserFlow() bool`

HasBrowserFlow returns a boolean if a field has been set.

### GetRegistrationFlow

`func (o *RealmRepresentation) GetRegistrationFlow() string`

GetRegistrationFlow returns the RegistrationFlow field if non-nil, zero value otherwise.

### GetRegistrationFlowOk

`func (o *RealmRepresentation) GetRegistrationFlowOk() (*string, bool)`

GetRegistrationFlowOk returns a tuple with the RegistrationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationFlow

`func (o *RealmRepresentation) SetRegistrationFlow(v string)`

SetRegistrationFlow sets RegistrationFlow field to given value.

### HasRegistrationFlow

`func (o *RealmRepresentation) HasRegistrationFlow() bool`

HasRegistrationFlow returns a boolean if a field has been set.

### GetDirectGrantFlow

`func (o *RealmRepresentation) GetDirectGrantFlow() string`

GetDirectGrantFlow returns the DirectGrantFlow field if non-nil, zero value otherwise.

### GetDirectGrantFlowOk

`func (o *RealmRepresentation) GetDirectGrantFlowOk() (*string, bool)`

GetDirectGrantFlowOk returns a tuple with the DirectGrantFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectGrantFlow

`func (o *RealmRepresentation) SetDirectGrantFlow(v string)`

SetDirectGrantFlow sets DirectGrantFlow field to given value.

### HasDirectGrantFlow

`func (o *RealmRepresentation) HasDirectGrantFlow() bool`

HasDirectGrantFlow returns a boolean if a field has been set.

### GetResetCredentialsFlow

`func (o *RealmRepresentation) GetResetCredentialsFlow() string`

GetResetCredentialsFlow returns the ResetCredentialsFlow field if non-nil, zero value otherwise.

### GetResetCredentialsFlowOk

`func (o *RealmRepresentation) GetResetCredentialsFlowOk() (*string, bool)`

GetResetCredentialsFlowOk returns a tuple with the ResetCredentialsFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCredentialsFlow

`func (o *RealmRepresentation) SetResetCredentialsFlow(v string)`

SetResetCredentialsFlow sets ResetCredentialsFlow field to given value.

### HasResetCredentialsFlow

`func (o *RealmRepresentation) HasResetCredentialsFlow() bool`

HasResetCredentialsFlow returns a boolean if a field has been set.

### GetClientAuthenticationFlow

`func (o *RealmRepresentation) GetClientAuthenticationFlow() string`

GetClientAuthenticationFlow returns the ClientAuthenticationFlow field if non-nil, zero value otherwise.

### GetClientAuthenticationFlowOk

`func (o *RealmRepresentation) GetClientAuthenticationFlowOk() (*string, bool)`

GetClientAuthenticationFlowOk returns a tuple with the ClientAuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationFlow

`func (o *RealmRepresentation) SetClientAuthenticationFlow(v string)`

SetClientAuthenticationFlow sets ClientAuthenticationFlow field to given value.

### HasClientAuthenticationFlow

`func (o *RealmRepresentation) HasClientAuthenticationFlow() bool`

HasClientAuthenticationFlow returns a boolean if a field has been set.

### GetDockerAuthenticationFlow

`func (o *RealmRepresentation) GetDockerAuthenticationFlow() string`

GetDockerAuthenticationFlow returns the DockerAuthenticationFlow field if non-nil, zero value otherwise.

### GetDockerAuthenticationFlowOk

`func (o *RealmRepresentation) GetDockerAuthenticationFlowOk() (*string, bool)`

GetDockerAuthenticationFlowOk returns a tuple with the DockerAuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerAuthenticationFlow

`func (o *RealmRepresentation) SetDockerAuthenticationFlow(v string)`

SetDockerAuthenticationFlow sets DockerAuthenticationFlow field to given value.

### HasDockerAuthenticationFlow

`func (o *RealmRepresentation) HasDockerAuthenticationFlow() bool`

HasDockerAuthenticationFlow returns a boolean if a field has been set.

### GetFirstBrokerLoginFlow

`func (o *RealmRepresentation) GetFirstBrokerLoginFlow() string`

GetFirstBrokerLoginFlow returns the FirstBrokerLoginFlow field if non-nil, zero value otherwise.

### GetFirstBrokerLoginFlowOk

`func (o *RealmRepresentation) GetFirstBrokerLoginFlowOk() (*string, bool)`

GetFirstBrokerLoginFlowOk returns a tuple with the FirstBrokerLoginFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstBrokerLoginFlow

`func (o *RealmRepresentation) SetFirstBrokerLoginFlow(v string)`

SetFirstBrokerLoginFlow sets FirstBrokerLoginFlow field to given value.

### HasFirstBrokerLoginFlow

`func (o *RealmRepresentation) HasFirstBrokerLoginFlow() bool`

HasFirstBrokerLoginFlow returns a boolean if a field has been set.

### GetAttributes

`func (o *RealmRepresentation) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RealmRepresentation) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RealmRepresentation) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RealmRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetKeycloakVersion

`func (o *RealmRepresentation) GetKeycloakVersion() string`

GetKeycloakVersion returns the KeycloakVersion field if non-nil, zero value otherwise.

### GetKeycloakVersionOk

`func (o *RealmRepresentation) GetKeycloakVersionOk() (*string, bool)`

GetKeycloakVersionOk returns a tuple with the KeycloakVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeycloakVersion

`func (o *RealmRepresentation) SetKeycloakVersion(v string)`

SetKeycloakVersion sets KeycloakVersion field to given value.

### HasKeycloakVersion

`func (o *RealmRepresentation) HasKeycloakVersion() bool`

HasKeycloakVersion returns a boolean if a field has been set.

### GetUserManagedAccessAllowed

`func (o *RealmRepresentation) GetUserManagedAccessAllowed() bool`

GetUserManagedAccessAllowed returns the UserManagedAccessAllowed field if non-nil, zero value otherwise.

### GetUserManagedAccessAllowedOk

`func (o *RealmRepresentation) GetUserManagedAccessAllowedOk() (*bool, bool)`

GetUserManagedAccessAllowedOk returns a tuple with the UserManagedAccessAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagedAccessAllowed

`func (o *RealmRepresentation) SetUserManagedAccessAllowed(v bool)`

SetUserManagedAccessAllowed sets UserManagedAccessAllowed field to given value.

### HasUserManagedAccessAllowed

`func (o *RealmRepresentation) HasUserManagedAccessAllowed() bool`

HasUserManagedAccessAllowed returns a boolean if a field has been set.

### GetOrganizationsEnabled

`func (o *RealmRepresentation) GetOrganizationsEnabled() bool`

GetOrganizationsEnabled returns the OrganizationsEnabled field if non-nil, zero value otherwise.

### GetOrganizationsEnabledOk

`func (o *RealmRepresentation) GetOrganizationsEnabledOk() (*bool, bool)`

GetOrganizationsEnabledOk returns a tuple with the OrganizationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationsEnabled

`func (o *RealmRepresentation) SetOrganizationsEnabled(v bool)`

SetOrganizationsEnabled sets OrganizationsEnabled field to given value.

### HasOrganizationsEnabled

`func (o *RealmRepresentation) HasOrganizationsEnabled() bool`

HasOrganizationsEnabled returns a boolean if a field has been set.

### GetOrganizations

`func (o *RealmRepresentation) GetOrganizations() []OrganizationRepresentation`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *RealmRepresentation) GetOrganizationsOk() (*[]OrganizationRepresentation, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *RealmRepresentation) SetOrganizations(v []OrganizationRepresentation)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *RealmRepresentation) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetVerifiableCredentialsEnabled

`func (o *RealmRepresentation) GetVerifiableCredentialsEnabled() bool`

GetVerifiableCredentialsEnabled returns the VerifiableCredentialsEnabled field if non-nil, zero value otherwise.

### GetVerifiableCredentialsEnabledOk

`func (o *RealmRepresentation) GetVerifiableCredentialsEnabledOk() (*bool, bool)`

GetVerifiableCredentialsEnabledOk returns a tuple with the VerifiableCredentialsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiableCredentialsEnabled

`func (o *RealmRepresentation) SetVerifiableCredentialsEnabled(v bool)`

SetVerifiableCredentialsEnabled sets VerifiableCredentialsEnabled field to given value.

### HasVerifiableCredentialsEnabled

`func (o *RealmRepresentation) HasVerifiableCredentialsEnabled() bool`

HasVerifiableCredentialsEnabled returns a boolean if a field has been set.

### GetAdminPermissionsEnabled

`func (o *RealmRepresentation) GetAdminPermissionsEnabled() bool`

GetAdminPermissionsEnabled returns the AdminPermissionsEnabled field if non-nil, zero value otherwise.

### GetAdminPermissionsEnabledOk

`func (o *RealmRepresentation) GetAdminPermissionsEnabledOk() (*bool, bool)`

GetAdminPermissionsEnabledOk returns a tuple with the AdminPermissionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermissionsEnabled

`func (o *RealmRepresentation) SetAdminPermissionsEnabled(v bool)`

SetAdminPermissionsEnabled sets AdminPermissionsEnabled field to given value.

### HasAdminPermissionsEnabled

`func (o *RealmRepresentation) HasAdminPermissionsEnabled() bool`

HasAdminPermissionsEnabled returns a boolean if a field has been set.

### GetSocial

`func (o *RealmRepresentation) GetSocial() bool`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *RealmRepresentation) GetSocialOk() (*bool, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *RealmRepresentation) SetSocial(v bool)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *RealmRepresentation) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetUpdateProfileOnInitialSocialLogin

`func (o *RealmRepresentation) GetUpdateProfileOnInitialSocialLogin() bool`

GetUpdateProfileOnInitialSocialLogin returns the UpdateProfileOnInitialSocialLogin field if non-nil, zero value otherwise.

### GetUpdateProfileOnInitialSocialLoginOk

`func (o *RealmRepresentation) GetUpdateProfileOnInitialSocialLoginOk() (*bool, bool)`

GetUpdateProfileOnInitialSocialLoginOk returns a tuple with the UpdateProfileOnInitialSocialLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProfileOnInitialSocialLogin

`func (o *RealmRepresentation) SetUpdateProfileOnInitialSocialLogin(v bool)`

SetUpdateProfileOnInitialSocialLogin sets UpdateProfileOnInitialSocialLogin field to given value.

### HasUpdateProfileOnInitialSocialLogin

`func (o *RealmRepresentation) HasUpdateProfileOnInitialSocialLogin() bool`

HasUpdateProfileOnInitialSocialLogin returns a boolean if a field has been set.

### GetSocialProviders

`func (o *RealmRepresentation) GetSocialProviders() map[string]string`

GetSocialProviders returns the SocialProviders field if non-nil, zero value otherwise.

### GetSocialProvidersOk

`func (o *RealmRepresentation) GetSocialProvidersOk() (*map[string]string, bool)`

GetSocialProvidersOk returns a tuple with the SocialProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProviders

`func (o *RealmRepresentation) SetSocialProviders(v map[string]string)`

SetSocialProviders sets SocialProviders field to given value.

### HasSocialProviders

`func (o *RealmRepresentation) HasSocialProviders() bool`

HasSocialProviders returns a boolean if a field has been set.

### GetApplicationScopeMappings

`func (o *RealmRepresentation) GetApplicationScopeMappings() map[string][]ScopeMappingRepresentation`

GetApplicationScopeMappings returns the ApplicationScopeMappings field if non-nil, zero value otherwise.

### GetApplicationScopeMappingsOk

`func (o *RealmRepresentation) GetApplicationScopeMappingsOk() (*map[string][]ScopeMappingRepresentation, bool)`

GetApplicationScopeMappingsOk returns a tuple with the ApplicationScopeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationScopeMappings

`func (o *RealmRepresentation) SetApplicationScopeMappings(v map[string][]ScopeMappingRepresentation)`

SetApplicationScopeMappings sets ApplicationScopeMappings field to given value.

### HasApplicationScopeMappings

`func (o *RealmRepresentation) HasApplicationScopeMappings() bool`

HasApplicationScopeMappings returns a boolean if a field has been set.

### GetApplications

`func (o *RealmRepresentation) GetApplications() []ApplicationRepresentation`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RealmRepresentation) GetApplicationsOk() (*[]ApplicationRepresentation, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RealmRepresentation) SetApplications(v []ApplicationRepresentation)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RealmRepresentation) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetOauthClients

`func (o *RealmRepresentation) GetOauthClients() []OAuthClientRepresentation`

GetOauthClients returns the OauthClients field if non-nil, zero value otherwise.

### GetOauthClientsOk

`func (o *RealmRepresentation) GetOauthClientsOk() (*[]OAuthClientRepresentation, bool)`

GetOauthClientsOk returns a tuple with the OauthClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClients

`func (o *RealmRepresentation) SetOauthClients(v []OAuthClientRepresentation)`

SetOauthClients sets OauthClients field to given value.

### HasOauthClients

`func (o *RealmRepresentation) HasOauthClients() bool`

HasOauthClients returns a boolean if a field has been set.

### GetClientTemplates

`func (o *RealmRepresentation) GetClientTemplates() []ClientTemplateRepresentation`

GetClientTemplates returns the ClientTemplates field if non-nil, zero value otherwise.

### GetClientTemplatesOk

`func (o *RealmRepresentation) GetClientTemplatesOk() (*[]ClientTemplateRepresentation, bool)`

GetClientTemplatesOk returns a tuple with the ClientTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTemplates

`func (o *RealmRepresentation) SetClientTemplates(v []ClientTemplateRepresentation)`

SetClientTemplates sets ClientTemplates field to given value.

### HasClientTemplates

`func (o *RealmRepresentation) HasClientTemplates() bool`

HasClientTemplates returns a boolean if a field has been set.

### GetOAuth2DeviceCodeLifespan

`func (o *RealmRepresentation) GetOAuth2DeviceCodeLifespan() int32`

GetOAuth2DeviceCodeLifespan returns the OAuth2DeviceCodeLifespan field if non-nil, zero value otherwise.

### GetOAuth2DeviceCodeLifespanOk

`func (o *RealmRepresentation) GetOAuth2DeviceCodeLifespanOk() (*int32, bool)`

GetOAuth2DeviceCodeLifespanOk returns a tuple with the OAuth2DeviceCodeLifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2DeviceCodeLifespan

`func (o *RealmRepresentation) SetOAuth2DeviceCodeLifespan(v int32)`

SetOAuth2DeviceCodeLifespan sets OAuth2DeviceCodeLifespan field to given value.

### HasOAuth2DeviceCodeLifespan

`func (o *RealmRepresentation) HasOAuth2DeviceCodeLifespan() bool`

HasOAuth2DeviceCodeLifespan returns a boolean if a field has been set.

### GetOAuth2DevicePollingInterval

`func (o *RealmRepresentation) GetOAuth2DevicePollingInterval() int32`

GetOAuth2DevicePollingInterval returns the OAuth2DevicePollingInterval field if non-nil, zero value otherwise.

### GetOAuth2DevicePollingIntervalOk

`func (o *RealmRepresentation) GetOAuth2DevicePollingIntervalOk() (*int32, bool)`

GetOAuth2DevicePollingIntervalOk returns a tuple with the OAuth2DevicePollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuth2DevicePollingInterval

`func (o *RealmRepresentation) SetOAuth2DevicePollingInterval(v int32)`

SetOAuth2DevicePollingInterval sets OAuth2DevicePollingInterval field to given value.

### HasOAuth2DevicePollingInterval

`func (o *RealmRepresentation) HasOAuth2DevicePollingInterval() bool`

HasOAuth2DevicePollingInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


