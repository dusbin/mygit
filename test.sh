#! /bin/bash
# export GOPATH=/go
# export GOROOT=/go  ~/.bashrc文件 默认的GOROOT不用改
echo "begin..."
./mygit/mygit
./mygit/mygit help
./mygit/mygit initrepo
./mygit/mygit dus
./mygit/mygit undefine
echo "end..."
