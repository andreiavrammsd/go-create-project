#!/usr/bin/env bash

if [[ -z "$GOPATH" ]]; then
    echo "GOPATH not defined"
    exit
fi

BIN="$GOPATH/bin/gop"
EXEC="/usr/local/bin/gop"

go build -o $BIN
sudo mv $BIN $EXEC

echo
echo "Installed at $EXEC"
echo "Use: gop project-name"
echo
