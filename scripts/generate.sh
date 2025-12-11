#!/bin/bash

protoc \
  --go_out=. \
  --go_opt=module=github.com/forzeyy/messenger-media-service \
  --go-grpc_out=. \
  --go-grpc_opt=module=github.com/forzeyy/messenger-media-service \
  ./api/proto/v1/*.proto