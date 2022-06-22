#!/bin/bash

CMDNAME=`basename $0`
script_dir="$( dirname -- "$( readlink -f -- "$0"; )"; )"

cd $script_dir

while getopts fv: OPT
do
  case $OPT in
    "f" ) FLG_FORCE="TRUE" ;;
    "v" ) FLG_VERSION="TRUE" ; VERSION="$OPTARG" ;;
      * ) echo "Usage: $CMDNAME [-f] [-v VALUE]" 1>&2
          exit 1 ;;
  esac
done

if [ "$FLG_VERSION" = "TRUE" ]; then
  :
else
  VERSION="2.2.1"
fi

if [ "$FLG_FORCE" != "TRUE" ]; then
  FILE=./VERSION
  if test -f "$FILE"; then
      line=`cat ${FILE}`
      if [ $line = $VERSION ]; then
        exit 0
      fi
  fi
fi

echo "Download latest binaries..."

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  url="https://github.com/shinolab/autd3/releases/download/v${VERSION}/autd3-v${VERSION}-linux-x64.tar.gz"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  url="https://github.com/shinolab/autd3/releases/download/v${VERSION}/autd3-v${VERSION}-macos-universal.tar.gz"
fi

curl -L -o "tmp.tar.gz" $url

tar xf "tmp.tar.gz"

mv -f ./bin/* ./
rm -rf ./bin
rm -rf ./firmware
rm -rf ./include
rm -rf ./lib
rm -rf tmp.tar.gz

echo $VERSION > ./VERSION

cd - &> /dev/null

echo "done"

exit 0
