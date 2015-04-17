#!/usr/bin/env bash

source /usr/lib/kurz/socket_send.sh
socket_send $1 $2 $3 /tmp/kurz.socket
