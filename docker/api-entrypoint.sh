#!/bin/sh
set -e

CONFIG_ROOT="${HOME:-/root}/.config/fabric"

if [ ! -d "$CONFIG_ROOT" ]; then
	mkdir -p "$CONFIG_ROOT"
fi

if [ ! -f "$CONFIG_ROOT/.env" ]; then
	touch "$CONFIG_ROOT/.env"
fi

exec "$@"
