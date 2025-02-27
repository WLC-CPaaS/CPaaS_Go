# ServiceGroupOutputShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to **int32** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Flags** | Pointer to **[]string** |  | [optional] 
**GroupEndpoints** | Pointer to [**map[string]ServiceEndpoint**](ServiceEndpoint.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceGroupOutputShort

`func NewServiceGroupOutputShort() *ServiceGroupOutputShort`

NewServiceGroupOutputShort instantiates a new ServiceGroupOutputShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupOutputShortWithDefaults

`func NewServiceGroupOutputShortWithDefaults() *ServiceGroupOutputShort`

NewServiceGroupOutputShortWithDefaults instantiates a new ServiceGroupOutputShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServiceGroupOutputShort) GetEndpoints() int32`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServiceGroupOutputShort) GetEndpointsOk() (*int32, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServiceGroupOutputShort) SetEndpoints(v int32)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ServiceGroupOutputShort) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetFeatures

`func (o *ServiceGroupOutputShort) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ServiceGroupOutputShort) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ServiceGroupOutputShort) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ServiceGroupOutputShort) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFlags

`func (o *ServiceGroupOutputShort) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *ServiceGroupOutputShort) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *ServiceGroupOutputShort) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *ServiceGroupOutputShort) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetGroupEndpoints

`func (o *ServiceGroupOutputShort) GetGroupEndpoints() map[string]ServiceEndpoint`

GetGroupEndpoints returns the GroupEndpoints field if non-nil, zero value otherwise.

### GetGroupEndpointsOk

`func (o *ServiceGroupOutputShort) GetGroupEndpointsOk() (*map[string]ServiceEndpoint, bool)`

GetGroupEndpointsOk returns a tuple with the GroupEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupEndpoints

`func (o *ServiceGroupOutputShort) SetGroupEndpoints(v map[string]ServiceEndpoint)`

SetGroupEndpoints sets GroupEndpoints field to given value.

### HasGroupEndpoints

`func (o *ServiceGroupOutputShort) HasGroupEndpoints() bool`

HasGroupEndpoints returns a boolean if a field has been set.

### GetId

`func (o *ServiceGroupOutputShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceGroupOutputShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceGroupOutputShort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceGroupOutputShort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceGroupOutputShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceGroupOutputShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceGroupOutputShort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceGroupOutputShort) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


