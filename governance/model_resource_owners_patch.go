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

// ResourceOwnersPatch struct for ResourceOwnersPatch
type ResourceOwnersPatch struct {
	// The `id` of the resource in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an app, or a bundle. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn          string                         `json:"resourceOrn"`
	Data                 []ResourceOwnersPatchDataInner `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _ResourceOwnersPatch ResourceOwnersPatch

// NewResourceOwnersPatch instantiates a new ResourceOwnersPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceOwnersPatch(resourceOrn string, data []ResourceOwnersPatchDataInner) *ResourceOwnersPatch {
	this := ResourceOwnersPatch{}
	this.ResourceOrn = resourceOrn
	this.Data = data
	return &this
}

// NewResourceOwnersPatchWithDefaults instantiates a new ResourceOwnersPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceOwnersPatchWithDefaults() *ResourceOwnersPatch {
	this := ResourceOwnersPatch{}
	return &this
}

// GetResourceOrn returns the ResourceOrn field value
func (o *ResourceOwnersPatch) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnersPatch) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *ResourceOwnersPatch) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetData returns the Data field value
func (o *ResourceOwnersPatch) GetData() []ResourceOwnersPatchDataInner {
	if o == nil {
		var ret []ResourceOwnersPatchDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ResourceOwnersPatch) GetDataOk() ([]ResourceOwnersPatchDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ResourceOwnersPatch) SetData(v []ResourceOwnersPatchDataInner) {
	o.Data = v
}

func (o ResourceOwnersPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceOwnersPatch) UnmarshalJSON(bytes []byte) (err error) {
	varResourceOwnersPatch := _ResourceOwnersPatch{}

	err = json.Unmarshal(bytes, &varResourceOwnersPatch)
	if err == nil {
		*o = ResourceOwnersPatch(varResourceOwnersPatch)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceOwnersPatch struct {
	value *ResourceOwnersPatch
	isSet bool
}

func (v NullableResourceOwnersPatch) Get() *ResourceOwnersPatch {
	return v.value
}

func (v *NullableResourceOwnersPatch) Set(val *ResourceOwnersPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceOwnersPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceOwnersPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceOwnersPatch(val *ResourceOwnersPatch) *NullableResourceOwnersPatch {
	return &NullableResourceOwnersPatch{value: val, isSet: true}
}

func (v NullableResourceOwnersPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceOwnersPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
