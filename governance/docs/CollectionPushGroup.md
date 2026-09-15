# CollectionPushGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Okta Group ID for the push group | [optional] 
**Name** | Pointer to **string** | The name of the push group | [optional] 
**Description** | Pointer to **string** | The description of the push group | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of push group logo resources | [optional] 
**ResourceId** | **string** | The unique resource ID for this resource (app, group, or push group). Use this identifier to reference the resource in collection-resource API calls, such as &#x60;GET&#x60;/&#x60;PUT&#x60;/&#x60;DELETE /v2/collections/{collectionId}/resources/{resourceId}&#x60;.  | 
**ResourceOrn** | **string** | The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.  | 

## Methods

### NewCollectionPushGroup

`func NewCollectionPushGroup(resourceId string, resourceOrn string, ) *CollectionPushGroup`

NewCollectionPushGroup instantiates a new CollectionPushGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionPushGroupWithDefaults

`func NewCollectionPushGroupWithDefaults() *CollectionPushGroup`

NewCollectionPushGroupWithDefaults instantiates a new CollectionPushGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionPushGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionPushGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionPushGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CollectionPushGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CollectionPushGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionPushGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionPushGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionPushGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CollectionPushGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionPushGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionPushGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionPushGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *CollectionPushGroup) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CollectionPushGroup) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CollectionPushGroup) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CollectionPushGroup) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetResourceId

`func (o *CollectionPushGroup) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionPushGroup) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionPushGroup) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceOrn

`func (o *CollectionPushGroup) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionPushGroup) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionPushGroup) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


