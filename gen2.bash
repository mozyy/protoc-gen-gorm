#!/bin/bash

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
        --js_out=import_style=commonjs,binary:. \
    ./options/gorm.proto

protoc \
    --proto_path=./ \
    --go_out=. --go_opt=paths=source_relative \
        --js_out=import_style=commonjs,binary:. \
    ./types/types.proto
