#!/bin/bash -e -x -v

# Build binary
env GOOS=linux GOARCH=amd64 go build -o planets-skill lambda-skill.go

# Zip binary
zip -j planets-skill.zip planets-skill

# Create Role
aws iam create-role --role-name lambda-planets-skill-executor --assume-role-policy-document file://$(pwd)/code/skill/trust-policy.json

# Add simple permissions
aws iam attach-role-policy --role-name lambda-planets-skill-executor --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

# Create function
aws lambda create-function --function-name planet-gopher --runtime go1.x --role ${ARN} --handler planets-skill --zip-file fileb://$(pwd)/code/skill/planets-skill.zip

# Update function
# aws lambda update-function-code --function-name planet-gopher --zip-file fileb://$(pwd)/code/skill/planets-skill.zip
