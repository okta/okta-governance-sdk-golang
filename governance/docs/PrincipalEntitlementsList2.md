# PrincipalEntitlementsList2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** | The Okta user in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format | 
**ResourceOrn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**Data** | [**[]PrincipalEntitlementRow**](PrincipalEntitlementRow.md) | The principal&#39;s effective-entitlement rows. Each row is the resource itself or one of its resource assets. | 
**Links** | [**ListLinks**](ListLinks.md) |  | 

## Methods

### NewPrincipalEntitlementsList2

`func NewPrincipalEntitlementsList2(principalOrn string, resourceOrn string, data []PrincipalEntitlementRow, links ListLinks, ) *PrincipalEntitlementsList2`

NewPrincipalEntitlementsList2 instantiates a new PrincipalEntitlementsList2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementsList2WithDefaults

`func NewPrincipalEntitlementsList2WithDefaults() *PrincipalEntitlementsList2`

NewPrincipalEntitlementsList2WithDefaults instantiates a new PrincipalEntitlementsList2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *PrincipalEntitlementsList2) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *PrincipalEntitlementsList2) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *PrincipalEntitlementsList2) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetResourceOrn

`func (o *PrincipalEntitlementsList2) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *PrincipalEntitlementsList2) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *PrincipalEntitlementsList2) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetData

`func (o *PrincipalEntitlementsList2) GetData() []PrincipalEntitlementRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrincipalEntitlementsList2) GetDataOk() (*[]PrincipalEntitlementRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrincipalEntitlementsList2) SetData(v []PrincipalEntitlementRow)`

SetData sets Data field to given value.


### GetLinks

`func (o *PrincipalEntitlementsList2) GetLinks() ListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PrincipalEntitlementsList2) GetLinksOk() (*ListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PrincipalEntitlementsList2) SetLinks(v ListLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


