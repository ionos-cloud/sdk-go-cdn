/*
 * IONOS Cloud - CDN Distribution API
 *
 * This API manages CDN distributions.
 *
 * API version: 1.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// RoutingRule struct for RoutingRule
type RoutingRule struct {
	// The scheme of the routing rule.
	Scheme *string `json:"scheme"`
	// The prefix of the routing rule.
	Prefix   *string   `json:"prefix"`
	Upstream *Upstream `json:"upstream"`
}

// NewRoutingRule instantiates a new RoutingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingRule(scheme string, prefix string, upstream Upstream) *RoutingRule {
	this := RoutingRule{}

	this.Scheme = &scheme
	this.Prefix = &prefix
	this.Upstream = &upstream

	return &this
}

// NewRoutingRuleWithDefaults instantiates a new RoutingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingRuleWithDefaults() *RoutingRule {
	this := RoutingRule{}
	return &this
}

// GetScheme returns the Scheme field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RoutingRule) GetScheme() *string {
	if o == nil {
		return nil
	}

	return o.Scheme

}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingRule) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Scheme, true
}

// SetScheme sets field value
func (o *RoutingRule) SetScheme(v string) {

	o.Scheme = &v

}

// HasScheme returns a boolean if a field has been set.
func (o *RoutingRule) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// GetPrefix returns the Prefix field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RoutingRule) GetPrefix() *string {
	if o == nil {
		return nil
	}

	return o.Prefix

}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingRule) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Prefix, true
}

// SetPrefix sets field value
func (o *RoutingRule) SetPrefix(v string) {

	o.Prefix = &v

}

// HasPrefix returns a boolean if a field has been set.
func (o *RoutingRule) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// GetUpstream returns the Upstream field value
// If the value is explicit nil, the zero value for Upstream will be returned
func (o *RoutingRule) GetUpstream() *Upstream {
	if o == nil {
		return nil
	}

	return o.Upstream

}

// GetUpstreamOk returns a tuple with the Upstream field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoutingRule) GetUpstreamOk() (*Upstream, bool) {
	if o == nil {
		return nil, false
	}

	return o.Upstream, true
}

// SetUpstream sets field value
func (o *RoutingRule) SetUpstream(v Upstream) {

	o.Upstream = &v

}

// HasUpstream returns a boolean if a field has been set.
func (o *RoutingRule) HasUpstream() bool {
	if o != nil && o.Upstream != nil {
		return true
	}

	return false
}

func (o RoutingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}

	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}

	if o.Upstream != nil {
		toSerialize["upstream"] = o.Upstream
	}

	return json.Marshal(toSerialize)
}

type NullableRoutingRule struct {
	value *RoutingRule
	isSet bool
}

func (v NullableRoutingRule) Get() *RoutingRule {
	return v.value
}

func (v *NullableRoutingRule) Set(val *RoutingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingRule(val *RoutingRule) *NullableRoutingRule {
	return &NullableRoutingRule{value: val, isSet: true}
}

func (v NullableRoutingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
