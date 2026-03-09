# GrantFullWithEntitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | [**GrantType**](GrantType.md) |  | 
**EntitlementBundleId** | Pointer to **string** | The entitlement bundle &#x60;id&#x60;  | [optional] 
**TargetPrincipalOrn** | **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | 
**TargetPrincipal** | [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | 
**Action** | [**GrantAction**](GrantAction.md) |  | [default to GRANTACTION_ALLOW]
**Actor** | [**GrantActor**](GrantActor.md) |  | [default to GRANTACTOR_API]
**TargetResourceOrn** | **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Entitlements** | [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Links** | [**ResourceGrantLinks**](ResourceGrantLinks.md) |  | 
**Status** | [**GrantStatus**](GrantStatus.md) |  | 
**Metadata** | Pointer to [**GrantMetadata**](GrantMetadata.md) |  | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewGrantFullWithEntitlements

`func NewGrantFullWithEntitlements(grantType GrantType, targetPrincipalOrn string, targetPrincipal TargetPrincipalFull, action GrantAction, actor GrantActor, targetResourceOrn string, target TargetResource, entitlements []EntitlementFull, links ResourceGrantLinks, status GrantStatus, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *GrantFullWithEntitlements`

NewGrantFullWithEntitlements instantiates a new GrantFullWithEntitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantFullWithEntitlementsWithDefaults

`func NewGrantFullWithEntitlementsWithDefaults() *GrantFullWithEntitlements`

NewGrantFullWithEntitlementsWithDefaults instantiates a new GrantFullWithEntitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GrantFullWithEntitlements) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GrantFullWithEntitlements) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GrantFullWithEntitlements) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetEntitlementBundleId

`func (o *GrantFullWithEntitlements) GetEntitlementBundleId() string`

GetEntitlementBundleId returns the EntitlementBundleId field if non-nil, zero value otherwise.

### GetEntitlementBundleIdOk

`func (o *GrantFullWithEntitlements) GetEntitlementBundleIdOk() (*string, bool)`

GetEntitlementBundleIdOk returns a tuple with the EntitlementBundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundleId

`func (o *GrantFullWithEntitlements) SetEntitlementBundleId(v string)`

SetEntitlementBundleId sets EntitlementBundleId field to given value.

### HasEntitlementBundleId

`func (o *GrantFullWithEntitlements) HasEntitlementBundleId() bool`

HasEntitlementBundleId returns a boolean if a field has been set.

### GetTargetPrincipalOrn

`func (o *GrantFullWithEntitlements) GetTargetPrincipalOrn() string`

GetTargetPrincipalOrn returns the TargetPrincipalOrn field if non-nil, zero value otherwise.

### GetTargetPrincipalOrnOk

`func (o *GrantFullWithEntitlements) GetTargetPrincipalOrnOk() (*string, bool)`

GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipalOrn

`func (o *GrantFullWithEntitlements) SetTargetPrincipalOrn(v string)`

SetTargetPrincipalOrn sets TargetPrincipalOrn field to given value.


### GetTargetPrincipal

`func (o *GrantFullWithEntitlements) GetTargetPrincipal() TargetPrincipalFull`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *GrantFullWithEntitlements) GetTargetPrincipalOk() (*TargetPrincipalFull, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *GrantFullWithEntitlements) SetTargetPrincipal(v TargetPrincipalFull)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetAction

`func (o *GrantFullWithEntitlements) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GrantFullWithEntitlements) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GrantFullWithEntitlements) SetAction(v GrantAction)`

SetAction sets Action field to given value.


### GetActor

`func (o *GrantFullWithEntitlements) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GrantFullWithEntitlements) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GrantFullWithEntitlements) SetActor(v GrantActor)`

SetActor sets Actor field to given value.


### GetTargetResourceOrn

`func (o *GrantFullWithEntitlements) GetTargetResourceOrn() string`

GetTargetResourceOrn returns the TargetResourceOrn field if non-nil, zero value otherwise.

### GetTargetResourceOrnOk

`func (o *GrantFullWithEntitlements) GetTargetResourceOrnOk() (*string, bool)`

GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceOrn

`func (o *GrantFullWithEntitlements) SetTargetResourceOrn(v string)`

SetTargetResourceOrn sets TargetResourceOrn field to given value.


### GetTarget

`func (o *GrantFullWithEntitlements) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GrantFullWithEntitlements) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GrantFullWithEntitlements) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetEntitlements

`func (o *GrantFullWithEntitlements) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GrantFullWithEntitlements) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GrantFullWithEntitlements) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.


### GetScheduleSettings

`func (o *GrantFullWithEntitlements) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GrantFullWithEntitlements) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GrantFullWithEntitlements) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *GrantFullWithEntitlements) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetLinks

`func (o *GrantFullWithEntitlements) GetLinks() ResourceGrantLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GrantFullWithEntitlements) GetLinksOk() (*ResourceGrantLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GrantFullWithEntitlements) SetLinks(v ResourceGrantLinks)`

SetLinks sets Links field to given value.


### GetStatus

`func (o *GrantFullWithEntitlements) GetStatus() GrantStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GrantFullWithEntitlements) GetStatusOk() (*GrantStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GrantFullWithEntitlements) SetStatus(v GrantStatus)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *GrantFullWithEntitlements) GetMetadata() GrantMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GrantFullWithEntitlements) GetMetadataOk() (*GrantMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GrantFullWithEntitlements) SetMetadata(v GrantMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GrantFullWithEntitlements) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *GrantFullWithEntitlements) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantFullWithEntitlements) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantFullWithEntitlements) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *GrantFullWithEntitlements) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GrantFullWithEntitlements) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GrantFullWithEntitlements) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *GrantFullWithEntitlements) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GrantFullWithEntitlements) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GrantFullWithEntitlements) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *GrantFullWithEntitlements) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GrantFullWithEntitlements) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GrantFullWithEntitlements) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *GrantFullWithEntitlements) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GrantFullWithEntitlements) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GrantFullWithEntitlements) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


