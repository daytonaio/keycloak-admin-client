# PathCacheConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxEntries** | Pointer to **int32** |  | [optional] 
**Lifespan** | Pointer to **int64** |  | [optional] 

## Methods

### NewPathCacheConfig

`func NewPathCacheConfig() *PathCacheConfig`

NewPathCacheConfig instantiates a new PathCacheConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathCacheConfigWithDefaults

`func NewPathCacheConfigWithDefaults() *PathCacheConfig`

NewPathCacheConfigWithDefaults instantiates a new PathCacheConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxEntries

`func (o *PathCacheConfig) GetMaxEntries() int32`

GetMaxEntries returns the MaxEntries field if non-nil, zero value otherwise.

### GetMaxEntriesOk

`func (o *PathCacheConfig) GetMaxEntriesOk() (*int32, bool)`

GetMaxEntriesOk returns a tuple with the MaxEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntries

`func (o *PathCacheConfig) SetMaxEntries(v int32)`

SetMaxEntries sets MaxEntries field to given value.

### HasMaxEntries

`func (o *PathCacheConfig) HasMaxEntries() bool`

HasMaxEntries returns a boolean if a field has been set.

### GetLifespan

`func (o *PathCacheConfig) GetLifespan() int64`

GetLifespan returns the Lifespan field if non-nil, zero value otherwise.

### GetLifespanOk

`func (o *PathCacheConfig) GetLifespanOk() (*int64, bool)`

GetLifespanOk returns a tuple with the Lifespan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifespan

`func (o *PathCacheConfig) SetLifespan(v int64)`

SetLifespan sets Lifespan field to given value.

### HasLifespan

`func (o *PathCacheConfig) HasLifespan() bool`

HasLifespan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


