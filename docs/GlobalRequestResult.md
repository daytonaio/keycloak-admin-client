# GlobalRequestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessRequests** | Pointer to **[]string** |  | [optional] 
**FailedRequests** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGlobalRequestResult

`func NewGlobalRequestResult() *GlobalRequestResult`

NewGlobalRequestResult instantiates a new GlobalRequestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalRequestResultWithDefaults

`func NewGlobalRequestResultWithDefaults() *GlobalRequestResult`

NewGlobalRequestResultWithDefaults instantiates a new GlobalRequestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessRequests

`func (o *GlobalRequestResult) GetSuccessRequests() []string`

GetSuccessRequests returns the SuccessRequests field if non-nil, zero value otherwise.

### GetSuccessRequestsOk

`func (o *GlobalRequestResult) GetSuccessRequestsOk() (*[]string, bool)`

GetSuccessRequestsOk returns a tuple with the SuccessRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRequests

`func (o *GlobalRequestResult) SetSuccessRequests(v []string)`

SetSuccessRequests sets SuccessRequests field to given value.

### HasSuccessRequests

`func (o *GlobalRequestResult) HasSuccessRequests() bool`

HasSuccessRequests returns a boolean if a field has been set.

### GetFailedRequests

`func (o *GlobalRequestResult) GetFailedRequests() []string`

GetFailedRequests returns the FailedRequests field if non-nil, zero value otherwise.

### GetFailedRequestsOk

`func (o *GlobalRequestResult) GetFailedRequestsOk() (*[]string, bool)`

GetFailedRequestsOk returns a tuple with the FailedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRequests

`func (o *GlobalRequestResult) SetFailedRequests(v []string)`

SetFailedRequests sets FailedRequests field to given value.

### HasFailedRequests

`func (o *GlobalRequestResult) HasFailedRequests() bool`

HasFailedRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


