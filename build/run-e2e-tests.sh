#!/bin/bash
# Copyright (c) 2021 Red Hat, Inc.
# Copyright Contributors to the Open Cluster Management project

echo " > Running run-e2e-test.sh"
set -e
sh tests/e2e/runTests.sh $1

exit 0