# SecurityAccessReviewsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCount** | **int32** | The number of active security access reviews | 
**PendingCount** | **int32** | The number of pending security access reviews | 
**ErrorCount** | **int32** | The number of errored security access reviews | 
**ClosedCount** | **int32** | The number of closed security access reviews | 

## Methods

### NewSecurityAccessReviewsStats

`func NewSecurityAccessReviewsStats(activeCount int32, pendingCount int32, errorCount int32, closedCount int32, ) *SecurityAccessReviewsStats`

NewSecurityAccessReviewsStats instantiates a new SecurityAccessReviewsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewsStatsWithDefaults

`func NewSecurityAccessReviewsStatsWithDefaults() *SecurityAccessReviewsStats`

NewSecurityAccessReviewsStatsWithDefaults instantiates a new SecurityAccessReviewsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCount

`func (o *SecurityAccessReviewsStats) GetActiveCount() int32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *SecurityAccessReviewsStats) GetActiveCountOk() (*int32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *SecurityAccessReviewsStats) SetActiveCount(v int32)`

SetActiveCount sets ActiveCount field to given value.


### GetPendingCount

`func (o *SecurityAccessReviewsStats) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *SecurityAccessReviewsStats) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *SecurityAccessReviewsStats) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.


### GetErrorCount

`func (o *SecurityAccessReviewsStats) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *SecurityAccessReviewsStats) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *SecurityAccessReviewsStats) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.


### GetClosedCount

`func (o *SecurityAccessReviewsStats) GetClosedCount() int32`

GetClosedCount returns the ClosedCount field if non-nil, zero value otherwise.

### GetClosedCountOk

`func (o *SecurityAccessReviewsStats) GetClosedCountOk() (*int32, bool)`

GetClosedCountOk returns a tuple with the ClosedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedCount

`func (o *SecurityAccessReviewsStats) SetClosedCount(v int32)`

SetClosedCount sets ClosedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


