# GetGrant200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | [**GrantType**](GrantType.md) |  | 
**EntitlementBundleId** | Pointer to **string** | The entitlement bundle &#x60;id&#x60;  | [optional] 
**TargetPrincipalOrn** | **string** | The Okta user &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**TargetPrincipal** | [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | 
**Action** | [**GrantAction**](GrantAction.md) |  | [default to GRANTACTION_ALLOW]
**Actor** | [**GrantActor**](GrantActor.md) |  | [default to GRANTACTOR_API]
**TargetResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Entitlements** | [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | 
**ScheduleSettings** | Pointer to [**ScheduleSettingsWriteable**](ScheduleSettingsWriteable.md) |  | [optional] 
**Links** | [**map[string]Link**](Link.md) |  | 
**Status** | [**GrantStatus**](GrantStatus.md) |  | 
**Metadata** | Pointer to [**GrantMetadata**](GrantMetadata.md) |  | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewGetGrant200Response

`func NewGetGrant200Response(grantType GrantType, targetPrincipalOrn string, targetPrincipal TargetPrincipalFull, action GrantAction, actor GrantActor, targetResourceOrn string, target TargetResource, entitlements []EntitlementFull, links map[string]Link, status GrantStatus, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *GetGrant200Response`

NewGetGrant200Response instantiates a new GetGrant200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGrant200ResponseWithDefaults

`func NewGetGrant200ResponseWithDefaults() *GetGrant200Response`

NewGetGrant200ResponseWithDefaults instantiates a new GetGrant200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *GetGrant200Response) GetGrantType() GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *GetGrant200Response) GetGrantTypeOk() (*GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *GetGrant200Response) SetGrantType(v GrantType)`

SetGrantType sets GrantType field to given value.


### GetEntitlementBundleId

`func (o *GetGrant200Response) GetEntitlementBundleId() string`

GetEntitlementBundleId returns the EntitlementBundleId field if non-nil, zero value otherwise.

### GetEntitlementBundleIdOk

`func (o *GetGrant200Response) GetEntitlementBundleIdOk() (*string, bool)`

GetEntitlementBundleIdOk returns a tuple with the EntitlementBundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementBundleId

`func (o *GetGrant200Response) SetEntitlementBundleId(v string)`

SetEntitlementBundleId sets EntitlementBundleId field to given value.

### HasEntitlementBundleId

`func (o *GetGrant200Response) HasEntitlementBundleId() bool`

HasEntitlementBundleId returns a boolean if a field has been set.

### GetTargetPrincipalOrn

`func (o *GetGrant200Response) GetTargetPrincipalOrn() string`

GetTargetPrincipalOrn returns the TargetPrincipalOrn field if non-nil, zero value otherwise.

### GetTargetPrincipalOrnOk

`func (o *GetGrant200Response) GetTargetPrincipalOrnOk() (*string, bool)`

GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipalOrn

`func (o *GetGrant200Response) SetTargetPrincipalOrn(v string)`

SetTargetPrincipalOrn sets TargetPrincipalOrn field to given value.


### GetTargetPrincipal

`func (o *GetGrant200Response) GetTargetPrincipal() TargetPrincipalFull`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *GetGrant200Response) GetTargetPrincipalOk() (*TargetPrincipalFull, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *GetGrant200Response) SetTargetPrincipal(v TargetPrincipalFull)`

SetTargetPrincipal sets TargetPrincipal field to given value.


### GetAction

`func (o *GetGrant200Response) GetAction() GrantAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetGrant200Response) GetActionOk() (*GrantAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetGrant200Response) SetAction(v GrantAction)`

SetAction sets Action field to given value.


### GetActor

`func (o *GetGrant200Response) GetActor() GrantActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *GetGrant200Response) GetActorOk() (*GrantActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *GetGrant200Response) SetActor(v GrantActor)`

SetActor sets Actor field to given value.


### GetTargetResourceOrn

`func (o *GetGrant200Response) GetTargetResourceOrn() string`

GetTargetResourceOrn returns the TargetResourceOrn field if non-nil, zero value otherwise.

### GetTargetResourceOrnOk

`func (o *GetGrant200Response) GetTargetResourceOrnOk() (*string, bool)`

GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceOrn

`func (o *GetGrant200Response) SetTargetResourceOrn(v string)`

SetTargetResourceOrn sets TargetResourceOrn field to given value.


### GetTarget

`func (o *GetGrant200Response) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetGrant200Response) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetGrant200Response) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetEntitlements

`func (o *GetGrant200Response) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GetGrant200Response) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GetGrant200Response) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.


### GetScheduleSettings

`func (o *GetGrant200Response) GetScheduleSettings() ScheduleSettingsWriteable`

GetScheduleSettings returns the ScheduleSettings field if non-nil, zero value otherwise.

### GetScheduleSettingsOk

`func (o *GetGrant200Response) GetScheduleSettingsOk() (*ScheduleSettingsWriteable, bool)`

GetScheduleSettingsOk returns a tuple with the ScheduleSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleSettings

`func (o *GetGrant200Response) SetScheduleSettings(v ScheduleSettingsWriteable)`

SetScheduleSettings sets ScheduleSettings field to given value.

### HasScheduleSettings

`func (o *GetGrant200Response) HasScheduleSettings() bool`

HasScheduleSettings returns a boolean if a field has been set.

### GetLinks

`func (o *GetGrant200Response) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetGrant200Response) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetGrant200Response) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetStatus

`func (o *GetGrant200Response) GetStatus() GrantStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGrant200Response) GetStatusOk() (*GrantStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGrant200Response) SetStatus(v GrantStatus)`

SetStatus sets Status field to given value.


### GetMetadata

`func (o *GetGrant200Response) GetMetadata() GrantMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetGrant200Response) GetMetadataOk() (*GrantMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetGrant200Response) SetMetadata(v GrantMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetGrant200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *GetGrant200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGrant200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGrant200Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *GetGrant200Response) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetGrant200Response) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetGrant200Response) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *GetGrant200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetGrant200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetGrant200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *GetGrant200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGrant200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGrant200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *GetGrant200Response) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetGrant200Response) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetGrant200Response) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


