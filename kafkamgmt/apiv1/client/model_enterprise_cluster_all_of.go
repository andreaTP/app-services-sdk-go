/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.14.0
 * Contact: rhosak-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkamgmtclient

import (
	"encoding/json"
)

// EnterpriseClusterAllOf struct for EnterpriseClusterAllOf
type EnterpriseClusterAllOf struct {
	// Indicates whether Kafkas created on this data plane cluster have to be accessed via private network
	AccessKafkasViaPrivateNetwork bool `json:"access_kafkas_via_private_network"`
	// ocm cluster id of the registered Enterprise cluster
	ClusterId *string `json:"cluster_id,omitempty"`
	// status of registered Enterprise cluster
	Status *string `json:"status,omitempty"`
}

// NewEnterpriseClusterAllOf instantiates a new EnterpriseClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseClusterAllOf(accessKafkasViaPrivateNetwork bool) *EnterpriseClusterAllOf {
	this := EnterpriseClusterAllOf{}
	this.AccessKafkasViaPrivateNetwork = accessKafkasViaPrivateNetwork
	return &this
}

// NewEnterpriseClusterAllOfWithDefaults instantiates a new EnterpriseClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseClusterAllOfWithDefaults() *EnterpriseClusterAllOf {
	this := EnterpriseClusterAllOf{}
	return &this
}

// GetAccessKafkasViaPrivateNetwork returns the AccessKafkasViaPrivateNetwork field value
func (o *EnterpriseClusterAllOf) GetAccessKafkasViaPrivateNetwork() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AccessKafkasViaPrivateNetwork
}

// GetAccessKafkasViaPrivateNetworkOk returns a tuple with the AccessKafkasViaPrivateNetwork field value
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOf) GetAccessKafkasViaPrivateNetworkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessKafkasViaPrivateNetwork, true
}

// SetAccessKafkasViaPrivateNetwork sets field value
func (o *EnterpriseClusterAllOf) SetAccessKafkasViaPrivateNetwork(v bool) {
	o.AccessKafkasViaPrivateNetwork = v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *EnterpriseClusterAllOf) GetClusterId() string {
	if o == nil || o.ClusterId == nil {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil || o.ClusterId == nil {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *EnterpriseClusterAllOf) HasClusterId() bool {
	if o != nil && o.ClusterId != nil {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *EnterpriseClusterAllOf) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnterpriseClusterAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseClusterAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnterpriseClusterAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnterpriseClusterAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o EnterpriseClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_kafkas_via_private_network"] = o.AccessKafkasViaPrivateNetwork
	}
	if o.ClusterId != nil {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseClusterAllOf struct {
	value *EnterpriseClusterAllOf
	isSet bool
}

func (v NullableEnterpriseClusterAllOf) Get() *EnterpriseClusterAllOf {
	return v.value
}

func (v *NullableEnterpriseClusterAllOf) Set(val *EnterpriseClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseClusterAllOf(val *EnterpriseClusterAllOf) *NullableEnterpriseClusterAllOf {
	return &NullableEnterpriseClusterAllOf{value: val, isSet: true}
}

func (v NullableEnterpriseClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


