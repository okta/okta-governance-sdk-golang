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
	if o == nil || o.ResourceType == nil {
		var ret ResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetResourceTypeOk() (*ResourceType, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
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
	if o == nil || o.IncludeAllEntitlementsAndBundles == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAllEntitlementsAndBundles
}

// GetIncludeAllEntitlementsAndBundlesOk returns a tuple with the IncludeAllEntitlementsAndBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetIncludeAllEntitlementsAndBundlesOk() (*bool, bool) {
	if o == nil || o.IncludeAllEntitlementsAndBundles == nil {
		return nil, false
	}
	return o.IncludeAllEntitlementsAndBundles, true
}

// HasIncludeAllEntitlementsAndBundles returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasIncludeAllEntitlementsAndBundles() bool {
	if o != nil && o.IncludeAllEntitlementsAndBundles != nil {
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
	if o == nil || o.EntitlementBundles == nil {
		var ret []EntitlementBundlesInner
		return ret
	}
	return o.EntitlementBundles
}

// GetEntitlementBundlesOk returns a tuple with the EntitlementBundles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetEntitlementBundlesOk() ([]EntitlementBundlesInner, bool) {
	if o == nil || o.EntitlementBundles == nil {
		return nil, false
	}
	return o.EntitlementBundles, true
}

// HasEntitlementBundles returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasEntitlementBundles() bool {
	if o != nil && o.EntitlementBundles != nil {
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
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementsInner
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetEntitlementsOk() ([]EntitlementsInner, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementsInner and assigns it to the Entitlements field.
func (o *TargetResourcesRequestInner) SetEntitlements(v []EntitlementsInner) {
	o.Entitlements = v
}

func (o TargetResourcesRequestInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.IncludeAllEntitlementsAndBundles != nil {
		toSerialize["includeAllEntitlementsAndBundles"] = o.IncludeAllEntitlementsAndBundles
	}
	if o.EntitlementBundles != nil {
		toSerialize["entitlementBundles"] = o.EntitlementBundles
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TargetResourcesRequestInner) UnmarshalJSON(bytes []byte) (err error) {
	varTargetResourcesRequestInner := _TargetResourcesRequestInner{}

	err = json.Unmarshal(bytes, &varTargetResourcesRequestInner)
	if err == nil {
		*o = TargetResourcesRequestInner(varTargetResourcesRequestInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "includeAllEntitlementsAndBundles")
		delete(additionalProperties, "entitlementBundles")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
