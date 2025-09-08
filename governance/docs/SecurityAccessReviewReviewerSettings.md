# SecurityAccessReviewReviewerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SecurityAccessReviewReviewerType**](SecurityAccessReviewReviewerType.md) |  | 
**UserSettings** | Pointer to [**SecurityAccessReviewUserReviewerSettings**](SecurityAccessReviewUserReviewerSettings.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewReviewerSettings

`func NewSecurityAccessReviewReviewerSettings(type_ SecurityAccessReviewReviewerType, ) *SecurityAccessReviewReviewerSettings`

NewSecurityAccessReviewReviewerSettings instantiates a new SecurityAccessReviewReviewerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewReviewerSettingsWithDefaults

`func NewSecurityAccessReviewReviewerSettingsWithDefaults() *SecurityAccessReviewReviewerSettings`

NewSecurityAccessReviewReviewerSettingsWithDefaults instantiates a new SecurityAccessReviewReviewerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecurityAccessReviewReviewerSettings) GetType() SecurityAccessReviewReviewerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewReviewerSettings) GetTypeOk() (*SecurityAccessReviewReviewerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewReviewerSettings) SetType(v SecurityAccessReviewReviewerType)`

SetType sets Type field to given value.


### GetUserSettings

`func (o *SecurityAccessReviewReviewerSettings) GetUserSettings() SecurityAccessReviewUserReviewerSettings`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *SecurityAccessReviewReviewerSettings) GetUserSettingsOk() (*SecurityAccessReviewUserReviewerSettings, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *SecurityAccessReviewReviewerSettings) SetUserSettings(v SecurityAccessReviewUserReviewerSettings)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *SecurityAccessReviewReviewerSettings) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


