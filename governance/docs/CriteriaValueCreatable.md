# CriteriaValueCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the criteria value. Supported type: &#x60;ENTITLEMENTS&#x60; | [optional] 
**Value** | Pointer to [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | [optional] 

## Methods

### NewCriteriaValueCreatable

`func NewCriteriaValueCreatable() *CriteriaValueCreatable`

NewCriteriaValueCreatable instantiates a new CriteriaValueCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriteriaValueCreatableWithDefaults

`func NewCriteriaValueCreatableWithDefaults() *CriteriaValueCreatable`

NewCriteriaValueCreatableWithDefaults instantiates a new CriteriaValueCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CriteriaValueCreatable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CriteriaValueCreatable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CriteriaValueCreatable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CriteriaValueCreatable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CriteriaValueCreatable) GetValue() []EntitlementCreatable`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CriteriaValueCreatable) GetValueOk() (*[]EntitlementCreatable, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CriteriaValueCreatable) SetValue(v []EntitlementCreatable)`

SetValue sets Value field to given value.

### HasValue

`func (o *CriteriaValueCreatable) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


