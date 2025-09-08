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
	"time"
)

// SecurityAccessReviewAccessItemAppInfo struct for SecurityAccessReviewAccessItemAppInfo
type SecurityAccessReviewAccessItemAppInfo struct {
	// The Okta app ID
	Id string `json:"id"`
	// The app name
	Name string `json:"name"`
	// The app logo URL
	LogoUrl string `json:"logoUrl"`
	// The app label
	Label *string `json:"label,omitempty"`
	// The last time the app was accessed by the user
	LastAccessDate *time.Time `json:"lastAccessDate,omitempty"`
	// The date the app was assigned to the user
	AssignedDate   *time.Time      `json:"assignedDate,omitempty"`
	AssignmentType *AssignmentType `json:"assignmentType,omitempty"`
	// The number of times the user has used the app in the last 90 days
	ApplicationUsage *int64 `json:"applicationUsage,omitempty"`
	// The last time the user/app pair was reviewed in an access certification campaign
	LastAccessCertificationReviewedDate *time.Time `json:"lastAccessCertificationReviewedDate,omitempty"`
	// The last time an action was taken on this app for the user in a security access review
	LastSecurityAccessReviewDate *time.Time `json:"lastSecurityAccessReviewDate,omitempty"`
	// Information about the last device used to log into the app
	LastDeviceLogin *string `json:"lastDeviceLogin,omitempty"`
	// Any active entitlements for this app
	ActiveEntitlements []EntitlementPropertyFull `json:"activeEntitlements,omitempty"`
	// All governance labels applied to the app
	GovernanceLabels     []GovernanceLabel `json:"governanceLabels,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewAccessItemAppInfo SecurityAccessReviewAccessItemAppInfo

// NewSecurityAccessReviewAccessItemAppInfo instantiates a new SecurityAccessReviewAccessItemAppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewAccessItemAppInfo(id string, name string, logoUrl string) *SecurityAccessReviewAccessItemAppInfo {
	this := SecurityAccessReviewAccessItemAppInfo{}
	this.Id = id
	this.Name = name
	this.LogoUrl = logoUrl
	return &this
}

// NewSecurityAccessReviewAccessItemAppInfoWithDefaults instantiates a new SecurityAccessReviewAccessItemAppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewAccessItemAppInfoWithDefaults() *SecurityAccessReviewAccessItemAppInfo {
	this := SecurityAccessReviewAccessItemAppInfo{}
	return &this
}

// GetId returns the Id field value
func (o *SecurityAccessReviewAccessItemAppInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SecurityAccessReviewAccessItemAppInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SecurityAccessReviewAccessItemAppInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SecurityAccessReviewAccessItemAppInfo) SetName(v string) {
	o.Name = v
}

// GetLogoUrl returns the LogoUrl field value
func (o *SecurityAccessReviewAccessItemAppInfo) GetLogoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoUrl, true
}

// SetLogoUrl sets field value
func (o *SecurityAccessReviewAccessItemAppInfo) SetLogoUrl(v string) {
	o.LogoUrl = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetLabel(v string) {
	o.Label = &v
}

// GetLastAccessDate returns the LastAccessDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessDate() time.Time {
	if o == nil || o.LastAccessDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessDate
}

// GetLastAccessDateOk returns a tuple with the LastAccessDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessDateOk() (*time.Time, bool) {
	if o == nil || o.LastAccessDate == nil {
		return nil, false
	}
	return o.LastAccessDate, true
}

// HasLastAccessDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasLastAccessDate() bool {
	if o != nil && o.LastAccessDate != nil {
		return true
	}

	return false
}

// SetLastAccessDate gets a reference to the given time.Time and assigns it to the LastAccessDate field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetLastAccessDate(v time.Time) {
	o.LastAccessDate = &v
}

// GetAssignedDate returns the AssignedDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignedDate() time.Time {
	if o == nil || o.AssignedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AssignedDate
}

// GetAssignedDateOk returns a tuple with the AssignedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignedDateOk() (*time.Time, bool) {
	if o == nil || o.AssignedDate == nil {
		return nil, false
	}
	return o.AssignedDate, true
}

// HasAssignedDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasAssignedDate() bool {
	if o != nil && o.AssignedDate != nil {
		return true
	}

	return false
}

// SetAssignedDate gets a reference to the given time.Time and assigns it to the AssignedDate field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetAssignedDate(v time.Time) {
	o.AssignedDate = &v
}

// GetAssignmentType returns the AssignmentType field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignmentType() AssignmentType {
	if o == nil || o.AssignmentType == nil {
		var ret AssignmentType
		return ret
	}
	return *o.AssignmentType
}

// GetAssignmentTypeOk returns a tuple with the AssignmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetAssignmentTypeOk() (*AssignmentType, bool) {
	if o == nil || o.AssignmentType == nil {
		return nil, false
	}
	return o.AssignmentType, true
}

// HasAssignmentType returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasAssignmentType() bool {
	if o != nil && o.AssignmentType != nil {
		return true
	}

	return false
}

// SetAssignmentType gets a reference to the given AssignmentType and assigns it to the AssignmentType field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetAssignmentType(v AssignmentType) {
	o.AssignmentType = &v
}

// GetApplicationUsage returns the ApplicationUsage field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetApplicationUsage() int64 {
	if o == nil || o.ApplicationUsage == nil {
		var ret int64
		return ret
	}
	return *o.ApplicationUsage
}

// GetApplicationUsageOk returns a tuple with the ApplicationUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetApplicationUsageOk() (*int64, bool) {
	if o == nil || o.ApplicationUsage == nil {
		return nil, false
	}
	return o.ApplicationUsage, true
}

// HasApplicationUsage returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasApplicationUsage() bool {
	if o != nil && o.ApplicationUsage != nil {
		return true
	}

	return false
}

// SetApplicationUsage gets a reference to the given int64 and assigns it to the ApplicationUsage field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetApplicationUsage(v int64) {
	o.ApplicationUsage = &v
}

// GetLastAccessCertificationReviewedDate returns the LastAccessCertificationReviewedDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessCertificationReviewedDate() time.Time {
	if o == nil || o.LastAccessCertificationReviewedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessCertificationReviewedDate
}

// GetLastAccessCertificationReviewedDateOk returns a tuple with the LastAccessCertificationReviewedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastAccessCertificationReviewedDateOk() (*time.Time, bool) {
	if o == nil || o.LastAccessCertificationReviewedDate == nil {
		return nil, false
	}
	return o.LastAccessCertificationReviewedDate, true
}

// HasLastAccessCertificationReviewedDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasLastAccessCertificationReviewedDate() bool {
	if o != nil && o.LastAccessCertificationReviewedDate != nil {
		return true
	}

	return false
}

// SetLastAccessCertificationReviewedDate gets a reference to the given time.Time and assigns it to the LastAccessCertificationReviewedDate field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetLastAccessCertificationReviewedDate(v time.Time) {
	o.LastAccessCertificationReviewedDate = &v
}

// GetLastSecurityAccessReviewDate returns the LastSecurityAccessReviewDate field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastSecurityAccessReviewDate() time.Time {
	if o == nil || o.LastSecurityAccessReviewDate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSecurityAccessReviewDate
}

// GetLastSecurityAccessReviewDateOk returns a tuple with the LastSecurityAccessReviewDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastSecurityAccessReviewDateOk() (*time.Time, bool) {
	if o == nil || o.LastSecurityAccessReviewDate == nil {
		return nil, false
	}
	return o.LastSecurityAccessReviewDate, true
}

// HasLastSecurityAccessReviewDate returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasLastSecurityAccessReviewDate() bool {
	if o != nil && o.LastSecurityAccessReviewDate != nil {
		return true
	}

	return false
}

// SetLastSecurityAccessReviewDate gets a reference to the given time.Time and assigns it to the LastSecurityAccessReviewDate field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetLastSecurityAccessReviewDate(v time.Time) {
	o.LastSecurityAccessReviewDate = &v
}

// GetLastDeviceLogin returns the LastDeviceLogin field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastDeviceLogin() string {
	if o == nil || o.LastDeviceLogin == nil {
		var ret string
		return ret
	}
	return *o.LastDeviceLogin
}

// GetLastDeviceLoginOk returns a tuple with the LastDeviceLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetLastDeviceLoginOk() (*string, bool) {
	if o == nil || o.LastDeviceLogin == nil {
		return nil, false
	}
	return o.LastDeviceLogin, true
}

// HasLastDeviceLogin returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasLastDeviceLogin() bool {
	if o != nil && o.LastDeviceLogin != nil {
		return true
	}

	return false
}

// SetLastDeviceLogin gets a reference to the given string and assigns it to the LastDeviceLogin field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetLastDeviceLogin(v string) {
	o.LastDeviceLogin = &v
}

// GetActiveEntitlements returns the ActiveEntitlements field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetActiveEntitlements() []EntitlementPropertyFull {
	if o == nil || o.ActiveEntitlements == nil {
		var ret []EntitlementPropertyFull
		return ret
	}
	return o.ActiveEntitlements
}

// GetActiveEntitlementsOk returns a tuple with the ActiveEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetActiveEntitlementsOk() ([]EntitlementPropertyFull, bool) {
	if o == nil || o.ActiveEntitlements == nil {
		return nil, false
	}
	return o.ActiveEntitlements, true
}

// HasActiveEntitlements returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasActiveEntitlements() bool {
	if o != nil && o.ActiveEntitlements != nil {
		return true
	}

	return false
}

// SetActiveEntitlements gets a reference to the given []EntitlementPropertyFull and assigns it to the ActiveEntitlements field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetActiveEntitlements(v []EntitlementPropertyFull) {
	o.ActiveEntitlements = v
}

// GetGovernanceLabels returns the GovernanceLabels field value if set, zero value otherwise.
func (o *SecurityAccessReviewAccessItemAppInfo) GetGovernanceLabels() []GovernanceLabel {
	if o == nil || o.GovernanceLabels == nil {
		var ret []GovernanceLabel
		return ret
	}
	return o.GovernanceLabels
}

// GetGovernanceLabelsOk returns a tuple with the GovernanceLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) GetGovernanceLabelsOk() ([]GovernanceLabel, bool) {
	if o == nil || o.GovernanceLabels == nil {
		return nil, false
	}
	return o.GovernanceLabels, true
}

// HasGovernanceLabels returns a boolean if a field has been set.
func (o *SecurityAccessReviewAccessItemAppInfo) HasGovernanceLabels() bool {
	if o != nil && o.GovernanceLabels != nil {
		return true
	}

	return false
}

// SetGovernanceLabels gets a reference to the given []GovernanceLabel and assigns it to the GovernanceLabels field.
func (o *SecurityAccessReviewAccessItemAppInfo) SetGovernanceLabels(v []GovernanceLabel) {
	o.GovernanceLabels = v
}

func (o SecurityAccessReviewAccessItemAppInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LastAccessDate != nil {
		toSerialize["lastAccessDate"] = o.LastAccessDate
	}
	if o.AssignedDate != nil {
		toSerialize["assignedDate"] = o.AssignedDate
	}
	if o.AssignmentType != nil {
		toSerialize["assignmentType"] = o.AssignmentType
	}
	if o.ApplicationUsage != nil {
		toSerialize["applicationUsage"] = o.ApplicationUsage
	}
	if o.LastAccessCertificationReviewedDate != nil {
		toSerialize["lastAccessCertificationReviewedDate"] = o.LastAccessCertificationReviewedDate
	}
	if o.LastSecurityAccessReviewDate != nil {
		toSerialize["lastSecurityAccessReviewDate"] = o.LastSecurityAccessReviewDate
	}
	if o.LastDeviceLogin != nil {
		toSerialize["lastDeviceLogin"] = o.LastDeviceLogin
	}
	if o.ActiveEntitlements != nil {
		toSerialize["activeEntitlements"] = o.ActiveEntitlements
	}
	if o.GovernanceLabels != nil {
		toSerialize["governanceLabels"] = o.GovernanceLabels
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewAccessItemAppInfo) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewAccessItemAppInfo := _SecurityAccessReviewAccessItemAppInfo{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewAccessItemAppInfo)
	if err == nil {
		*o = SecurityAccessReviewAccessItemAppInfo(varSecurityAccessReviewAccessItemAppInfo)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "logoUrl")
		delete(additionalProperties, "label")
		delete(additionalProperties, "lastAccessDate")
		delete(additionalProperties, "assignedDate")
		delete(additionalProperties, "assignmentType")
		delete(additionalProperties, "applicationUsage")
		delete(additionalProperties, "lastAccessCertificationReviewedDate")
		delete(additionalProperties, "lastSecurityAccessReviewDate")
		delete(additionalProperties, "lastDeviceLogin")
		delete(additionalProperties, "activeEntitlements")
		delete(additionalProperties, "governanceLabels")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewAccessItemAppInfo struct {
	value *SecurityAccessReviewAccessItemAppInfo
	isSet bool
}

func (v NullableSecurityAccessReviewAccessItemAppInfo) Get() *SecurityAccessReviewAccessItemAppInfo {
	return v.value
}

func (v *NullableSecurityAccessReviewAccessItemAppInfo) Set(val *SecurityAccessReviewAccessItemAppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAccessItemAppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAccessItemAppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAccessItemAppInfo(val *SecurityAccessReviewAccessItemAppInfo) *NullableSecurityAccessReviewAccessItemAppInfo {
	return &NullableSecurityAccessReviewAccessItemAppInfo{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAccessItemAppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAccessItemAppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
