// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrFPD(t *testing.T) {
	if os.Getenv("FPD") == "" {
		t.Skip("skipping test, set environment variable FPD")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_fpd.test", "auto_upgrade_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_fpd.test", "auto_upgrade_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_fpd.test", "auto_reload_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_fpd.test", "auto_reload_disable", "false"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrFPDConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrFPDConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_fpd.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-fpd-cfg:/fpd",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrFPDConfig_minimum() string {
	config := `resource "iosxr_fpd" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrFPDConfig_all() string {
	config := `resource "iosxr_fpd" "test" {` + "\n"
	config += `	auto_upgrade_enable = false` + "\n"
	config += `	auto_upgrade_disable = false` + "\n"
	config += `	auto_reload_enable = false` + "\n"
	config += `	auto_reload_disable = false` + "\n"
	config += `}` + "\n"
	return config
}
