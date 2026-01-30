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
	"time"
)

// checks if the SecurityAccessReviewSubAccessItemEntitlementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityAccessReviewSubAccessItemEntitlementInfo{}

// SecurityAccessReviewSubAccessItemEntitlementInfo The entitlement details for the sub-access item
type SecurityAccessReviewSubAccessItemEntitlementInfo struct {
	Type SecurityAccessReviewSubAccessItemEntitlementType `json:"type"`
	// A brief description of the entitlement value or bundle
	Description *string `json:"description,omitempty"`
	// A brief description of the entitlement associated with the value
	EntitlementDescription *string `json:"entitlementDescription,omitempty"`
	// The date the entitlement or bundle was assigned to the user
	AssignedDate   *time.Time      `json:"assignedDate,omitempty"`
	AssignmentType *AssignmentType `json:"assignmentType,omitempty"`
	// Collections assigning this resource
	CollectionsAssigning []CollectionInfoSparse `json:"collectionsAssigning,omitempty"`
	// The entitlements included in the entitlement bundle
	Entitlements []EntitlementPropertyFull `json:"entitlements,omitempty"`
	// The last time the user/entitlement pair was reviewed in an access certification campaign
	LastAccessCertificationReviewedDate *time.Time `json:"lastAccessCertificationReviewedDate,omitempty"`
	// The last time an action was taken on this entitlement for the user in a security access review
	LastSecurityAccessReviewDate *time.Time `json:"lastSecurityAccessReviewDate,omitempty"`
	// All governance labels applied to the entitlement value or bundle
	GovernanceLabels     []TargetGovernanceLabel `json:"governanceLabels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewSubAccessItemEntitlementInfo SecurityAccessReviewSubAccessItemEntitlementInfo

// NewSecurityAccessReviewSubAccessItemEntitlementInfo instantiates a new SecurityAccessReviewSubAccessItemEntitlementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewSubAccessItemEntitlementInfo(type_ SecurityAccessReviewSubAccessItemEntitlementType) *SecurityAccessReviewSubAccessItemEntitlementInfo {
	this := SecurityAccessReviewSubAccessItemEntitlementInfo{}
	this.Type = type_
	return &this
}

// NewSecurityAccessReviewSubAccessItemEntitlementInfoWithDefaults instantiates a new SecurityAccessReviewSubAccessItemEntitlementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewSubAccessItemEntitlementInfoWithDefaults() *SecurityAccessReviewSubAccessItemEntitlementInfo {
	this := SecurityAccessReviewSubAccessItemEntitlementInfo{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetType() SecurityAccessReviewSubAccessItemEntitlementType {
	if o == nil {
		var ret SecurityAccessReviewSubAccessItemEntitlementType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetTypeOk() (*SecurityAccessReviewSubAccessItemEntitlementType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetType(v SecurityAccessReviewSubAccessItemEntitlementType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEntitlementDescription returns the EntitlementDescription field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlementDescription() string {
	if o == nil || IsNil(o.EntitlementDescription) {
		var ret string
		return ret
	}
	return *o.EntitlementDescription
}

// GetEntitlementDescriptionOk returns a tuple with the EntitlementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementDescription) {
		return nil, false
	}
	return o.EntitlementDescription, true
}

// HasEntitlementDescription returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasEntitlementDescription() bool {
	if o != nil && !IsNil(o.EntitlementDescription) {
		return true
	}

	return false
}

// SetEntitlementDescription gets a reference to the given string and assigns it to the EntitlementDescription field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetEntitlementDescription(v string) {
	o.EntitlementDescription = &v
}

// GetAssignedDate returns the AssignedDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignedDate() time.Time {
	if o == nil || IsNil(o.AssignedDate) {
		var ret time.Time
		return ret
	}
	return *o.AssignedDate
}

// GetAssignedDateOk returns a tuple with the AssignedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AssignedDate) {
		return nil, false
	}
	return o.AssignedDate, true
}

// HasAssignedDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasAssignedDate() bool {
	if o != nil && !IsNil(o.AssignedDate) {
		return true
	}

	return false
}

// SetAssignedDate gets a reference to the given time.Time and assigns it to the AssignedDate field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetAssignedDate(v time.Time) {
	o.AssignedDate = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignmentType() AssignmentType {
	if o == nil || IsNil(o.AssignmentType) {
		var ret AssignmentType
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetAssignmentTypeOk() (*AssignmentType, bool) {
	if o == nil || IsNil(o.AssignmentType) {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasAssignmentType() bool {
	if o != nil && !IsNil(o.AssignmentType) {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given AssignmentType and assigns it to the AssignmentType field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetAssignmentType(v AssignmentType) {
	o.AssignmentType = &v
}

// GetCollectionsAssigning returns the CollectionsAssigning field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetCollectionsAssigning() []CollectionInfoSparse {
	if o == nil || IsNil(o.CollectionsAssigning) {
		var ret []CollectionInfoSparse
		return ret
	}
	return o.CollectionsAssigning
}

// GetCollectionsAssigningOk returns a tuple with the CollectionsAssigning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetCollectionsAssigningOk() ([]CollectionInfoSparse, bool) {
	if o == nil || IsNil(o.CollectionsAssigning) {
		return nil, false
	}
	return o.CollectionsAssigning, true
}

// HasCollectionsAssigning returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasCollectionsAssigning() bool {
	if o != nil && !IsNil(o.CollectionsAssigning) {
		return true
	}

	return false
}

// SetCollectionsAssigning gets a reference to the given []CollectionInfoSparse and assigns it to the CollectionsAssigning field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetCollectionsAssigning(v []CollectionInfoSparse) {
	o.CollectionsAssigning = v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlements() []EntitlementPropertyFull {
	if o == nil || IsNil(o.Entitlements) {
		var ret []EntitlementPropertyFull
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetEntitlementsOk() ([]EntitlementPropertyFull, bool) {
	if o == nil || IsNil(o.Entitlements) {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasEntitlements() bool {
	if o != nil && !IsNil(o.Entitlements) {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementPropertyFull and assigns it to the Entitlements field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetEntitlements(v []EntitlementPropertyFull) {
	o.Entitlements = v
}

// GetLastAccessCertificationReviewedDate returns the LastAccessCertificationReviewedDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastAccessCertificationReviewedDate() time.Time {
	if o == nil || IsNil(o.LastAccessCertificationReviewedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastAccessCertificationReviewedDate
}

// GetLastAccessCertificationReviewedDateOk returns a tuple with the LastAccessCertificationReviewedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastAccessCertificationReviewedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAccessCertificationReviewedDate) {
		return nil, false
	}
	return o.LastAccessCertificationReviewedDate, true
}

// HasLastAccessCertificationReviewedDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasLastAccessCertificationReviewedDate() bool {
	if o != nil && !IsNil(o.LastAccessCertificationReviewedDate) {
		return true
	}

	return false
}

// SetLastAccessCertificationReviewedDate gets a reference to the given time.Time and assigns it to the LastAccessCertificationReviewedDate field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetLastAccessCertificationReviewedDate(v time.Time) {
	o.LastAccessCertificationReviewedDate = &v
}

// GetLastSecurityAccessReviewDate returns the LastSecurityAccessReviewDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastSecurityAccessReviewDate() time.Time {
	if o == nil || IsNil(o.LastSecurityAccessReviewDate) {
		var ret time.Time
		return ret
	}
	return *o.LastSecurityAccessReviewDate
}

// GetLastSecurityAccessReviewDateOk returns a tuple with the LastSecurityAccessReviewDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetLastSecurityAccessReviewDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSecurityAccessReviewDate) {
		return nil, false
	}
	return o.LastSecurityAccessReviewDate, true
}

// HasLastSecurityAccessReviewDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasLastSecurityAccessReviewDate() bool {
	if o != nil && !IsNil(o.LastSecurityAccessReviewDate) {
		return true
	}

	return false
}

// SetLastSecurityAccessReviewDate gets a reference to the given time.Time and assigns it to the LastSecurityAccessReviewDate field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetLastSecurityAccessReviewDate(v time.Time) {
	o.LastSecurityAccessReviewDate = &v
}

// GetGovernanceLabels returns the GovernanceLabels field value if set, zero value otherwise.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetGovernanceLabels() []TargetGovernanceLabel {
	if o == nil || IsNil(o.GovernanceLabels) {
		var ret []TargetGovernanceLabel
		return ret
	}
	return o.GovernanceLabels
}

// GetGovernanceLabelsOk returns a tuple with the GovernanceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) GetGovernanceLabelsOk() ([]TargetGovernanceLabel, bool) {
	if o == nil || IsNil(o.GovernanceLabels) {
		return nil, false
	}
	return o.GovernanceLabels, true
}

// HasGovernanceLabels returns a boolean if a field has been set.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) HasGovernanceLabels() bool {
	if o != nil && !IsNil(o.GovernanceLabels) {
		return true
	}

	return false
}

// SetGovernanceLabels gets a reference to the given []TargetGovernanceLabel and assigns it to the GovernanceLabels field.
func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) SetGovernanceLabels(v []TargetGovernanceLabel) {
	o.GovernanceLabels = v
}

func (o SecurityAccessReviewSubAccessItemEntitlementInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityAccessReviewSubAccessItemEntitlementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EntitlementDescription) {
		toSerialize["entitlementDescription"] = o.EntitlementDescription
	}
	if !IsNil(o.AssignedDate) {
		toSerialize["assignedDate"] = o.AssignedDate
	}
	if !IsNil(o.AssignmentType) {
		toSerialize["assignmentType"] = o.AssignmentType
	}
	if !IsNil(o.CollectionsAssigning) {
		toSerialize["collectionsAssigning"] = o.CollectionsAssigning
	}
	if !IsNil(o.Entitlements) {
		toSerialize["entitlements"] = o.Entitlements
	}
	if !IsNil(o.LastAccessCertificationReviewedDate) {
		toSerialize["lastAccessCertificationReviewedDate"] = o.LastAccessCertificationReviewedDate
	}
	if !IsNil(o.LastSecurityAccessReviewDate) {
		toSerialize["lastSecurityAccessReviewDate"] = o.LastSecurityAccessReviewDate
	}
	if !IsNil(o.GovernanceLabels) {
		toSerialize["governanceLabels"] = o.GovernanceLabels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityAccessReviewSubAccessItemEntitlementInfo) UnmarshalJSON(data []byte) (err error) {
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

	varSecurityAccessReviewSubAccessItemEntitlementInfo := _SecurityAccessReviewSubAccessItemEntitlementInfo{}

	err = json.Unmarshal(data, &varSecurityAccessReviewSubAccessItemEntitlementInfo)

	if err != nil {
		return err
	}

	*o = SecurityAccessReviewSubAccessItemEntitlementInfo(varSecurityAccessReviewSubAccessItemEntitlementInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "entitlementDescription")
		delete(additionalProperties, "assignedDate")
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "collectionsAssigning")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "lastAccessCertificationReviewedDate")
		delete(additionalProperties, "lastSecurityAccessReviewDate")
		delete(additionalProperties, "governanceLabels")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityAccessReviewSubAccessItemEntitlementInfo struct {
	value *SecurityAccessReviewSubAccessItemEntitlementInfo
	isSet bool
}

func (v NullableSecurityAccessReviewSubAccessItemEntitlementInfo) Get() *SecurityAccessReviewSubAccessItemEntitlementInfo {
	return v.value
}

func (v *NullableSecurityAccessReviewSubAccessItemEntitlementInfo) Set(val *SecurityAccessReviewSubAccessItemEntitlementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewSubAccessItemEntitlementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewSubAccessItemEntitlementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewSubAccessItemEntitlementInfo(val *SecurityAccessReviewSubAccessItemEntitlementInfo) *NullableSecurityAccessReviewSubAccessItemEntitlementInfo {
	return &NullableSecurityAccessReviewSubAccessItemEntitlementInfo{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewSubAccessItemEntitlementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewSubAccessItemEntitlementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
