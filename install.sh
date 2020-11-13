#!/usr/bin/env bash

if [ "$EUID" -ne 0 ]
    then echo "The script must be run as root."
    exit
fi

mkdir -p /opt/ootbat/

cp * /opt/ootbat/
cp ootbat.service /lib/systemd/system/

chown user:user -R /opt/ootbat/

systemctl daemon-reload
systemctl restart ootbat
systemctl enable ootbat
