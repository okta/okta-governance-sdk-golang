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

// ResourceGrantLinks Links available on a single grant representation
type ResourceGrantLinks struct {
	Self                 Link                   `json:"self"`
	EntitlementBundle    *EntitlementBundleLink `json:"entitlementBundle,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceGrantLinks ResourceGrantLinks

// NewResourceGrantLinks instantiates a new ResourceGrantLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceGrantLinks(self Link) *ResourceGrantLinks {
	this := ResourceGrantLinks{}
	this.Self = self
	return &this
}

// NewResourceGrantLinksWithDefaults instantiates a new ResourceGrantLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceGrantLinksWithDefaults() *ResourceGrantLinks {
	this := ResourceGrantLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *ResourceGrantLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ResourceGrantLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ResourceGrantLinks) SetSelf(v Link) {
	o.Self = v
}

// GetEntitlementBundle returns the EntitlementBundle field value if set, zero value otherwise.
func (o *ResourceGrantLinks) GetEntitlementBundle() EntitlementBundleLink {
	if o == nil || o.EntitlementBundle == nil {
		var ret EntitlementBundleLink
		return ret
	}
	return *o.EntitlementBundle
}

// GetEntitlementBundleOk returns a tuple with the EntitlementBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceGrantLinks) GetEntitlementBundleOk() (*EntitlementBundleLink, bool) {
	if o == nil || o.EntitlementBundle == nil {
		return nil, false
	}
	return o.EntitlementBundle, true
}

// HasEntitlementBundle returns a boolean if a field has been set.
func (o *ResourceGrantLinks) HasEntitlementBundle() bool {
	if o != nil && o.EntitlementBundle != nil {
		return true
	}

	return false
}

// SetEntitlementBundle gets a reference to the given EntitlementBundleLink and assigns it to the EntitlementBundle field.
func (o *ResourceGrantLinks) SetEntitlementBundle(v EntitlementBundleLink) {
	o.EntitlementBundle = &v
}

func (o ResourceGrantLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.EntitlementBundle != nil {
		toSerialize["entitlementBundle"] = o.EntitlementBundle
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceGrantLinks) UnmarshalJSON(bytes []byte) (err error) {
	varResourceGrantLinks := _ResourceGrantLinks{}

	err = json.Unmarshal(bytes, &varResourceGrantLinks)
	if err == nil {
		*o = ResourceGrantLinks(varResourceGrantLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "entitlementBundle")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceGrantLinks struct {
	value *ResourceGrantLinks
	isSet bool
}

func (v NullableResourceGrantLinks) Get() *ResourceGrantLinks {
	return v.value
}

func (v *NullableResourceGrantLinks) Set(val *ResourceGrantLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceGrantLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceGrantLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceGrantLinks(val *ResourceGrantLinks) *NullableResourceGrantLinks {
	return &NullableResourceGrantLinks{value: val, isSet: true}
}

func (v NullableResourceGrantLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceGrantLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
