# EntitlementCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The &#x60;id&#x60; property of an entitlement | [optional] 
**Values** | Pointer to [**[]EntitlementValueCreatable**](EntitlementValueCreatable.md) | Collection of entitlement value identifiers | [optional] 

## Methods

### NewEntitlementCreatable

`func NewEntitlementCreatable() *EntitlementCreatable`

NewEntitlementCreatable instantiates a new EntitlementCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementCreatableWithDefaults

`func NewEntitlementCreatableWithDefaults() *EntitlementCreatable`

NewEntitlementCreatableWithDefaults instantiates a new EntitlementCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementCreatable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementCreatable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementCreatable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementCreatable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementCreatable) GetValues() []EntitlementValueCreatable`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementCreatable) GetValuesOk() (*[]EntitlementValueCreatable, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementCreatable) SetValues(v []EntitlementValueCreatable)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementCreatable) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


