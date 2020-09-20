#!/bin/sh
set -e

POSTGRES="psql --username ${POSTGRES_USER} -d ${POSTGRES_DB}"

echo -e "\e[33m$(date +"%c") \e[39mCreating schema..."
psql -d ${POSTGRES_DB} -a -U${POSTGRES_USER} -f /schema.sql

echo -e "\e[33m$(date +"%c") \e[39mCreating tables..."
psql -d ${POSTGRES_DB} -a -U${POSTGRES_USER} -f /tables.sql
