package provider

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSensitiveDataSource(t *testing.T) {
	randomValue := fmt.Sprintf("somerandom_%d", rand.Int())
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck: func() {
			os.Setenv("TF_ENV_MYTESTENVSENSITIVE", randomValue)
		},
		Steps: []resource.TestStep{
			{ // Read testing
				Config: `
data "env_sensitive" "test1" {
  id = "TF_ENV_MYTESTENVSENSITIVE"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_sensitive.test1", "id", "TF_ENV_MYTESTENVSENSITIVE"),
					resource.TestCheckResourceAttr("data.env_sensitive.test1", "value", randomValue),
				),
			},
			{ // required
				Config: `
data "env_sensitive" "test2" {
  id       = "TF_ENV_MYTESTENVSENSITIVE"
  required = true
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.env_sensitive.test2", "id", "TF_ENV_MYTESTENVSENSITIVE"),
					resource.TestCheckResourceAttr("data.env_sensitive.test2", "required", "true"),
					resource.TestCheckResourceAttr("data.env_sensitive.test2", "value", randomValue),
				),
			},
			{ // required and fails
				Config: `
data "env_sensitive" "test3" {
  id       = "TF_ENV_MY_NON_EXISTING_VAR"
  required = true
}
`,
				ExpectError: regexp.MustCompile("env \"TF_ENV_MY_NON_EXISTING_VAR\" not found"),
			},
		},
	})
}
