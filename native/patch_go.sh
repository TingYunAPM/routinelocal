#!/bin/bash

wget https://storage.googleapis.com/golang/go1.7.6.linux-amd64.tar.gz
wget https://github.com/golang/go/archive/go1.11.9.tar.gz
tar -zxvf go1.7.6.linux-amd64.tar.gz
tar -zxvf go1.11.9.tar.gz
export GOROOT_BOOTSTRAP=$(pwd)/go
cd go-go1.11.9
patch -p1 < ../go1.11.9.patch
#export GOROOT_BOOTSTRAP=$(go env | grep GOROOT | sed "s/GOROOT=//g"| sed "s/\"//g")
cd src
unset GOROOT
unset GOPATH
#if build faild see this:
#   https://github.com/golang/go/issues/27754
#I build pass on both centos7 and ubuntu16.04 
time bash -x ./all.bash

