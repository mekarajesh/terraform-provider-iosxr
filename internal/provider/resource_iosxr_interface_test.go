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

func TestAccIosxrInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "l2transport", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "point_to_point", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "multipoint", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "dampening_decay_half_life_value", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv4_point_to_point", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "service_policy_input.0.name", "PMAP-IN"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "service_policy_output.0.name", "PMAP-OUT"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "mtu", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "bandwidth", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "description", "My Interface Description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "load_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "vrf", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv4_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv4_netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_link_local_address", "fe80::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_link_local_zone", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_autoconfig", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_addresses.0.address", "2001::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_addresses.0.prefix_length", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "ipv6_addresses.0.zone", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_interface.test", "bundle_port_priority", "100"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrInterfacePrerequisitesConfig + testAccIosxrInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrInterfacePrerequisitesConfig + testAccIosxrInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_interface.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-interface-cfg:/interfaces/interface[interface-name=GigabitEthernet0/0/0/1]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccIosxrInterfacePrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=PMAP-IN]"
	attributes = {
		"policy-map-name" = "PMAP-IN"
	}
	lists = [
		{
			name = "class"
			key = "name,type"
			items = [
				{
					"name" = "class-default"
					"type" = "qos"
					"set/qos-group" = "0"
				},
			] 
		},
	]
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-policymap-classmap-cfg:/policy-map/type/qos[policy-map-name=PMAP-OUT]"
	attributes = {
		"policy-map-name" = "PMAP-OUT"
	}
	lists = [
		{
			name = "class"
			key = "name,type"
			items = [
				{
					"name" = "class-default"
					"type" = "qos"
					"set/dscp" = "0"
				},
			] 
		},
	]
}

`

func testAccIosxrInterfaceConfig_minimum() string {
	config := `resource "iosxr_interface" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrInterfaceConfig_all() string {
	config := `resource "iosxr_interface" "test" {` + "\n"
	config += `	interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `	l2transport = false` + "\n"
	config += `	point_to_point = false` + "\n"
	config += `	multipoint = false` + "\n"
	config += `	dampening_decay_half_life_value = 2` + "\n"
	config += `	ipv4_point_to_point = true` + "\n"
	config += `	service_policy_input = [{` + "\n"
	config += `		name = "PMAP-IN"` + "\n"
	config += `	}]` + "\n"
	config += `	service_policy_output = [{` + "\n"
	config += `		name = "PMAP-OUT"` + "\n"
	config += `	}]` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	bandwidth = 100000` + "\n"
	config += `	description = "My Interface Description"` + "\n"
	config += `	load_interval = 30` + "\n"
	config += `	vrf = "VRF1"` + "\n"
	config += `	ipv4_address = "1.1.1.1"` + "\n"
	config += `	ipv4_netmask = "255.255.255.0"` + "\n"
	config += `	ipv6_link_local_address = "fe80::1"` + "\n"
	config += `	ipv6_link_local_zone = "0"` + "\n"
	config += `	ipv6_autoconfig = false` + "\n"
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		address = "2001::1"` + "\n"
	config += `		prefix_length = 64` + "\n"
	config += `		zone = "0"` + "\n"
	config += `	}]` + "\n"
	config += `	bundle_port_priority = 100` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
