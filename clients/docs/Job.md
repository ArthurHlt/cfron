# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for the job. Use only lower case letters (unicode), digits, underscore and dash. | 
**Displayname** | Pointer to **string** | Nice name for the job. Optional. | [optional] 
**Schedule** | **string** | Cron expression for the job. | 
**Timezone** | Pointer to **string** | Timezone where the job will be executed. By default and when field is set to empty string, the job will run in local time. | [optional] 
**Owner** | Pointer to **string** | Owner of the job | [optional] 
**OwnerEmail** | Pointer to **string** | Email of the owner | [optional] 
**SuccessCount** | Pointer to **int32** | Number of successful executions | [optional] [readonly] 
**ErrorCount** | Pointer to **int32** | Number of failed executions | [optional] [readonly] 
**LastSuccess** | Pointer to **time.Time** | Last time this job executed successfully | [optional] [readonly] 
**LastError** | Pointer to **time.Time** | Last time this job failed | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Disabled state of the job | [optional] 
**Tags** | Pointer to **map[string]string** | Target nodes tags of this job | [optional] 
**Metadata** | Pointer to **map[string]string** | Extra metadata tags for this job | [optional] 
**Retries** | Pointer to **int32** | Number of times to retry a failed job execution | [optional] 
**ParentJob** | Pointer to **string** | The name/id of the job that will trigger the execution of this job | [optional] 
**DependentJobs** | Pointer to **[]string** | Array containing the jobs that depends on this one | [optional] [readonly] 
**Processors** | Pointer to **map[string]map[string]string** | Processor plugins used to process executions results of this job | [optional] 
**Concurrency** | Pointer to **string** | Concurrency policy for the job allow/forbid | [optional] 
**Executor** | Pointer to **string** | Executor plugin used to run the job | [optional] 
**ExecutorConfig** | Pointer to **map[string]string** | Executor plugin parameters | [optional] 
**Status** | Pointer to **string** | Status of the job | [optional] [readonly] 

## Methods

### NewJob

`func NewJob(name string, schedule string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayname

`func (o *Job) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *Job) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *Job) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.

### HasDisplayname

`func (o *Job) HasDisplayname() bool`

HasDisplayname returns a boolean if a field has been set.

### GetSchedule

`func (o *Job) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Job) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Job) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetTimezone

`func (o *Job) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Job) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Job) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Job) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetOwner

`func (o *Job) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Job) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Job) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Job) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Job) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Job) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Job) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *Job) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetSuccessCount

`func (o *Job) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *Job) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *Job) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *Job) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetErrorCount

`func (o *Job) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *Job) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *Job) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *Job) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetLastSuccess

`func (o *Job) GetLastSuccess() time.Time`

GetLastSuccess returns the LastSuccess field if non-nil, zero value otherwise.

### GetLastSuccessOk

`func (o *Job) GetLastSuccessOk() (*time.Time, bool)`

GetLastSuccessOk returns a tuple with the LastSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccess

`func (o *Job) SetLastSuccess(v time.Time)`

SetLastSuccess sets LastSuccess field to given value.

### HasLastSuccess

`func (o *Job) HasLastSuccess() bool`

HasLastSuccess returns a boolean if a field has been set.

### GetLastError

`func (o *Job) GetLastError() time.Time`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *Job) GetLastErrorOk() (*time.Time, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *Job) SetLastError(v time.Time)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *Job) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### GetDisabled

`func (o *Job) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Job) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Job) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Job) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTags

`func (o *Job) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Job) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Job) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Job) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *Job) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Job) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Job) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Job) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRetries

`func (o *Job) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Job) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Job) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *Job) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetParentJob

`func (o *Job) GetParentJob() string`

GetParentJob returns the ParentJob field if non-nil, zero value otherwise.

### GetParentJobOk

`func (o *Job) GetParentJobOk() (*string, bool)`

GetParentJobOk returns a tuple with the ParentJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentJob

`func (o *Job) SetParentJob(v string)`

SetParentJob sets ParentJob field to given value.

### HasParentJob

`func (o *Job) HasParentJob() bool`

HasParentJob returns a boolean if a field has been set.

### GetDependentJobs

`func (o *Job) GetDependentJobs() []string`

GetDependentJobs returns the DependentJobs field if non-nil, zero value otherwise.

### GetDependentJobsOk

`func (o *Job) GetDependentJobsOk() (*[]string, bool)`

GetDependentJobsOk returns a tuple with the DependentJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentJobs

`func (o *Job) SetDependentJobs(v []string)`

SetDependentJobs sets DependentJobs field to given value.

### HasDependentJobs

`func (o *Job) HasDependentJobs() bool`

HasDependentJobs returns a boolean if a field has been set.

### GetProcessors

`func (o *Job) GetProcessors() map[string]map[string]string`

GetProcessors returns the Processors field if non-nil, zero value otherwise.

### GetProcessorsOk

`func (o *Job) GetProcessorsOk() (*map[string]map[string]string, bool)`

GetProcessorsOk returns a tuple with the Processors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessors

`func (o *Job) SetProcessors(v map[string]map[string]string)`

SetProcessors sets Processors field to given value.

### HasProcessors

`func (o *Job) HasProcessors() bool`

HasProcessors returns a boolean if a field has been set.

### GetConcurrency

`func (o *Job) GetConcurrency() string`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *Job) GetConcurrencyOk() (*string, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *Job) SetConcurrency(v string)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *Job) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetExecutor

`func (o *Job) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *Job) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *Job) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *Job) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetExecutorConfig

`func (o *Job) GetExecutorConfig() map[string]string`

GetExecutorConfig returns the ExecutorConfig field if non-nil, zero value otherwise.

### GetExecutorConfigOk

`func (o *Job) GetExecutorConfigOk() (*map[string]string, bool)`

GetExecutorConfigOk returns a tuple with the ExecutorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutorConfig

`func (o *Job) SetExecutorConfig(v map[string]string)`

SetExecutorConfig sets ExecutorConfig field to given value.

### HasExecutorConfig

`func (o *Job) HasExecutorConfig() bool`

HasExecutorConfig returns a boolean if a field has been set.

### GetStatus

`func (o *Job) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Job) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


