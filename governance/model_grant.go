/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2025 - Present Okta, Inc.

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
	"fmt"
	"time"
)

// checks if the Grant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Grant{}

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
	if o == nil || IsNil(o.GrantMethod) {
		var ret GrantMethod
		return ret
	}
	return *o.GrantMethod
}

// GetGrantMethodOk returns a tuple with the GrantMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetGrantMethodOk() (*GrantMethod, bool) {
	if o == nil || IsNil(o.GrantMethod) {
		return nil, false
	}
	return o.GrantMethod, true
}

// HasGrantMethod returns a boolean if a field has been set.
func (o *Grant) HasGrantMethod() bool {
	if o != nil && !IsNil(o.GrantMethod) {
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
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Grant) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
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
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *Grant) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
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
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Grant) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
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
	if o == nil || IsNil(o.Bundle) {
		var ret Bundle
		return ret
	}
	return *o.Bundle
}

// GetBundleOk returns a tuple with the Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grant) GetBundleOk() (*Bundle, bool) {
	if o == nil || IsNil(o.Bundle) {
		return nil, false
	}
	return o.Bundle, true
}

// HasBundle returns a boolean if a field has been set.
func (o *Grant) HasBundle() bool {
	if o != nil && !IsNil(o.Bundle) {
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Grant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GrantMethod) {
		toSerialize["grantMethod"] = o.GrantMethod
	}
	toSerialize["grantType"] = o.GrantType
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	toSerialize["grant"] = o.Grant
	if !IsNil(o.Bundle) {
		toSerialize["bundle"] = o.Bundle
	}
	toSerialize["entitlements"] = o.Entitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Grant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grantType",
		"grant",
		"entitlements",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGrant := _Grant{}

	err = json.Unmarshal(data, &varGrant)

	if err != nil {
		return err
	}

	*o = Grant(varGrant)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grantMethod")
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "grant")
		delete(additionalProperties, "bundle")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
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
