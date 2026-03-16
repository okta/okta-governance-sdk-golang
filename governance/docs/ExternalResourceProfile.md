# ExternalResourceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Okta resource ID | 
**Name** | **string** | The display name for the resource | 
**Description** | Pointer to **string** | The description of the resource | [optional] 
**Parent** | Pointer to [**ExternalResourceProfileParent**](ExternalResourceProfileParent.md) |  | [optional] 
**Label** | Pointer to **string** | The label of the Okta resource | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of logo resources | [optional] 

## Methods

### NewExternalResourceProfile

`func NewExternalResourceProfile(id string, name string, ) *ExternalResourceProfile`

NewExternalResourceProfile instantiates a new ExternalResourceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalResourceProfileWithDefaults

`func NewExternalResourceProfileWithDefaults() *ExternalResourceProfile`

NewExternalResourceProfileWithDefaults instantiates a new ExternalResourceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalResourceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalResourceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalResourceProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExternalResourceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalResourceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalResourceProfile) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExternalResourceProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalResourceProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalResourceProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalResourceProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *ExternalResourceProfile) GetParent() ExternalResourceProfileParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ExternalResourceProfile) GetParentOk() (*ExternalResourceProfileParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ExternalResourceProfile) SetParent(v ExternalResourceProfileParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ExternalResourceProfile) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetLabel

`func (o *ExternalResourceProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ExternalResourceProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ExternalResourceProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ExternalResourceProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLogo

`func (o *ExternalResourceProfile) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ExternalResourceProfile) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ExternalResourceProfile) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ExternalResourceProfile) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


