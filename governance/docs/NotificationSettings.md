# NotificationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyReviewerWhenReviewAssigned** | Pointer to **bool** | A Boolean value to indicate whether a notification should be sent to the reviewer when actionable reviews are assigned. | [optional] [default to false]
**NotifyReviewerAtCampaignEnd** | Pointer to **bool** | A Boolean value to indicate whether a notification should be sent to the reviewers when campaign has come to an end. | [optional] [default to false]
**RemindersReviewerBeforeCampaignCloseInSecs** | Pointer to **[]int32** | Specifies, in seconds, the time a reminder is sent to reviewers before the campaign closes. You can send up to three notifications. For example, the following array, &#x60;[86400, 172800, 604800]&#x60;, sends reminder notifications 7 days, 2 days, and 1 day before the campaign closes. By default, reminders are sent 2 days and 1 day before the campaign closes. | [optional] [default to [86400,172800]]
**NotifyReviewerWhenOverdue** | Pointer to **NullableBool** | A boolean value to indicate whether a notification should be sent to the reviewer when reviews are over due. | [optional] [default to false]
**NotifyReviewerDuringMidpointOfReview** | Pointer to **NullableBool** | A boolean value to indicate whether a notification should be sent to the reviewer during the midpoint of the review process. | [optional] [default to false]
**NotifyReviewPeriodEnd** | Pointer to **NullableBool** | Applicable for multi level campaigns. A boolean value to indicate whether a notification should be sent to the reviewer when a given reviewer level period is about to end. | [optional] [default to false]

## Methods

### NewNotificationSettings

`func NewNotificationSettings() *NotificationSettings`

NewNotificationSettings instantiates a new NotificationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsWithDefaults

`func NewNotificationSettingsWithDefaults() *NotificationSettings`

NewNotificationSettingsWithDefaults instantiates a new NotificationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyReviewerWhenReviewAssigned

`func (o *NotificationSettings) GetNotifyReviewerWhenReviewAssigned() bool`

GetNotifyReviewerWhenReviewAssigned returns the NotifyReviewerWhenReviewAssigned field if non-nil, zero value otherwise.

### GetNotifyReviewerWhenReviewAssignedOk

`func (o *NotificationSettings) GetNotifyReviewerWhenReviewAssignedOk() (*bool, bool)`

GetNotifyReviewerWhenReviewAssignedOk returns a tuple with the NotifyReviewerWhenReviewAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReviewerWhenReviewAssigned

`func (o *NotificationSettings) SetNotifyReviewerWhenReviewAssigned(v bool)`

SetNotifyReviewerWhenReviewAssigned sets NotifyReviewerWhenReviewAssigned field to given value.

### HasNotifyReviewerWhenReviewAssigned

`func (o *NotificationSettings) HasNotifyReviewerWhenReviewAssigned() bool`

HasNotifyReviewerWhenReviewAssigned returns a boolean if a field has been set.

### GetNotifyReviewerAtCampaignEnd

`func (o *NotificationSettings) GetNotifyReviewerAtCampaignEnd() bool`

GetNotifyReviewerAtCampaignEnd returns the NotifyReviewerAtCampaignEnd field if non-nil, zero value otherwise.

### GetNotifyReviewerAtCampaignEndOk

`func (o *NotificationSettings) GetNotifyReviewerAtCampaignEndOk() (*bool, bool)`

GetNotifyReviewerAtCampaignEndOk returns a tuple with the NotifyReviewerAtCampaignEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReviewerAtCampaignEnd

`func (o *NotificationSettings) SetNotifyReviewerAtCampaignEnd(v bool)`

SetNotifyReviewerAtCampaignEnd sets NotifyReviewerAtCampaignEnd field to given value.

### HasNotifyReviewerAtCampaignEnd

`func (o *NotificationSettings) HasNotifyReviewerAtCampaignEnd() bool`

HasNotifyReviewerAtCampaignEnd returns a boolean if a field has been set.

### GetRemindersReviewerBeforeCampaignCloseInSecs

`func (o *NotificationSettings) GetRemindersReviewerBeforeCampaignCloseInSecs() []int32`

GetRemindersReviewerBeforeCampaignCloseInSecs returns the RemindersReviewerBeforeCampaignCloseInSecs field if non-nil, zero value otherwise.

### GetRemindersReviewerBeforeCampaignCloseInSecsOk

`func (o *NotificationSettings) GetRemindersReviewerBeforeCampaignCloseInSecsOk() (*[]int32, bool)`

GetRemindersReviewerBeforeCampaignCloseInSecsOk returns a tuple with the RemindersReviewerBeforeCampaignCloseInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemindersReviewerBeforeCampaignCloseInSecs

`func (o *NotificationSettings) SetRemindersReviewerBeforeCampaignCloseInSecs(v []int32)`

SetRemindersReviewerBeforeCampaignCloseInSecs sets RemindersReviewerBeforeCampaignCloseInSecs field to given value.

### HasRemindersReviewerBeforeCampaignCloseInSecs

`func (o *NotificationSettings) HasRemindersReviewerBeforeCampaignCloseInSecs() bool`

HasRemindersReviewerBeforeCampaignCloseInSecs returns a boolean if a field has been set.

### GetNotifyReviewerWhenOverdue

`func (o *NotificationSettings) GetNotifyReviewerWhenOverdue() bool`

GetNotifyReviewerWhenOverdue returns the NotifyReviewerWhenOverdue field if non-nil, zero value otherwise.

### GetNotifyReviewerWhenOverdueOk

`func (o *NotificationSettings) GetNotifyReviewerWhenOverdueOk() (*bool, bool)`

GetNotifyReviewerWhenOverdueOk returns a tuple with the NotifyReviewerWhenOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReviewerWhenOverdue

`func (o *NotificationSettings) SetNotifyReviewerWhenOverdue(v bool)`

SetNotifyReviewerWhenOverdue sets NotifyReviewerWhenOverdue field to given value.

### HasNotifyReviewerWhenOverdue

`func (o *NotificationSettings) HasNotifyReviewerWhenOverdue() bool`

HasNotifyReviewerWhenOverdue returns a boolean if a field has been set.

### SetNotifyReviewerWhenOverdueNil

`func (o *NotificationSettings) SetNotifyReviewerWhenOverdueNil(b bool)`

 SetNotifyReviewerWhenOverdueNil sets the value for NotifyReviewerWhenOverdue to be an explicit nil

### UnsetNotifyReviewerWhenOverdue
`func (o *NotificationSettings) UnsetNotifyReviewerWhenOverdue()`

UnsetNotifyReviewerWhenOverdue ensures that no value is present for NotifyReviewerWhenOverdue, not even an explicit nil
### GetNotifyReviewerDuringMidpointOfReview

`func (o *NotificationSettings) GetNotifyReviewerDuringMidpointOfReview() bool`

GetNotifyReviewerDuringMidpointOfReview returns the NotifyReviewerDuringMidpointOfReview field if non-nil, zero value otherwise.

### GetNotifyReviewerDuringMidpointOfReviewOk

`func (o *NotificationSettings) GetNotifyReviewerDuringMidpointOfReviewOk() (*bool, bool)`

GetNotifyReviewerDuringMidpointOfReviewOk returns a tuple with the NotifyReviewerDuringMidpointOfReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReviewerDuringMidpointOfReview

`func (o *NotificationSettings) SetNotifyReviewerDuringMidpointOfReview(v bool)`

SetNotifyReviewerDuringMidpointOfReview sets NotifyReviewerDuringMidpointOfReview field to given value.

### HasNotifyReviewerDuringMidpointOfReview

`func (o *NotificationSettings) HasNotifyReviewerDuringMidpointOfReview() bool`

HasNotifyReviewerDuringMidpointOfReview returns a boolean if a field has been set.

### SetNotifyReviewerDuringMidpointOfReviewNil

`func (o *NotificationSettings) SetNotifyReviewerDuringMidpointOfReviewNil(b bool)`

 SetNotifyReviewerDuringMidpointOfReviewNil sets the value for NotifyReviewerDuringMidpointOfReview to be an explicit nil

### UnsetNotifyReviewerDuringMidpointOfReview
`func (o *NotificationSettings) UnsetNotifyReviewerDuringMidpointOfReview()`

UnsetNotifyReviewerDuringMidpointOfReview ensures that no value is present for NotifyReviewerDuringMidpointOfReview, not even an explicit nil
### GetNotifyReviewPeriodEnd

`func (o *NotificationSettings) GetNotifyReviewPeriodEnd() bool`

GetNotifyReviewPeriodEnd returns the NotifyReviewPeriodEnd field if non-nil, zero value otherwise.

### GetNotifyReviewPeriodEndOk

`func (o *NotificationSettings) GetNotifyReviewPeriodEndOk() (*bool, bool)`

GetNotifyReviewPeriodEndOk returns a tuple with the NotifyReviewPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyReviewPeriodEnd

`func (o *NotificationSettings) SetNotifyReviewPeriodEnd(v bool)`

SetNotifyReviewPeriodEnd sets NotifyReviewPeriodEnd field to given value.

### HasNotifyReviewPeriodEnd

`func (o *NotificationSettings) HasNotifyReviewPeriodEnd() bool`

HasNotifyReviewPeriodEnd returns a boolean if a field has been set.

### SetNotifyReviewPeriodEndNil

`func (o *NotificationSettings) SetNotifyReviewPeriodEndNil(b bool)`

 SetNotifyReviewPeriodEndNil sets the value for NotifyReviewPeriodEnd to be an explicit nil

### UnsetNotifyReviewPeriodEnd
`func (o *NotificationSettings) UnsetNotifyReviewPeriodEnd()`

UnsetNotifyReviewPeriodEnd ensures that no value is present for NotifyReviewPeriodEnd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


