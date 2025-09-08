# ListResourceLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ResourceLabel**](ResourceLabel.md) | Resources with labels | 
**Links** | [**ListLinks**](ListLinks.md) |  | 

## Methods

### NewListResourceLabels

`func NewListResourceLabels(data []ResourceLabel, links ListLinks, ) *ListResourceLabels`

NewListResourceLabels instantiates a new ListResourceLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResourceLabelsWithDefaults

`func NewListResourceLabelsWithDefaults() *ListResourceLabels`

NewListResourceLabelsWithDefaults instantiates a new ListResourceLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListResourceLabels) GetData() []ResourceLabel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListResourceLabels) GetDataOk() (*[]ResourceLabel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListResourceLabels) SetData(v []ResourceLabel)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListResourceLabels) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListResourceLabels) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListResourceLabels) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


