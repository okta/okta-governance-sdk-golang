# SlackIntegrationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of integration | 
**IntegrationId** | **string** | The integration ID | 
**Notifications** | Pointer to [**SlackNotificationSettings**](SlackNotificationSettings.md) |  | [optional] 
**CanInitiateRequest** | Pointer to **bool** | Indicates whether requests can be initiated from this Slack integration | [optional] 
**CanApproveRequest** | Pointer to **bool** | Indicates whether requests can be approved from this Slack integration | [optional] 

## Methods

### NewSlackIntegrationSettings

`func NewSlackIntegrationSettings(type_ string, integrationId string, ) *SlackIntegrationSettings`

NewSlackIntegrationSettings instantiates a new SlackIntegrationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackIntegrationSettingsWithDefaults

`func NewSlackIntegrationSettingsWithDefaults() *SlackIntegrationSettings`

NewSlackIntegrationSettingsWithDefaults instantiates a new SlackIntegrationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SlackIntegrationSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlackIntegrationSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlackIntegrationSettings) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *SlackIntegrationSettings) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *SlackIntegrationSettings) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *SlackIntegrationSettings) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetNotifications

`func (o *SlackIntegrationSettings) GetNotifications() SlackNotificationSettings`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SlackIntegrationSettings) GetNotificationsOk() (*SlackNotificationSettings, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SlackIntegrationSettings) SetNotifications(v SlackNotificationSettings)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SlackIntegrationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetCanInitiateRequest

`func (o *SlackIntegrationSettings) GetCanInitiateRequest() bool`

GetCanInitiateRequest returns the CanInitiateRequest field if non-nil, zero value otherwise.

### GetCanInitiateRequestOk

`func (o *SlackIntegrationSettings) GetCanInitiateRequestOk() (*bool, bool)`

GetCanInitiateRequestOk returns a tuple with the CanInitiateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInitiateRequest

`func (o *SlackIntegrationSettings) SetCanInitiateRequest(v bool)`

SetCanInitiateRequest sets CanInitiateRequest field to given value.

### HasCanInitiateRequest

`func (o *SlackIntegrationSettings) HasCanInitiateRequest() bool`

HasCanInitiateRequest returns a boolean if a field has been set.

### GetCanApproveRequest

`func (o *SlackIntegrationSettings) GetCanApproveRequest() bool`

GetCanApproveRequest returns the CanApproveRequest field if non-nil, zero value otherwise.

### GetCanApproveRequestOk

`func (o *SlackIntegrationSettings) GetCanApproveRequestOk() (*bool, bool)`

GetCanApproveRequestOk returns a tuple with the CanApproveRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApproveRequest

`func (o *SlackIntegrationSettings) SetCanApproveRequest(v bool)`

SetCanApproveRequest sets CanApproveRequest field to given value.

### HasCanApproveRequest

`func (o *SlackIntegrationSettings) HasCanApproveRequest() bool`

HasCanApproveRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


