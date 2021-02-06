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
EVENT='
{
    "Records":[
        {
            "body":"[{\"shop_id\": \"adidas\", \"type\": \"20%\", \"data\": \"U20-7W6W-M9GT-3MS4-MLN11\" } ]"
        }
    ]
}
'

FUNCTION_NAME="receiver"
DYNAMODB_TABLE="telegram-notifier-lambda-Table-3J2VOEM2OKB3"
DYNAMODB_PREFIX="stage-"
SENDER_QUEUE_URL="https://sqs.eu-west-1.amazonaws.com/869607576501/telegram-sender20210121145307172200000008"
VKPOSTER_LAMBDA_NAME="promocode-vkposter"
ADIDAS_REFLINK="https://fas.st/C74PQ"
REEBOK_REFLINK="https://fas.st/OsaQV"

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
    -e VKPOSTER_LAMBDA_NAME="${VKPOSTER_LAMBDA_NAME}" \
    -e ADIDAS_REFLINK="${ADIDAS_REFLINK}" \
    -e REEBOK_REFLINK="${REEBOK_REFLINK}" \
    lambci/lambda:go1.x \
    "${FUNCTION_NAME}" "${EVENT}"

rm "${FUNCTION_NAME}"
