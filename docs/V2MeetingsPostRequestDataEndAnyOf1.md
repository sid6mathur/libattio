# V2MeetingsPostRequestDataEndAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | An ISO 8601 date indicating when an all day meeting ends. Note that dates are exclusive, meaning that the meeting ends before the specified time, not at it. For example, a one day meeting on June 3rd would end on June 4th, not June 3rd. | 

## Methods

### NewV2MeetingsPostRequestDataEndAnyOf1

`func NewV2MeetingsPostRequestDataEndAnyOf1(date string, ) *V2MeetingsPostRequestDataEndAnyOf1`

NewV2MeetingsPostRequestDataEndAnyOf1 instantiates a new V2MeetingsPostRequestDataEndAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataEndAnyOf1WithDefaults

`func NewV2MeetingsPostRequestDataEndAnyOf1WithDefaults() *V2MeetingsPostRequestDataEndAnyOf1`

NewV2MeetingsPostRequestDataEndAnyOf1WithDefaults instantiates a new V2MeetingsPostRequestDataEndAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *V2MeetingsPostRequestDataEndAnyOf1) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *V2MeetingsPostRequestDataEndAnyOf1) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *V2MeetingsPostRequestDataEndAnyOf1) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


