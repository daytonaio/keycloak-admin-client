# AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jti** | Pointer to **string** |  | [optional] 
**Exp** | Pointer to **int64** |  | [optional] 
**Nbf** | Pointer to **int64** |  | [optional] 
**Iat** | Pointer to **int64** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Typ** | Pointer to **string** |  | [optional] 
**Azp** | Pointer to **string** |  | [optional] 
**OtherClaims** | Pointer to **map[string]interface{}** |  | [optional] 
**Nonce** | Pointer to **string** |  | [optional] 
**AuthTime** | Pointer to **int64** |  | [optional] 
**Sid** | Pointer to **string** |  | [optional] 
**AtHash** | Pointer to **string** |  | [optional] 
**CHash** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**GivenName** | Pointer to **string** |  | [optional] 
**FamilyName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**PreferredUsername** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**Website** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Birthdate** | Pointer to **string** |  | [optional] 
**Zoneinfo** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PhoneNumberVerified** | Pointer to **bool** |  | [optional] 
**Address** | Pointer to [**AddressClaimSet**](AddressClaimSet.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** |  | [optional] 
**ClaimsLocales** | Pointer to **string** |  | [optional] 
**Acr** | Pointer to **string** |  | [optional] 
**SHash** | Pointer to **string** |  | [optional] 
**TrustedCerts** | Pointer to **[]string** |  | [optional] 
**AllowedOrigins** | Pointer to **[]string** |  | [optional] 
**RealmAccess** | Pointer to [**Access**](Access.md) |  | [optional] 
**ResourceAccess** | Pointer to [**map[string]Access**](Access.md) |  | [optional] 
**Authorization** | Pointer to [**Authorization**](Authorization.md) |  | [optional] 
**Cnf** | Pointer to [**Confirmation**](Confirmation.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessToken

`func NewAccessToken() *AccessToken`

NewAccessToken instantiates a new AccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenWithDefaults

`func NewAccessTokenWithDefaults() *AccessToken`

NewAccessTokenWithDefaults instantiates a new AccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJti

`func (o *AccessToken) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *AccessToken) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *AccessToken) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *AccessToken) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetExp

`func (o *AccessToken) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *AccessToken) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *AccessToken) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *AccessToken) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetNbf

`func (o *AccessToken) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *AccessToken) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *AccessToken) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *AccessToken) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetIat

`func (o *AccessToken) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *AccessToken) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *AccessToken) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *AccessToken) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *AccessToken) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *AccessToken) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *AccessToken) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *AccessToken) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetSub

`func (o *AccessToken) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AccessToken) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AccessToken) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AccessToken) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTyp

`func (o *AccessToken) GetTyp() string`

GetTyp returns the Typ field if non-nil, zero value otherwise.

### GetTypOk

`func (o *AccessToken) GetTypOk() (*string, bool)`

GetTypOk returns a tuple with the Typ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyp

`func (o *AccessToken) SetTyp(v string)`

SetTyp sets Typ field to given value.

### HasTyp

`func (o *AccessToken) HasTyp() bool`

HasTyp returns a boolean if a field has been set.

### GetAzp

`func (o *AccessToken) GetAzp() string`

GetAzp returns the Azp field if non-nil, zero value otherwise.

### GetAzpOk

`func (o *AccessToken) GetAzpOk() (*string, bool)`

GetAzpOk returns a tuple with the Azp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzp

`func (o *AccessToken) SetAzp(v string)`

SetAzp sets Azp field to given value.

### HasAzp

`func (o *AccessToken) HasAzp() bool`

HasAzp returns a boolean if a field has been set.

### GetOtherClaims

`func (o *AccessToken) GetOtherClaims() map[string]interface{}`

GetOtherClaims returns the OtherClaims field if non-nil, zero value otherwise.

### GetOtherClaimsOk

`func (o *AccessToken) GetOtherClaimsOk() (*map[string]interface{}, bool)`

GetOtherClaimsOk returns a tuple with the OtherClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherClaims

`func (o *AccessToken) SetOtherClaims(v map[string]interface{})`

SetOtherClaims sets OtherClaims field to given value.

### HasOtherClaims

`func (o *AccessToken) HasOtherClaims() bool`

HasOtherClaims returns a boolean if a field has been set.

### GetNonce

`func (o *AccessToken) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *AccessToken) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *AccessToken) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *AccessToken) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetAuthTime

`func (o *AccessToken) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *AccessToken) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *AccessToken) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *AccessToken) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetSid

`func (o *AccessToken) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *AccessToken) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *AccessToken) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *AccessToken) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetAtHash

`func (o *AccessToken) GetAtHash() string`

GetAtHash returns the AtHash field if non-nil, zero value otherwise.

### GetAtHashOk

`func (o *AccessToken) GetAtHashOk() (*string, bool)`

GetAtHashOk returns a tuple with the AtHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtHash

`func (o *AccessToken) SetAtHash(v string)`

SetAtHash sets AtHash field to given value.

### HasAtHash

`func (o *AccessToken) HasAtHash() bool`

HasAtHash returns a boolean if a field has been set.

### GetCHash

`func (o *AccessToken) GetCHash() string`

GetCHash returns the CHash field if non-nil, zero value otherwise.

### GetCHashOk

`func (o *AccessToken) GetCHashOk() (*string, bool)`

GetCHashOk returns a tuple with the CHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHash

`func (o *AccessToken) SetCHash(v string)`

SetCHash sets CHash field to given value.

### HasCHash

`func (o *AccessToken) HasCHash() bool`

HasCHash returns a boolean if a field has been set.

### GetName

`func (o *AccessToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGivenName

`func (o *AccessToken) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *AccessToken) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *AccessToken) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *AccessToken) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *AccessToken) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *AccessToken) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *AccessToken) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *AccessToken) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetMiddleName

`func (o *AccessToken) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *AccessToken) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *AccessToken) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *AccessToken) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNickname

`func (o *AccessToken) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *AccessToken) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *AccessToken) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *AccessToken) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPreferredUsername

`func (o *AccessToken) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *AccessToken) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *AccessToken) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *AccessToken) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### GetProfile

`func (o *AccessToken) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AccessToken) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AccessToken) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AccessToken) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetPicture

`func (o *AccessToken) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *AccessToken) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *AccessToken) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *AccessToken) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetWebsite

`func (o *AccessToken) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *AccessToken) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *AccessToken) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *AccessToken) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *AccessToken) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccessToken) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccessToken) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccessToken) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *AccessToken) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *AccessToken) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *AccessToken) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *AccessToken) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetGender

`func (o *AccessToken) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *AccessToken) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *AccessToken) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *AccessToken) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBirthdate

`func (o *AccessToken) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *AccessToken) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *AccessToken) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *AccessToken) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### GetZoneinfo

`func (o *AccessToken) GetZoneinfo() string`

GetZoneinfo returns the Zoneinfo field if non-nil, zero value otherwise.

### GetZoneinfoOk

`func (o *AccessToken) GetZoneinfoOk() (*string, bool)`

GetZoneinfoOk returns a tuple with the Zoneinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneinfo

`func (o *AccessToken) SetZoneinfo(v string)`

SetZoneinfo sets Zoneinfo field to given value.

### HasZoneinfo

`func (o *AccessToken) HasZoneinfo() bool`

HasZoneinfo returns a boolean if a field has been set.

### GetLocale

`func (o *AccessToken) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *AccessToken) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *AccessToken) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *AccessToken) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AccessToken) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AccessToken) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AccessToken) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AccessToken) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberVerified

`func (o *AccessToken) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *AccessToken) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *AccessToken) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *AccessToken) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### GetAddress

`func (o *AccessToken) GetAddress() AddressClaimSet`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccessToken) GetAddressOk() (*AddressClaimSet, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccessToken) SetAddress(v AddressClaimSet)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccessToken) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccessToken) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccessToken) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccessToken) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccessToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClaimsLocales

`func (o *AccessToken) GetClaimsLocales() string`

GetClaimsLocales returns the ClaimsLocales field if non-nil, zero value otherwise.

### GetClaimsLocalesOk

`func (o *AccessToken) GetClaimsLocalesOk() (*string, bool)`

GetClaimsLocalesOk returns a tuple with the ClaimsLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsLocales

`func (o *AccessToken) SetClaimsLocales(v string)`

SetClaimsLocales sets ClaimsLocales field to given value.

### HasClaimsLocales

`func (o *AccessToken) HasClaimsLocales() bool`

HasClaimsLocales returns a boolean if a field has been set.

### GetAcr

`func (o *AccessToken) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *AccessToken) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *AccessToken) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *AccessToken) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetSHash

`func (o *AccessToken) GetSHash() string`

GetSHash returns the SHash field if non-nil, zero value otherwise.

### GetSHashOk

`func (o *AccessToken) GetSHashOk() (*string, bool)`

GetSHashOk returns a tuple with the SHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHash

`func (o *AccessToken) SetSHash(v string)`

SetSHash sets SHash field to given value.

### HasSHash

`func (o *AccessToken) HasSHash() bool`

HasSHash returns a boolean if a field has been set.

### GetTrustedCerts

`func (o *AccessToken) GetTrustedCerts() []string`

GetTrustedCerts returns the TrustedCerts field if non-nil, zero value otherwise.

### GetTrustedCertsOk

`func (o *AccessToken) GetTrustedCertsOk() (*[]string, bool)`

GetTrustedCertsOk returns a tuple with the TrustedCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCerts

`func (o *AccessToken) SetTrustedCerts(v []string)`

SetTrustedCerts sets TrustedCerts field to given value.

### HasTrustedCerts

`func (o *AccessToken) HasTrustedCerts() bool`

HasTrustedCerts returns a boolean if a field has been set.

### GetAllowedOrigins

`func (o *AccessToken) GetAllowedOrigins() []string`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *AccessToken) GetAllowedOriginsOk() (*[]string, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *AccessToken) SetAllowedOrigins(v []string)`

SetAllowedOrigins sets AllowedOrigins field to given value.

### HasAllowedOrigins

`func (o *AccessToken) HasAllowedOrigins() bool`

HasAllowedOrigins returns a boolean if a field has been set.

### GetRealmAccess

`func (o *AccessToken) GetRealmAccess() Access`

GetRealmAccess returns the RealmAccess field if non-nil, zero value otherwise.

### GetRealmAccessOk

`func (o *AccessToken) GetRealmAccessOk() (*Access, bool)`

GetRealmAccessOk returns a tuple with the RealmAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmAccess

`func (o *AccessToken) SetRealmAccess(v Access)`

SetRealmAccess sets RealmAccess field to given value.

### HasRealmAccess

`func (o *AccessToken) HasRealmAccess() bool`

HasRealmAccess returns a boolean if a field has been set.

### GetResourceAccess

`func (o *AccessToken) GetResourceAccess() map[string]Access`

GetResourceAccess returns the ResourceAccess field if non-nil, zero value otherwise.

### GetResourceAccessOk

`func (o *AccessToken) GetResourceAccessOk() (*map[string]Access, bool)`

GetResourceAccessOk returns a tuple with the ResourceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAccess

`func (o *AccessToken) SetResourceAccess(v map[string]Access)`

SetResourceAccess sets ResourceAccess field to given value.

### HasResourceAccess

`func (o *AccessToken) HasResourceAccess() bool`

HasResourceAccess returns a boolean if a field has been set.

### GetAuthorization

`func (o *AccessToken) GetAuthorization() Authorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *AccessToken) GetAuthorizationOk() (*Authorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *AccessToken) SetAuthorization(v Authorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *AccessToken) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetCnf

`func (o *AccessToken) GetCnf() Confirmation`

GetCnf returns the Cnf field if non-nil, zero value otherwise.

### GetCnfOk

`func (o *AccessToken) GetCnfOk() (*Confirmation, bool)`

GetCnfOk returns a tuple with the Cnf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnf

`func (o *AccessToken) SetCnf(v Confirmation)`

SetCnf sets Cnf field to given value.

### HasCnf

`func (o *AccessToken) HasCnf() bool`

HasCnf returns a boolean if a field has been set.

### GetScope

`func (o *AccessToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccessToken) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


