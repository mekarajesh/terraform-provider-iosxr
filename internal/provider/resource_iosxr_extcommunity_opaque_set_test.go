// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrExtcommunityOpaqueSet(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_extcommunity_opaque_set.test", "set_name", "BLUE"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_extcommunity_opaque_set.test", "rpl", "extcommunity-set opaque BLUE\n  100\nend-set\n"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrExtcommunityOpaqueSetConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_extcommunity_opaque_set.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/sets/extended-community-opaque-sets/extended-community-opaque-set[set-name=BLUE]",
			},
		},
	})
}

func testAccIosxrExtcommunityOpaqueSetConfig_minimum() string {
	config := `resource "iosxr_extcommunity_opaque_set" "test" {` + "\n"
	config += `	set_name = "BLUE"` + "\n"
	config += `	rpl = "extcommunity-set opaque BLUE\n  100\nend-set\n"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrExtcommunityOpaqueSetConfig_all() string {
	config := `resource "iosxe_extcommunity_opaque_set" "test" {` + "\n"
	config += `	set_name = "BLUE"` + "\n"
	config += `	rpl = "extcommunity-set opaque BLUE\n  100\nend-set\n"` + "\n"
	config += `}` + "\n"
	return config
}
