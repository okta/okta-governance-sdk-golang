# GetentitlementBundle200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Status** | [**EntitlementBundleStatus**](EntitlementBundleStatus.md) |  | 
**Entitlements** | [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | 
**EntitlementsObjectType** | **string** |  | [default to "entitlement-bundle-full-with-entitlements"]
**Links** | [**map[string]Link**](Link.md) |  | 
**Name** | **string** | The unique name of the entitlement bundle | 
**Description** | Pointer to **string** | The human-readable description | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 

## Methods

### NewGetentitlementBundle200Response

`func NewGetentitlementBundle200Response(targetResourceOrn string, target TargetResource, status EntitlementBundleStatus, entitlements []EntitlementFull, entitlementsObjectType string, links map[string]Link, name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *GetentitlementBundle200Response`

NewGetentitlementBundle200Response instantiates a new GetentitlementBundle200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetentitlementBundle200ResponseWithDefaults

`func NewGetentitlementBundle200ResponseWithDefaults() *GetentitlementBundle200Response`

NewGetentitlementBundle200ResponseWithDefaults instantiates a new GetentitlementBundle200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetResourceOrn

`func (o *GetentitlementBundle200Response) GetTargetResourceOrn() string`

GetTargetResourceOrn returns the TargetResourceOrn field if non-nil, zero value otherwise.

### GetTargetResourceOrnOk

`func (o *GetentitlementBundle200Response) GetTargetResourceOrnOk() (*string, bool)`

GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceOrn

`func (o *GetentitlementBundle200Response) SetTargetResourceOrn(v string)`

SetTargetResourceOrn sets TargetResourceOrn field to given value.


### GetTarget

`func (o *GetentitlementBundle200Response) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetentitlementBundle200Response) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetentitlementBundle200Response) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetStatus

`func (o *GetentitlementBundle200Response) GetStatus() EntitlementBundleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetentitlementBundle200Response) GetStatusOk() (*EntitlementBundleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetentitlementBundle200Response) SetStatus(v EntitlementBundleStatus)`

SetStatus sets Status field to given value.


### GetEntitlements

`func (o *GetentitlementBundle200Response) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *GetentitlementBundle200Response) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *GetentitlementBundle200Response) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.


### GetEntitlementsObjectType

`func (o *GetentitlementBundle200Response) GetEntitlementsObjectType() string`

GetEntitlementsObjectType returns the EntitlementsObjectType field if non-nil, zero value otherwise.

### GetEntitlementsObjectTypeOk

`func (o *GetentitlementBundle200Response) GetEntitlementsObjectTypeOk() (*string, bool)`

GetEntitlementsObjectTypeOk returns a tuple with the EntitlementsObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementsObjectType

`func (o *GetentitlementBundle200Response) SetEntitlementsObjectType(v string)`

SetEntitlementsObjectType sets EntitlementsObjectType field to given value.


### GetLinks

`func (o *GetentitlementBundle200Response) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetentitlementBundle200Response) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetentitlementBundle200Response) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.


### GetName

`func (o *GetentitlementBundle200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetentitlementBundle200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetentitlementBundle200Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetentitlementBundle200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetentitlementBundle200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetentitlementBundle200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetentitlementBundle200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GetentitlementBundle200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetentitlementBundle200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetentitlementBundle200Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *GetentitlementBundle200Response) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetentitlementBundle200Response) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetentitlementBundle200Response) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *GetentitlementBundle200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetentitlementBundle200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetentitlementBundle200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *GetentitlementBundle200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetentitlementBundle200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetentitlementBundle200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *GetentitlementBundle200Response) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *GetentitlementBundle200Response) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *GetentitlementBundle200Response) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


