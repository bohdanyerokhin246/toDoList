#!/usr/bin/env sh

set -e

if [ -f .env ]; then
  # shellcheck disable=SC2046
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found!"
  exit 1
fi

# shellcheck disable=SC2039
if [[ -z "$DB_USER" || -z "$DB_PASSWORD" || -z "$DB_HOST_APP" || -z "$DB_PORT" || -z "$DB_NAME" || -z "$SSL_MODE" ]]; then
  echo "One or more environment variables are missing!"
  exit 1
fi

MIGRATIONS_PATH=${MIGRATIONS_PATH}

echo "Waiting for 10 seconds before starting migration..."
sleep 10

DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST_APP}:${DB_PORT}/${DB_NAME}?sslmode=${SSL_MODE}"


migrate -path "$MIGRATIONS_PATH" -database "$DATABASE_URL" up
