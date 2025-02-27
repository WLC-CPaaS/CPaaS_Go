# ServiceCallQueueOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentWrapupTime** | Pointer to **int32** |  | [optional] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**ForceAwayOnReject** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**QueueRouter** | Pointer to **string** |  | [optional] 
**QueueType** | Pointer to **string** |  | [optional] 
**RingTimeout** | Pointer to **int32** |  | [optional] 
**TickTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceCallQueueOutputFull

`func NewServiceCallQueueOutputFull() *ServiceCallQueueOutputFull`

NewServiceCallQueueOutputFull instantiates a new ServiceCallQueueOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallQueueOutputFullWithDefaults

`func NewServiceCallQueueOutputFullWithDefaults() *ServiceCallQueueOutputFull`

NewServiceCallQueueOutputFullWithDefaults instantiates a new ServiceCallQueueOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentWrapupTime

`func (o *ServiceCallQueueOutputFull) GetAgentWrapupTime() int32`

GetAgentWrapupTime returns the AgentWrapupTime field if non-nil, zero value otherwise.

### GetAgentWrapupTimeOk

`func (o *ServiceCallQueueOutputFull) GetAgentWrapupTimeOk() (*int32, bool)`

GetAgentWrapupTimeOk returns a tuple with the AgentWrapupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentWrapupTime

`func (o *ServiceCallQueueOutputFull) SetAgentWrapupTime(v int32)`

SetAgentWrapupTime sets AgentWrapupTime field to given value.

### HasAgentWrapupTime

`func (o *ServiceCallQueueOutputFull) HasAgentWrapupTime() bool`

HasAgentWrapupTime returns a boolean if a field has been set.

### GetFeatures

`func (o *ServiceCallQueueOutputFull) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceCallQueueOutputFull) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceCallQueueOutputFull) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ServiceCallQueueOutputFull) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetForceAwayOnReject

`func (o *ServiceCallQueueOutputFull) GetForceAwayOnReject() bool`

GetForceAwayOnReject returns the ForceAwayOnReject field if non-nil, zero value otherwise.

### GetForceAwayOnRejectOk

`func (o *ServiceCallQueueOutputFull) GetForceAwayOnRejectOk() (*bool, bool)`

GetForceAwayOnRejectOk returns a tuple with the ForceAwayOnReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAwayOnReject

`func (o *ServiceCallQueueOutputFull) SetForceAwayOnReject(v bool)`

SetForceAwayOnReject sets ForceAwayOnReject field to given value.

### HasForceAwayOnReject

`func (o *ServiceCallQueueOutputFull) HasForceAwayOnReject() bool`

HasForceAwayOnReject returns a boolean if a field has been set.

### GetId

`func (o *ServiceCallQueueOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceCallQueueOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceCallQueueOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceCallQueueOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceCallQueueOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceCallQueueOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceCallQueueOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceCallQueueOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueueRouter

`func (o *ServiceCallQueueOutputFull) GetQueueRouter() string`

GetQueueRouter returns the QueueRouter field if non-nil, zero value otherwise.

### GetQueueRouterOk

`func (o *ServiceCallQueueOutputFull) GetQueueRouterOk() (*string, bool)`

GetQueueRouterOk returns a tuple with the QueueRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRouter

`func (o *ServiceCallQueueOutputFull) SetQueueRouter(v string)`

SetQueueRouter sets QueueRouter field to given value.

### HasQueueRouter

`func (o *ServiceCallQueueOutputFull) HasQueueRouter() bool`

HasQueueRouter returns a boolean if a field has been set.

### GetQueueType

`func (o *ServiceCallQueueOutputFull) GetQueueType() string`

GetQueueType returns the QueueType field if non-nil, zero value otherwise.

### GetQueueTypeOk

`func (o *ServiceCallQueueOutputFull) GetQueueTypeOk() (*string, bool)`

GetQueueTypeOk returns a tuple with the QueueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueType

`func (o *ServiceCallQueueOutputFull) SetQueueType(v string)`

SetQueueType sets QueueType field to given value.

### HasQueueType

`func (o *ServiceCallQueueOutputFull) HasQueueType() bool`

HasQueueType returns a boolean if a field has been set.

### GetRingTimeout

`func (o *ServiceCallQueueOutputFull) GetRingTimeout() int32`

GetRingTimeout returns the RingTimeout field if non-nil, zero value otherwise.

### GetRingTimeoutOk

`func (o *ServiceCallQueueOutputFull) GetRingTimeoutOk() (*int32, bool)`

GetRingTimeoutOk returns a tuple with the RingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingTimeout

`func (o *ServiceCallQueueOutputFull) SetRingTimeout(v int32)`

SetRingTimeout sets RingTimeout field to given value.

### HasRingTimeout

`func (o *ServiceCallQueueOutputFull) HasRingTimeout() bool`

HasRingTimeout returns a boolean if a field has been set.

### GetTickTime

`func (o *ServiceCallQueueOutputFull) GetTickTime() int32`

GetTickTime returns the TickTime field if non-nil, zero value otherwise.

### GetTickTimeOk

`func (o *ServiceCallQueueOutputFull) GetTickTimeOk() (*int32, bool)`

GetTickTimeOk returns a tuple with the TickTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickTime

`func (o *ServiceCallQueueOutputFull) SetTickTime(v int32)`

SetTickTime sets TickTime field to given value.

### HasTickTime

`func (o *ServiceCallQueueOutputFull) HasTickTime() bool`

HasTickTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


