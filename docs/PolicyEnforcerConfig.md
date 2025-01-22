# PolicyEnforcerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforcementMode** | Pointer to [**EnforcementMode**](EnforcementMode.md) |  | [optional] 
**Paths** | Pointer to [**[]PathConfig**](PathConfig.md) |  | [optional] 
**PathCache** | Pointer to [**PathCacheConfig**](PathCacheConfig.md) |  | [optional] 
**LazyLoadPaths** | Pointer to **bool** |  | [optional] 
**OnDenyRedirectTo** | Pointer to **string** |  | [optional] 
**UserManagedAccess** | Pointer to **map[string]interface{}** |  | [optional] 
**ClaimInformationPoint** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**HttpMethodAsScope** | Pointer to **bool** |  | [optional] 
**Realm** | Pointer to **string** |  | [optional] 
**AuthServerUrl** | Pointer to **string** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyEnforcerConfig

`func NewPolicyEnforcerConfig() *PolicyEnforcerConfig`

NewPolicyEnforcerConfig instantiates a new PolicyEnforcerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEnforcerConfigWithDefaults

`func NewPolicyEnforcerConfigWithDefaults() *PolicyEnforcerConfig`

NewPolicyEnforcerConfigWithDefaults instantiates a new PolicyEnforcerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforcementMode

`func (o *PolicyEnforcerConfig) GetEnforcementMode() EnforcementMode`

GetEnforcementMode returns the EnforcementMode field if non-nil, zero value otherwise.

### GetEnforcementModeOk

`func (o *PolicyEnforcerConfig) GetEnforcementModeOk() (*EnforcementMode, bool)`

GetEnforcementModeOk returns a tuple with the EnforcementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementMode

`func (o *PolicyEnforcerConfig) SetEnforcementMode(v EnforcementMode)`

SetEnforcementMode sets EnforcementMode field to given value.

### HasEnforcementMode

`func (o *PolicyEnforcerConfig) HasEnforcementMode() bool`

HasEnforcementMode returns a boolean if a field has been set.

### GetPaths

`func (o *PolicyEnforcerConfig) GetPaths() []PathConfig`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *PolicyEnforcerConfig) GetPathsOk() (*[]PathConfig, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *PolicyEnforcerConfig) SetPaths(v []PathConfig)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *PolicyEnforcerConfig) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetPathCache

`func (o *PolicyEnforcerConfig) GetPathCache() PathCacheConfig`

GetPathCache returns the PathCache field if non-nil, zero value otherwise.

### GetPathCacheOk

`func (o *PolicyEnforcerConfig) GetPathCacheOk() (*PathCacheConfig, bool)`

GetPathCacheOk returns a tuple with the PathCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathCache

`func (o *PolicyEnforcerConfig) SetPathCache(v PathCacheConfig)`

SetPathCache sets PathCache field to given value.

### HasPathCache

`func (o *PolicyEnforcerConfig) HasPathCache() bool`

HasPathCache returns a boolean if a field has been set.

### GetLazyLoadPaths

`func (o *PolicyEnforcerConfig) GetLazyLoadPaths() bool`

GetLazyLoadPaths returns the LazyLoadPaths field if non-nil, zero value otherwise.

### GetLazyLoadPathsOk

`func (o *PolicyEnforcerConfig) GetLazyLoadPathsOk() (*bool, bool)`

GetLazyLoadPathsOk returns a tuple with the LazyLoadPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLazyLoadPaths

`func (o *PolicyEnforcerConfig) SetLazyLoadPaths(v bool)`

SetLazyLoadPaths sets LazyLoadPaths field to given value.

### HasLazyLoadPaths

`func (o *PolicyEnforcerConfig) HasLazyLoadPaths() bool`

HasLazyLoadPaths returns a boolean if a field has been set.

### GetOnDenyRedirectTo

`func (o *PolicyEnforcerConfig) GetOnDenyRedirectTo() string`

GetOnDenyRedirectTo returns the OnDenyRedirectTo field if non-nil, zero value otherwise.

### GetOnDenyRedirectToOk

`func (o *PolicyEnforcerConfig) GetOnDenyRedirectToOk() (*string, bool)`

GetOnDenyRedirectToOk returns a tuple with the OnDenyRedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDenyRedirectTo

`func (o *PolicyEnforcerConfig) SetOnDenyRedirectTo(v string)`

SetOnDenyRedirectTo sets OnDenyRedirectTo field to given value.

### HasOnDenyRedirectTo

`func (o *PolicyEnforcerConfig) HasOnDenyRedirectTo() bool`

HasOnDenyRedirectTo returns a boolean if a field has been set.

### GetUserManagedAccess

`func (o *PolicyEnforcerConfig) GetUserManagedAccess() map[string]interface{}`

GetUserManagedAccess returns the UserManagedAccess field if non-nil, zero value otherwise.

### GetUserManagedAccessOk

`func (o *PolicyEnforcerConfig) GetUserManagedAccessOk() (*map[string]interface{}, bool)`

GetUserManagedAccessOk returns a tuple with the UserManagedAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserManagedAccess

`func (o *PolicyEnforcerConfig) SetUserManagedAccess(v map[string]interface{})`

SetUserManagedAccess sets UserManagedAccess field to given value.

### HasUserManagedAccess

`func (o *PolicyEnforcerConfig) HasUserManagedAccess() bool`

HasUserManagedAccess returns a boolean if a field has been set.

### GetClaimInformationPoint

`func (o *PolicyEnforcerConfig) GetClaimInformationPoint() map[string]map[string]interface{}`

GetClaimInformationPoint returns the ClaimInformationPoint field if non-nil, zero value otherwise.

### GetClaimInformationPointOk

`func (o *PolicyEnforcerConfig) GetClaimInformationPointOk() (*map[string]map[string]interface{}, bool)`

GetClaimInformationPointOk returns a tuple with the ClaimInformationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimInformationPoint

`func (o *PolicyEnforcerConfig) SetClaimInformationPoint(v map[string]map[string]interface{})`

SetClaimInformationPoint sets ClaimInformationPoint field to given value.

### HasClaimInformationPoint

`func (o *PolicyEnforcerConfig) HasClaimInformationPoint() bool`

HasClaimInformationPoint returns a boolean if a field has been set.

### GetHttpMethodAsScope

`func (o *PolicyEnforcerConfig) GetHttpMethodAsScope() bool`

GetHttpMethodAsScope returns the HttpMethodAsScope field if non-nil, zero value otherwise.

### GetHttpMethodAsScopeOk

`func (o *PolicyEnforcerConfig) GetHttpMethodAsScopeOk() (*bool, bool)`

GetHttpMethodAsScopeOk returns a tuple with the HttpMethodAsScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethodAsScope

`func (o *PolicyEnforcerConfig) SetHttpMethodAsScope(v bool)`

SetHttpMethodAsScope sets HttpMethodAsScope field to given value.

### HasHttpMethodAsScope

`func (o *PolicyEnforcerConfig) HasHttpMethodAsScope() bool`

HasHttpMethodAsScope returns a boolean if a field has been set.

### GetRealm

`func (o *PolicyEnforcerConfig) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *PolicyEnforcerConfig) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *PolicyEnforcerConfig) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *PolicyEnforcerConfig) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetAuthServerUrl

`func (o *PolicyEnforcerConfig) GetAuthServerUrl() string`

GetAuthServerUrl returns the AuthServerUrl field if non-nil, zero value otherwise.

### GetAuthServerUrlOk

`func (o *PolicyEnforcerConfig) GetAuthServerUrlOk() (*string, bool)`

GetAuthServerUrlOk returns a tuple with the AuthServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerUrl

`func (o *PolicyEnforcerConfig) SetAuthServerUrl(v string)`

SetAuthServerUrl sets AuthServerUrl field to given value.

### HasAuthServerUrl

`func (o *PolicyEnforcerConfig) HasAuthServerUrl() bool`

HasAuthServerUrl returns a boolean if a field has been set.

### GetCredentials

`func (o *PolicyEnforcerConfig) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PolicyEnforcerConfig) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PolicyEnforcerConfig) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PolicyEnforcerConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetResource

`func (o *PolicyEnforcerConfig) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyEnforcerConfig) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyEnforcerConfig) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PolicyEnforcerConfig) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


