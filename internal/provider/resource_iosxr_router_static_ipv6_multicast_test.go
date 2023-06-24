// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrRouterStaticIPv6Multicast(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrRouterStaticIPv6MulticastConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "prefix_address", "1::"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "prefix_length", "64"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/1"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interfaces.0.description", "interface-description"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interfaces.0.tag", "100"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interfaces.0.distance_metric", "122"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interfaces.0.permanent", "true"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interfaces.0.metric", "10"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/2"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.address", "2::2"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.description", "interface-description"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.tag", "103"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.distance_metric", "144"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.permanent", "true"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_interface_addresses.0.metric", "10"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_addresses.0.address", "3::3"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_addresses.0.description", "ip-description"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_addresses.0.tag", "104"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_addresses.0.distance_metric", "155"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_addresses.0.track", "TRACK1"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "nexthop_addresses.0.metric", "10"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interfaces.0.interface_name", "GigabitEthernet0/0/0/3"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interfaces.0.description", "interface-description"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interfaces.0.tag", "100"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interfaces.0.distance_metric", "122"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interfaces.0.permanent", "true"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interfaces.0.metric", "10"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.interface_name", "GigabitEthernet0/0/0/4"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.address", "2::2"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.description", "interface-description"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.tag", "103"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.distance_metric", "144"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.permanent", "true"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_interface_addresses.0.metric", "10"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_addresses.0.address", "3::3"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_addresses.0.description", "ip-description"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_addresses.0.tag", "104"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_addresses.0.distance_metric", "155"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_addresses.0.track", "TRACK1"),
					resource.TestCheckResourceAttr("iosxr_router_static_ipv6_multicast.test", "vrfs.0.nexthop_addresses.0.metric", "10"),
				),
			},
			{
				ResourceName:  "iosxr_router_static_ipv6_multicast.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-router-static-cfg:/router/static/address-family/ipv6/multicast/prefixes/prefix[prefix-address=1::][prefix-length=%!d(string=64)]",
			},
		},
	})
}

func testAccIosxrRouterStaticIPv6MulticastConfig_minimum() string {
	return `
	resource "iosxr_router_static_ipv6_multicast" "test" {
		prefix_address = "1::"
		prefix_length = 64
	}
	`
}

func testAccIosxrRouterStaticIPv6MulticastConfig_all() string {
	return `
	resource "iosxr_router_static_ipv6_multicast" "test" {
		prefix_address = "1::"
		prefix_length = 64
		nexthop_interfaces = [{
			interface_name = "GigabitEthernet0/0/0/1"
			description = "interface-description"
			tag = 100
			distance_metric = 122
			permanent = true
			metric = 10
		}]
		nexthop_interface_addresses = [{
			interface_name = "GigabitEthernet0/0/0/2"
			address = "2::2"
			description = "interface-description"
			tag = 103
			distance_metric = 144
			permanent = true
			metric = 10
		}]
		nexthop_addresses = [{
			address = "3::3"
			description = "ip-description"
			tag = 104
			distance_metric = 155
			track = "TRACK1"
			metric = 10
		}]
		vrfs = [{
			vrf_name = "VRF1"
				nexthop_interfaces = [{
					interface_name = "GigabitEthernet0/0/0/3"
					description = "interface-description"
					tag = 100
					distance_metric = 122
					permanent = true
					metric = 10
				}]
				nexthop_interface_addresses = [{
					interface_name = "GigabitEthernet0/0/0/4"
					address = "2::2"
					description = "interface-description"
					tag = 103
					distance_metric = 144
					permanent = true
					metric = 10
				}]
				nexthop_addresses = [{
					address = "3::3"
					description = "ip-description"
					tag = 104
					distance_metric = 155
					track = "TRACK1"
					metric = 10
				}]
		}]
	}
	`
}
