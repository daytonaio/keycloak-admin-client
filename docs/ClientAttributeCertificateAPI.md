# \ClientAttributeCertificateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost**](ClientAttributeCertificateAPI.md#AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/download | Get a keystore file for the client, containing private key and public certificate
[**AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost**](ClientAttributeCertificateAPI.md#AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/generate-and-download | Generate a new keypair and certificate, and get the private key file  Generates a keypair and certificate and serves the private key in a specified keystore format. Only generated public certificate is saved in Keycloak DB - the private key is not.
[**AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost**](ClientAttributeCertificateAPI.md#AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/generate | Generate a new certificate with new key pair
[**AdminRealmsRealmClientsClientUuidCertificatesAttrGet**](ClientAttributeCertificateAPI.md#AdminRealmsRealmClientsClientUuidCertificatesAttrGet) | **Get** /admin/realms/{realm}/clients/{client-uuid}/certificates/{attr} | Get key info
[**AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost**](ClientAttributeCertificateAPI.md#AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/upload-certificate | Upload only certificate, not private key
[**AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost**](ClientAttributeCertificateAPI.md#AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost) | **Post** /admin/realms/{realm}/clients/{client-uuid}/certificates/{attr}/upload | Upload certificate and eventually private key



## AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost

> *os.File AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost(ctx, realm, clientUuid, attr).KeyStoreConfig(keyStoreConfig).Execute()

Get a keystore file for the client, containing private key and public certificate

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	attr := "attr_example" // string | 
	keyStoreConfig := *openapiclient.NewKeyStoreConfig() // KeyStoreConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost(context.Background(), realm, clientUuid, attr).KeyStoreConfig(keyStoreConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidCertificatesAttrDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **keyStoreConfig** | [**KeyStoreConfig**](KeyStoreConfig.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost

> *os.File AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost(ctx, realm, clientUuid, attr).KeyStoreConfig(keyStoreConfig).Execute()

Generate a new keypair and certificate, and get the private key file  Generates a keypair and certificate and serves the private key in a specified keystore format. Only generated public certificate is saved in Keycloak DB - the private key is not.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	attr := "attr_example" // string | 
	keyStoreConfig := *openapiclient.NewKeyStoreConfig() // KeyStoreConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost(context.Background(), realm, clientUuid, attr).KeyStoreConfig(keyStoreConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidCertificatesAttrGenerateAndDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **keyStoreConfig** | [**KeyStoreConfig**](KeyStoreConfig.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost

> CertificateRepresentation AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost(ctx, realm, clientUuid, attr).Execute()

Generate a new certificate with new key pair

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost(context.Background(), realm, clientUuid, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidCertificatesAttrGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidCertificatesAttrGet

> CertificateRepresentation AdminRealmsRealmClientsClientUuidCertificatesAttrGet(ctx, realm, clientUuid, attr).Execute()

Get key info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGet(context.Background(), realm, clientUuid, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidCertificatesAttrGet`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidCertificatesAttrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost

> CertificateRepresentation AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost(ctx, realm, clientUuid, attr).Execute()

Upload only certificate, not private key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost(context.Background(), realm, clientUuid, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadCertificatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost

> CertificateRepresentation AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost(ctx, realm, clientUuid, attr).Execute()

Upload certificate and eventually private key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/daytonaio/keycloak-admin-client"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	clientUuid := "clientUuid_example" // string | id of client (not client-id!)
	attr := "attr_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost(context.Background(), realm, clientUuid, attr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost`: CertificateRepresentation
	fmt.Fprintf(os.Stdout, "Response from `ClientAttributeCertificateAPI.AdminRealmsRealmClientsClientUuidCertificatesAttrUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**clientUuid** | **string** | id of client (not client-id!) | 
**attr** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmClientsClientUuidCertificatesAttrUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CertificateRepresentation**](CertificateRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

