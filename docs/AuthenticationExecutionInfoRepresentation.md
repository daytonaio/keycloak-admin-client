# AuthenticationExecutionInfoRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Requirement** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequirementChoices** | Pointer to **[]string** |  | [optional] 
**Configurable** | Pointer to **bool** |  | [optional] 
**AuthenticationFlow** | Pointer to **bool** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**AuthenticationConfig** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthenticationExecutionInfoRepresentation

`func NewAuthenticationExecutionInfoRepresentation() *AuthenticationExecutionInfoRepresentation`

NewAuthenticationExecutionInfoRepresentation instantiates a new AuthenticationExecutionInfoRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationExecutionInfoRepresentationWithDefaults

`func NewAuthenticationExecutionInfoRepresentationWithDefaults() *AuthenticationExecutionInfoRepresentation`

NewAuthenticationExecutionInfoRepresentationWithDefaults instantiates a new AuthenticationExecutionInfoRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationExecutionInfoRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationExecutionInfoRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationExecutionInfoRepresentation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationExecutionInfoRepresentation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequirement

`func (o *AuthenticationExecutionInfoRepresentation) GetRequirement() string`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *AuthenticationExecutionInfoRepresentation) GetRequirementOk() (*string, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *AuthenticationExecutionInfoRepresentation) SetRequirement(v string)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *AuthenticationExecutionInfoRepresentation) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.

### GetDisplayName

`func (o *AuthenticationExecutionInfoRepresentation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthenticationExecutionInfoRepresentation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthenticationExecutionInfoRepresentation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthenticationExecutionInfoRepresentation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAlias

`func (o *AuthenticationExecutionInfoRepresentation) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AuthenticationExecutionInfoRepresentation) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AuthenticationExecutionInfoRepresentation) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AuthenticationExecutionInfoRepresentation) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationExecutionInfoRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationExecutionInfoRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationExecutionInfoRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationExecutionInfoRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequirementChoices

`func (o *AuthenticationExecutionInfoRepresentation) GetRequirementChoices() []string`

GetRequirementChoices returns the RequirementChoices field if non-nil, zero value otherwise.

### GetRequirementChoicesOk

`func (o *AuthenticationExecutionInfoRepresentation) GetRequirementChoicesOk() (*[]string, bool)`

GetRequirementChoicesOk returns a tuple with the RequirementChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementChoices

`func (o *AuthenticationExecutionInfoRepresentation) SetRequirementChoices(v []string)`

SetRequirementChoices sets RequirementChoices field to given value.

### HasRequirementChoices

`func (o *AuthenticationExecutionInfoRepresentation) HasRequirementChoices() bool`

HasRequirementChoices returns a boolean if a field has been set.

### GetConfigurable

`func (o *AuthenticationExecutionInfoRepresentation) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *AuthenticationExecutionInfoRepresentation) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *AuthenticationExecutionInfoRepresentation) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *AuthenticationExecutionInfoRepresentation) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.

### GetAuthenticationFlow

`func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationFlow() bool`

GetAuthenticationFlow returns the AuthenticationFlow field if non-nil, zero value otherwise.

### GetAuthenticationFlowOk

`func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationFlowOk() (*bool, bool)`

GetAuthenticationFlowOk returns a tuple with the AuthenticationFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationFlow

`func (o *AuthenticationExecutionInfoRepresentation) SetAuthenticationFlow(v bool)`

SetAuthenticationFlow sets AuthenticationFlow field to given value.

### HasAuthenticationFlow

`func (o *AuthenticationExecutionInfoRepresentation) HasAuthenticationFlow() bool`

HasAuthenticationFlow returns a boolean if a field has been set.

### GetProviderId

`func (o *AuthenticationExecutionInfoRepresentation) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *AuthenticationExecutionInfoRepresentation) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *AuthenticationExecutionInfoRepresentation) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *AuthenticationExecutionInfoRepresentation) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetAuthenticationConfig

`func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationConfig() string`

GetAuthenticationConfig returns the AuthenticationConfig field if non-nil, zero value otherwise.

### GetAuthenticationConfigOk

`func (o *AuthenticationExecutionInfoRepresentation) GetAuthenticationConfigOk() (*string, bool)`

GetAuthenticationConfigOk returns a tuple with the AuthenticationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationConfig

`func (o *AuthenticationExecutionInfoRepresentation) SetAuthenticationConfig(v string)`

SetAuthenticationConfig sets AuthenticationConfig field to given value.

### HasAuthenticationConfig

`func (o *AuthenticationExecutionInfoRepresentation) HasAuthenticationConfig() bool`

HasAuthenticationConfig returns a boolean if a field has been set.

### GetFlowId

`func (o *AuthenticationExecutionInfoRepresentation) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *AuthenticationExecutionInfoRepresentation) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *AuthenticationExecutionInfoRepresentation) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *AuthenticationExecutionInfoRepresentation) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetLevel

`func (o *AuthenticationExecutionInfoRepresentation) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuthenticationExecutionInfoRepresentation) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuthenticationExecutionInfoRepresentation) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AuthenticationExecutionInfoRepresentation) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetIndex

`func (o *AuthenticationExecutionInfoRepresentation) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AuthenticationExecutionInfoRepresentation) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AuthenticationExecutionInfoRepresentation) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *AuthenticationExecutionInfoRepresentation) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPriority

`func (o *AuthenticationExecutionInfoRepresentation) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AuthenticationExecutionInfoRepresentation) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AuthenticationExecutionInfoRepresentation) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AuthenticationExecutionInfoRepresentation) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


