# EntitlementBundleFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Status** | [**EntitlementBundleStatus**](EntitlementBundleStatus.md) |  | 
**Entitlements** | Pointer to [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | [optional] 
**Links** | [**EntitlementBundleLinks**](EntitlementBundleLinks.md) |  | 
**Name** | **string** | The unique name of the entitlement bundle | 
**Description** | Pointer to **string** | The human-readable description | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewEntitlementBundleFull

`func NewEntitlementBundleFull(targetResourceOrn string, target TargetResource, status EntitlementBundleStatus, links EntitlementBundleLinks, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *EntitlementBundleFull`

NewEntitlementBundleFull instantiates a new EntitlementBundleFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementBundleFullWithDefaults

`func NewEntitlementBundleFullWithDefaults() *EntitlementBundleFull`

NewEntitlementBundleFullWithDefaults instantiates a new EntitlementBundleFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetResourceOrn

`func (o *EntitlementBundleFull) GetTargetResourceOrn() string`

GetTargetResourceOrn returns the TargetResourceOrn field if non-nil, zero value otherwise.

### GetTargetResourceOrnOk

`func (o *EntitlementBundleFull) GetTargetResourceOrnOk() (*string, bool)`

GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceOrn

`func (o *EntitlementBundleFull) SetTargetResourceOrn(v string)`

SetTargetResourceOrn sets TargetResourceOrn field to given value.


### GetTarget

`func (o *EntitlementBundleFull) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EntitlementBundleFull) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EntitlementBundleFull) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetStatus

`func (o *EntitlementBundleFull) GetStatus() EntitlementBundleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntitlementBundleFull) GetStatusOk() (*EntitlementBundleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntitlementBundleFull) SetStatus(v EntitlementBundleStatus)`

SetStatus sets Status field to given value.


### GetEntitlements

`func (o *EntitlementBundleFull) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementBundleFull) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementBundleFull) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *EntitlementBundleFull) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetLinks

`func (o *EntitlementBundleFull) GetLinks() EntitlementBundleLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementBundleFull) GetLinksOk() (*EntitlementBundleLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementBundleFull) SetLinks(v EntitlementBundleLinks)`

SetLinks sets Links field to given value.


### GetName

`func (o *EntitlementBundleFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementBundleFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementBundleFull) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntitlementBundleFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementBundleFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementBundleFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementBundleFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *EntitlementBundleFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementBundleFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementBundleFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *EntitlementBundleFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntitlementBundleFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntitlementBundleFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *EntitlementBundleFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementBundleFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementBundleFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *EntitlementBundleFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntitlementBundleFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntitlementBundleFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *EntitlementBundleFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntitlementBundleFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntitlementBundleFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


