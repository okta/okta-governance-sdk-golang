# LinkNext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Link name | [optional] 
**Rel** | Pointer to **string** | Link relation | [optional] 
**Href** | **string** | Link URI | 
**Type** | Pointer to **string** | The media type of the link. If omitted, it&#39;s implicitly &#x60;application/json&#x60;. | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Templated** | Pointer to **bool** | Indicates whether the link object&#39;s &#x60;href&#x60; property is a URI template | [optional] 
**Hints** | Pointer to **map[string][]string** | Link hints | [optional] 
**Title** | Pointer to **string** | Link title or label | [optional] 

## Methods

### NewLinkNext

`func NewLinkNext(href string, ) *LinkNext`

NewLinkNext instantiates a new LinkNext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkNextWithDefaults

`func NewLinkNextWithDefaults() *LinkNext`

NewLinkNextWithDefaults instantiates a new LinkNext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LinkNext) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkNext) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkNext) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinkNext) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRel

`func (o *LinkNext) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *LinkNext) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *LinkNext) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *LinkNext) HasRel() bool`

HasRel returns a boolean if a field has been set.

### GetHref

`func (o *LinkNext) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkNext) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkNext) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *LinkNext) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkNext) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkNext) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkNext) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMethod

`func (o *LinkNext) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LinkNext) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LinkNext) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LinkNext) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetTemplated

`func (o *LinkNext) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *LinkNext) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *LinkNext) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *LinkNext) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetHints

`func (o *LinkNext) GetHints() map[string][]string`

GetHints returns the Hints field if non-nil, zero value otherwise.

### GetHintsOk

`func (o *LinkNext) GetHintsOk() (*map[string][]string, bool)`

GetHintsOk returns a tuple with the Hints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHints

`func (o *LinkNext) SetHints(v map[string][]string)`

SetHints sets Hints field to given value.

### HasHints

`func (o *LinkNext) HasHints() bool`

HasHints returns a boolean if a field has been set.

### GetTitle

`func (o *LinkNext) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LinkNext) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LinkNext) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LinkNext) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


