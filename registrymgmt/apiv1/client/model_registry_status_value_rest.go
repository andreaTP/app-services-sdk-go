/*
 * Service Registry Fleet Manager
 *
 * Managed Service Registry cloud.redhat.com API Management API that lets you create new registry instances. Registry is a datastore for standard event schemas and API designs. Service Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Registry is an Managed version of upstream project called Apicurio Registry. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.
 *
 * API version: 0.0.5
 * Contact: rhosak-eval-support@redhat.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registrymgmtclient

import (
	"encoding/json"
	"fmt"
)

// RegistryStatusValueRest the model 'RegistryStatusValueRest'
type RegistryStatusValueRest string

// List of RegistryStatusValueRest
const (
	PROVISIONING RegistryStatusValueRest = "PROVISIONING"
	AVAILABLE RegistryStatusValueRest = "AVAILABLE"
	UNAVAILABLE RegistryStatusValueRest = "UNAVAILABLE"
)

var allowedRegistryStatusValueRestEnumValues = []RegistryStatusValueRest{
	"PROVISIONING",
	"AVAILABLE",
	"UNAVAILABLE",
}

func (v *RegistryStatusValueRest) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RegistryStatusValueRest(value)
	for _, existing := range allowedRegistryStatusValueRestEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RegistryStatusValueRest", value)
}

// NewRegistryStatusValueRestFromValue returns a pointer to a valid RegistryStatusValueRest
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRegistryStatusValueRestFromValue(v string) (*RegistryStatusValueRest, error) {
	ev := RegistryStatusValueRest(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RegistryStatusValueRest: valid values are %v", v, allowedRegistryStatusValueRestEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RegistryStatusValueRest) IsValid() bool {
	for _, existing := range allowedRegistryStatusValueRestEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RegistryStatusValueRest value
func (v RegistryStatusValueRest) Ptr() *RegistryStatusValueRest {
	return &v
}

type NullableRegistryStatusValueRest struct {
	value *RegistryStatusValueRest
	isSet bool
}

func (v NullableRegistryStatusValueRest) Get() *RegistryStatusValueRest {
	return v.value
}

func (v *NullableRegistryStatusValueRest) Set(val *RegistryStatusValueRest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryStatusValueRest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryStatusValueRest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryStatusValueRest(val *RegistryStatusValueRest) *NullableRegistryStatusValueRest {
	return &NullableRegistryStatusValueRest{value: val, isSet: true}
}

func (v NullableRegistryStatusValueRest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryStatusValueRest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

