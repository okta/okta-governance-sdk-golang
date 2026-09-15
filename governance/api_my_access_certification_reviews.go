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
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MyAccessCertificationReviewsAPI interface {

	/*
		GetCampaignBulkDecisionsJobStatus Retrieve the status of a bulk-review submission

		Retrieves the job status of a bulk-review submission. Use this request to monitor the status of a bulk-review submission.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param campaignId Unique identifier for the campaign
		@param jobId Unique identifier for the job
		@return ApiGetCampaignBulkDecisionsJobStatusRequest
	*/
	GetCampaignBulkDecisionsJobStatus(ctx context.Context, campaignId string, jobId string) ApiGetCampaignBulkDecisionsJobStatusRequest

	// GetCampaignBulkDecisionsJobStatusExecute executes the request
	//  @return BulkReviewSubmissionJobDetails
	GetCampaignBulkDecisionsJobStatusExecute(r ApiGetCampaignBulkDecisionsJobStatusRequest) (*BulkReviewSubmissionJobDetails, *APIResponse, error)

	/*
			ListCampaignReviews List my reviews for a campaign

			Lists the reviews assigned to you for a given campaign.

		Use the filter expression (`?filter=`) query parameter to return a subset of your reviews for a campaign.

		Pagination parameters are accepted, and standard link headers are in the response.

		Reviews are returned only for campaigns that launched and have an `ACTIVE` or `COMPLETED` status.
		`UNREVIEWED` decision status reviews have a null `decided` property. If remediation isn't completed, then the `remediationStatus` property is also null.

		The order criteria (`orderBy`) applies to the following properties: `decided`, `decision`, `remediationStatus`, and `created`.

		By default, results are sorted by `id`.

		> **Note:** Okta recommends adding a filter in your request to target the reviews returned. This provides reliable performance when your org contains a large number of campaigns and reviews. The response for this request can time out or fail if there are a large number of campaigns and reviews to process.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param campaignId Unique identifier for the campaign
			@return ApiListCampaignReviewsRequest
	*/
	ListCampaignReviews(ctx context.Context, campaignId string) ApiListCampaignReviewsRequest

	// ListCampaignReviewsExecute executes the request
	//  @return ReviewList
	ListCampaignReviewsExecute(r ApiListCampaignReviewsRequest) (*ReviewList, *APIResponse, error)

	/*
		ListMyResourceConnections List all resource connections for my review

		Lists all resource connections for the specified AI agent in the current user's assigned review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param campaignId Unique identifier for the campaign
		@param reviewId Unique identifier for the review
		@param agentId ID of the agent
		@return ApiListMyResourceConnectionsRequest
	*/
	ListMyResourceConnections(ctx context.Context, campaignId string, reviewId string, agentId string) ApiListMyResourceConnectionsRequest

	// ListMyResourceConnectionsExecute executes the request
	//  @return MyResourceConnections
	ListMyResourceConnectionsExecute(r ApiListMyResourceConnectionsRequest) (*MyResourceConnections, *APIResponse, error)

	/*
			SubmitCampaignBulkDecisions Submit a bulk-review decision

			Submits a decision for bulk reviews in a campaign.
		The request body includes the criteria used to select the reviews for applying the decision.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param campaignId Unique identifier for the campaign
			@return ApiSubmitCampaignBulkDecisionsRequest
	*/
	SubmitCampaignBulkDecisions(ctx context.Context, campaignId string) ApiSubmitCampaignBulkDecisionsRequest

	// SubmitCampaignBulkDecisionsExecute executes the request
	//  @return BulkReviewSubmissionDetails
	SubmitCampaignBulkDecisionsExecute(r ApiSubmitCampaignBulkDecisionsRequest) (*BulkReviewSubmissionDetails, *APIResponse, error)

	/*
			SubmitMyCampaignReviewActions Submit my review actions

			Submits actions for reviews that are assigned to you in the specified campaign.

		You can specify a maximum of 50 review actions in a request.

		For each review, specify the `APPROVE`, `REVOKE`, or `REASSIGN` action.

		If you reassign a review (`action=REASSIGN`), you must also include the
		`reviewerId` and `justification` parameters. All your `REASSIGN` reviews must have
		the same `reviewerId` in a single request.

		A successful request returns the updated reviews based on the submitted actions.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param campaignId Unique identifier for the campaign
			@return ApiSubmitMyCampaignReviewActionsRequest
	*/
	SubmitMyCampaignReviewActions(ctx context.Context, campaignId string) ApiSubmitMyCampaignReviewActionsRequest

	// SubmitMyCampaignReviewActionsExecute executes the request
	//  @return ReviewList
	SubmitMyCampaignReviewActionsExecute(r ApiSubmitMyCampaignReviewActionsRequest) (*ReviewList, *APIResponse, error)
}

// MyAccessCertificationReviewsAPIService MyAccessCertificationReviewsAPI service
type MyAccessCertificationReviewsAPIService service

type ApiGetCampaignBulkDecisionsJobStatusRequest struct {
	ctx        context.Context
	ApiService MyAccessCertificationReviewsAPI
	campaignId string
	jobId      string
	retryCount int32
}

func (r ApiGetCampaignBulkDecisionsJobStatusRequest) Execute() (*BulkReviewSubmissionJobDetails, *APIResponse, error) {
	return r.ApiService.GetCampaignBulkDecisionsJobStatusExecute(r)
}

/*
GetCampaignBulkDecisionsJobStatus Retrieve the status of a bulk-review submission

Retrieves the job status of a bulk-review submission. Use this request to monitor the status of a bulk-review submission.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param campaignId Unique identifier for the campaign
	@param jobId Unique identifier for the job
	@return ApiGetCampaignBulkDecisionsJobStatusRequest
*/
func (a *MyAccessCertificationReviewsAPIService) GetCampaignBulkDecisionsJobStatus(ctx context.Context, campaignId string, jobId string) ApiGetCampaignBulkDecisionsJobStatusRequest {
	return ApiGetCampaignBulkDecisionsJobStatusRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
		jobId:      jobId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return BulkReviewSubmissionJobDetails
func (a *MyAccessCertificationReviewsAPIService) GetCampaignBulkDecisionsJobStatusExecute(r ApiGetCampaignBulkDecisionsJobStatusRequest) (*BulkReviewSubmissionJobDetails, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BulkReviewSubmissionJobDetails
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyAccessCertificationReviewsAPIService.GetCampaignBulkDecisionsJobStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/my/campaigns/{campaignId}/reviews/bulk-decisions/jobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterToString(r.campaignId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiListCampaignReviewsRequest struct {
	ctx        context.Context
	ApiService MyAccessCertificationReviewsAPI
	campaignId string
	filter     *string
	after      *string
	limit      *int32
	orderBy    *[]string
	retryCount int32
}

// Apply various filters by using supported review filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  Supported filters are:  - &#x60;id&#x60;: string - &#x60;principalId&#x60;: string - &#x60;reviewerId&#x60;: string - &#x60;decision&#x60;: string (APPROVE, REVOKE, UNREVIEWED) - &#x60;resourceId&#x60;: string (&#x60;GroupId&#x60; or &#x60;AppId&#x60;) - &#x60;reviewerType&#x60;: string (&#x60;USER&#x60; or &#x60;GROUP&#x60; or &#x60;RESOURCE_OWNER&#x60;) - &#x60;reviewerLevel&#x60;: string (&#x60;FIRST&#x60; or &#x60;SECOND&#x60;) - &#x60;entitlementValueId&#x60;: string - &#x60;entitlementBundleId&#x60;: string
func (r ApiListCampaignReviewsRequest) Filter(filter string) ApiListCampaignReviewsRequest {
	r.filter = &filter
	return r
}

// Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListCampaignReviewsRequest) After(after string) ApiListCampaignReviewsRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListCampaignReviewsRequest) Limit(limit int32) ApiListCampaignReviewsRequest {
	r.limit = &limit
	return r
}

// Specifies a property to sort the results. The following properties are supported: - &#x60;decided&#x60; - &#x60;decision&#x60; - &#x60;remediationStatus&#x60; - &#x60;created&#x60;  Append &#x60;desc&#x60; or &#x60;asc&#x60; to indicate the sorting direction.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListCampaignReviewsRequest) OrderBy(orderBy []string) ApiListCampaignReviewsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiListCampaignReviewsRequest) Execute() (*ReviewList, *APIResponse, error) {
	return r.ApiService.ListCampaignReviewsExecute(r)
}

/*
ListCampaignReviews List my reviews for a campaign

Lists the reviews assigned to you for a given campaign.

Use the filter expression (`?filter=`) query parameter to return a subset of your reviews for a campaign.

Pagination parameters are accepted, and standard link headers are in the response.

Reviews are returned only for campaigns that launched and have an `ACTIVE` or `COMPLETED` status.
`UNREVIEWED` decision status reviews have a null `decided` property. If remediation isn't completed, then the `remediationStatus` property is also null.

The order criteria (`orderBy`) applies to the following properties: `decided`, `decision`, `remediationStatus`, and `created`.

By default, results are sorted by `id`.

> **Note:** Okta recommends adding a filter in your request to target the reviews returned. This provides reliable performance when your org contains a large number of campaigns and reviews. The response for this request can time out or fail if there are a large number of campaigns and reviews to process.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param campaignId Unique identifier for the campaign
	@return ApiListCampaignReviewsRequest
*/
func (a *MyAccessCertificationReviewsAPIService) ListCampaignReviews(ctx context.Context, campaignId string) ApiListCampaignReviewsRequest {
	return ApiListCampaignReviewsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return ReviewList
func (a *MyAccessCertificationReviewsAPIService) ListCampaignReviewsExecute(r ApiListCampaignReviewsRequest) (*ReviewList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReviewList
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyAccessCertificationReviewsAPIService.ListCampaignReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/my/campaigns/{campaignId}/reviews"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterToString(r.campaignId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, "csv"))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiListMyResourceConnectionsRequest struct {
	ctx        context.Context
	ApiService MyAccessCertificationReviewsAPI
	campaignId string
	reviewId   string
	agentId    string
	after      *string
	limit      *int32
	retryCount int32
}

// Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination).
func (r ApiListMyResourceConnectionsRequest) After(after string) ApiListMyResourceConnectionsRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListMyResourceConnectionsRequest) Limit(limit int32) ApiListMyResourceConnectionsRequest {
	r.limit = &limit
	return r
}

func (r ApiListMyResourceConnectionsRequest) Execute() (*MyResourceConnections, *APIResponse, error) {
	return r.ApiService.ListMyResourceConnectionsExecute(r)
}

/*
ListMyResourceConnections List all resource connections for my review

Lists all resource connections for the specified AI agent in the current user's assigned review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param campaignId Unique identifier for the campaign
	@param reviewId Unique identifier for the review
	@param agentId ID of the agent
	@return ApiListMyResourceConnectionsRequest
*/
func (a *MyAccessCertificationReviewsAPIService) ListMyResourceConnections(ctx context.Context, campaignId string, reviewId string, agentId string) ApiListMyResourceConnectionsRequest {
	return ApiListMyResourceConnectionsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
		reviewId:   reviewId,
		agentId:    agentId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return MyResourceConnections
func (a *MyAccessCertificationReviewsAPIService) ListMyResourceConnectionsExecute(r ApiListMyResourceConnectionsRequest) (*MyResourceConnections, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MyResourceConnections
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyAccessCertificationReviewsAPIService.ListMyResourceConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/my/campaigns/{campaignId}/reviews/{reviewId}/agent-resource-connections/{agentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterToString(r.campaignId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"reviewId"+"}", url.PathEscape(parameterToString(r.reviewId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"agentId"+"}", url.PathEscape(parameterToString(r.agentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiSubmitCampaignBulkDecisionsRequest struct {
	ctx                         context.Context
	ApiService                  MyAccessCertificationReviewsAPI
	campaignId                  string
	bulkReviewSubmissionMutable *BulkReviewSubmissionMutable
	retryCount                  int32
}

// Request body for submitting a decision for bulk reviews
func (r ApiSubmitCampaignBulkDecisionsRequest) BulkReviewSubmissionMutable(bulkReviewSubmissionMutable BulkReviewSubmissionMutable) ApiSubmitCampaignBulkDecisionsRequest {
	r.bulkReviewSubmissionMutable = &bulkReviewSubmissionMutable
	return r
}

func (r ApiSubmitCampaignBulkDecisionsRequest) Execute() (*BulkReviewSubmissionDetails, *APIResponse, error) {
	return r.ApiService.SubmitCampaignBulkDecisionsExecute(r)
}

/*
SubmitCampaignBulkDecisions Submit a bulk-review decision

Submits a decision for bulk reviews in a campaign.
The request body includes the criteria used to select the reviews for applying the decision.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param campaignId Unique identifier for the campaign
	@return ApiSubmitCampaignBulkDecisionsRequest
*/
func (a *MyAccessCertificationReviewsAPIService) SubmitCampaignBulkDecisions(ctx context.Context, campaignId string) ApiSubmitCampaignBulkDecisionsRequest {
	return ApiSubmitCampaignBulkDecisionsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return BulkReviewSubmissionDetails
func (a *MyAccessCertificationReviewsAPIService) SubmitCampaignBulkDecisionsExecute(r ApiSubmitCampaignBulkDecisionsRequest) (*BulkReviewSubmissionDetails, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BulkReviewSubmissionDetails
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyAccessCertificationReviewsAPIService.SubmitCampaignBulkDecisions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/my/campaigns/{campaignId}/reviews/bulk-decisions"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterToString(r.campaignId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bulkReviewSubmissionMutable == nil {
		return localVarReturnValue, nil, reportError("bulkReviewSubmissionMutable is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.bulkReviewSubmissionMutable
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}

type ApiSubmitMyCampaignReviewActionsRequest struct {
	ctx                      context.Context
	ApiService               MyAccessCertificationReviewsAPI
	campaignId               string
	reviewItemsActionMutable *ReviewItemsActionMutable
	retryCount               int32
}

// Submit my review actions request body
func (r ApiSubmitMyCampaignReviewActionsRequest) ReviewItemsActionMutable(reviewItemsActionMutable ReviewItemsActionMutable) ApiSubmitMyCampaignReviewActionsRequest {
	r.reviewItemsActionMutable = &reviewItemsActionMutable
	return r
}

func (r ApiSubmitMyCampaignReviewActionsRequest) Execute() (*ReviewList, *APIResponse, error) {
	return r.ApiService.SubmitMyCampaignReviewActionsExecute(r)
}

/*
SubmitMyCampaignReviewActions Submit my review actions

Submits actions for reviews that are assigned to you in the specified campaign.

You can specify a maximum of 50 review actions in a request.

For each review, specify the `APPROVE`, `REVOKE`, or `REASSIGN` action.

If you reassign a review (`action=REASSIGN`), you must also include the
`reviewerId` and `justification` parameters. All your `REASSIGN` reviews must have
the same `reviewerId` in a single request.

A successful request returns the updated reviews based on the submitted actions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param campaignId Unique identifier for the campaign
	@return ApiSubmitMyCampaignReviewActionsRequest
*/
func (a *MyAccessCertificationReviewsAPIService) SubmitMyCampaignReviewActions(ctx context.Context, campaignId string) ApiSubmitMyCampaignReviewActionsRequest {
	return ApiSubmitMyCampaignReviewActionsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return ReviewList
func (a *MyAccessCertificationReviewsAPIService) SubmitMyCampaignReviewActionsExecute(r ApiSubmitMyCampaignReviewActionsRequest) (*ReviewList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReviewList
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyAccessCertificationReviewsAPIService.SubmitMyCampaignReviewActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/my/campaigns/{campaignId}/reviews/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterToString(r.campaignId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reviewItemsActionMutable == nil {
		return localVarReturnValue, nil, reportError("reviewItemsActionMutable is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.reviewItemsActionMutable
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
			return localVarReturnValue, localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
				return localVarReturnValue, localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
		return localVarReturnValue, localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, localVarReturnValue)
	return localVarReturnValue, localAPIResponse, nil
}
