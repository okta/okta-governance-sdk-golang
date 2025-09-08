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

// OktaEntitlementBundleResource Identifies a unique Okta entitlement bundle resource
type OktaEntitlementBundleResource struct {
	// The entitlement bundle `id`
	ResourceId           string `json:"resourceId" validate:"regexp=enb[0-9a-zA-Z]+"`
	AdditionalProperties map[string]interface{}
}

type _OktaEntitlementBundleResource OktaEntitlementBundleResource

// NewOktaEntitlementBundleResource instantiates a new OktaEntitlementBundleResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaEntitlementBundleResource(resourceId string) *OktaEntitlementBundleResource {
	this := OktaEntitlementBundleResource{}
	this.ResourceId = resourceId
	return &this
}

// NewOktaEntitlementBundleResourceWithDefaults instantiates a new OktaEntitlementBundleResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaEntitlementBundleResourceWithDefaults() *OktaEntitlementBundleResource {
	this := OktaEntitlementBundleResource{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *OktaEntitlementBundleResource) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *OktaEntitlementBundleResource) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *OktaEntitlementBundleResource) SetResourceId(v string) {
	o.ResourceId = v
}

func (o OktaEntitlementBundleResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OktaEntitlementBundleResource) UnmarshalJSON(bytes []byte) (err error) {
	varOktaEntitlementBundleResource := _OktaEntitlementBundleResource{}

	err = json.Unmarshal(bytes, &varOktaEntitlementBundleResource)
	if err == nil {
		*o = OktaEntitlementBundleResource(varOktaEntitlementBundleResource)
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

type NullableOktaEntitlementBundleResource struct {
	value *OktaEntitlementBundleResource
	isSet bool
}

func (v NullableOktaEntitlementBundleResource) Get() *OktaEntitlementBundleResource {
	return v.value
}

func (v *NullableOktaEntitlementBundleResource) Set(val *OktaEntitlementBundleResource) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaEntitlementBundleResource) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaEntitlementBundleResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaEntitlementBundleResource(val *OktaEntitlementBundleResource) *NullableOktaEntitlementBundleResource {
	return &NullableOktaEntitlementBundleResource{value: val, isSet: true}
}

func (v NullableOktaEntitlementBundleResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaEntitlementBundleResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
