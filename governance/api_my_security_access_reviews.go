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

type MySecurityAccessReviewsAPI interface {

	/*
		AddMySecurityAccessReviewComment Add a comment for a security access review

		Adds a comment for a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiAddMySecurityAccessReviewCommentRequest
	*/
	AddMySecurityAccessReviewComment(ctx context.Context, securityAccessReviewId string) ApiAddMySecurityAccessReviewCommentRequest

	// AddMySecurityAccessReviewCommentExecute executes the request
	AddMySecurityAccessReviewCommentExecute(r ApiAddMySecurityAccessReviewCommentRequest) (*APIResponse, error)

	/*
		ExecuteMySecurityAccessReviewAccessesAction Execute an action on an access item

		Executes an action on an access or sub-access item in a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
		@return ApiExecuteMySecurityAccessReviewAccessesActionRequest
	*/
	ExecuteMySecurityAccessReviewAccessesAction(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiExecuteMySecurityAccessReviewAccessesActionRequest

	// ExecuteMySecurityAccessReviewAccessesActionExecute executes the request
	ExecuteMySecurityAccessReviewAccessesActionExecute(r ApiExecuteMySecurityAccessReviewAccessesActionRequest) (*APIResponse, error)

	/*
		ExecuteMySecurityAccessReviewAction Execute an action on a security access review

		Executes a specified action on a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiExecuteMySecurityAccessReviewActionRequest
	*/
	ExecuteMySecurityAccessReviewAction(ctx context.Context, securityAccessReviewId string) ApiExecuteMySecurityAccessReviewActionRequest

	// ExecuteMySecurityAccessReviewActionExecute executes the request
	ExecuteMySecurityAccessReviewActionExecute(r ApiExecuteMySecurityAccessReviewActionRequest) (*APIResponse, error)

	/*
			GenerateMySecurityAccessReviewAccessesSummary Generate a summary for an access item

			Generates a summary for an access item in a security access review

		> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
			@return ApiGenerateMySecurityAccessReviewAccessesSummaryRequest
	*/
	GenerateMySecurityAccessReviewAccessesSummary(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiGenerateMySecurityAccessReviewAccessesSummaryRequest

	// GenerateMySecurityAccessReviewAccessesSummaryExecute executes the request
	//  @return AiMessage
	GenerateMySecurityAccessReviewAccessesSummaryExecute(r ApiGenerateMySecurityAccessReviewAccessesSummaryRequest) (*AiMessage, *APIResponse, error)

	/*
			GenerateMySecurityAccessReviewSummary Generate a summary for a security access review

			Generates a summary for a security access review

		> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@return ApiGenerateMySecurityAccessReviewSummaryRequest
	*/
	GenerateMySecurityAccessReviewSummary(ctx context.Context, securityAccessReviewId string) ApiGenerateMySecurityAccessReviewSummaryRequest

	// GenerateMySecurityAccessReviewSummaryExecute executes the request
	//  @return AiMessage
	GenerateMySecurityAccessReviewSummaryExecute(r ApiGenerateMySecurityAccessReviewSummaryRequest) (*AiMessage, *APIResponse, error)

	/*
		GetMySecurityAccessReview Retrieve a security access review

		Retrieves a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiGetMySecurityAccessReviewRequest
	*/
	GetMySecurityAccessReview(ctx context.Context, securityAccessReviewId string) ApiGetMySecurityAccessReviewRequest

	// GetMySecurityAccessReviewExecute executes the request
	//  @return SecurityAccessReview
	GetMySecurityAccessReviewExecute(r ApiGetMySecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error)

	/*
		GetMySecurityAccessReviewPrincipalDetails Retrieve the principal for a security access review

		Retrieves the details of a security access review's principal target

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiGetMySecurityAccessReviewPrincipalDetailsRequest
	*/
	GetMySecurityAccessReviewPrincipalDetails(ctx context.Context, securityAccessReviewId string) ApiGetMySecurityAccessReviewPrincipalDetailsRequest

	// GetMySecurityAccessReviewPrincipalDetailsExecute executes the request
	//  @return SecurityAccessReviewPrincipal
	GetMySecurityAccessReviewPrincipalDetailsExecute(r ApiGetMySecurityAccessReviewPrincipalDetailsRequest) (*SecurityAccessReviewPrincipal, *APIResponse, error)

	/*
		GetMySecurityAccessReviewsStats Retrieve the statistics for security access reviews

		Retrieves the statistics for security access reviews

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetMySecurityAccessReviewsStatsRequest
	*/
	GetMySecurityAccessReviewsStats(ctx context.Context) ApiGetMySecurityAccessReviewsStatsRequest

	// GetMySecurityAccessReviewsStatsExecute executes the request
	//  @return SecurityAccessReviewsStats
	GetMySecurityAccessReviewsStatsExecute(r ApiGetMySecurityAccessReviewsStatsRequest) (*SecurityAccessReviewsStats, *APIResponse, error)

	/*
			ListMySecurityAccessReviewAccesses List the access items for a security access review

			Lists the access items for a specific security access review.

		Access items refer to the top-level resources that the security access review's target principal has access to.
		For example, a top-level resource can be an app, hence the access item describes the principal's access to that app.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@return ApiListMySecurityAccessReviewAccessesRequest
	*/
	ListMySecurityAccessReviewAccesses(ctx context.Context, securityAccessReviewId string) ApiListMySecurityAccessReviewAccessesRequest

	// ListMySecurityAccessReviewAccessesExecute executes the request
	//  @return SecurityAccessReviewAccessItems
	ListMySecurityAccessReviewAccessesExecute(r ApiListMySecurityAccessReviewAccessesRequest) (*SecurityAccessReviewAccessItems, *APIResponse, error)

	/*
		ListMySecurityAccessReviewAccessesAnomalies List the anomalies for an access item

		Lists the anomalies for an access item in a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
		@return ApiListMySecurityAccessReviewAccessesAnomaliesRequest
	*/
	ListMySecurityAccessReviewAccessesAnomalies(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiListMySecurityAccessReviewAccessesAnomaliesRequest

	// ListMySecurityAccessReviewAccessesAnomaliesExecute executes the request
	//  @return SecurityAccessReviewAnomalies
	ListMySecurityAccessReviewAccessesAnomaliesExecute(r ApiListMySecurityAccessReviewAccessesAnomaliesRequest) (*SecurityAccessReviewAnomalies, *APIResponse, error)

	/*
		ListMySecurityAccessReviewActions List all actions for a security access review

		Lists all of the actions available for a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiListMySecurityAccessReviewActionsRequest
	*/
	ListMySecurityAccessReviewActions(ctx context.Context, securityAccessReviewId string) ApiListMySecurityAccessReviewActionsRequest

	// ListMySecurityAccessReviewActionsExecute executes the request
	//  @return SecurityAccessReviewActions
	ListMySecurityAccessReviewActionsExecute(r ApiListMySecurityAccessReviewActionsRequest) (*SecurityAccessReviewActions, *APIResponse, error)

	/*
		ListMySecurityAccessReviewHistory List the history of a security access review

		Lists the history of actions and changes for a security access review

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param securityAccessReviewId The ID of the security access review
		@return ApiListMySecurityAccessReviewHistoryRequest
	*/
	ListMySecurityAccessReviewHistory(ctx context.Context, securityAccessReviewId string) ApiListMySecurityAccessReviewHistoryRequest

	// ListMySecurityAccessReviewHistoryExecute executes the request
	//  @return SecurityAccessReviewHistoryItems
	ListMySecurityAccessReviewHistoryExecute(r ApiListMySecurityAccessReviewHistoryRequest) (*SecurityAccessReviewHistoryItems, *APIResponse, error)

	/*
			ListMySecurityAccessReviewSubAccesses List the sub-access items for an access item

			Lists the sub-access items for an access item from a security access review.

		A sub-access item refers to the access of a resource that is a part of a top-level resource in an access item.
		For example, an access item can describe the access of app, and the sub-access items can describe the access of groups, entitlement values, or entitlement bundles that belong to the app.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param securityAccessReviewId The ID of the security access review
			@param securityAccessReviewAccessId The ID of the access item in a security access review
			@return ApiListMySecurityAccessReviewSubAccessesRequest
	*/
	ListMySecurityAccessReviewSubAccesses(ctx context.Context, securityAccessReviewId string, securityAccessReviewAccessId string) ApiListMySecurityAccessReviewSubAccessesRequest

	// ListMySecurityAccessReviewSubAccessesExecute executes the request
	//  @return SecurityAccessReviewSubAccessItems
	ListMySecurityAccessReviewSubAccessesExecute(r ApiListMySecurityAccessReviewSubAccessesRequest) (*SecurityAccessReviewSubAccessItems, *APIResponse, error)

	/*
		ListMySecurityAccessReviews List the security access reviews

		Lists the security access reviews

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiListMySecurityAccessReviewsRequest
	*/
	ListMySecurityAccessReviews(ctx context.Context) ApiListMySecurityAccessReviewsRequest

	// ListMySecurityAccessReviewsExecute executes the request
	//  @return SecurityAccessReviews
	ListMySecurityAccessReviewsExecute(r ApiListMySecurityAccessReviewsRequest) (*SecurityAccessReviews, *APIResponse, error)
}

// MySecurityAccessReviewsAPIService MySecurityAccessReviewsAPI service
type MySecurityAccessReviewsAPIService service

type ApiAddMySecurityAccessReviewCommentRequest struct {
	ctx                                context.Context
	ApiService                         MySecurityAccessReviewsAPI
	securityAccessReviewId             string
	securityAccessReviewCommentRequest *SecurityAccessReviewCommentRequest
	retryCount                         int32
}

func (r ApiAddMySecurityAccessReviewCommentRequest) SecurityAccessReviewCommentRequest(securityAccessReviewCommentRequest SecurityAccessReviewCommentRequest) ApiAddMySecurityAccessReviewCommentRequest {
	r.securityAccessReviewCommentRequest = &securityAccessReviewCommentRequest
	return r
}

func (r ApiAddMySecurityAccessReviewCommentRequest) Execute() (*APIResponse, error) {
	return r.ApiService.AddMySecurityAccessReviewCommentExecute(r)
}

/*
AddMySecurityAccessReviewComment Add a comment for a security access review

Adds a comment for a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiAddMySecurityAccessReviewCommentRequest
*/
func (a *MySecurityAccessReviewsAPIService) AddMySecurityAccessReviewComment(ctx context.Context, securityAccessReviewId string) ApiAddMySecurityAccessReviewCommentRequest {
	return ApiAddMySecurityAccessReviewCommentRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
func (a *MySecurityAccessReviewsAPIService) AddMySecurityAccessReviewCommentExecute(r ApiAddMySecurityAccessReviewCommentRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.AddMySecurityAccessReviewComment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/comment"
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

type ApiExecuteMySecurityAccessReviewAccessesActionRequest struct {
	ctx                                       context.Context
	ApiService                                MySecurityAccessReviewsAPI
	securityAccessReviewId                    string
	securityAccessReviewTargetId              string
	securityAccessReviewAccessesActionRequest *SecurityAccessReviewAccessesActionRequest
	retryCount                                int32
}

func (r ApiExecuteMySecurityAccessReviewAccessesActionRequest) SecurityAccessReviewAccessesActionRequest(securityAccessReviewAccessesActionRequest SecurityAccessReviewAccessesActionRequest) ApiExecuteMySecurityAccessReviewAccessesActionRequest {
	r.securityAccessReviewAccessesActionRequest = &securityAccessReviewAccessesActionRequest
	return r
}

func (r ApiExecuteMySecurityAccessReviewAccessesActionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ExecuteMySecurityAccessReviewAccessesActionExecute(r)
}

/*
ExecuteMySecurityAccessReviewAccessesAction Execute an action on an access item

Executes an action on an access or sub-access item in a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
	@return ApiExecuteMySecurityAccessReviewAccessesActionRequest
*/
func (a *MySecurityAccessReviewsAPIService) ExecuteMySecurityAccessReviewAccessesAction(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiExecuteMySecurityAccessReviewAccessesActionRequest {
	return ApiExecuteMySecurityAccessReviewAccessesActionRequest{
		ApiService:                   a,
		ctx:                          ctx,
		securityAccessReviewId:       securityAccessReviewId,
		securityAccessReviewTargetId: securityAccessReviewTargetId,
		retryCount:                   0,
	}
}

// Execute executes the request
func (a *MySecurityAccessReviewsAPIService) ExecuteMySecurityAccessReviewAccessesActionExecute(r ApiExecuteMySecurityAccessReviewAccessesActionRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ExecuteMySecurityAccessReviewAccessesAction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/actions"
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

type ApiExecuteMySecurityAccessReviewActionRequest struct {
	ctx                               context.Context
	ApiService                        MySecurityAccessReviewsAPI
	securityAccessReviewId            string
	securityAccessReviewActionRequest *SecurityAccessReviewActionRequest
	retryCount                        int32
}

func (r ApiExecuteMySecurityAccessReviewActionRequest) SecurityAccessReviewActionRequest(securityAccessReviewActionRequest SecurityAccessReviewActionRequest) ApiExecuteMySecurityAccessReviewActionRequest {
	r.securityAccessReviewActionRequest = &securityAccessReviewActionRequest
	return r
}

func (r ApiExecuteMySecurityAccessReviewActionRequest) Execute() (*APIResponse, error) {
	return r.ApiService.ExecuteMySecurityAccessReviewActionExecute(r)
}

/*
ExecuteMySecurityAccessReviewAction Execute an action on a security access review

Executes a specified action on a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiExecuteMySecurityAccessReviewActionRequest
*/
func (a *MySecurityAccessReviewsAPIService) ExecuteMySecurityAccessReviewAction(ctx context.Context, securityAccessReviewId string) ApiExecuteMySecurityAccessReviewActionRequest {
	return ApiExecuteMySecurityAccessReviewActionRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
func (a *MySecurityAccessReviewsAPIService) ExecuteMySecurityAccessReviewActionExecute(r ApiExecuteMySecurityAccessReviewActionRequest) (*APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ExecuteMySecurityAccessReviewAction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/actions"
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

type ApiGenerateMySecurityAccessReviewAccessesSummaryRequest struct {
	ctx                          context.Context
	ApiService                   MySecurityAccessReviewsAPI
	securityAccessReviewId       string
	securityAccessReviewTargetId string
	retryCount                   int32
}

func (r ApiGenerateMySecurityAccessReviewAccessesSummaryRequest) Execute() (*AiMessage, *APIResponse, error) {
	return r.ApiService.GenerateMySecurityAccessReviewAccessesSummaryExecute(r)
}

/*
GenerateMySecurityAccessReviewAccessesSummary Generate a summary for an access item

# Generates a summary for an access item in a security access review

> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
	@return ApiGenerateMySecurityAccessReviewAccessesSummaryRequest
*/
func (a *MySecurityAccessReviewsAPIService) GenerateMySecurityAccessReviewAccessesSummary(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiGenerateMySecurityAccessReviewAccessesSummaryRequest {
	return ApiGenerateMySecurityAccessReviewAccessesSummaryRequest{
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
func (a *MySecurityAccessReviewsAPIService) GenerateMySecurityAccessReviewAccessesSummaryExecute(r ApiGenerateMySecurityAccessReviewAccessesSummaryRequest) (*AiMessage, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.GenerateMySecurityAccessReviewAccessesSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/summary"
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

type ApiGenerateMySecurityAccessReviewSummaryRequest struct {
	ctx                    context.Context
	ApiService             MySecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiGenerateMySecurityAccessReviewSummaryRequest) Execute() (*AiMessage, *APIResponse, error) {
	return r.ApiService.GenerateMySecurityAccessReviewSummaryExecute(r)
}

/*
GenerateMySecurityAccessReviewSummary Generate a summary for a security access review

# Generates a summary for a security access review

> **Note:** The [`governanceAI.securityAccessReview.enabled`](https://developer.okta.com/docs/api/iga/openapi/governance.api/tag/Org-Governance-Settings/#tag/Org-Governance-Settings/operation/updateOrgSettings!path=governanceAI&t=request) org governance setting must be enabled for users to generate summaries for security access reviews.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiGenerateMySecurityAccessReviewSummaryRequest
*/
func (a *MySecurityAccessReviewsAPIService) GenerateMySecurityAccessReviewSummary(ctx context.Context, securityAccessReviewId string) ApiGenerateMySecurityAccessReviewSummaryRequest {
	return ApiGenerateMySecurityAccessReviewSummaryRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return AiMessage
func (a *MySecurityAccessReviewsAPIService) GenerateMySecurityAccessReviewSummaryExecute(r ApiGenerateMySecurityAccessReviewSummaryRequest) (*AiMessage, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.GenerateMySecurityAccessReviewSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/summary"
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

type ApiGetMySecurityAccessReviewRequest struct {
	ctx                    context.Context
	ApiService             MySecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiGetMySecurityAccessReviewRequest) Execute() (*SecurityAccessReview, *APIResponse, error) {
	return r.ApiService.GetMySecurityAccessReviewExecute(r)
}

/*
GetMySecurityAccessReview Retrieve a security access review

Retrieves a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiGetMySecurityAccessReviewRequest
*/
func (a *MySecurityAccessReviewsAPIService) GetMySecurityAccessReview(ctx context.Context, securityAccessReviewId string) ApiGetMySecurityAccessReviewRequest {
	return ApiGetMySecurityAccessReviewRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReview
func (a *MySecurityAccessReviewsAPIService) GetMySecurityAccessReviewExecute(r ApiGetMySecurityAccessReviewRequest) (*SecurityAccessReview, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.GetMySecurityAccessReview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}"
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

type ApiGetMySecurityAccessReviewPrincipalDetailsRequest struct {
	ctx                    context.Context
	ApiService             MySecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiGetMySecurityAccessReviewPrincipalDetailsRequest) Execute() (*SecurityAccessReviewPrincipal, *APIResponse, error) {
	return r.ApiService.GetMySecurityAccessReviewPrincipalDetailsExecute(r)
}

/*
GetMySecurityAccessReviewPrincipalDetails Retrieve the principal for a security access review

Retrieves the details of a security access review's principal target

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiGetMySecurityAccessReviewPrincipalDetailsRequest
*/
func (a *MySecurityAccessReviewsAPIService) GetMySecurityAccessReviewPrincipalDetails(ctx context.Context, securityAccessReviewId string) ApiGetMySecurityAccessReviewPrincipalDetailsRequest {
	return ApiGetMySecurityAccessReviewPrincipalDetailsRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewPrincipal
func (a *MySecurityAccessReviewsAPIService) GetMySecurityAccessReviewPrincipalDetailsExecute(r ApiGetMySecurityAccessReviewPrincipalDetailsRequest) (*SecurityAccessReviewPrincipal, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.GetMySecurityAccessReviewPrincipalDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/principal"
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

type ApiGetMySecurityAccessReviewsStatsRequest struct {
	ctx        context.Context
	ApiService MySecurityAccessReviewsAPI
	retryCount int32
}

func (r ApiGetMySecurityAccessReviewsStatsRequest) Execute() (*SecurityAccessReviewsStats, *APIResponse, error) {
	return r.ApiService.GetMySecurityAccessReviewsStatsExecute(r)
}

/*
GetMySecurityAccessReviewsStats Retrieve the statistics for security access reviews

Retrieves the statistics for security access reviews

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMySecurityAccessReviewsStatsRequest
*/
func (a *MySecurityAccessReviewsAPIService) GetMySecurityAccessReviewsStats(ctx context.Context) ApiGetMySecurityAccessReviewsStatsRequest {
	return ApiGetMySecurityAccessReviewsStatsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewsStats
func (a *MySecurityAccessReviewsAPIService) GetMySecurityAccessReviewsStatsExecute(r ApiGetMySecurityAccessReviewsStatsRequest) (*SecurityAccessReviewsStats, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.GetMySecurityAccessReviewsStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/stats"

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

type ApiListMySecurityAccessReviewAccessesRequest struct {
	ctx                    context.Context
	ApiService             MySecurityAccessReviewsAPI
	securityAccessReviewId string
	filter                 *string
	orderBy                *[]string
	after                  *string
	limit                  *int32
	retryCount             int32
}

// A [filter](https://developer.okta.com/docs/api/#filter) expression that filters access items.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListMySecurityAccessReviewAccessesRequest) Filter(filter string) ApiListMySecurityAccessReviewAccessesRequest {
	r.filter = &filter
	return r
}

// The field to sort the results, in ascending (asc) or descending (desc) order. Sorting is applied to only one field.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListMySecurityAccessReviewAccessesRequest) OrderBy(orderBy []string) ApiListMySecurityAccessReviewAccessesRequest {
	r.orderBy = &orderBy
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListMySecurityAccessReviewAccessesRequest) After(after string) ApiListMySecurityAccessReviewAccessesRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListMySecurityAccessReviewAccessesRequest) Limit(limit int32) ApiListMySecurityAccessReviewAccessesRequest {
	r.limit = &limit
	return r
}

func (r ApiListMySecurityAccessReviewAccessesRequest) Execute() (*SecurityAccessReviewAccessItems, *APIResponse, error) {
	return r.ApiService.ListMySecurityAccessReviewAccessesExecute(r)
}

/*
ListMySecurityAccessReviewAccesses List the access items for a security access review

Lists the access items for a specific security access review.

Access items refer to the top-level resources that the security access review's target principal has access to.
For example, a top-level resource can be an app, hence the access item describes the principal's access to that app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiListMySecurityAccessReviewAccessesRequest
*/
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewAccesses(ctx context.Context, securityAccessReviewId string) ApiListMySecurityAccessReviewAccessesRequest {
	return ApiListMySecurityAccessReviewAccessesRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewAccessItems
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewAccessesExecute(r ApiListMySecurityAccessReviewAccessesRequest) (*SecurityAccessReviewAccessItems, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ListMySecurityAccessReviewAccesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses"
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

type ApiListMySecurityAccessReviewAccessesAnomaliesRequest struct {
	ctx                          context.Context
	ApiService                   MySecurityAccessReviewsAPI
	securityAccessReviewId       string
	securityAccessReviewTargetId string
	retryCount                   int32
}

func (r ApiListMySecurityAccessReviewAccessesAnomaliesRequest) Execute() (*SecurityAccessReviewAnomalies, *APIResponse, error) {
	return r.ApiService.ListMySecurityAccessReviewAccessesAnomaliesExecute(r)
}

/*
ListMySecurityAccessReviewAccessesAnomalies List the anomalies for an access item

Lists the anomalies for an access item in a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewTargetId The ID of the access or sub-access item in a security access review
	@return ApiListMySecurityAccessReviewAccessesAnomaliesRequest
*/
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewAccessesAnomalies(ctx context.Context, securityAccessReviewId string, securityAccessReviewTargetId string) ApiListMySecurityAccessReviewAccessesAnomaliesRequest {
	return ApiListMySecurityAccessReviewAccessesAnomaliesRequest{
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
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewAccessesAnomaliesExecute(r ApiListMySecurityAccessReviewAccessesAnomaliesRequest) (*SecurityAccessReviewAnomalies, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ListMySecurityAccessReviewAccessesAnomalies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/anomalies"
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

type ApiListMySecurityAccessReviewActionsRequest struct {
	ctx                    context.Context
	ApiService             MySecurityAccessReviewsAPI
	securityAccessReviewId string
	retryCount             int32
}

func (r ApiListMySecurityAccessReviewActionsRequest) Execute() (*SecurityAccessReviewActions, *APIResponse, error) {
	return r.ApiService.ListMySecurityAccessReviewActionsExecute(r)
}

/*
ListMySecurityAccessReviewActions List all actions for a security access review

Lists all of the actions available for a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiListMySecurityAccessReviewActionsRequest
*/
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewActions(ctx context.Context, securityAccessReviewId string) ApiListMySecurityAccessReviewActionsRequest {
	return ApiListMySecurityAccessReviewActionsRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewActions
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewActionsExecute(r ApiListMySecurityAccessReviewActionsRequest) (*SecurityAccessReviewActions, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ListMySecurityAccessReviewActions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/actions"
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

type ApiListMySecurityAccessReviewHistoryRequest struct {
	ctx                    context.Context
	ApiService             MySecurityAccessReviewsAPI
	securityAccessReviewId string
	after                  *string
	limit                  *int32
	retryCount             int32
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListMySecurityAccessReviewHistoryRequest) After(after string) ApiListMySecurityAccessReviewHistoryRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListMySecurityAccessReviewHistoryRequest) Limit(limit int32) ApiListMySecurityAccessReviewHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiListMySecurityAccessReviewHistoryRequest) Execute() (*SecurityAccessReviewHistoryItems, *APIResponse, error) {
	return r.ApiService.ListMySecurityAccessReviewHistoryExecute(r)
}

/*
ListMySecurityAccessReviewHistory List the history of a security access review

Lists the history of actions and changes for a security access review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@return ApiListMySecurityAccessReviewHistoryRequest
*/
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewHistory(ctx context.Context, securityAccessReviewId string) ApiListMySecurityAccessReviewHistoryRequest {
	return ApiListMySecurityAccessReviewHistoryRequest{
		ApiService:             a,
		ctx:                    ctx,
		securityAccessReviewId: securityAccessReviewId,
		retryCount:             0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviewHistoryItems
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewHistoryExecute(r ApiListMySecurityAccessReviewHistoryRequest) (*SecurityAccessReviewHistoryItems, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ListMySecurityAccessReviewHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/history"
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

type ApiListMySecurityAccessReviewSubAccessesRequest struct {
	ctx                          context.Context
	ApiService                   MySecurityAccessReviewsAPI
	securityAccessReviewId       string
	securityAccessReviewAccessId string
	filter                       *string
	orderBy                      *[]string
	after                        *string
	limit                        *int32
	retryCount                   int32
}

// A [filter](https://developer.okta.com/docs/api/#filter) expression that filters sub-access items.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListMySecurityAccessReviewSubAccessesRequest) Filter(filter string) ApiListMySecurityAccessReviewSubAccessesRequest {
	r.filter = &filter
	return r
}

// A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListMySecurityAccessReviewSubAccessesRequest) OrderBy(orderBy []string) ApiListMySecurityAccessReviewSubAccessesRequest {
	r.orderBy = &orderBy
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListMySecurityAccessReviewSubAccessesRequest) After(after string) ApiListMySecurityAccessReviewSubAccessesRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListMySecurityAccessReviewSubAccessesRequest) Limit(limit int32) ApiListMySecurityAccessReviewSubAccessesRequest {
	r.limit = &limit
	return r
}

func (r ApiListMySecurityAccessReviewSubAccessesRequest) Execute() (*SecurityAccessReviewSubAccessItems, *APIResponse, error) {
	return r.ApiService.ListMySecurityAccessReviewSubAccessesExecute(r)
}

/*
ListMySecurityAccessReviewSubAccesses List the sub-access items for an access item

Lists the sub-access items for an access item from a security access review.

A sub-access item refers to the access of a resource that is a part of a top-level resource in an access item.
For example, an access item can describe the access of app, and the sub-access items can describe the access of groups, entitlement values, or entitlement bundles that belong to the app.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param securityAccessReviewId The ID of the security access review
	@param securityAccessReviewAccessId The ID of the access item in a security access review
	@return ApiListMySecurityAccessReviewSubAccessesRequest
*/
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewSubAccesses(ctx context.Context, securityAccessReviewId string, securityAccessReviewAccessId string) ApiListMySecurityAccessReviewSubAccessesRequest {
	return ApiListMySecurityAccessReviewSubAccessesRequest{
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
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewSubAccessesExecute(r ApiListMySecurityAccessReviewSubAccessesRequest) (*SecurityAccessReviewSubAccessItems, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ListMySecurityAccessReviewSubAccesses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewAccessId}/sub-accesses"
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

type ApiListMySecurityAccessReviewsRequest struct {
	ctx        context.Context
	ApiService MySecurityAccessReviewsAPI
	filter     *string
	orderBy    *[]string
	after      *string
	limit      *int32
	retryCount int32
}

// A [filter](https://developer.okta.com/docs/api/#filter) expression that filters security access reviews. The &#x60;eq&#x60; and &#x60;co&#x60; [operators](https://developer.okta.com/docs/api/#operators) are supported for string properties. The &#x60;gt&#x60; and &#x60;lt&#x60; operators are supported for date properties.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).
func (r ApiListMySecurityAccessReviewsRequest) Filter(filter string) ApiListMySecurityAccessReviewsRequest {
	r.filter = &filter
	return r
}

// The field to sort the results, in ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;) order. Sorting is applied to only one field.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)
func (r ApiListMySecurityAccessReviewsRequest) OrderBy(orderBy []string) ApiListMySecurityAccessReviewsRequest {
	r.orderBy = &orderBy
	return r
}

// The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request.
func (r ApiListMySecurityAccessReviewsRequest) After(after string) ApiListMySecurityAccessReviewsRequest {
	r.after = &after
	return r
}

// The maximum number of records returned in a response
func (r ApiListMySecurityAccessReviewsRequest) Limit(limit int32) ApiListMySecurityAccessReviewsRequest {
	r.limit = &limit
	return r
}

func (r ApiListMySecurityAccessReviewsRequest) Execute() (*SecurityAccessReviews, *APIResponse, error) {
	return r.ApiService.ListMySecurityAccessReviewsExecute(r)
}

/*
ListMySecurityAccessReviews List the security access reviews

Lists the security access reviews

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListMySecurityAccessReviewsRequest
*/
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviews(ctx context.Context) ApiListMySecurityAccessReviewsRequest {
	return ApiListMySecurityAccessReviewsRequest{
		ApiService: a,
		ctx:        ctx,
		retryCount: 0,
	}
}

// Execute executes the request
//
//	@return SecurityAccessReviews
func (a *MySecurityAccessReviewsAPIService) ListMySecurityAccessReviewsExecute(r ApiListMySecurityAccessReviewsRequest) (*SecurityAccessReviews, *APIResponse, error) {
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
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySecurityAccessReviewsAPIService.ListMySecurityAccessReviews")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/governance/api/v2/my/security-access-reviews"

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
