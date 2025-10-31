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

// checks if the TargetResourcesRequestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetResourcesRequestInner{}

// TargetResourcesRequestInner Represents a resource that will be part of access certification
type TargetResourcesRequestInner struct {
	// The resource ID that is being reviewed. The `resourceId` can have a different value based on the `resourceType`. When the `resourceType = GROUP`, the value is a group ID. When the `resourceType = APPLICATION`, the value is an application ID.
	ResourceId   string        `json:"resourceId"`
	ResourceType *ResourceType `json:"resourceType,omitempty"`
	// Include all entitlements and entitlement bundles for this application. Only applicable if the `resourcetype = APPLICATION` and Entitlement Management is enabled.
	IncludeAllEntitlementsAndBundles *bool `json:"includeAllEntitlementsAndBundles,omitempty"`
	// An array of entitlement bundles associated with `resourceId` that should be chosen as target when creating reviews. Only applicable if the `resourceType = APPLICATION` and Entitlement Management is enabled.
	EntitlementBundles []EntitlementBundlesInner `json:"entitlementBundles,omitempty"`
	// An array of entitlements associated with `resourceId` that should be chosen as target when creating reviews. Only applicable if `resourceType = APPLICATION` and Entitlement Management is enabled.
	Entitlements         []EntitlementsInner `json:"entitlements,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TargetResourcesRequestInner TargetResourcesRequestInner

// NewTargetResourcesRequestInner instantiates a new TargetResourcesRequestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetResourcesRequestInner(resourceId string) *TargetResourcesRequestInner {
	this := TargetResourcesRequestInner{}
	this.ResourceId = resourceId
	return &this
}

// NewTargetResourcesRequestInnerWithDefaults instantiates a new TargetResourcesRequestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetResourcesRequestInnerWithDefaults() *TargetResourcesRequestInner {
	this := TargetResourcesRequestInner{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *TargetResourcesRequestInner) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *TargetResourcesRequestInner) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetResourceType() ResourceType {
	if o == nil || IsNil(o.ResourceType) {
		var ret ResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetResourceTypeOk() (*ResourceType, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceType and assigns it to the ResourceType field.
func (o *TargetResourcesRequestInner) SetResourceType(v ResourceType) {
	o.ResourceType = &v
}

// GetIncludeAllEntitlementsAndBundles returns the IncludeAllEntitlementsAndBundles field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetIncludeAllEntitlementsAndBundles() bool {
	if o == nil || IsNil(o.IncludeAllEntitlementsAndBundles) {
		var ret bool
		return ret
	}
	return *o.IncludeAllEntitlementsAndBundles
}

// GetIncludeAllEntitlementsAndBundlesOk returns a tuple with the IncludeAllEntitlementsAndBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetIncludeAllEntitlementsAndBundlesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAllEntitlementsAndBundles) {
		return nil, false
	}
	return o.IncludeAllEntitlementsAndBundles, true
}

// HasIncludeAllEntitlementsAndBundles returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasIncludeAllEntitlementsAndBundles() bool {
	if o != nil && !IsNil(o.IncludeAllEntitlementsAndBundles) {
		return true
	}

	return false
}

// SetIncludeAllEntitlementsAndBundles gets a reference to the given bool and assigns it to the IncludeAllEntitlementsAndBundles field.
func (o *TargetResourcesRequestInner) SetIncludeAllEntitlementsAndBundles(v bool) {
	o.IncludeAllEntitlementsAndBundles = &v
}

// GetEntitlementBundles returns the EntitlementBundles field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetEntitlementBundles() []EntitlementBundlesInner {
	if o == nil || IsNil(o.EntitlementBundles) {
		var ret []EntitlementBundlesInner
		return ret
	}
	return o.EntitlementBundles
}

// GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetEntitlementBundlesOk() ([]EntitlementBundlesInner, bool) {
	if o == nil || IsNil(o.EntitlementBundles) {
		return nil, false
	}
	return o.EntitlementBundles, true
}

// HasEntitlementBundles returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasEntitlementBundles() bool {
	if o != nil && !IsNil(o.EntitlementBundles) {
		return true
	}

	return false
}

// SetEntitlementBundles gets a reference to the given []EntitlementBundlesInner and assigns it to the EntitlementBundles field.
func (o *TargetResourcesRequestInner) SetEntitlementBundles(v []EntitlementBundlesInner) {
	o.EntitlementBundles = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetEntitlements() []EntitlementsInner {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementsInner
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetEntitlementsOk() ([]EntitlementsInner, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementsInner and assigns it to the Entitlements field.
func (o *TargetResourcesRequestInner) SetEntitlements(v []EntitlementsInner) {
	o.Entitlements = v
}

func (o TargetResourcesRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetResourcesRequestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.IncludeAllEntitlementsAndBundles) {
		toSerialize["includeAllEntitlementsAndBundles"] = o.IncludeAllEntitlementsAndBundles
	}
	if !IsNil(o.EntitlementBundles) {
		toSerialize["entitlementBundles"] = o.EntitlementBundles
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TargetResourcesRequestInner) UnmarshalJSON(data []byte) (err error) {
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

	varTargetResourcesRequestInner := _TargetResourcesRequestInner{}

	err = json.Unmarshal(data, &varTargetResourcesRequestInner)

	if err != nil {
		return err
	}

	*o = TargetResourcesRequestInner(varTargetResourcesRequestInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "includeAllEntitlementsAndBundles")
		delete(additionalProperties, "entitlementBundles")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTargetResourcesRequestInner struct {
	value *TargetResourcesRequestInner
	isSet bool
}

func (v NullableTargetResourcesRequestInner) Get() *TargetResourcesRequestInner {
	return v.value
}

func (v *NullableTargetResourcesRequestInner) Set(val *TargetResourcesRequestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetResourcesRequestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetResourcesRequestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetResourcesRequestInner(val *TargetResourcesRequestInner) *NullableTargetResourcesRequestInner {
	return &NullableTargetResourcesRequestInner{value: val, isSet: true}
}

func (v NullableTargetResourcesRequestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetResourcesRequestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
