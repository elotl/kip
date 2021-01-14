#!/bin/sh
sudo /usr/local/bin/anka license activate -f $ANKA_LICENSE
sudo /usr/local/bin/anka license accept-eula || true
sudo /usr/local/bin/anka license validate
sudo curl -o /usr/local/bin/itzo https://itzo-kip-dev-download.s3.amazonaws.com/itzo-darwin-820
sudo chmod +x /usr/local/bin/itzo
sudo /usr/local/bin/itzo -v=2 > /var/log/itzo.log 2>&1 &


