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

// checks if the PrincipalAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalAccess{}

// PrincipalAccess Full representation of a principal access
type PrincipalAccess struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ParentResourceOrn string         `json:"parentResourceOrn"`
	Parent            TargetResource `json:"parent"`
	// The Okta user `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetPrincipalOrn string              `json:"targetPrincipalOrn"`
	TargetPrincipal    TargetPrincipalFull `json:"targetPrincipal"`
	// The date on which the user access expires. Date in ISO 8601 format.
	ExpirationTime *time.Time `json:"expirationTime,omitempty"`
	// The time zone, in IANA format, for the end date of the user access.
	TimeZone             *string `json:"timeZone,omitempty"`
	Base                 *Grant  `json:"base,omitempty"`
	Additional           []Grant `json:"additional,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalAccess PrincipalAccess

// NewPrincipalAccess instantiates a new PrincipalAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalAccess(parentResourceOrn string, parent TargetResource, targetPrincipalOrn string, targetPrincipal TargetPrincipalFull) *PrincipalAccess {
	this := PrincipalAccess{}
	this.ParentResourceOrn = parentResourceOrn
	this.Parent = parent
	this.TargetPrincipalOrn = targetPrincipalOrn
	this.TargetPrincipal = targetPrincipal
	return &this
}

// NewPrincipalAccessWithDefaults instantiates a new PrincipalAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalAccessWithDefaults() *PrincipalAccess {
	this := PrincipalAccess{}
	return &this
}

// GetParentResourceOrn returns the ParentResourceOrn field value
func (o *PrincipalAccess) GetParentResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentResourceOrn
}

// GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetParentResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentResourceOrn, true
}

// SetParentResourceOrn sets field value
func (o *PrincipalAccess) SetParentResourceOrn(v string) {
	o.ParentResourceOrn = v
}

// GetParent returns the Parent field value
func (o *PrincipalAccess) GetParent() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetParentOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *PrincipalAccess) SetParent(v TargetResource) {
	o.Parent = v
}

// GetTargetPrincipalOrn returns the TargetPrincipalOrn field value
func (o *PrincipalAccess) GetTargetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetPrincipalOrn
}

// GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetTargetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipalOrn, true
}

// SetTargetPrincipalOrn sets field value
func (o *PrincipalAccess) SetTargetPrincipalOrn(v string) {
	o.TargetPrincipalOrn = v
}

// GetTargetPrincipal returns the TargetPrincipal field value
func (o *PrincipalAccess) GetTargetPrincipal() TargetPrincipalFull {
	if o == nil {
		var ret TargetPrincipalFull
		return ret
	}

	return o.TargetPrincipal
}

// GetTargetPrincipalOk returns a tuple with the TargetPrincipal field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetTargetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetPrincipal, true
}

// SetTargetPrincipal sets field value
func (o *PrincipalAccess) SetTargetPrincipal(v TargetPrincipalFull) {
	o.TargetPrincipal = v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *PrincipalAccess) GetExpirationTime() time.Time {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *PrincipalAccess) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given time.Time and assigns it to the ExpirationTime field.
func (o *PrincipalAccess) SetExpirationTime(v time.Time) {
	o.ExpirationTime = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *PrincipalAccess) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *PrincipalAccess) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *PrincipalAccess) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *PrincipalAccess) GetBase() Grant {
	if o == nil || IsNil(o.Base) {
		var ret Grant
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetBaseOk() (*Grant, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *PrincipalAccess) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given Grant and assigns it to the Base field.
func (o *PrincipalAccess) SetBase(v Grant) {
	o.Base = &v
}

// GetAdditional returns the Additional field value if set, zero value otherwise.
func (o *PrincipalAccess) GetAdditional() []Grant {
	if o == nil || IsNil(o.Additional) {
		var ret []Grant
		return ret
	}
	return o.Additional
}

// GetAdditionalOk returns a tuple with the Additional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccess) GetAdditionalOk() ([]Grant, bool) {
	if o == nil || IsNil(o.Additional) {
		return nil, false
	}
	return o.Additional, true
}

// HasAdditional returns a boolean if a field has been set.
func (o *PrincipalAccess) HasAdditional() bool {
	if o != nil && !IsNil(o.Additional) {
		return true
	}

	return false
}

// SetAdditional gets a reference to the given []Grant and assigns it to the Additional field.
func (o *PrincipalAccess) SetAdditional(v []Grant) {
	o.Additional = v
}

func (o PrincipalAccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parentResourceOrn"] = o.ParentResourceOrn
	toSerialize["parent"] = o.Parent
	toSerialize["targetPrincipalOrn"] = o.TargetPrincipalOrn
	toSerialize["targetPrincipal"] = o.TargetPrincipal
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !IsNil(o.Additional) {
		toSerialize["additional"] = o.Additional
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalAccess) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parentResourceOrn",
		"parent",
		"targetPrincipalOrn",
		"targetPrincipal",
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

	varPrincipalAccess := _PrincipalAccess{}

	err = json.Unmarshal(data, &varPrincipalAccess)

	if err != nil {
		return err
	}

	*o = PrincipalAccess(varPrincipalAccess)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "parentResourceOrn")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "targetPrincipalOrn")
		delete(additionalProperties, "targetPrincipal")
		delete(additionalProperties, "expirationTime")
		delete(additionalProperties, "timeZone")
		delete(additionalProperties, "base")
		delete(additionalProperties, "additional")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalAccess struct {
	value *PrincipalAccess
	isSet bool
}

func (v NullablePrincipalAccess) Get() *PrincipalAccess {
	return v.value
}

func (v *NullablePrincipalAccess) Set(val *PrincipalAccess) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalAccess) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalAccess(val *PrincipalAccess) *NullablePrincipalAccess {
	return &NullablePrincipalAccess{value: val, isSet: true}
}

func (v NullablePrincipalAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
