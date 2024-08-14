#!/bin/bash
set -e

retry() {
    local cmd="$1"
    local host="$2"
    local port="$3"
    local retries=15
    local wait=5

    until nc -z "$host" "$port" > /dev/null 2>&1; do
        retries=$((retries-1))
        if [ $retries -le 0 ]; then
            echo "Service $host:$port is not available after multiple attempts, exiting."
            exit 1
        fi
        echo "Waiting for $host:$port to be available..."
        sleep $wait
    done

    echo "Service $host:$port is available, running command: $cmd"
    exec $cmd
}

# Choose command based on input arguments
if [ "$1" == "/simulator" ]; then
    retry "/simulator" "$RABBITMQ_HOST" "$RABBITMQ_PORT"
elif [ "$1" == "/processor" ]; then
    retry "/processor" "$RABBITMQ_HOST" "$RABBITMQ_PORT"
else
    exec "$@"
fi
