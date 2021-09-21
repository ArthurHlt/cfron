# Go API client for openapi

You can communicate with Dkron using a RESTful JSON API over HTTP. Dkron nodes usually listen on port `8080` for API requests. All examples in this section assume that you've found a running leader at `localhost:8080`.

Dkron implements a RESTful JSON API over HTTP to communicate with software clients. Dkron listens in port `8080` by default. All examples in this section assume that you're using the default port.

Default API responses are unformatted JSON add the `pretty=true` param to format the response.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**Busy**](docs/DefaultApi.md#busy) | **Get** /busy | 
*DefaultApi* | [**GetIsLeader**](docs/DefaultApi.md#getisleader) | **Get** /isleader | 
*DefaultApi* | [**GetLeader**](docs/DefaultApi.md#getleader) | **Get** /leader | 
*DefaultApi* | [**Leave**](docs/DefaultApi.md#leave) | **Post** /leave | 
*DefaultApi* | [**Status**](docs/DefaultApi.md#status) | **Get** / | 
*ExecutionsApi* | [**ListExecutionsByJob**](docs/ExecutionsApi.md#listexecutionsbyjob) | **Get** /jobs/{job_name}/executions | 
*JobsApi* | [**CreateOrUpdateJob**](docs/JobsApi.md#createorupdatejob) | **Post** /jobs | 
*JobsApi* | [**DeleteJob**](docs/JobsApi.md#deletejob) | **Delete** /jobs/{job_name} | 
*JobsApi* | [**GetJobs**](docs/JobsApi.md#getjobs) | **Get** /jobs | 
*JobsApi* | [**Restore**](docs/JobsApi.md#restore) | **Post** /restore | 
*JobsApi* | [**RunJob**](docs/JobsApi.md#runjob) | **Post** /jobs/{job_name} | 
*JobsApi* | [**ShowJobByName**](docs/JobsApi.md#showjobbyname) | **Get** /jobs/{job_name} | 
*JobsApi* | [**ToggleJob**](docs/JobsApi.md#togglejob) | **Post** /jobs/{job_name}/toggle | 
*MembersApi* | [**GetMember**](docs/MembersApi.md#getmember) | **Get** /members | 


## Documentation For Models

 - [Execution](docs/Execution.md)
 - [Job](docs/Job.md)
 - [Member](docs/Member.md)
 - [Status](docs/Status.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


