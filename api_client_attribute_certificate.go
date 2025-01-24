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
	"os"
)


// ClientAttributeCertificateAPIService ClientAttributeCertificateAPI service
type ClientAttributeCertificateAPIService service

type ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest struct {
	ctx context.Context
	ApiService *ClientAttributeCertificateAPIService
	realm string
	clientUuid string
	attr string
	keyStoreConfig *KeyStoreConfig
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest) KeyStoreConfig(keyStoreConfig KeyStoreConfig) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest {
	r.keyStoreConfig = &keyStoreConfig
	return r
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostExecute(r)
}

/*
AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost Get a keystore file for the client, containing private key and public certificate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param clientUuid id of client (not client-id!)
 @param attr
 @return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest
*/
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost(ctx context.Context, realm string, clientUuid string, attr string) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest {
	return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		clientUuid: clientUuid,
		attr: attr,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostExecute(r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAttributeCertificateAPIService.AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client-uuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attr"+"}", url.PathEscape(parameterValueToString(r.attr, "attr")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.keyStoreConfig
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

type ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest struct {
	ctx context.Context
	ApiService *ClientAttributeCertificateAPIService
	realm string
	clientUuid string
	attr string
	keyStoreConfig *KeyStoreConfig
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest) KeyStoreConfig(keyStoreConfig KeyStoreConfig) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest {
	r.keyStoreConfig = &keyStoreConfig
	return r
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostExecute(r)
}

/*
AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost Generate a new keypair and certificate, and get the private key file  Generates a keypair and certificate and serves the private key in a specified keystore format. Only generated public certificate is saved in Keycloak DB - the private key is not.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param clientUuid id of client (not client-id!)
 @param attr
 @return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest
*/
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost(ctx context.Context, realm string, clientUuid string, attr string) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest {
	return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		clientUuid: clientUuid,
		attr: attr,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostExecute(r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAttributeCertificateAPIService.AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/generate-and-download"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client-uuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attr"+"}", url.PathEscape(parameterValueToString(r.attr, "attr")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.keyStoreConfig
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

type ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest struct {
	ctx context.Context
	ApiService *ClientAttributeCertificateAPIService
	realm string
	clientUuid string
	attr string
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest) Execute() (*CertificateRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostExecute(r)
}

/*
AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost Generate a new certificate with new key pair

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param clientUuid id of client (not client-id!)
 @param attr
 @return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest
*/
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost(ctx context.Context, realm string, clientUuid string, attr string) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest {
	return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		clientUuid: clientUuid,
		attr: attr,
	}
}

// Execute executes the request
//  @return CertificateRepresentation
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostExecute(r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest) (*CertificateRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAttributeCertificateAPIService.AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/generate"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client-uuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attr"+"}", url.PathEscape(parameterValueToString(r.attr, "attr")), -1)

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

type ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest struct {
	ctx context.Context
	ApiService *ClientAttributeCertificateAPIService
	realm string
	clientUuid string
	attr string
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest) Execute() (*CertificateRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsClientUuidCertificatesAttrGetExecute(r)
}

/*
AdminRealmsRealmClientsClientUuidCertificatesAttrGet Get key info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param clientUuid id of client (not client-id!)
 @param attr
 @return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest
*/
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrGet(ctx context.Context, realm string, clientUuid string, attr string) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest {
	return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		clientUuid: clientUuid,
		attr: attr,
	}
}

// Execute executes the request
//  @return CertificateRepresentation
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrGetExecute(r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest) (*CertificateRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAttributeCertificateAPIService.AdminRealmsRealmClientsClientUuidCertificatesAttrGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client-uuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attr"+"}", url.PathEscape(parameterValueToString(r.attr, "attr")), -1)

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

type ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest struct {
	ctx context.Context
	ApiService *ClientAttributeCertificateAPIService
	realm string
	clientUuid string
	attr string
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest) Execute() (*CertificateRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostExecute(r)
}

/*
AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost Upload only certificate, not private key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param clientUuid id of client (not client-id!)
 @param attr
 @return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest
*/
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost(ctx context.Context, realm string, clientUuid string, attr string) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest {
	return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		clientUuid: clientUuid,
		attr: attr,
	}
}

// Execute executes the request
//  @return CertificateRepresentation
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostExecute(r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest) (*CertificateRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAttributeCertificateAPIService.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/upload-certificate"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client-uuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attr"+"}", url.PathEscape(parameterValueToString(r.attr, "attr")), -1)

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

type ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest struct {
	ctx context.Context
	ApiService *ClientAttributeCertificateAPIService
	realm string
	clientUuid string
	attr string
}

func (r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest) Execute() (*CertificateRepresentation, *http.Response, error) {
	return r.ApiService.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostExecute(r)
}

/*
AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost Upload certificate and eventually private key

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param realm realm name (not id!)
 @param clientUuid id of client (not client-id!)
 @param attr
 @return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest
*/
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost(ctx context.Context, realm string, clientUuid string, attr string) ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest {
	return ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest{
		ApiService: a,
		ctx: ctx,
		realm: realm,
		clientUuid: clientUuid,
		attr: attr,
	}
}

// Execute executes the request
//  @return CertificateRepresentation
func (a *ClientAttributeCertificateAPIService) AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostExecute(r ApiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest) (*CertificateRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CertificateRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientAttributeCertificateAPIService.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/upload"
	localVarPath = strings.Replace(localVarPath, "{"+"realm"+"}", url.PathEscape(parameterValueToString(r.realm, "realm")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"client-uuid"+"}", url.PathEscape(parameterValueToString(r.clientUuid, "clientUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attr"+"}", url.PathEscape(parameterValueToString(r.attr, "attr")), -1)

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
