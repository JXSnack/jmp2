#!/usr/bin/env bash

set -e # stop after first error

# ensure sudo usage, in order to move file without issues
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

go build -buildvcs=false -o jp2-bin # build
mv jp2-bin /usr/local/bin # move
chmod a+x /usr/local/bin/jp2-bin # grant usage permissions

echo "Done"
