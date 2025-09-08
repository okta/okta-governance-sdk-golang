# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelId** | **string** | The &#x60;id&#x60; property of a label | 
**Name** | **string** | Name of the label | 
**Values** | [**[]LabelValue**](LabelValue.md) | List of label values | 
**Links** | [**LinkSelf**](LinkSelf.md) |  | 

## Methods

### NewLabel

`func NewLabel(labelId string, name string, values []LabelValue, links LinkSelf, ) *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelId

`func (o *Label) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *Label) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *Label) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetName

`func (o *Label) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Label) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Label) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *Label) GetValues() []LabelValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Label) GetValuesOk() (*[]LabelValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Label) SetValues(v []LabelValue)`

SetValues sets Values field to given value.


### GetLinks

`func (o *Label) GetLinks() LinkSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Label) GetLinksOk() (*LinkSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Label) SetLinks(v LinkSelf)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


