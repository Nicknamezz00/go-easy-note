<<<<<<< HEAD
#! /usr/bin/env bash
CURDIR=$(cd "$(dirname "$0")" || exit; pwd)
=======
#!/usr/bin/env bash

CURDIR=$(cd $(dirname $0); pwd)
>>>>>>> 6de7d20470d73c692ebfb2c6bdfbd0670657b54b

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

<<<<<<< HEAD
exec "$CURDIR/bin/user"
=======
# self define
export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_AGENT_HOST="127.0.0.1"
export JAEGER_AGENT_PORT=6831


exec "$CURDIR/bin/user"
>>>>>>> 6de7d20470d73c692ebfb2c6bdfbd0670657b54b
