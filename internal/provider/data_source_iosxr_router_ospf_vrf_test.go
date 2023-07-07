// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrRouterOSPFVRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "mpls_ldp_sync", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "hello_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "dead_interval", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "priority", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "mtu_ignore_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "mtu_ignore_disable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "passive_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "passive_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "router_id", "10.11.12.13"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_connected_tag", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_connected_metric_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_static", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_static_tag", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_static_metric_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "bfd_fast_detect", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "bfd_minimum_interval", "300"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "bfd_multiplier", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "default_information_originate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "default_information_originate_always", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "default_information_originate_metric_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "areas.0.area_id", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_bgp.0.as_number", "65001"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_bgp.0.tag", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_bgp.0.metric_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_isis.0.instance_name", "P1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_isis.0.level_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_isis.0.level_2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_isis.0.level_1_2", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_isis.0.tag", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_isis.0.metric_type", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_ospf.0.instance_name", "OSPF2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_ospf.0.match_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_ospf.0.match_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_ospf.0.match_nssa_external", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_ospf.0.tag", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_router_ospf_vrf.test", "redistribute_ospf.0.metric_type", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrRouterOSPFVRFConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrRouterOSPFVRFConfig() string {
	config := `resource "iosxr_router_ospf_vrf" "test" {` + "\n"
	config += `	delete_mode = "attributes"` + "\n"
	config += `	process_name = "OSPF1"` + "\n"
	config += `	vrf_name = "VRF1"` + "\n"
	config += `	mpls_ldp_sync = false` + "\n"
	config += `	hello_interval = 10` + "\n"
	config += `	dead_interval = 40` + "\n"
	config += `	priority = 10` + "\n"
	config += `	mtu_ignore_enable = true` + "\n"
	config += `	mtu_ignore_disable = false` + "\n"
	config += `	passive_enable = false` + "\n"
	config += `	passive_disable = true` + "\n"
	config += `	router_id = "10.11.12.13"` + "\n"
	config += `	redistribute_connected = true` + "\n"
	config += `	redistribute_connected_tag = 1` + "\n"
	config += `	redistribute_connected_metric_type = "1"` + "\n"
	config += `	redistribute_static = true` + "\n"
	config += `	redistribute_static_tag = 2` + "\n"
	config += `	redistribute_static_metric_type = "1"` + "\n"
	config += `	bfd_fast_detect = true` + "\n"
	config += `	bfd_minimum_interval = 300` + "\n"
	config += `	bfd_multiplier = 3` + "\n"
	config += `	default_information_originate = true` + "\n"
	config += `	default_information_originate_always = true` + "\n"
	config += `	default_information_originate_metric_type = 1` + "\n"
	config += `	areas = [{` + "\n"
	config += `		area_id = "0"` + "\n"
	config += `	}]` + "\n"
	config += `	redistribute_bgp = [{` + "\n"
	config += `		as_number = "65001"` + "\n"
	config += `		tag = 3` + "\n"
	config += `		metric_type = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	redistribute_isis = [{` + "\n"
	config += `		instance_name = "P1"` + "\n"
	config += `		level_1 = true` + "\n"
	config += `		level_2 = false` + "\n"
	config += `		level_1_2 = false` + "\n"
	config += `		tag = 3` + "\n"
	config += `		metric_type = "1"` + "\n"
	config += `	}]` + "\n"
	config += `	redistribute_ospf = [{` + "\n"
	config += `		instance_name = "OSPF2"` + "\n"
	config += `		match_internal = true` + "\n"
	config += `		match_external = false` + "\n"
	config += `		match_nssa_external = false` + "\n"
	config += `		tag = 4` + "\n"
	config += `		metric_type = "1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_router_ospf_vrf" "test" {
			process_name = "OSPF1"
			vrf_name = "VRF1"
			depends_on = [iosxr_router_ospf_vrf.test]
		}
	`
	return config
}
