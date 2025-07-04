#!/usr/bin/env bash

set -e # stop after first error

# ensure sudo usage, in order to move file without issues
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

go build -buildvcs=false -o jp2 # build
mv jp2 /usr/local/bin # move
chmod a+x /usr/local/bin/jp2 # grant usage permissions

echo "Done"
