/*
 * Svix API
 *
 * Welcome to the Svix API documentation!  Useful links: [Homepage](https://www.svix.com) | [Support email](mailto:support+docs@svix.com) | [Blog](https://www.svix.com/blog/) | [Slack Community](https://www.svix.com/slack/)  # Introduction  This is the reference documentation and schemas for the [Svix webhook service](https://www.svix.com) API. For tutorials and other documentation please refer to [the documentation](https://docs.svix.com).  ## Main concepts  In Svix you have four important entities you will be interacting with:  - `messages`: these are the webhooks being sent. They can have contents and a few other properties. - `application`: this is where `messages` are sent to. Usually you want to create one application for each user on your platform. - `endpoint`: endpoints are the URLs messages will be sent to. Each application can have multiple `endpoints` and each message sent to that application will be sent to all of them (unless they are not subscribed to the sent event type). - `event-type`: event types are identifiers denoting the type of the message being sent. Event types are primarily used to decide which events are sent to which endpoint.   ## Authentication  Get your authentication token (`AUTH_TOKEN`) from the [Svix dashboard](https://dashboard.svix.com) and use it as part of the `Authorization` header as such: `Authorization: Bearer ${AUTH_TOKEN}`.  <SecurityDefinitions />   ## Code samples  The code samples assume you already have the respective libraries installed and you know how to use them. For the latest information on how to do that, please refer to [the documentation](https://docs.svix.com/).   ## Idempotency  Svix supports [idempotency](https://en.wikipedia.org/wiki/Idempotence) for safely retrying requests without accidentally performing the same operation twice. This is useful when an API call is disrupted in transit and you do not receive a response.  To perform an idempotent request, pass the idempotency key in the `Idempotency-Key` header to the request. The idempotency key should be a unique value generated by the client. You can create the key in however way you like, though we suggest using UUID v4, or any other string with enough entropy to avoid collisions.  Svix's idempotency works by saving the resulting status code and body of the first request made for any given idempotency key for any successful request. Subsequent requests with the same key return the same result.  Please note that idempotency is only supported for `POST` requests.   ## Cross-Origin Resource Sharing  This API features Cross-Origin Resource Sharing (CORS) implemented in compliance with [W3C spec](https://www.w3.org/TR/cors/). And that allows cross-domain communication from the browser. All responses have a wildcard same-origin which makes them completely public and accessible to everyone, including any code on any site. 
 *
 * API version: 1.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FontSizeConfig struct for FontSizeConfig
type FontSizeConfig struct {
	Base NullableInt32 `json:"base,omitempty"`
}

// NewFontSizeConfig instantiates a new FontSizeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFontSizeConfig() *FontSizeConfig {
	this := FontSizeConfig{}
	return &this
}

// NewFontSizeConfigWithDefaults instantiates a new FontSizeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFontSizeConfigWithDefaults() *FontSizeConfig {
	this := FontSizeConfig{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FontSizeConfig) GetBase() int32 {
	if o == nil || o.Base.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Base.Get()
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FontSizeConfig) GetBaseOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Base.Get(), o.Base.IsSet()
}

// HasBase returns a boolean if a field has been set.
func (o *FontSizeConfig) HasBase() bool {
	if o != nil && o.Base.IsSet() {
		return true
	}

	return false
}

// SetBase gets a reference to the given NullableInt32 and assigns it to the Base field.
func (o *FontSizeConfig) SetBase(v int32) {
	o.Base.Set(&v)
}
// SetBaseNil sets the value for Base to be an explicit nil
func (o *FontSizeConfig) SetBaseNil() {
	o.Base.Set(nil)
}

// UnsetBase ensures that no value is present for Base, not even an explicit nil
func (o *FontSizeConfig) UnsetBase() {
	o.Base.Unset()
}

func (o FontSizeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Base.IsSet() {
		toSerialize["base"] = o.Base.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableFontSizeConfig struct {
	value *FontSizeConfig
	isSet bool
}

func (v NullableFontSizeConfig) Get() *FontSizeConfig {
	return v.value
}

func (v *NullableFontSizeConfig) Set(val *FontSizeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFontSizeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFontSizeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFontSizeConfig(val *FontSizeConfig) *NullableFontSizeConfig {
	return &NullableFontSizeConfig{value: val, isSet: true}
}

func (v NullableFontSizeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFontSizeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

