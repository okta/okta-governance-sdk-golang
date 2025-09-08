# CollectionResourcePropertiesReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | The Okta &#x60;app.id&#x60; of the resource  See the [List applications](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Application/#tag/Application/operation/listApplications) endpoint to retrieve application IDs.  | [optional] 
**Links** | Pointer to [**CollectionResourceLinks**](CollectionResourceLinks.md) |  | [optional] 

## Methods

### NewCollectionResourcePropertiesReadOnly

`func NewCollectionResourcePropertiesReadOnly() *CollectionResourcePropertiesReadOnly`

NewCollectionResourcePropertiesReadOnly instantiates a new CollectionResourcePropertiesReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourcePropertiesReadOnlyWithDefaults

`func NewCollectionResourcePropertiesReadOnlyWithDefaults() *CollectionResourcePropertiesReadOnly`

NewCollectionResourcePropertiesReadOnlyWithDefaults instantiates a new CollectionResourcePropertiesReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *CollectionResourcePropertiesReadOnly) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *CollectionResourcePropertiesReadOnly) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *CollectionResourcePropertiesReadOnly) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *CollectionResourcePropertiesReadOnly) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetLinks

`func (o *CollectionResourcePropertiesReadOnly) GetLinks() CollectionResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CollectionResourcePropertiesReadOnly) GetLinksOk() (*CollectionResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CollectionResourcePropertiesReadOnly) SetLinks(v CollectionResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CollectionResourcePropertiesReadOnly) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


