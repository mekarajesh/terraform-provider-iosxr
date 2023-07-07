// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrVRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "description", "My VRF Description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "vpn_id", "1000:1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_multicast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_flowspec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_policy", "ROUTE_POLICY_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_multicast", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_flowspec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "rd_two_byte_as_as_number", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "rd_two_byte_as_index", "123"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_two_byte_as_format.0.as_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_two_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_two_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_four_byte_as_format.0.as_number", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_four_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_four_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_ip_address_format.0.ip_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_ip_address_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_import_route_target_ip_address_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_two_byte_as_format.0.as_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_two_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_two_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_four_byte_as_format.0.as_number", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_four_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_four_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_ip_address_format.0.ip_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_ip_address_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv4_unicast_export_route_target_ip_address_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_two_byte_as_format.0.as_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_two_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_two_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_four_byte_as_format.0.as_number", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_four_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_four_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_ip_address_format.0.ip_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_ip_address_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_import_route_target_ip_address_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_two_byte_as_format.0.as_number", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_two_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_two_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_four_byte_as_format.0.as_number", "100000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_four_byte_as_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_four_byte_as_format.0.stitching", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_ip_address_format.0.ip_address", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_ip_address_format.0.index", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_vrf.test", "address_family_ipv6_unicast_export_route_target_ip_address_format.0.stitching", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrVRFPrerequisitesConfig + testAccDataSourceIosxrVRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

const testAccDataSourceIosxrVRFPrerequisitesConfig = `
resource "iosxr_gnmi" "PreReq0" {
	path = "Cisco-IOS-XR-um-route-policy-cfg:/routing-policy/route-policies/route-policy[route-policy-name=ROUTE_POLICY_1]"
	attributes = {
		"route-policy-name" = "ROUTE_POLICY_1"
		"rpl-route-policy" = "route-policy ROUTE_POLICY_1\n  pass\nend-policy\n"
	}
}

`

func testAccDataSourceIosxrVRFConfig() string {
	config := `resource "iosxr_vrf" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	vrf_name = "VRF3"` + "\n"
	config += `	description = "My VRF Description"` + "\n"
	config += `	vpn_id = "1000:1000"` + "\n"
	config += `	address_family_ipv4_unicast = true` + "\n"
	config += `	address_family_ipv4_unicast_import_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	address_family_ipv4_unicast_export_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	address_family_ipv4_multicast = true` + "\n"
	config += `	address_family_ipv4_flowspec = true` + "\n"
	config += `	address_family_ipv6_unicast = true` + "\n"
	config += `	address_family_ipv6_unicast_import_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	address_family_ipv6_unicast_export_route_policy = "ROUTE_POLICY_1"` + "\n"
	config += `	address_family_ipv6_multicast = true` + "\n"
	config += `	address_family_ipv6_flowspec = true` + "\n"
	config += `	rd_two_byte_as_as_number = "123"` + "\n"
	config += `	rd_two_byte_as_index = 123` + "\n"
	config += `	address_family_ipv4_unicast_import_route_target_two_byte_as_format = [{` + "\n"
	config += `		as_number = 1` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv4_unicast_import_route_target_four_byte_as_format = [{` + "\n"
	config += `		as_number = 100000` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv4_unicast_import_route_target_ip_address_format = [{` + "\n"
	config += `		ip_address = "1.1.1.1"` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv4_unicast_export_route_target_two_byte_as_format = [{` + "\n"
	config += `		as_number = 1` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv4_unicast_export_route_target_four_byte_as_format = [{` + "\n"
	config += `		as_number = 100000` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv4_unicast_export_route_target_ip_address_format = [{` + "\n"
	config += `		ip_address = "1.1.1.1"` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv6_unicast_import_route_target_two_byte_as_format = [{` + "\n"
	config += `		as_number = 1` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv6_unicast_import_route_target_four_byte_as_format = [{` + "\n"
	config += `		as_number = 100000` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv6_unicast_import_route_target_ip_address_format = [{` + "\n"
	config += `		ip_address = "1.1.1.1"` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv6_unicast_export_route_target_two_byte_as_format = [{` + "\n"
	config += `		as_number = 1` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv6_unicast_export_route_target_four_byte_as_format = [{` + "\n"
	config += `		as_number = 100000` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	address_family_ipv6_unicast_export_route_target_ip_address_format = [{` + "\n"
	config += `		ip_address = "1.1.1.1"` + "\n"
	config += `		index = 1` + "\n"
	config += `		stitching = true` + "\n"
	config += `	}]` + "\n"
	config += `	depends_on = [iosxr_gnmi.PreReq0, ]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_vrf" "test" {
			vrf_name = "VRF3"
			depends_on = [iosxr_vrf.test]
		}
	`
	return config
}
