# ServiceDocsE911LocationsURIApiOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceE911LocationURI**](ServiceE911LocationURI.md) |  | [optional] 
**RequestId** | Pointer to **string** | Unique id for each request | [optional] 
**StatusCode** | Pointer to **int32** | HTTP response status code | [optional] 

## Methods

### NewServiceDocsE911LocationsURIApiOutput

`func NewServiceDocsE911LocationsURIApiOutput() *ServiceDocsE911LocationsURIApiOutput`

NewServiceDocsE911LocationsURIApiOutput instantiates a new ServiceDocsE911LocationsURIApiOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDocsE911LocationsURIApiOutputWithDefaults

`func NewServiceDocsE911LocationsURIApiOutputWithDefaults() *ServiceDocsE911LocationsURIApiOutput`

NewServiceDocsE911LocationsURIApiOutputWithDefaults instantiates a new ServiceDocsE911LocationsURIApiOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceDocsE911LocationsURIApiOutput) GetData() []ServiceE911LocationURI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceDocsE911LocationsURIApiOutput) GetDataOk() (*[]ServiceE911LocationURI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceDocsE911LocationsURIApiOutput) SetData(v []ServiceE911LocationURI)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceDocsE911LocationsURIApiOutput) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ServiceDocsE911LocationsURIApiOutput) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceDocsE911LocationsURIApiOutput) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceDocsE911LocationsURIApiOutput) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceDocsE911LocationsURIApiOutput) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceDocsE911LocationsURIApiOutput) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceDocsE911LocationsURIApiOutput) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceDocsE911LocationsURIApiOutput) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceDocsE911LocationsURIApiOutput) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


