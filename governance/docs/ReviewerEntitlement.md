# ReviewerEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The entitlement &#x60;id&#x60; | 
**Name** | **string** | The entitlement display name | 
**ExternalValue** | Pointer to **string** | The value of the entitlement property | [optional] 

## Methods

### NewReviewerEntitlement

`func NewReviewerEntitlement(id string, name string, ) *ReviewerEntitlement`

NewReviewerEntitlement instantiates a new ReviewerEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerEntitlementWithDefaults

`func NewReviewerEntitlementWithDefaults() *ReviewerEntitlement`

NewReviewerEntitlementWithDefaults instantiates a new ReviewerEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewerEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewerEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewerEntitlement) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ReviewerEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewerEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewerEntitlement) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *ReviewerEntitlement) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *ReviewerEntitlement) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *ReviewerEntitlement) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *ReviewerEntitlement) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


