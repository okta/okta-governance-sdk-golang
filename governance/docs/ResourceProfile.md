# ResourceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Okta application instance &#x60;id&#x60; | [optional] 
**Name** | Pointer to **string** | The name of the Okta application | [optional] 
**Label** | Pointer to **string** | The label of the Okta application | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of app logo resources | [optional] 

## Methods

### NewResourceProfile

`func NewResourceProfile() *ResourceProfile`

NewResourceProfile instantiates a new ResourceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceProfileWithDefaults

`func NewResourceProfileWithDefaults() *ResourceProfile`

NewResourceProfileWithDefaults instantiates a new ResourceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *ResourceProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ResourceProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ResourceProfile) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ResourceProfile) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLogo

`func (o *ResourceProfile) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ResourceProfile) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ResourceProfile) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ResourceProfile) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


