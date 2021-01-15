#!/bin/bash

anka license activate -f $ANKA_LICENSE
anka license accept-eula || true
anka license validate

wget -O /usr/local/bin/itzo https://itzo-kip-dev-download.s3.amazonaws.com/itzo-darwin-818
chmod 755 /usr/local/bin/*
/usr/local/bin/itzo -v=2 -use-anka > /var/log/itzo.log 2>&1 &
