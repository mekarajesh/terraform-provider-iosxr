// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrMPLSOAM(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_echo_disable_vendor_extension", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_echo_reply_mode_control_channel_allow_reverse_lsp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_dpm_pps", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_mpls_oam.test", "oam_dpm_interval", "60"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrMPLSOAMConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrMPLSOAMConfig() string {
	config := `resource "iosxr_mpls_oam" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	oam = true` + "\n"
	config += `	oam_echo_disable_vendor_extension = false` + "\n"
	config += `	oam_echo_reply_mode_control_channel_allow_reverse_lsp = false` + "\n"
	config += `	oam_dpm_pps = 10` + "\n"
	config += `	oam_dpm_interval = 60` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_mpls_oam" "test" {
			depends_on = [iosxr_mpls_oam.test]
		}
	`
	return config
}
