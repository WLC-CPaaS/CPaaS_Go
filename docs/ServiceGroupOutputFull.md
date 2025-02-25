# ServiceGroupOutputFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**map[string]ServiceEndpoint**](ServiceEndpoint.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceGroupOutputFull

`func NewServiceGroupOutputFull() *ServiceGroupOutputFull`

NewServiceGroupOutputFull instantiates a new ServiceGroupOutputFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceGroupOutputFullWithDefaults

`func NewServiceGroupOutputFullWithDefaults() *ServiceGroupOutputFull`

NewServiceGroupOutputFullWithDefaults instantiates a new ServiceGroupOutputFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServiceGroupOutputFull) GetEndpoints() map[string]ServiceEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServiceGroupOutputFull) GetEndpointsOk() (*map[string]ServiceEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServiceGroupOutputFull) SetEndpoints(v map[string]ServiceEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ServiceGroupOutputFull) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetId

`func (o *ServiceGroupOutputFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceGroupOutputFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceGroupOutputFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceGroupOutputFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceGroupOutputFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceGroupOutputFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceGroupOutputFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceGroupOutputFull) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


