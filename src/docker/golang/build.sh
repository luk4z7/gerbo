#!/bin/bash
set -e
echo "[build.sh:building binary]"
cd $BUILDPATH && go build -o /gerbo && rm -rf /tmp/*
echo "[build.sh:launching binary]"
/gerbo