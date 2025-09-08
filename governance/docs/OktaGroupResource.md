# OktaGroupResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | Okta group ID  &gt; **Note:** See [List all groups](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/listGroups) for reference on how to retrieve group IDs. | 

## Methods

### NewOktaGroupResource

`func NewOktaGroupResource(resourceId string, ) *OktaGroupResource`

NewOktaGroupResource instantiates a new OktaGroupResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaGroupResourceWithDefaults

`func NewOktaGroupResourceWithDefaults() *OktaGroupResource`

NewOktaGroupResourceWithDefaults instantiates a new OktaGroupResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *OktaGroupResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OktaGroupResource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OktaGroupResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


