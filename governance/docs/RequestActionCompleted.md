# RequestActionCompleted

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

### NewRequestActionCompleted

`func NewRequestActionCompleted(status string, action RequestActionEnum, actionId string, actionStatus string, actionAttempted time.Time, resourceName string, resourceId string, resourceType string, ) *RequestActionCompleted`

NewRequestActionCompleted instantiates a new RequestActionCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestActionCompletedWithDefaults

`func NewRequestActionCompletedWithDefaults() *RequestActionCompleted`

NewRequestActionCompletedWithDefaults instantiates a new RequestActionCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RequestActionCompleted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestActionCompleted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestActionCompleted) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAction

`func (o *RequestActionCompleted) GetAction() RequestActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RequestActionCompleted) GetActionOk() (*RequestActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RequestActionCompleted) SetAction(v RequestActionEnum)`

SetAction sets Action field to given value.


### GetActionId

`func (o *RequestActionCompleted) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *RequestActionCompleted) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *RequestActionCompleted) SetActionId(v string)`

SetActionId sets ActionId field to given value.


### GetActionStatus

`func (o *RequestActionCompleted) GetActionStatus() string`

GetActionStatus returns the ActionStatus field if non-nil, zero value otherwise.

### GetActionStatusOk

`func (o *RequestActionCompleted) GetActionStatusOk() (*string, bool)`

GetActionStatusOk returns a tuple with the ActionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionStatus

`func (o *RequestActionCompleted) SetActionStatus(v string)`

SetActionStatus sets ActionStatus field to given value.


### GetActionAttempted

`func (o *RequestActionCompleted) GetActionAttempted() time.Time`

GetActionAttempted returns the ActionAttempted field if non-nil, zero value otherwise.

### GetActionAttemptedOk

`func (o *RequestActionCompleted) GetActionAttemptedOk() (*time.Time, bool)`

GetActionAttemptedOk returns a tuple with the ActionAttempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAttempted

`func (o *RequestActionCompleted) SetActionAttempted(v time.Time)`

SetActionAttempted sets ActionAttempted field to given value.


### GetAccessRemoved

`func (o *RequestActionCompleted) GetAccessRemoved() time.Time`

GetAccessRemoved returns the AccessRemoved field if non-nil, zero value otherwise.

### GetAccessRemovedOk

`func (o *RequestActionCompleted) GetAccessRemovedOk() (*time.Time, bool)`

GetAccessRemovedOk returns a tuple with the AccessRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRemoved

`func (o *RequestActionCompleted) SetAccessRemoved(v time.Time)`

SetAccessRemoved sets AccessRemoved field to given value.

### HasAccessRemoved

`func (o *RequestActionCompleted) HasAccessRemoved() bool`

HasAccessRemoved returns a boolean if a field has been set.

### SetAccessRemovedNil

`func (o *RequestActionCompleted) SetAccessRemovedNil(b bool)`

 SetAccessRemovedNil sets the value for AccessRemoved to be an explicit nil

### UnsetAccessRemoved
`func (o *RequestActionCompleted) UnsetAccessRemoved()`

UnsetAccessRemoved ensures that no value is present for AccessRemoved, not even an explicit nil
### GetResourceName

`func (o *RequestActionCompleted) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *RequestActionCompleted) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *RequestActionCompleted) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceId

`func (o *RequestActionCompleted) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *RequestActionCompleted) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *RequestActionCompleted) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *RequestActionCompleted) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *RequestActionCompleted) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *RequestActionCompleted) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


