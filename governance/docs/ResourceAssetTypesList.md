# ResourceAssetTypesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourceAssetType**](ResourceAssetType.md) | List of resource asset types for the resource | 
**Links** | [**ListLinks**](ListLinks.md) |  | 

## Methods

### NewResourceAssetTypesList

`func NewResourceAssetTypesList(data []ResourceAssetType, links ListLinks, ) *ResourceAssetTypesList`

NewResourceAssetTypesList instantiates a new ResourceAssetTypesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAssetTypesListWithDefaults

`func NewResourceAssetTypesListWithDefaults() *ResourceAssetTypesList`

NewResourceAssetTypesListWithDefaults instantiates a new ResourceAssetTypesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResourceAssetTypesList) GetData() []ResourceAssetType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceAssetTypesList) GetDataOk() (*[]ResourceAssetType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceAssetTypesList) SetData(v []ResourceAssetType)`

SetData sets Data field to given value.


### GetLinks

`func (o *ResourceAssetTypesList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceAssetTypesList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceAssetTypesList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


