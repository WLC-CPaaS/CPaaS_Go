# ServiceVOIPGroupAddEdit2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | [**map[string]ServiceEndpoint**](ServiceEndpoint.md) |  | 
**Name** | **string** |  | 

## Methods

### NewServiceVOIPGroupAddEdit2

`func NewServiceVOIPGroupAddEdit2(endpoints map[string]ServiceEndpoint, name string, ) *ServiceVOIPGroupAddEdit2`

NewServiceVOIPGroupAddEdit2 instantiates a new ServiceVOIPGroupAddEdit2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVOIPGroupAddEdit2WithDefaults

`func NewServiceVOIPGroupAddEdit2WithDefaults() *ServiceVOIPGroupAddEdit2`

NewServiceVOIPGroupAddEdit2WithDefaults instantiates a new ServiceVOIPGroupAddEdit2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServiceVOIPGroupAddEdit2) GetEndpoints() map[string]ServiceEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServiceVOIPGroupAddEdit2) GetEndpointsOk() (*map[string]ServiceEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServiceVOIPGroupAddEdit2) SetEndpoints(v map[string]ServiceEndpoint)`

SetEndpoints sets Endpoints field to given value.


### GetName

`func (o *ServiceVOIPGroupAddEdit2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceVOIPGroupAddEdit2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceVOIPGroupAddEdit2) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


