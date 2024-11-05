#!/bin/sh

set -e

while ! nc -z "$DB_HOST" "$DB_PORT"; do
  sleep 0.1
done

echo "Database is up - running migrations"

migrate -path=/migrations -database "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up

echo "Migrations completed - starting application"

exec ./main
