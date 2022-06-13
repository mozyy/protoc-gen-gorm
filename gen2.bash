#!/bin/bash

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
    --js_out=import_style=commonjs,binary:. \
    --grpc-web_out=import_style=typescript,mode=grpcweb:.\
    ./options/gorm.proto

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
    --js_out=import_style=commonjs,binary:. \
    --grpc-web_out=import_style=typescript,mode=grpcweb:.\
    ./types/types.proto
