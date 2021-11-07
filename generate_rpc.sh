#!/usr/bin/env bash

protoc -I ./app/extra --go_out=. config.proto
