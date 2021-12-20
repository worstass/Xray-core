#!/usr/bin/env bash

protoc -I ./app/extra --go_out=. config.proto
protoc --go_out=.  proxy/tun/config.proto
