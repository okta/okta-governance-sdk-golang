# ReviewerEntitlementValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The entitlement value id | 
**Name** | **string** | The entitlement value display name | 
**ExternalValue** | Pointer to **string** | The value of the entitlement property value | [optional] 
**Entitlement** | Pointer to [**ReviewerEntitlement**](ReviewerEntitlement.md) |  | [optional] 

## Methods

### NewReviewerEntitlementValue

`func NewReviewerEntitlementValue(id string, name string, ) *ReviewerEntitlementValue`

NewReviewerEntitlementValue instantiates a new ReviewerEntitlementValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerEntitlementValueWithDefaults

`func NewReviewerEntitlementValueWithDefaults() *ReviewerEntitlementValue`

NewReviewerEntitlementValueWithDefaults instantiates a new ReviewerEntitlementValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewerEntitlementValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewerEntitlementValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewerEntitlementValue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ReviewerEntitlementValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewerEntitlementValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewerEntitlementValue) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *ReviewerEntitlementValue) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *ReviewerEntitlementValue) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *ReviewerEntitlementValue) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *ReviewerEntitlementValue) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetEntitlement

`func (o *ReviewerEntitlementValue) GetEntitlement() ReviewerEntitlement`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *ReviewerEntitlementValue) GetEntitlementOk() (*ReviewerEntitlement, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *ReviewerEntitlementValue) SetEntitlement(v ReviewerEntitlement)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *ReviewerEntitlementValue) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


