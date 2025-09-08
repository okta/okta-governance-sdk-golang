/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.2.0
Contact: devex-public@okta.com
*/

package governance

import (
	"encoding/json"
	"time"
)

// SecurityAccessReviewPrincipalLastLoginInfo struct for SecurityAccessReviewPrincipalLastLoginInfo
type SecurityAccessReviewPrincipalLastLoginInfo struct {
	// Last login date
	Date     *time.Time                             `json:"date,omitempty"`
	Location *SecurityAccessReviewPrincipalLocation `json:"location,omitempty"`
	// Last login device
	Device *string `json:"device,omitempty"`
	// Last login IP address
	IpAddress            *string `json:"ipAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewPrincipalLastLoginInfo SecurityAccessReviewPrincipalLastLoginInfo

// NewSecurityAccessReviewPrincipalLastLoginInfo instantiates a new SecurityAccessReviewPrincipalLastLoginInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewPrincipalLastLoginInfo() *SecurityAccessReviewPrincipalLastLoginInfo {
	this := SecurityAccessReviewPrincipalLastLoginInfo{}
	return &this
}

// NewSecurityAccessReviewPrincipalLastLoginInfoWithDefaults instantiates a new SecurityAccessReviewPrincipalLastLoginInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewPrincipalLastLoginInfoWithDefaults() *SecurityAccessReviewPrincipalLastLoginInfo {
	this := SecurityAccessReviewPrincipalLastLoginInfo{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetDate(v time.Time) {
	o.Date = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetLocation() SecurityAccessReviewPrincipalLocation {
	if o == nil || o.Location == nil {
		var ret SecurityAccessReviewPrincipalLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetLocationOk() (*SecurityAccessReviewPrincipalLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SecurityAccessReviewPrincipalLocation and assigns it to the Location field.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetLocation(v SecurityAccessReviewPrincipalLocation) {
	o.Location = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetDevice(v string) {
	o.Device = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *SecurityAccessReviewPrincipalLastLoginInfo) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o SecurityAccessReviewPrincipalLastLoginInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewPrincipalLastLoginInfo) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewPrincipalLastLoginInfo := _SecurityAccessReviewPrincipalLastLoginInfo{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewPrincipalLastLoginInfo)
	if err == nil {
		*o = SecurityAccessReviewPrincipalLastLoginInfo(varSecurityAccessReviewPrincipalLastLoginInfo)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "date")
		delete(additionalProperties, "location")
		delete(additionalProperties, "device")
		delete(additionalProperties, "ipAddress")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewPrincipalLastLoginInfo struct {
	value *SecurityAccessReviewPrincipalLastLoginInfo
	isSet bool
}

func (v NullableSecurityAccessReviewPrincipalLastLoginInfo) Get() *SecurityAccessReviewPrincipalLastLoginInfo {
	return v.value
}

func (v *NullableSecurityAccessReviewPrincipalLastLoginInfo) Set(val *SecurityAccessReviewPrincipalLastLoginInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewPrincipalLastLoginInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewPrincipalLastLoginInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewPrincipalLastLoginInfo(val *SecurityAccessReviewPrincipalLastLoginInfo) *NullableSecurityAccessReviewPrincipalLastLoginInfo {
	return &NullableSecurityAccessReviewPrincipalLastLoginInfo{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewPrincipalLastLoginInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewPrincipalLastLoginInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
