// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrIPv4PrefixList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "prefix_list_name", "LIST1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.sequence_number", "4096"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.remark", "REMARK"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.permission", "deny"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.prefix", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.mask", "255.255.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.match_prefix_length_eq", "12"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.match_prefix_length_ge", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ipv4_prefix_list.test", "sequences.0.match_prefix_length_le", "32"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrIPv4PrefixListConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_ipv4_prefix_list.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-ipv4-prefix-list-cfg:/ipv4/prefix-lists/prefix-list[prefix-list-name=LIST1]",
			},
		},
	})
}

func testAccIosxrIPv4PrefixListConfig_minimum() string {
	config := `resource "iosxr_ipv4_prefix_list" "test" {` + "\n"
	config += `	prefix_list_name = "LIST1"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrIPv4PrefixListConfig_all() string {
	config := `resource "iosxe_ipv4_prefix_list" "test" {` + "\n"
	config += `	prefix_list_name = "LIST1"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `		sequence_number = 4096` + "\n"
	config += `		remark = "REMARK"` + "\n"
	config += `		permission = "deny"` + "\n"
	config += `		prefix = "10.1.1.1"` + "\n"
	config += `		mask = "255.255.0.0"` + "\n"
	config += `		match_prefix_length_eq = 12` + "\n"
	config += `		match_prefix_length_ge = 22` + "\n"
	config += `		match_prefix_length_le = 32` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}
