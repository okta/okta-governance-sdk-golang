# WritableStandardProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Writable unique key on Create. Not modifiable on update. | [optional] 
**Description** | Pointer to **string** | Human readable description. | [optional] 

## Methods

### NewWritableStandardProperties

`func NewWritableStandardProperties() *WritableStandardProperties`

NewWritableStandardProperties instantiates a new WritableStandardProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableStandardPropertiesWithDefaults

`func NewWritableStandardPropertiesWithDefaults() *WritableStandardProperties`

NewWritableStandardPropertiesWithDefaults instantiates a new WritableStandardProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritableStandardProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableStandardProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableStandardProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WritableStandardProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WritableStandardProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableStandardProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableStandardProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableStandardProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


