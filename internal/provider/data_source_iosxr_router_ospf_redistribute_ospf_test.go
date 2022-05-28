// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIosxrRouterOSPFRedistributeOSPF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterOSPFRedistributeOSPFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.iosxr_router_ospf_redistribute_ospf.test", "match_internal", "true"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf_redistribute_ospf.test", "match_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf_redistribute_ospf.test", "match_nssa_external", "false"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf_redistribute_ospf.test", "tag", "4"),
					resource.TestCheckResourceAttr("data.iosxr_router_ospf_redistribute_ospf.test", "metric_type", "1"),
				),
			},
		},
	})
}

const testAccDataSourceIosxrRouterOSPFRedistributeOSPFConfig = `

resource "iosxr_router_ospf_redistribute_ospf" "test" {
	process_name = "OSPF1"
	instance_name = "OSPF2"
	match_internal = true
	match_external = false
	match_nssa_external = false
	tag = 4
	metric_type = "1"
}

data "iosxr_router_ospf_redistribute_ospf" "test" {
	process_name = "OSPF1"
	instance_name = "OSPF2"
	depends_on = [iosxr_router_ospf_redistribute_ospf.test]
}
`
