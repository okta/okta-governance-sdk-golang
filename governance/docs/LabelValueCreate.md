# LabelValueCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the label value | 
**Metadata** | Pointer to [**LabelMetadata**](LabelMetadata.md) |  | [optional] 

## Methods

### NewLabelValueCreate

`func NewLabelValueCreate(name string, ) *LabelValueCreate`

NewLabelValueCreate instantiates a new LabelValueCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelValueCreateWithDefaults

`func NewLabelValueCreateWithDefaults() *LabelValueCreate`

NewLabelValueCreateWithDefaults instantiates a new LabelValueCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LabelValueCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabelValueCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabelValueCreate) SetName(v string)`

SetName sets Name field to given value.


### GetMetadata

`func (o *LabelValueCreate) GetMetadata() LabelMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LabelValueCreate) GetMetadataOk() (*LabelMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LabelValueCreate) SetMetadata(v LabelMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LabelValueCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


