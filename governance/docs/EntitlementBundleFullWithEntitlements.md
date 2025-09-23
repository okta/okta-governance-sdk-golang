# EntitlementBundleFullWithEntitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Status** | [**EntitlementBundleStatus**](EntitlementBundleStatus.md) |  | 
**Entitlements** | Pointer to [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | [optional] 
**Links** | [**EntitlementBundleLinks**](EntitlementBundleLinks.md) |  | 
**Name** | **string** | The unique name of the entitlement bundle | 
**Description** | Pointer to **string** | The human-readable description | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewEntitlementBundleFullWithEntitlements

`func NewEntitlementBundleFullWithEntitlements(targetResourceOrn string, target TargetResource, status EntitlementBundleStatus, links EntitlementBundleLinks, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *EntitlementBundleFullWithEntitlements`

NewEntitlementBundleFullWithEntitlements instantiates a new EntitlementBundleFullWithEntitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementBundleFullWithEntitlementsWithDefaults

`func NewEntitlementBundleFullWithEntitlementsWithDefaults() *EntitlementBundleFullWithEntitlements`

NewEntitlementBundleFullWithEntitlementsWithDefaults instantiates a new EntitlementBundleFullWithEntitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetResourceOrn

`func (o *EntitlementBundleFullWithEntitlements) GetTargetResourceOrn() string`

GetTargetResourceOrn returns the TargetResourceOrn field if non-nil, zero value otherwise.

### GetTargetResourceOrnOk

`func (o *EntitlementBundleFullWithEntitlements) GetTargetResourceOrnOk() (*string, bool)`

GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceOrn

`func (o *EntitlementBundleFullWithEntitlements) SetTargetResourceOrn(v string)`

SetTargetResourceOrn sets TargetResourceOrn field to given value.


### GetTarget

`func (o *EntitlementBundleFullWithEntitlements) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EntitlementBundleFullWithEntitlements) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EntitlementBundleFullWithEntitlements) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetStatus

`func (o *EntitlementBundleFullWithEntitlements) GetStatus() EntitlementBundleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntitlementBundleFullWithEntitlements) GetStatusOk() (*EntitlementBundleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntitlementBundleFullWithEntitlements) SetStatus(v EntitlementBundleStatus)`

SetStatus sets Status field to given value.


### GetEntitlements

`func (o *EntitlementBundleFullWithEntitlements) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementBundleFullWithEntitlements) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementBundleFullWithEntitlements) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *EntitlementBundleFullWithEntitlements) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetLinks

`func (o *EntitlementBundleFullWithEntitlements) GetLinks() EntitlementBundleLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementBundleFullWithEntitlements) GetLinksOk() (*EntitlementBundleLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementBundleFullWithEntitlements) SetLinks(v EntitlementBundleLinks)`

SetLinks sets Links field to given value.


### GetName

`func (o *EntitlementBundleFullWithEntitlements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementBundleFullWithEntitlements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementBundleFullWithEntitlements) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntitlementBundleFullWithEntitlements) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementBundleFullWithEntitlements) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementBundleFullWithEntitlements) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementBundleFullWithEntitlements) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *EntitlementBundleFullWithEntitlements) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementBundleFullWithEntitlements) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementBundleFullWithEntitlements) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *EntitlementBundleFullWithEntitlements) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntitlementBundleFullWithEntitlements) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntitlementBundleFullWithEntitlements) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *EntitlementBundleFullWithEntitlements) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementBundleFullWithEntitlements) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementBundleFullWithEntitlements) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *EntitlementBundleFullWithEntitlements) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntitlementBundleFullWithEntitlements) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntitlementBundleFullWithEntitlements) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *EntitlementBundleFullWithEntitlements) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntitlementBundleFullWithEntitlements) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntitlementBundleFullWithEntitlements) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


