terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
}
provider "aws" {
  region = "us-east-1"
}

data "aws_caller_identity" "current" {}


data "archive_file" "lambda_zip" {
  type        = "zip"
  source_dir  = "resources/lambda/react"
  output_path = "resources/lambda/out/reactapp.zip"
  depends_on  = [aws_s3_bucket_object.upload_folder_prod]
}

resource "aws_lambda_function" "react_lambda_app" {
  filename      = "resources/lambda/out/reactapp.zip"
  function_name = "blog-application"
  handler       = "index.handler"
  runtime       = "nodejs14.x"
  role          = aws_iam_role.blog_app_lambda.arn
  depends_on    = [data.archive_file.lambda_zip, null_resource.file_replacement_lambda_react]
}

resource "aws_iam_role" "blog_app_lambda" {
  name = "blog_app_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}