# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobName** | Pointer to **string** | job name | [optional] 
**StartedAt** | Pointer to **time.Time** | start time of the execution | [optional] 
**FinishedAt** | Pointer to **time.Time** | when the execution finished running | [optional] 
**Success** | Pointer to **bool** | the execution run successfuly | [optional] 
**Output** | Pointer to **string** | partial output of the command execution | [optional] 
**NodeName** | Pointer to **string** | name of the node that executed the command | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobName

`func (o *Execution) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *Execution) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *Execution) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *Execution) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetStartedAt

`func (o *Execution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Execution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Execution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Execution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *Execution) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Execution) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Execution) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Execution) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetSuccess

`func (o *Execution) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Execution) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Execution) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Execution) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetOutput

`func (o *Execution) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Execution) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Execution) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Execution) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetNodeName

`func (o *Execution) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *Execution) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *Execution) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *Execution) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


