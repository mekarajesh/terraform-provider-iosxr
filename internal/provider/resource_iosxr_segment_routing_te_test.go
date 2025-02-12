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

func TestAccIosxrSegmentRoutingTE(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "logging_pcep_peer_status", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "logging_policy_status", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_report_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_source_address", "88.88.88.8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_delegation_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_dead_timer", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_initiated_state", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pcc_initiated_orphan", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pce_peers.0.pce_address", "66.66.66.6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "pce_peers.0.precedence", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.dynamic_anycast_sid_inclusion", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.dynamic_metric_type", "te"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.color", "266"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.srv6_locator_name", "LOC11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.srv6_locator_behavior", "ub6-insert-reduced"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.srv6_locator_binding_sid_type", "srv6-dynamic"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.source_address", "fccc:0:213::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.source_address_type", "end-point-type-ipv6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.effective_metric_value", "4444"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.effective_metric_type", "igp"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.constraint_segments_protection_type", "protected-only"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "on_demand_colors.0.constraint_segments_sid_algorithm", "128"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.srv6_locator_name", "Locator11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.srv6_locator_binding_sid_type", "srv6-dynamic"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.srv6_locator_behavior", "ub6-insert-reduced"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.source_address", "fccc:0:103::1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.source_address_type", "end-point-type-ipv6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_color_endpoint_color", "65534"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_color_endpoint_type", "end-point-type-ipv6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_segment_routing_te.test", "policies.0.policy_color_endpoint_address", "fccc:0:215::1"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrSegmentRoutingTEConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrSegmentRoutingTEConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_segment_routing_te.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-segment-routing-ms-cfg:/sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrSegmentRoutingTEConfig_minimum() string {
	config := `resource "iosxr_segment_routing_te" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrSegmentRoutingTEConfig_all() string {
	config := `resource "iosxr_segment_routing_te" "test" {` + "\n"
	config += `	logging_pcep_peer_status = true` + "\n"
	config += `	logging_policy_status = true` + "\n"
	config += `	pcc_report_all = true` + "\n"
	config += `	pcc_source_address = "88.88.88.8"` + "\n"
	config += `	pcc_delegation_timeout = 10` + "\n"
	config += `	pcc_dead_timer = 60` + "\n"
	config += `	pcc_initiated_state = 15` + "\n"
	config += `	pcc_initiated_orphan = 10` + "\n"
	config += `	pce_peers = [{` + "\n"
	config += `		pce_address = "66.66.66.6"` + "\n"
	config += `		precedence = 122` + "\n"
	config += `	}]` + "\n"
	config += `	on_demand_colors = [{` + "\n"
	config += `		dynamic_anycast_sid_inclusion = true` + "\n"
	config += `		dynamic_metric_type = "te"` + "\n"
	config += `		color = 266` + "\n"
	config += `		srv6_locator_name = "LOC11"` + "\n"
	config += `		srv6_locator_behavior = "ub6-insert-reduced"` + "\n"
	config += `		srv6_locator_binding_sid_type = "srv6-dynamic"` + "\n"
	config += `		source_address = "fccc:0:213::1"` + "\n"
	config += `		source_address_type = "end-point-type-ipv6"` + "\n"
	config += `		effective_metric_value = 4444` + "\n"
	config += `		effective_metric_type = "igp"` + "\n"
	config += `		constraint_segments_protection_type = "protected-only"` + "\n"
	config += `		constraint_segments_sid_algorithm = 128` + "\n"
	config += `	}]` + "\n"
	config += `	policies = [{` + "\n"
	config += `		policy_name = "POLICY1"` + "\n"
	config += `		srv6_locator_name = "Locator11"` + "\n"
	config += `		srv6_locator_binding_sid_type = "srv6-dynamic"` + "\n"
	config += `		srv6_locator_behavior = "ub6-insert-reduced"` + "\n"
	config += `		source_address = "fccc:0:103::1"` + "\n"
	config += `		source_address_type = "end-point-type-ipv6"` + "\n"
	config += `		policy_color_endpoint_color = 65534` + "\n"
	config += `		policy_color_endpoint_type = "end-point-type-ipv6"` + "\n"
	config += `		policy_color_endpoint_address = "fccc:0:215::1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
