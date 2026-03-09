# RuleConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | Pointer to **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | [optional] 
**ResourceOrn** | Pointer to **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 
**RuleId** | Pointer to **string** | Unique identifier for rule object | [optional] 
**RuleName** | Pointer to **string** | The name of a resource rule causing the conflict | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object | [optional] 
**Description** | Pointer to **string** | Description for the risk rule | [optional] 
**Type** | Pointer to **string** | Risk rule type | [optional] 
**ConflictCriteria** | Pointer to [**ConflictCriteria**](ConflictCriteria.md) |  | [optional] 

## Methods

### NewRuleConflict

`func NewRuleConflict() *RuleConflict`

NewRuleConflict instantiates a new RuleConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConflictWithDefaults

`func NewRuleConflictWithDefaults() *RuleConflict`

NewRuleConflictWithDefaults instantiates a new RuleConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *RuleConflict) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *RuleConflict) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *RuleConflict) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.

### HasPrincipalOrn

`func (o *RuleConflict) HasPrincipalOrn() bool`

HasPrincipalOrn returns a boolean if a field has been set.

### GetResourceOrn

`func (o *RuleConflict) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *RuleConflict) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *RuleConflict) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.

### HasResourceOrn

`func (o *RuleConflict) HasResourceOrn() bool`

HasResourceOrn returns a boolean if a field has been set.

### GetRuleId

`func (o *RuleConflict) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *RuleConflict) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *RuleConflict) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *RuleConflict) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleName

`func (o *RuleConflict) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *RuleConflict) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *RuleConflict) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *RuleConflict) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetId

`func (o *RuleConflict) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleConflict) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleConflict) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RuleConflict) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *RuleConflict) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleConflict) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleConflict) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleConflict) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RuleConflict) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleConflict) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleConflict) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RuleConflict) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConflictCriteria

`func (o *RuleConflict) GetConflictCriteria() ConflictCriteria`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *RuleConflict) GetConflictCriteriaOk() (*ConflictCriteria, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *RuleConflict) SetConflictCriteria(v ConflictCriteria)`

SetConflictCriteria sets ConflictCriteria field to given value.

### HasConflictCriteria

`func (o *RuleConflict) HasConflictCriteria() bool`

HasConflictCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


