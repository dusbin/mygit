#! /bin/bash
# export GOPATH=/go
# export GOROOT=/go  ~/.bashrc文件 默认的GOROOT不用改
echo "begin..."
export GOPATH=/mnt/hgfs/mygitrepo/mygit/mygit/
#export GOROOT=/usr/lib/go-1.6
cd mygit/
go build .
echo "end..."
