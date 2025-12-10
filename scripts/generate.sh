#!/bin/bash

protoc \
  --go_out=./api/gen/v1 \
  --go_opt=paths=source_relative \
  --go-grpc_out=./api/gen/v1 \
  --go-grpc_opt=paths=source_relative \
  ./api/proto/v1/*.proto