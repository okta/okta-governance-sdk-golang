# RemediationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessApproved** | [**ApprovedRemediationAction**](ApprovedRemediationAction.md) |  | [default to APPROVEDREMEDIATIONACTION_NO_ACTION]
**AccessRevoked** | [**RevokedRemediationAction**](RevokedRemediationAction.md) |  | [default to REVOKEDREMEDIATIONACTION_NO_ACTION]
**NoResponse** | [**NoResponseRemediationAction**](NoResponseRemediationAction.md) |  | [default to NORESPONSEREMEDIATIONACTION_NO_ACTION]
**AutoRemediationSettings** | Pointer to [**AutoRemediationSettings**](AutoRemediationSettings.md) |  | [optional] 

## Methods

### NewRemediationSettings

`func NewRemediationSettings(accessApproved ApprovedRemediationAction, accessRevoked RevokedRemediationAction, noResponse NoResponseRemediationAction, ) *RemediationSettings`

NewRemediationSettings instantiates a new RemediationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationSettingsWithDefaults

`func NewRemediationSettingsWithDefaults() *RemediationSettings`

NewRemediationSettingsWithDefaults instantiates a new RemediationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessApproved

`func (o *RemediationSettings) GetAccessApproved() ApprovedRemediationAction`

GetAccessApproved returns the AccessApproved field if non-nil, zero value otherwise.

### GetAccessApprovedOk

`func (o *RemediationSettings) GetAccessApprovedOk() (*ApprovedRemediationAction, bool)`

GetAccessApprovedOk returns a tuple with the AccessApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessApproved

`func (o *RemediationSettings) SetAccessApproved(v ApprovedRemediationAction)`

SetAccessApproved sets AccessApproved field to given value.


### GetAccessRevoked

`func (o *RemediationSettings) GetAccessRevoked() RevokedRemediationAction`

GetAccessRevoked returns the AccessRevoked field if non-nil, zero value otherwise.

### GetAccessRevokedOk

`func (o *RemediationSettings) GetAccessRevokedOk() (*RevokedRemediationAction, bool)`

GetAccessRevokedOk returns a tuple with the AccessRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRevoked

`func (o *RemediationSettings) SetAccessRevoked(v RevokedRemediationAction)`

SetAccessRevoked sets AccessRevoked field to given value.


### GetNoResponse

`func (o *RemediationSettings) GetNoResponse() NoResponseRemediationAction`

GetNoResponse returns the NoResponse field if non-nil, zero value otherwise.

### GetNoResponseOk

`func (o *RemediationSettings) GetNoResponseOk() (*NoResponseRemediationAction, bool)`

GetNoResponseOk returns a tuple with the NoResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoResponse

`func (o *RemediationSettings) SetNoResponse(v NoResponseRemediationAction)`

SetNoResponse sets NoResponse field to given value.


### GetAutoRemediationSettings

`func (o *RemediationSettings) GetAutoRemediationSettings() AutoRemediationSettings`

GetAutoRemediationSettings returns the AutoRemediationSettings field if non-nil, zero value otherwise.

### GetAutoRemediationSettingsOk

`func (o *RemediationSettings) GetAutoRemediationSettingsOk() (*AutoRemediationSettings, bool)`

GetAutoRemediationSettingsOk returns a tuple with the AutoRemediationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRemediationSettings

`func (o *RemediationSettings) SetAutoRemediationSettings(v AutoRemediationSettings)`

SetAutoRemediationSettings sets AutoRemediationSettings field to given value.

### HasAutoRemediationSettings

`func (o *RemediationSettings) HasAutoRemediationSettings() bool`

HasAutoRemediationSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


