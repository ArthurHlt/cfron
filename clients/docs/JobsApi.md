# \JobsApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateJob**](JobsApi.md#CreateOrUpdateJob) | **Post** /jobs |
[**DeleteJob**](JobsApi.md#DeleteJob) | **Delete** /jobs/{job_name} |
[**GetJobs**](JobsApi.md#GetJobs) | **Get** /jobs |
[**Restore**](JobsApi.md#Restore) | **Post** /restore |
[**RunJob**](JobsApi.md#RunJob) | **Post** /jobs/{job_name} |
[**ShowJobByName**](JobsApi.md#ShowJobByName) | **Get** /jobs/{job_name} |
[**ToggleJob**](JobsApi.md#ToggleJob) | **Post** /jobs/{job_name}/toggle |

## CreateOrUpdateJob

> Job CreateOrUpdateJob(ctx).Body(body).Runoncreate(runoncreate).Execute()

### Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/orange-cloudfoundry/cfron/clients"
	"os"
)

func main() {
	body := *clients.NewJob("job1", "@every 10s") // Job | Updated job object
	runoncreate := true                           // bool | If present, regardless of any value, causes the job to be run immediately after being succesfully created or updated. (optional)

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.CreateOrUpdateJob(context.Background()).Body(body).Runoncreate(runoncreate).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.CreateOrUpdateJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.CreateOrUpdateJob`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateJobRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Job**](Job.md) | Updated job object |
**runoncreate** | **bool** | If present, regardless of any value, causes the job to be run immediately after being succesfully created or updated. |

### Return type

[**Job**](job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteJob

> Job DeleteJob(ctx, jobName).Execute()

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
	jobName := "jobName_example" // string | The job that needs to be deleted.

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.DeleteJob(context.Background(), jobName).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.DeleteJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.DeleteJob`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job that needs to be deleted. |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Job**](job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetJobs

> []Job GetJobs(ctx).Metadata(metadata).Sort(sort).Order(order).Q(q).Start(start).End(end).Execute()

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
	metadata := []string{"Inner_example"} // []string | Filter jobs by metadata (optional)
	sort := "sort_example"                // string | Sorting field (optional)
	order := "order_example"              // string | Sort order (ASC/DESC) (optional)
	q := "q_example"                      // string | Filter query text (optional)
	start := int32(56)                    // int32 | Start index (optional)
	end := int32(56)                      // int32 | End index (optional)

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.GetJobs(context.Background()).Metadata(metadata).Sort(sort).Order(order).Q(q).Start(start).End(end).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobs`: []Job
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobs`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**metadata** | **[]string** | Filter jobs by metadata |
**sort** | **string** | Sorting field |
**order** | **string** | Sort order (ASC/DESC) |
**q** | **string** | Filter query text |
**start** | **int32** | Start index |
**end** | **int32** | End index |

### Return type

[**[]Job**](job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Restore

> []string Restore(ctx).File(file).Execute()

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
	file := os.NewFile(1234, "some_file") // *os.File | Json file that needs to be restored.

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.Restore(context.Background()).File(file).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.Restore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Restore`: []string
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.Restore`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**file** | ***os.File** | Json file that needs to be restored. |

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RunJob

> Job RunJob(ctx, jobName).Execute()

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
	jobName := "jobName_example" // string | The job that needs to be run.

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.RunJob(context.Background(), jobName).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.RunJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.RunJob`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job that needs to be run. |

### Other Parameters

Other parameters are passed through a pointer to a apiRunJobRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Job**](job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ShowJobByName

> Job ShowJobByName(ctx, jobName).Execute()

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
	jobName := "jobName_example" // string | The job that needs to be fetched.

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.ShowJobByName(context.Background(), jobName).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ShowJobByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowJobByName`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.ShowJobByName`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job that needs to be fetched. |

### Other Parameters

Other parameters are passed through a pointer to a apiShowJobByNameRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Job**](job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ToggleJob

> Job ToggleJob(ctx, jobName).Execute()

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
	jobName := "jobName_example" // string | The job that needs to be toggled.

	configuration := clients.NewConfiguration()
	api_client := clients.NewAPIClient(configuration)
	resp, r, err := api_client.JobsApi.ToggleJob(context.Background(), jobName).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ToggleJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleJob`: Job
	fmt.Fprintf(os.Stdout, "Response from `JobsApi.ToggleJob`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job that needs to be toggled. |

### Other Parameters

Other parameters are passed through a pointer to a apiToggleJobRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Job**](job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

