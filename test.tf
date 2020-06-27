variable "bucket" {
  default = "aviltest"
}

variable "region" {
  default = "us-east-1"
}

provider "aws" {
  region = var.region
  access_key = "AKIAWVIUOJZXDIRPSJU2"
  secret_key = "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"
}

resource "aws_s3_bucket" "b3" {
  bucket = var.bucket
  acl    = "private"

  tags = {
    Name  = var.bucket
    Environment = "dev"
  }
}

resource "local_file" "file1" {
    content    = timestamp()
    filename = "${path.module}/test1.txt"
}

resource "local_file" "file2" {
    content    = timestamp()
    filename = "${path.module}/test2.txt"
}


resource "aws_s3_bucket_object" "object1" {
  bucket = aws_s3_bucket.b3.bucket
  key    = "test1.txt"
  source = "test1.txt"
}

resource "aws_s3_bucket_object" "object2" {
  bucket = aws_s3_bucket.b3.bucket
  key    = "test2.txt"
  source = "test2.txt"
}

