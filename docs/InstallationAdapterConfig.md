# InstallationAdapterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realm** | Pointer to **string** |  | [optional] 
**RealmPublicKey** | Pointer to **string** |  | [optional] 
**AuthServerUrl** | Pointer to **string** |  | [optional] 
**SslRequired** | Pointer to **string** |  | [optional] 
**BearerOnly** | Pointer to **bool** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**PublicClient** | Pointer to **bool** |  | [optional] 
**VerifyTokenAudience** | Pointer to **bool** |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**UseResourceRoleMappings** | Pointer to **bool** |  | [optional] 
**ConfidentialPort** | Pointer to **int32** |  | [optional] 
**PolicyEnforcer** | Pointer to [**PolicyEnforcerConfig**](PolicyEnforcerConfig.md) |  | [optional] 

## Methods

### NewInstallationAdapterConfig

`func NewInstallationAdapterConfig() *InstallationAdapterConfig`

NewInstallationAdapterConfig instantiates a new InstallationAdapterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationAdapterConfigWithDefaults

`func NewInstallationAdapterConfigWithDefaults() *InstallationAdapterConfig`

NewInstallationAdapterConfigWithDefaults instantiates a new InstallationAdapterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealm

`func (o *InstallationAdapterConfig) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *InstallationAdapterConfig) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *InstallationAdapterConfig) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *InstallationAdapterConfig) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetRealmPublicKey

`func (o *InstallationAdapterConfig) GetRealmPublicKey() string`

GetRealmPublicKey returns the RealmPublicKey field if non-nil, zero value otherwise.

### GetRealmPublicKeyOk

`func (o *InstallationAdapterConfig) GetRealmPublicKeyOk() (*string, bool)`

GetRealmPublicKeyOk returns a tuple with the RealmPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmPublicKey

`func (o *InstallationAdapterConfig) SetRealmPublicKey(v string)`

SetRealmPublicKey sets RealmPublicKey field to given value.

### HasRealmPublicKey

`func (o *InstallationAdapterConfig) HasRealmPublicKey() bool`

HasRealmPublicKey returns a boolean if a field has been set.

### GetAuthServerUrl

`func (o *InstallationAdapterConfig) GetAuthServerUrl() string`

GetAuthServerUrl returns the AuthServerUrl field if non-nil, zero value otherwise.

### GetAuthServerUrlOk

`func (o *InstallationAdapterConfig) GetAuthServerUrlOk() (*string, bool)`

GetAuthServerUrlOk returns a tuple with the AuthServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServerUrl

`func (o *InstallationAdapterConfig) SetAuthServerUrl(v string)`

SetAuthServerUrl sets AuthServerUrl field to given value.

### HasAuthServerUrl

`func (o *InstallationAdapterConfig) HasAuthServerUrl() bool`

HasAuthServerUrl returns a boolean if a field has been set.

### GetSslRequired

`func (o *InstallationAdapterConfig) GetSslRequired() string`

GetSslRequired returns the SslRequired field if non-nil, zero value otherwise.

### GetSslRequiredOk

`func (o *InstallationAdapterConfig) GetSslRequiredOk() (*string, bool)`

GetSslRequiredOk returns a tuple with the SslRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRequired

`func (o *InstallationAdapterConfig) SetSslRequired(v string)`

SetSslRequired sets SslRequired field to given value.

### HasSslRequired

`func (o *InstallationAdapterConfig) HasSslRequired() bool`

HasSslRequired returns a boolean if a field has been set.

### GetBearerOnly

`func (o *InstallationAdapterConfig) GetBearerOnly() bool`

GetBearerOnly returns the BearerOnly field if non-nil, zero value otherwise.

### GetBearerOnlyOk

`func (o *InstallationAdapterConfig) GetBearerOnlyOk() (*bool, bool)`

GetBearerOnlyOk returns a tuple with the BearerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerOnly

`func (o *InstallationAdapterConfig) SetBearerOnly(v bool)`

SetBearerOnly sets BearerOnly field to given value.

### HasBearerOnly

`func (o *InstallationAdapterConfig) HasBearerOnly() bool`

HasBearerOnly returns a boolean if a field has been set.

### GetResource

`func (o *InstallationAdapterConfig) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *InstallationAdapterConfig) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *InstallationAdapterConfig) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *InstallationAdapterConfig) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPublicClient

`func (o *InstallationAdapterConfig) GetPublicClient() bool`

GetPublicClient returns the PublicClient field if non-nil, zero value otherwise.

### GetPublicClientOk

`func (o *InstallationAdapterConfig) GetPublicClientOk() (*bool, bool)`

GetPublicClientOk returns a tuple with the PublicClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicClient

`func (o *InstallationAdapterConfig) SetPublicClient(v bool)`

SetPublicClient sets PublicClient field to given value.

### HasPublicClient

`func (o *InstallationAdapterConfig) HasPublicClient() bool`

HasPublicClient returns a boolean if a field has been set.

### GetVerifyTokenAudience

`func (o *InstallationAdapterConfig) GetVerifyTokenAudience() bool`

GetVerifyTokenAudience returns the VerifyTokenAudience field if non-nil, zero value otherwise.

### GetVerifyTokenAudienceOk

`func (o *InstallationAdapterConfig) GetVerifyTokenAudienceOk() (*bool, bool)`

GetVerifyTokenAudienceOk returns a tuple with the VerifyTokenAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTokenAudience

`func (o *InstallationAdapterConfig) SetVerifyTokenAudience(v bool)`

SetVerifyTokenAudience sets VerifyTokenAudience field to given value.

### HasVerifyTokenAudience

`func (o *InstallationAdapterConfig) HasVerifyTokenAudience() bool`

HasVerifyTokenAudience returns a boolean if a field has been set.

### GetCredentials

`func (o *InstallationAdapterConfig) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *InstallationAdapterConfig) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *InstallationAdapterConfig) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *InstallationAdapterConfig) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetUseResourceRoleMappings

`func (o *InstallationAdapterConfig) GetUseResourceRoleMappings() bool`

GetUseResourceRoleMappings returns the UseResourceRoleMappings field if non-nil, zero value otherwise.

### GetUseResourceRoleMappingsOk

`func (o *InstallationAdapterConfig) GetUseResourceRoleMappingsOk() (*bool, bool)`

GetUseResourceRoleMappingsOk returns a tuple with the UseResourceRoleMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseResourceRoleMappings

`func (o *InstallationAdapterConfig) SetUseResourceRoleMappings(v bool)`

SetUseResourceRoleMappings sets UseResourceRoleMappings field to given value.

### HasUseResourceRoleMappings

`func (o *InstallationAdapterConfig) HasUseResourceRoleMappings() bool`

HasUseResourceRoleMappings returns a boolean if a field has been set.

### GetConfidentialPort

`func (o *InstallationAdapterConfig) GetConfidentialPort() int32`

GetConfidentialPort returns the ConfidentialPort field if non-nil, zero value otherwise.

### GetConfidentialPortOk

`func (o *InstallationAdapterConfig) GetConfidentialPortOk() (*int32, bool)`

GetConfidentialPortOk returns a tuple with the ConfidentialPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialPort

`func (o *InstallationAdapterConfig) SetConfidentialPort(v int32)`

SetConfidentialPort sets ConfidentialPort field to given value.

### HasConfidentialPort

`func (o *InstallationAdapterConfig) HasConfidentialPort() bool`

HasConfidentialPort returns a boolean if a field has been set.

### GetPolicyEnforcer

`func (o *InstallationAdapterConfig) GetPolicyEnforcer() PolicyEnforcerConfig`

GetPolicyEnforcer returns the PolicyEnforcer field if non-nil, zero value otherwise.

### GetPolicyEnforcerOk

`func (o *InstallationAdapterConfig) GetPolicyEnforcerOk() (*PolicyEnforcerConfig, bool)`

GetPolicyEnforcerOk returns a tuple with the PolicyEnforcer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEnforcer

`func (o *InstallationAdapterConfig) SetPolicyEnforcer(v PolicyEnforcerConfig)`

SetPolicyEnforcer sets PolicyEnforcer field to given value.

### HasPolicyEnforcer

`func (o *InstallationAdapterConfig) HasPolicyEnforcer() bool`

HasPolicyEnforcer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


