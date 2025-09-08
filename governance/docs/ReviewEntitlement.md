# ReviewEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The entitlement id | 
**Name** | **string** | The entitlement name | 
**Values** | Pointer to [**[]ReviewerEntitlementValue**](ReviewerEntitlementValue.md) | Collection of entitlement values. | [optional] 

## Methods

### NewReviewEntitlement

`func NewReviewEntitlement(id string, name string, ) *ReviewEntitlement`

NewReviewEntitlement instantiates a new ReviewEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewEntitlementWithDefaults

`func NewReviewEntitlementWithDefaults() *ReviewEntitlement`

NewReviewEntitlementWithDefaults instantiates a new ReviewEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewEntitlement) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ReviewEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewEntitlement) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *ReviewEntitlement) GetValues() []ReviewerEntitlementValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ReviewEntitlement) GetValuesOk() (*[]ReviewerEntitlementValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ReviewEntitlement) SetValues(v []ReviewerEntitlementValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *ReviewEntitlement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


