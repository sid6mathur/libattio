# V2MeetingsPostRequestDataExternalRefAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IcalUid** | **string** | The ical uid of the meeting. | 
**Provider** | **string** | The email provider used to sync the meeting. | 
**OriginalStartTime** | Pointer to **string** | The original start time of the meeting. Use a timestamp with a specified offset for all day and non-all day meetings. This property is required for recurring event exceptions and optional otherwise. | [optional] 
**IsRecurring** | **bool** | Whether or not the meeting is recurring. | 

## Methods

### NewV2MeetingsPostRequestDataExternalRefAnyOf

`func NewV2MeetingsPostRequestDataExternalRefAnyOf(icalUid string, provider string, isRecurring bool, ) *V2MeetingsPostRequestDataExternalRefAnyOf`

NewV2MeetingsPostRequestDataExternalRefAnyOf instantiates a new V2MeetingsPostRequestDataExternalRefAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataExternalRefAnyOfWithDefaults

`func NewV2MeetingsPostRequestDataExternalRefAnyOfWithDefaults() *V2MeetingsPostRequestDataExternalRefAnyOf`

NewV2MeetingsPostRequestDataExternalRefAnyOfWithDefaults instantiates a new V2MeetingsPostRequestDataExternalRefAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcalUid

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetIcalUid() string`

GetIcalUid returns the IcalUid field if non-nil, zero value otherwise.

### GetIcalUidOk

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetIcalUidOk() (*string, bool)`

GetIcalUidOk returns a tuple with the IcalUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcalUid

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) SetIcalUid(v string)`

SetIcalUid sets IcalUid field to given value.


### GetProvider

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetOriginalStartTime

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetOriginalStartTime() string`

GetOriginalStartTime returns the OriginalStartTime field if non-nil, zero value otherwise.

### GetOriginalStartTimeOk

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetOriginalStartTimeOk() (*string, bool)`

GetOriginalStartTimeOk returns a tuple with the OriginalStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalStartTime

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) SetOriginalStartTime(v string)`

SetOriginalStartTime sets OriginalStartTime field to given value.

### HasOriginalStartTime

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) HasOriginalStartTime() bool`

HasOriginalStartTime returns a boolean if a field has been set.

### GetIsRecurring

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetIsRecurring() bool`

GetIsRecurring returns the IsRecurring field if non-nil, zero value otherwise.

### GetIsRecurringOk

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) GetIsRecurringOk() (*bool, bool)`

GetIsRecurringOk returns a tuple with the IsRecurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecurring

`func (o *V2MeetingsPostRequestDataExternalRefAnyOf) SetIsRecurring(v bool)`

SetIsRecurring sets IsRecurring field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


