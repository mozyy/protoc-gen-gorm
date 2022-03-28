#!/bin/bash

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
    --js_out=import_style=commonjs,binary:. \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:.\
    ./options/gorm.proto

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
    --js_out=import_style=commonjs,binary:. \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:.\
    ./types/types.proto
