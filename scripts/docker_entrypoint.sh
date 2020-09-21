#!/bin/sh
set -e

# If the intention is to serve, make sure postgres is ready
# $1 = subcommand to be passed as an arg to the binary
if [ "$#" -eq 1 ] && [ "$1" = "serve" ]; then
	echo -e "\e[33m$(date +"%c")\t\e[39mInitializing database"
	sh scripts/wait_for_postgres.sh
fi

echo -e "\e[33m$(date +"%c")\t\e[92mLaunching application"
bin/"${PWD##*/}" "$@"

exec "$@"
