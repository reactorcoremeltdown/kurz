#!/usr/bin/env bash

echo -n "$1∙$2∙$3" | socat stdio unix-connect:/tmp/kurz.socket
