#!/bin/sh
rm -r bin/

for i in $(find . -maxdepth 1 -type d -name "[!.]*")
do
	cd $i
	go build -o ../bin/$i -ldflags '-s -w'
	cd ..
done
