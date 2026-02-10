//sai 
package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEC2VPC(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
    TerraformDir: "../example/basic",

    VarFiles: []string{
        "ec2_test.tfvars",
    },

}

	// Destroy infra after test finishes
	defer terraform.Destroy(t, terraformOptions)

	// Init and Apply Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Fetch Terraform outputs
	instanceID := terraform.Output(t, terraformOptions, "instance_id")

	// Validate
	assert.NotEmpty(t, instanceID)
}