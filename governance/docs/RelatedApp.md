# RelatedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Okta resource instance ID | [optional] 
**Name** | Pointer to **string** | The name of the Okta resource | [optional] 
**Label** | Pointer to **string** | The label of the Okta app. Only populated for &#x60;APPLICATION&#x60; type resources. | [optional] 
**Description** | Pointer to **string** | The description of the resource | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of resource logo resources | [optional] 
**InCollection** | Pointer to **bool** | Indicates whether this app is also a resource in the current collection. Only present for &#x60;GROUP&#x60; type resources. | [optional] 

## Methods

### NewRelatedApp

`func NewRelatedApp() *RelatedApp`

NewRelatedApp instantiates a new RelatedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedAppWithDefaults

`func NewRelatedAppWithDefaults() *RelatedApp`

NewRelatedAppWithDefaults instantiates a new RelatedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedApp) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedApp) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedApp) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelatedApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelatedApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelatedApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelatedApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *RelatedApp) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RelatedApp) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RelatedApp) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RelatedApp) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *RelatedApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelatedApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelatedApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelatedApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogo

`func (o *RelatedApp) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *RelatedApp) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *RelatedApp) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *RelatedApp) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetInCollection

`func (o *RelatedApp) GetInCollection() bool`

GetInCollection returns the InCollection field if non-nil, zero value otherwise.

### GetInCollectionOk

`func (o *RelatedApp) GetInCollectionOk() (*bool, bool)`

GetInCollectionOk returns a tuple with the InCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInCollection

`func (o *RelatedApp) SetInCollection(v bool)`

SetInCollection sets InCollection field to given value.

### HasInCollection

`func (o *RelatedApp) HasInCollection() bool`

HasInCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


