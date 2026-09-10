# ResourceProfileV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Okta resource instance ID | [optional] 
**Name** | Pointer to **string** | The name of the Okta resource | [optional] 
**Label** | Pointer to **string** | The label of the Okta app. Only populated for &#x60;APPLICATION&#x60; type resources. | [optional] 
**Description** | Pointer to **string** | The description of the resource | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of resource logo resources | [optional] 

## Methods

### NewResourceProfileV2

`func NewResourceProfileV2() *ResourceProfileV2`

NewResourceProfileV2 instantiates a new ResourceProfileV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceProfileV2WithDefaults

`func NewResourceProfileV2WithDefaults() *ResourceProfileV2`

NewResourceProfileV2WithDefaults instantiates a new ResourceProfileV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceProfileV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceProfileV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceProfileV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceProfileV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceProfileV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceProfileV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceProfileV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceProfileV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *ResourceProfileV2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ResourceProfileV2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ResourceProfileV2) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ResourceProfileV2) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceProfileV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceProfileV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceProfileV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceProfileV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *ResourceProfileV2) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ResourceProfileV2) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ResourceProfileV2) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ResourceProfileV2) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


