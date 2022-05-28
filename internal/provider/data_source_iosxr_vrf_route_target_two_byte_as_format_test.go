// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrVRFRouteTargetTwoByteASFormat(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrVRFRouteTargetTwoByteASFormatConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceIosxrVRFRouteTargetTwoByteASFormatConfig = `

resource "iosxr_vrf_route_target_two_byte_as_format" "test" {
	vrf_name = "VRF1"
	address_family = "ipv4"
	sub_address_family = "unicast"
	direction = "import"
	as_number = 1
	index = 1
	stitching = true
}

data "iosxr_vrf_route_target_two_byte_as_format" "test" {
	vrf_name = "VRF1"
	address_family = "ipv4"
	sub_address_family = "unicast"
	direction = "import"
	as_number = 1
	index = 1
	stitching = true
	depends_on = [iosxr_vrf_route_target_two_byte_as_format.test]
}
`
