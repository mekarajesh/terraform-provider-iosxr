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

func TestAccDataSourceIosxrRouterISISAddressFamily(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_narrow", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_wide", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_transition", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.level_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.narrow", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.wide", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "metric_style_levels.0.transition", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "router_id_ip_address", "1050:0000:0000:0000:0005:0600:300c:326b"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "default_information_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_delay_interval", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_link_priority_limit_critical", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_link_priority_limit_high", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_link_priority_limit_medium", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_prefix_priority_limit_critical", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_prefix_priority_limit_high", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "fast_reroute_per_prefix_priority_limit_medium", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "microloop_avoidance_protected", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "microloop_avoidance_segment_routing", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "advertise_passive_only", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "advertise_link_attributes", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_ldp_auto_config", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_traffic_eng_level_1_2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "mpls_traffic_eng_level_1", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_interval_maximum_wait", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_interval_initial_wait", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_interval_secondary_wait", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_prefix_priorities.0.priority", "critical"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "spf_prefix_priorities.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "maximum_redistributed_prefixes", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "maximum_redistributed_prefixes_levels.0.level_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "maximum_redistributed_prefixes_levels.0.maximum_prefixes", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "redistribute_isis.0.instance_id", "CORE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "redistribute_isis.0.route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "segment_routing_srv6_locators.0.locator_name", "AlgoLocator"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_isis_address_family.test", "segment_routing_srv6_locators.0.level", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterISISAddressFamilyPrerequisitesConfig + testAccDataSourceIosxrRouterISISAddressFamilyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxrRouterISISAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
	attributes = {
		"route-policy-name" = "ROUTE_POLICY_1"
		"rpl-route-policy" = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
	}
}

`

func testAccDataSourceIosxrRouterISISAddressFamilyConfig() string {
	config := `resource "iosxr_router_isis_address_family" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	process_id = "P1"` + "\n"
	config += `	af_name = "ipv6"` + "\n"
	config += `	saf_name = "unicast"` + "\n"
	config += `	metric_style_narrow = false` + "\n"
	config += `	metric_style_wide = true` + "\n"
	config += `	metric_style_transition = false` + "\n"
	config += `	metric_style_levels = [{` + "\n"
	config += `		level_id = 1` + "\n"
	config += `		narrow = false` + "\n"
	config += `		wide = true` + "\n"
	config += `		transition = false` + "\n"
	config += `	}]` + "\n"
	config += `	router_id_ip_address = "1050:0000:0000:0000:0005:0600:300c:326b"` + "\n"
	config += `	default_information_originate = true` + "\n"
	config += `	fast_reroute_delay_interval = 300` + "\n"
	config += `	fast_reroute_per_link_priority_limit_critical = true` + "\n"
	config += `	fast_reroute_per_link_priority_limit_high = false` + "\n"
	config += `	fast_reroute_per_link_priority_limit_medium = false` + "\n"
	config += `	fast_reroute_per_prefix_priority_limit_critical = true` + "\n"
	config += `	fast_reroute_per_prefix_priority_limit_high = false` + "\n"
	config += `	fast_reroute_per_prefix_priority_limit_medium = false` + "\n"
	config += `	microloop_avoidance_protected = false` + "\n"
	config += `	microloop_avoidance_segment_routing = true` + "\n"
	config += `	advertise_passive_only = true` + "\n"
	config += `	advertise_link_attributes = true` + "\n"
	config += `	mpls_ldp_auto_config = false` + "\n"
	config += `	mpls_traffic_eng_level_1_2 = false` + "\n"
	config += `	mpls_traffic_eng_level_1 = false` + "\n"
	config += `	spf_interval_maximum_wait = 5000` + "\n"
	config += `	spf_interval_initial_wait = 50` + "\n"
	config += `	spf_interval_secondary_wait = 200` + "\n"
	config += `	spf_prefix_priorities = [{` + "\n"
	config += `		priority = "critical"` + "\n"
	config += `		tag = 100` + "\n"
	config += `	}]` + "\n"
	config += `	maximum_redistributed_prefixes = 100` + "\n"
	config += `	maximum_redistributed_prefixes_levels = [{` + "\n"
	config += `		level_id = 1` + "\n"
	config += `		maximum_prefixes = 1000` + "\n"
	config += `	}]` + "\n"
	config += `	redistribute_isis = [{` + "\n"
	config += `		instance_id = "CORE"` + "\n"
	config += `		route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	}]` + "\n"
	config += `	segment_routing_srv6_locators = [{` + "\n"
	config += `		locator_name = "AlgoLocator"` + "\n"
	config += `		level = 1` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_router_isis_address_family" "test" {
			process_id = "P1"
			af_name = "ipv6"
			saf_name = "unicast"
			depends_on = [iosxr_router_isis_address_family.test]
		}
	`
	return config
}
