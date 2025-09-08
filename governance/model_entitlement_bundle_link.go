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

// EntitlementBundleLink Entitlement bundle details link. This link is only present when `grantType` is `ENTITLEMENT-BUNDLE`
type EntitlementBundleLink struct {
	// Link name
	Name *string `json:"name,omitempty"`
	// Link relation
	Rel *string `json:"rel,omitempty"`
	// Link URI
	Href string `json:"href"`
	// The media type of the link. If omitted, it's implicitly `application/json`.
	Type *string `json:"type,omitempty"`
	// Deprecated
	Method *string `json:"method,omitempty"`
	// Indicates whether the link object's `href` property is a URI template.
	Templated *bool `json:"templated,omitempty"`
	// Link hints
	Hints *map[string][]string `json:"hints,omitempty"`
	// Link title or label
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementBundleLink EntitlementBundleLink

// NewEntitlementBundleLink instantiates a new EntitlementBundleLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBundleLink(href string) *EntitlementBundleLink {
	this := EntitlementBundleLink{}
	this.Href = href
	return &this
}

// NewEntitlementBundleLinkWithDefaults instantiates a new EntitlementBundleLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBundleLinkWithDefaults() *EntitlementBundleLink {
	this := EntitlementBundleLink{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementBundleLink) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementBundleLink) SetName(v string) {
	o.Name = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *EntitlementBundleLink) GetRel() string {
	if o == nil || o.Rel == nil {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetRelOk() (*string, bool) {
	if o == nil || o.Rel == nil {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasRel() bool {
	if o != nil && o.Rel != nil {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *EntitlementBundleLink) SetRel(v string) {
	o.Rel = &v
}

// GetHref returns the Href field value
func (o *EntitlementBundleLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *EntitlementBundleLink) SetHref(v string) {
	o.Href = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntitlementBundleLink) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntitlementBundleLink) SetType(v string) {
	o.Type = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
// Deprecated
func (o *EntitlementBundleLink) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *EntitlementBundleLink) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
// Deprecated
func (o *EntitlementBundleLink) SetMethod(v string) {
	o.Method = &v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *EntitlementBundleLink) GetTemplated() bool {
	if o == nil || o.Templated == nil {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetTemplatedOk() (*bool, bool) {
	if o == nil || o.Templated == nil {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasTemplated() bool {
	if o != nil && o.Templated != nil {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *EntitlementBundleLink) SetTemplated(v bool) {
	o.Templated = &v
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *EntitlementBundleLink) GetHints() map[string][]string {
	if o == nil || o.Hints == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetHintsOk() (*map[string][]string, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given map[string][]string and assigns it to the Hints field.
func (o *EntitlementBundleLink) SetHints(v map[string][]string) {
	o.Hints = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *EntitlementBundleLink) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleLink) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *EntitlementBundleLink) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *EntitlementBundleLink) SetTitle(v string) {
	o.Title = &v
}

func (o EntitlementBundleLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Rel != nil {
		toSerialize["rel"] = o.Rel
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Templated != nil {
		toSerialize["templated"] = o.Templated
	}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementBundleLink) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementBundleLink := _EntitlementBundleLink{}

	err = json.Unmarshal(bytes, &varEntitlementBundleLink)
	if err == nil {
		*o = EntitlementBundleLink(varEntitlementBundleLink)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "rel")
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "method")
		delete(additionalProperties, "templated")
		delete(additionalProperties, "hints")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementBundleLink struct {
	value *EntitlementBundleLink
	isSet bool
}

func (v NullableEntitlementBundleLink) Get() *EntitlementBundleLink {
	return v.value
}

func (v *NullableEntitlementBundleLink) Set(val *EntitlementBundleLink) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundleLink) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundleLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundleLink(val *EntitlementBundleLink) *NullableEntitlementBundleLink {
	return &NullableEntitlementBundleLink{value: val, isSet: true}
}

func (v NullableEntitlementBundleLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundleLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
