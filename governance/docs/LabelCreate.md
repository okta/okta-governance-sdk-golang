# LabelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the label category | 
**Values** | [**[]LabelValueCreate**](LabelValueCreate.md) | List of label values | 

## Methods

### NewLabelCreate

`func NewLabelCreate(name string, values []LabelValueCreate, ) *LabelCreate`

NewLabelCreate instantiates a new LabelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelCreateWithDefaults

`func NewLabelCreateWithDefaults() *LabelCreate`

NewLabelCreateWithDefaults instantiates a new LabelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LabelCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelCreate) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *LabelCreate) GetValues() []LabelValueCreate`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LabelCreate) GetValuesOk() (*[]LabelValueCreate, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LabelCreate) SetValues(v []LabelValueCreate)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


