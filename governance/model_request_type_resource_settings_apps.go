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

// RequestTypeResourceSettingsApps Specifies that all resources must be Okta apps
type RequestTypeResourceSettingsApps struct {
	Type string `json:"type"`
	// List of Okta applications
	TargetResources      []OktaApplicationResource `json:"targetResources"`
	AdditionalProperties map[string]interface{}
}

type _RequestTypeResourceSettingsApps RequestTypeResourceSettingsApps

// NewRequestTypeResourceSettingsApps instantiates a new RequestTypeResourceSettingsApps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestTypeResourceSettingsApps(type_ string, targetResources []OktaApplicationResource) *RequestTypeResourceSettingsApps {
	this := RequestTypeResourceSettingsApps{}
	this.Type = type_
	this.TargetResources = targetResources
	return &this
}

// NewRequestTypeResourceSettingsAppsWithDefaults instantiates a new RequestTypeResourceSettingsApps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestTypeResourceSettingsAppsWithDefaults() *RequestTypeResourceSettingsApps {
	this := RequestTypeResourceSettingsApps{}
	return &this
}

// GetType returns the Type field value
func (o *RequestTypeResourceSettingsApps) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsApps) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RequestTypeResourceSettingsApps) SetType(v string) {
	o.Type = v
}

// GetTargetResources returns the TargetResources field value
func (o *RequestTypeResourceSettingsApps) GetTargetResources() []OktaApplicationResource {
	if o == nil {
		var ret []OktaApplicationResource
		return ret
	}

	return o.TargetResources
}

// GetTargetResourcesOk returns a tuple with the TargetResources field value
// and a boolean to check if the value has been set.
func (o *RequestTypeResourceSettingsApps) GetTargetResourcesOk() ([]OktaApplicationResource, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetResources, true
}

// SetTargetResources sets field value
func (o *RequestTypeResourceSettingsApps) SetTargetResources(v []OktaApplicationResource) {
	o.TargetResources = v
}

func (o RequestTypeResourceSettingsApps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["targetResources"] = o.TargetResources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestTypeResourceSettingsApps) UnmarshalJSON(bytes []byte) (err error) {
	varRequestTypeResourceSettingsApps := _RequestTypeResourceSettingsApps{}

	err = json.Unmarshal(bytes, &varRequestTypeResourceSettingsApps)
	if err == nil {
		*o = RequestTypeResourceSettingsApps(varRequestTypeResourceSettingsApps)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "targetResources")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestTypeResourceSettingsApps struct {
	value *RequestTypeResourceSettingsApps
	isSet bool
}

func (v NullableRequestTypeResourceSettingsApps) Get() *RequestTypeResourceSettingsApps {
	return v.value
}

func (v *NullableRequestTypeResourceSettingsApps) Set(val *RequestTypeResourceSettingsApps) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestTypeResourceSettingsApps) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestTypeResourceSettingsApps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestTypeResourceSettingsApps(val *RequestTypeResourceSettingsApps) *NullableRequestTypeResourceSettingsApps {
	return &NullableRequestTypeResourceSettingsApps{value: val, isSet: true}
}

func (v NullableRequestTypeResourceSettingsApps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestTypeResourceSettingsApps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
