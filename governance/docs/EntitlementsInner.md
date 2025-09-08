# EntitlementsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The entitlement id | 
**IncludeAllValues** | Pointer to **bool** | whether to include all entitlement values. If &#x60;false&#x60; we must provide the &#x60;values&#x60; property | [optional] 
**Values** | Pointer to [**[]EntitlementValue**](EntitlementValue.md) | entitlement value ids | [optional] 

## Methods

### NewEntitlementsInner

`func NewEntitlementsInner(id string, ) *EntitlementsInner`

NewEntitlementsInner instantiates a new EntitlementsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsInnerWithDefaults

`func NewEntitlementsInnerWithDefaults() *EntitlementsInner`

NewEntitlementsInnerWithDefaults instantiates a new EntitlementsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementsInner) SetId(v string)`

SetId sets Id field to given value.


### GetIncludeAllValues

`func (o *EntitlementsInner) GetIncludeAllValues() bool`

GetIncludeAllValues returns the IncludeAllValues field if non-nil, zero value otherwise.

### GetIncludeAllValuesOk

`func (o *EntitlementsInner) GetIncludeAllValuesOk() (*bool, bool)`

GetIncludeAllValuesOk returns a tuple with the IncludeAllValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllValues

`func (o *EntitlementsInner) SetIncludeAllValues(v bool)`

SetIncludeAllValues sets IncludeAllValues field to given value.

### HasIncludeAllValues

`func (o *EntitlementsInner) HasIncludeAllValues() bool`

HasIncludeAllValues returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementsInner) GetValues() []EntitlementValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementsInner) GetValuesOk() (*[]EntitlementValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementsInner) SetValues(v []EntitlementValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementsInner) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


