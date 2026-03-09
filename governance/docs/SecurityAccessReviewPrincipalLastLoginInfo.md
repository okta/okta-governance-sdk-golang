# SecurityAccessReviewPrincipalLastLoginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | Last sign-in date | [optional] 
**Location** | Pointer to [**SecurityAccessReviewPrincipalLocation**](SecurityAccessReviewPrincipalLocation.md) |  | [optional] 
**Device** | Pointer to **string** | Last sign-in device | [optional] 
**IpAddress** | Pointer to **string** | Last sign-in IP address | [optional] 

## Methods

### NewSecurityAccessReviewPrincipalLastLoginInfo

`func NewSecurityAccessReviewPrincipalLastLoginInfo() *SecurityAccessReviewPrincipalLastLoginInfo`

NewSecurityAccessReviewPrincipalLastLoginInfo instantiates a new SecurityAccessReviewPrincipalLastLoginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewPrincipalLastLoginInfoWithDefaults

`func NewSecurityAccessReviewPrincipalLastLoginInfoWithDefaults() *SecurityAccessReviewPrincipalLastLoginInfo`

NewSecurityAccessReviewPrincipalLastLoginInfoWithDefaults instantiates a new SecurityAccessReviewPrincipalLastLoginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLocation

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetLocation() SecurityAccessReviewPrincipalLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetLocationOk() (*SecurityAccessReviewPrincipalLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetLocation(v SecurityAccessReviewPrincipalLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDevice

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetIpAddress

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


