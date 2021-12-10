#!/bin/sh

. ./conf.sh

./xq migrate && ./xq serve
