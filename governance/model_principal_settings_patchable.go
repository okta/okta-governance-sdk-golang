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

// PrincipalSettingsPatchable Governance settings for a principal
type PrincipalSettingsPatchable struct {
	Delegates            *DelegatesPatchable `json:"delegates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalSettingsPatchable PrincipalSettingsPatchable

// NewPrincipalSettingsPatchable instantiates a new PrincipalSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalSettingsPatchable() *PrincipalSettingsPatchable {
	this := PrincipalSettingsPatchable{}
	return &this
}

// NewPrincipalSettingsPatchableWithDefaults instantiates a new PrincipalSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalSettingsPatchableWithDefaults() *PrincipalSettingsPatchable {
	this := PrincipalSettingsPatchable{}
	return &this
}

// GetDelegates returns the Delegates field value if set, zero value otherwise.
func (o *PrincipalSettingsPatchable) GetDelegates() DelegatesPatchable {
	if o == nil || o.Delegates == nil {
		var ret DelegatesPatchable
		return ret
	}
	return *o.Delegates
}

// GetDelegatesOk returns a tuple with the Delegates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalSettingsPatchable) GetDelegatesOk() (*DelegatesPatchable, bool) {
	if o == nil || o.Delegates == nil {
		return nil, false
	}
	return o.Delegates, true
}

// HasDelegates returns a boolean if a field has been set.
func (o *PrincipalSettingsPatchable) HasDelegates() bool {
	if o != nil && o.Delegates != nil {
		return true
	}

	return false
}

// SetDelegates gets a reference to the given DelegatesPatchable and assigns it to the Delegates field.
func (o *PrincipalSettingsPatchable) SetDelegates(v DelegatesPatchable) {
	o.Delegates = &v
}

func (o PrincipalSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delegates != nil {
		toSerialize["delegates"] = o.Delegates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalSettingsPatchable) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalSettingsPatchable := _PrincipalSettingsPatchable{}

	err = json.Unmarshal(bytes, &varPrincipalSettingsPatchable)
	if err == nil {
		*o = PrincipalSettingsPatchable(varPrincipalSettingsPatchable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "delegates")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrincipalSettingsPatchable struct {
	value *PrincipalSettingsPatchable
	isSet bool
}

func (v NullablePrincipalSettingsPatchable) Get() *PrincipalSettingsPatchable {
	return v.value
}

func (v *NullablePrincipalSettingsPatchable) Set(val *PrincipalSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalSettingsPatchable(val *PrincipalSettingsPatchable) *NullablePrincipalSettingsPatchable {
	return &NullablePrincipalSettingsPatchable{value: val, isSet: true}
}

func (v NullablePrincipalSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
