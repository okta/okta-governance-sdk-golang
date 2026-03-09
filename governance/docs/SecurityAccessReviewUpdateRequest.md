# SecurityAccessReviewUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **time.Time** | The date and time when the security access review closes. | [optional] 
**ReviewerSettings** | Pointer to [**SecurityAccessReviewReviewerSettings**](SecurityAccessReviewReviewerSettings.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewUpdateRequest

`func NewSecurityAccessReviewUpdateRequest() *SecurityAccessReviewUpdateRequest`

NewSecurityAccessReviewUpdateRequest instantiates a new SecurityAccessReviewUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewUpdateRequestWithDefaults

`func NewSecurityAccessReviewUpdateRequestWithDefaults() *SecurityAccessReviewUpdateRequest`

NewSecurityAccessReviewUpdateRequestWithDefaults instantiates a new SecurityAccessReviewUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *SecurityAccessReviewUpdateRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SecurityAccessReviewUpdateRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SecurityAccessReviewUpdateRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SecurityAccessReviewUpdateRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetReviewerSettings

`func (o *SecurityAccessReviewUpdateRequest) GetReviewerSettings() SecurityAccessReviewReviewerSettings`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *SecurityAccessReviewUpdateRequest) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *SecurityAccessReviewUpdateRequest) SetReviewerSettings(v SecurityAccessReviewReviewerSettings)`

SetReviewerSettings sets ReviewerSettings field to given value.

### HasReviewerSettings

`func (o *SecurityAccessReviewUpdateRequest) HasReviewerSettings() bool`

HasReviewerSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


