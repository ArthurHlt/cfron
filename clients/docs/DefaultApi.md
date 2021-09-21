# \DefaultApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Busy**](DefaultApi.md#Busy) | **Get** /busy | 
[**GetIsLeader**](DefaultApi.md#GetIsLeader) | **Get** /isleader | 
[**GetLeader**](DefaultApi.md#GetLeader) | **Get** /leader | 
[**Leave**](DefaultApi.md#Leave) | **Post** /leave | 
[**Status**](DefaultApi.md#Status) | **Get** / | 



## Busy

> []Execution Busy(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/orange-cloudfoundry/cfron/clients"
)

func main() {

    configuration := clients.NewConfiguration()
    api_client := clients.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Busy(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Busy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Busy`: []Execution
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Busy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBusyRequest struct via the builder pattern


### Return type

[**[]Execution**](execution.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsLeader

> GetIsLeader(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/orange-cloudfoundry/cfron/clients"
)

func main() {

    configuration := clients.NewConfiguration()
    api_client := clients.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIsLeader(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIsLeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIsLeaderRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLeader

> Member GetLeader(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/orange-cloudfoundry/cfron/clients"
)

func main() {

    configuration := clients.NewConfiguration()
    api_client := clients.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetLeader(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLeader``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLeader`: Member
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLeader`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLeaderRequest struct via the builder pattern


### Return type

[**Member**](member.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Leave

> []Member Leave(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/orange-cloudfoundry/cfron/clients"
)

func main() {

    configuration := clients.NewConfiguration()
    api_client := clients.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Leave(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Leave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Leave`: []Member
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Leave`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLeaveRequest struct via the builder pattern


### Return type

[**[]Member**](member.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> Status Status(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/orange-cloudfoundry/cfron/clients"
)

func main() {

    configuration := clients.NewConfiguration()
    api_client := clients.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.Status(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Status``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status`: Status
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Status`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


### Return type

[**Status**](status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

