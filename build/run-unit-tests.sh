#!/bin/bash

echo " > Running run-unit-tests.sh"
set -e

make deps
make test
make coverage

exit 0