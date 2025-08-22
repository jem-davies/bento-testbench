provider "aws" {
  region = "eu-west-2"
}

resource "random_string" "bucket_suffix" {
  length  = 6
  special = false
  upper   = false
}

resource "aws_s3_bucket" "large_message_bucket" {
  bucket = "large-message-bucket-${random_string.bucket_suffix.result}"
  force_destroy = true
}

resource "aws_sqs_queue" "large_message_queue" {
  name = "large-message-queue"
}

output "bucket_name" {
  value = aws_s3_bucket.large_message_bucket.id
}

output "queue_url" {
  value = aws_sqs_queue.large_message_queue.url
}
