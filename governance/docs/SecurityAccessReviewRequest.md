# SecurityAccessReviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | **string** | The Okta user ID in the security access review | 
**SecurityAccessReviewName** | **string** | The name of the security access review | 
**EndTime** | Pointer to **time.Time** | The end time of the security access review. If not specified this defaults to seven days from the start time. | [optional] 
**ReviewerSettings** | [**SecurityAccessReviewReviewerSettings**](SecurityAccessReviewReviewerSettings.md) |  | 

## Methods

### NewSecurityAccessReviewRequest

`func NewSecurityAccessReviewRequest(principalId string, securityAccessReviewName string, reviewerSettings SecurityAccessReviewReviewerSettings, ) *SecurityAccessReviewRequest`

NewSecurityAccessReviewRequest instantiates a new SecurityAccessReviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewRequestWithDefaults

`func NewSecurityAccessReviewRequestWithDefaults() *SecurityAccessReviewRequest`

NewSecurityAccessReviewRequestWithDefaults instantiates a new SecurityAccessReviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *SecurityAccessReviewRequest) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *SecurityAccessReviewRequest) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *SecurityAccessReviewRequest) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetSecurityAccessReviewName

`func (o *SecurityAccessReviewRequest) GetSecurityAccessReviewName() string`

GetSecurityAccessReviewName returns the SecurityAccessReviewName field if non-nil, zero value otherwise.

### GetSecurityAccessReviewNameOk

`func (o *SecurityAccessReviewRequest) GetSecurityAccessReviewNameOk() (*string, bool)`

GetSecurityAccessReviewNameOk returns a tuple with the SecurityAccessReviewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityAccessReviewName

`func (o *SecurityAccessReviewRequest) SetSecurityAccessReviewName(v string)`

SetSecurityAccessReviewName sets SecurityAccessReviewName field to given value.


### GetEndTime

`func (o *SecurityAccessReviewRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SecurityAccessReviewRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SecurityAccessReviewRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SecurityAccessReviewRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetReviewerSettings

`func (o *SecurityAccessReviewRequest) GetReviewerSettings() SecurityAccessReviewReviewerSettings`

GetReviewerSettings returns the ReviewerSettings field if non-nil, zero value otherwise.

### GetReviewerSettingsOk

`func (o *SecurityAccessReviewRequest) GetReviewerSettingsOk() (*SecurityAccessReviewReviewerSettings, bool)`

GetReviewerSettingsOk returns a tuple with the ReviewerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSettings

`func (o *SecurityAccessReviewRequest) SetReviewerSettings(v SecurityAccessReviewReviewerSettings)`

SetReviewerSettings sets ReviewerSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


