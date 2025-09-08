# AssignResourceLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOrns** | **[]string** | Resources assigned to labels | 
**LabelValueIds** | **[]string** | Labels assigned to resources in the &#x60;resourceOrns&#x60; list | 

## Methods

### NewAssignResourceLabels

`func NewAssignResourceLabels(resourceOrns []string, labelValueIds []string, ) *AssignResourceLabels`

NewAssignResourceLabels instantiates a new AssignResourceLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignResourceLabelsWithDefaults

`func NewAssignResourceLabelsWithDefaults() *AssignResourceLabels`

NewAssignResourceLabelsWithDefaults instantiates a new AssignResourceLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOrns

`func (o *AssignResourceLabels) GetResourceOrns() []string`

GetResourceOrns returns the ResourceOrns field if non-nil, zero value otherwise.

### GetResourceOrnsOk

`func (o *AssignResourceLabels) GetResourceOrnsOk() (*[]string, bool)`

GetResourceOrnsOk returns a tuple with the ResourceOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrns

`func (o *AssignResourceLabels) SetResourceOrns(v []string)`

SetResourceOrns sets ResourceOrns field to given value.


### GetLabelValueIds

`func (o *AssignResourceLabels) GetLabelValueIds() []string`

GetLabelValueIds returns the LabelValueIds field if non-nil, zero value otherwise.

### GetLabelValueIdsOk

`func (o *AssignResourceLabels) GetLabelValueIdsOk() (*[]string, bool)`

GetLabelValueIdsOk returns a tuple with the LabelValueIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelValueIds

`func (o *AssignResourceLabels) SetLabelValueIds(v []string)`

SetLabelValueIds sets LabelValueIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


