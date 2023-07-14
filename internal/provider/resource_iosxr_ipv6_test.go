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

func TestAccIosxrIPv6(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "hop_limit", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "icmp_error_interval", "2111"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "icmp_error_interval_bucket_size", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "source_route", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_timeout", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_max_packets", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_reassembler_drop_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_frag_hdr_incomplete_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "assembler_overlap_frag_drop_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "path_mtu_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv6.test", "path_mtu_timeout", "10"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrIPv6Config_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrIPv6Config_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_ipv6.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-ipv6-cfg:/ipv6",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrIPv6Config_minimum() string {
	config := `resource "iosxr_ipv6" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrIPv6Config_all() string {
	config := `resource "iosxr_ipv6" "test" {` + "\n"
	config += `	hop_limit = 123` + "\n"
	config += `	icmp_error_interval = 2111` + "\n"
	config += `	icmp_error_interval_bucket_size = 123` + "\n"
	config += `	source_route = true` + "\n"
	config += `	assembler_timeout = 50` + "\n"
	config += `	assembler_max_packets = 40` + "\n"
	config += `	assembler_reassembler_drop_enable = true` + "\n"
	config += `	assembler_frag_hdr_incomplete_enable = true` + "\n"
	config += `	assembler_overlap_frag_drop_enable = true` + "\n"
	config += `	path_mtu_enable = true` + "\n"
	config += `	path_mtu_timeout = 10` + "\n"
	config += `}` + "\n"
	return config
}
