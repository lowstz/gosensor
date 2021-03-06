#!/bin/bash
# this build script get from https://github.com/cyfdecyf/cow/blob/master/script/build.sh
cd "$( dirname "${BASH_SOURCE[0]}" )/.."

version=`grep 'version' gosensor.go  | sed 's/^[[:space:]]*version = //' | sed 's/"//g'`
echo "creating gosensor binary version $version"

mkdir -p build
build() {
    local dirname
    local goos
    local goarch
    local goarm
    local cgo
    local armv

    goos="GOOS=$1"
    goarch="GOARCH=$2"
    arch=$3
    if [[ $2 == "arm" ]]; then
        armv=`echo $arch | grep -o [0-9]`
        goarm="GOARM=$armv"
    fi

    if [[ $1 == "darwin" ]]; then
        # Enable CGO for OS X so change network location will not cause problem.
        cgo="CGO_ENABLED=1"
    else
        cgo="CGO_ENABLED=0"
    fi

    dirname=build/$arch-$version
    echo $dirname
    mkdir -p $dirname
    echo "building $dirname"
    echo $cgo $goos $goarch $goarm go build
    eval $cgo $goos $goarch $goarm go build || exit 1

    mv gosensor $dirname/
    cp monitor.json.example $dirname/monitor.json
}

build darwin amd64 mac64
build darwin 386 mac32
build linux amd64 linux64
build linux 386 linux32
build linux arm linux-armv5tel
build linux arm linux-armv6l
build linux arm linux-armv7l