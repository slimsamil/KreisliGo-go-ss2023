#!/bin/sh

set -e

# Wait for DB
if [ -n "$DB_CONNECT" ]; then
    /go/src/app/wait-for-it.sh "$DB_CONNECT" -t 20
fi

# Wait for banktransfer
if [ -n "$BANKTRANSFER_CONNECT" ]; then
  /go/src/app/wait-for-it.sh "$BANKTRANSFER_CONNECT" -t 20
fi

# Run the main container command.
exec "$@"
