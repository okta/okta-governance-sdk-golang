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

// TargetResource Representation of a resource
type TargetResource struct {
	// The Okta `app.id` of the resource  See the [List applications](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Application/#tag/Application/operation/listApplications) endpoint to retrieve application IDs.
	ExternalId           string        `json:"externalId"`
	Type                 ResourceType2 `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _TargetResource TargetResource

// NewTargetResource instantiates a new TargetResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetResource(externalId string, type_ ResourceType2) *TargetResource {
	this := TargetResource{}
	this.ExternalId = externalId
	this.Type = type_
	return &this
}

// NewTargetResourceWithDefaults instantiates a new TargetResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetResourceWithDefaults() *TargetResource {
	this := TargetResource{}
	return &this
}

// GetExternalId returns the ExternalId field value
func (o *TargetResource) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *TargetResource) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *TargetResource) SetExternalId(v string) {
	o.ExternalId = v
}

// GetType returns the Type field value
func (o *TargetResource) GetType() ResourceType2 {
	if o == nil {
		var ret ResourceType2
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TargetResource) GetTypeOk() (*ResourceType2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TargetResource) SetType(v ResourceType2) {
	o.Type = v
}

func (o TargetResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["externalId"] = o.ExternalId
	}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TargetResource) UnmarshalJSON(bytes []byte) (err error) {
	varTargetResource := _TargetResource{}

	err = json.Unmarshal(bytes, &varTargetResource)
	if err == nil {
		*o = TargetResource(varTargetResource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTargetResource struct {
	value *TargetResource
	isSet bool
}

func (v NullableTargetResource) Get() *TargetResource {
	return v.value
}

func (v *NullableTargetResource) Set(val *TargetResource) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetResource) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetResource(val *TargetResource) *NullableTargetResource {
	return &NullableTargetResource{value: val, isSet: true}
}

func (v NullableTargetResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
