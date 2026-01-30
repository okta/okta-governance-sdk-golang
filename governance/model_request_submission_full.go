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

// checks if the RequestSubmissionFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestSubmissionFull{}

// RequestSubmissionFull Full representation of a request submission
type RequestSubmissionFull struct {
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
	// The status of the request
	Status string `json:"status"`
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
	// How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported, for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  > **Note:** Mixes of units, as well as month/year/week designations, are not supported. For example, `P40DT65H`, `P40M`, `P1W` and `P1Y` are not supported.
	AccessDuration NullableString `json:"accessDuration,omitempty"`
	// The date the granted access is scheduled for recovation. Only set if request.accessDuration exists, and request.grantStatus is GRANTED.
	RevocationScheduled NullableTime `json:"revocationScheduled,omitempty"`
	// The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.
	RequesterFieldValues []RequestFieldValue `json:"requesterFieldValues,omitempty"`
	RiskAssessment       *RiskAssessment     `json:"riskAssessment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestSubmissionFull RequestSubmissionFull

// NewRequestSubmissionFull instantiates a new RequestSubmissionFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSubmissionFull(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks2, status string, requestedBy ClientCredentialPrincipal, requestedFor TargetPrincipal, requested Requested) *RequestSubmissionFull {
	this := RequestSubmissionFull{}
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

// NewRequestSubmissionFullWithDefaults instantiates a new RequestSubmissionFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSubmissionFullWithDefaults() *RequestSubmissionFull {
	this := RequestSubmissionFull{}
	return &this
}

// GetId returns the Id field value
func (o *RequestSubmissionFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestSubmissionFull) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *RequestSubmissionFull) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *RequestSubmissionFull) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreated returns the Created field value
func (o *RequestSubmissionFull) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *RequestSubmissionFull) SetCreated(v time.Time) {
	o.Created = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *RequestSubmissionFull) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RequestSubmissionFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value
func (o *RequestSubmissionFull) GetLastUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedBy, true
}

// SetLastUpdatedBy sets field value
func (o *RequestSubmissionFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = v
}

// GetLinks returns the Links field value
func (o *RequestSubmissionFull) GetLinks() RequestLinks2 {
	if o == nil {
		var ret RequestLinks2
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetLinksOk() (*RequestLinks2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RequestSubmissionFull) SetLinks(v RequestLinks2) {
	o.Links = v
}

// GetStatus returns the Status field value
func (o *RequestSubmissionFull) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RequestSubmissionFull) SetStatus(v string) {
	o.Status = v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSubmissionFull) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSubmissionFull) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableTime and assigns it to the Resolved field.
func (o *RequestSubmissionFull) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *RequestSubmissionFull) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *RequestSubmissionFull) UnsetResolved() {
	o.Resolved.Unset()
}

// GetGrantStatus returns the GrantStatus field value if set, zero value otherwise.
func (o *RequestSubmissionFull) GetGrantStatus() RequestGrantStatus {
	if o == nil || IsNil(o.GrantStatus) {
		var ret RequestGrantStatus
		return ret
	}
	return *o.GrantStatus
}

// GetGrantStatusOk returns a tuple with the GrantStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetGrantStatusOk() (*RequestGrantStatus, bool) {
	if o == nil || IsNil(o.GrantStatus) {
		return nil, false
	}
	return o.GrantStatus, true
}

// HasGrantStatus returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasGrantStatus() bool {
	if o != nil && !IsNil(o.GrantStatus) {
		return true
	}

	return false
}

// SetGrantStatus gets a reference to the given RequestGrantStatus and assigns it to the GrantStatus field.
func (o *RequestSubmissionFull) SetGrantStatus(v RequestGrantStatus) {
	o.GrantStatus = &v
}

// GetGranted returns the Granted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSubmissionFull) GetGranted() time.Time {
	if o == nil || IsNil(o.Granted.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Granted.Get()
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSubmissionFull) GetGrantedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Granted.Get(), o.Granted.IsSet()
}

// HasGranted returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasGranted() bool {
	if o != nil && o.Granted.IsSet() {
		return true
	}

	return false
}

// SetGranted gets a reference to the given NullableTime and assigns it to the Granted field.
func (o *RequestSubmissionFull) SetGranted(v time.Time) {
	o.Granted.Set(&v)
}

// SetGrantedNil sets the value for Granted to be an explicit nil
func (o *RequestSubmissionFull) SetGrantedNil() {
	o.Granted.Set(nil)
}

// UnsetGranted ensures that no value is present for Granted, not even an explicit nil
func (o *RequestSubmissionFull) UnsetGranted() {
	o.Granted.Unset()
}

// GetRevocationStatus returns the RevocationStatus field value if set, zero value otherwise.
func (o *RequestSubmissionFull) GetRevocationStatus() RequestRevocationStatus {
	if o == nil || IsNil(o.RevocationStatus) {
		var ret RequestRevocationStatus
		return ret
	}
	return *o.RevocationStatus
}

// GetRevocationStatusOk returns a tuple with the RevocationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetRevocationStatusOk() (*RequestRevocationStatus, bool) {
	if o == nil || IsNil(o.RevocationStatus) {
		return nil, false
	}
	return o.RevocationStatus, true
}

// HasRevocationStatus returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasRevocationStatus() bool {
	if o != nil && !IsNil(o.RevocationStatus) {
		return true
	}

	return false
}

// SetRevocationStatus gets a reference to the given RequestRevocationStatus and assigns it to the RevocationStatus field.
func (o *RequestSubmissionFull) SetRevocationStatus(v RequestRevocationStatus) {
	o.RevocationStatus = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSubmissionFull) GetRevoked() time.Time {
	if o == nil || IsNil(o.Revoked.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Revoked.Get()
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSubmissionFull) GetRevokedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revoked.Get(), o.Revoked.IsSet()
}

// HasRevoked returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasRevoked() bool {
	if o != nil && o.Revoked.IsSet() {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given NullableTime and assigns it to the Revoked field.
func (o *RequestSubmissionFull) SetRevoked(v time.Time) {
	o.Revoked.Set(&v)
}

// SetRevokedNil sets the value for Revoked to be an explicit nil
func (o *RequestSubmissionFull) SetRevokedNil() {
	o.Revoked.Set(nil)
}

// UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
func (o *RequestSubmissionFull) UnsetRevoked() {
	o.Revoked.Unset()
}

// GetRequestedBy returns the RequestedBy field value
func (o *RequestSubmissionFull) GetRequestedBy() ClientCredentialPrincipal {
	if o == nil {
		var ret ClientCredentialPrincipal
		return ret
	}

	return o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetRequestedByOk() (*ClientCredentialPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedBy, true
}

// SetRequestedBy sets field value
func (o *RequestSubmissionFull) SetRequestedBy(v ClientCredentialPrincipal) {
	o.RequestedBy = v
}

// GetRequestedFor returns the RequestedFor field value
func (o *RequestSubmissionFull) GetRequestedFor() TargetPrincipal {
	if o == nil {
		var ret TargetPrincipal
		return ret
	}

	return o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetRequestedForOk() (*TargetPrincipal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedFor, true
}

// SetRequestedFor sets field value
func (o *RequestSubmissionFull) SetRequestedFor(v TargetPrincipal) {
	o.RequestedFor = v
}

// GetRequested returns the Requested field value
func (o *RequestSubmissionFull) GetRequested() Requested {
	if o == nil {
		var ret Requested
		return ret
	}

	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetRequestedOk() (*Requested, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value
func (o *RequestSubmissionFull) SetRequested(v Requested) {
	o.Requested = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSubmissionFull) GetAccessDuration() string {
	if o == nil || IsNil(o.AccessDuration.Get()) {
		var ret string
		return ret
	}
	return *o.AccessDuration.Get()
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSubmissionFull) GetAccessDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessDuration.Get(), o.AccessDuration.IsSet()
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasAccessDuration() bool {
	if o != nil && o.AccessDuration.IsSet() {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given NullableString and assigns it to the AccessDuration field.
func (o *RequestSubmissionFull) SetAccessDuration(v string) {
	o.AccessDuration.Set(&v)
}

// SetAccessDurationNil sets the value for AccessDuration to be an explicit nil
func (o *RequestSubmissionFull) SetAccessDurationNil() {
	o.AccessDuration.Set(nil)
}

// UnsetAccessDuration ensures that no value is present for AccessDuration, not even an explicit nil
func (o *RequestSubmissionFull) UnsetAccessDuration() {
	o.AccessDuration.Unset()
}

// GetRevocationScheduled returns the RevocationScheduled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSubmissionFull) GetRevocationScheduled() time.Time {
	if o == nil || IsNil(o.RevocationScheduled.Get()) {
		var ret time.Time
		return ret
	}
	return *o.RevocationScheduled.Get()
}

// GetRevocationScheduledOk returns a tuple with the RevocationScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSubmissionFull) GetRevocationScheduledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationScheduled.Get(), o.RevocationScheduled.IsSet()
}

// HasRevocationScheduled returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasRevocationScheduled() bool {
	if o != nil && o.RevocationScheduled.IsSet() {
		return true
	}

	return false
}

// SetRevocationScheduled gets a reference to the given NullableTime and assigns it to the RevocationScheduled field.
func (o *RequestSubmissionFull) SetRevocationScheduled(v time.Time) {
	o.RevocationScheduled.Set(&v)
}

// SetRevocationScheduledNil sets the value for RevocationScheduled to be an explicit nil
func (o *RequestSubmissionFull) SetRevocationScheduledNil() {
	o.RevocationScheduled.Set(nil)
}

// UnsetRevocationScheduled ensures that no value is present for RevocationScheduled, not even an explicit nil
func (o *RequestSubmissionFull) UnsetRevocationScheduled() {
	o.RevocationScheduled.Unset()
}

// GetRequesterFieldValues returns the RequesterFieldValues field value if set, zero value otherwise.
func (o *RequestSubmissionFull) GetRequesterFieldValues() []RequestFieldValue {
	if o == nil || IsNil(o.RequesterFieldValues) {
		var ret []RequestFieldValue
		return ret
	}
	return o.RequesterFieldValues
}

// GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetRequesterFieldValuesOk() ([]RequestFieldValue, bool) {
	if o == nil || IsNil(o.RequesterFieldValues) {
		return nil, false
	}
	return o.RequesterFieldValues, true
}

// HasRequesterFieldValues returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasRequesterFieldValues() bool {
	if o != nil && !IsNil(o.RequesterFieldValues) {
		return true
	}

	return false
}

// SetRequesterFieldValues gets a reference to the given []RequestFieldValue and assigns it to the RequesterFieldValues field.
func (o *RequestSubmissionFull) SetRequesterFieldValues(v []RequestFieldValue) {
	o.RequesterFieldValues = v
}

// GetRiskAssessment returns the RiskAssessment field value if set, zero value otherwise.
func (o *RequestSubmissionFull) GetRiskAssessment() RiskAssessment {
	if o == nil || IsNil(o.RiskAssessment) {
		var ret RiskAssessment
		return ret
	}
	return *o.RiskAssessment
}

// GetRiskAssessmentOk returns a tuple with the RiskAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSubmissionFull) GetRiskAssessmentOk() (*RiskAssessment, bool) {
	if o == nil || IsNil(o.RiskAssessment) {
		return nil, false
	}
	return o.RiskAssessment, true
}

// HasRiskAssessment returns a boolean if a field has been set.
func (o *RequestSubmissionFull) HasRiskAssessment() bool {
	if o != nil && !IsNil(o.RiskAssessment) {
		return true
	}

	return false
}

// SetRiskAssessment gets a reference to the given RiskAssessment and assigns it to the RiskAssessment field.
func (o *RequestSubmissionFull) SetRiskAssessment(v RiskAssessment) {
	o.RiskAssessment = &v
}

func (o RequestSubmissionFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSubmissionFull) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RiskAssessment) {
		toSerialize["riskAssessment"] = o.RiskAssessment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestSubmissionFull) UnmarshalJSON(data []byte) (err error) {
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

	varRequestSubmissionFull := _RequestSubmissionFull{}

	err = json.Unmarshal(data, &varRequestSubmissionFull)

	if err != nil {
		return err
	}

	*o = RequestSubmissionFull(varRequestSubmissionFull)

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
		delete(additionalProperties, "riskAssessment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestSubmissionFull struct {
	value *RequestSubmissionFull
	isSet bool
}

func (v NullableRequestSubmissionFull) Get() *RequestSubmissionFull {
	return v.value
}

func (v *NullableRequestSubmissionFull) Set(val *RequestSubmissionFull) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSubmissionFull) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSubmissionFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSubmissionFull(val *RequestSubmissionFull) *NullableRequestSubmissionFull {
	return &NullableRequestSubmissionFull{value: val, isSet: true}
}

func (v NullableRequestSubmissionFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSubmissionFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
