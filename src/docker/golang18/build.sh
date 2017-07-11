#!/bin/bash
set -e
echo "[build.sh:building binary]"
cd $BUILDPATH && go build -o /gerbo && cp /gerbo /usr/bin/ && chmod u+x /usr/bin/gerbo && rm -rf /tmp/*
echo "[build.sh:launching binary]"
gerbo