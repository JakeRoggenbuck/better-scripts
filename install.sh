#!/bin/sh

# Clean out old builds
rm -r build/

for i in $(find . -maxdepth 1 -type d -name "[!.]*")
do
	cd $i
	go build -o ../build/$i -ldflags '-s -w'
	cd ..
done
