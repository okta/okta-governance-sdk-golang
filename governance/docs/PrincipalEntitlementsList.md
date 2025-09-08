# PrincipalEntitlementsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PrincipalEntitlement**](PrincipalEntitlement.md) | Principal entitlements list | [optional] 
**Links** | Pointer to [**PrincipalEntitlementsListLinks**](PrincipalEntitlementsListLinks.md) |  | [optional] 

## Methods

### NewPrincipalEntitlementsList

`func NewPrincipalEntitlementsList() *PrincipalEntitlementsList`

NewPrincipalEntitlementsList instantiates a new PrincipalEntitlementsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementsListWithDefaults

`func NewPrincipalEntitlementsListWithDefaults() *PrincipalEntitlementsList`

NewPrincipalEntitlementsListWithDefaults instantiates a new PrincipalEntitlementsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PrincipalEntitlementsList) GetData() []PrincipalEntitlement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrincipalEntitlementsList) GetDataOk() (*[]PrincipalEntitlement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrincipalEntitlementsList) SetData(v []PrincipalEntitlement)`

SetData sets Data field to given value.

### HasData

`func (o *PrincipalEntitlementsList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *PrincipalEntitlementsList) GetLinks() PrincipalEntitlementsListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrincipalEntitlementsList) GetLinksOk() (*PrincipalEntitlementsListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrincipalEntitlementsList) SetLinks(v PrincipalEntitlementsListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PrincipalEntitlementsList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


