# CriteriaValueCreatableEntitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator. Must be &#x60;ENTITLEMENTS&#x60;. | 
**Value** | [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | 

## Methods

### NewCriteriaValueCreatableEntitlements

`func NewCriteriaValueCreatableEntitlements(type_ string, value []EntitlementCreatable, ) *CriteriaValueCreatableEntitlements`

NewCriteriaValueCreatableEntitlements instantiates a new CriteriaValueCreatableEntitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriteriaValueCreatableEntitlementsWithDefaults

`func NewCriteriaValueCreatableEntitlementsWithDefaults() *CriteriaValueCreatableEntitlements`

NewCriteriaValueCreatableEntitlementsWithDefaults instantiates a new CriteriaValueCreatableEntitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CriteriaValueCreatableEntitlements) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CriteriaValueCreatableEntitlements) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CriteriaValueCreatableEntitlements) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CriteriaValueCreatableEntitlements) GetValue() []EntitlementCreatable`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CriteriaValueCreatableEntitlements) GetValueOk() (*[]EntitlementCreatable, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CriteriaValueCreatableEntitlements) SetValue(v []EntitlementCreatable)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


