// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrL2VPNBridgeGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_l2vpn_bridge_group.test", "group_name", "BG123"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrL2VPNBridgeGroupConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_l2vpn_bridge_group.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-l2vpn-cfg:/l2vpn/bridge/groups/group[group-name=BG123]",
			},
		},
	})
}

func testAccIosxrL2VPNBridgeGroupConfig_minimum() string {
	config := `resource "iosxr_l2vpn_bridge_group" "test" {` + "\n"
	config += `	group_name = "BG123"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrL2VPNBridgeGroupConfig_all() string {
	config := `resource "iosxe_l2vpn_bridge_group" "test" {` + "\n"
	config += `	group_name = "BG123"` + "\n"
	config += `}` + "\n"
	return config
}
