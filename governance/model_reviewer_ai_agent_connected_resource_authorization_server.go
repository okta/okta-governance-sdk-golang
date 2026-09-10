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

// checks if the ReviewerAiAgentConnectedResourceAuthorizationServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewerAiAgentConnectedResourceAuthorizationServer{}

// ReviewerAiAgentConnectedResourceAuthorizationServer Authorization server details for an AI agent connection
type ReviewerAiAgentConnectedResourceAuthorizationServer struct {
	// Okta Resource Name (ORN) of the authorization server
	Orn *string `json:"orn,omitempty"`
	// Name of the authorization server
	Name *string `json:"name,omitempty"`
	// Issuer URL of the connected authorization server
	IssuerUrl *string `json:"issuerUrl,omitempty"`
	// Resource indicator used when requesting tokens
	ResourceIndicator *string                            `json:"resourceIndicator,omitempty"`
	ScopeCondition    *AuthorizationServerScopeCondition `json:"scopeCondition,omitempty"`
	// Connected authorization server scopes that the AI agent has access to
	Scopes               []string `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewerAiAgentConnectedResourceAuthorizationServer ReviewerAiAgentConnectedResourceAuthorizationServer

// NewReviewerAiAgentConnectedResourceAuthorizationServer instantiates a new ReviewerAiAgentConnectedResourceAuthorizationServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewerAiAgentConnectedResourceAuthorizationServer() *ReviewerAiAgentConnectedResourceAuthorizationServer {
	this := ReviewerAiAgentConnectedResourceAuthorizationServer{}
	return &this
}

// NewReviewerAiAgentConnectedResourceAuthorizationServerWithDefaults instantiates a new ReviewerAiAgentConnectedResourceAuthorizationServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewerAiAgentConnectedResourceAuthorizationServerWithDefaults() *ReviewerAiAgentConnectedResourceAuthorizationServer {
	this := ReviewerAiAgentConnectedResourceAuthorizationServer{}
	return &this
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetOrn(v string) {
	o.Orn = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetName(v string) {
	o.Name = &v
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetIssuerUrl() string {
	if o == nil || IsNil(o.IssuerUrl) {
		var ret string
		return ret
	}
	return *o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetIssuerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerUrl) {
		return nil, false
	}
	return o.IssuerUrl, true
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasIssuerUrl() bool {
	if o != nil && !IsNil(o.IssuerUrl) {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given string and assigns it to the IssuerUrl field.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetIssuerUrl(v string) {
	o.IssuerUrl = &v
}

// GetResourceIndicator returns the ResourceIndicator field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetResourceIndicator() string {
	if o == nil || IsNil(o.ResourceIndicator) {
		var ret string
		return ret
	}
	return *o.ResourceIndicator
}

// GetResourceIndicatorOk returns a tuple with the ResourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetResourceIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceIndicator) {
		return nil, false
	}
	return o.ResourceIndicator, true
}

// HasResourceIndicator returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasResourceIndicator() bool {
	if o != nil && !IsNil(o.ResourceIndicator) {
		return true
	}

	return false
}

// SetResourceIndicator gets a reference to the given string and assigns it to the ResourceIndicator field.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetResourceIndicator(v string) {
	o.ResourceIndicator = &v
}

// GetScopeCondition returns the ScopeCondition field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopeCondition() AuthorizationServerScopeCondition {
	if o == nil || IsNil(o.ScopeCondition) {
		var ret AuthorizationServerScopeCondition
		return ret
	}
	return *o.ScopeCondition
}

// GetScopeConditionOk returns a tuple with the ScopeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopeConditionOk() (*AuthorizationServerScopeCondition, bool) {
	if o == nil || IsNil(o.ScopeCondition) {
		return nil, false
	}
	return o.ScopeCondition, true
}

// HasScopeCondition returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasScopeCondition() bool {
	if o != nil && !IsNil(o.ScopeCondition) {
		return true
	}

	return false
}

// SetScopeCondition gets a reference to the given AuthorizationServerScopeCondition and assigns it to the ScopeCondition field.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetScopeCondition(v AuthorizationServerScopeCondition) {
	o.ScopeCondition = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) SetScopes(v []string) {
	o.Scopes = v
}

func (o ReviewerAiAgentConnectedResourceAuthorizationServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewerAiAgentConnectedResourceAuthorizationServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IssuerUrl) {
		toSerialize["issuerUrl"] = o.IssuerUrl
	}
	if !IsNil(o.ResourceIndicator) {
		toSerialize["resourceIndicator"] = o.ResourceIndicator
	}
	if !IsNil(o.ScopeCondition) {
		toSerialize["scopeCondition"] = o.ScopeCondition
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReviewerAiAgentConnectedResourceAuthorizationServer) UnmarshalJSON(data []byte) (err error) {
	varReviewerAiAgentConnectedResourceAuthorizationServer := _ReviewerAiAgentConnectedResourceAuthorizationServer{}

	err = json.Unmarshal(data, &varReviewerAiAgentConnectedResourceAuthorizationServer)

	if err != nil {
		return err
	}

	*o = ReviewerAiAgentConnectedResourceAuthorizationServer(varReviewerAiAgentConnectedResourceAuthorizationServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "name")
		delete(additionalProperties, "issuerUrl")
		delete(additionalProperties, "resourceIndicator")
		delete(additionalProperties, "scopeCondition")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewerAiAgentConnectedResourceAuthorizationServer struct {
	value *ReviewerAiAgentConnectedResourceAuthorizationServer
	isSet bool
}

func (v NullableReviewerAiAgentConnectedResourceAuthorizationServer) Get() *ReviewerAiAgentConnectedResourceAuthorizationServer {
	return v.value
}

func (v *NullableReviewerAiAgentConnectedResourceAuthorizationServer) Set(val *ReviewerAiAgentConnectedResourceAuthorizationServer) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewerAiAgentConnectedResourceAuthorizationServer) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewerAiAgentConnectedResourceAuthorizationServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewerAiAgentConnectedResourceAuthorizationServer(val *ReviewerAiAgentConnectedResourceAuthorizationServer) *NullableReviewerAiAgentConnectedResourceAuthorizationServer {
	return &NullableReviewerAiAgentConnectedResourceAuthorizationServer{value: val, isSet: true}
}

func (v NullableReviewerAiAgentConnectedResourceAuthorizationServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewerAiAgentConnectedResourceAuthorizationServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
