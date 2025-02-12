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

func TestAccIosxrIPv4AccessList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "access_list_name", "ACCESS1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.sequence_number", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_protocol", "tcp"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_source_address", "18.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_source_wildcard_mask", "0.255.255.255"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_source_port_range_start", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_source_port_range_end", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_destination_host", "11.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_destination_port_eq", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_dscp", "cs1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_ttl_eq", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_nexthop1_ipv4", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_nexthop2_ipv4", "3.4.5.6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list.test", "sequences.0.permit_log", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrIPv4AccessListConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrIPv4AccessListConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_ipv4_access_list.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-ipv4-access-list-cfg:/ipv4/access-lists/access-list[access-list-name=ACCESS1]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrIPv4AccessListConfig_minimum() string {
	config := `resource "iosxr_ipv4_access_list" "test" {` + "\n"
	config += `	access_list_name = "ACCESS1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrIPv4AccessListConfig_all() string {
	config := `resource "iosxr_ipv4_access_list" "test" {` + "\n"
	config += `	access_list_name = "ACCESS1"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `		sequence_number = 11` + "\n"
	config += `		permit_protocol = "tcp"` + "\n"
	config += `		permit_source_address = "18.0.0.0"` + "\n"
	config += `		permit_source_wildcard_mask = "0.255.255.255"` + "\n"
	config += `		permit_source_port_range_start = "100"` + "\n"
	config += `		permit_source_port_range_end = "200"` + "\n"
	config += `		permit_destination_host = "11.1.1.1"` + "\n"
	config += `		permit_destination_port_eq = "300"` + "\n"
	config += `		permit_dscp = "cs1"` + "\n"
	config += `		permit_ttl_eq = 10` + "\n"
	config += `		permit_nexthop1_ipv4 = "1.2.3.4"` + "\n"
	config += `		permit_nexthop2_ipv4 = "3.4.5.6"` + "\n"
	config += `		permit_log = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
