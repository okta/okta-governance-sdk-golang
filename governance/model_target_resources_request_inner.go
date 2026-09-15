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

// TargetResourcesRequestInner A resource that's included in the access certification campaign
type TargetResourcesRequestInner struct {
	// The resource ID to review. The `resourceId` value depends on the `resourceType` object.
	ResourceId   string        `json:"resourceId"`
	ResourceType *ResourceType `json:"resourceType,omitempty"`
	// Only applicable if the `resourceType` is `APPLICATION` and entitlement management is enabled.  If `true`, include all entitlements and entitlement bundles for the app.
	IncludeAllEntitlementsAndBundles *bool `json:"includeAllEntitlementsAndBundles,omitempty"`
	// Only applicable if the `resourceType` is `APPLICATION` and entitlement management is enabled.  A list of entitlement bundles (associated with the app specified in `resourceId`) selected as targets for review.
	EntitlementBundles []EntitlementBundlesInner `json:"entitlementBundles,omitempty"`
	// Only applicable if `resourceType` is `APPLICATION` and entitlement management is enabled. A list of entitlements (associated with the app specified in `resourceId`) selected as targets for review.
	Entitlements []EntitlementsInner `json:"entitlements,omitempty"`
	// <x-lifecycle class=\"beta\"></x-lifecycle><br> Only applicable when `targetResources.resourceType` is `APPLICATION` or `GOVERNANCE_LABEL_VALUE`.  A list of governance label values assigned to resources to target for the review.
	EntitlementValueScopeByGovernanceLabelValues []GovernanceLabelValuesInner `json:"entitlementValueScopeByGovernanceLabelValues,omitempty"`
	// <x-lifecycle class=\"beta\"></x-lifecycle><br> If `true`, entitlements bundles are also included as additional targets to review. Only entitlement bundles with values that are targeted through the `entitlementValueScopeByGovernanceLabelValues` list are included.
	IncludeBundlesHavingSameEntitlementValues *bool `json:"includeBundlesHavingSameEntitlementValues,omitempty"`
	// <x-lifecycle class=\"beta\"></x-lifecycle><br> Only applicable when `targetResources.resourceType` is `APPLICATION` or `GOVERNANCE_LABEL_VALUE`.  A list of governance label values assigned to resources to target for the review.
	EntitlementBundleScopeByGovernanceLabelValues []GovernanceLabelValuesInner `json:"entitlementBundleScopeByGovernanceLabelValues,omitempty"`
	// Only applicable if the `resourceType` is `APPLICATION`.  A list of app service accounts (associated with the app specified in `resourceId`) selected as targets for review.
	AppServiceAccounts []AppServiceAccountsInner `json:"appServiceAccounts,omitempty"`
	// Only applicable if `resourceSettings.type` is `APP_SERVICE_ACCOUNT` and `resourceSettings.targetResources.resourceType` is `APPLICATION`:  * If true, includes all SaaS app service accounts associated with the app specified in `resourceSettings.targetResources.resourceId`. * If false, includes only the SaaS app service accounts specified in the list of `resourceSettings.targetResources.appServiceAccounts`. The list of service account IDs in `resourceSettings.targetResources.appServiceAccounts` must be associated with the app in `resourceSettings.targetResources.resourceId`.
	IncludeAllAppServiceAccounts *bool `json:"includeAllAppServiceAccounts,omitempty"`
	AdditionalProperties         map[string]interface{}
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

// GetEntitlementValueScopeByGovernanceLabelValues returns the EntitlementValueScopeByGovernanceLabelValues field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetEntitlementValueScopeByGovernanceLabelValues() []GovernanceLabelValuesInner {
	if o == nil || IsNil(o.EntitlementValueScopeByGovernanceLabelValues) {
		var ret []GovernanceLabelValuesInner
		return ret
	}
	return o.EntitlementValueScopeByGovernanceLabelValues
}

// GetEntitlementValueScopeByGovernanceLabelValuesOk returns a tuple with the EntitlementValueScopeByGovernanceLabelValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetEntitlementValueScopeByGovernanceLabelValuesOk() ([]GovernanceLabelValuesInner, bool) {
	if o == nil || IsNil(o.EntitlementValueScopeByGovernanceLabelValues) {
		return nil, false
	}
	return o.EntitlementValueScopeByGovernanceLabelValues, true
}

// HasEntitlementValueScopeByGovernanceLabelValues returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasEntitlementValueScopeByGovernanceLabelValues() bool {
	if o != nil && !IsNil(o.EntitlementValueScopeByGovernanceLabelValues) {
		return true
	}

	return false
}

// SetEntitlementValueScopeByGovernanceLabelValues gets a reference to the given []GovernanceLabelValuesInner and assigns it to the EntitlementValueScopeByGovernanceLabelValues field.
func (o *TargetResourcesRequestInner) SetEntitlementValueScopeByGovernanceLabelValues(v []GovernanceLabelValuesInner) {
	o.EntitlementValueScopeByGovernanceLabelValues = v
}

// GetIncludeBundlesHavingSameEntitlementValues returns the IncludeBundlesHavingSameEntitlementValues field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetIncludeBundlesHavingSameEntitlementValues() bool {
	if o == nil || IsNil(o.IncludeBundlesHavingSameEntitlementValues) {
		var ret bool
		return ret
	}
	return *o.IncludeBundlesHavingSameEntitlementValues
}

// GetIncludeBundlesHavingSameEntitlementValuesOk returns a tuple with the IncludeBundlesHavingSameEntitlementValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetIncludeBundlesHavingSameEntitlementValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeBundlesHavingSameEntitlementValues) {
		return nil, false
	}
	return o.IncludeBundlesHavingSameEntitlementValues, true
}

// HasIncludeBundlesHavingSameEntitlementValues returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasIncludeBundlesHavingSameEntitlementValues() bool {
	if o != nil && !IsNil(o.IncludeBundlesHavingSameEntitlementValues) {
		return true
	}

	return false
}

// SetIncludeBundlesHavingSameEntitlementValues gets a reference to the given bool and assigns it to the IncludeBundlesHavingSameEntitlementValues field.
func (o *TargetResourcesRequestInner) SetIncludeBundlesHavingSameEntitlementValues(v bool) {
	o.IncludeBundlesHavingSameEntitlementValues = &v
}

// GetEntitlementBundleScopeByGovernanceLabelValues returns the EntitlementBundleScopeByGovernanceLabelValues field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetEntitlementBundleScopeByGovernanceLabelValues() []GovernanceLabelValuesInner {
	if o == nil || IsNil(o.EntitlementBundleScopeByGovernanceLabelValues) {
		var ret []GovernanceLabelValuesInner
		return ret
	}
	return o.EntitlementBundleScopeByGovernanceLabelValues
}

// GetEntitlementBundleScopeByGovernanceLabelValuesOk returns a tuple with the EntitlementBundleScopeByGovernanceLabelValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetEntitlementBundleScopeByGovernanceLabelValuesOk() ([]GovernanceLabelValuesInner, bool) {
	if o == nil || IsNil(o.EntitlementBundleScopeByGovernanceLabelValues) {
		return nil, false
	}
	return o.EntitlementBundleScopeByGovernanceLabelValues, true
}

// HasEntitlementBundleScopeByGovernanceLabelValues returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasEntitlementBundleScopeByGovernanceLabelValues() bool {
	if o != nil && !IsNil(o.EntitlementBundleScopeByGovernanceLabelValues) {
		return true
	}

	return false
}

// SetEntitlementBundleScopeByGovernanceLabelValues gets a reference to the given []GovernanceLabelValuesInner and assigns it to the EntitlementBundleScopeByGovernanceLabelValues field.
func (o *TargetResourcesRequestInner) SetEntitlementBundleScopeByGovernanceLabelValues(v []GovernanceLabelValuesInner) {
	o.EntitlementBundleScopeByGovernanceLabelValues = v
}

// GetAppServiceAccounts returns the AppServiceAccounts field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetAppServiceAccounts() []AppServiceAccountsInner {
	if o == nil || IsNil(o.AppServiceAccounts) {
		var ret []AppServiceAccountsInner
		return ret
	}
	return o.AppServiceAccounts
}

// GetAppServiceAccountsOk returns a tuple with the AppServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetAppServiceAccountsOk() ([]AppServiceAccountsInner, bool) {
	if o == nil || IsNil(o.AppServiceAccounts) {
		return nil, false
	}
	return o.AppServiceAccounts, true
}

// HasAppServiceAccounts returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasAppServiceAccounts() bool {
	if o != nil && !IsNil(o.AppServiceAccounts) {
		return true
	}

	return false
}

// SetAppServiceAccounts gets a reference to the given []AppServiceAccountsInner and assigns it to the AppServiceAccounts field.
func (o *TargetResourcesRequestInner) SetAppServiceAccounts(v []AppServiceAccountsInner) {
	o.AppServiceAccounts = v
}

// GetIncludeAllAppServiceAccounts returns the IncludeAllAppServiceAccounts field value if set, zero value otherwise.
func (o *TargetResourcesRequestInner) GetIncludeAllAppServiceAccounts() bool {
	if o == nil || IsNil(o.IncludeAllAppServiceAccounts) {
		var ret bool
		return ret
	}
	return *o.IncludeAllAppServiceAccounts
}

// GetIncludeAllAppServiceAccountsOk returns a tuple with the IncludeAllAppServiceAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetResourcesRequestInner) GetIncludeAllAppServiceAccountsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAllAppServiceAccounts) {
		return nil, false
	}
	return o.IncludeAllAppServiceAccounts, true
}

// HasIncludeAllAppServiceAccounts returns a boolean if a field has been set.
func (o *TargetResourcesRequestInner) HasIncludeAllAppServiceAccounts() bool {
	if o != nil && !IsNil(o.IncludeAllAppServiceAccounts) {
		return true
	}

	return false
}

// SetIncludeAllAppServiceAccounts gets a reference to the given bool and assigns it to the IncludeAllAppServiceAccounts field.
func (o *TargetResourcesRequestInner) SetIncludeAllAppServiceAccounts(v bool) {
	o.IncludeAllAppServiceAccounts = &v
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
	if !IsNil(o.EntitlementValueScopeByGovernanceLabelValues) {
		toSerialize["entitlementValueScopeByGovernanceLabelValues"] = o.EntitlementValueScopeByGovernanceLabelValues
	}
	if !IsNil(o.IncludeBundlesHavingSameEntitlementValues) {
		toSerialize["includeBundlesHavingSameEntitlementValues"] = o.IncludeBundlesHavingSameEntitlementValues
	}
	if !IsNil(o.EntitlementBundleScopeByGovernanceLabelValues) {
		toSerialize["entitlementBundleScopeByGovernanceLabelValues"] = o.EntitlementBundleScopeByGovernanceLabelValues
	}
	if !IsNil(o.AppServiceAccounts) {
		toSerialize["appServiceAccounts"] = o.AppServiceAccounts
	}
	if !IsNil(o.IncludeAllAppServiceAccounts) {
		toSerialize["includeAllAppServiceAccounts"] = o.IncludeAllAppServiceAccounts
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
		delete(additionalProperties, "entitlementValueScopeByGovernanceLabelValues")
		delete(additionalProperties, "includeBundlesHavingSameEntitlementValues")
		delete(additionalProperties, "entitlementBundleScopeByGovernanceLabelValues")
		delete(additionalProperties, "appServiceAccounts")
		delete(additionalProperties, "includeAllAppServiceAccounts")
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
