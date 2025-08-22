import os
import boto3  # pip install boto3
import sqs_extended_client  # pip install amazon-sqs-extended-client

# From Terraform Outputs
bucket_name = "large-message-bucket-ipjv5p"
queue_url = "https://sqs.<AWS_REGION>.amazonaws.com/<AWS_ACCOUNT_ID>/large-message-queue"

# Configure SQS extended client for large payloads
sqs_extended_client = boto3.client("sqs", region_name="eu-west-2")
sqs_extended_client.large_payload_support = bucket_name
sqs_extended_client.always_through_s3 = True
sqs_extended_client.use_legacy_attribute = False

# Folder containing input files
input_folder = "./input_data_same"

for filename in ["one.txt", "two.txt", "three.txt"]:
    filepath = os.path.join(input_folder, filename)
    
    with open(filepath, "rb") as f:  # read file as bytes
        file_content = f.read()

    # SQS requires string, so decode (if text) or base64 encode (if binary)
    try:
        message_body = file_content.decode("utf-8")  # works if it's plain text
    except UnicodeDecodeError:
        import base64
        message_body = base64.b64encode(file_content).decode("utf-8")

    print(f"Sending file {filename} as SQS message...")

    send_message_response = sqs_extended_client.send_message(
        QueueUrl=queue_url,
        MessageBody=message_body
    )

    response_code = send_message_response["ResponseMetadata"]["HTTPStatusCode"]
    print(f"{filename} → send_message_response code: {response_code}")
