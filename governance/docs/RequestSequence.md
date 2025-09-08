# RequestSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique idenfitifer of the request sequence | 
**Name** | **string** | Name of the request sequence | 
**Description** | **string** | Description of the request sequence | 
**CompatibleResourceTypes** | Pointer to [**[]CompatibleResourceTypes**](CompatibleResourceTypes.md) |  | [optional] 
**Link** | **string** | Link to edit the request sequence | 

## Methods

### NewRequestSequence

`func NewRequestSequence(id string, name string, description string, link string, ) *RequestSequence`

NewRequestSequence instantiates a new RequestSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSequenceWithDefaults

`func NewRequestSequenceWithDefaults() *RequestSequence`

NewRequestSequenceWithDefaults instantiates a new RequestSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestSequence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestSequence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestSequence) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RequestSequence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestSequence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestSequence) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RequestSequence) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestSequence) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestSequence) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCompatibleResourceTypes

`func (o *RequestSequence) GetCompatibleResourceTypes() []CompatibleResourceTypes`

GetCompatibleResourceTypes returns the CompatibleResourceTypes field if non-nil, zero value otherwise.

### GetCompatibleResourceTypesOk

`func (o *RequestSequence) GetCompatibleResourceTypesOk() (*[]CompatibleResourceTypes, bool)`

GetCompatibleResourceTypesOk returns a tuple with the CompatibleResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleResourceTypes

`func (o *RequestSequence) SetCompatibleResourceTypes(v []CompatibleResourceTypes)`

SetCompatibleResourceTypes sets CompatibleResourceTypes field to given value.

### HasCompatibleResourceTypes

`func (o *RequestSequence) HasCompatibleResourceTypes() bool`

HasCompatibleResourceTypes returns a boolean if a field has been set.

### GetLink

`func (o *RequestSequence) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *RequestSequence) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *RequestSequence) SetLink(v string)`

SetLink sets Link field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


