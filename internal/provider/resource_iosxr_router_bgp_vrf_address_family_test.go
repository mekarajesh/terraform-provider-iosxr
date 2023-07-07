// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterBGPVRFAddressFamily(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "af_name", "ipv4-unicast"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "maximum_paths_ebgp_multipath", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "maximum_paths_ibgp_multipath", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "label_mode_per_ce", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "label_mode_per_vrf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_connected_metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_static_metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "segment_routing_srv6_locator", "LocAlgo11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "segment_routing_srv6_alloc_mode_per_vrf", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.address", "10.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.masklength", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.as_set", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.as_confed_set", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "aggregate_addresses.0.summary_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "networks.0.address", "10.1.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "networks.0.masklength", "16"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.router_tag", "OSPF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_internal_external", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_internal_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_external_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.match_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_bgp_vrf_address_family.test", "redistribute_ospf.0.metric", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterBGPVRFAddressFamilyPrerequisitesConfig + testAccIosxrRouterBGPVRFAddressFamilyConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_router_bgp_vrf_address_family.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]/vrfs/vrf[vrf-name=VRF1]/address-families/address-family[af-name=ipv4-unicast]",
			},
		},
	})
}

const testAccIosxrRouterBGPVRFAddressFamilyPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-vrf-cfg:/vrfs/vrf[vrf-name=VRF1]/Cisco-IOS-XR-um-router-bgp-cfg:rd/Cisco-IOS-XR-um-router-bgp-cfg:two-byte-as"
	attributes = {
		"as-number" = "1"
		"index" = "1"
	}
}

resource "iosxr_gnmi" "PreReq1" {
	path = "Cisco-IOS-XR-um-router-bgp-cfg:/router/bgp/as[as-number=65001]"
	attributes = {
		"as-number" = "65001"
	}
	lists = [
		{
			name = "address-families/address-family"
			key = "af-name"
			items = [
				{
					"af-name" = "vpnv4-unicast"
				},
			] 
		},
	]
}

`

func testAccIosxrRouterBGPVRFAddressFamilyConfig_minimum() string {
	config := `resource "iosxr_router_bgp_vrf_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	af_name = "ipv4-unicast"` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, iosxr_gnmi.PreReq1, ]` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterBGPVRFAddressFamilyConfig_all() string {
	config := `resource "iosxe_router_bgp_vrf_address_family" "test" {` + "\n"
	config += `	as_number = "65001"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	af_name = "ipv4-unicast"` + "\n"
	config += `	maximum_paths_ebgp_multipath = 10` + "\n"
	config += `	maximum_paths_ibgp_multipath = 10` + "\n"
	config += `	label_mode_per_ce = false` + "\n"
	config += `	label_mode_per_vrf = false` + "\n"
	config += `	redistribute_connected = true` + "\n"
	config += `	redistribute_connected_metric = 10` + "\n"
	config += `	redistribute_static = true` + "\n"
	config += `	redistribute_static_metric = 10` + "\n"
	config += `	segment_routing_srv6_locator = "LocAlgo11"` + "\n"
	config += `	segment_routing_srv6_alloc_mode_per_vrf = true` + "\n"
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
