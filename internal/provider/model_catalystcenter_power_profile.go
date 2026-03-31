// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type PowerProfile struct {
	Id          types.String        `tfsdk:"id"`
	ProfileName types.String        `tfsdk:"profile_name"`
	Description types.String        `tfsdk:"description"`
	Rules       []PowerProfileRules `tfsdk:"rules"`
}

type PowerProfileRules struct {
	InterfaceType  types.String `tfsdk:"interface_type"`
	InterfaceId    types.String `tfsdk:"interface_id"`
	ParameterType  types.String `tfsdk:"parameter_type"`
	ParameterValue types.String `tfsdk:"parameter_value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data PowerProfile) getPath() string {
	return "/dna/intent/api/v1/wirelessSettings/powerProfiles"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

func (data PowerProfile) getPathDelete() string {
	return "/dna/intent/api/v1/wirelessSettings/powerProfiles"
}

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

func (data PowerProfile) getPathGet() string {
	return "/dna/intent/api/v1/wirelessSettings/powerProfiles"
}

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data PowerProfile) toBody(ctx context.Context, state PowerProfile) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.ProfileName.IsNull() {
		body, _ = sjson.Set(body, "profileName", data.ProfileName.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Rules) > 0 {
		body, _ = sjson.Set(body, "rules", []interface{}{})
		for _, item := range data.Rules {
			itemBody := ""
			if !item.InterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceType", item.InterfaceType.ValueString())
			}
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interfaceId", item.InterfaceId.ValueString())
			}
			if !item.ParameterType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parameterType", item.ParameterType.ValueString())
			}
			if !item.ParameterValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "parameterValue", item.ParameterValue.ValueString())
			}
			body, _ = sjson.SetRaw(body, "rules.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *PowerProfile) fromBody(ctx context.Context, res gjson.Result) {
	// Retrieve the 'id' attribute, if Data Source doesn't require id
	if value := res.Get("response.0.id"); value.Exists() {
		data.Id = types.StringValue(value.String())
	} else {
		data.Id = types.StringNull()
	}
	if value := res.Get("response.0.profileName"); value.Exists() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("response.0.description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("response.0.rules"); value.Exists() && len(value.Array()) > 0 {
		data.Rules = make([]PowerProfileRules, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := PowerProfileRules{}
			if cValue := v.Get("interfaceType"); cValue.Exists() {
				item.InterfaceType = types.StringValue(cValue.String())
			} else {
				item.InterfaceType = types.StringNull()
			}
			if cValue := v.Get("interfaceId"); cValue.Exists() {
				item.InterfaceId = types.StringValue(cValue.String())
			} else {
				item.InterfaceId = types.StringNull()
			}
			if cValue := v.Get("parameterType"); cValue.Exists() {
				item.ParameterType = types.StringValue(cValue.String())
			} else {
				item.ParameterType = types.StringNull()
			}
			if cValue := v.Get("parameterValue"); cValue.Exists() {
				item.ParameterValue = types.StringValue(cValue.String())
			} else {
				item.ParameterValue = types.StringNull()
			}
			data.Rules = append(data.Rules, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *PowerProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.profileName"); value.Exists() && !data.ProfileName.IsNull() {
		data.ProfileName = types.StringValue(value.String())
	} else {
		data.ProfileName = types.StringNull()
	}
	if value := res.Get("response.0.description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	for i := range data.Rules {
		keys := [...]string{"interfaceType", "interfaceId", "parameterType", "parameterValue"}
		keyValues := [...]string{data.Rules[i].InterfaceType.ValueString(), data.Rules[i].InterfaceId.ValueString(), data.Rules[i].ParameterType.ValueString(), data.Rules[i].ParameterValue.ValueString()}

		var r gjson.Result
		res.Get("response.0.rules").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() == keyValues[ik] {
						found = true
						continue
					}
					found = false
					break
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		if value := r.Get("interfaceType"); value.Exists() && !data.Rules[i].InterfaceType.IsNull() {
			data.Rules[i].InterfaceType = types.StringValue(value.String())
		} else {
			data.Rules[i].InterfaceType = types.StringNull()
		}
		if value := r.Get("interfaceId"); value.Exists() && !data.Rules[i].InterfaceId.IsNull() {
			data.Rules[i].InterfaceId = types.StringValue(value.String())
		} else {
			data.Rules[i].InterfaceId = types.StringNull()
		}
		if value := r.Get("parameterType"); value.Exists() && !data.Rules[i].ParameterType.IsNull() {
			data.Rules[i].ParameterType = types.StringValue(value.String())
		} else {
			data.Rules[i].ParameterType = types.StringNull()
		}
		if value := r.Get("parameterValue"); value.Exists() && !data.Rules[i].ParameterValue.IsNull() {
			data.Rules[i].ParameterValue = types.StringValue(value.String())
		} else {
			data.Rules[i].ParameterValue = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *PowerProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Description.IsNull() {
		return false
	}
	if len(data.Rules) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
