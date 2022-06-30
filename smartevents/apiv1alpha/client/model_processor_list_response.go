/*
 * Red Hat Openshift SmartEvents Fleet Manager
 *
 * The API exposed by the fleet manager of the SmartEvents service.
 *
 * API version: 0.0.1
 * Contact: openbridge-dev@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package smarteventsclient

import (
	"encoding/json"
)

// ProcessorListResponse struct for ProcessorListResponse
type ProcessorListResponse struct {
	Kind *string `json:"kind,omitempty"`
	Items *[]ProcessorResponse `json:"items,omitempty"`
	Page *int64 `json:"page,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Total *int64 `json:"total,omitempty"`
}

// NewProcessorListResponse instantiates a new ProcessorListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorListResponse() *ProcessorListResponse {
	this := ProcessorListResponse{}
	return &this
}

// NewProcessorListResponseWithDefaults instantiates a new ProcessorListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorListResponseWithDefaults() *ProcessorListResponse {
	this := ProcessorListResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ProcessorListResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorListResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ProcessorListResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ProcessorListResponse) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ProcessorListResponse) GetItems() []ProcessorResponse {
	if o == nil || o.Items == nil {
		var ret []ProcessorResponse
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorListResponse) GetItemsOk() (*[]ProcessorResponse, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ProcessorListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ProcessorResponse and assigns it to the Items field.
func (o *ProcessorListResponse) SetItems(v []ProcessorResponse) {
	o.Items = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ProcessorListResponse) GetPage() int64 {
	if o == nil || o.Page == nil {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorListResponse) GetPageOk() (*int64, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ProcessorListResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *ProcessorListResponse) SetPage(v int64) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ProcessorListResponse) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorListResponse) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ProcessorListResponse) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ProcessorListResponse) SetSize(v int64) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ProcessorListResponse) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ProcessorListResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ProcessorListResponse) SetTotal(v int64) {
	o.Total = &v
}

func (o ProcessorListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableProcessorListResponse struct {
	value *ProcessorListResponse
	isSet bool
}

func (v NullableProcessorListResponse) Get() *ProcessorListResponse {
	return v.value
}

func (v *NullableProcessorListResponse) Set(val *ProcessorListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorListResponse(val *ProcessorListResponse) *NullableProcessorListResponse {
	return &NullableProcessorListResponse{value: val, isSet: true}
}

func (v NullableProcessorListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

