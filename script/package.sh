#!/bin/bash
#! This script package binary file to deb or rpm package.
cd "$( dirname "${BASH_SOURCE[0]}" )/.."

version=`grep 'version' gosensor.go  | sed 's/^[[:space:]]*version = //' | sed 's/"//g'`
echo "creating gosensor package version $version"

rm -r build/dist
mkdir -p build/dist

package() {
	local dirname
	local binary_path
	local configfile_path
	local fpmarch
	local target
	local arch

	target=$1  # deb or rpm
	arch=$2    # 32 or 64

	if [ $target == "deb" ]; then
		if [ $arch == "32" ]; then
			fpmarch="i386"
		elif [ $arch == "64" ]; then
			fpmarch="amd64"
		else
			echo "Use only 64 and 32"
			exit 2
		fi
	elif [ $target == "rpm" ]; then
		if [ $arch == "32" ]; then
			fpmarch="i386"
		elif [ $arch == "64" ]; then
			fpmarch="x86_64"
		else
			echo "Use only 64 and 32"
			exit 2
		fi
	else
		echo "Use only deb and rpm"
		exit 1
	fi

	echo "target:${target}, fpmarch:${fpmarch}"

	binary_path=build/linux${arch}-$version/gosensor
	configfile_path=build/linux${arch}-$version/monitor.json
	echo $binary_path
	# stat $binary_path
	echo $configfile_path
	# stat $configfile_path

	if [ $target == "deb" ]; then
		fpm  -a $fpmarch -s dir -t $target -n "gosensor" -v $version \
		${binary_path}=/usr/local/bin/gosensor ${configfile_path}=/etc/gosensor/monitor.json
	else
		fpm  -a $fpmarch -s dir -t $target -n "gosensor" -v $version --epoch 0 \
		${binary_path}=/usr/local/bin/gosensor ${configfile_path}=/etc/gosensor/monitor.json
	fi

	mv *.$target build/dist
}

package deb 32
package deb 64
package rpm 32
package rpm 64