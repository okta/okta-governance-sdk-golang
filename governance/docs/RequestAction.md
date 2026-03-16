# RequestAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of a completed approval | 
**Action** | [**RequestActionEnum**](RequestActionEnum.md) |  | 
**ActionId** | **string** | A unique identifier of the action taken in the request | 
**ActionStatus** | **string** | Whether the action was successful or not | 
**ActionAttempted** | **time.Time** | When the action was attempted | 
**AccessRemoved** | Pointer to **NullableTime** | When access was removed if the associated Request Type has an &#x60;accessDuration&#x60;. null if the Request Type does not have an &#x60;accessDuration&#x60;, or the duration has not yet expired. | [optional] 
**ResourceName** | **string** | Human readable name of the resource | 
**ResourceId** | **string** | The Okta &#x60;app.id&#x60;, or &#x60;group.id&#x60; of the resource that can be requested with this Request Type.  * See [List applications](https://developer.okta.com/docs/reference/api/apps/#list-applications) to retrieve app IDs. * See [List groups](https://developer.okta.com/docs/reference/api/groups/#list-groups) to retrieve group IDs.  | 
**ResourceType** | **string** | The type of resource | 

## Methods

### NewRequestAction

`func NewRequestAction(status string, action RequestActionEnum, actionId string, actionStatus string, actionAttempted time.Time, resourceName string, resourceId string, resourceType string, ) *RequestAction`

NewRequestAction instantiates a new RequestAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestActionWithDefaults

`func NewRequestActionWithDefaults() *RequestAction`

NewRequestActionWithDefaults instantiates a new RequestAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestAction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestAction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestAction) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAction

`func (o *RequestAction) GetAction() RequestActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestAction) GetActionOk() (*RequestActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestAction) SetAction(v RequestActionEnum)`

SetAction sets Action field to given value.


### GetActionId

`func (o *RequestAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *RequestAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *RequestAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetActionStatus

`func (o *RequestAction) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *RequestAction) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *RequestAction) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.


### GetActionAttempted

`func (o *RequestAction) GetActionAttempted() time.Time`

GetActionAttempted returns the ActionAttempted field if non-nil, zero value otherwise.

### GetActionAttemptedOk

`func (o *RequestAction) GetActionAttemptedOk() (*time.Time, bool)`

GetActionAttemptedOk returns a tuple with the ActionAttempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAttempted

`func (o *RequestAction) SetActionAttempted(v time.Time)`

SetActionAttempted sets ActionAttempted field to given value.


### GetAccessRemoved

`func (o *RequestAction) GetAccessRemoved() time.Time`

GetAccessRemoved returns the AccessRemoved field if non-nil, zero value otherwise.

### GetAccessRemovedOk

`func (o *RequestAction) GetAccessRemovedOk() (*time.Time, bool)`

GetAccessRemovedOk returns a tuple with the AccessRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoved

`func (o *RequestAction) SetAccessRemoved(v time.Time)`

SetAccessRemoved sets AccessRemoved field to given value.

### HasAccessRemoved

`func (o *RequestAction) HasAccessRemoved() bool`

HasAccessRemoved returns a boolean if a field has been set.

### SetAccessRemovedNil

`func (o *RequestAction) SetAccessRemovedNil(b bool)`

 SetAccessRemovedNil sets the value for AccessRemoved to be an explicit nil

### UnsetAccessRemoved
`func (o *RequestAction) UnsetAccessRemoved()`

UnsetAccessRemoved ensures that no value is present for AccessRemoved, not even an explicit nil
### GetResourceName

`func (o *RequestAction) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *RequestAction) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *RequestAction) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceId

`func (o *RequestAction) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *RequestAction) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *RequestAction) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *RequestAction) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *RequestAction) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *RequestAction) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


