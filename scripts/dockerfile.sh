#!/bin/bash

set -eux

# 二进制文件名
PROJECT="caesar"
# 启动文件PATH
MAIN_PATH="main.go"

# build flag
LDFLAGS="-s -w"
#			   -X "main.BuildVersion=${VERSION}"
#			   -X "main.BuildDate=$(shell /bin/date "+%F %T")"

function create() {
  #判断文件夹是否存在，不存在则创建
  if [ ! -d "$1" ]; then
    mkdir -p "$1"
  fi
}

function delete() {
  #判断文件夹是否存在，不存在则创建
  if [ -d "$1" ]; then
    rm -r "$1"
  fi
}

function build() {

  delete build/package/agent
  delete build/package/caesar.tar
  create build/package/agent
  create build/package/agent/results
  cp -r assets build/package/agent
  cp configs/config.yml build/package/agent
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "${LDFLAGS}" -o build/package/agent/"${PROJECT}" "${MAIN_PATH}"
  cd build/package/
  docker build -t caesar .
  docker save caesar:latest -o caesar.tar

}

build
