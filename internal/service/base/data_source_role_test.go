package base_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-pingone/internal/acctest"
)

func TestAccRoleDataSource_ByNameFull(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("pingone_role.%s", resourceName)
	dataSourceFullName := fmt.Sprintf("data.%s", resourceFullName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckEnvironment(t) },
		ProviderFactories: acctest.ProviderFactories,
		ErrorCheck:        acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: testAccRoleDataSourceConfig_ByNameFull(resourceName, "Organization Admin"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "Organization Admin"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "description"),
				),
			},
			{
				Config: testAccRoleDataSourceConfig_ByNameFull(resourceName, "Environment Admin"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "Environment Admin"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "description"),
				),
			},
			{
				Config: testAccRoleDataSourceConfig_ByNameFull(resourceName, "Identity Data Admin"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "Identity Data Admin"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "description"),
				),
			},
			{
				Config: testAccRoleDataSourceConfig_ByNameFull(resourceName, "Client Application Developer"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "Client Application Developer"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "description"),
				),
			},
			{
				Config: testAccRoleDataSourceConfig_ByNameFull(resourceName, "Identity Data Read Only"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "Identity Data Read Only"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "description"),
				),
			},
			{
				Config: testAccRoleDataSourceConfig_ByNameFull(resourceName, "Configuration Read Only"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "Configuration Read Only"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "description"),
				),
			},
		},
	})
}

func testAccRoleDataSourceConfig_ByNameFull(resourceName, name string) string {
	return fmt.Sprintf(`
		%[1]s

data "pingone_role" "%[2]s" {
  name = "%[3]s"
}`, acctest.GenericSandboxEnvironment(), resourceName, name)
}
