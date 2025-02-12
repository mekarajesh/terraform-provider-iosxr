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

func TestAccIosxrKeyChain(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "name", "KEY11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.key_name", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.key_string_password", "00071A150754"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.cryptographic_algorithm", "hmac-md5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_hour", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_minute", "52"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_second", "55"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_day_of_month", "21"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_month", "january"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_start_time_year", "2023"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.accept_lifetime_infinite", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_start_time_hour", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_start_time_minute", "36"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_start_time_second", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_start_time_day_of_month", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_start_time_month", "january"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_start_time_year", "2023"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_key_chain.test", "keys.0.send_lifetime_infinite", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrKeyChainConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrKeyChainConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_key_chain.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-key-chain-cfg:/key/chains/chain[key-chain-name=KEY11]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrKeyChainConfig_minimum() string {
	config := `resource "iosxr_key_chain" "test" {` + "\n"
	config += `	name = "KEY11"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrKeyChainConfig_all() string {
	config := `resource "iosxr_key_chain" "test" {` + "\n"
	config += `	name = "KEY11"` + "\n"
	config += `	keys = [{` + "\n"
	config += `		key_name = "1"` + "\n"
	config += `		key_string_password = "00071A150754"` + "\n"
	config += `		cryptographic_algorithm = "hmac-md5"` + "\n"
	config += `		accept_lifetime_start_time_hour = 11` + "\n"
	config += `		accept_lifetime_start_time_minute = 52` + "\n"
	config += `		accept_lifetime_start_time_second = 55` + "\n"
	config += `		accept_lifetime_start_time_day_of_month = 21` + "\n"
	config += `		accept_lifetime_start_time_month = "january"` + "\n"
	config += `		accept_lifetime_start_time_year = 2023` + "\n"
	config += `		accept_lifetime_infinite = true` + "\n"
	config += `		send_lifetime_start_time_hour = 8` + "\n"
	config += `		send_lifetime_start_time_minute = 36` + "\n"
	config += `		send_lifetime_start_time_second = 22` + "\n"
	config += `		send_lifetime_start_time_day_of_month = 15` + "\n"
	config += `		send_lifetime_start_time_month = "january"` + "\n"
	config += `		send_lifetime_start_time_year = 2023` + "\n"
	config += `		send_lifetime_infinite = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
