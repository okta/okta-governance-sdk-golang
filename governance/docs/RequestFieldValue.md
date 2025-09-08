# RequestFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of a &#x60;requesterField&#x60;.  For &#x60;OKTA&#x60; request approval, this is typically the ID of an input field defined in an approval sequence.  For &#x60;EXTERNAL&#x60; request approval, this is the ID of an input field relevant to the external approval system.  The following are reserved identifiers: - &#x60;ACCESS_DURATION&#x60;: An identifier used to specify the duration of the access being requested. (**Note:**   This identifier is to be grandfathered. The new name will be &#x60;OKTA_ACCESS_DURATION&#x60;.) - &#x60;OKTA_REQUESTED_FOR&#x60;: An identifier used to specify who the resource is being requested for - &#x60;OKTA_ACCESS_DURATION&#x60;: An identifier used to specify the duration of the requested access. If a request   field is included in this ID, the field value must be a &#x60;DURATION&#x60; expression in ISO format. The maximum   length of this value is 5 characters. For example, &#x60;P1D&#x60; represents a duration of one day. - Any identifiers with &#x60;OKTA_&#x60; prefix are reserved for future use.  | 
**Label** | Pointer to **string** | A human-readable description of &#x60;requesterField&#x60;. It&#39;s used for display purposes and is optional | [optional] 
**Type** | Pointer to [**RequestFieldType**](RequestFieldType.md) |  | [optional] 
**Value** | Pointer to **string** | The value of &#x60;requesterField&#x60;, which depends on the type of the field: - &#x60;TEXT&#x60;: A multi-line string - &#x60;SELECT&#x60;: An option predefined by the approval system - &#x60;ISO_DATE&#x60;: An ISO formatted date string. for example, &#x60;2022-05-05T14:15:22Z&#x60; - &#x60;DURATION&#x60;: An ISO format duration expression with a maximum length of 5 characters. For example,   &#x60;P1D&#x60; indicates a duration of one day. - &#x60;OKTA_USER_ID&#x60;: Okta User ID  If the field type is one of the previously listed, this property is required.  | [optional] 
**Values** | Pointer to **[]string** | The values of &#x60;requesterField&#x60; with the type &#x60;MULTISELECT&#x60;.  If the field type is &#x60;MULTISELECT&#x60;, this property is required.  | [optional] 

## Methods

### NewRequestFieldValue

`func NewRequestFieldValue(id string, ) *RequestFieldValue`

NewRequestFieldValue instantiates a new RequestFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFieldValueWithDefaults

`func NewRequestFieldValueWithDefaults() *RequestFieldValue`

NewRequestFieldValueWithDefaults instantiates a new RequestFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestFieldValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestFieldValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestFieldValue) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *RequestFieldValue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RequestFieldValue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RequestFieldValue) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RequestFieldValue) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *RequestFieldValue) GetType() RequestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestFieldValue) GetTypeOk() (*RequestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestFieldValue) SetType(v RequestFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *RequestFieldValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *RequestFieldValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestFieldValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestFieldValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RequestFieldValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *RequestFieldValue) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RequestFieldValue) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RequestFieldValue) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RequestFieldValue) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


