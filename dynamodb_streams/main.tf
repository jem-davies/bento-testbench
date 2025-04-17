provider "aws" {
  region = "us-east-1"
}

resource "aws_dynamodb_table" "users" {
  name = "Users"
  billing_mode = "PROVISIONED"
  read_capacity = 10
  write_capacity = 5

  hash_key = "userId"
  attribute {
    name = "userId"
    type = "S"  # String data type
  }

  stream_enabled = true
  stream_view_type = "NEW_AND_OLD_IMAGES" 
}
