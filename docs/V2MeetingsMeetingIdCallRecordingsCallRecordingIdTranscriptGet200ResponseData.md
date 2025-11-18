# V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataId**](V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataId.md) |  | 
**Transcript** | [**[]V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataTranscriptInner**](V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataTranscriptInner.md) | The transcript segments with speech, timing, and speaker information. | 
**RawTranscript** | **string** | The raw transcript of the call recording. | 
**WebUrl** | **string** | A URL that links directly to the call recording transcript in the Attio web application. | 

## Methods

### NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData

`func NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData(id V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataId, transcript []V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataTranscriptInner, rawTranscript string, webUrl string, ) *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData`

NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData instantiates a new V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataWithDefaults

`func NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataWithDefaults() *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData`

NewV2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataWithDefaults instantiates a new V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetId() V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetIdOk() (*V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) SetId(v V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataId)`

SetId sets Id field to given value.


### GetTranscript

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetTranscript() []V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataTranscriptInner`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetTranscriptOk() (*[]V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataTranscriptInner, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) SetTranscript(v []V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseDataTranscriptInner)`

SetTranscript sets Transcript field to given value.


### GetRawTranscript

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetRawTranscript() string`

GetRawTranscript returns the RawTranscript field if non-nil, zero value otherwise.

### GetRawTranscriptOk

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetRawTranscriptOk() (*string, bool)`

GetRawTranscriptOk returns a tuple with the RawTranscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTranscript

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) SetRawTranscript(v string)`

SetRawTranscript sets RawTranscript field to given value.


### GetWebUrl

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *V2MeetingsMeetingIdCallRecordingsCallRecordingIdTranscriptGet200ResponseData) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


