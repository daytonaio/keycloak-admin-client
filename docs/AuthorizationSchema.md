# AuthorizationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceTypes** | Pointer to [**map[string]ResourceType**](ResourceType.md) |  | [optional] 

## Methods

### NewAuthorizationSchema

`func NewAuthorizationSchema() *AuthorizationSchema`

NewAuthorizationSchema instantiates a new AuthorizationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationSchemaWithDefaults

`func NewAuthorizationSchemaWithDefaults() *AuthorizationSchema`

NewAuthorizationSchemaWithDefaults instantiates a new AuthorizationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceTypes

`func (o *AuthorizationSchema) GetResourceTypes() map[string]ResourceType`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *AuthorizationSchema) GetResourceTypesOk() (*map[string]ResourceType, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *AuthorizationSchema) SetResourceTypes(v map[string]ResourceType)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *AuthorizationSchema) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


