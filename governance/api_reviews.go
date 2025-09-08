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
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/okta/okta-sdk-golang/v5/okta"
)

type ReviewsAPI interface {
	/*
			GetReview Retrieve a review

			Retrieves the full representation of a specific review.

		More information is returned than the abbreviated representation in a List reviews operation.
		Also note that, if reviews are still `UNREVIEWED`, then the property `decided` would be null. If the remediation is not completed, then `remediationStatus` would be null too.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param reviewId The `id` of the review
			@return ApiGetReviewRequest
	*/
	GetReview(ctx context.Context, reviewId string) ApiGetReviewRequest

	// GetReviewExecute executes the request
	//  @return ReviewFull
	GetReviewExecute(r ApiGetReviewRequest) (*ReviewFull, *APIResponse, error)

	/*
			ListReviews List all reviews

			Lists reviews for your organization.

		You can return a subset of reviews if a filter expression (`?filter=`) is provided.

		Supported filters are:

		- `campaignId`: string
		- `principalId`: string
		- `reviewerId`: string
		- `decision`: string (APPROVE, REVOKE, UNREVIEWED)
		- `resourceId`: string (`GroupId` or `AppId`)
		- `reviewerType`: string (`USER` or `GROUP` or `RESOURCE_OWNER`)
		- `reviewerLevel`: string (`FIRST` or `SECOND`)
		- `entitlementValueId`: string
		- `entitlementBundleId`: string

		Pagination parameters are accepted, and standard link headers are in the response.

		Reviews exist only for campaigns that have been launched with a status of `ACTIVE` or `COMPLETED`.
		Also note that, if reviews are still `UNREVIEWED`, then the property `decided` would be null. If the remediation is not completed, then `remediationStatus` would be null too.

		The order criteria (`orderBy`) applies to the following properties: `decided`, `decision`, `remediationStatus`, and `created`.

		By default, results are sorted by `created`.

		Note: Calling this endpoint without a filter would require fetching a large amount of data. If your org has a large number of campaigns and reviews,
		the request might time out or fail. To ensure reliable performance, we strongly recommend that you fetch data on a per-campaign basis using the provided filters.


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiListReviewsRequest
	*/
	ListReviews(ctx context.Context) ApiListReviewsRequest

	// ListReviewsExecute executes the request
	//  @return ReviewList
	ListReviewsExecute(r ApiListReviewsRequest) (*ReviewList, *APIResponse, error)

	/*
			ReassignReviews Reassign the reviews

			Reassigns a review to another reviewer.

		Reassigning a review is useful in cases where the currently assigned reviewer is not able to make a decision before the campaign ends.

		Only reviews belonging to campaigns with a status of `READY` can be reassigned. And only those reviews with decision `UNREVIEWED` can be reassigned.

		If reviews are currently being reviewed by a group (when `reviewerType` is `GROUP` or `RESOURCE_OWNER`), on reassignment, reviews is directly assigned to the chosen new reviewer and `reviewerType` is changed to `USER`.

		A valid reassignment changes an existing reviewer to a new reviewer (`userId`), and appends the change to the reassignment history.

		To reassign a set of reviews, you must specify:

		- the Okta `user.id` of the new reviewer

		- a list of `reviewIds` to be re-assigned

		- The `reviewerLevel` at which the reviews needs to be re-assigned (Applicable only for multi level campaigns)

		- a `note` justifying the reassignment decision for the specified reviews


			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param campaignId The `id` of the campaign
			@return ApiReassignReviewsRequest
	*/
	ReassignReviews(ctx context.Context, campaignId string) ApiReassignReviewsRequest

	// ReassignReviewsExecute executes the request
	//  @return ReviewReassignList
	ReassignReviewsExecute(r ApiReassignReviewsRequest) (*ReviewReassignList, *APIResponse, error)
}

// ReviewsAPIService ReviewsAPI service
type ReviewsAPIService service

type ApiGetReviewRequest struct {
	ctx        context.Context
	ApiService ReviewsAPI
	reviewId   string
	retryCount int32
}

func (r ApiGetReviewRequest) Execute() (*ReviewFull, *APIResponse, error) {
	return r.ApiService.GetReviewExecute(r)
}

/*
GetReview Retrieve a review

Retrieves the full representation of a specific review.

More information is returned than the abbreviated representation in a List reviews operation.
Also note that, if reviews are still `UNREVIEWED`, then the property `decided` would be null. If the remediation is not completed, then `remediationStatus` would be null too.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reviewId The `id` of the review
	@return ApiGetReviewRequest
*/
func (a *ReviewsAPIService) GetReview(ctx context.Context, reviewId string) ApiGetReviewRequest {
	return ApiGetReviewRequest{
		ApiService: a,
		ctx:        ctx,
		reviewId:   reviewId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return ReviewFull
func (a *ReviewsAPIService) GetReviewExecute(r ApiGetReviewRequest) (*ReviewFull, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReviewFull
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsAPIService.GetReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/reviews/{reviewId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reviewId"+"}", url.PathEscape(parameterToString(r.reviewId, "")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(okta.ContextAPIKeys).(map[string]okta.APIKey); ok {
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListReviewsRequest struct {
	ctx        context.Context
	ApiService ReviewsAPI
	filter     *string
	after      *string
	limit      *int32
	orderBy    *[]string
	retryCount int32
}

// Apply various filters by using supported review filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListReviewsRequest) Filter(filter string) ApiListReviewsRequest {
	r.filter = &filter
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListReviewsRequest) After(after string) ApiListReviewsRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListReviewsRequest) Limit(limit int32) ApiListReviewsRequest {
	r.limit = &limit
	return r
}

// A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListReviewsRequest) OrderBy(orderBy []string) ApiListReviewsRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiListReviewsRequest) Execute() (*ReviewList, *APIResponse, error) {
	return r.ApiService.ListReviewsExecute(r)
}

/*
ListReviews List all reviews

Lists reviews for your organization.

You can return a subset of reviews if a filter expression (`?filter=`) is provided.

Supported filters are:

- `campaignId`: string
- `principalId`: string
- `reviewerId`: string
- `decision`: string (APPROVE, REVOKE, UNREVIEWED)
- `resourceId`: string (`GroupId` or `AppId`)
- `reviewerType`: string (`USER` or `GROUP` or `RESOURCE_OWNER`)
- `reviewerLevel`: string (`FIRST` or `SECOND`)
- `entitlementValueId`: string
- `entitlementBundleId`: string

Pagination parameters are accepted, and standard link headers are in the response.

Reviews exist only for campaigns that have been launched with a status of `ACTIVE` or `COMPLETED`.
Also note that, if reviews are still `UNREVIEWED`, then the property `decided` would be null. If the remediation is not completed, then `remediationStatus` would be null too.

The order criteria (`orderBy`) applies to the following properties: `decided`, `decision`, `remediationStatus`, and `created`.

By default, results are sorted by `created`.

Note: Calling this endpoint without a filter would require fetching a large amount of data. If your org has a large number of campaigns and reviews,
the request might time out or fail. To ensure reliable performance, we strongly recommend that you fetch data on a per-campaign basis using the provided filters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListReviewsRequest
*/
func (a *ReviewsAPIService) ListReviews(ctx context.Context) ApiListReviewsRequest {
	return ApiListReviewsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return ReviewList
func (a *ReviewsAPIService) ListReviewsExecute(r ApiListReviewsRequest) (*ReviewList, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsAPIService.ListReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/reviews"

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
		if auth, ok := r.ctx.Value(okta.ContextAPIKeys).(map[string]okta.APIKey); ok {
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiReassignReviewsRequest struct {
	ctx             context.Context
	ApiService      ReviewsAPI
	campaignId      string
	reviewsReassign *ReviewsReassign
	retryCount      int32
}

// The operation payload for reviews reassignment
func (r ApiReassignReviewsRequest) ReviewsReassign(reviewsReassign ReviewsReassign) ApiReassignReviewsRequest {
	r.reviewsReassign = &reviewsReassign
	return r
}

func (r ApiReassignReviewsRequest) Execute() (*ReviewReassignList, *APIResponse, error) {
	return r.ApiService.ReassignReviewsExecute(r)
}

/*
ReassignReviews Reassign the reviews

Reassigns a review to another reviewer.

Reassigning a review is useful in cases where the currently assigned reviewer is not able to make a decision before the campaign ends.

Only reviews belonging to campaigns with a status of `READY` can be reassigned. And only those reviews with decision `UNREVIEWED` can be reassigned.

If reviews are currently being reviewed by a group (when `reviewerType` is `GROUP` or `RESOURCE_OWNER`), on reassignment, reviews is directly assigned to the chosen new reviewer and `reviewerType` is changed to `USER`.

A valid reassignment changes an existing reviewer to a new reviewer (`userId`), and appends the change to the reassignment history.

To reassign a set of reviews, you must specify:

- the Okta `user.id` of the new reviewer

- a list of `reviewIds` to be re-assigned

- The `reviewerLevel` at which the reviews needs to be re-assigned (Applicable only for multi level campaigns)

- a `note` justifying the reassignment decision for the specified reviews

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param campaignId The `id` of the campaign
	@return ApiReassignReviewsRequest
*/
func (a *ReviewsAPIService) ReassignReviews(ctx context.Context, campaignId string) ApiReassignReviewsRequest {
	return ApiReassignReviewsRequest{
		ApiService: a,
		ctx:        ctx,
		campaignId: campaignId,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return ReviewReassignList
func (a *ReviewsAPIService) ReassignReviewsExecute(r ApiReassignReviewsRequest) (*ReviewReassignList, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReviewReassignList
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReviewsAPIService.ReassignReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v1/campaigns/{campaignId}/reviews/reassign"
	localVarPath = strings.Replace(localVarPath, "{"+"campaignId"+"}", url.PathEscape(parameterToString(r.campaignId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.reviewsReassign == nil {
		return localVarReturnValue, nil, reportError("reviewsReassign is required and must be specified")
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
	localVarPostBody = r.reviewsReassign
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(okta.ContextAPIKeys).(map[string]okta.APIKey); ok {
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
