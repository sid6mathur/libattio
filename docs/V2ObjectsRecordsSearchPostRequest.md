# V2ObjectsRecordsSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query string to search for. An empty string returns a default set of results. | 
**Limit** | Pointer to **float32** | The maximum number of results to return. Defaults to 25. | [optional] [default to 25]
**Objects** | **[]string** | Specifies which objects to filter results by. At least one object must be specified. Accepts object slugs or IDs. | 
**RequestAs** | [**V2ObjectsRecordsSearchPostRequestRequestAs**](V2ObjectsRecordsSearchPostRequestRequestAs.md) |  | 

## Methods

### NewV2ObjectsRecordsSearchPostRequest

`func NewV2ObjectsRecordsSearchPostRequest(query string, objects []string, requestAs V2ObjectsRecordsSearchPostRequestRequestAs, ) *V2ObjectsRecordsSearchPostRequest`

NewV2ObjectsRecordsSearchPostRequest instantiates a new V2ObjectsRecordsSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsRecordsSearchPostRequestWithDefaults

`func NewV2ObjectsRecordsSearchPostRequestWithDefaults() *V2ObjectsRecordsSearchPostRequest`

NewV2ObjectsRecordsSearchPostRequestWithDefaults instantiates a new V2ObjectsRecordsSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *V2ObjectsRecordsSearchPostRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V2ObjectsRecordsSearchPostRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V2ObjectsRecordsSearchPostRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetLimit

`func (o *V2ObjectsRecordsSearchPostRequest) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V2ObjectsRecordsSearchPostRequest) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V2ObjectsRecordsSearchPostRequest) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V2ObjectsRecordsSearchPostRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetObjects

`func (o *V2ObjectsRecordsSearchPostRequest) GetObjects() []string`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *V2ObjectsRecordsSearchPostRequest) GetObjectsOk() (*[]string, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *V2ObjectsRecordsSearchPostRequest) SetObjects(v []string)`

SetObjects sets Objects field to given value.


### GetRequestAs

`func (o *V2ObjectsRecordsSearchPostRequest) GetRequestAs() V2ObjectsRecordsSearchPostRequestRequestAs`

GetRequestAs returns the RequestAs field if non-nil, zero value otherwise.

### GetRequestAsOk

`func (o *V2ObjectsRecordsSearchPostRequest) GetRequestAsOk() (*V2ObjectsRecordsSearchPostRequestRequestAs, bool)`

GetRequestAsOk returns a tuple with the RequestAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAs

`func (o *V2ObjectsRecordsSearchPostRequest) SetRequestAs(v V2ObjectsRecordsSearchPostRequestRequestAs)`

SetRequestAs sets RequestAs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


