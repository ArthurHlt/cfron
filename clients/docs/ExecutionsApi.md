# \ExecutionsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExecutionsByJob**](ExecutionsApi.md#ListExecutionsByJob) | **Get** /jobs/{job_name}/executions | 



## ListExecutionsByJob

> []Execution ListExecutionsByJob(ctx, jobName).Execute()





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
    jobName := "jobName_example" // string | The job that owns the executions to be fetched.

    configuration := clients.NewConfiguration()
    api_client := clients.NewAPIClient(configuration)
    resp, r, err := api_client.ExecutionsApi.ListExecutionsByJob(context.Background(), jobName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ExecutionsApi.ListExecutionsByJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExecutionsByJob`: []Execution
    fmt.Fprintf(os.Stdout, "Response from `ExecutionsApi.ListExecutionsByJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job that owns the executions to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExecutionsByJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

