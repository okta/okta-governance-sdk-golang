# ListLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Label**](Label.md) | All labels on the current page | 
**Links** | [**ListLinks**](ListLinks.md) |  | 

## Methods

### NewListLabels

`func NewListLabels(data []Label, links ListLinks, ) *ListLabels`

NewListLabels instantiates a new ListLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLabelsWithDefaults

`func NewListLabelsWithDefaults() *ListLabels`

NewListLabelsWithDefaults instantiates a new ListLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListLabels) GetData() []Label`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListLabels) GetDataOk() (*[]Label, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListLabels) SetData(v []Label)`

SetData sets Data field to given value.


### GetLinks

`func (o *ListLabels) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListLabels) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListLabels) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


