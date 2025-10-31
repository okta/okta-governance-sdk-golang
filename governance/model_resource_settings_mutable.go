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

// checks if the ResourceSettingsMutable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSettingsMutable{}

// ResourceSettingsMutable Resource specific properties
type ResourceSettingsMutable struct {
	Type CampaignResourceType `json:"type"`
	//  Represents a resource that will be part of Access certifications. If the app is enabled for Access Certifications, it's possible to review entitlements and entitlement bundles.  You can review all entitlements by specifying `includeEntitlements: true` and/or restrict it by setting the property `onlyIncludeOutOfPolicyEntitlements: true`, both of which are `false` by default.  If `includeEntitlements: false`, you need to specify a list of `entitlementBundles` and/or `entitlements`.
	TargetResources []TargetResourcesRequestInner `json:"targetResources,omitempty"`
	// An array of resources that are excluded from the review
	ExcludedResources []ResourceSettingsMutableExcludedResourcesInner `json:"excludedResources,omitempty"`
	// Only include individually assigned apps. This is only applicable if campaign type is `USER`.
	IndividuallyAssignedAppsOnly *bool `json:"individuallyAssignedAppsOnly,omitempty"`
	// Only include individually assigned groups. This is only applicable if campaign type is `USER`.
	IndividuallyAssignedGroupsOnly *bool `json:"individuallyAssignedGroupsOnly,omitempty"`
	// Include entitlements for this application. This property is only applicable if `resourcetype = APPLICATION` and Entitlement Management is enabled.
	IncludeEntitlements *bool `json:"includeEntitlements,omitempty"`
	// Only include out-of-policy entitlements. Only applicable if `resourcetype = APPLICATION` and Entitlement Management is enabled.
	OnlyIncludeOutOfPolicyEntitlements *bool `json:"onlyIncludeOutOfPolicyEntitlements,omitempty"`
	// Include admin roles.
	IncludeAdminRoles    *bool `json:"includeAdminRoles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSettingsMutable ResourceSettingsMutable

// NewResourceSettingsMutable instantiates a new ResourceSettingsMutable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSettingsMutable(type_ CampaignResourceType) *ResourceSettingsMutable {
	this := ResourceSettingsMutable{}
	this.Type = type_
	return &this
}

// NewResourceSettingsMutableWithDefaults instantiates a new ResourceSettingsMutable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSettingsMutableWithDefaults() *ResourceSettingsMutable {
	this := ResourceSettingsMutable{}
	return &this
}

// GetType returns the Type field value
func (o *ResourceSettingsMutable) GetType() CampaignResourceType {
	if o == nil {
		var ret CampaignResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetTypeOk() (*CampaignResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceSettingsMutable) SetType(v CampaignResourceType) {
	o.Type = v
}

// GetTargetResources returns the TargetResources field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetTargetResources() []TargetResourcesRequestInner {
	if o == nil || IsNil(o.TargetResources) {
		var ret []TargetResourcesRequestInner
		return ret
	}
	return o.TargetResources
}

// GetTargetResourcesOk returns a tuple with the TargetResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetTargetResourcesOk() ([]TargetResourcesRequestInner, bool) {
	if o == nil || IsNil(o.TargetResources) {
		return nil, false
	}
	return o.TargetResources, true
}

// HasTargetResources returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasTargetResources() bool {
	if o != nil && !IsNil(o.TargetResources) {
		return true
	}

	return false
}

// SetTargetResources gets a reference to the given []TargetResourcesRequestInner and assigns it to the TargetResources field.
func (o *ResourceSettingsMutable) SetTargetResources(v []TargetResourcesRequestInner) {
	o.TargetResources = v
}

// GetExcludedResources returns the ExcludedResources field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetExcludedResources() []ResourceSettingsMutableExcludedResourcesInner {
	if o == nil || IsNil(o.ExcludedResources) {
		var ret []ResourceSettingsMutableExcludedResourcesInner
		return ret
	}
	return o.ExcludedResources
}

// GetExcludedResourcesOk returns a tuple with the ExcludedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetExcludedResourcesOk() ([]ResourceSettingsMutableExcludedResourcesInner, bool) {
	if o == nil || IsNil(o.ExcludedResources) {
		return nil, false
	}
	return o.ExcludedResources, true
}

// HasExcludedResources returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasExcludedResources() bool {
	if o != nil && !IsNil(o.ExcludedResources) {
		return true
	}

	return false
}

// SetExcludedResources gets a reference to the given []ResourceSettingsMutableExcludedResourcesInner and assigns it to the ExcludedResources field.
func (o *ResourceSettingsMutable) SetExcludedResources(v []ResourceSettingsMutableExcludedResourcesInner) {
	o.ExcludedResources = v
}

// GetIndividuallyAssignedAppsOnly returns the IndividuallyAssignedAppsOnly field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetIndividuallyAssignedAppsOnly() bool {
	if o == nil || IsNil(o.IndividuallyAssignedAppsOnly) {
		var ret bool
		return ret
	}
	return *o.IndividuallyAssignedAppsOnly
}

// GetIndividuallyAssignedAppsOnlyOk returns a tuple with the IndividuallyAssignedAppsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetIndividuallyAssignedAppsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IndividuallyAssignedAppsOnly) {
		return nil, false
	}
	return o.IndividuallyAssignedAppsOnly, true
}

// HasIndividuallyAssignedAppsOnly returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasIndividuallyAssignedAppsOnly() bool {
	if o != nil && !IsNil(o.IndividuallyAssignedAppsOnly) {
		return true
	}

	return false
}

// SetIndividuallyAssignedAppsOnly gets a reference to the given bool and assigns it to the IndividuallyAssignedAppsOnly field.
func (o *ResourceSettingsMutable) SetIndividuallyAssignedAppsOnly(v bool) {
	o.IndividuallyAssignedAppsOnly = &v
}

// GetIndividuallyAssignedGroupsOnly returns the IndividuallyAssignedGroupsOnly field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetIndividuallyAssignedGroupsOnly() bool {
	if o == nil || IsNil(o.IndividuallyAssignedGroupsOnly) {
		var ret bool
		return ret
	}
	return *o.IndividuallyAssignedGroupsOnly
}

// GetIndividuallyAssignedGroupsOnlyOk returns a tuple with the IndividuallyAssignedGroupsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetIndividuallyAssignedGroupsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IndividuallyAssignedGroupsOnly) {
		return nil, false
	}
	return o.IndividuallyAssignedGroupsOnly, true
}

// HasIndividuallyAssignedGroupsOnly returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasIndividuallyAssignedGroupsOnly() bool {
	if o != nil && !IsNil(o.IndividuallyAssignedGroupsOnly) {
		return true
	}

	return false
}

// SetIndividuallyAssignedGroupsOnly gets a reference to the given bool and assigns it to the IndividuallyAssignedGroupsOnly field.
func (o *ResourceSettingsMutable) SetIndividuallyAssignedGroupsOnly(v bool) {
	o.IndividuallyAssignedGroupsOnly = &v
}

// GetIncludeEntitlements returns the IncludeEntitlements field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetIncludeEntitlements() bool {
	if o == nil || IsNil(o.IncludeEntitlements) {
		var ret bool
		return ret
	}
	return *o.IncludeEntitlements
}

// GetIncludeEntitlementsOk returns a tuple with the IncludeEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetIncludeEntitlementsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeEntitlements) {
		return nil, false
	}
	return o.IncludeEntitlements, true
}

// HasIncludeEntitlements returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasIncludeEntitlements() bool {
	if o != nil && !IsNil(o.IncludeEntitlements) {
		return true
	}

	return false
}

// SetIncludeEntitlements gets a reference to the given bool and assigns it to the IncludeEntitlements field.
func (o *ResourceSettingsMutable) SetIncludeEntitlements(v bool) {
	o.IncludeEntitlements = &v
}

// GetOnlyIncludeOutOfPolicyEntitlements returns the OnlyIncludeOutOfPolicyEntitlements field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetOnlyIncludeOutOfPolicyEntitlements() bool {
	if o == nil || IsNil(o.OnlyIncludeOutOfPolicyEntitlements) {
		var ret bool
		return ret
	}
	return *o.OnlyIncludeOutOfPolicyEntitlements
}

// GetOnlyIncludeOutOfPolicyEntitlementsOk returns a tuple with the OnlyIncludeOutOfPolicyEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetOnlyIncludeOutOfPolicyEntitlementsOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyIncludeOutOfPolicyEntitlements) {
		return nil, false
	}
	return o.OnlyIncludeOutOfPolicyEntitlements, true
}

// HasOnlyIncludeOutOfPolicyEntitlements returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasOnlyIncludeOutOfPolicyEntitlements() bool {
	if o != nil && !IsNil(o.OnlyIncludeOutOfPolicyEntitlements) {
		return true
	}

	return false
}

// SetOnlyIncludeOutOfPolicyEntitlements gets a reference to the given bool and assigns it to the OnlyIncludeOutOfPolicyEntitlements field.
func (o *ResourceSettingsMutable) SetOnlyIncludeOutOfPolicyEntitlements(v bool) {
	o.OnlyIncludeOutOfPolicyEntitlements = &v
}

// GetIncludeAdminRoles returns the IncludeAdminRoles field value if set, zero value otherwise.
func (o *ResourceSettingsMutable) GetIncludeAdminRoles() bool {
	if o == nil || IsNil(o.IncludeAdminRoles) {
		var ret bool
		return ret
	}
	return *o.IncludeAdminRoles
}

// GetIncludeAdminRolesOk returns a tuple with the IncludeAdminRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSettingsMutable) GetIncludeAdminRolesOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAdminRoles) {
		return nil, false
	}
	return o.IncludeAdminRoles, true
}

// HasIncludeAdminRoles returns a boolean if a field has been set.
func (o *ResourceSettingsMutable) HasIncludeAdminRoles() bool {
	if o != nil && !IsNil(o.IncludeAdminRoles) {
		return true
	}

	return false
}

// SetIncludeAdminRoles gets a reference to the given bool and assigns it to the IncludeAdminRoles field.
func (o *ResourceSettingsMutable) SetIncludeAdminRoles(v bool) {
	o.IncludeAdminRoles = &v
}

func (o ResourceSettingsMutable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSettingsMutable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.TargetResources) {
		toSerialize["targetResources"] = o.TargetResources
	}
	if !IsNil(o.ExcludedResources) {
		toSerialize["excludedResources"] = o.ExcludedResources
	}
	if !IsNil(o.IndividuallyAssignedAppsOnly) {
		toSerialize["individuallyAssignedAppsOnly"] = o.IndividuallyAssignedAppsOnly
	}
	if !IsNil(o.IndividuallyAssignedGroupsOnly) {
		toSerialize["individuallyAssignedGroupsOnly"] = o.IndividuallyAssignedGroupsOnly
	}
	if !IsNil(o.IncludeEntitlements) {
		toSerialize["includeEntitlements"] = o.IncludeEntitlements
	}
	if !IsNil(o.OnlyIncludeOutOfPolicyEntitlements) {
		toSerialize["onlyIncludeOutOfPolicyEntitlements"] = o.OnlyIncludeOutOfPolicyEntitlements
	}
	if !IsNil(o.IncludeAdminRoles) {
		toSerialize["includeAdminRoles"] = o.IncludeAdminRoles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSettingsMutable) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varResourceSettingsMutable := _ResourceSettingsMutable{}

	err = json.Unmarshal(data, &varResourceSettingsMutable)

	if err != nil {
		return err
	}

	*o = ResourceSettingsMutable(varResourceSettingsMutable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "targetResources")
		delete(additionalProperties, "excludedResources")
		delete(additionalProperties, "individuallyAssignedAppsOnly")
		delete(additionalProperties, "individuallyAssignedGroupsOnly")
		delete(additionalProperties, "includeEntitlements")
		delete(additionalProperties, "onlyIncludeOutOfPolicyEntitlements")
		delete(additionalProperties, "includeAdminRoles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSettingsMutable struct {
	value *ResourceSettingsMutable
	isSet bool
}

func (v NullableResourceSettingsMutable) Get() *ResourceSettingsMutable {
	return v.value
}

func (v *NullableResourceSettingsMutable) Set(val *ResourceSettingsMutable) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSettingsMutable) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSettingsMutable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSettingsMutable(val *ResourceSettingsMutable) *NullableResourceSettingsMutable {
	return &NullableResourceSettingsMutable{value: val, isSet: true}
}

func (v NullableResourceSettingsMutable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSettingsMutable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
