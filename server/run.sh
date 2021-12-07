#!/bin/sh

export SERVER_LISTEN_ADDR=":50000"
export SERVER_SERVICE="pet/games/xq"

./server serve
