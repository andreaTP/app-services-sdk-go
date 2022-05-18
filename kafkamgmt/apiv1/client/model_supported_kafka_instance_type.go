/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.7.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// SupportedKafkaInstanceType Supported Kafka instance type
type SupportedKafkaInstanceType struct {

	// Unique identifier of the Kafka instance type.
	Id *string `json:"id,omitempty"`

	// Human readable name of the supported Kafka instance type
	DisplayName *string `json:"display_name,omitempty"`

	//  A list of Kafka instance sizes available for this instance type
	Sizes *[]SupportedKafkaSize `json:"sizes,omitempty"`

}

// NewSupportedKafkaInstanceType instantiates a new SupportedKafkaInstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedKafkaInstanceType() *SupportedKafkaInstanceType {
	this := SupportedKafkaInstanceType{}
	return &this
}

// NewSupportedKafkaInstanceTypeWithDefaults instantiates a new SupportedKafkaInstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedKafkaInstanceTypeWithDefaults() *SupportedKafkaInstanceType {
	this := SupportedKafkaInstanceType{}




	return &this
}


// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceType) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceType) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportedKafkaInstanceType) SetId(v string) {
	o.Id = &v
}


// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceType) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceType) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceType) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SupportedKafkaInstanceType) SetDisplayName(v string) {
	o.DisplayName = &v
}


// GetSizes returns the Sizes field value if set, zero value otherwise.
func (o *SupportedKafkaInstanceType) GetSizes() []SupportedKafkaSize {
	if o == nil || o.Sizes == nil {
		var ret []SupportedKafkaSize
		return ret
	}
	return *o.Sizes
}

// GetSizesOk returns a tuple with the Sizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedKafkaInstanceType) GetSizesOk() (*[]SupportedKafkaSize, bool) {
	if o == nil || o.Sizes == nil {
		return nil, false
	}
	return o.Sizes, true
}

// HasSizes returns a boolean if a field has been set.
func (o *SupportedKafkaInstanceType) HasSizes() bool {
	if o != nil && o.Sizes != nil {
		return true
	}

	return false
}

// SetSizes gets a reference to the given []SupportedKafkaSize and assigns it to the Sizes field.
func (o *SupportedKafkaInstanceType) SetSizes(v []SupportedKafkaSize) {
	o.Sizes = &v
}


func (o SupportedKafkaInstanceType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
    
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
    
	if o.Sizes != nil {
		toSerialize["sizes"] = o.Sizes
	}
    
	return json.Marshal(toSerialize)
}

type NullableSupportedKafkaInstanceType struct {
	value *SupportedKafkaInstanceType
	isSet bool
}

func (v NullableSupportedKafkaInstanceType) Get() *SupportedKafkaInstanceType {
	return v.value
}

func (v *NullableSupportedKafkaInstanceType) Set(val *SupportedKafkaInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedKafkaInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedKafkaInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedKafkaInstanceType(val *SupportedKafkaInstanceType) *NullableSupportedKafkaInstanceType {
	return &NullableSupportedKafkaInstanceType{value: val, isSet: true}
}

func (v NullableSupportedKafkaInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedKafkaInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

