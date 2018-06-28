#!/bin/bash

# Build binary
env GOOS=linux GOARCH=amd64 go build -o planets main.go

# Zip binary
zip -j planets.zip planets

# Create Role
aws iam create-role --role-name lambda-planets-executor --assume-role-policy-document file://$(pwd)/trust-policy.json

# Add simple permissions
aws iam attach-role-policy --role-name lambda-planets-executor --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

# Create function
aws lambda create-function --function-name planets --runtime go1.x --role ${ARN} --handler planets --zip-file fileb://$(pwd)/planets.zip

# Test the Lambda
aws lambda invoke --function-name planets --payload '{"planet": "sun"}' /tmp/planets-output.json

aws lambda invoke --function-name planets --payload '{"planet": ""}' /tmp/planets-output.json

aws lambda invoke --function-name planets /tmp/planets-output.json
