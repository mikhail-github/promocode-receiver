#!/bin/bash

set -euo pipefail

FUNCTION_NAME="receiver"

rm "${FUNCTION_NAME}.zip" "${FUNCTION_NAME}" || true

GOOS=linux GOARCH=amd64 \
go build -o "${FUNCTION_NAME}" .

zip -9yr "${FUNCTION_NAME}.zip" "${FUNCTION_NAME}"
