# \ServiceaccountsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceaccountsApi.md#CreateServiceAccount) | **Post** /apis/service_accounts/v1 | Create service account
[**DeleteServiceAccount**](ServiceaccountsApi.md#DeleteServiceAccount) | **Delete** /apis/service_accounts/v1/{id} | Delete service account by id
[**GetServiceAccount**](ServiceaccountsApi.md#GetServiceAccount) | **Get** /apis/service_accounts/v1/{id} | Get service account by id
[**GetServiceAccounts**](ServiceaccountsApi.md#GetServiceAccounts) | **Get** /apis/service_accounts/v1 | List all service accounts
[**ResetServiceAccountSecret**](ServiceaccountsApi.md#ResetServiceAccountSecret) | **Post** /apis/service_accounts/v1/{id}/resetSecret | Reset service account secret by id
[**UpdateServiceAccount**](ServiceaccountsApi.md#UpdateServiceAccount) | **Patch** /apis/service_accounts/v1/{id} | Update service account



## CreateServiceAccount

> ServiceAccountData CreateServiceAccount(ctx).ServiceAccountCreateRequestData(serviceAccountCreateRequestData).Execute()

Create service account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceAccountCreateRequestData := *openapiclient.NewServiceAccountCreateRequestData("Name_example", "Description_example") // ServiceAccountCreateRequestData | 'name' and 'description' of the service account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceaccountsApi.CreateServiceAccount(context.Background()).ServiceAccountCreateRequestData(serviceAccountCreateRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.CreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccount`: ServiceAccountData
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountCreateRequestData** | [**ServiceAccountCreateRequestData**](ServiceAccountCreateRequestData.md) | &#39;name&#39; and &#39;description&#39; of the service account | 

### Return type

[**ServiceAccountData**](ServiceAccountData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> DeleteServiceAccount(ctx, id).Execute()

Delete service account by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceaccountsApi.DeleteServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.DeleteServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccount

> ServiceAccountData GetServiceAccount(ctx, id).Execute()

Get service account by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceaccountsApi.GetServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.GetServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccount`: ServiceAccountData
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.GetServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccountData**](ServiceAccountData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccounts

> []ServiceAccountData GetServiceAccounts(ctx).First(first).Max(max).Execute()

List all service accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    first := int32(56) // int32 |  (optional) (default to 0)
    max := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceaccountsApi.GetServiceAccounts(context.Background()).First(first).Max(max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.GetServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccounts`: []ServiceAccountData
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.GetServiceAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **first** | **int32** |  | [default to 0]
 **max** | **int32** |  | [default to 20]

### Return type

[**[]ServiceAccountData**](ServiceAccountData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetServiceAccountSecret

> ServiceAccountData ResetServiceAccountSecret(ctx, id).Execute()

Reset service account secret by id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceaccountsApi.ResetServiceAccountSecret(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.ResetServiceAccountSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetServiceAccountSecret`: ServiceAccountData
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.ResetServiceAccountSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetServiceAccountSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccountData**](ServiceAccountData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAccount

> ServiceAccountData UpdateServiceAccount(ctx, id).ServiceAccountRequestData(serviceAccountRequestData).Execute()

Update service account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    serviceAccountRequestData := *openapiclient.NewServiceAccountRequestData() // ServiceAccountRequestData | 'name' and 'description' of the service account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceaccountsApi.UpdateServiceAccount(context.Background(), id).ServiceAccountRequestData(serviceAccountRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceaccountsApi.UpdateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAccount`: ServiceAccountData
    fmt.Fprintf(os.Stdout, "Response from `ServiceaccountsApi.UpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountRequestData** | [**ServiceAccountRequestData**](ServiceAccountRequestData.md) | &#39;name&#39; and &#39;description&#39; of the service account | 

### Return type

[**ServiceAccountData**](ServiceAccountData.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
