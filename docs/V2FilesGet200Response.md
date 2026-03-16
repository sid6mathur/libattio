# V2FilesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]V2FilesGet200ResponseDataInner**](V2FilesGet200ResponseDataInner.md) |  | 
**Pagination** | [**V2MeetingsGet200ResponsePagination**](V2MeetingsGet200ResponsePagination.md) |  | 

## Methods

### NewV2FilesGet200Response

`func NewV2FilesGet200Response(data []V2FilesGet200ResponseDataInner, pagination V2MeetingsGet200ResponsePagination, ) *V2FilesGet200Response`

NewV2FilesGet200Response instantiates a new V2FilesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2FilesGet200ResponseWithDefaults

`func NewV2FilesGet200ResponseWithDefaults() *V2FilesGet200Response`

NewV2FilesGet200ResponseWithDefaults instantiates a new V2FilesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V2FilesGet200Response) GetData() []V2FilesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V2FilesGet200Response) GetDataOk() (*[]V2FilesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V2FilesGet200Response) SetData(v []V2FilesGet200ResponseDataInner)`

SetData sets Data field to given value.


### GetPagination

`func (o *V2FilesGet200Response) GetPagination() V2MeetingsGet200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V2FilesGet200Response) GetPaginationOk() (*V2MeetingsGet200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V2FilesGet200Response) SetPagination(v V2MeetingsGet200ResponsePagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


