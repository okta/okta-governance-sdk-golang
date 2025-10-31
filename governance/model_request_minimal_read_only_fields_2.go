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
	"time"
)

// checks if the RequestMinimalReadOnlyFields2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestMinimalReadOnlyFields2{}

// RequestMinimalReadOnlyFields2 Properties contained by default in sparse representation
type RequestMinimalReadOnlyFields2 struct {
	Status *RequestStatus `json:"status,omitempty"`
	// The date the request was resolved. The property may transition from having a value to null if the request is reopened.
	Resolved    NullableTime        `json:"resolved,omitempty"`
	GrantStatus *RequestGrantStatus `json:"grantStatus,omitempty"`
	// The date the approved access was granted. Only set if request.status is APPROVED.
	Granted          NullableTime             `json:"granted,omitempty"`
	RevocationStatus *RequestRevocationStatus `json:"revocationStatus,omitempty"`
	// The date the granted access was revoked. Only set if request.grantStatus is GRANTED and request.revocationStatus is REVOKED.
	Revoked              NullableTime     `json:"revoked,omitempty"`
	RequestedBy          *TargetPrincipal `json:"requestedBy,omitempty"`
	RequestedFor         *TargetPrincipal `json:"requestedFor,omitempty"`
	Requested            *Requested       `json:"requested,omitempty"`
	Links                *RequestLinks2   `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestMinimalReadOnlyFields2 RequestMinimalReadOnlyFields2

// NewRequestMinimalReadOnlyFields2 instantiates a new RequestMinimalReadOnlyFields2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestMinimalReadOnlyFields2() *RequestMinimalReadOnlyFields2 {
	this := RequestMinimalReadOnlyFields2{}
	return &this
}

// NewRequestMinimalReadOnlyFields2WithDefaults instantiates a new RequestMinimalReadOnlyFields2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestMinimalReadOnlyFields2WithDefaults() *RequestMinimalReadOnlyFields2 {
	this := RequestMinimalReadOnlyFields2{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetStatus() RequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret RequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetStatusOk() (*RequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RequestStatus and assigns it to the Status field.
func (o *RequestMinimalReadOnlyFields2) SetStatus(v RequestStatus) {
	o.Status = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestMinimalReadOnlyFields2) GetResolved() time.Time {
	if o == nil || IsNil(o.Resolved.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Resolved.Get()
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMinimalReadOnlyFields2) GetResolvedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resolved.Get(), o.Resolved.IsSet()
}

// HasResolved returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasResolved() bool {
	if o != nil && o.Resolved.IsSet() {
		return true
	}

	return false
}

// SetResolved gets a reference to the given NullableTime and assigns it to the Resolved field.
func (o *RequestMinimalReadOnlyFields2) SetResolved(v time.Time) {
	o.Resolved.Set(&v)
}

// SetResolvedNil sets the value for Resolved to be an explicit nil
func (o *RequestMinimalReadOnlyFields2) SetResolvedNil() {
	o.Resolved.Set(nil)
}

// UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
func (o *RequestMinimalReadOnlyFields2) UnsetResolved() {
	o.Resolved.Unset()
}

// GetGrantStatus returns the GrantStatus field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetGrantStatus() RequestGrantStatus {
	if o == nil || IsNil(o.GrantStatus) {
		var ret RequestGrantStatus
		return ret
	}
	return *o.GrantStatus
}

// GetGrantStatusOk returns a tuple with the GrantStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetGrantStatusOk() (*RequestGrantStatus, bool) {
	if o == nil || IsNil(o.GrantStatus) {
		return nil, false
	}
	return o.GrantStatus, true
}

// HasGrantStatus returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasGrantStatus() bool {
	if o != nil && !IsNil(o.GrantStatus) {
		return true
	}

	return false
}

// SetGrantStatus gets a reference to the given RequestGrantStatus and assigns it to the GrantStatus field.
func (o *RequestMinimalReadOnlyFields2) SetGrantStatus(v RequestGrantStatus) {
	o.GrantStatus = &v
}

// GetGranted returns the Granted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestMinimalReadOnlyFields2) GetGranted() time.Time {
	if o == nil || IsNil(o.Granted.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Granted.Get()
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMinimalReadOnlyFields2) GetGrantedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Granted.Get(), o.Granted.IsSet()
}

// HasGranted returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasGranted() bool {
	if o != nil && o.Granted.IsSet() {
		return true
	}

	return false
}

// SetGranted gets a reference to the given NullableTime and assigns it to the Granted field.
func (o *RequestMinimalReadOnlyFields2) SetGranted(v time.Time) {
	o.Granted.Set(&v)
}

// SetGrantedNil sets the value for Granted to be an explicit nil
func (o *RequestMinimalReadOnlyFields2) SetGrantedNil() {
	o.Granted.Set(nil)
}

// UnsetGranted ensures that no value is present for Granted, not even an explicit nil
func (o *RequestMinimalReadOnlyFields2) UnsetGranted() {
	o.Granted.Unset()
}

// GetRevocationStatus returns the RevocationStatus field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetRevocationStatus() RequestRevocationStatus {
	if o == nil || IsNil(o.RevocationStatus) {
		var ret RequestRevocationStatus
		return ret
	}
	return *o.RevocationStatus
}

// GetRevocationStatusOk returns a tuple with the RevocationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetRevocationStatusOk() (*RequestRevocationStatus, bool) {
	if o == nil || IsNil(o.RevocationStatus) {
		return nil, false
	}
	return o.RevocationStatus, true
}

// HasRevocationStatus returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasRevocationStatus() bool {
	if o != nil && !IsNil(o.RevocationStatus) {
		return true
	}

	return false
}

// SetRevocationStatus gets a reference to the given RequestRevocationStatus and assigns it to the RevocationStatus field.
func (o *RequestMinimalReadOnlyFields2) SetRevocationStatus(v RequestRevocationStatus) {
	o.RevocationStatus = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestMinimalReadOnlyFields2) GetRevoked() time.Time {
	if o == nil || IsNil(o.Revoked.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Revoked.Get()
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestMinimalReadOnlyFields2) GetRevokedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revoked.Get(), o.Revoked.IsSet()
}

// HasRevoked returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasRevoked() bool {
	if o != nil && o.Revoked.IsSet() {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given NullableTime and assigns it to the Revoked field.
func (o *RequestMinimalReadOnlyFields2) SetRevoked(v time.Time) {
	o.Revoked.Set(&v)
}

// SetRevokedNil sets the value for Revoked to be an explicit nil
func (o *RequestMinimalReadOnlyFields2) SetRevokedNil() {
	o.Revoked.Set(nil)
}

// UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
func (o *RequestMinimalReadOnlyFields2) UnsetRevoked() {
	o.Revoked.Unset()
}

// GetRequestedBy returns the RequestedBy field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetRequestedBy() TargetPrincipal {
	if o == nil || IsNil(o.RequestedBy) {
		var ret TargetPrincipal
		return ret
	}
	return *o.RequestedBy
}

// GetRequestedByOk returns a tuple with the RequestedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetRequestedByOk() (*TargetPrincipal, bool) {
	if o == nil || IsNil(o.RequestedBy) {
		return nil, false
	}
	return o.RequestedBy, true
}

// HasRequestedBy returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasRequestedBy() bool {
	if o != nil && !IsNil(o.RequestedBy) {
		return true
	}

	return false
}

// SetRequestedBy gets a reference to the given TargetPrincipal and assigns it to the RequestedBy field.
func (o *RequestMinimalReadOnlyFields2) SetRequestedBy(v TargetPrincipal) {
	o.RequestedBy = &v
}

// GetRequestedFor returns the RequestedFor field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetRequestedFor() TargetPrincipal {
	if o == nil || IsNil(o.RequestedFor) {
		var ret TargetPrincipal
		return ret
	}
	return *o.RequestedFor
}

// GetRequestedForOk returns a tuple with the RequestedFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetRequestedForOk() (*TargetPrincipal, bool) {
	if o == nil || IsNil(o.RequestedFor) {
		return nil, false
	}
	return o.RequestedFor, true
}

// HasRequestedFor returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasRequestedFor() bool {
	if o != nil && !IsNil(o.RequestedFor) {
		return true
	}

	return false
}

// SetRequestedFor gets a reference to the given TargetPrincipal and assigns it to the RequestedFor field.
func (o *RequestMinimalReadOnlyFields2) SetRequestedFor(v TargetPrincipal) {
	o.RequestedFor = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetRequested() Requested {
	if o == nil || IsNil(o.Requested) {
		var ret Requested
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetRequestedOk() (*Requested, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasRequested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given Requested and assigns it to the Requested field.
func (o *RequestMinimalReadOnlyFields2) SetRequested(v Requested) {
	o.Requested = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RequestMinimalReadOnlyFields2) GetLinks() RequestLinks2 {
	if o == nil || IsNil(o.Links) {
		var ret RequestLinks2
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestMinimalReadOnlyFields2) GetLinksOk() (*RequestLinks2, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RequestMinimalReadOnlyFields2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RequestLinks2 and assigns it to the Links field.
func (o *RequestMinimalReadOnlyFields2) SetLinks(v RequestLinks2) {
	o.Links = &v
}

func (o RequestMinimalReadOnlyFields2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestMinimalReadOnlyFields2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
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
	if !IsNil(o.RequestedBy) {
		toSerialize["requestedBy"] = o.RequestedBy
	}
	if !IsNil(o.RequestedFor) {
		toSerialize["requestedFor"] = o.RequestedFor
	}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestMinimalReadOnlyFields2) UnmarshalJSON(data []byte) (err error) {
	varRequestMinimalReadOnlyFields2 := _RequestMinimalReadOnlyFields2{}

	err = json.Unmarshal(data, &varRequestMinimalReadOnlyFields2)

	if err != nil {
		return err
	}

	*o = RequestMinimalReadOnlyFields2(varRequestMinimalReadOnlyFields2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "resolved")
		delete(additionalProperties, "grantStatus")
		delete(additionalProperties, "granted")
		delete(additionalProperties, "revocationStatus")
		delete(additionalProperties, "revoked")
		delete(additionalProperties, "requestedBy")
		delete(additionalProperties, "requestedFor")
		delete(additionalProperties, "requested")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestMinimalReadOnlyFields2 struct {
	value *RequestMinimalReadOnlyFields2
	isSet bool
}

func (v NullableRequestMinimalReadOnlyFields2) Get() *RequestMinimalReadOnlyFields2 {
	return v.value
}

func (v *NullableRequestMinimalReadOnlyFields2) Set(val *RequestMinimalReadOnlyFields2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMinimalReadOnlyFields2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMinimalReadOnlyFields2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMinimalReadOnlyFields2(val *RequestMinimalReadOnlyFields2) *NullableRequestMinimalReadOnlyFields2 {
	return &NullableRequestMinimalReadOnlyFields2{value: val, isSet: true}
}

func (v NullableRequestMinimalReadOnlyFields2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMinimalReadOnlyFields2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
