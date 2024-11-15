# \MimirApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Mimir**](MimirApi.md#Mimir) | **Get** /mayachain/mimir | 
[**MimirAdmin**](MimirApi.md#MimirAdmin) | **Get** /mayachain/mimir/admin | 
[**MimirKey**](MimirApi.md#MimirKey) | **Get** /mayachain/mimir/key/{key} | 
[**MimirNode**](MimirApi.md#MimirNode) | **Get** /mayachain/mimir/node/{address} | 
[**MimirNodes**](MimirApi.md#MimirNodes) | **Get** /mayachain/mimir/nodes_all | 



## Mimir

> map[string]string Mimir(ctx).Height(height).Execute()





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
    height := int64(789) // int64 | optional block height, defaults to current tip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MimirApi.Mimir(context.Background()).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MimirApi.Mimir``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Mimir`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `MimirApi.Mimir`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMimirRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **int64** | optional block height, defaults to current tip | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MimirAdmin

> map[string]string MimirAdmin(ctx).Height(height).Execute()





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
    height := int64(789) // int64 | optional block height, defaults to current tip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MimirApi.MimirAdmin(context.Background()).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MimirApi.MimirAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MimirAdmin`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `MimirApi.MimirAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMimirAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **int64** | optional block height, defaults to current tip | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MimirKey

> int64 MimirKey(ctx, key).Height(height).Execute()





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
    key := "key_example" // string | the mimir key to lookup
    height := int64(789) // int64 | optional block height, defaults to current tip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MimirApi.MimirKey(context.Background(), key).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MimirApi.MimirKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MimirKey`: int64
    fmt.Fprintf(os.Stdout, "Response from `MimirApi.MimirKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | the mimir key to lookup | 

### Other Parameters

Other parameters are passed through a pointer to a apiMimirKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **height** | **int64** | optional block height, defaults to current tip | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MimirNode

> map[string]string MimirNode(ctx, address).Height(height).Execute()





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
    address := "maya1zupk5lmc84r2dh738a9g3zscavannjy3nzplwt" // string | 
    height := int64(789) // int64 | optional block height, defaults to current tip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MimirApi.MimirNode(context.Background(), address).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MimirApi.MimirNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MimirNode`: map[string]string
    fmt.Fprintf(os.Stdout, "Response from `MimirApi.MimirNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMimirNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **height** | **int64** | optional block height, defaults to current tip | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MimirNodes

> MimirNodesResponse MimirNodes(ctx).Height(height).Execute()





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
    height := int64(789) // int64 | optional block height, defaults to current tip (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MimirApi.MimirNodes(context.Background()).Height(height).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MimirApi.MimirNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MimirNodes`: MimirNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `MimirApi.MimirNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMimirNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **height** | **int64** | optional block height, defaults to current tip | 

### Return type

[**MimirNodesResponse**](MimirNodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

