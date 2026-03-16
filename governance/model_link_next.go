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

// checks if the LinkNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkNext{}

// LinkNext A link to the next collection of resources if pagination is required
type LinkNext struct {
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
	// Indicates whether the link object's `href` property is a URI template
	Templated *bool `json:"templated,omitempty"`
	// Link hints
	Hints *map[string][]string `json:"hints,omitempty"`
	// Link title or label
	Title                *string `json:"title,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkNext LinkNext

// NewLinkNext instantiates a new LinkNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkNext(href string) *LinkNext {
	this := LinkNext{}
	this.Href = href
	return &this
}

// NewLinkNextWithDefaults instantiates a new LinkNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkNextWithDefaults() *LinkNext {
	this := LinkNext{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinkNext) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNext) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinkNext) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinkNext) SetName(v string) {
	o.Name = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *LinkNext) GetRel() string {
	if o == nil || IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNext) GetRelOk() (*string, bool) {
	if o == nil || IsNil(o.Rel) {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *LinkNext) HasRel() bool {
	if o != nil && !IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *LinkNext) SetRel(v string) {
	o.Rel = &v
}

// GetHref returns the Href field value
func (o *LinkNext) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LinkNext) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LinkNext) SetHref(v string) {
	o.Href = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkNext) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNext) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkNext) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LinkNext) SetType(v string) {
	o.Type = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
// Deprecated
func (o *LinkNext) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LinkNext) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LinkNext) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
// Deprecated
func (o *LinkNext) SetMethod(v string) {
	o.Method = &v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *LinkNext) GetTemplated() bool {
	if o == nil || IsNil(o.Templated) {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNext) GetTemplatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Templated) {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *LinkNext) HasTemplated() bool {
	if o != nil && !IsNil(o.Templated) {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *LinkNext) SetTemplated(v bool) {
	o.Templated = &v
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *LinkNext) GetHints() map[string][]string {
	if o == nil || IsNil(o.Hints) {
		var ret map[string][]string
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNext) GetHintsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Hints) {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *LinkNext) HasHints() bool {
	if o != nil && !IsNil(o.Hints) {
		return true
	}

	return false
}

// SetHints gets a reference to the given map[string][]string and assigns it to the Hints field.
func (o *LinkNext) SetHints(v map[string][]string) {
	o.Hints = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LinkNext) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkNext) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LinkNext) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LinkNext) SetTitle(v string) {
	o.Title = &v
}

func (o LinkNext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	toSerialize["href"] = o.Href
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Templated) {
		toSerialize["templated"] = o.Templated
	}
	if !IsNil(o.Hints) {
		toSerialize["hints"] = o.Hints
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkNext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
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

	varLinkNext := _LinkNext{}

	err = json.Unmarshal(data, &varLinkNext)

	if err != nil {
		return err
	}

	*o = LinkNext(varLinkNext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "rel")
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "method")
		delete(additionalProperties, "templated")
		delete(additionalProperties, "hints")
		delete(additionalProperties, "title")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkNext struct {
	value *LinkNext
	isSet bool
}

func (v NullableLinkNext) Get() *LinkNext {
	return v.value
}

func (v *NullableLinkNext) Set(val *LinkNext) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkNext) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkNext(val *LinkNext) *NullableLinkNext {
	return &NullableLinkNext{value: val, isSet: true}
}

func (v NullableLinkNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
