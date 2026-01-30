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

// checks if the AssignResourceLabels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignResourceLabels{}

// AssignResourceLabels The properties to assign resources to labels. All resources in the `resourceOrns` list are assigned to all labels in the `labelValueIds`  list.
type AssignResourceLabels struct {
	// Resources assigned to labels (in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format)
	ResourceOrns []string `json:"resourceOrns"`
	// Labels assigned to resources in the `resourceOrns` list
	LabelValueIds        []string `json:"labelValueIds"`
	AdditionalProperties map[string]interface{}
}

type _AssignResourceLabels AssignResourceLabels

// NewAssignResourceLabels instantiates a new AssignResourceLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignResourceLabels(resourceOrns []string, labelValueIds []string) *AssignResourceLabels {
	this := AssignResourceLabels{}
	this.ResourceOrns = resourceOrns
	this.LabelValueIds = labelValueIds
	return &this
}

// NewAssignResourceLabelsWithDefaults instantiates a new AssignResourceLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignResourceLabelsWithDefaults() *AssignResourceLabels {
	this := AssignResourceLabels{}
	return &this
}

// GetResourceOrns returns the ResourceOrns field value
func (o *AssignResourceLabels) GetResourceOrns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceOrns
}

// GetResourceOrnsOk returns a tuple with the ResourceOrns field value
// and a boolean to check if the value has been set.
func (o *AssignResourceLabels) GetResourceOrnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceOrns, true
}

// SetResourceOrns sets field value
func (o *AssignResourceLabels) SetResourceOrns(v []string) {
	o.ResourceOrns = v
}

// GetLabelValueIds returns the LabelValueIds field value
func (o *AssignResourceLabels) GetLabelValueIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LabelValueIds
}

// GetLabelValueIdsOk returns a tuple with the LabelValueIds field value
// and a boolean to check if the value has been set.
func (o *AssignResourceLabels) GetLabelValueIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LabelValueIds, true
}

// SetLabelValueIds sets field value
func (o *AssignResourceLabels) SetLabelValueIds(v []string) {
	o.LabelValueIds = v
}

func (o AssignResourceLabels) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignResourceLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceOrns"] = o.ResourceOrns
	toSerialize["labelValueIds"] = o.LabelValueIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignResourceLabels) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceOrns",
		"labelValueIds",
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

	varAssignResourceLabels := _AssignResourceLabels{}

	err = json.Unmarshal(data, &varAssignResourceLabels)

	if err != nil {
		return err
	}

	*o = AssignResourceLabels(varAssignResourceLabels)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceOrns")
		delete(additionalProperties, "labelValueIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignResourceLabels struct {
	value *AssignResourceLabels
	isSet bool
}

func (v NullableAssignResourceLabels) Get() *AssignResourceLabels {
	return v.value
}

func (v *NullableAssignResourceLabels) Set(val *AssignResourceLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignResourceLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignResourceLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignResourceLabels(val *AssignResourceLabels) *NullableAssignResourceLabels {
	return &NullableAssignResourceLabels{value: val, isSet: true}
}

func (v NullableAssignResourceLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignResourceLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
