# OktaApplicationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | An Okta app ID.  See [list apps](https://developer.okta.com/docs/reference/api/apps/#list-applications) endpoint for reference on how to retrieve app ids.  | 

## Methods

### NewOktaApplicationResource

`func NewOktaApplicationResource(resourceId string, ) *OktaApplicationResource`

NewOktaApplicationResource instantiates a new OktaApplicationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaApplicationResourceWithDefaults

`func NewOktaApplicationResourceWithDefaults() *OktaApplicationResource`

NewOktaApplicationResourceWithDefaults instantiates a new OktaApplicationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *OktaApplicationResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OktaApplicationResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OktaApplicationResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


