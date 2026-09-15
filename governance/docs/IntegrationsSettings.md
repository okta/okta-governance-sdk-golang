# IntegrationsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Integration type | 
**IntegrationId** | **string** | The integration ID | 
**Notifications** | Pointer to [**NotificationsSettings**](NotificationsSettings.md) |  | [optional] 

## Methods

### NewIntegrationsSettings

`func NewIntegrationsSettings(type_ string, integrationId string, ) *IntegrationsSettings`

NewIntegrationsSettings instantiates a new IntegrationsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsSettingsWithDefaults

`func NewIntegrationsSettingsWithDefaults() *IntegrationsSettings`

NewIntegrationsSettingsWithDefaults instantiates a new IntegrationsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IntegrationsSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationsSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationsSettings) SetType(v string)`

SetType sets Type field to given value.


### GetIntegrationId

`func (o *IntegrationsSettings) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationsSettings) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationsSettings) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetNotifications

`func (o *IntegrationsSettings) GetNotifications() NotificationsSettings`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *IntegrationsSettings) GetNotificationsOk() (*NotificationsSettings, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *IntegrationsSettings) SetNotifications(v NotificationsSettings)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *IntegrationsSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


