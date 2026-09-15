# CollectionFullWithFilterContextV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of a resource collection | 
**Description** | Pointer to **string** | The human-readable description | [optional] 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Orn** | **string** | The &#x60;id&#x60; of the collection in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn). | 
**Counts** | Pointer to [**CollectionCountsV2**](CollectionCountsV2.md) |  | [optional] 
**ResourceRelationship** | Pointer to [**CollectionResourceRelationship**](CollectionResourceRelationship.md) |  | [optional] 
**PrincipalAssignment** | Pointer to [**CollectionPrincipalAssignment**](CollectionPrincipalAssignment.md) |  | [optional] 

## Methods

### NewCollectionFullWithFilterContextV2

`func NewCollectionFullWithFilterContextV2(name string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links CollectionLinks, orn string, ) *CollectionFullWithFilterContextV2`

NewCollectionFullWithFilterContextV2 instantiates a new CollectionFullWithFilterContextV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFullWithFilterContextV2WithDefaults

`func NewCollectionFullWithFilterContextV2WithDefaults() *CollectionFullWithFilterContextV2`

NewCollectionFullWithFilterContextV2WithDefaults instantiates a new CollectionFullWithFilterContextV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CollectionFullWithFilterContextV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionFullWithFilterContextV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionFullWithFilterContextV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CollectionFullWithFilterContextV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionFullWithFilterContextV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionFullWithFilterContextV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionFullWithFilterContextV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CollectionFullWithFilterContextV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionFullWithFilterContextV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionFullWithFilterContextV2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *CollectionFullWithFilterContextV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CollectionFullWithFilterContextV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CollectionFullWithFilterContextV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *CollectionFullWithFilterContextV2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CollectionFullWithFilterContextV2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CollectionFullWithFilterContextV2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *CollectionFullWithFilterContextV2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CollectionFullWithFilterContextV2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CollectionFullWithFilterContextV2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *CollectionFullWithFilterContextV2) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *CollectionFullWithFilterContextV2) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *CollectionFullWithFilterContextV2) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *CollectionFullWithFilterContextV2) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionFullWithFilterContextV2) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionFullWithFilterContextV2) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetOrn

`func (o *CollectionFullWithFilterContextV2) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *CollectionFullWithFilterContextV2) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *CollectionFullWithFilterContextV2) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetCounts

`func (o *CollectionFullWithFilterContextV2) GetCounts() CollectionCountsV2`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *CollectionFullWithFilterContextV2) GetCountsOk() (*CollectionCountsV2, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *CollectionFullWithFilterContextV2) SetCounts(v CollectionCountsV2)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *CollectionFullWithFilterContextV2) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetResourceRelationship

`func (o *CollectionFullWithFilterContextV2) GetResourceRelationship() CollectionResourceRelationship`

GetResourceRelationship returns the ResourceRelationship field if non-nil, zero value otherwise.

### GetResourceRelationshipOk

`func (o *CollectionFullWithFilterContextV2) GetResourceRelationshipOk() (*CollectionResourceRelationship, bool)`

GetResourceRelationshipOk returns a tuple with the ResourceRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRelationship

`func (o *CollectionFullWithFilterContextV2) SetResourceRelationship(v CollectionResourceRelationship)`

SetResourceRelationship sets ResourceRelationship field to given value.

### HasResourceRelationship

`func (o *CollectionFullWithFilterContextV2) HasResourceRelationship() bool`

HasResourceRelationship returns a boolean if a field has been set.

### GetPrincipalAssignment

`func (o *CollectionFullWithFilterContextV2) GetPrincipalAssignment() CollectionPrincipalAssignment`

GetPrincipalAssignment returns the PrincipalAssignment field if non-nil, zero value otherwise.

### GetPrincipalAssignmentOk

`func (o *CollectionFullWithFilterContextV2) GetPrincipalAssignmentOk() (*CollectionPrincipalAssignment, bool)`

GetPrincipalAssignmentOk returns a tuple with the PrincipalAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalAssignment

`func (o *CollectionFullWithFilterContextV2) SetPrincipalAssignment(v CollectionPrincipalAssignment)`

SetPrincipalAssignment sets PrincipalAssignment field to given value.

### HasPrincipalAssignment

`func (o *CollectionFullWithFilterContextV2) HasPrincipalAssignment() bool`

HasPrincipalAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


