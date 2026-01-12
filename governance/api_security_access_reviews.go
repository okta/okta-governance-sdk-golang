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

type SecurityAccessReviewsAPI interface {

	/*
		AddSecurityAccessReviewComment Add a comment for a security access review

		Adds a comment for a specific security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiAddSecurityAccessReviewCommentRequest
	*/
	AddSecurityAccessReviewComment(ctx context.Context, securityAccessReviewId string) ApiAddSecurityAccessReviewCommentRequest

	// AddSecurityAccessReviewCommentExecute executes the request
	AddSecurityAccessReviewCommentExecute(r ApiAddSecurityAccessReviewCommentRequest) (*APIResponse, error)

	/*
		CreateSecurityAccessReview Create a security access review

		Creates a security access review for a target principal (such as a user)

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateSecurityAccessReviewRequest
	*/
	CreateSecurityAccessReview(ctx context.Context) ApiCreateSecurityAccessReviewRequest

	// CreateSecurityAccessReviewExecute executes the request
	//  @return SecurityAccessReview
	CreateSecurityAccessReviewExecute(r ApiCreateSecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error)

	/*
		ExecuteSecurityAccessReviewAccessesAction Execute an action on an access item

		Executes an action on an access or sub-access item in a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
		@return ApiExecuteSecurityAccessReviewAccessesActionRequest
	*/
	ExecuteSecurityAccessReviewAccessesAction(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiExecuteSecurityAccessReviewAccessesActionRequest

	// ExecuteSecurityAccessReviewAccessesActionExecute executes the request
	ExecuteSecurityAccessReviewAccessesActionExecute(r ApiExecuteSecurityAccessReviewAccessesActionRequest) (*APIResponse, error)

	/*
		ExecuteSecurityAccessReviewAction Execute an action on a security access review

		Executes a specified action on a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiExecuteSecurityAccessReviewActionRequest
	*/
	ExecuteSecurityAccessReviewAction(ctx context.Context, securityAccessReviewId string) ApiExecuteSecurityAccessReviewActionRequest

	// ExecuteSecurityAccessReviewActionExecute executes the request
	ExecuteSecurityAccessReviewActionExecute(r ApiExecuteSecurityAccessReviewActionRequest) (*APIResponse, error)

	/*
			GenerateSecurityAccessReviewAccessesSummary Generate a summary for an access item

			Generates a summary for an access item in a security access review

		> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
			@return ApiGenerateSecurityAccessReviewAccessesSummaryRequest
	*/
	GenerateSecurityAccessReviewAccessesSummary(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiGenerateSecurityAccessReviewAccessesSummaryRequest

	// GenerateSecurityAccessReviewAccessesSummaryExecute executes the request
	//  @return AiMessage
	GenerateSecurityAccessReviewAccessesSummaryExecute(r ApiGenerateSecurityAccessReviewAccessesSummaryRequest) (*AiMessage, *APIResponse, error)

	/*
			GenerateSecurityAccessReviewSummary Generate a summary for a security access review

			Generates a summary for a security access review

		> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@return ApiGenerateSecurityAccessReviewSummaryRequest
	*/
	GenerateSecurityAccessReviewSummary(ctx context.Context, securityAccessReviewId string) ApiGenerateSecurityAccessReviewSummaryRequest

	// GenerateSecurityAccessReviewSummaryExecute executes the request
	//  @return AiMessage
	GenerateSecurityAccessReviewSummaryExecute(r ApiGenerateSecurityAccessReviewSummaryRequest) (*AiMessage, *APIResponse, error)

	/*
		GetSecurityAccessReview Retrieve a security access review

		Retrieves a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiGetSecurityAccessReviewRequest
	*/
	GetSecurityAccessReview(ctx context.Context, securityAccessReviewId string) ApiGetSecurityAccessReviewRequest

	// GetSecurityAccessReviewExecute executes the request
	//  @return SecurityAccessReview
	GetSecurityAccessReviewExecute(r ApiGetSecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error)

	/*
		GetSecurityAccessReviewPrincipalDetails Retrieve the principal for a security access review

		Retrieves the details of a security access review's principal target

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiGetSecurityAccessReviewPrincipalDetailsRequest
	*/
	GetSecurityAccessReviewPrincipalDetails(ctx context.Context, securityAccessReviewId string) ApiGetSecurityAccessReviewPrincipalDetailsRequest

	// GetSecurityAccessReviewPrincipalDetailsExecute executes the request
	//  @return SecurityAccessReviewPrincipal
	GetSecurityAccessReviewPrincipalDetailsExecute(r ApiGetSecurityAccessReviewPrincipalDetailsRequest) (*SecurityAccessReviewPrincipal, *APIResponse, error)

	/*
		GetSecurityAccessReviewsStats Retrieve the statistics for security access reviews

		Retrieves the statistics for security access reviews in an org

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetSecurityAccessReviewsStatsRequest
	*/
	GetSecurityAccessReviewsStats(ctx context.Context) ApiGetSecurityAccessReviewsStatsRequest

	// GetSecurityAccessReviewsStatsExecute executes the request
	//  @return SecurityAccessReviewsStats
	GetSecurityAccessReviewsStatsExecute(r ApiGetSecurityAccessReviewsStatsRequest) (*SecurityAccessReviewsStats, *APIResponse, error)

	/*
			ListSecurityAccessReviewAccesses List the access items for a security access review

			Lists the access items for a specific security access review.

		Access items refer to the top-level resources that the security access review's target principal has access to.
		For example, a top-level resource can be an app, hence the access item describes the principal's access to that app.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@return ApiListSecurityAccessReviewAccessesRequest
	*/
	ListSecurityAccessReviewAccesses(ctx context.Context, securityAccessReviewId string) ApiListSecurityAccessReviewAccessesRequest

	// ListSecurityAccessReviewAccessesExecute executes the request
	//  @return SecurityAccessReviewAccessItems
	ListSecurityAccessReviewAccessesExecute(r ApiListSecurityAccessReviewAccessesRequest) (*SecurityAccessReviewAccessItems, *APIResponse, error)

	/*
		ListSecurityAccessReviewAccessesAnomalies List the anomalies for an access item

		Lists the anomalies for an access item

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
		@return ApiListSecurityAccessReviewAccessesAnomaliesRequest
	*/
	ListSecurityAccessReviewAccessesAnomalies(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiListSecurityAccessReviewAccessesAnomaliesRequest

	// ListSecurityAccessReviewAccessesAnomaliesExecute executes the request
	//  @return SecurityAccessReviewAnomalies
	ListSecurityAccessReviewAccessesAnomaliesExecute(r ApiListSecurityAccessReviewAccessesAnomaliesRequest) (*SecurityAccessReviewAnomalies, *APIResponse, error)

	/*
		ListSecurityAccessReviewActions List the actions for a security access review

		Lists the available actions for a specific security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiListSecurityAccessReviewActionsRequest
	*/
	ListSecurityAccessReviewActions(ctx context.Context, securityAccessReviewId string) ApiListSecurityAccessReviewActionsRequest

	// ListSecurityAccessReviewActionsExecute executes the request
	//  @return SecurityAccessReviewActions
	ListSecurityAccessReviewActionsExecute(r ApiListSecurityAccessReviewActionsRequest) (*SecurityAccessReviewActions, *APIResponse, error)

	/*
		ListSecurityAccessReviewHistory List the history of a security access review

		Lists the history of actions and changes for a specific security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiListSecurityAccessReviewHistoryRequest
	*/
	ListSecurityAccessReviewHistory(ctx context.Context, securityAccessReviewId string) ApiListSecurityAccessReviewHistoryRequest

	// ListSecurityAccessReviewHistoryExecute executes the request
	//  @return SecurityAccessReviewHistoryItems
	ListSecurityAccessReviewHistoryExecute(r ApiListSecurityAccessReviewHistoryRequest) (*SecurityAccessReviewHistoryItems, *APIResponse, error)

	/*
			ListSecurityAccessReviewSubAccesses List the sub-access items for an access item

			Lists the sub-access items for an access item from a security access review.

		A sub-access item refers to the access of a resource that is a part of a top-level resource in an access item.
		For example, an access item can describe the access of app, and the sub-access items can describe the access of groups, entitlement values, or entitlement bundles that belong to the app.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@param securityAccessReviewAccessId The ID of the access item in a security access review
			@return ApiListSecurityAccessReviewSubAccessesRequest
	*/
	ListSecurityAccessReviewSubAccesses(ctx context.Context, securityAccessReviewId string, securityAccessReviewAccessId string) ApiListSecurityAccessReviewSubAccessesRequest

	// ListSecurityAccessReviewSubAccessesExecute executes the request
	//  @return SecurityAccessReviewSubAccessItems
	ListSecurityAccessReviewSubAccessesExecute(r ApiListSecurityAccessReviewSubAccessesRequest) (*SecurityAccessReviewSubAccessItems, *APIResponse, error)

	/*
		ListSecurityAccessReviews List all security access reviews

		Lists all security access reviews

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListSecurityAccessReviewsRequest
	*/
	ListSecurityAccessReviews(ctx context.Context) ApiListSecurityAccessReviewsRequest

	// ListSecurityAccessReviewsExecute executes the request
	//  @return SecurityAccessReviews
	ListSecurityAccessReviewsExecute(r ApiListSecurityAccessReviewsRequest) (*SecurityAccessReviews, *APIResponse, error)

	/*
		UpdateSecurityAccessReview Update a security access review

		Updates a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiUpdateSecurityAccessReviewRequest
	*/
	UpdateSecurityAccessReview(ctx context.Context, securityAccessReviewId string) ApiUpdateSecurityAccessReviewRequest

	// UpdateSecurityAccessReviewExecute executes the request
	//  @return SecurityAccessReview
	UpdateSecurityAccessReviewExecute(r ApiUpdateSecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error)
}

// SecurityAccessReviewsAPIService SecurityAccessReviewsAPI service
type SecurityAccessReviewsAPIService service

type ApiAddSecurityAccessReviewCommentRequest struct {
	ctx                                context.Context
	ApiService                         SecurityAccessReviewsAPI
	securityAccessReviewId             string
	securityAccessReviewCommentRequest *SecurityAccessReviewCommentRequest
	retryCount                         int32
}

func (r ApiAddSecurityAccessReviewCommentRequest) SecurityAccessReviewCommentRequest(securityAccessReviewCommentRequest SecurityAccessReviewCommentRequest) ApiAddSecurityAccessReviewCommentRequest {
	r.securityAccessReviewCommentRequest = &securityAccessReviewCommentRequest
	return r
}

func (r ApiAddSecurityAccessReviewCommentRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AddSecurityAccessReviewCommentExecute(r)
}

/*
AddSecurityAccessReviewComment Add a comment for a security access review

Adds a comment for a specific security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiAddSecurityAccessReviewCommentRequest
*/
func (a *SecurityAccessReviewsAPIService) AddSecurityAccessReviewComment(ctx context.Context, securityAccessReviewId string) ApiAddSecurityAccessReviewCommentRequest {
	return ApiAddSecurityAccessReviewCommentRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
func (a *SecurityAccessReviewsAPIService) AddSecurityAccessReviewCommentExecute(r ApiAddSecurityAccessReviewCommentRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.AddSecurityAccessReviewComment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/comment"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityAccessReviewCommentRequest == nil {
		return nil, reportError("securityAccessReviewCommentRequest is required and must be specified")
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
	localVarPostBody = r.securityAccessReviewCommentRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
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
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiCreateSecurityAccessReviewRequest struct {
	ctx                         context.Context
	ApiService                  SecurityAccessReviewsAPI
	securityAccessReviewRequest *SecurityAccessReviewRequest
	retryCount                  int32
}

func (r ApiCreateSecurityAccessReviewRequest) SecurityAccessReviewRequest(securityAccessReviewRequest SecurityAccessReviewRequest) ApiCreateSecurityAccessReviewRequest {
	r.securityAccessReviewRequest = &securityAccessReviewRequest
	return r
}

func (r ApiCreateSecurityAccessReviewRequest) Execute() (*SecurityAccessReview, *APIResponse, error) {
	return r.ApiService.CreateSecurityAccessReviewExecute(r)
}

/*
CreateSecurityAccessReview Create a security access review

Creates a security access review for a target principal (such as a user)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateSecurityAccessReviewRequest
*/
func (a *SecurityAccessReviewsAPIService) CreateSecurityAccessReview(ctx context.Context) ApiCreateSecurityAccessReviewRequest {
	return ApiCreateSecurityAccessReviewRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReview
func (a *SecurityAccessReviewsAPIService) CreateSecurityAccessReviewExecute(r ApiCreateSecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReview
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.CreateSecurityAccessReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityAccessReviewRequest == nil {
		return localVarReturnValue, nil, reportError("securityAccessReviewRequest is required and must be specified")
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
	localVarPostBody = r.securityAccessReviewRequest
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

type ApiExecuteSecurityAccessReviewAccessesActionRequest struct {
	ctx                                       context.Context
	ApiService                                SecurityAccessReviewsAPI
	securityAccessReviewId                    string
	securityAccessReviewTargetId              string
	securityAccessReviewAccessesActionRequest *SecurityAccessReviewAccessesActionRequest
	retryCount                                int32
}

func (r ApiExecuteSecurityAccessReviewAccessesActionRequest) SecurityAccessReviewAccessesActionRequest(securityAccessReviewAccessesActionRequest SecurityAccessReviewAccessesActionRequest) ApiExecuteSecurityAccessReviewAccessesActionRequest {
	r.securityAccessReviewAccessesActionRequest = &securityAccessReviewAccessesActionRequest
	return r
}

func (r ApiExecuteSecurityAccessReviewAccessesActionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ExecuteSecurityAccessReviewAccessesActionExecute(r)
}

/*
ExecuteSecurityAccessReviewAccessesAction Execute an action on an access item

Executes an action on an access or sub-access item in a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
	@return ApiExecuteSecurityAccessReviewAccessesActionRequest
*/
func (a *SecurityAccessReviewsAPIService) ExecuteSecurityAccessReviewAccessesAction(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiExecuteSecurityAccessReviewAccessesActionRequest {
	return ApiExecuteSecurityAccessReviewAccessesActionRequest{
		ApiService:                   a,
		ctx:                          ctx,
		securityAccessReviewId:       securityAccessReviewId,
		securityAccessReviewTargetId: securityAccessReviewTargetId,
		retryCount:                   0,
	}
}

// Execute executes the request
func (a *SecurityAccessReviewsAPIService) ExecuteSecurityAccessReviewAccessesActionExecute(r ApiExecuteSecurityAccessReviewAccessesActionRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ExecuteSecurityAccessReviewAccessesAction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewTargetId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewTargetId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityAccessReviewAccessesActionRequest == nil {
		return nil, reportError("securityAccessReviewAccessesActionRequest is required and must be specified")
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
	localVarPostBody = r.securityAccessReviewAccessesActionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
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
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiExecuteSecurityAccessReviewActionRequest struct {
	ctx                               context.Context
	ApiService                        SecurityAccessReviewsAPI
	securityAccessReviewId            string
	securityAccessReviewActionRequest *SecurityAccessReviewActionRequest
	retryCount                        int32
}

func (r ApiExecuteSecurityAccessReviewActionRequest) SecurityAccessReviewActionRequest(securityAccessReviewActionRequest SecurityAccessReviewActionRequest) ApiExecuteSecurityAccessReviewActionRequest {
	r.securityAccessReviewActionRequest = &securityAccessReviewActionRequest
	return r
}

func (r ApiExecuteSecurityAccessReviewActionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ExecuteSecurityAccessReviewActionExecute(r)
}

/*
ExecuteSecurityAccessReviewAction Execute an action on a security access review

Executes a specified action on a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiExecuteSecurityAccessReviewActionRequest
*/
func (a *SecurityAccessReviewsAPIService) ExecuteSecurityAccessReviewAction(ctx context.Context, securityAccessReviewId string) ApiExecuteSecurityAccessReviewActionRequest {
	return ApiExecuteSecurityAccessReviewActionRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
func (a *SecurityAccessReviewsAPIService) ExecuteSecurityAccessReviewActionExecute(r ApiExecuteSecurityAccessReviewActionRequest) (*APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ExecuteSecurityAccessReviewAction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityAccessReviewActionRequest == nil {
		return nil, reportError("securityAccessReviewActionRequest is required and must be specified")
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
	localVarPostBody = r.securityAccessReviewActionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}
	localVarHTTPResponse, err = a.client.do(r.ctx, req)
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, &GenericOpenAPIError{error: err.Error()}
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, err
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
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
			localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
			return localAPIResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
				return localAPIResponse, newErr
			}
			newErr.model = v
		}
		localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
		return localAPIResponse, newErr
	}

	localAPIResponse = newAPIResponse(localVarHTTPResponse, a.client, nil)
	return localAPIResponse, nil
}

type ApiGenerateSecurityAccessReviewAccessesSummaryRequest struct {
	ctx                          context.Context
	ApiService                   SecurityAccessReviewsAPI
	securityAccessReviewId       string
	securityAccessReviewTargetId string
	retryCount                   int32
}

func (r ApiGenerateSecurityAccessReviewAccessesSummaryRequest) Execute() (*AiMessage, *APIResponse, error) {
	return r.ApiService.GenerateSecurityAccessReviewAccessesSummaryExecute(r)
}

/*
GenerateSecurityAccessReviewAccessesSummary Generate a summary for an access item

# Generates a summary for an access item in a security access review

> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
	@return ApiGenerateSecurityAccessReviewAccessesSummaryRequest
*/
func (a *SecurityAccessReviewsAPIService) GenerateSecurityAccessReviewAccessesSummary(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiGenerateSecurityAccessReviewAccessesSummaryRequest {
	return ApiGenerateSecurityAccessReviewAccessesSummaryRequest{
		ApiService:                   a,
		ctx:                          ctx,
		securityAccessReviewId:       securityAccessReviewId,
		securityAccessReviewTargetId: securityAccessReviewTargetId,
		retryCount:                   0,
	}
}

// Execute executes the request
//
//	@return AiMessage
func (a *SecurityAccessReviewsAPIService) GenerateSecurityAccessReviewAccessesSummaryExecute(r ApiGenerateSecurityAccessReviewAccessesSummaryRequest) (*AiMessage, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiMessage
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.GenerateSecurityAccessReviewAccessesSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewTargetId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewTargetId, "")), -1)

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

type ApiGenerateSecurityAccessReviewSummaryRequest struct {
	ctx                    context.Context
	ApiService             SecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiGenerateSecurityAccessReviewSummaryRequest) Execute() (*AiMessage, *APIResponse, error) {
	return r.ApiService.GenerateSecurityAccessReviewSummaryExecute(r)
}

/*
GenerateSecurityAccessReviewSummary Generate a summary for a security access review

# Generates a summary for a security access review

> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiGenerateSecurityAccessReviewSummaryRequest
*/
func (a *SecurityAccessReviewsAPIService) GenerateSecurityAccessReviewSummary(ctx context.Context, securityAccessReviewId string) ApiGenerateSecurityAccessReviewSummaryRequest {
	return ApiGenerateSecurityAccessReviewSummaryRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return AiMessage
func (a *SecurityAccessReviewsAPIService) GenerateSecurityAccessReviewSummaryExecute(r ApiGenerateSecurityAccessReviewSummaryRequest) (*AiMessage, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AiMessage
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.GenerateSecurityAccessReviewSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

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

type ApiGetSecurityAccessReviewRequest struct {
	ctx                    context.Context
	ApiService             SecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiGetSecurityAccessReviewRequest) Execute() (*SecurityAccessReview, *APIResponse, error) {
	return r.ApiService.GetSecurityAccessReviewExecute(r)
}

/*
GetSecurityAccessReview Retrieve a security access review

Retrieves a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiGetSecurityAccessReviewRequest
*/
func (a *SecurityAccessReviewsAPIService) GetSecurityAccessReview(ctx context.Context, securityAccessReviewId string) ApiGetSecurityAccessReviewRequest {
	return ApiGetSecurityAccessReviewRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReview
func (a *SecurityAccessReviewsAPIService) GetSecurityAccessReviewExecute(r ApiGetSecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReview
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.GetSecurityAccessReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

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

type ApiGetSecurityAccessReviewPrincipalDetailsRequest struct {
	ctx                    context.Context
	ApiService             SecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiGetSecurityAccessReviewPrincipalDetailsRequest) Execute() (*SecurityAccessReviewPrincipal, *APIResponse, error) {
	return r.ApiService.GetSecurityAccessReviewPrincipalDetailsExecute(r)
}

/*
GetSecurityAccessReviewPrincipalDetails Retrieve the principal for a security access review

Retrieves the details of a security access review's principal target

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiGetSecurityAccessReviewPrincipalDetailsRequest
*/
func (a *SecurityAccessReviewsAPIService) GetSecurityAccessReviewPrincipalDetails(ctx context.Context, securityAccessReviewId string) ApiGetSecurityAccessReviewPrincipalDetailsRequest {
	return ApiGetSecurityAccessReviewPrincipalDetailsRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewPrincipal
func (a *SecurityAccessReviewsAPIService) GetSecurityAccessReviewPrincipalDetailsExecute(r ApiGetSecurityAccessReviewPrincipalDetailsRequest) (*SecurityAccessReviewPrincipal, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewPrincipal
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.GetSecurityAccessReviewPrincipalDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/principal"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

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

type ApiGetSecurityAccessReviewsStatsRequest struct {
	ctx        context.Context
	ApiService SecurityAccessReviewsAPI
	retryCount int32
}

func (r ApiGetSecurityAccessReviewsStatsRequest) Execute() (*SecurityAccessReviewsStats, *APIResponse, error) {
	return r.ApiService.GetSecurityAccessReviewsStatsExecute(r)
}

/*
GetSecurityAccessReviewsStats Retrieve the statistics for security access reviews

Retrieves the statistics for security access reviews in an org

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSecurityAccessReviewsStatsRequest
*/
func (a *SecurityAccessReviewsAPIService) GetSecurityAccessReviewsStats(ctx context.Context) ApiGetSecurityAccessReviewsStatsRequest {
	return ApiGetSecurityAccessReviewsStatsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewsStats
func (a *SecurityAccessReviewsAPIService) GetSecurityAccessReviewsStatsExecute(r ApiGetSecurityAccessReviewsStatsRequest) (*SecurityAccessReviewsStats, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewsStats
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.GetSecurityAccessReviewsStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/stats"

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

type ApiListSecurityAccessReviewAccessesRequest struct {
	ctx                    context.Context
	ApiService             SecurityAccessReviewsAPI
	securityAccessReviewId string
	filter                 *string
	orderBy                *[]string
	after                  *string
	limit                  *int32
	retryCount             int32
}

// A [filter](https://developer.okta.com/docs/api/#filter) expression that filters access items.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListSecurityAccessReviewAccessesRequest) Filter(filter string) ApiListSecurityAccessReviewAccessesRequest {
	r.filter = &filter
	return r
}

// The field to sort the results, in ascending (asc) or descending (desc) order. Sorting is applied to only one field.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListSecurityAccessReviewAccessesRequest) OrderBy(orderBy []string) ApiListSecurityAccessReviewAccessesRequest {
	r.orderBy = &orderBy
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListSecurityAccessReviewAccessesRequest) After(after string) ApiListSecurityAccessReviewAccessesRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListSecurityAccessReviewAccessesRequest) Limit(limit int32) ApiListSecurityAccessReviewAccessesRequest {
	r.limit = &limit
	return r
}

func (r ApiListSecurityAccessReviewAccessesRequest) Execute() (*SecurityAccessReviewAccessItems, *APIResponse, error) {
	return r.ApiService.ListSecurityAccessReviewAccessesExecute(r)
}

/*
ListSecurityAccessReviewAccesses List the access items for a security access review

Lists the access items for a specific security access review.

Access items refer to the top-level resources that the security access review's target principal has access to.
For example, a top-level resource can be an app, hence the access item describes the principal's access to that app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiListSecurityAccessReviewAccessesRequest
*/
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewAccesses(ctx context.Context, securityAccessReviewId string) ApiListSecurityAccessReviewAccessesRequest {
	return ApiListSecurityAccessReviewAccessesRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewAccessItems
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewAccessesExecute(r ApiListSecurityAccessReviewAccessesRequest) (*SecurityAccessReviewAccessItems, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewAccessItems
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ListSecurityAccessReviewAccesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, "csv"))
	}
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

type ApiListSecurityAccessReviewAccessesAnomaliesRequest struct {
	ctx                          context.Context
	ApiService                   SecurityAccessReviewsAPI
	securityAccessReviewId       string
	securityAccessReviewTargetId string
	retryCount                   int32
}

func (r ApiListSecurityAccessReviewAccessesAnomaliesRequest) Execute() (*SecurityAccessReviewAnomalies, *APIResponse, error) {
	return r.ApiService.ListSecurityAccessReviewAccessesAnomaliesExecute(r)
}

/*
ListSecurityAccessReviewAccessesAnomalies List the anomalies for an access item

Lists the anomalies for an access item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
	@return ApiListSecurityAccessReviewAccessesAnomaliesRequest
*/
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewAccessesAnomalies(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiListSecurityAccessReviewAccessesAnomaliesRequest {
	return ApiListSecurityAccessReviewAccessesAnomaliesRequest{
		ApiService:                   a,
		ctx:                          ctx,
		securityAccessReviewId:       securityAccessReviewId,
		securityAccessReviewTargetId: securityAccessReviewTargetId,
		retryCount:                   0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewAnomalies
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewAccessesAnomaliesExecute(r ApiListSecurityAccessReviewAccessesAnomaliesRequest) (*SecurityAccessReviewAnomalies, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewAnomalies
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ListSecurityAccessReviewAccessesAnomalies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/anomalies"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewTargetId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewTargetId, "")), -1)

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

type ApiListSecurityAccessReviewActionsRequest struct {
	ctx                    context.Context
	ApiService             SecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiListSecurityAccessReviewActionsRequest) Execute() (*SecurityAccessReviewActions, *APIResponse, error) {
	return r.ApiService.ListSecurityAccessReviewActionsExecute(r)
}

/*
ListSecurityAccessReviewActions List the actions for a security access review

Lists the available actions for a specific security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiListSecurityAccessReviewActionsRequest
*/
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewActions(ctx context.Context, securityAccessReviewId string) ApiListSecurityAccessReviewActionsRequest {
	return ApiListSecurityAccessReviewActionsRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewActions
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewActionsExecute(r ApiListSecurityAccessReviewActionsRequest) (*SecurityAccessReviewActions, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewActions
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ListSecurityAccessReviewActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/actions"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

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

type ApiListSecurityAccessReviewHistoryRequest struct {
	ctx                    context.Context
	ApiService             SecurityAccessReviewsAPI
	securityAccessReviewId string
	after                  *string
	limit                  *int32
	retryCount             int32
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListSecurityAccessReviewHistoryRequest) After(after string) ApiListSecurityAccessReviewHistoryRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListSecurityAccessReviewHistoryRequest) Limit(limit int32) ApiListSecurityAccessReviewHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiListSecurityAccessReviewHistoryRequest) Execute() (*SecurityAccessReviewHistoryItems, *APIResponse, error) {
	return r.ApiService.ListSecurityAccessReviewHistoryExecute(r)
}

/*
ListSecurityAccessReviewHistory List the history of a security access review

Lists the history of actions and changes for a specific security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiListSecurityAccessReviewHistoryRequest
*/
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewHistory(ctx context.Context, securityAccessReviewId string) ApiListSecurityAccessReviewHistoryRequest {
	return ApiListSecurityAccessReviewHistoryRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewHistoryItems
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewHistoryExecute(r ApiListSecurityAccessReviewHistoryRequest) (*SecurityAccessReviewHistoryItems, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewHistoryItems
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ListSecurityAccessReviewHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

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

type ApiListSecurityAccessReviewSubAccessesRequest struct {
	ctx                          context.Context
	ApiService                   SecurityAccessReviewsAPI
	securityAccessReviewId       string
	securityAccessReviewAccessId string
	filter                       *string
	orderBy                      *[]string
	after                        *string
	limit                        *int32
	retryCount                   int32
}

// A [filter](https://developer.okta.com/docs/api/#filter) expression that filters sub-access items.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListSecurityAccessReviewSubAccessesRequest) Filter(filter string) ApiListSecurityAccessReviewSubAccessesRequest {
	r.filter = &filter
	return r
}

// A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListSecurityAccessReviewSubAccessesRequest) OrderBy(orderBy []string) ApiListSecurityAccessReviewSubAccessesRequest {
	r.orderBy = &orderBy
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListSecurityAccessReviewSubAccessesRequest) After(after string) ApiListSecurityAccessReviewSubAccessesRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListSecurityAccessReviewSubAccessesRequest) Limit(limit int32) ApiListSecurityAccessReviewSubAccessesRequest {
	r.limit = &limit
	return r
}

func (r ApiListSecurityAccessReviewSubAccessesRequest) Execute() (*SecurityAccessReviewSubAccessItems, *APIResponse, error) {
	return r.ApiService.ListSecurityAccessReviewSubAccessesExecute(r)
}

/*
ListSecurityAccessReviewSubAccesses List the sub-access items for an access item

Lists the sub-access items for an access item from a security access review.

A sub-access item refers to the access of a resource that is a part of a top-level resource in an access item.
For example, an access item can describe the access of app, and the sub-access items can describe the access of groups, entitlement values, or entitlement bundles that belong to the app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewAccessId The ID of the access item in a security access review
	@return ApiListSecurityAccessReviewSubAccessesRequest
*/
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewSubAccesses(ctx context.Context, securityAccessReviewId string, securityAccessReviewAccessId string) ApiListSecurityAccessReviewSubAccessesRequest {
	return ApiListSecurityAccessReviewSubAccessesRequest{
		ApiService:                   a,
		ctx:                          ctx,
		securityAccessReviewId:       securityAccessReviewId,
		securityAccessReviewAccessId: securityAccessReviewAccessId,
		retryCount:                   0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewSubAccessItems
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewSubAccessesExecute(r ApiListSecurityAccessReviewSubAccessesRequest) (*SecurityAccessReviewSubAccessItems, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviewSubAccessItems
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ListSecurityAccessReviewSubAccesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewAccessId}/sub-accesses"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewAccessId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewAccessId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, "csv"))
	}
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

type ApiListSecurityAccessReviewsRequest struct {
	ctx        context.Context
	ApiService SecurityAccessReviewsAPI
	filter     *string
	orderBy    *[]string
	after      *string
	limit      *int32
	retryCount int32
}

// A [filter](https://developer.okta.com/docs/api/#filter) expression that filters security access reviews. The &#x60;eq&#x60; and &#x60;co&#x60; [operators](https://developer.okta.com/docs/api/#operators) are supported for string properties. The &#x60;gt&#x60; and &#x60;lt&#x60; operators are supported for date properties.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListSecurityAccessReviewsRequest) Filter(filter string) ApiListSecurityAccessReviewsRequest {
	r.filter = &filter
	return r
}

// The field to sort the results, in ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;) order. Sorting is applied to only one field.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListSecurityAccessReviewsRequest) OrderBy(orderBy []string) ApiListSecurityAccessReviewsRequest {
	r.orderBy = &orderBy
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListSecurityAccessReviewsRequest) After(after string) ApiListSecurityAccessReviewsRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListSecurityAccessReviewsRequest) Limit(limit int32) ApiListSecurityAccessReviewsRequest {
	r.limit = &limit
	return r
}

func (r ApiListSecurityAccessReviewsRequest) Execute() (*SecurityAccessReviews, *APIResponse, error) {
	return r.ApiService.ListSecurityAccessReviewsExecute(r)
}

/*
ListSecurityAccessReviews List all security access reviews

Lists all security access reviews

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListSecurityAccessReviewsRequest
*/
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviews(ctx context.Context) ApiListSecurityAccessReviewsRequest {
	return ApiListSecurityAccessReviewsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviews
func (a *SecurityAccessReviewsAPIService) ListSecurityAccessReviewsExecute(r ApiListSecurityAccessReviewsRequest) (*SecurityAccessReviews, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReviews
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.ListSecurityAccessReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("orderBy", parameterToString(*r.orderBy, "csv"))
	}
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

type ApiUpdateSecurityAccessReviewRequest struct {
	ctx                               context.Context
	ApiService                        SecurityAccessReviewsAPI
	securityAccessReviewId            string
	securityAccessReviewUpdateRequest *SecurityAccessReviewUpdateRequest
	retryCount                        int32
}

func (r ApiUpdateSecurityAccessReviewRequest) SecurityAccessReviewUpdateRequest(securityAccessReviewUpdateRequest SecurityAccessReviewUpdateRequest) ApiUpdateSecurityAccessReviewRequest {
	r.securityAccessReviewUpdateRequest = &securityAccessReviewUpdateRequest
	return r
}

func (r ApiUpdateSecurityAccessReviewRequest) Execute() (*SecurityAccessReview, *APIResponse, error) {
	return r.ApiService.UpdateSecurityAccessReviewExecute(r)
}

/*
UpdateSecurityAccessReview Update a security access review

Updates a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiUpdateSecurityAccessReviewRequest
*/
func (a *SecurityAccessReviewsAPIService) UpdateSecurityAccessReview(ctx context.Context, securityAccessReviewId string) ApiUpdateSecurityAccessReviewRequest {
	return ApiUpdateSecurityAccessReviewRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReview
func (a *SecurityAccessReviewsAPIService) UpdateSecurityAccessReviewExecute(r ApiUpdateSecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityAccessReview
		localVarHTTPResponse *http.Response
		localAPIResponse     *APIResponse
		err                  error
	)

	if a.client.cfg.Okta.Client.RequestTimeout > 0 {
		localctx, cancel := context.WithTimeout(r.ctx, time.Second*time.Duration(a.client.cfg.Okta.Client.RequestTimeout))
		r.ctx = localctx
		defer cancel()
	}
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecurityAccessReviewsAPIService.UpdateSecurityAccessReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/security-access-reviews/{securityAccessReviewId}"
	localVarPath = strings.Replace(localVarPath, "{"+"securityAccessReviewId"+"}", url.PathEscape(parameterToString(r.securityAccessReviewId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.securityAccessReviewUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("securityAccessReviewUpdateRequest is required and must be specified")
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
	localVarPostBody = r.securityAccessReviewUpdateRequest
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
