#!/bin/sh
set -e

cmd="$@"
timer="1"
seconds_passed=1

echo -e "\e[33m$(date +"%c") \e[39mWaiting for Postgres"

# until pg_isready 1>/dev/null 2>&1; do
until nc -z postgres 5432 1>/dev/null 2>&1; do
	if [ $(expr $seconds_passed % 5) == 0 ]; then
		echo -e "\e[33m$(date +"%c")\t\e[39mPostgres is unavailable, waiting" >&2
	fi

	sleep $timer
	seconds_passed=$((seconds_passed + 1))
done

sleep 5 # To make sure the database is up by the time we exit this script

echo -e "\e[33m$(date +"%c")\t\e[92mPostgres is ready\e[39m"
