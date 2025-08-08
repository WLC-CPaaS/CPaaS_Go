# ServiceSystemStatusOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpaasServices** | Pointer to [**ServiceSystemStatusCPAASService**](ServiceSystemStatusCPAASService.md) |  | [optional] 
**MessagingServices** | Pointer to [**ServiceSystemStatusMessagingService**](ServiceSystemStatusMessagingService.md) |  | [optional] 
**SupportServices** | Pointer to [**ServiceSystemStatusSupportService**](ServiceSystemStatusSupportService.md) |  | [optional] 

## Methods

### NewServiceSystemStatusOutput

`func NewServiceSystemStatusOutput() *ServiceSystemStatusOutput`

NewServiceSystemStatusOutput instantiates a new ServiceSystemStatusOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSystemStatusOutputWithDefaults

`func NewServiceSystemStatusOutputWithDefaults() *ServiceSystemStatusOutput`

NewServiceSystemStatusOutputWithDefaults instantiates a new ServiceSystemStatusOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpaasServices

`func (o *ServiceSystemStatusOutput) GetCpaasServices() ServiceSystemStatusCPAASService`

GetCpaasServices returns the CpaasServices field if non-nil, zero value otherwise.

### GetCpaasServicesOk

`func (o *ServiceSystemStatusOutput) GetCpaasServicesOk() (*ServiceSystemStatusCPAASService, bool)`

GetCpaasServicesOk returns a tuple with the CpaasServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpaasServices

`func (o *ServiceSystemStatusOutput) SetCpaasServices(v ServiceSystemStatusCPAASService)`

SetCpaasServices sets CpaasServices field to given value.

### HasCpaasServices

`func (o *ServiceSystemStatusOutput) HasCpaasServices() bool`

HasCpaasServices returns a boolean if a field has been set.

### GetMessagingServices

`func (o *ServiceSystemStatusOutput) GetMessagingServices() ServiceSystemStatusMessagingService`

GetMessagingServices returns the MessagingServices field if non-nil, zero value otherwise.

### GetMessagingServicesOk

`func (o *ServiceSystemStatusOutput) GetMessagingServicesOk() (*ServiceSystemStatusMessagingService, bool)`

GetMessagingServicesOk returns a tuple with the MessagingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingServices

`func (o *ServiceSystemStatusOutput) SetMessagingServices(v ServiceSystemStatusMessagingService)`

SetMessagingServices sets MessagingServices field to given value.

### HasMessagingServices

`func (o *ServiceSystemStatusOutput) HasMessagingServices() bool`

HasMessagingServices returns a boolean if a field has been set.

### GetSupportServices

`func (o *ServiceSystemStatusOutput) GetSupportServices() ServiceSystemStatusSupportService`

GetSupportServices returns the SupportServices field if non-nil, zero value otherwise.

### GetSupportServicesOk

`func (o *ServiceSystemStatusOutput) GetSupportServicesOk() (*ServiceSystemStatusSupportService, bool)`

GetSupportServicesOk returns a tuple with the SupportServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportServices

`func (o *ServiceSystemStatusOutput) SetSupportServices(v ServiceSystemStatusSupportService)`

SetSupportServices sets SupportServices field to given value.

### HasSupportServices

`func (o *ServiceSystemStatusOutput) HasSupportServices() bool`

HasSupportServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


