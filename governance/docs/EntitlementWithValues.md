# EntitlementWithValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The &#x60;id&#x60; property of an entitlement | 
**Name** | **string** | The display name for an entitlement property | 
**ExternalValue** | Pointer to **string** | The value of an entitlement property | [optional] 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | Pointer to **bool** | Indicates if the entitlement property can hold multiple values. If this property is &#x60;true&#x60;, then the &#x60;dataType&#x60; property is set to &#x60;array&#x60;. | [optional] 
**Required** | Pointer to **bool** | The property that determines if the entitlement property is a required attribute | [optional] 
**DataType** | Pointer to [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | [optional] 
**Values** | [**[]PrincipalEntitlementValue**](PrincipalEntitlementValue.md) | The first page of the principal&#39;s effective values for this entitlement property. Up to 20 values are returned inline, ordered by name in ascending order. If more values remain, the &#x60;_links.next&#x60; reference points to &#x60;GET /governance/api/v2/principal-entitlements/values&#x60; with the query parameters for the next page. When the &#x60;_links.next&#x60; reference is absent, &#x60;values&#x60; is the complete set.  | 
**Links** | Pointer to [**EntitlementWithValuesAllOfLinks**](EntitlementWithValuesAllOfLinks.md) |  | [optional] 

## Methods

### NewEntitlementWithValues

`func NewEntitlementWithValues(id string, name string, values []PrincipalEntitlementValue, ) *EntitlementWithValues`

NewEntitlementWithValues instantiates a new EntitlementWithValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementWithValuesWithDefaults

`func NewEntitlementWithValuesWithDefaults() *EntitlementWithValues`

NewEntitlementWithValuesWithDefaults instantiates a new EntitlementWithValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementWithValues) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementWithValues) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementWithValues) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntitlementWithValues) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementWithValues) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementWithValues) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *EntitlementWithValues) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementWithValues) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementWithValues) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *EntitlementWithValues) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementWithValues) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementWithValues) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementWithValues) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementWithValues) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *EntitlementWithValues) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *EntitlementWithValues) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *EntitlementWithValues) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *EntitlementWithValues) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.

### GetRequired

`func (o *EntitlementWithValues) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EntitlementWithValues) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EntitlementWithValues) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EntitlementWithValues) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *EntitlementWithValues) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EntitlementWithValues) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EntitlementWithValues) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *EntitlementWithValues) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementWithValues) GetValues() []PrincipalEntitlementValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementWithValues) GetValuesOk() (*[]PrincipalEntitlementValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementWithValues) SetValues(v []PrincipalEntitlementValue)`

SetValues sets Values field to given value.


### GetLinks

`func (o *EntitlementWithValues) GetLinks() EntitlementWithValuesAllOfLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementWithValues) GetLinksOk() (*EntitlementWithValuesAllOfLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementWithValues) SetLinks(v EntitlementWithValuesAllOfLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntitlementWithValues) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


