#!/bin/bash

# required versions
REQUIRED_GO="1.24.3"
REQUIRED_SQLITE="3.49.2"

# check if version is greater than or equal to
version_ge() {
  [ "$(printf '%s\n' "$@" | sort -V | head -n1)" != "$1" ]
}

# check each requirement
check_version() {
  tool=$1
  required=$2
  current=$3

  echo -n "Checking $tool... "
  if version_ge "$required" "$current"; then
    echo "Installed: $current (FAIL — requires $required+)"
    return 1
  else
    echo "Installed: $current (OK)"
    return 0
  fi
}

# check for go
if command -v go >/dev/null 2>&1; then
  go_version=$(go version | awk '{print $3}' | sed 's/go//')
else
  echo "Go is not installed."
  go_version="0"
fi

# check for sqlite3
if command -v sqlite3 >/dev/null 2>&1; then
  sqlite_version=$(sqlite3 --version | awk '{print $1}')
else
  echo "SQLite3 is not installed."
  sqlite_version="0"
fi

echo

# set fail flag
fail=0

# update fail flag if any requirement is not met
check_version "Go" "$REQUIRED_GO" "$go_version" || fail=1
check_version "SQLite3" "$REQUIRED_SQLITE" "$sqlite_version" || fail=1

if [ "$fail" -eq 1 ]; then
  echo -e "\nOne or more dependencies are missing or outdated."
  exit 1
else
  echo -e "\nAll dependencies met. You're good to go!"
  exit 0
fi
