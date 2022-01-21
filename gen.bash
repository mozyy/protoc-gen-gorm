#!/bin/bash

go install .

protoc \
    --proto_path=./ \
    --proto_path=$HOME/.local/include \
    --go_out=. --go_opt=paths=source_relative \
    --gorm_out=. --gorm_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --js_out=import_style=commonjs,binary:. \
    ./example/base/test.proto