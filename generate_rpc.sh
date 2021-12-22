#!/usr/bin/env bash

protoc --go_out=. app/extra/config.proto
protoc --go_out=. app/tun/config.proto