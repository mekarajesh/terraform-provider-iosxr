// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIosxrEVPNInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrEVPNInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "interface_name", "Bundle-Ether12"),
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "core_isolation_group", "11"),
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "ethernet_segment_identifier_type_zero_esi", "01.00.01.01.00.00.00.01.1"),
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_all_active", "false"),
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_port_active", "false"),
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_single_active", "true"),
					resource.TestCheckResourceAttr("iosxr_evpn_interface.test", "ethernet_segment_load_balancing_mode_single_flow_active", "false"),
				),
			},
			{
				ResourceName:  "iosxr_evpn_interface.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:evpn/interface/interface[interface-name=Bundle-Ether12]",
			},
		},
	})
}

func testAccIosxrEVPNInterfaceConfig_minimum() string {
	return `
	resource "iosxr_evpn_interface" "test" {
		interface_name = "Bundle-Ether12"
	}
	`
}

func testAccIosxrEVPNInterfaceConfig_all() string {
	return `
	resource "iosxr_evpn_interface" "test" {
		interface_name = "Bundle-Ether12"
		core_isolation_group = 11
		ethernet_segment_identifier_type_zero_esi = "01.00.01.01.00.00.00.01.1"
		ethernet_segment_load_balancing_mode_all_active = false
		ethernet_segment_load_balancing_mode_port_active = false
		ethernet_segment_load_balancing_mode_single_active = true
		ethernet_segment_load_balancing_mode_single_flow_active = false
	}
	`
}
