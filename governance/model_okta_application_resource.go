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

// OktaApplicationResource Identifies a unique Okta resource
type OktaApplicationResource struct {
	// An Okta app ID.  See [list apps](https://developer.okta.com/docs/reference/api/apps/#list-applications) endpoint for reference on how to retrieve app ids.
	ResourceId           string `json:"resourceId" validate:"regexp=0oa[0-9a-zA-Z]+"`
	AdditionalProperties map[string]interface{}
}

type _OktaApplicationResource OktaApplicationResource

// NewOktaApplicationResource instantiates a new OktaApplicationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaApplicationResource(resourceId string) *OktaApplicationResource {
	this := OktaApplicationResource{}
	this.ResourceId = resourceId
	return &this
}

// NewOktaApplicationResourceWithDefaults instantiates a new OktaApplicationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaApplicationResourceWithDefaults() *OktaApplicationResource {
	this := OktaApplicationResource{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *OktaApplicationResource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *OktaApplicationResource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *OktaApplicationResource) SetResourceId(v string) {
	o.ResourceId = v
}

func (o OktaApplicationResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaApplicationResource) UnmarshalJSON(bytes []byte) (err error) {
	varOktaApplicationResource := _OktaApplicationResource{}

	err = json.Unmarshal(bytes, &varOktaApplicationResource)
	if err == nil {
		*o = OktaApplicationResource(varOktaApplicationResource)
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

type NullableOktaApplicationResource struct {
	value *OktaApplicationResource
	isSet bool
}

func (v NullableOktaApplicationResource) Get() *OktaApplicationResource {
	return v.value
}

func (v *NullableOktaApplicationResource) Set(val *OktaApplicationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaApplicationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaApplicationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaApplicationResource(val *OktaApplicationResource) *NullableOktaApplicationResource {
	return &NullableOktaApplicationResource{value: val, isSet: true}
}

func (v NullableOktaApplicationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaApplicationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
