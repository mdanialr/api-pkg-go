#!/usr/bin/env bash
set -euo pipefail

# All credit to https://github.com/ncruces/go-coverage-report

# Get the script's directory after resolving a possible symlink.
SCRIPT_DIR="$(dirname -- "$(readlink -f "${BASH_SOURCE[0]}")")"

OUT_DIR="${1-$SCRIPT_DIR}"
OUT_FILE="$(mktemp)"

# Get coverage for all packages in the current directory; store next to script.
go test ./... -coverprofile "$OUT_FILE"

# Extract total coverage: the decimal number from the last line of the function report.
COVERAGE=$(go tool cover -func="$OUT_FILE" | tail -1 | grep -Eo '[0-9]+\.[0-9]')

echo "coverage: $COVERAGE% of statements"

date "+%s,$COVERAGE" >> "$OUT_DIR/coverage.log"
sort -u -o "$OUT_DIR/coverage.log" "$OUT_DIR/coverage.log"

# Pick a color for the badge.
if awk "BEGIN {exit !($COVERAGE >= 90)}"; then
	COLOR=lime
elif awk "BEGIN {exit !($COVERAGE >= 80)}"; then
	COLOR=limegreen
elif awk "BEGIN {exit !($COVERAGE >= 70)}"; then
	COLOR=green
elif awk "BEGIN {exit !($COVERAGE >= 60)}"; then
	COLOR=yellow
elif awk "BEGIN {exit !($COVERAGE >= 50)}"; then
	COLOR=orange
else
	COLOR=red
fi

# Download the badge; store next to script.
curl -s "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > "$OUT_DIR/coverage.svg"

# When running as a pre-commit hook, add the report and badge to the commit.
if [[ -n "${GIT_INDEX_FILE-}" ]]; then
	git add "$OUT_DIR/coverage.html" "$OUT_DIR/coverage.svg"
fi
