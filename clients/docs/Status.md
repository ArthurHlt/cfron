# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **map[string]interface{}** | Node basic details | [optional] [readonly] 
**Serf** | Pointer to **map[string]interface{}** | Serf status | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Tags asociated with this node | [optional] [readonly] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *Status) GetAgent() map[string]interface{}`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Status) GetAgentOk() (*map[string]interface{}, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Status) SetAgent(v map[string]interface{})`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *Status) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetSerf

`func (o *Status) GetSerf() map[string]interface{}`

GetSerf returns the Serf field if non-nil, zero value otherwise.

### GetSerfOk

`func (o *Status) GetSerfOk() (*map[string]interface{}, bool)`

GetSerfOk returns a tuple with the Serf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerf

`func (o *Status) SetSerf(v map[string]interface{})`

SetSerf sets Serf field to given value.

### HasSerf

`func (o *Status) HasSerf() bool`

HasSerf returns a boolean if a field has been set.

### GetTags

`func (o *Status) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Status) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Status) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Status) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


