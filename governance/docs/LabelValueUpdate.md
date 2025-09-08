# LabelValueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the label value | [optional] 
**Metadata** | Pointer to [**LabelMetadata**](LabelMetadata.md) |  | [optional] 

## Methods

### NewLabelValueUpdate

`func NewLabelValueUpdate() *LabelValueUpdate`

NewLabelValueUpdate instantiates a new LabelValueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelValueUpdateWithDefaults

`func NewLabelValueUpdateWithDefaults() *LabelValueUpdate`

NewLabelValueUpdateWithDefaults instantiates a new LabelValueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LabelValueUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelValueUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelValueUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LabelValueUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *LabelValueUpdate) GetMetadata() LabelMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LabelValueUpdate) GetMetadataOk() (*LabelMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LabelValueUpdate) SetMetadata(v LabelMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LabelValueUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


