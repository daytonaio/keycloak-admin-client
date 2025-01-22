# UserRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**Attributes** | Pointer to **map[string][]string** |  | [optional] 
**UserProfileMetadata** | Pointer to [**UserProfileMetadata**](UserProfileMetadata.md) |  | [optional] 
**Self** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**CreatedTimestamp** | Pointer to **int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Totp** | Pointer to **bool** |  | [optional] 
**FederationLink** | Pointer to **string** |  | [optional] 
**ServiceAccountClientId** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to [**[]CredentialRepresentation**](CredentialRepresentation.md) |  | [optional] 
**DisableableCredentialTypes** | Pointer to **[]string** |  | [optional] 
**RequiredActions** | Pointer to **[]string** |  | [optional] 
**FederatedIdentities** | Pointer to [**[]FederatedIdentityRepresentation**](FederatedIdentityRepresentation.md) |  | [optional] 
**RealmRoles** | Pointer to **[]string** |  | [optional] 
**ClientRoles** | Pointer to **map[string][]string** |  | [optional] 
**ClientConsents** | Pointer to [**[]UserConsentRepresentation**](UserConsentRepresentation.md) |  | [optional] 
**NotBefore** | Pointer to **int32** |  | [optional] 
**ApplicationRoles** | Pointer to **map[string][]string** |  | [optional] 
**SocialLinks** | Pointer to [**[]SocialLinkRepresentation**](SocialLinkRepresentation.md) |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Access** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewUserRepresentation

`func NewUserRepresentation() *UserRepresentation`

NewUserRepresentation instantiates a new UserRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRepresentationWithDefaults

`func NewUserRepresentationWithDefaults() *UserRepresentation`

NewUserRepresentationWithDefaults instantiates a new UserRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *UserRepresentation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRepresentation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRepresentation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserRepresentation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *UserRepresentation) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserRepresentation) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserRepresentation) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserRepresentation) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserRepresentation) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserRepresentation) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserRepresentation) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserRepresentation) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *UserRepresentation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRepresentation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRepresentation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRepresentation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *UserRepresentation) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserRepresentation) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserRepresentation) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UserRepresentation) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetAttributes

`func (o *UserRepresentation) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserRepresentation) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserRepresentation) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UserRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUserProfileMetadata

`func (o *UserRepresentation) GetUserProfileMetadata() UserProfileMetadata`

GetUserProfileMetadata returns the UserProfileMetadata field if non-nil, zero value otherwise.

### GetUserProfileMetadataOk

`func (o *UserRepresentation) GetUserProfileMetadataOk() (*UserProfileMetadata, bool)`

GetUserProfileMetadataOk returns a tuple with the UserProfileMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfileMetadata

`func (o *UserRepresentation) SetUserProfileMetadata(v UserProfileMetadata)`

SetUserProfileMetadata sets UserProfileMetadata field to given value.

### HasUserProfileMetadata

`func (o *UserRepresentation) HasUserProfileMetadata() bool`

HasUserProfileMetadata returns a boolean if a field has been set.

### GetSelf

`func (o *UserRepresentation) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *UserRepresentation) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *UserRepresentation) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *UserRepresentation) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetOrigin

`func (o *UserRepresentation) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UserRepresentation) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UserRepresentation) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *UserRepresentation) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *UserRepresentation) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *UserRepresentation) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *UserRepresentation) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *UserRepresentation) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetEnabled

`func (o *UserRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTotp

`func (o *UserRepresentation) GetTotp() bool`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *UserRepresentation) GetTotpOk() (*bool, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *UserRepresentation) SetTotp(v bool)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *UserRepresentation) HasTotp() bool`

HasTotp returns a boolean if a field has been set.

### GetFederationLink

`func (o *UserRepresentation) GetFederationLink() string`

GetFederationLink returns the FederationLink field if non-nil, zero value otherwise.

### GetFederationLinkOk

`func (o *UserRepresentation) GetFederationLinkOk() (*string, bool)`

GetFederationLinkOk returns a tuple with the FederationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationLink

`func (o *UserRepresentation) SetFederationLink(v string)`

SetFederationLink sets FederationLink field to given value.

### HasFederationLink

`func (o *UserRepresentation) HasFederationLink() bool`

HasFederationLink returns a boolean if a field has been set.

### GetServiceAccountClientId

`func (o *UserRepresentation) GetServiceAccountClientId() string`

GetServiceAccountClientId returns the ServiceAccountClientId field if non-nil, zero value otherwise.

### GetServiceAccountClientIdOk

`func (o *UserRepresentation) GetServiceAccountClientIdOk() (*string, bool)`

GetServiceAccountClientIdOk returns a tuple with the ServiceAccountClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountClientId

`func (o *UserRepresentation) SetServiceAccountClientId(v string)`

SetServiceAccountClientId sets ServiceAccountClientId field to given value.

### HasServiceAccountClientId

`func (o *UserRepresentation) HasServiceAccountClientId() bool`

HasServiceAccountClientId returns a boolean if a field has been set.

### GetCredentials

`func (o *UserRepresentation) GetCredentials() []CredentialRepresentation`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UserRepresentation) GetCredentialsOk() (*[]CredentialRepresentation, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UserRepresentation) SetCredentials(v []CredentialRepresentation)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UserRepresentation) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDisableableCredentialTypes

`func (o *UserRepresentation) GetDisableableCredentialTypes() []string`

GetDisableableCredentialTypes returns the DisableableCredentialTypes field if non-nil, zero value otherwise.

### GetDisableableCredentialTypesOk

`func (o *UserRepresentation) GetDisableableCredentialTypesOk() (*[]string, bool)`

GetDisableableCredentialTypesOk returns a tuple with the DisableableCredentialTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableableCredentialTypes

`func (o *UserRepresentation) SetDisableableCredentialTypes(v []string)`

SetDisableableCredentialTypes sets DisableableCredentialTypes field to given value.

### HasDisableableCredentialTypes

`func (o *UserRepresentation) HasDisableableCredentialTypes() bool`

HasDisableableCredentialTypes returns a boolean if a field has been set.

### GetRequiredActions

`func (o *UserRepresentation) GetRequiredActions() []string`

GetRequiredActions returns the RequiredActions field if non-nil, zero value otherwise.

### GetRequiredActionsOk

`func (o *UserRepresentation) GetRequiredActionsOk() (*[]string, bool)`

GetRequiredActionsOk returns a tuple with the RequiredActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredActions

`func (o *UserRepresentation) SetRequiredActions(v []string)`

SetRequiredActions sets RequiredActions field to given value.

### HasRequiredActions

`func (o *UserRepresentation) HasRequiredActions() bool`

HasRequiredActions returns a boolean if a field has been set.

### GetFederatedIdentities

`func (o *UserRepresentation) GetFederatedIdentities() []FederatedIdentityRepresentation`

GetFederatedIdentities returns the FederatedIdentities field if non-nil, zero value otherwise.

### GetFederatedIdentitiesOk

`func (o *UserRepresentation) GetFederatedIdentitiesOk() (*[]FederatedIdentityRepresentation, bool)`

GetFederatedIdentitiesOk returns a tuple with the FederatedIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedIdentities

`func (o *UserRepresentation) SetFederatedIdentities(v []FederatedIdentityRepresentation)`

SetFederatedIdentities sets FederatedIdentities field to given value.

### HasFederatedIdentities

`func (o *UserRepresentation) HasFederatedIdentities() bool`

HasFederatedIdentities returns a boolean if a field has been set.

### GetRealmRoles

`func (o *UserRepresentation) GetRealmRoles() []string`

GetRealmRoles returns the RealmRoles field if non-nil, zero value otherwise.

### GetRealmRolesOk

`func (o *UserRepresentation) GetRealmRolesOk() (*[]string, bool)`

GetRealmRolesOk returns a tuple with the RealmRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmRoles

`func (o *UserRepresentation) SetRealmRoles(v []string)`

SetRealmRoles sets RealmRoles field to given value.

### HasRealmRoles

`func (o *UserRepresentation) HasRealmRoles() bool`

HasRealmRoles returns a boolean if a field has been set.

### GetClientRoles

`func (o *UserRepresentation) GetClientRoles() map[string][]string`

GetClientRoles returns the ClientRoles field if non-nil, zero value otherwise.

### GetClientRolesOk

`func (o *UserRepresentation) GetClientRolesOk() (*map[string][]string, bool)`

GetClientRolesOk returns a tuple with the ClientRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRoles

`func (o *UserRepresentation) SetClientRoles(v map[string][]string)`

SetClientRoles sets ClientRoles field to given value.

### HasClientRoles

`func (o *UserRepresentation) HasClientRoles() bool`

HasClientRoles returns a boolean if a field has been set.

### GetClientConsents

`func (o *UserRepresentation) GetClientConsents() []UserConsentRepresentation`

GetClientConsents returns the ClientConsents field if non-nil, zero value otherwise.

### GetClientConsentsOk

`func (o *UserRepresentation) GetClientConsentsOk() (*[]UserConsentRepresentation, bool)`

GetClientConsentsOk returns a tuple with the ClientConsents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConsents

`func (o *UserRepresentation) SetClientConsents(v []UserConsentRepresentation)`

SetClientConsents sets ClientConsents field to given value.

### HasClientConsents

`func (o *UserRepresentation) HasClientConsents() bool`

HasClientConsents returns a boolean if a field has been set.

### GetNotBefore

`func (o *UserRepresentation) GetNotBefore() int32`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *UserRepresentation) GetNotBeforeOk() (*int32, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *UserRepresentation) SetNotBefore(v int32)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *UserRepresentation) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetApplicationRoles

`func (o *UserRepresentation) GetApplicationRoles() map[string][]string`

GetApplicationRoles returns the ApplicationRoles field if non-nil, zero value otherwise.

### GetApplicationRolesOk

`func (o *UserRepresentation) GetApplicationRolesOk() (*map[string][]string, bool)`

GetApplicationRolesOk returns a tuple with the ApplicationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationRoles

`func (o *UserRepresentation) SetApplicationRoles(v map[string][]string)`

SetApplicationRoles sets ApplicationRoles field to given value.

### HasApplicationRoles

`func (o *UserRepresentation) HasApplicationRoles() bool`

HasApplicationRoles returns a boolean if a field has been set.

### GetSocialLinks

`func (o *UserRepresentation) GetSocialLinks() []SocialLinkRepresentation`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *UserRepresentation) GetSocialLinksOk() (*[]SocialLinkRepresentation, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *UserRepresentation) SetSocialLinks(v []SocialLinkRepresentation)`

SetSocialLinks sets SocialLinks field to given value.

### HasSocialLinks

`func (o *UserRepresentation) HasSocialLinks() bool`

HasSocialLinks returns a boolean if a field has been set.

### GetGroups

`func (o *UserRepresentation) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserRepresentation) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserRepresentation) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserRepresentation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAccess

`func (o *UserRepresentation) GetAccess() map[string]bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UserRepresentation) GetAccessOk() (*map[string]bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UserRepresentation) SetAccess(v map[string]bool)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UserRepresentation) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


