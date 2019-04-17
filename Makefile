# Copyright © 2019 Sonoko Mizuki, Ltd. All rights reserved.

all: build

build:
	go build -o build/app

test: 
	go test -v ./...

clean:
	-rm build/*

.PHONY: build test clean

