# EntitlementsListObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentResourceOrn** | **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Parent** | [**TargetResource**](TargetResource.md) |  | 
**Links** | Pointer to [**EntitlementLinks**](EntitlementLinks.md) |  | [optional] 
**Metadata** | Pointer to [**ListMetadata**](ListMetadata.md) |  | [optional] 
**Id** | **string** | The &#x60;id&#x60; property of an entitlement | 
**Name** | **string** | The display name for an entitlement property | 
**ExternalValue** | **string** | The value of an entitlement property | 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | **bool** | Indicate if the entitlement property can hold multiple values. If this property is &#x60;true&#x60;, then the &#x60;dataType&#x60; property is set to  &#x60;array&#x60;. | 
**Required** | Pointer to **bool** | The property that determines if the entitlement property is a required attribute | [optional] 
**DataType** | [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | 

## Methods

### NewEntitlementsListObject

`func NewEntitlementsListObject(parentResourceOrn string, parent TargetResource, id string, name string, externalValue string, multiValue bool, dataType EntitlementPropertyDatatype, ) *EntitlementsListObject`

NewEntitlementsListObject instantiates a new EntitlementsListObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsListObjectWithDefaults

`func NewEntitlementsListObjectWithDefaults() *EntitlementsListObject`

NewEntitlementsListObjectWithDefaults instantiates a new EntitlementsListObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentResourceOrn

`func (o *EntitlementsListObject) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *EntitlementsListObject) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *EntitlementsListObject) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.


### GetParent

`func (o *EntitlementsListObject) GetParent() TargetResource`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EntitlementsListObject) GetParentOk() (*TargetResource, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EntitlementsListObject) SetParent(v TargetResource)`

SetParent sets Parent field to given value.


### GetLinks

`func (o *EntitlementsListObject) GetLinks() EntitlementLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementsListObject) GetLinksOk() (*EntitlementLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementsListObject) SetLinks(v EntitlementLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntitlementsListObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetadata

`func (o *EntitlementsListObject) GetMetadata() ListMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EntitlementsListObject) GetMetadataOk() (*ListMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EntitlementsListObject) SetMetadata(v ListMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EntitlementsListObject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetId

`func (o *EntitlementsListObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementsListObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementsListObject) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntitlementsListObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementsListObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementsListObject) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *EntitlementsListObject) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementsListObject) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementsListObject) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetDescription

`func (o *EntitlementsListObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementsListObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementsListObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementsListObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *EntitlementsListObject) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *EntitlementsListObject) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *EntitlementsListObject) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.


### GetRequired

`func (o *EntitlementsListObject) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EntitlementsListObject) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EntitlementsListObject) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EntitlementsListObject) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *EntitlementsListObject) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EntitlementsListObject) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EntitlementsListObject) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


