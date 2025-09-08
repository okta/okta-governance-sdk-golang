# RiskRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of a resource rule causing a conflict | 
**Description** | Pointer to **string** | The human readable description | [optional] 
**ResourceName** | Pointer to **string** | Human readable name of the resource | [optional] 

## Methods

### NewRiskRule

`func NewRiskRule(name string, ) *RiskRule`

NewRiskRule instantiates a new RiskRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleWithDefaults

`func NewRiskRuleWithDefaults() *RiskRule`

NewRiskRuleWithDefaults instantiates a new RiskRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RiskRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RiskRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceName

`func (o *RiskRule) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *RiskRule) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *RiskRule) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *RiskRule) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


