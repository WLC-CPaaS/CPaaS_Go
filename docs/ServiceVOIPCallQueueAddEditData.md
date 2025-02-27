# ServiceVOIPCallQueueAddEditData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentWrapupTime** | Pointer to **int32** |  | [optional] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**ForceAwayOnReject** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**QueueRouter** | Pointer to **string** |  | [optional] 
**QueueType** | Pointer to **string** |  | [optional] 
**RingTimeout** | Pointer to **int32** |  | [optional] 
**TickTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceVOIPCallQueueAddEditData

`func NewServiceVOIPCallQueueAddEditData(name string, ) *ServiceVOIPCallQueueAddEditData`

NewServiceVOIPCallQueueAddEditData instantiates a new ServiceVOIPCallQueueAddEditData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPCallQueueAddEditDataWithDefaults

`func NewServiceVOIPCallQueueAddEditDataWithDefaults() *ServiceVOIPCallQueueAddEditData`

NewServiceVOIPCallQueueAddEditDataWithDefaults instantiates a new ServiceVOIPCallQueueAddEditData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentWrapupTime

`func (o *ServiceVOIPCallQueueAddEditData) GetAgentWrapupTime() int32`

GetAgentWrapupTime returns the AgentWrapupTime field if non-nil, zero value otherwise.

### GetAgentWrapupTimeOk

`func (o *ServiceVOIPCallQueueAddEditData) GetAgentWrapupTimeOk() (*int32, bool)`

GetAgentWrapupTimeOk returns a tuple with the AgentWrapupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentWrapupTime

`func (o *ServiceVOIPCallQueueAddEditData) SetAgentWrapupTime(v int32)`

SetAgentWrapupTime sets AgentWrapupTime field to given value.

### HasAgentWrapupTime

`func (o *ServiceVOIPCallQueueAddEditData) HasAgentWrapupTime() bool`

HasAgentWrapupTime returns a boolean if a field has been set.

### GetFeatures

`func (o *ServiceVOIPCallQueueAddEditData) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceVOIPCallQueueAddEditData) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceVOIPCallQueueAddEditData) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ServiceVOIPCallQueueAddEditData) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetForceAwayOnReject

`func (o *ServiceVOIPCallQueueAddEditData) GetForceAwayOnReject() bool`

GetForceAwayOnReject returns the ForceAwayOnReject field if non-nil, zero value otherwise.

### GetForceAwayOnRejectOk

`func (o *ServiceVOIPCallQueueAddEditData) GetForceAwayOnRejectOk() (*bool, bool)`

GetForceAwayOnRejectOk returns a tuple with the ForceAwayOnReject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAwayOnReject

`func (o *ServiceVOIPCallQueueAddEditData) SetForceAwayOnReject(v bool)`

SetForceAwayOnReject sets ForceAwayOnReject field to given value.

### HasForceAwayOnReject

`func (o *ServiceVOIPCallQueueAddEditData) HasForceAwayOnReject() bool`

HasForceAwayOnReject returns a boolean if a field has been set.

### GetName

`func (o *ServiceVOIPCallQueueAddEditData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPCallQueueAddEditData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPCallQueueAddEditData) SetName(v string)`

SetName sets Name field to given value.


### GetQueueRouter

`func (o *ServiceVOIPCallQueueAddEditData) GetQueueRouter() string`

GetQueueRouter returns the QueueRouter field if non-nil, zero value otherwise.

### GetQueueRouterOk

`func (o *ServiceVOIPCallQueueAddEditData) GetQueueRouterOk() (*string, bool)`

GetQueueRouterOk returns a tuple with the QueueRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRouter

`func (o *ServiceVOIPCallQueueAddEditData) SetQueueRouter(v string)`

SetQueueRouter sets QueueRouter field to given value.

### HasQueueRouter

`func (o *ServiceVOIPCallQueueAddEditData) HasQueueRouter() bool`

HasQueueRouter returns a boolean if a field has been set.

### GetQueueType

`func (o *ServiceVOIPCallQueueAddEditData) GetQueueType() string`

GetQueueType returns the QueueType field if non-nil, zero value otherwise.

### GetQueueTypeOk

`func (o *ServiceVOIPCallQueueAddEditData) GetQueueTypeOk() (*string, bool)`

GetQueueTypeOk returns a tuple with the QueueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueType

`func (o *ServiceVOIPCallQueueAddEditData) SetQueueType(v string)`

SetQueueType sets QueueType field to given value.

### HasQueueType

`func (o *ServiceVOIPCallQueueAddEditData) HasQueueType() bool`

HasQueueType returns a boolean if a field has been set.

### GetRingTimeout

`func (o *ServiceVOIPCallQueueAddEditData) GetRingTimeout() int32`

GetRingTimeout returns the RingTimeout field if non-nil, zero value otherwise.

### GetRingTimeoutOk

`func (o *ServiceVOIPCallQueueAddEditData) GetRingTimeoutOk() (*int32, bool)`

GetRingTimeoutOk returns a tuple with the RingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingTimeout

`func (o *ServiceVOIPCallQueueAddEditData) SetRingTimeout(v int32)`

SetRingTimeout sets RingTimeout field to given value.

### HasRingTimeout

`func (o *ServiceVOIPCallQueueAddEditData) HasRingTimeout() bool`

HasRingTimeout returns a boolean if a field has been set.

### GetTickTime

`func (o *ServiceVOIPCallQueueAddEditData) GetTickTime() int32`

GetTickTime returns the TickTime field if non-nil, zero value otherwise.

### GetTickTimeOk

`func (o *ServiceVOIPCallQueueAddEditData) GetTickTimeOk() (*int32, bool)`

GetTickTimeOk returns a tuple with the TickTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickTime

`func (o *ServiceVOIPCallQueueAddEditData) SetTickTime(v int32)`

SetTickTime sets TickTime field to given value.

### HasTickTime

`func (o *ServiceVOIPCallQueueAddEditData) HasTickTime() bool`

HasTickTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


