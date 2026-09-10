# EntitlementDriftEntitlementReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementId** | **string** | The &#x60;id&#x60; of the entitlement the drift applies to | 
**Id** | Pointer to **string** | The &#x60;id&#x60; of the entitlement value, after Okta has registered it | [optional] 
**Value** | Pointer to **string** | The entitlement value as seen in the import, or as Okta governed it before this import | [optional] 
**ExternalId** | Pointer to **string** | The entitlement value&#39;s ID in the source app | [optional] 

## Methods

### NewEntitlementDriftEntitlementReference

`func NewEntitlementDriftEntitlementReference(entitlementId string, ) *EntitlementDriftEntitlementReference`

NewEntitlementDriftEntitlementReference instantiates a new EntitlementDriftEntitlementReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDriftEntitlementReferenceWithDefaults

`func NewEntitlementDriftEntitlementReferenceWithDefaults() *EntitlementDriftEntitlementReference`

NewEntitlementDriftEntitlementReferenceWithDefaults instantiates a new EntitlementDriftEntitlementReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementId

`func (o *EntitlementDriftEntitlementReference) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementDriftEntitlementReference) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementDriftEntitlementReference) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetId

`func (o *EntitlementDriftEntitlementReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementDriftEntitlementReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementDriftEntitlementReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementDriftEntitlementReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *EntitlementDriftEntitlementReference) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntitlementDriftEntitlementReference) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntitlementDriftEntitlementReference) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntitlementDriftEntitlementReference) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExternalId

`func (o *EntitlementDriftEntitlementReference) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntitlementDriftEntitlementReference) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntitlementDriftEntitlementReference) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntitlementDriftEntitlementReference) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


