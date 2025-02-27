# ServiceDocsPingGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServicePingOutput**](ServicePingOutput.md) |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewServiceDocsPingGet

`func NewServiceDocsPingGet() *ServiceDocsPingGet`

NewServiceDocsPingGet instantiates a new ServiceDocsPingGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsPingGetWithDefaults

`func NewServiceDocsPingGetWithDefaults() *ServiceDocsPingGet`

NewServiceDocsPingGetWithDefaults instantiates a new ServiceDocsPingGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsPingGet) GetData() ServicePingOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsPingGet) GetDataOk() (*ServicePingOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsPingGet) SetData(v ServicePingOutput)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsPingGet) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsPingGet) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsPingGet) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsPingGet) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsPingGet) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsPingGet) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsPingGet) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsPingGet) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsPingGet) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


