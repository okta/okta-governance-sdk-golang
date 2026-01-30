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

// checks if the OktaEntitlementBundleResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaEntitlementBundleResource{}

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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaEntitlementBundleResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaEntitlementBundleResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
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

	varOktaEntitlementBundleResource := _OktaEntitlementBundleResource{}

	err = json.Unmarshal(data, &varOktaEntitlementBundleResource)

	if err != nil {
		return err
	}

	*o = OktaEntitlementBundleResource(varOktaEntitlementBundleResource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		o.AdditionalProperties = additionalProperties
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
