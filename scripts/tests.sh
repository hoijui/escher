#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2021 Robin Vobruba <hoijui.quaero@gmail.com>
#
# SPDX-License-Identifier: Unlicense

# Runs all the go and Escher (unit-)tests.
# NOTE
# * Requires the `escher` command available on the PATH.

# Exit immediately on each error and unset variable;
# see: https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail/
set -Eeuo pipefail
#set -Eeu

script_dir=$(dirname "$(readlink -f "${BASH_SOURCE[0]}")")
repo_root="$(cd "$script_dir"; cd ..; pwd)"
# NOTE We do not use this path,
#      even though it would make the script position independent,
#      because it would break (or worse: run the wrong code)
#      when working on a fork of the repository.
#src_dir="$GOPATH/src/github.com/hoijui/escher/src/"
# This way of defning src_dir ensures that we can use relative paths,
# while the script may still be called from anywhere,
# as long as the sources are to be found
# under the same relative path within the escher repo.
src_dir="$repo_root/src"

if ! which escher > /dev/null
then
	>&2 echo "Error: Could not find 'escher' in PATH"
	exit 1
fi

echo
echo "Running Go(lang) tests ..."
cd "$repo_root"
find . -name "*_test.go" \
	| while read -r go_test
do
	test_dir=$(dirname "$go_test")
	cd "$test_dir"

	echo
	echo "GO TESTS $go_test ..."
	go test
	cd "$repo_root"
done

echo
if [ "${CI:-NOT_SET}" = "NOT_SET" ]
then
	echo "Running Escher tests ..."
	ESCHER=$src_dir escher "*test.All"
else
	>&2 echo "WARNING: Not running Escher tests on CI, as there is a mysterious bug" # HACK
	>&2 echo "WARNING: that makes them fail only there; of which the output on the CI Job is:"
	>&2 printf "WARNING: \"end of session\\nunexpected EOF\""
fi

echo
echo "done."
