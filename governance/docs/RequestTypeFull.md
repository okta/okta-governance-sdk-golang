# RequestTypeFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** | The ID of the team that administers this request type. | 
**ResourceSettings** | [**RequestTypeResourceSettingsReadable**](RequestTypeResourceSettingsReadable.md) |  | 
**RequestSettings** | [**RequestTypeRequestSettingsReadable**](RequestTypeRequestSettingsReadable.md) |  | [default to {"type":"EVERYONE","requesterFields":[]}]
**ApprovalSettings** | [**RequestTypeApprovalSettingsReadable**](RequestTypeApprovalSettingsReadable.md) |  | 
**AccessDuration** | **NullableString** | How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported, for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  &gt; **Note:** Mixes of units, as well as month/year/week designations, are not supported. For example, &#x60;P40DT65H&#x60;, &#x60;P40M&#x60;, &#x60;P1W&#x60; and &#x60;P1Y&#x60; are not supported. | 
**Status** | [**RequestTypeStatus**](RequestTypeStatus.md) |  | 
**LastUpdateSource** | [**RequestTypeLastUpdateSource**](RequestTypeLastUpdateSource.md) |  | 
**Links** | [**RequestTypeLinks**](RequestTypeLinks.md) |  | 
**Name** | **string** | Writable unique key on Create. Not modifiable on update. | 
**Description** | **string** | Human readable description. | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewRequestTypeFull

`func NewRequestTypeFull(ownerId string, resourceSettings RequestTypeResourceSettingsReadable, requestSettings RequestTypeRequestSettingsReadable, approvalSettings RequestTypeApprovalSettingsReadable, accessDuration NullableString, status RequestTypeStatus, lastUpdateSource RequestTypeLastUpdateSource, links RequestTypeLinks, name string, description string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *RequestTypeFull`

NewRequestTypeFull instantiates a new RequestTypeFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeFullWithDefaults

`func NewRequestTypeFullWithDefaults() *RequestTypeFull`

NewRequestTypeFullWithDefaults instantiates a new RequestTypeFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *RequestTypeFull) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *RequestTypeFull) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *RequestTypeFull) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetResourceSettings

`func (o *RequestTypeFull) GetResourceSettings() RequestTypeResourceSettingsReadable`

GetResourceSettings returns the ResourceSettings field if non-nil, zero value otherwise.

### GetResourceSettingsOk

`func (o *RequestTypeFull) GetResourceSettingsOk() (*RequestTypeResourceSettingsReadable, bool)`

GetResourceSettingsOk returns a tuple with the ResourceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSettings

`func (o *RequestTypeFull) SetResourceSettings(v RequestTypeResourceSettingsReadable)`

SetResourceSettings sets ResourceSettings field to given value.


### GetRequestSettings

`func (o *RequestTypeFull) GetRequestSettings() RequestTypeRequestSettingsReadable`

GetRequestSettings returns the RequestSettings field if non-nil, zero value otherwise.

### GetRequestSettingsOk

`func (o *RequestTypeFull) GetRequestSettingsOk() (*RequestTypeRequestSettingsReadable, bool)`

GetRequestSettingsOk returns a tuple with the RequestSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSettings

`func (o *RequestTypeFull) SetRequestSettings(v RequestTypeRequestSettingsReadable)`

SetRequestSettings sets RequestSettings field to given value.


### GetApprovalSettings

`func (o *RequestTypeFull) GetApprovalSettings() RequestTypeApprovalSettingsReadable`

GetApprovalSettings returns the ApprovalSettings field if non-nil, zero value otherwise.

### GetApprovalSettingsOk

`func (o *RequestTypeFull) GetApprovalSettingsOk() (*RequestTypeApprovalSettingsReadable, bool)`

GetApprovalSettingsOk returns a tuple with the ApprovalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalSettings

`func (o *RequestTypeFull) SetApprovalSettings(v RequestTypeApprovalSettingsReadable)`

SetApprovalSettings sets ApprovalSettings field to given value.


### GetAccessDuration

`func (o *RequestTypeFull) GetAccessDuration() string`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *RequestTypeFull) GetAccessDurationOk() (*string, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *RequestTypeFull) SetAccessDuration(v string)`

SetAccessDuration sets AccessDuration field to given value.


### SetAccessDurationNil

`func (o *RequestTypeFull) SetAccessDurationNil(b bool)`

 SetAccessDurationNil sets the value for AccessDuration to be an explicit nil

### UnsetAccessDuration
`func (o *RequestTypeFull) UnsetAccessDuration()`

UnsetAccessDuration ensures that no value is present for AccessDuration, not even an explicit nil
### GetStatus

`func (o *RequestTypeFull) GetStatus() RequestTypeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestTypeFull) GetStatusOk() (*RequestTypeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestTypeFull) SetStatus(v RequestTypeStatus)`

SetStatus sets Status field to given value.


### GetLastUpdateSource

`func (o *RequestTypeFull) GetLastUpdateSource() RequestTypeLastUpdateSource`

GetLastUpdateSource returns the LastUpdateSource field if non-nil, zero value otherwise.

### GetLastUpdateSourceOk

`func (o *RequestTypeFull) GetLastUpdateSourceOk() (*RequestTypeLastUpdateSource, bool)`

GetLastUpdateSourceOk returns a tuple with the LastUpdateSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateSource

`func (o *RequestTypeFull) SetLastUpdateSource(v RequestTypeLastUpdateSource)`

SetLastUpdateSource sets LastUpdateSource field to given value.


### GetLinks

`func (o *RequestTypeFull) GetLinks() RequestTypeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestTypeFull) GetLinksOk() (*RequestTypeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestTypeFull) SetLinks(v RequestTypeLinks)`

SetLinks sets Links field to given value.


### GetName

`func (o *RequestTypeFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestTypeFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestTypeFull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestTypeFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeFull) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *RequestTypeFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestTypeFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestTypeFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestTypeFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestTypeFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestTypeFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestTypeFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestTypeFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestTypeFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestTypeFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestTypeFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestTypeFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestTypeFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestTypeFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestTypeFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


