# \MembersApi

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMember**](MembersApi.md#GetMember) | **Get** /members | 



## GetMember

> []Member GetMember(ctx).Execute()





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
    resp, r, err := api_client.MembersApi.GetMember(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `MembersApi.GetMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMember`: []Member
    fmt.Fprintf(os.Stdout, "Response from `MembersApi.GetMember`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberRequest struct via the builder pattern


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

