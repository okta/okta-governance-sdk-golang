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
	"fmt"
)

// CampaignReviewerType Identifies the kind of reviewer for Access Certification. For example, a reviewer can be a user or an expression.
type CampaignReviewerType string

// List of campaign-reviewer-type
const (
	CAMPAIGNREVIEWERTYPE_USER                CampaignReviewerType = "USER"
	CAMPAIGNREVIEWERTYPE_REVIEWER_EXPRESSION CampaignReviewerType = "REVIEWER_EXPRESSION"
	CAMPAIGNREVIEWERTYPE_GROUP               CampaignReviewerType = "GROUP"
	CAMPAIGNREVIEWERTYPE_RESOURCE_OWNER      CampaignReviewerType = "RESOURCE_OWNER"
	CAMPAIGNREVIEWERTYPE_MULTI_LEVEL         CampaignReviewerType = "MULTI_LEVEL"
)

// All allowed values of CampaignReviewerType enum
var AllowedCampaignReviewerTypeEnumValues = []CampaignReviewerType{
	"USER",
	"REVIEWER_EXPRESSION",
	"GROUP",
	"RESOURCE_OWNER",
	"MULTI_LEVEL",
}

func (v *CampaignReviewerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignReviewerType(value)
	for _, existing := range AllowedCampaignReviewerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignReviewerType", value)
}

// NewCampaignReviewerTypeFromValue returns a pointer to a valid CampaignReviewerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignReviewerTypeFromValue(v string) (*CampaignReviewerType, error) {
	ev := CampaignReviewerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignReviewerType: valid values are %v", v, AllowedCampaignReviewerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignReviewerType) IsValid() bool {
	for _, existing := range AllowedCampaignReviewerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign-reviewer-type value
func (v CampaignReviewerType) Ptr() *CampaignReviewerType {
	return &v
}

type NullableCampaignReviewerType struct {
	value *CampaignReviewerType
	isSet bool
}

func (v NullableCampaignReviewerType) Get() *CampaignReviewerType {
	return v.value
}

func (v *NullableCampaignReviewerType) Set(val *CampaignReviewerType) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReviewerType) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReviewerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReviewerType(val *CampaignReviewerType) *NullableCampaignReviewerType {
	return &NullableCampaignReviewerType{value: val, isSet: true}
}

func (v NullableCampaignReviewerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReviewerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
