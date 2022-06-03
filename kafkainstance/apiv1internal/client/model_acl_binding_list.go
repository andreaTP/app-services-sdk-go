/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.11.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// ACLBindingList A page of ACL binding entries
type ACLBindingList struct {

	Items *[]AclBinding `json:"items,omitempty"`

}

// NewACLBindingList instantiates a new ACLBindingList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLBindingList() *ACLBindingList {
	this := ACLBindingList{}
	return &this
}

// NewACLBindingListWithDefaults instantiates a new ACLBindingList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLBindingListWithDefaults() *ACLBindingList {
	this := ACLBindingList{}


	return &this
}


// GetItems returns the Items field value if set, zero value otherwise.
func (o *ACLBindingList) GetItems() []AclBinding {
	if o == nil || o.Items == nil {
		var ret []AclBinding
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLBindingList) GetItemsOk() (*[]AclBinding, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ACLBindingList) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AclBinding and assigns it to the Items field.
func (o *ACLBindingList) SetItems(v []AclBinding) {
	o.Items = &v
}


func (o ACLBindingList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
    
	return json.Marshal(toSerialize)
}

type NullableACLBindingList struct {
	value *ACLBindingList
	isSet bool
}

func (v NullableACLBindingList) Get() *ACLBindingList {
	return v.value
}

func (v *NullableACLBindingList) Set(val *ACLBindingList) {
	v.value = val
	v.isSet = true
}

func (v NullableACLBindingList) IsSet() bool {
	return v.isSet
}

func (v *NullableACLBindingList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLBindingList(val *ACLBindingList) *NullableACLBindingList {
	return &NullableACLBindingList{value: val, isSet: true}
}

func (v NullableACLBindingList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLBindingList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
