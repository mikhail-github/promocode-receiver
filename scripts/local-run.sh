#!/bin/bash

set -euo pipefail

EVENT='
{
    "Records":[
        {
            "body":"[{\"shop_id\": \"reebok\", \"type\": \"20%\", \"data\": \"Q4NZ-CGFL-6KKH-7WZ1\" } ]"
        }
    ]
}
'

FUNCTION_NAME="receiver"
DYNAMODB_TABLE="telegram-notifier-lambda-Table-UTE0466VLHR8"
DYNAMODB_PREFIX="stage-"
SENDER_QUEUE_URL="https://sqs.eu-west-1.amazonaws.com/869607576501/telegram-sender20201104094935267000000002"

AWS_ACCESS_KEY_ID=""
AWS_SECRET_ACCESS_KEY=""
AWS_REGION="eu-west-1"


GOOS=linux GOARCH=amd64 \
go build -o "${FUNCTION_NAME}" .

docker run --rm \
    -v "$PWD":/var/task:ro,delegated \
    -v "$PWD/tmp":/tmp:rw \
    -e AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}" \
    -e AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}" \
    -e AWS_REGION="${AWS_REGION}" \
    -e LOG_LEVEL="debug" \
    -e DYNAMODB_TABLE="${DYNAMODB_TABLE}" \
    -e DYNAMODB_PREFIX="${DYNAMODB_PREFIX}" \
    -e SENDER_QUEUE_URL="${SENDER_QUEUE_URL}" \
    lambci/lambda:go1.x \
    "${FUNCTION_NAME}" "${EVENT}"

rm "${FUNCTION_NAME}"