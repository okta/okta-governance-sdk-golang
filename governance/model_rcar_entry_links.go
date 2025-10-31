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

// checks if the RcarEntryLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RcarEntryLinks{}

// RcarEntryLinks Links available for a resource catalog entry. * `createRequest`: returns the link to create the request only if the entry is requestable * `parent`: returns the parent entry only if a parent exists * `submittedRequest`: returns the submitted request only if the authenticated user has a submitted request for the entry and it's not available for approval * `pendingRequest`: returns the pending request only if the authenticated user has a submitted request for the entry and the request is available for approval * `relatedEntity`: returns the resource that's requested
type RcarEntryLinks struct {
	Self Link `json:"self"`
	// Links to the catalog entry logo resources
	Logo                 []Link `json:"logo,omitempty"`
	Parent               *Link  `json:"parent,omitempty"`
	CreateRequest        *Link  `json:"createRequest,omitempty"`
	SubmittedRequest     *Link  `json:"submittedRequest,omitempty"`
	PendingRequest       *Link  `json:"pendingRequest,omitempty"`
	RelatedEntity        *Link  `json:"relatedEntity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RcarEntryLinks RcarEntryLinks

// NewRcarEntryLinks instantiates a new RcarEntryLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcarEntryLinks(self Link) *RcarEntryLinks {
	this := RcarEntryLinks{}
	this.Self = self
	return &this
}

// NewRcarEntryLinksWithDefaults instantiates a new RcarEntryLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcarEntryLinksWithDefaults() *RcarEntryLinks {
	this := RcarEntryLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RcarEntryLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RcarEntryLinks) SetSelf(v Link) {
	o.Self = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *RcarEntryLinks) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *RcarEntryLinks) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *RcarEntryLinks) SetLogo(v []Link) {
	o.Logo = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *RcarEntryLinks) GetParent() Link {
	if o == nil || IsNil(o.Parent) {
		var ret Link
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetParentOk() (*Link, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *RcarEntryLinks) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given Link and assigns it to the Parent field.
func (o *RcarEntryLinks) SetParent(v Link) {
	o.Parent = &v
}

// GetCreateRequest returns the CreateRequest field value if set, zero value otherwise.
func (o *RcarEntryLinks) GetCreateRequest() Link {
	if o == nil || IsNil(o.CreateRequest) {
		var ret Link
		return ret
	}
	return *o.CreateRequest
}

// GetCreateRequestOk returns a tuple with the CreateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetCreateRequestOk() (*Link, bool) {
	if o == nil || IsNil(o.CreateRequest) {
		return nil, false
	}
	return o.CreateRequest, true
}

// HasCreateRequest returns a boolean if a field has been set.
func (o *RcarEntryLinks) HasCreateRequest() bool {
	if o != nil && !IsNil(o.CreateRequest) {
		return true
	}

	return false
}

// SetCreateRequest gets a reference to the given Link and assigns it to the CreateRequest field.
func (o *RcarEntryLinks) SetCreateRequest(v Link) {
	o.CreateRequest = &v
}

// GetSubmittedRequest returns the SubmittedRequest field value if set, zero value otherwise.
func (o *RcarEntryLinks) GetSubmittedRequest() Link {
	if o == nil || IsNil(o.SubmittedRequest) {
		var ret Link
		return ret
	}
	return *o.SubmittedRequest
}

// GetSubmittedRequestOk returns a tuple with the SubmittedRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetSubmittedRequestOk() (*Link, bool) {
	if o == nil || IsNil(o.SubmittedRequest) {
		return nil, false
	}
	return o.SubmittedRequest, true
}

// HasSubmittedRequest returns a boolean if a field has been set.
func (o *RcarEntryLinks) HasSubmittedRequest() bool {
	if o != nil && !IsNil(o.SubmittedRequest) {
		return true
	}

	return false
}

// SetSubmittedRequest gets a reference to the given Link and assigns it to the SubmittedRequest field.
func (o *RcarEntryLinks) SetSubmittedRequest(v Link) {
	o.SubmittedRequest = &v
}

// GetPendingRequest returns the PendingRequest field value if set, zero value otherwise.
func (o *RcarEntryLinks) GetPendingRequest() Link {
	if o == nil || IsNil(o.PendingRequest) {
		var ret Link
		return ret
	}
	return *o.PendingRequest
}

// GetPendingRequestOk returns a tuple with the PendingRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetPendingRequestOk() (*Link, bool) {
	if o == nil || IsNil(o.PendingRequest) {
		return nil, false
	}
	return o.PendingRequest, true
}

// HasPendingRequest returns a boolean if a field has been set.
func (o *RcarEntryLinks) HasPendingRequest() bool {
	if o != nil && !IsNil(o.PendingRequest) {
		return true
	}

	return false
}

// SetPendingRequest gets a reference to the given Link and assigns it to the PendingRequest field.
func (o *RcarEntryLinks) SetPendingRequest(v Link) {
	o.PendingRequest = &v
}

// GetRelatedEntity returns the RelatedEntity field value if set, zero value otherwise.
func (o *RcarEntryLinks) GetRelatedEntity() Link {
	if o == nil || IsNil(o.RelatedEntity) {
		var ret Link
		return ret
	}
	return *o.RelatedEntity
}

// GetRelatedEntityOk returns a tuple with the RelatedEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntryLinks) GetRelatedEntityOk() (*Link, bool) {
	if o == nil || IsNil(o.RelatedEntity) {
		return nil, false
	}
	return o.RelatedEntity, true
}

// HasRelatedEntity returns a boolean if a field has been set.
func (o *RcarEntryLinks) HasRelatedEntity() bool {
	if o != nil && !IsNil(o.RelatedEntity) {
		return true
	}

	return false
}

// SetRelatedEntity gets a reference to the given Link and assigns it to the RelatedEntity field.
func (o *RcarEntryLinks) SetRelatedEntity(v Link) {
	o.RelatedEntity = &v
}

func (o RcarEntryLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RcarEntryLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.CreateRequest) {
		toSerialize["createRequest"] = o.CreateRequest
	}
	if !IsNil(o.SubmittedRequest) {
		toSerialize["submittedRequest"] = o.SubmittedRequest
	}
	if !IsNil(o.PendingRequest) {
		toSerialize["pendingRequest"] = o.PendingRequest
	}
	if !IsNil(o.RelatedEntity) {
		toSerialize["relatedEntity"] = o.RelatedEntity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RcarEntryLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
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

	varRcarEntryLinks := _RcarEntryLinks{}

	err = json.Unmarshal(data, &varRcarEntryLinks)

	if err != nil {
		return err
	}

	*o = RcarEntryLinks(varRcarEntryLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "createRequest")
		delete(additionalProperties, "submittedRequest")
		delete(additionalProperties, "pendingRequest")
		delete(additionalProperties, "relatedEntity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRcarEntryLinks struct {
	value *RcarEntryLinks
	isSet bool
}

func (v NullableRcarEntryLinks) Get() *RcarEntryLinks {
	return v.value
}

func (v *NullableRcarEntryLinks) Set(val *RcarEntryLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRcarEntryLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRcarEntryLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcarEntryLinks(val *RcarEntryLinks) *NullableRcarEntryLinks {
	return &NullableRcarEntryLinks{value: val, isSet: true}
}

func (v NullableRcarEntryLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcarEntryLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
