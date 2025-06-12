# ModelsGenerateConfigFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | [**ModelsConfigFileParameter**](ModelsConfigFileParameter.md) |  | 
**TemplateFileId** | **string** |  | 
**TemplateId** | **string** |  | 

## Methods

### NewModelsGenerateConfigFileRequest

`func NewModelsGenerateConfigFileRequest(parameters ModelsConfigFileParameter, templateFileId string, templateId string, ) *ModelsGenerateConfigFileRequest`

NewModelsGenerateConfigFileRequest instantiates a new ModelsGenerateConfigFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsGenerateConfigFileRequestWithDefaults

`func NewModelsGenerateConfigFileRequestWithDefaults() *ModelsGenerateConfigFileRequest`

NewModelsGenerateConfigFileRequestWithDefaults instantiates a new ModelsGenerateConfigFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *ModelsGenerateConfigFileRequest) GetParameters() ModelsConfigFileParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ModelsGenerateConfigFileRequest) GetParametersOk() (*ModelsConfigFileParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ModelsGenerateConfigFileRequest) SetParameters(v ModelsConfigFileParameter)`

SetParameters sets Parameters field to given value.


### GetTemplateFileId

`func (o *ModelsGenerateConfigFileRequest) GetTemplateFileId() string`

GetTemplateFileId returns the TemplateFileId field if non-nil, zero value otherwise.

### GetTemplateFileIdOk

`func (o *ModelsGenerateConfigFileRequest) GetTemplateFileIdOk() (*string, bool)`

GetTemplateFileIdOk returns a tuple with the TemplateFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFileId

`func (o *ModelsGenerateConfigFileRequest) SetTemplateFileId(v string)`

SetTemplateFileId sets TemplateFileId field to given value.


### GetTemplateId

`func (o *ModelsGenerateConfigFileRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ModelsGenerateConfigFileRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ModelsGenerateConfigFileRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


