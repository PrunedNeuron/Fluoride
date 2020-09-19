#!/bin/bash
set -e

POSTGRES="psql --username ${POSTGRES_USER} -d ${POSTGRES_DB}"

echo "Creating schema..."
psql -d ${POSTGRES_DB} -a -U${POSTGRES_USER} -f /schema.sql

echo "Creating tables..."
psql -d ${POSTGRES_DB} -a -U${POSTGRES_USER} -f /tables.sql