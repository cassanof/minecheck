SHELL := /bin/bash
.PHONY: all clean

base-dir := $(shell pwd)

build:
	mkdir -p out; \
		GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o minecheck; \
		mv minecheck $(base-dir)/out/minecheck;

run:
	mkdir -p out; \
		GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o minecheck; \
		mv minecheck $(base-dir)/out/minecheck; \
		$(base-dir)/out/minecheck;


install:
	mkdir -p out; \
		GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o minecheck; \
		mv minecheck $(base-dir)/out/minecheck; \
		chmod -R 777 $(base-dir)/out/; \
		mv $(base-dir)/out/minecheck /usr/bin/minecheck; \
		chown root:root /usr/bin/minecheck; \
		chmod 0755 /usr/bin/minecheck; 

clean:
	rm -fr out;
