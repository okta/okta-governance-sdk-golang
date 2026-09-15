# ResolveMyTaskV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Specify the task resolution value. Required for &#x60;APPROVAL&#x60; and &#x60;QUESTION&#x60; type tasks.   * For &#x60;APPROVAL&#x60; type tasks, the only supported values are &#x60;APPROVED&#x60; or &#x60;DENIED&#x60;.   * For &#x60;QUESTION&#x60; type tasks, the value is the answer to the question, and can be any string.   * For custom &#x60;TODO&#x60; type tasks, you don&#39;t need to specify a resolution value.  | [optional] 

## Methods

### NewResolveMyTaskV2Request

`func NewResolveMyTaskV2Request() *ResolveMyTaskV2Request`

NewResolveMyTaskV2Request instantiates a new ResolveMyTaskV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolveMyTaskV2RequestWithDefaults

`func NewResolveMyTaskV2RequestWithDefaults() *ResolveMyTaskV2Request`

NewResolveMyTaskV2RequestWithDefaults instantiates a new ResolveMyTaskV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ResolveMyTaskV2Request) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResolveMyTaskV2Request) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResolveMyTaskV2Request) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResolveMyTaskV2Request) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


