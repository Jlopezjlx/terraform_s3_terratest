# terraform_s3_terratest


Terraform script which create a s3 bucket and add two file in it, these files have the timestamp when the script was executed.

To be able to run this, you must have:

- Go installed
- Terraform inatalled
- an aws account with and IAM user with full access to s3

## Steps to run it

### Locally

1- Set the env variables needed:

AWS_ACCESS_KEY_ID

AWS_SECRET_ACCESS_KEY

TF_VAR_bucket

TF_VAR_region

2- Execute the next command

go mod init terraform_test.go

terraform plan

terraform apply -auto-approve

terraform destroy

go test -v (just execute this line if you want to test the terraform script)

3- This also can be run using bash scripts

run_test_script.sh, execute this if you want to run the tests (you need to do steps 1)

run.sh, execute this if you want to execute the terraform script