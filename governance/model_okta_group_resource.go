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

// OktaGroupResource Identifies a unique Okta group resource
type OktaGroupResource struct {
	// Okta group ID  > **Note:** See [List all groups](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/listGroups) for reference on how to retrieve group IDs.
	ResourceId           string `json:"resourceId" validate:"regexp=00g[0-9a-zA-Z]+"`
	AdditionalProperties map[string]interface{}
}

type _OktaGroupResource OktaGroupResource

// NewOktaGroupResource instantiates a new OktaGroupResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaGroupResource(resourceId string) *OktaGroupResource {
	this := OktaGroupResource{}
	this.ResourceId = resourceId
	return &this
}

// NewOktaGroupResourceWithDefaults instantiates a new OktaGroupResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaGroupResourceWithDefaults() *OktaGroupResource {
	this := OktaGroupResource{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *OktaGroupResource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *OktaGroupResource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *OktaGroupResource) SetResourceId(v string) {
	o.ResourceId = v
}

func (o OktaGroupResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaGroupResource) UnmarshalJSON(bytes []byte) (err error) {
	varOktaGroupResource := _OktaGroupResource{}

	err = json.Unmarshal(bytes, &varOktaGroupResource)
	if err == nil {
		*o = OktaGroupResource(varOktaGroupResource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOktaGroupResource struct {
	value *OktaGroupResource
	isSet bool
}

func (v NullableOktaGroupResource) Get() *OktaGroupResource {
	return v.value
}

func (v *NullableOktaGroupResource) Set(val *OktaGroupResource) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaGroupResource) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaGroupResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaGroupResource(val *OktaGroupResource) *NullableOktaGroupResource {
	return &NullableOktaGroupResource{value: val, isSet: true}
}

func (v NullableOktaGroupResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaGroupResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
