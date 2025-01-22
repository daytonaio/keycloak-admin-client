# PolicyEvaluationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]EvaluationResultRepresentation**](EvaluationResultRepresentation.md) |  | [optional] 
**Entitlements** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**DecisionEffect**](DecisionEffect.md) |  | [optional] 
**Rpt** | Pointer to [**AccessToken**](AccessToken.md) |  | [optional] 

## Methods

### NewPolicyEvaluationResponse

`func NewPolicyEvaluationResponse() *PolicyEvaluationResponse`

NewPolicyEvaluationResponse instantiates a new PolicyEvaluationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationResponseWithDefaults

`func NewPolicyEvaluationResponseWithDefaults() *PolicyEvaluationResponse`

NewPolicyEvaluationResponseWithDefaults instantiates a new PolicyEvaluationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PolicyEvaluationResponse) GetResults() []EvaluationResultRepresentation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PolicyEvaluationResponse) GetResultsOk() (*[]EvaluationResultRepresentation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PolicyEvaluationResponse) SetResults(v []EvaluationResultRepresentation)`

SetResults sets Results field to given value.

### HasResults

`func (o *PolicyEvaluationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetEntitlements

`func (o *PolicyEvaluationResponse) GetEntitlements() bool`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *PolicyEvaluationResponse) GetEntitlementsOk() (*bool, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *PolicyEvaluationResponse) SetEntitlements(v bool)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *PolicyEvaluationResponse) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetStatus

`func (o *PolicyEvaluationResponse) GetStatus() DecisionEffect`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PolicyEvaluationResponse) GetStatusOk() (*DecisionEffect, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PolicyEvaluationResponse) SetStatus(v DecisionEffect)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PolicyEvaluationResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRpt

`func (o *PolicyEvaluationResponse) GetRpt() AccessToken`

GetRpt returns the Rpt field if non-nil, zero value otherwise.

### GetRptOk

`func (o *PolicyEvaluationResponse) GetRptOk() (*AccessToken, bool)`

GetRptOk returns a tuple with the Rpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpt

`func (o *PolicyEvaluationResponse) SetRpt(v AccessToken)`

SetRpt sets Rpt field to given value.

### HasRpt

`func (o *PolicyEvaluationResponse) HasRpt() bool`

HasRpt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


