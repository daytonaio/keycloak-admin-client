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


// ComponentAPIService ComponentAPI service
type ComponentAPIService service

type ApiAdminRealmsRealmComponentsGetRequest struct {
	ctx context.Context
	ApiService *ComponentAPIService
	realm string
	name *string
	parent *string
	type_ *string
}

func (r ApiAdminRealmsRealmComponentsGetRequest) Name(name string) ApiAdminRealmsRealmComponentsGetRequest {
	r.name = &name
	return r
}

func (r ApiAdminRealmsRealmComponentsGetRequest) Parent(parent string) ApiAdminRealmsRealmComponentsGetRequest {
	r.parent = &parent
	return r
}

func (r ApiAdminRealmsRealmComponentsGetRequest) Type_(type_ string) ApiAdminRealmsRealmComponentsGetRequest {
	r.type_ = &type_
	return r
}

func (r ApiAdminRealmsRealmComponentsGetRequest) Execute() ([]ComponentRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmComponentsGetExecute(r)
}

/*
AdminRealmsRealmComponentsGet Method for AdminRealmsRealmComponentsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @return ApiAdminRealmsRealmComponentsGetRequest
*/
func (a *ComponentAPIService) AdminRealmsRealmComponentsGet(ctx context.Context, realm string) ApiAdminRealmsRealmComponentsGetRequest {
	return ApiAdminRealmsRealmComponentsGetRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
	}
}

// Execute executes the request
//  @return []ComponentRepresentation
func (a *ComponentAPIService) AdminRealmsRealmComponentsGetExecute(r ApiAdminRealmsRealmComponentsGetRequest) ([]ComponentRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ComponentRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentAPIService.AdminRealmsRealmComponentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/components"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.parent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parent", r.parent, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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

type ApiAdminRealmsRealmComponentsIdDeleteRequest struct {
	ctx context.Context
	ApiService *ComponentAPIService
	realm string
	id string
}

func (r ApiAdminRealmsRealmComponentsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AdminRealmsRealmComponentsIdDeleteExecute(r)
}

/*
AdminRealmsRealmComponentsIdDelete Method for AdminRealmsRealmComponentsIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param id
 @return ApiAdminRealmsRealmComponentsIdDeleteRequest
*/
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdDelete(ctx context.Context, realm string, id string) ApiAdminRealmsRealmComponentsIdDeleteRequest {
	return ApiAdminRealmsRealmComponentsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		id: id,
	}
}

// Execute executes the request
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdDeleteExecute(r ApiAdminRealmsRealmComponentsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentAPIService.AdminRealmsRealmComponentsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/components/{id}"
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

type ApiAdminRealmsRealmComponentsIdGetRequest struct {
	ctx context.Context
	ApiService *ComponentAPIService
	realm string
	id string
}

func (r ApiAdminRealmsRealmComponentsIdGetRequest) Execute() (*ComponentRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmComponentsIdGetExecute(r)
}

/*
AdminRealmsRealmComponentsIdGet Method for AdminRealmsRealmComponentsIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param id
 @return ApiAdminRealmsRealmComponentsIdGetRequest
*/
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdGet(ctx context.Context, realm string, id string) ApiAdminRealmsRealmComponentsIdGetRequest {
	return ApiAdminRealmsRealmComponentsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		id: id,
	}
}

// Execute executes the request
//  @return ComponentRepresentation
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdGetExecute(r ApiAdminRealmsRealmComponentsIdGetRequest) (*ComponentRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ComponentRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentAPIService.AdminRealmsRealmComponentsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/components/{id}"
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

type ApiAdminRealmsRealmComponentsIdPutRequest struct {
	ctx context.Context
	ApiService *ComponentAPIService
	realm string
	id string
	componentRepresentation *ComponentRepresentation
}

func (r ApiAdminRealmsRealmComponentsIdPutRequest) ComponentRepresentation(componentRepresentation ComponentRepresentation) ApiAdminRealmsRealmComponentsIdPutRequest {
	r.componentRepresentation = &componentRepresentation
	return r
}

func (r ApiAdminRealmsRealmComponentsIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.AdminRealmsRealmComponentsIdPutExecute(r)
}

/*
AdminRealmsRealmComponentsIdPut Method for AdminRealmsRealmComponentsIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param id
 @return ApiAdminRealmsRealmComponentsIdPutRequest
*/
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdPut(ctx context.Context, realm string, id string) ApiAdminRealmsRealmComponentsIdPutRequest {
	return ApiAdminRealmsRealmComponentsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		id: id,
	}
}

// Execute executes the request
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdPutExecute(r ApiAdminRealmsRealmComponentsIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentAPIService.AdminRealmsRealmComponentsIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/components/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.componentRepresentation
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

type ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest struct {
	ctx context.Context
	ApiService *ComponentAPIService
	realm string
	id string
	type_ *string
}

func (r ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest) Type_(type_ string) ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest {
	r.type_ = &type_
	return r
}

func (r ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest) Execute() ([]ComponentTypeRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmComponentsIdSubComponentTypesGetExecute(r)
}

/*
AdminRealmsRealmComponentsIdSubComponentTypesGet List of subcomponent types that are available to configure for a particular parent component.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param id
 @return ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest
*/
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdSubComponentTypesGet(ctx context.Context, realm string, id string) ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest {
	return ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		id: id,
	}
}

// Execute executes the request
//  @return []ComponentTypeRepresentation
func (a *ComponentAPIService) AdminRealmsRealmComponentsIdSubComponentTypesGetExecute(r ApiAdminRealmsRealmComponentsIdSubComponentTypesGetRequest) ([]ComponentTypeRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ComponentTypeRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentAPIService.AdminRealmsRealmComponentsIdSubComponentTypesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/components/{id}/sub-component-types"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
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

type ApiAdminRealmsRealmComponentsPostRequest struct {
	ctx context.Context
	ApiService *ComponentAPIService
	realm string
	componentRepresentation *ComponentRepresentation
}

func (r ApiAdminRealmsRealmComponentsPostRequest) ComponentRepresentation(componentRepresentation ComponentRepresentation) ApiAdminRealmsRealmComponentsPostRequest {
	r.componentRepresentation = &componentRepresentation
	return r
}

func (r ApiAdminRealmsRealmComponentsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.AdminRealmsRealmComponentsPostExecute(r)
}

/*
AdminRealmsRealmComponentsPost Method for AdminRealmsRealmComponentsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @return ApiAdminRealmsRealmComponentsPostRequest
*/
func (a *ComponentAPIService) AdminRealmsRealmComponentsPost(ctx context.Context, realm string) ApiAdminRealmsRealmComponentsPostRequest {
	return ApiAdminRealmsRealmComponentsPostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
	}
}

// Execute executes the request
func (a *ComponentAPIService) AdminRealmsRealmComponentsPostExecute(r ApiAdminRealmsRealmComponentsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentAPIService.AdminRealmsRealmComponentsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/components"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.componentRepresentation
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
