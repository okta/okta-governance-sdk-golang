# LabelValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelValueId** | **string** | The ID of a label value | 
**Name** | **string** | The label value | 
**Metadata** | Pointer to [**LabelMetadata**](LabelMetadata.md) |  | [optional] 

## Methods

### NewLabelValue

`func NewLabelValue(labelValueId string, name string, ) *LabelValue`

NewLabelValue instantiates a new LabelValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelValueWithDefaults

`func NewLabelValueWithDefaults() *LabelValue`

NewLabelValueWithDefaults instantiates a new LabelValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelValueId

`func (o *LabelValue) GetLabelValueId() string`

GetLabelValueId returns the LabelValueId field if non-nil, zero value otherwise.

### GetLabelValueIdOk

`func (o *LabelValue) GetLabelValueIdOk() (*string, bool)`

GetLabelValueIdOk returns a tuple with the LabelValueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelValueId

`func (o *LabelValue) SetLabelValueId(v string)`

SetLabelValueId sets LabelValueId field to given value.


### GetName

`func (o *LabelValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelValue) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *LabelValue) GetMetadata() LabelMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LabelValue) GetMetadataOk() (*LabelMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LabelValue) SetMetadata(v LabelMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LabelValue) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


