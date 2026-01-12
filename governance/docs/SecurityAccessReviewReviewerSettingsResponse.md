# SecurityAccessReviewReviewerSettingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SecurityAccessReviewReviewerType**](SecurityAccessReviewReviewerType.md) |  | 
**UserSettings** | Pointer to [**SecurityAccessReviewUserReviewerSettingsAddnlDetails**](SecurityAccessReviewUserReviewerSettingsAddnlDetails.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewReviewerSettingsResponse

`func NewSecurityAccessReviewReviewerSettingsResponse(type_ SecurityAccessReviewReviewerType, ) *SecurityAccessReviewReviewerSettingsResponse`

NewSecurityAccessReviewReviewerSettingsResponse instantiates a new SecurityAccessReviewReviewerSettingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewReviewerSettingsResponseWithDefaults

`func NewSecurityAccessReviewReviewerSettingsResponseWithDefaults() *SecurityAccessReviewReviewerSettingsResponse`

NewSecurityAccessReviewReviewerSettingsResponseWithDefaults instantiates a new SecurityAccessReviewReviewerSettingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecurityAccessReviewReviewerSettingsResponse) GetType() SecurityAccessReviewReviewerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewReviewerSettingsResponse) GetTypeOk() (*SecurityAccessReviewReviewerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewReviewerSettingsResponse) SetType(v SecurityAccessReviewReviewerType)`

SetType sets Type field to given value.


### GetUserSettings

`func (o *SecurityAccessReviewReviewerSettingsResponse) GetUserSettings() SecurityAccessReviewUserReviewerSettingsAddnlDetails`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *SecurityAccessReviewReviewerSettingsResponse) GetUserSettingsOk() (*SecurityAccessReviewUserReviewerSettingsAddnlDetails, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *SecurityAccessReviewReviewerSettingsResponse) SetUserSettings(v SecurityAccessReviewUserReviewerSettingsAddnlDetails)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *SecurityAccessReviewReviewerSettingsResponse) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


