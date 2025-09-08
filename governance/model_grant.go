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

// Grant representation of a base grant
type Grant struct {
	GrantMethod *GrantMethod `json:"grantMethod,omitempty"`
	GrantType   GrantType    `json:"grantType"`
	// The date on which the user received an access. Date in ISO 8601 format.
	StartTime *time.Time `json:"startTime,omitempty"`
	// The date on which the user access expires. Date in ISO 8601 format.
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	// The time zone, in IANA format, for the end date of the user access.
	TimeZone             *string               `json:"timeZone,omitempty"`
	Grant                GrantDetails          `json:"grant"`
	Bundle               *Bundle               `json:"bundle,omitempty"`
	Entitlements         []GrantedEntitlements `json:"entitlements"`
	AdditionalProperties map[string]interface{}
}

type _Grant Grant

// NewGrant instantiates a new Grant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrant(grantType GrantType, grant GrantDetails, entitlements []GrantedEntitlements) *Grant {
	this := Grant{}
	this.GrantType = grantType
	this.Grant = grant
	this.Entitlements = entitlements
	return &this
}

// NewGrantWithDefaults instantiates a new Grant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantWithDefaults() *Grant {
	this := Grant{}
	return &this
}

// GetGrantMethod returns the GrantMethod field value if set, zero value otherwise.
func (o *Grant) GetGrantMethod() GrantMethod {
	if o == nil || o.GrantMethod == nil {
		var ret GrantMethod
		return ret
	}
	return *o.GrantMethod
}

// GetGrantMethodOk returns a tuple with the GrantMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetGrantMethodOk() (*GrantMethod, bool) {
	if o == nil || o.GrantMethod == nil {
		return nil, false
	}
	return o.GrantMethod, true
}

// HasGrantMethod returns a boolean if a field has been set.
func (o *Grant) HasGrantMethod() bool {
	if o != nil && o.GrantMethod != nil {
		return true
	}

	return false
}

// SetGrantMethod gets a reference to the given GrantMethod and assigns it to the GrantMethod field.
func (o *Grant) SetGrantMethod(v GrantMethod) {
	o.GrantMethod = &v
}

// GetGrantType returns the GrantType field value
func (o *Grant) GetGrantType() GrantType {
	if o == nil {
		var ret GrantType
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *Grant) GetGrantTypeOk() (*GrantType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *Grant) SetGrantType(v GrantType) {
	o.GrantType = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Grant) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Grant) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Grant) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *Grant) GetExpirationTime() time.Time {
	if o == nil || o.ExpirationTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationTime == nil {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *Grant) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime != nil {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *Grant) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *Grant) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Grant) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *Grant) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetGrant returns the Grant field value
func (o *Grant) GetGrant() GrantDetails {
	if o == nil {
		var ret GrantDetails
		return ret
	}

	return o.Grant
}

// GetGrantOk returns a tuple with the Grant field value
// and a boolean to check if the value has been set.
func (o *Grant) GetGrantOk() (*GrantDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grant, true
}

// SetGrant sets field value
func (o *Grant) SetGrant(v GrantDetails) {
	o.Grant = v
}

// GetBundle returns the Bundle field value if set, zero value otherwise.
func (o *Grant) GetBundle() Bundle {
	if o == nil || o.Bundle == nil {
		var ret Bundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetBundleOk() (*Bundle, bool) {
	if o == nil || o.Bundle == nil {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *Grant) HasBundle() bool {
	if o != nil && o.Bundle != nil {
		return true
	}

	return false
}

// SetBundle gets a reference to the given Bundle and assigns it to the Bundle field.
func (o *Grant) SetBundle(v Bundle) {
	o.Bundle = &v
}

// GetEntitlements returns the Entitlements field value
func (o *Grant) GetEntitlements() []GrantedEntitlements {
	if o == nil {
		var ret []GrantedEntitlements
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *Grant) GetEntitlementsOk() ([]GrantedEntitlements, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *Grant) SetEntitlements(v []GrantedEntitlements) {
	o.Entitlements = v
}

func (o Grant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantMethod != nil {
		toSerialize["grantMethod"] = o.GrantMethod
	}
	if true {
		toSerialize["grantType"] = o.GrantType
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.ExpirationTime != nil {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if true {
		toSerialize["grant"] = o.Grant
	}
	if o.Bundle != nil {
		toSerialize["bundle"] = o.Bundle
	}
	if true {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Grant) UnmarshalJSON(bytes []byte) (err error) {
	varGrant := _Grant{}

	err = json.Unmarshal(bytes, &varGrant)
	if err == nil {
		*o = Grant(varGrant)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "grantMethod")
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "grant")
		delete(additionalProperties, "bundle")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGrant struct {
	value *Grant
	isSet bool
}

func (v NullableGrant) Get() *Grant {
	return v.value
}

func (v *NullableGrant) Set(val *Grant) {
	v.value = val
	v.isSet = true
}

func (v NullableGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrant(val *Grant) *NullableGrant {
	return &NullableGrant{value: val, isSet: true}
}

func (v NullableGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
