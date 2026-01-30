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

// checks if the RequestTypeResourceSettingsGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestTypeResourceSettingsGroups{}

// RequestTypeResourceSettingsGroups Specifies that all resources must be Okta groups
type RequestTypeResourceSettingsGroups struct {
	Type string `json:"type"`
	// List of Okta groups
	TargetResources      []OktaGroupResource `json:"targetResources"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeResourceSettingsGroups RequestTypeResourceSettingsGroups

// NewRequestTypeResourceSettingsGroups instantiates a new RequestTypeResourceSettingsGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeResourceSettingsGroups(type_ string, targetResources []OktaGroupResource) *RequestTypeResourceSettingsGroups {
	this := RequestTypeResourceSettingsGroups{}
	this.Type = type_
	this.TargetResources = targetResources
	return &this
}

// NewRequestTypeResourceSettingsGroupsWithDefaults instantiates a new RequestTypeResourceSettingsGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeResourceSettingsGroupsWithDefaults() *RequestTypeResourceSettingsGroups {
	this := RequestTypeResourceSettingsGroups{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeResourceSettingsGroups) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsGroups) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeResourceSettingsGroups) SetType(v string) {
	o.Type = v
}

// GetTargetResources returns the TargetResources field value
func (o *RequestTypeResourceSettingsGroups) GetTargetResources() []OktaGroupResource {
	if o == nil {
		var ret []OktaGroupResource
		return ret
	}

	return o.TargetResources
}

// GetTargetResourcesOk returns a tuple with the TargetResources field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsGroups) GetTargetResourcesOk() ([]OktaGroupResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetResources, true
}

// SetTargetResources sets field value
func (o *RequestTypeResourceSettingsGroups) SetTargetResources(v []OktaGroupResource) {
	o.TargetResources = v
}

func (o RequestTypeResourceSettingsGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestTypeResourceSettingsGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["targetResources"] = o.TargetResources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestTypeResourceSettingsGroups) UnmarshalJSON(data []byte) (err error) {
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

	varRequestTypeResourceSettingsGroups := _RequestTypeResourceSettingsGroups{}

	err = json.Unmarshal(data, &varRequestTypeResourceSettingsGroups)

	if err != nil {
		return err
	}

	*o = RequestTypeResourceSettingsGroups(varRequestTypeResourceSettingsGroups)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "targetResources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestTypeResourceSettingsGroups struct {
	value *RequestTypeResourceSettingsGroups
	isSet bool
}

func (v NullableRequestTypeResourceSettingsGroups) Get() *RequestTypeResourceSettingsGroups {
	return v.value
}

func (v *NullableRequestTypeResourceSettingsGroups) Set(val *RequestTypeResourceSettingsGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeResourceSettingsGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeResourceSettingsGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeResourceSettingsGroups(val *RequestTypeResourceSettingsGroups) *NullableRequestTypeResourceSettingsGroups {
	return &NullableRequestTypeResourceSettingsGroups{value: val, isSet: true}
}

func (v NullableRequestTypeResourceSettingsGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeResourceSettingsGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
