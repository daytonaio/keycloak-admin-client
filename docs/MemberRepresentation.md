# MemberRepresentation

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
**MembershipType** | Pointer to [**MembershipType**](MembershipType.md) |  | [optional] 

## Methods

### NewMemberRepresentation

`func NewMemberRepresentation() *MemberRepresentation`

NewMemberRepresentation instantiates a new MemberRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberRepresentationWithDefaults

`func NewMemberRepresentationWithDefaults() *MemberRepresentation`

NewMemberRepresentationWithDefaults instantiates a new MemberRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *MemberRepresentation) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MemberRepresentation) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MemberRepresentation) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MemberRepresentation) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *MemberRepresentation) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MemberRepresentation) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MemberRepresentation) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MemberRepresentation) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *MemberRepresentation) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MemberRepresentation) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MemberRepresentation) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MemberRepresentation) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *MemberRepresentation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberRepresentation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberRepresentation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MemberRepresentation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *MemberRepresentation) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *MemberRepresentation) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *MemberRepresentation) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *MemberRepresentation) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetAttributes

`func (o *MemberRepresentation) GetAttributes() map[string][]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MemberRepresentation) GetAttributesOk() (*map[string][]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MemberRepresentation) SetAttributes(v map[string][]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *MemberRepresentation) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetUserProfileMetadata

`func (o *MemberRepresentation) GetUserProfileMetadata() UserProfileMetadata`

GetUserProfileMetadata returns the UserProfileMetadata field if non-nil, zero value otherwise.

### GetUserProfileMetadataOk

`func (o *MemberRepresentation) GetUserProfileMetadataOk() (*UserProfileMetadata, bool)`

GetUserProfileMetadataOk returns a tuple with the UserProfileMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfileMetadata

`func (o *MemberRepresentation) SetUserProfileMetadata(v UserProfileMetadata)`

SetUserProfileMetadata sets UserProfileMetadata field to given value.

### HasUserProfileMetadata

`func (o *MemberRepresentation) HasUserProfileMetadata() bool`

HasUserProfileMetadata returns a boolean if a field has been set.

### GetSelf

`func (o *MemberRepresentation) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MemberRepresentation) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MemberRepresentation) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MemberRepresentation) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetOrigin

`func (o *MemberRepresentation) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MemberRepresentation) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MemberRepresentation) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *MemberRepresentation) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *MemberRepresentation) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *MemberRepresentation) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *MemberRepresentation) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *MemberRepresentation) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetEnabled

`func (o *MemberRepresentation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MemberRepresentation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MemberRepresentation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MemberRepresentation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTotp

`func (o *MemberRepresentation) GetTotp() bool`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *MemberRepresentation) GetTotpOk() (*bool, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *MemberRepresentation) SetTotp(v bool)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *MemberRepresentation) HasTotp() bool`

HasTotp returns a boolean if a field has been set.

### GetFederationLink

`func (o *MemberRepresentation) GetFederationLink() string`

GetFederationLink returns the FederationLink field if non-nil, zero value otherwise.

### GetFederationLinkOk

`func (o *MemberRepresentation) GetFederationLinkOk() (*string, bool)`

GetFederationLinkOk returns a tuple with the FederationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationLink

`func (o *MemberRepresentation) SetFederationLink(v string)`

SetFederationLink sets FederationLink field to given value.

### HasFederationLink

`func (o *MemberRepresentation) HasFederationLink() bool`

HasFederationLink returns a boolean if a field has been set.

### GetServiceAccountClientId

`func (o *MemberRepresentation) GetServiceAccountClientId() string`

GetServiceAccountClientId returns the ServiceAccountClientId field if non-nil, zero value otherwise.

### GetServiceAccountClientIdOk

`func (o *MemberRepresentation) GetServiceAccountClientIdOk() (*string, bool)`

GetServiceAccountClientIdOk returns a tuple with the ServiceAccountClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountClientId

`func (o *MemberRepresentation) SetServiceAccountClientId(v string)`

SetServiceAccountClientId sets ServiceAccountClientId field to given value.

### HasServiceAccountClientId

`func (o *MemberRepresentation) HasServiceAccountClientId() bool`

HasServiceAccountClientId returns a boolean if a field has been set.

### GetCredentials

`func (o *MemberRepresentation) GetCredentials() []CredentialRepresentation`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *MemberRepresentation) GetCredentialsOk() (*[]CredentialRepresentation, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *MemberRepresentation) SetCredentials(v []CredentialRepresentation)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *MemberRepresentation) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetDisableableCredentialTypes

`func (o *MemberRepresentation) GetDisableableCredentialTypes() []string`

GetDisableableCredentialTypes returns the DisableableCredentialTypes field if non-nil, zero value otherwise.

### GetDisableableCredentialTypesOk

`func (o *MemberRepresentation) GetDisableableCredentialTypesOk() (*[]string, bool)`

GetDisableableCredentialTypesOk returns a tuple with the DisableableCredentialTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableableCredentialTypes

`func (o *MemberRepresentation) SetDisableableCredentialTypes(v []string)`

SetDisableableCredentialTypes sets DisableableCredentialTypes field to given value.

### HasDisableableCredentialTypes

`func (o *MemberRepresentation) HasDisableableCredentialTypes() bool`

HasDisableableCredentialTypes returns a boolean if a field has been set.

### GetRequiredActions

`func (o *MemberRepresentation) GetRequiredActions() []string`

GetRequiredActions returns the RequiredActions field if non-nil, zero value otherwise.

### GetRequiredActionsOk

`func (o *MemberRepresentation) GetRequiredActionsOk() (*[]string, bool)`

GetRequiredActionsOk returns a tuple with the RequiredActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredActions

`func (o *MemberRepresentation) SetRequiredActions(v []string)`

SetRequiredActions sets RequiredActions field to given value.

### HasRequiredActions

`func (o *MemberRepresentation) HasRequiredActions() bool`

HasRequiredActions returns a boolean if a field has been set.

### GetFederatedIdentities

`func (o *MemberRepresentation) GetFederatedIdentities() []FederatedIdentityRepresentation`

GetFederatedIdentities returns the FederatedIdentities field if non-nil, zero value otherwise.

### GetFederatedIdentitiesOk

`func (o *MemberRepresentation) GetFederatedIdentitiesOk() (*[]FederatedIdentityRepresentation, bool)`

GetFederatedIdentitiesOk returns a tuple with the FederatedIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedIdentities

`func (o *MemberRepresentation) SetFederatedIdentities(v []FederatedIdentityRepresentation)`

SetFederatedIdentities sets FederatedIdentities field to given value.

### HasFederatedIdentities

`func (o *MemberRepresentation) HasFederatedIdentities() bool`

HasFederatedIdentities returns a boolean if a field has been set.

### GetRealmRoles

`func (o *MemberRepresentation) GetRealmRoles() []string`

GetRealmRoles returns the RealmRoles field if non-nil, zero value otherwise.

### GetRealmRolesOk

`func (o *MemberRepresentation) GetRealmRolesOk() (*[]string, bool)`

GetRealmRolesOk returns a tuple with the RealmRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmRoles

`func (o *MemberRepresentation) SetRealmRoles(v []string)`

SetRealmRoles sets RealmRoles field to given value.

### HasRealmRoles

`func (o *MemberRepresentation) HasRealmRoles() bool`

HasRealmRoles returns a boolean if a field has been set.

### GetClientRoles

`func (o *MemberRepresentation) GetClientRoles() map[string][]string`

GetClientRoles returns the ClientRoles field if non-nil, zero value otherwise.

### GetClientRolesOk

`func (o *MemberRepresentation) GetClientRolesOk() (*map[string][]string, bool)`

GetClientRolesOk returns a tuple with the ClientRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRoles

`func (o *MemberRepresentation) SetClientRoles(v map[string][]string)`

SetClientRoles sets ClientRoles field to given value.

### HasClientRoles

`func (o *MemberRepresentation) HasClientRoles() bool`

HasClientRoles returns a boolean if a field has been set.

### GetClientConsents

`func (o *MemberRepresentation) GetClientConsents() []UserConsentRepresentation`

GetClientConsents returns the ClientConsents field if non-nil, zero value otherwise.

### GetClientConsentsOk

`func (o *MemberRepresentation) GetClientConsentsOk() (*[]UserConsentRepresentation, bool)`

GetClientConsentsOk returns a tuple with the ClientConsents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConsents

`func (o *MemberRepresentation) SetClientConsents(v []UserConsentRepresentation)`

SetClientConsents sets ClientConsents field to given value.

### HasClientConsents

`func (o *MemberRepresentation) HasClientConsents() bool`

HasClientConsents returns a boolean if a field has been set.

### GetNotBefore

`func (o *MemberRepresentation) GetNotBefore() int32`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *MemberRepresentation) GetNotBeforeOk() (*int32, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *MemberRepresentation) SetNotBefore(v int32)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *MemberRepresentation) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetApplicationRoles

`func (o *MemberRepresentation) GetApplicationRoles() map[string][]string`

GetApplicationRoles returns the ApplicationRoles field if non-nil, zero value otherwise.

### GetApplicationRolesOk

`func (o *MemberRepresentation) GetApplicationRolesOk() (*map[string][]string, bool)`

GetApplicationRolesOk returns a tuple with the ApplicationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationRoles

`func (o *MemberRepresentation) SetApplicationRoles(v map[string][]string)`

SetApplicationRoles sets ApplicationRoles field to given value.

### HasApplicationRoles

`func (o *MemberRepresentation) HasApplicationRoles() bool`

HasApplicationRoles returns a boolean if a field has been set.

### GetSocialLinks

`func (o *MemberRepresentation) GetSocialLinks() []SocialLinkRepresentation`

GetSocialLinks returns the SocialLinks field if non-nil, zero value otherwise.

### GetSocialLinksOk

`func (o *MemberRepresentation) GetSocialLinksOk() (*[]SocialLinkRepresentation, bool)`

GetSocialLinksOk returns a tuple with the SocialLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLinks

`func (o *MemberRepresentation) SetSocialLinks(v []SocialLinkRepresentation)`

SetSocialLinks sets SocialLinks field to given value.

### HasSocialLinks

`func (o *MemberRepresentation) HasSocialLinks() bool`

HasSocialLinks returns a boolean if a field has been set.

### GetGroups

`func (o *MemberRepresentation) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MemberRepresentation) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MemberRepresentation) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MemberRepresentation) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetAccess

`func (o *MemberRepresentation) GetAccess() map[string]bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MemberRepresentation) GetAccessOk() (*map[string]bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MemberRepresentation) SetAccess(v map[string]bool)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MemberRepresentation) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetMembershipType

`func (o *MemberRepresentation) GetMembershipType() MembershipType`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *MemberRepresentation) GetMembershipTypeOk() (*MembershipType, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *MemberRepresentation) SetMembershipType(v MembershipType)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *MemberRepresentation) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


