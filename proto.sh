#!/bin/bash

protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src `pwd`/proto/*.proto
echo Proto done.
