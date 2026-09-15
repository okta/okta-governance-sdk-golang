# CollectionResourceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementManagementEnabled** | **bool** | Indicates whether Entitlement Management is enabled for this resource | 
**HasPushMapping** | Pointer to **bool** | Indicates whether this group resource has push group mappings to applications (is a push group). Only present when &#x60;type&#x60; is &#x60;GROUP&#x60;. | [optional] 
**Type** | **string** | The type of resource | 
**Counts** | Pointer to [**CollectionResourceCounts**](CollectionResourceCounts.md) |  | [optional] 

## Methods

### NewCollectionResourceConfiguration

`func NewCollectionResourceConfiguration(entitlementManagementEnabled bool, type_ string, ) *CollectionResourceConfiguration`

NewCollectionResourceConfiguration instantiates a new CollectionResourceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceConfigurationWithDefaults

`func NewCollectionResourceConfigurationWithDefaults() *CollectionResourceConfiguration`

NewCollectionResourceConfigurationWithDefaults instantiates a new CollectionResourceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementManagementEnabled

`func (o *CollectionResourceConfiguration) GetEntitlementManagementEnabled() bool`

GetEntitlementManagementEnabled returns the EntitlementManagementEnabled field if non-nil, zero value otherwise.

### GetEntitlementManagementEnabledOk

`func (o *CollectionResourceConfiguration) GetEntitlementManagementEnabledOk() (*bool, bool)`

GetEntitlementManagementEnabledOk returns a tuple with the EntitlementManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementManagementEnabled

`func (o *CollectionResourceConfiguration) SetEntitlementManagementEnabled(v bool)`

SetEntitlementManagementEnabled sets EntitlementManagementEnabled field to given value.


### GetHasPushMapping

`func (o *CollectionResourceConfiguration) GetHasPushMapping() bool`

GetHasPushMapping returns the HasPushMapping field if non-nil, zero value otherwise.

### GetHasPushMappingOk

`func (o *CollectionResourceConfiguration) GetHasPushMappingOk() (*bool, bool)`

GetHasPushMappingOk returns a tuple with the HasPushMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPushMapping

`func (o *CollectionResourceConfiguration) SetHasPushMapping(v bool)`

SetHasPushMapping sets HasPushMapping field to given value.

### HasHasPushMapping

`func (o *CollectionResourceConfiguration) HasHasPushMapping() bool`

HasHasPushMapping returns a boolean if a field has been set.

### GetType

`func (o *CollectionResourceConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionResourceConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionResourceConfiguration) SetType(v string)`

SetType sets Type field to given value.


### GetCounts

`func (o *CollectionResourceConfiguration) GetCounts() CollectionResourceCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *CollectionResourceConfiguration) GetCountsOk() (*CollectionResourceCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *CollectionResourceConfiguration) SetCounts(v CollectionResourceCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *CollectionResourceConfiguration) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


