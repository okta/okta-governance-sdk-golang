# TargetGovernanceLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelId** | **string** | The ID of a label | 
**Name** | **string** | Name of the label | 
**Values** | [**[]LabelValue**](LabelValue.md) | List of label values | 

## Methods

### NewTargetGovernanceLabel

`func NewTargetGovernanceLabel(labelId string, name string, values []LabelValue, ) *TargetGovernanceLabel`

NewTargetGovernanceLabel instantiates a new TargetGovernanceLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGovernanceLabelWithDefaults

`func NewTargetGovernanceLabelWithDefaults() *TargetGovernanceLabel`

NewTargetGovernanceLabelWithDefaults instantiates a new TargetGovernanceLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelId

`func (o *TargetGovernanceLabel) GetLabelId() string`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *TargetGovernanceLabel) GetLabelIdOk() (*string, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *TargetGovernanceLabel) SetLabelId(v string)`

SetLabelId sets LabelId field to given value.


### GetName

`func (o *TargetGovernanceLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetGovernanceLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetGovernanceLabel) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *TargetGovernanceLabel) GetValues() []LabelValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TargetGovernanceLabel) GetValuesOk() (*[]LabelValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TargetGovernanceLabel) SetValues(v []LabelValue)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


