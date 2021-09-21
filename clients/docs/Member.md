# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Node name | [optional] 
**Addr** | Pointer to **string** | IP Address | [optional] 
**Port** | Pointer to **int32** | Port number | [optional] 
**Tags** | Pointer to **map[string]string** | Tags asociated with this node | [optional] 
**Status** | Pointer to **int32** | The serf status of the node see: https://godoc.org/github.com/hashicorp/serf/serf#MemberStatus | [optional] 
**ProtocolMin** | Pointer to **int32** | Serf protocol minimum version this node can understand or speak | [optional] 
**ProtocolMax** | Pointer to **int32** | Serf protocol maximum version this node can understand or speak | [optional] 
**ProtocolCur** | Pointer to **int32** | Serf protocol current version this node can understand or speak | [optional] 
**DelegateMin** | Pointer to **int32** | Serf delegate protocol minimum version this node can understand or speak | [optional] 
**DelegateMax** | Pointer to **int32** | Serf delegate protocol maximum version this node can understand or speak | [optional] 
**DelegateCur** | Pointer to **int32** | Serf delegate protocol current version this node can understand or speak | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Member) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddr

`func (o *Member) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *Member) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *Member) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *Member) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetPort

`func (o *Member) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Member) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Member) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Member) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTags

`func (o *Member) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Member) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Member) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Member) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetStatus

`func (o *Member) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Member) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Member) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Member) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProtocolMin

`func (o *Member) GetProtocolMin() int32`

GetProtocolMin returns the ProtocolMin field if non-nil, zero value otherwise.

### GetProtocolMinOk

`func (o *Member) GetProtocolMinOk() (*int32, bool)`

GetProtocolMinOk returns a tuple with the ProtocolMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMin

`func (o *Member) SetProtocolMin(v int32)`

SetProtocolMin sets ProtocolMin field to given value.

### HasProtocolMin

`func (o *Member) HasProtocolMin() bool`

HasProtocolMin returns a boolean if a field has been set.

### GetProtocolMax

`func (o *Member) GetProtocolMax() int32`

GetProtocolMax returns the ProtocolMax field if non-nil, zero value otherwise.

### GetProtocolMaxOk

`func (o *Member) GetProtocolMaxOk() (*int32, bool)`

GetProtocolMaxOk returns a tuple with the ProtocolMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolMax

`func (o *Member) SetProtocolMax(v int32)`

SetProtocolMax sets ProtocolMax field to given value.

### HasProtocolMax

`func (o *Member) HasProtocolMax() bool`

HasProtocolMax returns a boolean if a field has been set.

### GetProtocolCur

`func (o *Member) GetProtocolCur() int32`

GetProtocolCur returns the ProtocolCur field if non-nil, zero value otherwise.

### GetProtocolCurOk

`func (o *Member) GetProtocolCurOk() (*int32, bool)`

GetProtocolCurOk returns a tuple with the ProtocolCur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolCur

`func (o *Member) SetProtocolCur(v int32)`

SetProtocolCur sets ProtocolCur field to given value.

### HasProtocolCur

`func (o *Member) HasProtocolCur() bool`

HasProtocolCur returns a boolean if a field has been set.

### GetDelegateMin

`func (o *Member) GetDelegateMin() int32`

GetDelegateMin returns the DelegateMin field if non-nil, zero value otherwise.

### GetDelegateMinOk

`func (o *Member) GetDelegateMinOk() (*int32, bool)`

GetDelegateMinOk returns a tuple with the DelegateMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateMin

`func (o *Member) SetDelegateMin(v int32)`

SetDelegateMin sets DelegateMin field to given value.

### HasDelegateMin

`func (o *Member) HasDelegateMin() bool`

HasDelegateMin returns a boolean if a field has been set.

### GetDelegateMax

`func (o *Member) GetDelegateMax() int32`

GetDelegateMax returns the DelegateMax field if non-nil, zero value otherwise.

### GetDelegateMaxOk

`func (o *Member) GetDelegateMaxOk() (*int32, bool)`

GetDelegateMaxOk returns a tuple with the DelegateMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateMax

`func (o *Member) SetDelegateMax(v int32)`

SetDelegateMax sets DelegateMax field to given value.

### HasDelegateMax

`func (o *Member) HasDelegateMax() bool`

HasDelegateMax returns a boolean if a field has been set.

### GetDelegateCur

`func (o *Member) GetDelegateCur() int32`

GetDelegateCur returns the DelegateCur field if non-nil, zero value otherwise.

### GetDelegateCurOk

`func (o *Member) GetDelegateCurOk() (*int32, bool)`

GetDelegateCurOk returns a tuple with the DelegateCur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegateCur

`func (o *Member) SetDelegateCur(v int32)`

SetDelegateCur sets DelegateCur field to given value.

### HasDelegateCur

`func (o *Member) HasDelegateCur() bool`

HasDelegateCur returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


