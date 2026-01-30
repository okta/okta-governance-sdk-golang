# EntitlementAccessDetailsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentObjectType** | [**AssignmentObjectType**](AssignmentObjectType.md) |  | 
**AssignmentObject** | Pointer to [**AssignmentObject**](AssignmentObject.md) |  | [optional] 
**AssignmentMethod** | [**AssignmentMethod**](AssignmentMethod.md) |  | 
**Entitlements** | [**[]EntitlementFull**](EntitlementFull.md) | Collection of entitlements with associated values | 
**AccessDuration** | Pointer to [**AssignmentProperties**](AssignmentProperties.md) |  | [optional] 
**Grant** | Pointer to [**GrantObject**](GrantObject.md) |  | [optional] 

## Methods

### NewEntitlementAccessDetailsObject

`func NewEntitlementAccessDetailsObject(assignmentObjectType AssignmentObjectType, assignmentMethod AssignmentMethod, entitlements []EntitlementFull, ) *EntitlementAccessDetailsObject`

NewEntitlementAccessDetailsObject instantiates a new EntitlementAccessDetailsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementAccessDetailsObjectWithDefaults

`func NewEntitlementAccessDetailsObjectWithDefaults() *EntitlementAccessDetailsObject`

NewEntitlementAccessDetailsObjectWithDefaults instantiates a new EntitlementAccessDetailsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentObjectType

`func (o *EntitlementAccessDetailsObject) GetAssignmentObjectType() AssignmentObjectType`

GetAssignmentObjectType returns the AssignmentObjectType field if non-nil, zero value otherwise.

### GetAssignmentObjectTypeOk

`func (o *EntitlementAccessDetailsObject) GetAssignmentObjectTypeOk() (*AssignmentObjectType, bool)`

GetAssignmentObjectTypeOk returns a tuple with the AssignmentObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentObjectType

`func (o *EntitlementAccessDetailsObject) SetAssignmentObjectType(v AssignmentObjectType)`

SetAssignmentObjectType sets AssignmentObjectType field to given value.


### GetAssignmentObject

`func (o *EntitlementAccessDetailsObject) GetAssignmentObject() AssignmentObject`

GetAssignmentObject returns the AssignmentObject field if non-nil, zero value otherwise.

### GetAssignmentObjectOk

`func (o *EntitlementAccessDetailsObject) GetAssignmentObjectOk() (*AssignmentObject, bool)`

GetAssignmentObjectOk returns a tuple with the AssignmentObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentObject

`func (o *EntitlementAccessDetailsObject) SetAssignmentObject(v AssignmentObject)`

SetAssignmentObject sets AssignmentObject field to given value.

### HasAssignmentObject

`func (o *EntitlementAccessDetailsObject) HasAssignmentObject() bool`

HasAssignmentObject returns a boolean if a field has been set.

### GetAssignmentMethod

`func (o *EntitlementAccessDetailsObject) GetAssignmentMethod() AssignmentMethod`

GetAssignmentMethod returns the AssignmentMethod field if non-nil, zero value otherwise.

### GetAssignmentMethodOk

`func (o *EntitlementAccessDetailsObject) GetAssignmentMethodOk() (*AssignmentMethod, bool)`

GetAssignmentMethodOk returns a tuple with the AssignmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentMethod

`func (o *EntitlementAccessDetailsObject) SetAssignmentMethod(v AssignmentMethod)`

SetAssignmentMethod sets AssignmentMethod field to given value.


### GetEntitlements

`func (o *EntitlementAccessDetailsObject) GetEntitlements() []EntitlementFull`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementAccessDetailsObject) GetEntitlementsOk() (*[]EntitlementFull, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementAccessDetailsObject) SetEntitlements(v []EntitlementFull)`

SetEntitlements sets Entitlements field to given value.


### GetAccessDuration

`func (o *EntitlementAccessDetailsObject) GetAccessDuration() AssignmentProperties`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *EntitlementAccessDetailsObject) GetAccessDurationOk() (*AssignmentProperties, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *EntitlementAccessDetailsObject) SetAccessDuration(v AssignmentProperties)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *EntitlementAccessDetailsObject) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.

### GetGrant

`func (o *EntitlementAccessDetailsObject) GetGrant() GrantObject`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *EntitlementAccessDetailsObject) GetGrantOk() (*GrantObject, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *EntitlementAccessDetailsObject) SetGrant(v GrantObject)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *EntitlementAccessDetailsObject) HasGrant() bool`

HasGrant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


