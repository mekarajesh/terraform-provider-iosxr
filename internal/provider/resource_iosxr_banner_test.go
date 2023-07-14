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

func TestAccIosxrBanner(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_banner.test", "banner_type", "login"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_banner.test", "line", " Hello user !"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrBannerConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrBannerConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_banner.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-banner-cfg:/banners/banner[banner-type=login]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrBannerConfig_minimum() string {
	config := `resource "iosxr_banner" "test" {` + "\n"
	config += `	banner_type = "login"` + "\n"
	config += `	line = " Hello user !"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrBannerConfig_all() string {
	config := `resource "iosxr_banner" "test" {` + "\n"
	config += `	banner_type = "login"` + "\n"
	config += `	line = " Hello user !"` + "\n"
	config += `}` + "\n"
	return config
}
