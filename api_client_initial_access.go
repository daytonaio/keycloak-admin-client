/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keycloak_admin_client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ClientInitialAccessAPIService ClientInitialAccessAPI service
type ClientInitialAccessAPIService service

type ApiAdminRealmsRealmClientsInitialAccessGetRequest struct {
	ctx context.Context
	ApiService *ClientInitialAccessAPIService
	realm string
}

func (r ApiAdminRealmsRealmClientsInitialAccessGetRequest) Execute() ([]ClientInitialAccessPresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsInitialAccessGetExecute(r)
}

/*
AdminRealmsRealmClientsInitialAccessGet Method for AdminRealmsRealmClientsInitialAccessGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @return ApiAdminRealmsRealmClientsInitialAccessGetRequest
*/
func (a *ClientInitialAccessAPIService) AdminRealmsRealmClientsInitialAccessGet(ctx context.Context, realm string) ApiAdminRealmsRealmClientsInitialAccessGetRequest {
	return ApiAdminRealmsRealmClientsInitialAccessGetRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
	}
}

// Execute executes the request
//  @return []ClientInitialAccessPresentation
func (a *ClientInitialAccessAPIService) AdminRealmsRealmClientsInitialAccessGetExecute(r ApiAdminRealmsRealmClientsInitialAccessGetRequest) ([]ClientInitialAccessPresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ClientInitialAccessPresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientInitialAccessAPIService.AdminRealmsRealmClientsInitialAccessGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients-initial-access"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

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

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiAdminRealmsRealmClientsInitialAccessIdDeleteRequest struct {
	ctx context.Context
	ApiService *ClientInitialAccessAPIService
	realm string
	id string
}

func (r ApiAdminRealmsRealmClientsInitialAccessIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsInitialAccessIdDeleteExecute(r)
}

/*
AdminRealmsRealmClientsInitialAccessIdDelete Method for AdminRealmsRealmClientsInitialAccessIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param id
 @return ApiAdminRealmsRealmClientsInitialAccessIdDeleteRequest
*/
func (a *ClientInitialAccessAPIService) AdminRealmsRealmClientsInitialAccessIdDelete(ctx context.Context, realm string, id string) ApiAdminRealmsRealmClientsInitialAccessIdDeleteRequest {
	return ApiAdminRealmsRealmClientsInitialAccessIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		id: id,
	}
}

// Execute executes the request
func (a *ClientInitialAccessAPIService) AdminRealmsRealmClientsInitialAccessIdDeleteExecute(r ApiAdminRealmsRealmClientsInitialAccessIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientInitialAccessAPIService.AdminRealmsRealmClientsInitialAccessIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients-initial-access/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAdminRealmsRealmClientsInitialAccessPostRequest struct {
	ctx context.Context
	ApiService *ClientInitialAccessAPIService
	realm string
	clientInitialAccessCreatePresentation *ClientInitialAccessCreatePresentation
}

func (r ApiAdminRealmsRealmClientsInitialAccessPostRequest) ClientInitialAccessCreatePresentation(clientInitialAccessCreatePresentation ClientInitialAccessCreatePresentation) ApiAdminRealmsRealmClientsInitialAccessPostRequest {
	r.clientInitialAccessCreatePresentation = &clientInitialAccessCreatePresentation
	return r
}

func (r ApiAdminRealmsRealmClientsInitialAccessPostRequest) Execute() (*ClientInitialAccessCreatePresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsInitialAccessPostExecute(r)
}

/*
AdminRealmsRealmClientsInitialAccessPost Create a new initial access token.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @return ApiAdminRealmsRealmClientsInitialAccessPostRequest
*/
func (a *ClientInitialAccessAPIService) AdminRealmsRealmClientsInitialAccessPost(ctx context.Context, realm string) ApiAdminRealmsRealmClientsInitialAccessPostRequest {
	return ApiAdminRealmsRealmClientsInitialAccessPostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
	}
}

// Execute executes the request
//  @return ClientInitialAccessCreatePresentation
func (a *ClientInitialAccessAPIService) AdminRealmsRealmClientsInitialAccessPostExecute(r ApiAdminRealmsRealmClientsInitialAccessPostRequest) (*ClientInitialAccessCreatePresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ClientInitialAccessCreatePresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientInitialAccessAPIService.AdminRealmsRealmClientsInitialAccessPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients-initial-access"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.clientInitialAccessCreatePresentation
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
