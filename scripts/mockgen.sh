#!/usr/bin/env bash

set -euo pipefail

if ! [ -x "$(command -v mockgen)" ]; then
    echo "mockgen.sh requires mockgen";

    exit 1;
fi

MOCK_PKG=mocks
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR=$SCRIPT_DIR/..
DEST_DIR=$ROOT_DIR/internal/mocks

# Service mocks
mockgen -source=$ROOT_DIR/service/rs3.go -package=$MOCK_PKG -destination $DEST_DIR/rs3_service_mock.go

# Helper mocks
mockgen -source=$ROOT_DIR/internal/helpers/request.go -package=$MOCK_PKG -destination $DEST_DIR/request_mock.go
