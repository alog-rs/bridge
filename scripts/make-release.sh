#!/usr/bin/env bash

set -euo pipefail

log() {
    echo "[${0##*/}]: $1" >&2;
}

fatal() {
    log "<FATAL> $1";
    exit 1;
}

if ! [ -x "$(command -v git)" ]; then
    fatal "This script requires git to be installed"
fi

NEW_VERSION="${NEW_VERSION:?Please specify a valid NEW_VERSION environment variable (EX: v0.0.1)}"
RELEASE_TYPE="${RELEASE_TYPE:?Please specify a valid RELEASE_TYPE environment variable (EX: \'RC\' or \'release\')}"
REPO_PREFIX="${REPO_PREFIX:?Please specify a valid REPO_PREFIX environment variable (EX: registry.digitalocean.com/alog-rs)}"

if [ "${RELEASE_TYPE}" != 'RC' ] && [ "${RELEASE_TYPE}" != 'release' ]; then
    fatal "RELEASE_TYPE must be either 'RC' or 'release' - (got: '${RELEASE_TYPE}')"
fi

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_ROOT="${SCRIPT_DIR}/.."
cd $REPO_ROOT

if [[ ! "${NEW_VERSION}" =~ (^v[[:digit:]]+\.[[:digit:]]+\.[[:digit:]]+)(\-rc)*$ ]]; then
    fatal "NEW_VERSION must be in the form 'v0.0.0(-rc)' - (got: ${NEW_VERSION})"
else
    if [ "${BASH_REMATCH[2]}" == '-rc' ] && [ "${RELEASE_TYPE}" != 'RC' ]; then
        fatal "NEW_RELEASE is a release candidate ('${NEW_VERSION}') but RELEASE_TYPE is not 'RC'"
    fi

    if [[ -z "${BASH_REMATCH[2]}" ]] && [ "${RELEASE_TYPE}" != 'release' ]; then
        fatal "NEW_RELEASE is a release ('${NEW_VERSION}') but RELEASE_TYPE is not 'release'"
    fi
fi

if [[ $(git status -s | wc -l) -gt 0 ]]; then
    fatal "Can't have uncomitted changes"
fi

log "Initializing alog.rs ${RELEASE_TYPE} '${NEW_VERSION}' to '${REPO_PREFIX}'"

# Create release PR
git checkout -b "${RELEASE_TYPE}/${NEW_VERSION}"
git commit --allow-empty -s -S -m "${RELEASE_TYPE}: ${NEW_VERSION}"
git tag -s -a -m "${RELEASE_TYPE}/${NEW_VERSION}" "${NEW_VERSION}"

if [[ -z "${SKIP_IMAGE_PUSH:-}" ]]; then
    # Build and push release images
    log "Building and pushing release images to ${REPO_PREFIX}"

    skaffold config set local-cluster false
    PRODUCTION_ENV=1 skaffold build --default-repo="${REPO_PREFIX}" --tag="${NEW_VERSION}"
    skaffold config unset local-cluster
fi

log "Pushing release manifest to ${RELEASE_TYPE}/${NEW_VERSION}..."

git push --set-upstream origin "${RELEASE_TYPE}/${NEW_VERSION}"
git push --tags

log "Successfully tagged release ${NEW_VERSION}."