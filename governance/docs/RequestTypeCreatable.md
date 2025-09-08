# RequestTypeCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RequestTypeCreatableStatus**](RequestTypeCreatableStatus.md) |  | [optional] [default to REQUESTTYPECREATABLESTATUS_DRAFT]
**OwnerId** | **string** | The ID of the team that administers this request type. | 
**ResourceSettings** | [**RequestTypeResourceSettingsMutable**](RequestTypeResourceSettingsMutable.md) |  | 
**RequestSettings** | Pointer to [**RequestTypeRequestSettingsMutable**](RequestTypeRequestSettingsMutable.md) |  | [optional] [default to {"type":"EVERYONE","requesterFields":[]}]
**ApprovalSettings** | [**RequestTypeApprovalSettingsMutable**](RequestTypeApprovalSettingsMutable.md) |  | 
**AccessDuration** | Pointer to **NullableString** | How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported, for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  &gt; **Note:** Mixes of units, as well as month/year/week designations, are not supported. For example, &#x60;P40DT65H&#x60;, &#x60;P40M&#x60;, &#x60;P1W&#x60; and &#x60;P1Y&#x60; are not supported. | [optional] 
**Name** | **string** | Writable unique key on Create. Not modifiable on update. | 
**Description** | Pointer to **string** | Human readable description. | [optional] 

## Methods

### NewRequestTypeCreatable

`func NewRequestTypeCreatable(ownerId string, resourceSettings RequestTypeResourceSettingsMutable, approvalSettings RequestTypeApprovalSettingsMutable, name string, ) *RequestTypeCreatable`

NewRequestTypeCreatable instantiates a new RequestTypeCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeCreatableWithDefaults

`func NewRequestTypeCreatableWithDefaults() *RequestTypeCreatable`

NewRequestTypeCreatableWithDefaults instantiates a new RequestTypeCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestTypeCreatable) GetStatus() RequestTypeCreatableStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestTypeCreatable) GetStatusOk() (*RequestTypeCreatableStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestTypeCreatable) SetStatus(v RequestTypeCreatableStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestTypeCreatable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOwnerId

`func (o *RequestTypeCreatable) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *RequestTypeCreatable) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *RequestTypeCreatable) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetResourceSettings

`func (o *RequestTypeCreatable) GetResourceSettings() RequestTypeResourceSettingsMutable`

GetResourceSettings returns the ResourceSettings field if non-nil, zero value otherwise.

### GetResourceSettingsOk

`func (o *RequestTypeCreatable) GetResourceSettingsOk() (*RequestTypeResourceSettingsMutable, bool)`

GetResourceSettingsOk returns a tuple with the ResourceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettings

`func (o *RequestTypeCreatable) SetResourceSettings(v RequestTypeResourceSettingsMutable)`

SetResourceSettings sets ResourceSettings field to given value.


### GetRequestSettings

`func (o *RequestTypeCreatable) GetRequestSettings() RequestTypeRequestSettingsMutable`

GetRequestSettings returns the RequestSettings field if non-nil, zero value otherwise.

### GetRequestSettingsOk

`func (o *RequestTypeCreatable) GetRequestSettingsOk() (*RequestTypeRequestSettingsMutable, bool)`

GetRequestSettingsOk returns a tuple with the RequestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSettings

`func (o *RequestTypeCreatable) SetRequestSettings(v RequestTypeRequestSettingsMutable)`

SetRequestSettings sets RequestSettings field to given value.

### HasRequestSettings

`func (o *RequestTypeCreatable) HasRequestSettings() bool`

HasRequestSettings returns a boolean if a field has been set.

### GetApprovalSettings

`func (o *RequestTypeCreatable) GetApprovalSettings() RequestTypeApprovalSettingsMutable`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *RequestTypeCreatable) GetApprovalSettingsOk() (*RequestTypeApprovalSettingsMutable, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *RequestTypeCreatable) SetApprovalSettings(v RequestTypeApprovalSettingsMutable)`

SetApprovalSettings sets ApprovalSettings field to given value.


### GetAccessDuration

`func (o *RequestTypeCreatable) GetAccessDuration() string`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *RequestTypeCreatable) GetAccessDurationOk() (*string, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *RequestTypeCreatable) SetAccessDuration(v string)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *RequestTypeCreatable) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.

### SetAccessDurationNil

`func (o *RequestTypeCreatable) SetAccessDurationNil(b bool)`

 SetAccessDurationNil sets the value for AccessDuration to be an explicit nil

### UnsetAccessDuration
`func (o *RequestTypeCreatable) UnsetAccessDuration()`

UnsetAccessDuration ensures that no value is present for AccessDuration, not even an explicit nil
### GetName

`func (o *RequestTypeCreatable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTypeCreatable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTypeCreatable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestTypeCreatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeCreatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeCreatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTypeCreatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


