// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrLogging(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "ipv4_dscp", "cs6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "trap", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "events_display_location", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "events_level", "informational"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "console", "disable"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "monitor", "disable"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "buffered_logging_buffer_size", "4000000"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "buffered_level", "debugging"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "facility_level", "local7"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "hostnameprefix", "HOSTNAME01"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_logging.test", "suppress_duplicates", "true"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrLoggingConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrLoggingConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_logging.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-logging-cfg:/logging",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrLoggingConfig_minimum() string {
	config := `resource "iosxr_logging" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrLoggingConfig_all() string {
	config := `resource "iosxr_logging" "test" {` + "\n"
	config += `	ipv4_dscp = "cs6"` + "\n"
	config += `	trap = "informational"` + "\n"
	config += `	events_display_location = true` + "\n"
	config += `	events_level = "informational"` + "\n"
	config += `	console = "disable"` + "\n"
	config += `	monitor = "disable"` + "\n"
	config += `	buffered_logging_buffer_size = 4000000` + "\n"
	config += `	buffered_level = "debugging"` + "\n"
	config += `	facility_level = "local7"` + "\n"
	config += `	hostnameprefix = "HOSTNAME01"` + "\n"
	config += `	suppress_duplicates = true` + "\n"
	config += `}` + "\n"
	return config
}
