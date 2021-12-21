#!/usr/bin/env bash

protoc -I ./app/extra --go_out=. config.proto
protoc -I ./app/tun --go_out=. config.proto
