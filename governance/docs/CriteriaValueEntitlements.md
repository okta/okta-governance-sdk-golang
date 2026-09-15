# CriteriaValueEntitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator. Must be &#x60;ENTITLEMENTS&#x60;. | 
**Value** | [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | 

## Methods

### NewCriteriaValueEntitlements

`func NewCriteriaValueEntitlements(type_ string, value []EntitlementFull, ) *CriteriaValueEntitlements`

NewCriteriaValueEntitlements instantiates a new CriteriaValueEntitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriteriaValueEntitlementsWithDefaults

`func NewCriteriaValueEntitlementsWithDefaults() *CriteriaValueEntitlements`

NewCriteriaValueEntitlementsWithDefaults instantiates a new CriteriaValueEntitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CriteriaValueEntitlements) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CriteriaValueEntitlements) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CriteriaValueEntitlements) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CriteriaValueEntitlements) GetValue() []EntitlementFull`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CriteriaValueEntitlements) GetValueOk() (*[]EntitlementFull, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CriteriaValueEntitlements) SetValue(v []EntitlementFull)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


