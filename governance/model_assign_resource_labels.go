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

// AssignResourceLabels The properties to assign resources to labels. All resources in the `resourceOrns` list are assigned to all labels in the `labelValueIds`  list.
type AssignResourceLabels struct {
	// Resources assigned to labels
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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceOrns"] = o.ResourceOrns
	}
	if true {
		toSerialize["labelValueIds"] = o.LabelValueIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssignResourceLabels) UnmarshalJSON(bytes []byte) (err error) {
	varAssignResourceLabels := _AssignResourceLabels{}

	err = json.Unmarshal(bytes, &varAssignResourceLabels)
	if err == nil {
		*o = AssignResourceLabels(varAssignResourceLabels)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceOrns")
		delete(additionalProperties, "labelValueIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
