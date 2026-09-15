# EntitlementReconciliationConfigFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the configuration. Present only after a configuration has been saved for the resource. | [optional] 
**CreatedBy** | Pointer to **string** | The &#x60;id&#x60; of the Okta user who created the resource | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the resource was created | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [optional] [readonly] 
**Mode** | [**ReconciliationMode**](ReconciliationMode.md) |  | 
**Additive** | Pointer to [**EntitlementReconciliationDirectionFull**](EntitlementReconciliationDirectionFull.md) |  | [optional] 
**Subtractive** | Pointer to [**EntitlementReconciliationDirectionFull**](EntitlementReconciliationDirectionFull.md) |  | [optional] 
**ResourceOrn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**Links** | [**EntitlementReconciliationConfigLinks**](EntitlementReconciliationConfigLinks.md) |  | 

## Methods

### NewEntitlementReconciliationConfigFull

`func NewEntitlementReconciliationConfigFull(mode ReconciliationMode, resourceOrn string, links EntitlementReconciliationConfigLinks, ) *EntitlementReconciliationConfigFull`

NewEntitlementReconciliationConfigFull instantiates a new EntitlementReconciliationConfigFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementReconciliationConfigFullWithDefaults

`func NewEntitlementReconciliationConfigFullWithDefaults() *EntitlementReconciliationConfigFull`

NewEntitlementReconciliationConfigFullWithDefaults instantiates a new EntitlementReconciliationConfigFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementReconciliationConfigFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementReconciliationConfigFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementReconciliationConfigFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementReconciliationConfigFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EntitlementReconciliationConfigFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntitlementReconciliationConfigFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntitlementReconciliationConfigFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EntitlementReconciliationConfigFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreated

`func (o *EntitlementReconciliationConfigFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementReconciliationConfigFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementReconciliationConfigFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntitlementReconciliationConfigFull) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EntitlementReconciliationConfigFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntitlementReconciliationConfigFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntitlementReconciliationConfigFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EntitlementReconciliationConfigFull) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EntitlementReconciliationConfigFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntitlementReconciliationConfigFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntitlementReconciliationConfigFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EntitlementReconciliationConfigFull) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetMode

`func (o *EntitlementReconciliationConfigFull) GetMode() ReconciliationMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EntitlementReconciliationConfigFull) GetModeOk() (*ReconciliationMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EntitlementReconciliationConfigFull) SetMode(v ReconciliationMode)`

SetMode sets Mode field to given value.


### GetAdditive

`func (o *EntitlementReconciliationConfigFull) GetAdditive() EntitlementReconciliationDirectionFull`

GetAdditive returns the Additive field if non-nil, zero value otherwise.

### GetAdditiveOk

`func (o *EntitlementReconciliationConfigFull) GetAdditiveOk() (*EntitlementReconciliationDirectionFull, bool)`

GetAdditiveOk returns a tuple with the Additive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditive

`func (o *EntitlementReconciliationConfigFull) SetAdditive(v EntitlementReconciliationDirectionFull)`

SetAdditive sets Additive field to given value.

### HasAdditive

`func (o *EntitlementReconciliationConfigFull) HasAdditive() bool`

HasAdditive returns a boolean if a field has been set.

### GetSubtractive

`func (o *EntitlementReconciliationConfigFull) GetSubtractive() EntitlementReconciliationDirectionFull`

GetSubtractive returns the Subtractive field if non-nil, zero value otherwise.

### GetSubtractiveOk

`func (o *EntitlementReconciliationConfigFull) GetSubtractiveOk() (*EntitlementReconciliationDirectionFull, bool)`

GetSubtractiveOk returns a tuple with the Subtractive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtractive

`func (o *EntitlementReconciliationConfigFull) SetSubtractive(v EntitlementReconciliationDirectionFull)`

SetSubtractive sets Subtractive field to given value.

### HasSubtractive

`func (o *EntitlementReconciliationConfigFull) HasSubtractive() bool`

HasSubtractive returns a boolean if a field has been set.

### GetResourceOrn

`func (o *EntitlementReconciliationConfigFull) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *EntitlementReconciliationConfigFull) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *EntitlementReconciliationConfigFull) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetLinks

`func (o *EntitlementReconciliationConfigFull) GetLinks() EntitlementReconciliationConfigLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementReconciliationConfigFull) GetLinksOk() (*EntitlementReconciliationConfigLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementReconciliationConfigFull) SetLinks(v EntitlementReconciliationConfigLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


