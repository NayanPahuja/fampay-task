#!/bin/sh
set -e


# Run migrations
make migrate-up

# Start the application
exec ./bin/backend