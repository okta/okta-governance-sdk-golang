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

// checks if the RequestFullCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestFullCommon{}

// RequestFullCommon Shared base for admin and enduser request-full models.
type RequestFullCommon struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The `id` of the Okta user who created the resource
	CreatedBy string `json:"createdBy"`
	// The ISO 8601 formatted date and time when the resource was created
	Created time.Time `json:"created"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy string        `json:"lastUpdatedBy"`
	Links         RequestLinks2 `json:"_links"`
	Status        RequestStatus `json:"status"`
	// The date the request was resolved. The property may transition from having a value to null if the request is reopened.
	Resolved    NullableTime        `json:"resolved,omitempty"`
	GrantStatus *RequestGrantStatus `json:"grantStatus,omitempty"`
	// The date the approved access was granted. Only set if request.status is APPROVED.
	Granted          NullableTime             `json:"granted,omitempty"`
	RevocationStatus *RequestRevocationStatus `json:"revocationStatus,omitempty"`
	// The date the granted access was revoked. Only set if request.grantStatus is GRANTED and request.revocationStatus is REVOKED.
	Revoked      NullableTime              `json:"revoked,omitempty"`
	RequestedBy  ClientCredentialPrincipal `json:"requestedBy"`
	RequestedFor TargetPrincipal           `json:"requestedFor"`
	Requested    Requested                 `json:"requested"`
	// How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  > **Note:** Mixes of units, as well as month/year/week designations, aren't supported. For example, `P40DT65H`, `P40M`, `P1W`, and `P1Y` aren't supported.
	AccessDuration NullableString `json:"accessDuration,omitempty"`
	// The date the granted access is scheduled for recovation. Only set if request.accessDuration exists, and request.grantStatus is GRANTED.
	RevocationScheduled NullableTime `json:"revocationScheduled,omitempty"`
	// The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.
	RequesterFieldValues []RequestFieldValue `json:"requesterFieldValues,omitempty"`
	RequestApproval      *RequestApproval2   `json:"requestApproval,omitempty"`
	RiskAssessment       *RiskAssessment     `json:"riskAssessment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestFullCommon RequestFullCommon

// NewRequestFullCommon instantiates a new RequestFullCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestFullCommon(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks2, status RequestStatus, requestedBy ClientCredentialPrincipal, requestedFor TargetPrincipal, requested Requested) *RequestFullCommon {
	this := RequestFullCommon{}
	this.Id = id
	this.CreatedBy = createdBy
	this.Created = created
	this.LastUpdated = lastUpdated
	this.LastUpdatedBy = lastUpdatedBy
	this.Links = links
	this.Status = status
	this.RequestedBy = requestedBy
	this.RequestedFor = requestedFor
	this.Requested = requested
	return &this
}

// NewRequestFullCommonWithDefaults instantiates a new RequestFullCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestFullCommonWithDefaults() *RequestFullCommon {
	this := RequestFullCommon{}
	return &this
}

// GetId returns the Id field value
func (o *RequestFullCommon) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestFullCommon) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestFullCommon) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestFullCommon) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestFullCommon) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestFullCommon) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestFullCommon) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestFullCommon) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestFullCommon) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestFullCommon) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *RequestFullCommon) GetLinks() RequestLinks2 {
	if o == nil {
		var ret RequestLinks2
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetLinksOk() (*RequestLinks2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestFullCommon) SetLinks(v RequestLinks2) {
	o.Links = v
}

// GetStatus returns the Status field value
func (o *RequestFullCommon) GetStatus() RequestStatus {
	if o == nil {
		var ret RequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetStatusOk() (*RequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestFullCommon) SetStatus(v RequestStatus) {
	o.Status = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestFullCommon) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullCommon) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *RequestFullCommon) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableTime and assigns it to the Resolved field.
func (o *RequestFullCommon) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *RequestFullCommon) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *RequestFullCommon) UnsetResolved() {
	o.Resolved.Unset()
}

// GetGrantStatus returns the GrantStatus field value if set, zero value otherwise.
func (o *RequestFullCommon) GetGrantStatus() RequestGrantStatus {
	if o == nil || IsNil(o.GrantStatus) {
		var ret RequestGrantStatus
		return ret
	}
	return *o.GrantStatus
}

// GetGrantStatusOk returns a tuple with the GrantStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetGrantStatusOk() (*RequestGrantStatus, bool) {
	if o == nil || IsNil(o.GrantStatus) {
		return nil, false
	}
	return o.GrantStatus, true
}

// HasGrantStatus returns a boolean if a field has been set.
func (o *RequestFullCommon) HasGrantStatus() bool {
	if o != nil && !IsNil(o.GrantStatus) {
		return true
	}

	return false
}

// SetGrantStatus gets a reference to the given RequestGrantStatus and assigns it to the GrantStatus field.
func (o *RequestFullCommon) SetGrantStatus(v RequestGrantStatus) {
	o.GrantStatus = &v
}

// GetGranted returns the Granted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestFullCommon) GetGranted() time.Time {
	if o == nil || IsNil(o.Granted.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Granted.Get()
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullCommon) GetGrantedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Granted.Get(), o.Granted.IsSet()
}

// HasGranted returns a boolean if a field has been set.
func (o *RequestFullCommon) HasGranted() bool {
	if o != nil && o.Granted.IsSet() {
		return true
	}

	return false
}

// SetGranted gets a reference to the given NullableTime and assigns it to the Granted field.
func (o *RequestFullCommon) SetGranted(v time.Time) {
	o.Granted.Set(&v)
}

// SetGrantedNil sets the value for Granted to be an explicit nil
func (o *RequestFullCommon) SetGrantedNil() {
	o.Granted.Set(nil)
}

// UnsetGranted ensures that no value is present for Granted, not even an explicit nil
func (o *RequestFullCommon) UnsetGranted() {
	o.Granted.Unset()
}

// GetRevocationStatus returns the RevocationStatus field value if set, zero value otherwise.
func (o *RequestFullCommon) GetRevocationStatus() RequestRevocationStatus {
	if o == nil || IsNil(o.RevocationStatus) {
		var ret RequestRevocationStatus
		return ret
	}
	return *o.RevocationStatus
}

// GetRevocationStatusOk returns a tuple with the RevocationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRevocationStatusOk() (*RequestRevocationStatus, bool) {
	if o == nil || IsNil(o.RevocationStatus) {
		return nil, false
	}
	return o.RevocationStatus, true
}

// HasRevocationStatus returns a boolean if a field has been set.
func (o *RequestFullCommon) HasRevocationStatus() bool {
	if o != nil && !IsNil(o.RevocationStatus) {
		return true
	}

	return false
}

// SetRevocationStatus gets a reference to the given RequestRevocationStatus and assigns it to the RevocationStatus field.
func (o *RequestFullCommon) SetRevocationStatus(v RequestRevocationStatus) {
	o.RevocationStatus = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestFullCommon) GetRevoked() time.Time {
	if o == nil || IsNil(o.Revoked.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Revoked.Get()
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullCommon) GetRevokedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revoked.Get(), o.Revoked.IsSet()
}

// HasRevoked returns a boolean if a field has been set.
func (o *RequestFullCommon) HasRevoked() bool {
	if o != nil && o.Revoked.IsSet() {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given NullableTime and assigns it to the Revoked field.
func (o *RequestFullCommon) SetRevoked(v time.Time) {
	o.Revoked.Set(&v)
}

// SetRevokedNil sets the value for Revoked to be an explicit nil
func (o *RequestFullCommon) SetRevokedNil() {
	o.Revoked.Set(nil)
}

// UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
func (o *RequestFullCommon) UnsetRevoked() {
	o.Revoked.Unset()
}

// GetRequestedBy returns the RequestedBy field value
func (o *RequestFullCommon) GetRequestedBy() ClientCredentialPrincipal {
	if o == nil {
		var ret ClientCredentialPrincipal
		return ret
	}

	return o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRequestedByOk() (*ClientCredentialPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedBy, true
}

// SetRequestedBy sets field value
func (o *RequestFullCommon) SetRequestedBy(v ClientCredentialPrincipal) {
	o.RequestedBy = v
}

// GetRequestedFor returns the RequestedFor field value
func (o *RequestFullCommon) GetRequestedFor() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRequestedForOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *RequestFullCommon) SetRequestedFor(v TargetPrincipal) {
	o.RequestedFor = v
}

// GetRequested returns the Requested field value
func (o *RequestFullCommon) GetRequested() Requested {
	if o == nil {
		var ret Requested
		return ret
	}

	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRequestedOk() (*Requested, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value
func (o *RequestFullCommon) SetRequested(v Requested) {
	o.Requested = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestFullCommon) GetAccessDuration() string {
	if o == nil || IsNil(o.AccessDuration.Get()) {
		var ret string
		return ret
	}
	return *o.AccessDuration.Get()
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullCommon) GetAccessDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessDuration.Get(), o.AccessDuration.IsSet()
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *RequestFullCommon) HasAccessDuration() bool {
	if o != nil && o.AccessDuration.IsSet() {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given NullableString and assigns it to the AccessDuration field.
func (o *RequestFullCommon) SetAccessDuration(v string) {
	o.AccessDuration.Set(&v)
}

// SetAccessDurationNil sets the value for AccessDuration to be an explicit nil
func (o *RequestFullCommon) SetAccessDurationNil() {
	o.AccessDuration.Set(nil)
}

// UnsetAccessDuration ensures that no value is present for AccessDuration, not even an explicit nil
func (o *RequestFullCommon) UnsetAccessDuration() {
	o.AccessDuration.Unset()
}

// GetRevocationScheduled returns the RevocationScheduled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestFullCommon) GetRevocationScheduled() time.Time {
	if o == nil || IsNil(o.RevocationScheduled.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RevocationScheduled.Get()
}

// GetRevocationScheduledOk returns a tuple with the RevocationScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestFullCommon) GetRevocationScheduledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationScheduled.Get(), o.RevocationScheduled.IsSet()
}

// HasRevocationScheduled returns a boolean if a field has been set.
func (o *RequestFullCommon) HasRevocationScheduled() bool {
	if o != nil && o.RevocationScheduled.IsSet() {
		return true
	}

	return false
}

// SetRevocationScheduled gets a reference to the given NullableTime and assigns it to the RevocationScheduled field.
func (o *RequestFullCommon) SetRevocationScheduled(v time.Time) {
	o.RevocationScheduled.Set(&v)
}

// SetRevocationScheduledNil sets the value for RevocationScheduled to be an explicit nil
func (o *RequestFullCommon) SetRevocationScheduledNil() {
	o.RevocationScheduled.Set(nil)
}

// UnsetRevocationScheduled ensures that no value is present for RevocationScheduled, not even an explicit nil
func (o *RequestFullCommon) UnsetRevocationScheduled() {
	o.RevocationScheduled.Unset()
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise.
func (o *RequestFullCommon) GetRequesterFieldValues() []RequestFieldValue {
	if o == nil || IsNil(o.RequesterFieldValues) {
		var ret []RequestFieldValue
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRequesterFieldValuesOk() ([]RequestFieldValue, bool) {
	if o == nil || IsNil(o.RequesterFieldValues) {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *RequestFullCommon) HasRequesterFieldValues() bool {
	if o != nil && !IsNil(o.RequesterFieldValues) {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []RequestFieldValue and assigns it to the RequesterFieldValues field.
func (o *RequestFullCommon) SetRequesterFieldValues(v []RequestFieldValue) {
	o.RequesterFieldValues = v
}

// GetRequestApproval returns the RequestApproval field value if set, zero value otherwise.
func (o *RequestFullCommon) GetRequestApproval() RequestApproval2 {
	if o == nil || IsNil(o.RequestApproval) {
		var ret RequestApproval2
		return ret
	}
	return *o.RequestApproval
}

// GetRequestApprovalOk returns a tuple with the RequestApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRequestApprovalOk() (*RequestApproval2, bool) {
	if o == nil || IsNil(o.RequestApproval) {
		return nil, false
	}
	return o.RequestApproval, true
}

// HasRequestApproval returns a boolean if a field has been set.
func (o *RequestFullCommon) HasRequestApproval() bool {
	if o != nil && !IsNil(o.RequestApproval) {
		return true
	}

	return false
}

// SetRequestApproval gets a reference to the given RequestApproval2 and assigns it to the RequestApproval field.
func (o *RequestFullCommon) SetRequestApproval(v RequestApproval2) {
	o.RequestApproval = &v
}

// GetRiskAssessment returns the RiskAssessment field value if set, zero value otherwise.
func (o *RequestFullCommon) GetRiskAssessment() RiskAssessment {
	if o == nil || IsNil(o.RiskAssessment) {
		var ret RiskAssessment
		return ret
	}
	return *o.RiskAssessment
}

// GetRiskAssessmentOk returns a tuple with the RiskAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestFullCommon) GetRiskAssessmentOk() (*RiskAssessment, bool) {
	if o == nil || IsNil(o.RiskAssessment) {
		return nil, false
	}
	return o.RiskAssessment, true
}

// HasRiskAssessment returns a boolean if a field has been set.
func (o *RequestFullCommon) HasRiskAssessment() bool {
	if o != nil && !IsNil(o.RiskAssessment) {
		return true
	}

	return false
}

// SetRiskAssessment gets a reference to the given RiskAssessment and assigns it to the RiskAssessment field.
func (o *RequestFullCommon) SetRiskAssessment(v RiskAssessment) {
	o.RiskAssessment = &v
}

func (o RequestFullCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestFullCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["created"] = o.Created
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	toSerialize["_links"] = o.Links
	toSerialize["status"] = o.Status
	if o.Resolved.IsSet() {
		toSerialize["resolved"] = o.Resolved.Get()
	}
	if !IsNil(o.GrantStatus) {
		toSerialize["grantStatus"] = o.GrantStatus
	}
	if o.Granted.IsSet() {
		toSerialize["granted"] = o.Granted.Get()
	}
	if !IsNil(o.RevocationStatus) {
		toSerialize["revocationStatus"] = o.RevocationStatus
	}
	if o.Revoked.IsSet() {
		toSerialize["revoked"] = o.Revoked.Get()
	}
	toSerialize["requestedBy"] = o.RequestedBy
	toSerialize["requestedFor"] = o.RequestedFor
	toSerialize["requested"] = o.Requested
	if o.AccessDuration.IsSet() {
		toSerialize["accessDuration"] = o.AccessDuration.Get()
	}
	if o.RevocationScheduled.IsSet() {
		toSerialize["revocationScheduled"] = o.RevocationScheduled.Get()
	}
	if !IsNil(o.RequesterFieldValues) {
		toSerialize["requesterFieldValues"] = o.RequesterFieldValues
	}
	if !IsNil(o.RequestApproval) {
		toSerialize["requestApproval"] = o.RequestApproval
	}
	if !IsNil(o.RiskAssessment) {
		toSerialize["riskAssessment"] = o.RiskAssessment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestFullCommon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"created",
		"lastUpdated",
		"lastUpdatedBy",
		"_links",
		"status",
		"requestedBy",
		"requestedFor",
		"requested",
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

	varRequestFullCommon := _RequestFullCommon{}

	err = json.Unmarshal(data, &varRequestFullCommon)

	if err != nil {
		return err
	}

	*o = RequestFullCommon(varRequestFullCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "status")
		delete(additionalProperties, "resolved")
		delete(additionalProperties, "grantStatus")
		delete(additionalProperties, "granted")
		delete(additionalProperties, "revocationStatus")
		delete(additionalProperties, "revoked")
		delete(additionalProperties, "requestedBy")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requested")
		delete(additionalProperties, "accessDuration")
		delete(additionalProperties, "revocationScheduled")
		delete(additionalProperties, "requesterFieldValues")
		delete(additionalProperties, "requestApproval")
		delete(additionalProperties, "riskAssessment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestFullCommon struct {
	value *RequestFullCommon
	isSet bool
}

func (v NullableRequestFullCommon) Get() *RequestFullCommon {
	return v.value
}

func (v *NullableRequestFullCommon) Set(val *RequestFullCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFullCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFullCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFullCommon(val *RequestFullCommon) *NullableRequestFullCommon {
	return &NullableRequestFullCommon{value: val, isSet: true}
}

func (v NullableRequestFullCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFullCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
