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
)

// checks if the ReviewerAiAgentConnectedResourceApiServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceApiServer{}

// ReviewerAiAgentConnectedResourceApiServer API server details for an AI agent connection
type ReviewerAiAgentConnectedResourceApiServer struct {
	// ID of the connected API server
	Id *string `json:"id,omitempty"`
	// Okta Resource Name (ORN) of the connected API server
	Orn *string `json:"orn,omitempty"`
	// Name of the connected API server
	Name *string `json:"name,omitempty"`
	// URL of the connected API server
	ResourceUrl *string `json:"resourceUrl,omitempty"`
	// Connected API server scopes that the AI agent has access to
	Scopes               []string `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceApiServer ReviewerAiAgentConnectedResourceApiServer

// NewReviewerAiAgentConnectedResourceApiServer instantiates a new ReviewerAiAgentConnectedResourceApiServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceApiServer() *ReviewerAiAgentConnectedResourceApiServer {
	this := ReviewerAiAgentConnectedResourceApiServer{}
	return &this
}

// NewReviewerAiAgentConnectedResourceApiServerWithDefaults instantiates a new ReviewerAiAgentConnectedResourceApiServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceApiServerWithDefaults() *ReviewerAiAgentConnectedResourceApiServer {
	this := ReviewerAiAgentConnectedResourceApiServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReviewerAiAgentConnectedResourceApiServer) SetId(v string) {
	o.Id = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceApiServer) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceApiServer) SetName(v string) {
	o.Name = &v
}

// GetResourceUrl returns the ResourceUrl field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetResourceUrl() string {
	if o == nil || IsNil(o.ResourceUrl) {
		var ret string
		return ret
	}
	return *o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetResourceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceUrl) {
		return nil, false
	}
	return o.ResourceUrl, true
}

// HasResourceUrl returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) HasResourceUrl() bool {
	if o != nil && !IsNil(o.ResourceUrl) {
		return true
	}

	return false
}

// SetResourceUrl gets a reference to the given string and assigns it to the ResourceUrl field.
func (o *ReviewerAiAgentConnectedResourceApiServer) SetResourceUrl(v string) {
	o.ResourceUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceApiServer) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ReviewerAiAgentConnectedResourceApiServer) SetScopes(v []string) {
	o.Scopes = v
}

func (o ReviewerAiAgentConnectedResourceApiServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceApiServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ResourceUrl) {
		toSerialize["resourceUrl"] = o.ResourceUrl
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResourceApiServer) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceApiServer := _ReviewerAiAgentConnectedResourceApiServer{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceApiServer)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceApiServer(varReviewerAiAgentConnectedResourceApiServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "resourceUrl")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceApiServer struct {
	value *ReviewerAiAgentConnectedResourceApiServer
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceApiServer) Get() *ReviewerAiAgentConnectedResourceApiServer {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceApiServer) Set(val *ReviewerAiAgentConnectedResourceApiServer) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceApiServer) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceApiServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceApiServer(val *ReviewerAiAgentConnectedResourceApiServer) *NullableReviewerAiAgentConnectedResourceApiServer {
	return &NullableReviewerAiAgentConnectedResourceApiServer{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceApiServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceApiServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
