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
)

// RequestOnBehalfOfSettingsPatchable Specifies if and for whom a requester may request the resource for.
type RequestOnBehalfOfSettingsPatchable struct {
	// Indicates that users who can request this resource could also request for another requester of the same resource. This property can only be `true`. If request on behalf of should not be allowed then `requestOnBehalfOfSettings` should be nullified.
	Allowed bool `json:"allowed"`
	// Which requesters the resource requester can request on behalf of. If onlyFor is not specified then any requester may request a resource on the behalf of any other user
	OnlyFor              []RequestOnBehalfOfType `json:"onlyFor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestOnBehalfOfSettingsPatchable RequestOnBehalfOfSettingsPatchable

// NewRequestOnBehalfOfSettingsPatchable instantiates a new RequestOnBehalfOfSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestOnBehalfOfSettingsPatchable(allowed bool) *RequestOnBehalfOfSettingsPatchable {
	this := RequestOnBehalfOfSettingsPatchable{}
	this.Allowed = allowed
	return &this
}

// NewRequestOnBehalfOfSettingsPatchableWithDefaults instantiates a new RequestOnBehalfOfSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestOnBehalfOfSettingsPatchableWithDefaults() *RequestOnBehalfOfSettingsPatchable {
	this := RequestOnBehalfOfSettingsPatchable{}
	return &this
}

// GetAllowed returns the Allowed field value
func (o *RequestOnBehalfOfSettingsPatchable) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *RequestOnBehalfOfSettingsPatchable) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *RequestOnBehalfOfSettingsPatchable) SetAllowed(v bool) {
	o.Allowed = v
}

// GetOnlyFor returns the OnlyFor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestOnBehalfOfSettingsPatchable) GetOnlyFor() []RequestOnBehalfOfType {
	if o == nil {
		var ret []RequestOnBehalfOfType
		return ret
	}
	return o.OnlyFor
}

// GetOnlyForOk returns a tuple with the OnlyFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestOnBehalfOfSettingsPatchable) GetOnlyForOk() ([]RequestOnBehalfOfType, bool) {
	if o == nil || o.OnlyFor == nil {
		return nil, false
	}
	return o.OnlyFor, true
}

// HasOnlyFor returns a boolean if a field has been set.
func (o *RequestOnBehalfOfSettingsPatchable) HasOnlyFor() bool {
	if o != nil && o.OnlyFor != nil {
		return true
	}

	return false
}

// SetOnlyFor gets a reference to the given []RequestOnBehalfOfType and assigns it to the OnlyFor field.
func (o *RequestOnBehalfOfSettingsPatchable) SetOnlyFor(v []RequestOnBehalfOfType) {
	o.OnlyFor = v
}

func (o RequestOnBehalfOfSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allowed"] = o.Allowed
	}
	if o.OnlyFor != nil {
		toSerialize["onlyFor"] = o.OnlyFor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestOnBehalfOfSettingsPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varRequestOnBehalfOfSettingsPatchable := _RequestOnBehalfOfSettingsPatchable{}

	err = json.Unmarshal(bytes, &varRequestOnBehalfOfSettingsPatchable)
	if err == nil {
		*o = RequestOnBehalfOfSettingsPatchable(varRequestOnBehalfOfSettingsPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allowed")
		delete(additionalProperties, "onlyFor")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestOnBehalfOfSettingsPatchable struct {
	value *RequestOnBehalfOfSettingsPatchable
	isSet bool
}

func (v NullableRequestOnBehalfOfSettingsPatchable) Get() *RequestOnBehalfOfSettingsPatchable {
	return v.value
}

func (v *NullableRequestOnBehalfOfSettingsPatchable) Set(val *RequestOnBehalfOfSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestOnBehalfOfSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestOnBehalfOfSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestOnBehalfOfSettingsPatchable(val *RequestOnBehalfOfSettingsPatchable) *NullableRequestOnBehalfOfSettingsPatchable {
	return &NullableRequestOnBehalfOfSettingsPatchable{value: val, isSet: true}
}

func (v NullableRequestOnBehalfOfSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestOnBehalfOfSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
