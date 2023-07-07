// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterStaticIPv4Unicast(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "prefix_address", "100.0.1.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "prefix_length", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interfaces.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interfaces.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interfaces.0.distance_metric", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interfaces.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interfaces.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.address", "11.11.11.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.tag", "103"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.distance_metric", "144"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_interface_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_addresses.0.address", "100.0.2.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_addresses.0.description", "ip-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_addresses.0.tag", "104"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_addresses.0.distance_metric", "155"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_addresses.0.track", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "nexthop_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.vrf_name", "VRF1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interfaces.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interfaces.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interfaces.0.distance_metric", "122"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interfaces.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interfaces.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.address", "11.11.11.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.description", "interface-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.tag", "103"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.distance_metric", "144"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.permanent", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_interface_addresses.0.metric", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_addresses.0.address", "100.0.2.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_addresses.0.description", "ip-description"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_addresses.0.tag", "104"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_addresses.0.distance_metric", "155"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_addresses.0.track", "TRACK1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_router_static_ipv4_unicast.test", "vrfs.0.nexthop_addresses.0.metric", "10"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterStaticIPv4UnicastConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_router_static_ipv4_unicast.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-static-cfg:/router/static/address-family/ipv4/unicast/prefixes/prefix[prefix-address=100.0.1.0][prefix-length=%!d(string=24)]",
			},
		},
	})
}

func testAccIosxrRouterStaticIPv4UnicastConfig_minimum() string {
	config := `resource "iosxr_router_static_ipv4_unicast" "test" {` + "\n"
	config += `	prefix_address = "100.0.1.0"` + "\n"
	config += `	prefix_length = 24` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrRouterStaticIPv4UnicastConfig_all() string {
	config := `resource "iosxe_router_static_ipv4_unicast" "test" {` + "\n"
	config += `	prefix_address = "100.0.1.0"` + "\n"
	config += `	prefix_length = 24` + "\n"
	config += `	nexthop_interfaces = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		description = "interface-description"` + "\n"
	config += `		tag = 100` + "\n"
	config += `		distance_metric = 122` + "\n"
	config += `		permanent = true` + "\n"
	config += `		metric = 10` + "\n"
	config += `	}]` + "\n"
	config += `	nexthop_interface_addresses = [{` + "\n"
	config += `		interface_name = "GigabitEthernet0/0/0/2"` + "\n"
	config += `		address = "11.11.11.1"` + "\n"
	config += `		description = "interface-description"` + "\n"
	config += `		tag = 103` + "\n"
	config += `		distance_metric = 144` + "\n"
	config += `		permanent = true` + "\n"
	config += `		metric = 10` + "\n"
	config += `	}]` + "\n"
	config += `	nexthop_addresses = [{` + "\n"
	config += `		address = "100.0.2.0"` + "\n"
	config += `		description = "ip-description"` + "\n"
	config += `		tag = 104` + "\n"
	config += `		distance_metric = 155` + "\n"
	config += `		track = "TRACK1"` + "\n"
	config += `		metric = 10` + "\n"
	config += `	}]` + "\n"
	config += `	vrfs = [{` + "\n"
	config += `		vrf_name = "VRF1"` + "\n"
	config += `		nexthop_interfaces = [{` + "\n"
	config += `			interface_name = "GigabitEthernet0/0/0/3"` + "\n"
	config += `			description = "interface-description"` + "\n"
	config += `			tag = 100` + "\n"
	config += `			distance_metric = 122` + "\n"
	config += `			permanent = true` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		nexthop_interface_addresses = [{` + "\n"
	config += `			interface_name = "GigabitEthernet0/0/0/4"` + "\n"
	config += `			address = "11.11.11.1"` + "\n"
	config += `			description = "interface-description"` + "\n"
	config += `			tag = 103` + "\n"
	config += `			distance_metric = 144` + "\n"
	config += `			permanent = true` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `		nexthop_addresses = [{` + "\n"
	config += `			address = "100.0.2.0"` + "\n"
	config += `			description = "ip-description"` + "\n"
	config += `			tag = 104` + "\n"
	config += `			distance_metric = 155` + "\n"
	config += `			track = "TRACK1"` + "\n"
	config += `			metric = 10` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
