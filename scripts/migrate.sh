#!/bin/bash
set -e

echo "Running database migrations..."

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

MYSQL_CONTAINER=mysql-container

MYSQL_CMD="docker exec -i $MYSQL_CONTAINER mysql -u${DB_USER} -p${DB_PASSWORD}"

echo "Checking if database exists..."
$MYSQL_CMD -e "CREATE DATABASE IF NOT EXISTS ${DB_NAME};"

echo "Running migration files..."

for file in migrations/*.up.sql; do
  echo "Applying migration: $file"
  cat "$file" | $MYSQL_CMD ${DB_NAME}
done

echo "Migrations completed successfully!"