# ResourceLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | Pointer to **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 
**Profile** | Pointer to [**ExternalResourceProfile**](ExternalResourceProfile.md) |  | [optional] 
**Labels** | Pointer to [**[]Label**](Label.md) | List of assigned labels. | [optional] 
**Links** | [**LinkSelf**](LinkSelf.md) |  | 

## Methods

### NewResourceLabel

`func NewResourceLabel(links LinkSelf, ) *ResourceLabel`

NewResourceLabel instantiates a new ResourceLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLabelWithDefaults

`func NewResourceLabelWithDefaults() *ResourceLabel`

NewResourceLabelWithDefaults instantiates a new ResourceLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *ResourceLabel) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceLabel) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceLabel) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ResourceLabel) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProfile

`func (o *ResourceLabel) GetProfile() ExternalResourceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResourceLabel) GetProfileOk() (*ExternalResourceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResourceLabel) SetProfile(v ExternalResourceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ResourceLabel) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetLabels

`func (o *ResourceLabel) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ResourceLabel) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ResourceLabel) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ResourceLabel) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLinks

`func (o *ResourceLabel) GetLinks() LinkSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ResourceLabel) GetLinksOk() (*LinkSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ResourceLabel) SetLinks(v LinkSelf)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


