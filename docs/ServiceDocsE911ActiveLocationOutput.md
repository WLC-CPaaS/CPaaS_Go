# ServiceDocsE911ActiveLocationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceE911ActiveLocationOutput**](ServiceE911ActiveLocationOutput.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 

## Methods

### NewServiceDocsE911ActiveLocationOutput

`func NewServiceDocsE911ActiveLocationOutput() *ServiceDocsE911ActiveLocationOutput`

NewServiceDocsE911ActiveLocationOutput instantiates a new ServiceDocsE911ActiveLocationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsE911ActiveLocationOutputWithDefaults

`func NewServiceDocsE911ActiveLocationOutputWithDefaults() *ServiceDocsE911ActiveLocationOutput`

NewServiceDocsE911ActiveLocationOutputWithDefaults instantiates a new ServiceDocsE911ActiveLocationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsE911ActiveLocationOutput) GetData() ServiceE911ActiveLocationOutput`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsE911ActiveLocationOutput) GetDataOk() (*ServiceE911ActiveLocationOutput, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsE911ActiveLocationOutput) SetData(v ServiceE911ActiveLocationOutput)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsE911ActiveLocationOutput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsE911ActiveLocationOutput) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsE911ActiveLocationOutput) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsE911ActiveLocationOutput) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsE911ActiveLocationOutput) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsE911ActiveLocationOutput) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsE911ActiveLocationOutput) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsE911ActiveLocationOutput) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsE911ActiveLocationOutput) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


