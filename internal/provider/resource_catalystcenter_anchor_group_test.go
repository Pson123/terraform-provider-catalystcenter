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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccCcAnchorGroup(t *testing.T) {
	if os.Getenv("WIRELESS") == "" {
		t.Skip("skipping test, set environment variable WIRELESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "anchor_group_name", "AG1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.device_name", "WLC-ANCHOR"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.ip_address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.anchor_priority", "PRIMARY"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.managed_anchor_wlc", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.peer_device_type", "IOS-XE"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.mac_address", "aa:bb:cc:00:00:01"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_anchor_group.test", "mobility_anchors.0.mobility_group_name", "default"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcAnchorGroupConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "catalystcenter_anchor_group.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccCcAnchorGroupConfig_minimum() string {
	config := `resource "catalystcenter_anchor_group" "test" {` + "\n"
	config += `	anchor_group_name = "AG1"` + "\n"
	config += `	mobility_anchors = [{` + "\n"
	config += `	  device_name = "WLC-ANCHOR"` + "\n"
	config += `	  ip_address = "10.0.0.1"` + "\n"
	config += `	  anchor_priority = "PRIMARY"` + "\n"
	config += `	  managed_anchor_wlc = false` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccCcAnchorGroupConfig_all() string {
	config := `resource "catalystcenter_anchor_group" "test" {` + "\n"
	config += `	anchor_group_name = "AG1"` + "\n"
	config += `	mobility_anchors = [{` + "\n"
	config += `	  device_name = "WLC-ANCHOR"` + "\n"
	config += `	  ip_address = "10.0.0.1"` + "\n"
	config += `	  anchor_priority = "PRIMARY"` + "\n"
	config += `	  managed_anchor_wlc = false` + "\n"
	config += `	  peer_device_type = "IOS-XE"` + "\n"
	config += `	  mac_address = "aa:bb:cc:00:00:01"` + "\n"
	config += `	  mobility_group_name = "default"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
