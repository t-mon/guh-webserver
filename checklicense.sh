#!/bin/sh

# check the license headers
results=`licensecheck -r -c '\.(go)$' ./ | grep -v "MIT"`
if ! test -z "$results"; then
  echo " -> License check FAILED! Please check following files:"
  licensecheck -r -c '\.(go)$' ./ | grep -v "MIT"
  exit 1
fi

echo "License check PASSED!"
