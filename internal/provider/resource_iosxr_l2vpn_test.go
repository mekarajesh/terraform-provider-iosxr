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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrL2VPN(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn.test", "description", "My L2VPN Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn.test", "router_id", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn.test", "xconnect_groups.0.group_name", "P2P"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrL2VPNConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrL2VPNConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_l2vpn.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrL2VPNConfig_minimum() string {
	config := `resource "iosxr_l2vpn" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrL2VPNConfig_all() string {
	config := `resource "iosxr_l2vpn" "test" {` + "\n"
	config += `	description = "My L2VPN Description"` + "\n"
	config += `	router_id = "1.2.3.4"` + "\n"
	config += `	xconnect_groups = [{` + "\n"
	config += `		group_name = "P2P"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
