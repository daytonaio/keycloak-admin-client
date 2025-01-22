# ClientPolicyConditionRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClientPolicyConditionRepresentation

`func NewClientPolicyConditionRepresentation() *ClientPolicyConditionRepresentation`

NewClientPolicyConditionRepresentation instantiates a new ClientPolicyConditionRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientPolicyConditionRepresentationWithDefaults

`func NewClientPolicyConditionRepresentationWithDefaults() *ClientPolicyConditionRepresentation`

NewClientPolicyConditionRepresentationWithDefaults instantiates a new ClientPolicyConditionRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *ClientPolicyConditionRepresentation) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ClientPolicyConditionRepresentation) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ClientPolicyConditionRepresentation) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ClientPolicyConditionRepresentation) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConfiguration

`func (o *ClientPolicyConditionRepresentation) GetConfiguration() []string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ClientPolicyConditionRepresentation) GetConfigurationOk() (*[]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ClientPolicyConditionRepresentation) SetConfiguration(v []string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ClientPolicyConditionRepresentation) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


