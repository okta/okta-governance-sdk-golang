# RiskRuleCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the risk rule criteria | [optional] 
**Operation** | Pointer to **string** | Operation to be performed on the criteria value | [optional] 
**Value** | Pointer to [**RiskRuleCriteriaValue**](RiskRuleCriteriaValue.md) |  | [optional] 

## Methods

### NewRiskRuleCriteria

`func NewRiskRuleCriteria() *RiskRuleCriteria`

NewRiskRuleCriteria instantiates a new RiskRuleCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleCriteriaWithDefaults

`func NewRiskRuleCriteriaWithDefaults() *RiskRuleCriteria`

NewRiskRuleCriteriaWithDefaults instantiates a new RiskRuleCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RiskRuleCriteria) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRuleCriteria) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRuleCriteria) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskRuleCriteria) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *RiskRuleCriteria) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RiskRuleCriteria) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RiskRuleCriteria) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RiskRuleCriteria) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *RiskRuleCriteria) GetValue() RiskRuleCriteriaValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskRuleCriteria) GetValueOk() (*RiskRuleCriteriaValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskRuleCriteria) SetValue(v RiskRuleCriteriaValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskRuleCriteria) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


