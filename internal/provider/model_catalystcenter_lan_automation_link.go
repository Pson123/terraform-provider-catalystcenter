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
type LANAutomationLink struct {
	Id                               types.String `tfsdk:"id"`
	Action                           types.String `tfsdk:"action"`
	PrimaryDeviceManagementIpAddress types.String `tfsdk:"primary_device_management_ip_address"`
	PrimaryDeviceInterfaceName       types.String `tfsdk:"primary_device_interface_name"`
	PeerDeviceManagementIpAddress    types.String `tfsdk:"peer_device_management_ip_address"`
	PeerDeviceInterfaceName          types.String `tfsdk:"peer_device_interface_name"`
	IpPoolName                       types.String `tfsdk:"ip_pool_name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data LANAutomationLink) getPath() string {
	return "/dna/intent/api/v1/lan-automation/updateDevice"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin getFallbackPath

// End of section. //template:end getFallbackPath

// Section below is generated&owned by "gen/generator.go". //template:begin getPathDelete

// End of section. //template:end getPathDelete

// Section below is generated&owned by "gen/generator.go". //template:begin getPathGet

// End of section. //template:end getPathGet

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPost

// End of section. //template:end getPathPost

// Section below is generated&owned by "gen/generator.go". //template:begin getPathPut

// End of section. //template:end getPathPut

// Section below is generated&owned by "gen/generator.go". //template:begin getPathIdQuery

// End of section. //template:end getPathIdQuery

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data LANAutomationLink) toBody(ctx context.Context, state LANAutomationLink) string {
	body := ""
	put := false
	if state.Id.ValueString() != "" {
		put = true
	}
	_ = put
	if !data.PrimaryDeviceManagementIpAddress.IsNull() {
		body, _ = sjson.Set(body, "linkUpdate.sourceDeviceManagementIPAddress", data.PrimaryDeviceManagementIpAddress.ValueString())
	}
	if !data.PrimaryDeviceInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "linkUpdate.sourceDeviceInterfaceName", data.PrimaryDeviceInterfaceName.ValueString())
	}
	if !data.PeerDeviceManagementIpAddress.IsNull() {
		body, _ = sjson.Set(body, "linkUpdate.destinationDeviceManagementIPAddress", data.PeerDeviceManagementIpAddress.ValueString())
	}
	if !data.PeerDeviceInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "linkUpdate.destinationDeviceInterfaceName", data.PeerDeviceInterfaceName.ValueString())
	}
	if !data.IpPoolName.IsNull() {
		body, _ = sjson.Set(body, "linkUpdate.ipPoolName", data.IpPoolName.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *LANAutomationLink) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("linkUpdate.sourceDeviceManagementIPAddress"); value.Exists() {
		data.PrimaryDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("linkUpdate.sourceDeviceInterfaceName"); value.Exists() {
		data.PrimaryDeviceInterfaceName = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceInterfaceName = types.StringNull()
	}
	if value := res.Get("linkUpdate.destinationDeviceManagementIPAddress"); value.Exists() {
		data.PeerDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PeerDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("linkUpdate.destinationDeviceInterfaceName"); value.Exists() {
		data.PeerDeviceInterfaceName = types.StringValue(value.String())
	} else {
		data.PeerDeviceInterfaceName = types.StringNull()
	}
	if value := res.Get("linkUpdate.ipPoolName"); value.Exists() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *LANAutomationLink) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("linkUpdate.sourceDeviceManagementIPAddress"); value.Exists() && !data.PrimaryDeviceManagementIpAddress.IsNull() {
		data.PrimaryDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("linkUpdate.sourceDeviceInterfaceName"); value.Exists() && !data.PrimaryDeviceInterfaceName.IsNull() {
		data.PrimaryDeviceInterfaceName = types.StringValue(value.String())
	} else {
		data.PrimaryDeviceInterfaceName = types.StringNull()
	}
	if value := res.Get("linkUpdate.destinationDeviceManagementIPAddress"); value.Exists() && !data.PeerDeviceManagementIpAddress.IsNull() {
		data.PeerDeviceManagementIpAddress = types.StringValue(value.String())
	} else {
		data.PeerDeviceManagementIpAddress = types.StringNull()
	}
	if value := res.Get("linkUpdate.destinationDeviceInterfaceName"); value.Exists() && !data.PeerDeviceInterfaceName.IsNull() {
		data.PeerDeviceInterfaceName = types.StringValue(value.String())
	} else {
		data.PeerDeviceInterfaceName = types.StringNull()
	}
	if value := res.Get("linkUpdate.ipPoolName"); value.Exists() && !data.IpPoolName.IsNull() {
		data.IpPoolName = types.StringValue(value.String())
	} else {
		data.IpPoolName = types.StringNull()
	}
}

// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *LANAutomationLink) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.PrimaryDeviceManagementIpAddress.IsNull() {
		return false
	}
	if !data.PrimaryDeviceInterfaceName.IsNull() {
		return false
	}
	if !data.PeerDeviceManagementIpAddress.IsNull() {
		return false
	}
	if !data.PeerDeviceInterfaceName.IsNull() {
		return false
	}
	if !data.IpPoolName.IsNull() {
		return false
	}
	return true
}

// End of section. //template:end isNull
