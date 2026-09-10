# RequestFieldChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the choice. Stable across calls and used as the answer value. | 
**Label** | **string** | Display label for the choice | 
**Description** | Pointer to **string** | Optional description providing additional context for the choice | [optional] 
**Category** | Pointer to **string** | Optional category label for the choice. | [optional] 

## Methods

### NewRequestFieldChoice

`func NewRequestFieldChoice(id string, label string, ) *RequestFieldChoice`

NewRequestFieldChoice instantiates a new RequestFieldChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFieldChoiceWithDefaults

`func NewRequestFieldChoiceWithDefaults() *RequestFieldChoice`

NewRequestFieldChoiceWithDefaults instantiates a new RequestFieldChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestFieldChoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestFieldChoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestFieldChoice) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *RequestFieldChoice) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RequestFieldChoice) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RequestFieldChoice) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *RequestFieldChoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestFieldChoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestFieldChoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestFieldChoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *RequestFieldChoice) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RequestFieldChoice) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RequestFieldChoice) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *RequestFieldChoice) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


