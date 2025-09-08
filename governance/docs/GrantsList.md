# GrantsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GrantFull**](GrantFull.md) |  | [optional] 
**Links** | Pointer to [**GrantListLinks**](GrantListLinks.md) |  | [optional] 

## Methods

### NewGrantsList

`func NewGrantsList() *GrantsList`

NewGrantsList instantiates a new GrantsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantsListWithDefaults

`func NewGrantsListWithDefaults() *GrantsList`

NewGrantsListWithDefaults instantiates a new GrantsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GrantsList) GetData() []GrantFull`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GrantsList) GetDataOk() (*[]GrantFull, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GrantsList) SetData(v []GrantFull)`

SetData sets Data field to given value.

### HasData

`func (o *GrantsList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *GrantsList) GetLinks() GrantListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GrantsList) GetLinksOk() (*GrantListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GrantsList) SetLinks(v GrantListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GrantsList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


