#!/bin/sh
# run: Wrapper for Makefile with some env vars set

if [[ "$1" != "help" && "$1" != *docs* ]]; then
	# run script to set up basic env variables
	# run the script in the current shell to preserve env
	# since it is run in the same shell, it has access to the args
	. ./scripts/setup_env.sh
fi

# run Makefile with the passed target args
make "$@"
