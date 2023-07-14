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

func TestAccIosxrSNMPServerVRFHost(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_snmp_server_vrf_host.test", "address", "11.11.11.11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_snmp_server_vrf_host.test", "unencrypted_strings.0.community_string", "COMMUNITY1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_snmp_server_vrf_host.test", "unencrypted_strings.0.version_v3_security_level", "auth"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrSNMPServerVRFHostConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrSNMPServerVRFHostConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_snmp_server_vrf_host.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-snmp-server-cfg:/snmp-server/vrfs/vrf[vrf-name=VRF1]/hosts/host[address=11.11.11.11]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrSNMPServerVRFHostConfig_minimum() string {
	config := `resource "iosxr_snmp_server_vrf_host" "test" {` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	address = "11.11.11.11"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrSNMPServerVRFHostConfig_all() string {
	config := `resource "iosxr_snmp_server_vrf_host" "test" {` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	address = "11.11.11.11"` + "\n"
	config += `	unencrypted_strings = [{` + "\n"
	config += `		community_string = "COMMUNITY1"` + "\n"
	config += `		version_v3_security_level = "auth"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
