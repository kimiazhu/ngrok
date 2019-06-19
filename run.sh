#!/usr/bin/env bash
# Copyright 2019 ZHU HAIHUA <kimiazhu@gmail.com>.
# All rights reserved.
# Use of this source code is governed by MIT
# license that can be found in the LICENSE file.
#
# Description: run
# Author: ZHU HAIHUA <kimiazhu@gmail.com>
# Since: 2019-06-19 15:22

set -o errexit

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
PLAIN='\033[0m'

CWD=$(pwd)

function exits() {
  if [[ -n "$2" ]]; then
    echo $2
  fi
  cd $CWD
  exit $1
}

# 获取参数1的绝对目录位置，请通过反引号获取返回值
function absdir() {
  if [[ -d "$1" ]]; then
    cd $1
    echo `pwd`
	elif [[ -f "$1" ]]; then
		cd `dirname $1`
		echo `pwd`
  else
    exits 1 "no such file or directory: $1"
  fi
}

#================== common function end ==================

_pid=$(pidof ngrokd)
if [[ -n "$_pid" ]]; then
  echo "killing pid: $_pid"
  kill $_pid
  sleep 5
  if [[ -n "$(pidof ngrokd)" ]]; then
    exits 1 "kill failed"
  else
    echo "kill successfully"
  fi
fi
nohup ./ngrokd -domain="ngrok.kchu.me" -httpAddr=":80" > /dev/null 2>&1 &
sleep 1
echo "Done. New pid: `pidof ngrokd`"
exits 0
