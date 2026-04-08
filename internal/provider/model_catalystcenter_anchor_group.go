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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type AnchorGroup struct {
	Id              types.String                 `tfsdk:"id"`
	AnchorGroupName types.String                 `tfsdk:"anchor_group_name"`
	MobilityAnchors []AnchorGroupMobilityAnchors `tfsdk:"mobility_anchors"`
}

type AnchorGroupMobilityAnchors struct {
	DeviceName        types.String `tfsdk:"device_name"`
	IpAddress         types.String `tfsdk:"ip_address"`
	AnchorPriority    types.String `tfsdk:"anchor_priority"`
	ManagedAnchorWlc  types.Bool   `tfsdk:"managed_anchor_wlc"`
	PeerDeviceType    types.String `tfsdk:"peer_device_type"`
	MacAddress        types.String `tfsdk:"mac_address"`
	MobilityGroupName types.String `tfsdk:"mobility_group_name"`
	PrivateIp         types.String `tfsdk:"private_ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data AnchorGroup) getPath() string {
	return "/dna/intent/api/v1/wirelessSettings/anchorGroups"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data AnchorGroup) toBody(ctx context.Context, state AnchorGroup) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.AnchorGroupName.IsNull() {
		body, _ = sjson.Set(body, "anchorGroupName", data.AnchorGroupName.ValueString())
	}
	if len(data.MobilityAnchors) > 0 {
		body, _ = sjson.Set(body, "mobilityAnchors", []interface{}{})
		for _, item := range data.MobilityAnchors {
			itemBody := ""
			if !item.DeviceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceName", item.DeviceName.ValueString())
			}
			if !item.IpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress", item.IpAddress.ValueString())
			}
			if !item.AnchorPriority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "anchorPriority", item.AnchorPriority.ValueString())
			}
			if !item.ManagedAnchorWlc.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "managedAnchorWlc", item.ManagedAnchorWlc.ValueBool())
			}
			if !item.PeerDeviceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "peerDeviceType", item.PeerDeviceType.ValueString())
			}
			if !item.MacAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "macAddress", item.MacAddress.ValueString())
			}
			if !item.MobilityGroupName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "mobilityGroupName", item.MobilityGroupName.ValueString())
			}
			if !item.PrivateIp.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "privateIp", item.PrivateIp.ValueString())
			}
			body, _ = sjson.SetRaw(body, "mobilityAnchors.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *AnchorGroup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("anchorGroupName"); value.Exists() {
		data.AnchorGroupName = types.StringValue(value.String())
	} else {
		data.AnchorGroupName = types.StringNull()
	}
	if value := res.Get("mobilityAnchors"); value.Exists() && len(value.Array()) > 0 {
		data.MobilityAnchors = make([]AnchorGroupMobilityAnchors, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := AnchorGroupMobilityAnchors{}
			if cValue := v.Get("deviceName"); cValue.Exists() {
				item.DeviceName = types.StringValue(cValue.String())
			} else {
				item.DeviceName = types.StringNull()
			}
			if cValue := v.Get("ipAddress"); cValue.Exists() {
				item.IpAddress = types.StringValue(cValue.String())
			} else {
				item.IpAddress = types.StringNull()
			}
			if cValue := v.Get("anchorPriority"); cValue.Exists() {
				item.AnchorPriority = types.StringValue(cValue.String())
			} else {
				item.AnchorPriority = types.StringNull()
			}
			if cValue := v.Get("managedAnchorWlc"); cValue.Exists() {
				item.ManagedAnchorWlc = types.BoolValue(cValue.Bool())
			} else {
				item.ManagedAnchorWlc = types.BoolNull()
			}
			if cValue := v.Get("peerDeviceType"); cValue.Exists() {
				item.PeerDeviceType = types.StringValue(cValue.String())
			} else {
				item.PeerDeviceType = types.StringNull()
			}
			if cValue := v.Get("macAddress"); cValue.Exists() {
				item.MacAddress = types.StringValue(cValue.String())
			} else {
				item.MacAddress = types.StringNull()
			}
			if cValue := v.Get("mobilityGroupName"); cValue.Exists() {
				item.MobilityGroupName = types.StringValue(cValue.String())
			} else {
				item.MobilityGroupName = types.StringNull()
			}
			if cValue := v.Get("privateIp"); cValue.Exists() {
				item.PrivateIp = types.StringValue(cValue.String())
			} else {
				item.PrivateIp = types.StringNull()
			}
			data.MobilityAnchors = append(data.MobilityAnchors, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *AnchorGroup) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("anchorGroupName"); value.Exists() && !data.AnchorGroupName.IsNull() {
		data.AnchorGroupName = types.StringValue(value.String())
	} else {
		data.AnchorGroupName = types.StringNull()
	}
	for i := range data.MobilityAnchors {
		keys := [...]string{"deviceName", "ipAddress", "anchorPriority", "managedAnchorWlc", "peerDeviceType", "macAddress", "mobilityGroupName", "privateIp"}
		keyValues := [...]string{data.MobilityAnchors[i].DeviceName.ValueString(), data.MobilityAnchors[i].IpAddress.ValueString(), data.MobilityAnchors[i].AnchorPriority.ValueString(), strconv.FormatBool(data.MobilityAnchors[i].ManagedAnchorWlc.ValueBool()), data.MobilityAnchors[i].PeerDeviceType.ValueString(), data.MobilityAnchors[i].MacAddress.ValueString(), data.MobilityAnchors[i].MobilityGroupName.ValueString(), data.MobilityAnchors[i].PrivateIp.ValueString()}

		var r gjson.Result
		res.Get("mobilityAnchors").ForEach(
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
		if value := r.Get("deviceName"); value.Exists() && !data.MobilityAnchors[i].DeviceName.IsNull() {
			data.MobilityAnchors[i].DeviceName = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].DeviceName = types.StringNull()
		}
		if value := r.Get("ipAddress"); value.Exists() && !data.MobilityAnchors[i].IpAddress.IsNull() {
			data.MobilityAnchors[i].IpAddress = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].IpAddress = types.StringNull()
		}
		if value := r.Get("anchorPriority"); value.Exists() && !data.MobilityAnchors[i].AnchorPriority.IsNull() {
			data.MobilityAnchors[i].AnchorPriority = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].AnchorPriority = types.StringNull()
		}
		if value := r.Get("managedAnchorWlc"); value.Exists() && !data.MobilityAnchors[i].ManagedAnchorWlc.IsNull() {
			data.MobilityAnchors[i].ManagedAnchorWlc = types.BoolValue(value.Bool())
		} else {
			data.MobilityAnchors[i].ManagedAnchorWlc = types.BoolNull()
		}
		if value := r.Get("peerDeviceType"); value.Exists() && !data.MobilityAnchors[i].PeerDeviceType.IsNull() {
			data.MobilityAnchors[i].PeerDeviceType = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].PeerDeviceType = types.StringNull()
		}
		if value := r.Get("macAddress"); value.Exists() && !data.MobilityAnchors[i].MacAddress.IsNull() {
			data.MobilityAnchors[i].MacAddress = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].MacAddress = types.StringNull()
		}
		if value := r.Get("mobilityGroupName"); value.Exists() && !data.MobilityAnchors[i].MobilityGroupName.IsNull() {
			data.MobilityAnchors[i].MobilityGroupName = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].MobilityGroupName = types.StringNull()
		}
		if value := r.Get("privateIp"); value.Exists() && !data.MobilityAnchors[i].PrivateIp.IsNull() {
			data.MobilityAnchors[i].PrivateIp = types.StringValue(value.String())
		} else {
			data.MobilityAnchors[i].PrivateIp = types.StringNull()
		}
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *AnchorGroup) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.AnchorGroupName.IsNull() {
		return false
	}
	if len(data.MobilityAnchors) > 0 {
		return false
	}
	return true
}

// End of section. //template:end isNull
