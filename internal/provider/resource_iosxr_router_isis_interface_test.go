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

func TestAccIosxrRouterISISInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "circuit_type", "level-1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "hello_padding_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "hello_padding_sometimes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "point_to_point", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "passive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "suppressed", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "hello_password_keychain", "KEY_CHAIN_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_isis_interface.test", "bfd_fast_detect_ipv6", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterISISInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterISISInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_isis_interface.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-router-isis-cfg:/router/isis/processes/process[process-id=P1]/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrRouterISISInterfaceConfig_minimum() string {
	config := `resource "iosxr_router_isis_interface" "test" {` + "\n"
	config += `	process_id = "P1"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterISISInterfaceConfig_all() string {
	config := `resource "iosxr_router_isis_interface" "test" {` + "\n"
	config += `	process_id = "P1"` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	circuit_type = "level-1"` + "\n"
	config += `	hello_padding_disable = true` + "\n"
	config += `	hello_padding_sometimes = false` + "\n"
	config += `	priority = 10` + "\n"
	config += `	point_to_point = false` + "\n"
	config += `	passive = false` + "\n"
	config += `	suppressed = false` + "\n"
	config += `	shutdown = false` + "\n"
	config += `	hello_password_keychain = "KEY_CHAIN_1"` + "\n"
	config += `	bfd_fast_detect_ipv6 = true` + "\n"
	config += `}` + "\n"
	return config
}
