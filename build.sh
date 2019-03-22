#!/bin/bash -x
set -e


OUT_DIR=./pb

rm -rf $OUT_DIR/*.{go,json}

protoc -I pb -I /usr/local/include -I $GOPATH/src --go_out=plugins=grpc:$OUT_DIR pb/*.proto
