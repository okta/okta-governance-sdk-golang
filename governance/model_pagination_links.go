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

// PaginationLinks Links available in principal assignments list response
type PaginationLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginationLinks PaginationLinks

// NewPaginationLinks instantiates a new PaginationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationLinks(self Link) *PaginationLinks {
	this := PaginationLinks{}
	this.Self = self
	return &this
}

// NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationLinksWithDefaults() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *PaginationLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *PaginationLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginationLinks) GetNext() Link {
	if o == nil || o.Next == nil {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationLinks) GetNextOk() (*Link, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *PaginationLinks) SetNext(v Link) {
	o.Next = &v
}

func (o PaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaginationLinks) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationLinks := _PaginationLinks{}

	err = json.Unmarshal(bytes, &varPaginationLinks)
	if err == nil {
		*o = PaginationLinks(varPaginationLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePaginationLinks struct {
	value *PaginationLinks
	isSet bool
}

func (v NullablePaginationLinks) Get() *PaginationLinks {
	return v.value
}

func (v *NullablePaginationLinks) Set(val *PaginationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationLinks(val *PaginationLinks) *NullablePaginationLinks {
	return &NullablePaginationLinks{value: val, isSet: true}
}

func (v NullablePaginationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
