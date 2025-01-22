# IDToken

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

## Methods

### NewIDToken

`func NewIDToken() *IDToken`

NewIDToken instantiates a new IDToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDTokenWithDefaults

`func NewIDTokenWithDefaults() *IDToken`

NewIDTokenWithDefaults instantiates a new IDToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJti

`func (o *IDToken) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *IDToken) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *IDToken) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *IDToken) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetExp

`func (o *IDToken) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *IDToken) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *IDToken) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *IDToken) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetNbf

`func (o *IDToken) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *IDToken) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *IDToken) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *IDToken) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetIat

`func (o *IDToken) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *IDToken) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *IDToken) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *IDToken) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *IDToken) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *IDToken) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *IDToken) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *IDToken) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetSub

`func (o *IDToken) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IDToken) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IDToken) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IDToken) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetTyp

`func (o *IDToken) GetTyp() string`

GetTyp returns the Typ field if non-nil, zero value otherwise.

### GetTypOk

`func (o *IDToken) GetTypOk() (*string, bool)`

GetTypOk returns a tuple with the Typ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyp

`func (o *IDToken) SetTyp(v string)`

SetTyp sets Typ field to given value.

### HasTyp

`func (o *IDToken) HasTyp() bool`

HasTyp returns a boolean if a field has been set.

### GetAzp

`func (o *IDToken) GetAzp() string`

GetAzp returns the Azp field if non-nil, zero value otherwise.

### GetAzpOk

`func (o *IDToken) GetAzpOk() (*string, bool)`

GetAzpOk returns a tuple with the Azp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzp

`func (o *IDToken) SetAzp(v string)`

SetAzp sets Azp field to given value.

### HasAzp

`func (o *IDToken) HasAzp() bool`

HasAzp returns a boolean if a field has been set.

### GetOtherClaims

`func (o *IDToken) GetOtherClaims() map[string]interface{}`

GetOtherClaims returns the OtherClaims field if non-nil, zero value otherwise.

### GetOtherClaimsOk

`func (o *IDToken) GetOtherClaimsOk() (*map[string]interface{}, bool)`

GetOtherClaimsOk returns a tuple with the OtherClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherClaims

`func (o *IDToken) SetOtherClaims(v map[string]interface{})`

SetOtherClaims sets OtherClaims field to given value.

### HasOtherClaims

`func (o *IDToken) HasOtherClaims() bool`

HasOtherClaims returns a boolean if a field has been set.

### GetNonce

`func (o *IDToken) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *IDToken) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *IDToken) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *IDToken) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetAuthTime

`func (o *IDToken) GetAuthTime() int64`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *IDToken) GetAuthTimeOk() (*int64, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *IDToken) SetAuthTime(v int64)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *IDToken) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetSid

`func (o *IDToken) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IDToken) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IDToken) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IDToken) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetAtHash

`func (o *IDToken) GetAtHash() string`

GetAtHash returns the AtHash field if non-nil, zero value otherwise.

### GetAtHashOk

`func (o *IDToken) GetAtHashOk() (*string, bool)`

GetAtHashOk returns a tuple with the AtHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtHash

`func (o *IDToken) SetAtHash(v string)`

SetAtHash sets AtHash field to given value.

### HasAtHash

`func (o *IDToken) HasAtHash() bool`

HasAtHash returns a boolean if a field has been set.

### GetCHash

`func (o *IDToken) GetCHash() string`

GetCHash returns the CHash field if non-nil, zero value otherwise.

### GetCHashOk

`func (o *IDToken) GetCHashOk() (*string, bool)`

GetCHashOk returns a tuple with the CHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCHash

`func (o *IDToken) SetCHash(v string)`

SetCHash sets CHash field to given value.

### HasCHash

`func (o *IDToken) HasCHash() bool`

HasCHash returns a boolean if a field has been set.

### GetName

`func (o *IDToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IDToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IDToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IDToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGivenName

`func (o *IDToken) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *IDToken) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *IDToken) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.

### HasGivenName

`func (o *IDToken) HasGivenName() bool`

HasGivenName returns a boolean if a field has been set.

### GetFamilyName

`func (o *IDToken) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *IDToken) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *IDToken) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *IDToken) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetMiddleName

`func (o *IDToken) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *IDToken) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *IDToken) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *IDToken) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNickname

`func (o *IDToken) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *IDToken) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *IDToken) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *IDToken) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPreferredUsername

`func (o *IDToken) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *IDToken) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *IDToken) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *IDToken) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### GetProfile

`func (o *IDToken) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *IDToken) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *IDToken) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *IDToken) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetPicture

`func (o *IDToken) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *IDToken) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *IDToken) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *IDToken) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetWebsite

`func (o *IDToken) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *IDToken) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *IDToken) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *IDToken) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetEmail

`func (o *IDToken) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IDToken) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IDToken) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IDToken) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *IDToken) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *IDToken) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *IDToken) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *IDToken) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetGender

`func (o *IDToken) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *IDToken) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *IDToken) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *IDToken) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetBirthdate

`func (o *IDToken) GetBirthdate() string`

GetBirthdate returns the Birthdate field if non-nil, zero value otherwise.

### GetBirthdateOk

`func (o *IDToken) GetBirthdateOk() (*string, bool)`

GetBirthdateOk returns a tuple with the Birthdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthdate

`func (o *IDToken) SetBirthdate(v string)`

SetBirthdate sets Birthdate field to given value.

### HasBirthdate

`func (o *IDToken) HasBirthdate() bool`

HasBirthdate returns a boolean if a field has been set.

### GetZoneinfo

`func (o *IDToken) GetZoneinfo() string`

GetZoneinfo returns the Zoneinfo field if non-nil, zero value otherwise.

### GetZoneinfoOk

`func (o *IDToken) GetZoneinfoOk() (*string, bool)`

GetZoneinfoOk returns a tuple with the Zoneinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneinfo

`func (o *IDToken) SetZoneinfo(v string)`

SetZoneinfo sets Zoneinfo field to given value.

### HasZoneinfo

`func (o *IDToken) HasZoneinfo() bool`

HasZoneinfo returns a boolean if a field has been set.

### GetLocale

`func (o *IDToken) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *IDToken) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *IDToken) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *IDToken) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *IDToken) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IDToken) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IDToken) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IDToken) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberVerified

`func (o *IDToken) GetPhoneNumberVerified() bool`

GetPhoneNumberVerified returns the PhoneNumberVerified field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedOk

`func (o *IDToken) GetPhoneNumberVerifiedOk() (*bool, bool)`

GetPhoneNumberVerifiedOk returns a tuple with the PhoneNumberVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerified

`func (o *IDToken) SetPhoneNumberVerified(v bool)`

SetPhoneNumberVerified sets PhoneNumberVerified field to given value.

### HasPhoneNumberVerified

`func (o *IDToken) HasPhoneNumberVerified() bool`

HasPhoneNumberVerified returns a boolean if a field has been set.

### GetAddress

`func (o *IDToken) GetAddress() AddressClaimSet`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IDToken) GetAddressOk() (*AddressClaimSet, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IDToken) SetAddress(v AddressClaimSet)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IDToken) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IDToken) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IDToken) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IDToken) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IDToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetClaimsLocales

`func (o *IDToken) GetClaimsLocales() string`

GetClaimsLocales returns the ClaimsLocales field if non-nil, zero value otherwise.

### GetClaimsLocalesOk

`func (o *IDToken) GetClaimsLocalesOk() (*string, bool)`

GetClaimsLocalesOk returns a tuple with the ClaimsLocales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsLocales

`func (o *IDToken) SetClaimsLocales(v string)`

SetClaimsLocales sets ClaimsLocales field to given value.

### HasClaimsLocales

`func (o *IDToken) HasClaimsLocales() bool`

HasClaimsLocales returns a boolean if a field has been set.

### GetAcr

`func (o *IDToken) GetAcr() string`

GetAcr returns the Acr field if non-nil, zero value otherwise.

### GetAcrOk

`func (o *IDToken) GetAcrOk() (*string, bool)`

GetAcrOk returns a tuple with the Acr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcr

`func (o *IDToken) SetAcr(v string)`

SetAcr sets Acr field to given value.

### HasAcr

`func (o *IDToken) HasAcr() bool`

HasAcr returns a boolean if a field has been set.

### GetSHash

`func (o *IDToken) GetSHash() string`

GetSHash returns the SHash field if non-nil, zero value otherwise.

### GetSHashOk

`func (o *IDToken) GetSHashOk() (*string, bool)`

GetSHashOk returns a tuple with the SHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSHash

`func (o *IDToken) SetSHash(v string)`

SetSHash sets SHash field to given value.

### HasSHash

`func (o *IDToken) HasSHash() bool`

HasSHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


