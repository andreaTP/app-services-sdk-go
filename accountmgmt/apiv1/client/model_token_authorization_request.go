/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// TokenAuthorizationRequest struct for TokenAuthorizationRequest
type TokenAuthorizationRequest struct {
	AuthorizationToken *string `json:"authorization_token,omitempty"`
}

// NewTokenAuthorizationRequest instantiates a new TokenAuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAuthorizationRequest() *TokenAuthorizationRequest {
	this := TokenAuthorizationRequest{}
	return &this
}

// NewTokenAuthorizationRequestWithDefaults instantiates a new TokenAuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAuthorizationRequestWithDefaults() *TokenAuthorizationRequest {
	this := TokenAuthorizationRequest{}
	return &this
}

// GetAuthorizationToken returns the AuthorizationToken field value if set, zero value otherwise.
func (o *TokenAuthorizationRequest) GetAuthorizationToken() string {
	if o == nil || o.AuthorizationToken == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationToken
}

// GetAuthorizationTokenOk returns a tuple with the AuthorizationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationRequest) GetAuthorizationTokenOk() (*string, bool) {
	if o == nil || o.AuthorizationToken == nil {
		return nil, false
	}
	return o.AuthorizationToken, true
}

// HasAuthorizationToken returns a boolean if a field has been set.
func (o *TokenAuthorizationRequest) HasAuthorizationToken() bool {
	if o != nil && o.AuthorizationToken != nil {
		return true
	}

	return false
}

// SetAuthorizationToken gets a reference to the given string and assigns it to the AuthorizationToken field.
func (o *TokenAuthorizationRequest) SetAuthorizationToken(v string) {
	o.AuthorizationToken = &v
}

func (o TokenAuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizationToken != nil {
		toSerialize["authorization_token"] = o.AuthorizationToken
	}
	return json.Marshal(toSerialize)
}

type NullableTokenAuthorizationRequest struct {
	value *TokenAuthorizationRequest
	isSet bool
}

func (v NullableTokenAuthorizationRequest) Get() *TokenAuthorizationRequest {
	return v.value
}

func (v *NullableTokenAuthorizationRequest) Set(val *TokenAuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAuthorizationRequest(val *TokenAuthorizationRequest) *NullableTokenAuthorizationRequest {
	return &NullableTokenAuthorizationRequest{value: val, isSet: true}
}

func (v NullableTokenAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

