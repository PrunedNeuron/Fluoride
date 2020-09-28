#!/bin/sh
set -e

# set application env to development if unset
if [ -z "${APPLICATION_ENV}" ]; then
	export APPLICATION_ENV=development
fi

# enable buildkit if running docker commands
if [[ "$1" =~ "docker" || "$1" =~ "compose" ]]; then
	echo -e "\e[33m$(date +"%c")\t\e[39mEnabling BuildKit before running docker commands"
	export COMPOSE_DOCKER_CLI_BUILD=1
	export DOCKER_BUILDKIT=1
fi

case "$1" in
"serve" | "run" | "build")
	# set target os and arch for go binary build
	echo -e "\e[33m$(date +"%c")\t\e[39mSetting OS and ARCH env variables using the \e[92m\`go env\`\e[39m command"
	export TARGETOS=$(go env GOOS)
	export TARGETARCH=$(go env GOARCH)
	;;
esac
