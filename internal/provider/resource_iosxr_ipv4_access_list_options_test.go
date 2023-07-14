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

func TestAccIosxrIPv4AccessListOptions(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list_options.test", "log_update_threshold", "214748"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_access_list_options.test", "log_update_rate", "1000"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrIPv4AccessListOptionsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrIPv4AccessListOptionsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_ipv4_access_list_options.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-ipv4-access-list-cfg:/ipv4/access-list-options",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrIPv4AccessListOptionsConfig_minimum() string {
	config := `resource "iosxr_ipv4_access_list_options" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrIPv4AccessListOptionsConfig_all() string {
	config := `resource "iosxr_ipv4_access_list_options" "test" {` + "\n"
	config += `	log_update_threshold = 214748` + "\n"
	config += `	log_update_rate = 1000` + "\n"
	config += `}` + "\n"
	return config
}
