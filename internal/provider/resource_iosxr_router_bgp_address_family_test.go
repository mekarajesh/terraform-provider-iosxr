// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterBGPAddressFamily(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "af_name", "ipv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "additional_paths_send", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "additional_paths_receive", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "additional_paths_selection_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "advertise_best_external", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "allocate_label_all", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "label_mode_per_ce", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "label_mode_per_vrf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_connected_metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_static_metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "aggregate_addresses.0.address", "10.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "aggregate_addresses.0.masklength", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "aggregate_addresses.0.as_set", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "aggregate_addresses.0.as_confed_set", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "aggregate_addresses.0.summary_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "networks.0.address", "10.1.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "networks.0.masklength", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.instance_name", "P1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_one", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_one_two", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_one_two_one_inter_area", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_one_one_inter_area", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_two", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_two_one_inter_area", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.level_one_inter_area", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_isis.0.metric", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.router_tag", "OSPF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.match_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.match_internal_external", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.match_internal_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.match_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.match_external_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.match_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_address_family.test", "redistribute_ospf.0.metric", "100"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrRouterBGPAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPAddressFamilyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrRouterBGPAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPAddressFamilyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_router_bgp_address_family.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/address-families/address-family[af-name=ipv4-unicast]",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

const testAccIosxrRouterBGPAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
	attributes = {
		"route-policy-name" = "ROUTE_POLICY_1"
		"rpl-route-policy" = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
	}
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
}

`

func testAccIosxrRouterBGPAddressFamilyConfig_minimum() string {
	config := `resource "iosxr_router_bgp_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	af_name = "ipv4-unicast"` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterBGPAddressFamilyConfig_all() string {
	config := `resource "iosxr_router_bgp_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	af_name = "ipv4-unicast"` + "\n"
	config += `	additional_paths_send = true` + "\n"
	config += `	additional_paths_receive = true` + "\n"
	config += `	additional_paths_selection_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	advertise_best_external = true` + "\n"
	config += `	allocate_label_all = true` + "\n"
	config += `	label_mode_per_ce = false` + "\n"
	config += `	label_mode_per_vrf = false` + "\n"
	config += `	redistribute_connected = true` + "\n"
	config += `	redistribute_connected_metric = 10` + "\n"
	config += `	redistribute_static = true` + "\n"
	config += `	redistribute_static_metric = 10` + "\n"
	config += `	aggregate_addresses = [{` + "\n"
	config += `		address = "10.0.0.0"` + "\n"
	config += `		masklength = 8` + "\n"
	config += `		as_set = false` + "\n"
	config += `		as_confed_set = false` + "\n"
	config += `		summary_only = false` + "\n"
	config += `	}]` + "\n"
	config += `	networks = [{` + "\n"
	config += `		address = "10.1.0.0"` + "\n"
	config += `		masklength = 16` + "\n"
	config += `	}]` + "\n"
	config += `	redistribute_isis = [{` + "\n"
	config += `		instance_name = "P1"` + "\n"
	config += `		level_one = true` + "\n"
	config += `		level_one_two = true` + "\n"
	config += `		level_one_two_one_inter_area = false` + "\n"
	config += `		level_one_one_inter_area = false` + "\n"
	config += `		level_two = false` + "\n"
	config += `		level_two_one_inter_area = false` + "\n"
	config += `		level_one_inter_area = false` + "\n"
	config += `		metric = 100` + "\n"
	config += `	}]` + "\n"
	config += `	redistribute_ospf = [{` + "\n"
	config += `		router_tag = "OSPF1"` + "\n"
	config += `		match_internal = true` + "\n"
	config += `		match_internal_external = true` + "\n"
	config += `		match_internal_nssa_external = false` + "\n"
	config += `		match_external = false` + "\n"
	config += `		match_external_nssa_external = false` + "\n"
	config += `		match_nssa_external = false` + "\n"
	config += `		metric = 100` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}
