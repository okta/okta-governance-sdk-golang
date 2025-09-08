# RuleConflictResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOrn** | Pointer to **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 

## Methods

### NewRuleConflictResource

`func NewRuleConflictResource() *RuleConflictResource`

NewRuleConflictResource instantiates a new RuleConflictResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConflictResourceWithDefaults

`func NewRuleConflictResourceWithDefaults() *RuleConflictResource`

NewRuleConflictResourceWithDefaults instantiates a new RuleConflictResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOrn

`func (o *RuleConflictResource) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *RuleConflictResource) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *RuleConflictResource) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.

### HasResourceOrn

`func (o *RuleConflictResource) HasResourceOrn() bool`

HasResourceOrn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


