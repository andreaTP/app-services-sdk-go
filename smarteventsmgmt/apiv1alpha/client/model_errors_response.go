/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsmgmtclient

import (
	"encoding/json"
)

// ErrorsResponse struct for ErrorsResponse
type ErrorsResponse struct {
	Kind *string `json:"kind,omitempty"`
	Items *[]ErrorResponse `json:"items,omitempty"`
}

// NewErrorsResponse instantiates a new ErrorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorsResponse() *ErrorsResponse {
	this := ErrorsResponse{}
	return &this
}

// NewErrorsResponseWithDefaults instantiates a new ErrorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsResponseWithDefaults() *ErrorsResponse {
	this := ErrorsResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ErrorsResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ErrorsResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ErrorsResponse) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ErrorsResponse) GetItems() []ErrorResponse {
	if o == nil || o.Items == nil {
		var ret []ErrorResponse
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorsResponse) GetItemsOk() (*[]ErrorResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ErrorsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ErrorResponse and assigns it to the Items field.
func (o *ErrorsResponse) SetItems(v []ErrorResponse) {
	o.Items = &v
}

func (o ErrorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableErrorsResponse struct {
	value *ErrorsResponse
	isSet bool
}

func (v NullableErrorsResponse) Get() *ErrorsResponse {
	return v.value
}

func (v *NullableErrorsResponse) Set(val *ErrorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorsResponse(val *ErrorsResponse) *NullableErrorsResponse {
	return &NullableErrorsResponse{value: val, isSet: true}
}

func (v NullableErrorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

