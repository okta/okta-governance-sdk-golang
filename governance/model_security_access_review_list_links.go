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
)

// checks if the SecurityAccessReviewListLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewListLinks{}

// SecurityAccessReviewListLinks Links available in security access review, access, or sub access list response
type SecurityAccessReviewListLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewListLinks SecurityAccessReviewListLinks

// NewSecurityAccessReviewListLinks instantiates a new SecurityAccessReviewListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewListLinks(self Link) *SecurityAccessReviewListLinks {
	this := SecurityAccessReviewListLinks{}
	this.Self = self
	return &this
}

// NewSecurityAccessReviewListLinksWithDefaults instantiates a new SecurityAccessReviewListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewListLinksWithDefaults() *SecurityAccessReviewListLinks {
	this := SecurityAccessReviewListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *SecurityAccessReviewListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *SecurityAccessReviewListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SecurityAccessReviewListLinks) GetNext() Link {
	if o == nil || IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewListLinks) GetNextOk() (*Link, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SecurityAccessReviewListLinks) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *SecurityAccessReviewListLinks) SetNext(v Link) {
	o.Next = &v
}

func (o SecurityAccessReviewListLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewListLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewListLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varSecurityAccessReviewListLinks := _SecurityAccessReviewListLinks{}

	err = json.Unmarshal(data, &varSecurityAccessReviewListLinks)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewListLinks(varSecurityAccessReviewListLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewListLinks struct {
	value *SecurityAccessReviewListLinks
	isSet bool
}

func (v NullableSecurityAccessReviewListLinks) Get() *SecurityAccessReviewListLinks {
	return v.value
}

func (v *NullableSecurityAccessReviewListLinks) Set(val *SecurityAccessReviewListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewListLinks(val *SecurityAccessReviewListLinks) *NullableSecurityAccessReviewListLinks {
	return &NullableSecurityAccessReviewListLinks{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
