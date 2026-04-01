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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceCcPowerProfile(t *testing.T) {
	if os.Getenv("WIRELESS") == "" {
		t.Skip("skipping test, set environment variable WIRELESS")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_power_profile.test", "profile_name", "Low_Power_Profile"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_power_profile.test", "description", "Reduce power consumption on APs"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_power_profile.test", "rules.0.interface_type", "RADIO"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_power_profile.test", "rules.0.interface_id", "6GHZ"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_power_profile.test", "rules.0.parameter_type", "STATE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_power_profile.test", "rules.0.parameter_value", "DISABLE"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcPowerProfileConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceCcPowerProfileConfig() string {
	config := `resource "catalystcenter_power_profile" "test" {` + "\n"
	config += `	profile_name = "Low_Power_Profile"` + "\n"
	config += `	description = "Reduce power consumption on APs"` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  interface_type = "RADIO"` + "\n"
	config += `	  interface_id = "6GHZ"` + "\n"
	config += `	  parameter_type = "STATE"` + "\n"
	config += `	  parameter_value = "DISABLE"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_power_profile" "test" {
			profile_name = "Low_Power_Profile"
			depends_on = [catalystcenter_power_profile.test]
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
