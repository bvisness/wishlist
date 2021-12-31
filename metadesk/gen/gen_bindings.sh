#!/bin/bash

set -euo pipefail

rm -f md.go
mv ext.go ext.go.tmp
go run ./gen/gen.go; mv ext.go.tmp ext.go
