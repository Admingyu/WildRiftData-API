#!/bin/bash

echo y | yum install -y golang.x86_64

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
export GOPATH=/root/volume/WildRiftData-API
go env

pwd
go run main.go

