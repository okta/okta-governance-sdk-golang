# IndirectAccessList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IndirectAccess**](IndirectAccess.md) | Array of indirect access | [optional] 
**Links** | Pointer to [**ListLinks**](ListLinks.md) |  | [optional] 

## Methods

### NewIndirectAccessList

`func NewIndirectAccessList() *IndirectAccessList`

NewIndirectAccessList instantiates a new IndirectAccessList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndirectAccessListWithDefaults

`func NewIndirectAccessListWithDefaults() *IndirectAccessList`

NewIndirectAccessListWithDefaults instantiates a new IndirectAccessList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IndirectAccessList) GetData() []IndirectAccess`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IndirectAccessList) GetDataOk() (*[]IndirectAccess, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IndirectAccessList) SetData(v []IndirectAccess)`

SetData sets Data field to given value.

### HasData

`func (o *IndirectAccessList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *IndirectAccessList) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IndirectAccessList) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IndirectAccessList) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IndirectAccessList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


