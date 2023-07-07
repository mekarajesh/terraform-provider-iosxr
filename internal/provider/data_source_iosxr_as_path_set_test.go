// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceIosxrASPathSet(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.iosxr_as_path_set.test", "rpl", "as-path-set TEST1\n  length ge 10\nend-set\n"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceIosxrASPathSetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccDataSourceIosxrASPathSetConfig() string {
	config := `resource "iosxr_as_path_set" "test" {` + "\n"
	config += `	set_name = "TEST1"` + "\n"
	config += `	rpl = "as-path-set TEST1\n  length ge 10\nend-set\n"` + "\n"
	config += `}` + "\n"

	config += `
		data "iosxr_as_path_set" "test" {
			set_name = "TEST1"
			depends_on = [iosxr_as_path_set.test]
		}
	`
	return config
}
