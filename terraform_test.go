package aws

import (
	"fmt"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestS3BucketandFiles(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// assert bucket was created
	aws.AssertS3BucketExists(t, "us-east-1", os.Getenv("TF_VAR_bucket"))

	// getting files
	var file1 = aws.GetS3ObjectContents(t, os.Getenv("TF_VAR_region"), os.Getenv("TF_VAR_bucket"), "test1.txt")
	fmt.Println(file1)
	var file2 = aws.GetS3ObjectContents(t, os.Getenv("TF_VAR_region"), os.Getenv("TF_VAR_bucket"), "test2.txt")
	fmt.Println(file2)
}
