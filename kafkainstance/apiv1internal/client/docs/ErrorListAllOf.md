# ErrorListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Error**](Error.md) |  | [optional] 
**Total** | Pointer to **interface{}** | Total number of errors returned in this request | [optional] 


## Methods

### NewErrorListAllOf

`func NewErrorListAllOf() *ErrorListAllOf`

NewErrorListAllOf instantiates a new ErrorListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorListAllOfWithDefaults

`func NewErrorListAllOfWithDefaults() *ErrorListAllOf`

NewErrorListAllOfWithDefaults instantiates a new ErrorListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetItems

`func (o *ErrorListAllOf) GetItems() []Error`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ErrorListAllOf) GetItemsOk() (*[]Error, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ErrorListAllOf) SetItems(v []Error)`

SetItems sets Items field to given value.

### HasItems

`func (o *ErrorListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


### GetTotal

`func (o *ErrorListAllOf) GetTotal() interface{}`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ErrorListAllOf) GetTotalOk() (*interface{}, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ErrorListAllOf) SetTotal(v interface{})`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ErrorListAllOf) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ErrorListAllOf) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ErrorListAllOf) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
