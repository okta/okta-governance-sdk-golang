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

// checks if the RequestTypeResourceSettingsEntitlementBundles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeResourceSettingsEntitlementBundles{}

// RequestTypeResourceSettingsEntitlementBundles Specifies that all resources must be Okta entitlement bundles
type RequestTypeResourceSettingsEntitlementBundles struct {
	Type string `json:"type"`
	// List of Okta entitlement bundles
	TargetResources      []OktaEntitlementBundleResource `json:"targetResources"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeResourceSettingsEntitlementBundles RequestTypeResourceSettingsEntitlementBundles

// NewRequestTypeResourceSettingsEntitlementBundles instantiates a new RequestTypeResourceSettingsEntitlementBundles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeResourceSettingsEntitlementBundles(type_ string, targetResources []OktaEntitlementBundleResource) *RequestTypeResourceSettingsEntitlementBundles {
	this := RequestTypeResourceSettingsEntitlementBundles{}
	this.Type = type_
	this.TargetResources = targetResources
	return &this
}

// NewRequestTypeResourceSettingsEntitlementBundlesWithDefaults instantiates a new RequestTypeResourceSettingsEntitlementBundles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeResourceSettingsEntitlementBundlesWithDefaults() *RequestTypeResourceSettingsEntitlementBundles {
	this := RequestTypeResourceSettingsEntitlementBundles{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeResourceSettingsEntitlementBundles) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsEntitlementBundles) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeResourceSettingsEntitlementBundles) SetType(v string) {
	o.Type = v
}

// GetTargetResources returns the TargetResources field value
func (o *RequestTypeResourceSettingsEntitlementBundles) GetTargetResources() []OktaEntitlementBundleResource {
	if o == nil {
		var ret []OktaEntitlementBundleResource
		return ret
	}

	return o.TargetResources
}

// GetTargetResourcesOk returns a tuple with the TargetResources field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsEntitlementBundles) GetTargetResourcesOk() ([]OktaEntitlementBundleResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetResources, true
}

// SetTargetResources sets field value
func (o *RequestTypeResourceSettingsEntitlementBundles) SetTargetResources(v []OktaEntitlementBundleResource) {
	o.TargetResources = v
}

func (o RequestTypeResourceSettingsEntitlementBundles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeResourceSettingsEntitlementBundles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["targetResources"] = o.TargetResources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeResourceSettingsEntitlementBundles) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"targetResources",
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

	varRequestTypeResourceSettingsEntitlementBundles := _RequestTypeResourceSettingsEntitlementBundles{}

	err = json.Unmarshal(data, &varRequestTypeResourceSettingsEntitlementBundles)

	if err != nil {
		return err
	}

	*o = RequestTypeResourceSettingsEntitlementBundles(varRequestTypeResourceSettingsEntitlementBundles)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "targetResources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeResourceSettingsEntitlementBundles struct {
	value *RequestTypeResourceSettingsEntitlementBundles
	isSet bool
}

func (v NullableRequestTypeResourceSettingsEntitlementBundles) Get() *RequestTypeResourceSettingsEntitlementBundles {
	return v.value
}

func (v *NullableRequestTypeResourceSettingsEntitlementBundles) Set(val *RequestTypeResourceSettingsEntitlementBundles) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeResourceSettingsEntitlementBundles) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeResourceSettingsEntitlementBundles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeResourceSettingsEntitlementBundles(val *RequestTypeResourceSettingsEntitlementBundles) *NullableRequestTypeResourceSettingsEntitlementBundles {
	return &NullableRequestTypeResourceSettingsEntitlementBundles{value: val, isSet: true}
}

func (v NullableRequestTypeResourceSettingsEntitlementBundles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeResourceSettingsEntitlementBundles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
