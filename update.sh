#!/bin/bash

function update() {
    echo "updating"
    [ -e ./netbox ] && rm -Rf ./netbox
    [ -e ./go-netbox ] && rm -Rf ./go-netbox

    git clone --depth 1 git@github.com:netbox-community/go-netbox.git || exit 1
    cd go-netbox || exit 1

    echo -n "replacing "
    grep -Rl "\"github.com/netbox-community/" * | while read f; do
        echo -n .
        sed -e 's+"github.com/netbox-community/go-netbox+"github.com/nzions/not-netbox+g' $f >/tmp/nb.$$
        cp /tmp/nb.$$ $f
    done
    echo ""

    [ -e /tmp/nb.$$ ] && rm /tmp/nb.$$

    cp -R netbox ../netbox
    git merge mods
    go mod tidy

    [ -e ./go-netbox ] && rm -Rf ./go-netbox
}

case $1 in
*)
    update
    ;;
esac
