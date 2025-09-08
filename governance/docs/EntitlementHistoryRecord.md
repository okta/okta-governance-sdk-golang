# EntitlementHistoryRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **time.Time** | The start date and time when the entitlements became effective | [optional] 
**EndDate** | Pointer to **time.Time** | The end date and time when the entitlements were superseded (if empty, the entitlements are currently effective) | [optional] 
**Lifecycle** | Pointer to **string** | The current status of the entitlements for this history entry | [optional] 
**Entitlements** | Pointer to [**[]EntitlementDetail**](EntitlementDetail.md) | List of entitlements for this history entry | [optional] 

## Methods

### NewEntitlementHistoryRecord

`func NewEntitlementHistoryRecord() *EntitlementHistoryRecord`

NewEntitlementHistoryRecord instantiates a new EntitlementHistoryRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementHistoryRecordWithDefaults

`func NewEntitlementHistoryRecordWithDefaults() *EntitlementHistoryRecord`

NewEntitlementHistoryRecordWithDefaults instantiates a new EntitlementHistoryRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *EntitlementHistoryRecord) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EntitlementHistoryRecord) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EntitlementHistoryRecord) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EntitlementHistoryRecord) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *EntitlementHistoryRecord) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *EntitlementHistoryRecord) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *EntitlementHistoryRecord) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *EntitlementHistoryRecord) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLifecycle

`func (o *EntitlementHistoryRecord) GetLifecycle() string`

GetLifecycle returns the Lifecycle field if non-nil, zero value otherwise.

### GetLifecycleOk

`func (o *EntitlementHistoryRecord) GetLifecycleOk() (*string, bool)`

GetLifecycleOk returns a tuple with the Lifecycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycle

`func (o *EntitlementHistoryRecord) SetLifecycle(v string)`

SetLifecycle sets Lifecycle field to given value.

### HasLifecycle

`func (o *EntitlementHistoryRecord) HasLifecycle() bool`

HasLifecycle returns a boolean if a field has been set.

### GetEntitlements

`func (o *EntitlementHistoryRecord) GetEntitlements() []EntitlementDetail`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementHistoryRecord) GetEntitlementsOk() (*[]EntitlementDetail, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementHistoryRecord) SetEntitlements(v []EntitlementDetail)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *EntitlementHistoryRecord) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


