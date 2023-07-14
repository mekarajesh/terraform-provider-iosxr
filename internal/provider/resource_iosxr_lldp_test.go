// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrLLDP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "holdtime", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "timer", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "reinit", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "subinterfaces_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "priorityaddr_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "extended_show_width_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "tlv_select_management_address_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "tlv_select_port_description_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "tlv_select_system_capabilities_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "tlv_select_system_description_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_lldp.test", "tlv_select_system_name_disable", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrLLDPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrLLDPConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_lldp.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-lldp-cfg:/lldp",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrLLDPConfig_minimum() string {
	config := `resource "iosxr_lldp" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrLLDPConfig_all() string {
	config := `resource "iosxr_lldp" "test" {` + "\n"
	config += `	holdtime = 50` + "\n"
	config += `	timer = 6` + "\n"
	config += `	reinit = 3` + "\n"
	config += `	subinterfaces_enable = true` + "\n"
	config += `	priorityaddr_enable = true` + "\n"
	config += `	extended_show_width_enable = true` + "\n"
	config += `	tlv_select_management_address_disable = true` + "\n"
	config += `	tlv_select_port_description_disable = true` + "\n"
	config += `	tlv_select_system_capabilities_disable = true` + "\n"
	config += `	tlv_select_system_description_disable = true` + "\n"
	config += `	tlv_select_system_name_disable = true` + "\n"
	config += `}` + "\n"
	return config
}
