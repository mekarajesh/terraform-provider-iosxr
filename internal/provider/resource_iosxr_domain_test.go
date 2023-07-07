// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrDomain(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "domains.0.domain_name", "DOMAIN1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "domains.0.order", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "lookup_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "lookup_source_interface", "Loopback2147483647"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "name", "DOMAIN"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "ipv4_hosts.0.host_name", "HOST_NAME"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "ipv4_hosts.0.ip_address.0", "10.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "name_servers.0.address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "name_servers.0.order", "345"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "ipv6_hosts.0.host_name", "HOST_NAME_IPV6"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "ipv6_hosts.0.ipv6_address.0", "10::10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "multicast", "DOMAIN1_ACC"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain.test", "default_flows_disable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIosxrDomainConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				ResourceName:  "iosxr_domain.test",
				ImportState:   true,
				ImportStateId: "Cisco-IOS-XR-um-domain-cfg:/domain",
			},
		},
	})
}

func testAccIosxrDomainConfig_minimum() string {
	config := `resource "iosxr_domain" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrDomainConfig_all() string {
	config := `resource "iosxe_domain" "test" {` + "\n"
	config += `	domains = [{` + "\n"
	config += `		domain_name = "DOMAIN1"` + "\n"
	config += `		order = 0` + "\n"
	config += `	}]` + "\n"
	config += `	lookup_disable = true` + "\n"
	config += `	lookup_source_interface = "Loopback2147483647"` + "\n"
	config += `	name = "DOMAIN"` + "\n"
	config += `	ipv4_hosts = [{` + "\n"
	config += `		host_name = "HOST_NAME"` + "\n"
	config += `		ip_address = ["10.0.0.0"]` + "\n"
	config += `	}]` + "\n"
	config += `	name_servers = [{` + "\n"
	config += `		address = "10.0.0.1"` + "\n"
	config += `		order = 345` + "\n"
	config += `	}]` + "\n"
	config += `	ipv6_hosts = [{` + "\n"
	config += `		host_name = "HOST_NAME_IPV6"` + "\n"
	config += `		ipv6_address = ["10::10"]` + "\n"
	config += `	}]` + "\n"
	config += `	multicast = "DOMAIN1_ACC"` + "\n"
	config += `	default_flows_disable = true` + "\n"
	config += `}` + "\n"
	return config
}
